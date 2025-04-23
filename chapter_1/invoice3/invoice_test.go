package invoice3

import (
	"testing"
)

func TestStatement(t *testing.T) {
	t.Run("ReturnsCorrectOutputForValidInvoice", statementReturnsCorrectOutputForValidInvoice)
	t.Run("HandlesUnknownPlayIDGracefully", statementHandlesUnknownPlayIDGracefully)
	t.Run("HandlesUnknownPlayTypeGracefully", statementHandlesUnknownPlayTypeGracefully)
	t.Run("HandlesEmptyPerformances", statementHandlesEmptyPerformances)
	t.Run("HandlesZeroAudience", statementHandlesZeroAudience)
	t.Run("HandlesMultiplePerformancesWithZeroAudience", statementHandlesMultiplePerformancesWithZeroAudience)
	t.Run("HandlesLargeAudienceNumbers", statementHandlesLargeAudienceNumbers)
	t.Run("HandlesSinglePerformance", statementHandlesSinglePerformance)
}

func statementReturnsCorrectOutputForValidInvoice(t *testing.T) {
	plays := map[string]Play{
		"hamlet":  {Name: "Hamlet", Type: "tragedy"},
		"as-like": {Name: "As You Like It", Type: "comedy"},
		"othello": {Name: "Othello", Type: "tragedy"},
	}

	invoice := Invoice{
		Customer: "BigCo",
		Performances: []Performance{
			{PlayID: "hamlet", Audience: 55},
			{PlayID: "as-like", Audience: 35},
			{PlayID: "othello", Audience: 40},
		},
	}

	expected := `Statement for BigCo
 Hamlet: $650.00 (55 seats)
 As You Like It: $580.00 (35 seats)
 Othello: $500.00 (40 seats)
Amount owed is $1730.00
You earned 47 credits
`
	result := plainTextStatement(invoice, plays)
	if result != expected {
		t.Errorf("Expected:\n%s\nGot:\n%s", expected, result)
	}
}

func statementHandlesUnknownPlayIDGracefully(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic for unknown play ID, but no panic occurred")
		}
	}()

	plays := map[string]Play{
		"hamlet": {Name: "Hamlet", Type: "tragedy"},
	}

	invoice := Invoice{
		Customer: "BigCo",
		Performances: []Performance{
			{PlayID: "unknown", Audience: 55},
		},
	}

	plainTextStatement(invoice, plays)
}

func statementHandlesUnknownPlayTypeGracefully(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic for unknown play type, but no panic occurred")
		}
	}()

	plays := map[string]Play{
		"hamlet": {Name: "Hamlet", Type: "unknown"},
	}

	invoice := Invoice{
		Customer: "BigCo",
		Performances: []Performance{
			{PlayID: "hamlet", Audience: 55},
		},
	}

	plainTextStatement(invoice, plays)
}

func statementHandlesEmptyPerformances(t *testing.T) {
	plays := map[string]Play{
		"hamlet": {Name: "Hamlet", Type: "tragedy"},
	}

	invoice := Invoice{
		Customer:     "BigCo",
		Performances: []Performance{},
	}

	expected := `Statement for BigCo
Amount owed is $0.00
You earned 0 credits
`

	result := plainTextStatement(invoice, plays)
	if result != expected {
		t.Errorf("Expected:\n%s\nGot:\n%s", expected, result)
	}
}

func statementHandlesZeroAudience(t *testing.T) {
	plays := map[string]Play{
		"hamlet": {Name: "Hamlet", Type: "tragedy"},
	}

	invoice := Invoice{
		Customer: "BigCo",
		Performances: []Performance{
			{PlayID: "hamlet", Audience: 0},
		},
	}

	expected := `Statement for BigCo
 Hamlet: $400.00 (0 seats)
Amount owed is $400.00
You earned 0 credits
`

	result := plainTextStatement(invoice, plays)
	if result != expected {
		t.Errorf("Expected:\n%s\nGot:\n%s", expected, result)
	}
}

func statementHandlesMultiplePerformancesWithZeroAudience(t *testing.T) {
	plays := map[string]Play{
		"hamlet":  {Name: "Hamlet", Type: "tragedy"},
		"as-like": {Name: "As You Like It", Type: "comedy"},
	}

	invoice := Invoice{
		Customer: "BigCo",
		Performances: []Performance{
			{PlayID: "hamlet", Audience: 0},
			{PlayID: "as-like", Audience: 0},
		},
	}

	expected := `Statement for BigCo
 Hamlet: $400.00 (0 seats)
 As You Like It: $300.00 (0 seats)
Amount owed is $700.00
You earned 0 credits
`

	result := plainTextStatement(invoice, plays)
	if result != expected {
		t.Errorf("Expected:\n%s\nGot:\n%s", expected, result)
	}
}

func statementHandlesLargeAudienceNumbers(t *testing.T) {
	plays := map[string]Play{
		"hamlet":  {Name: "Hamlet", Type: "tragedy"},
		"as-like": {Name: "As You Like It", Type: "comedy"},
	}

	invoice := Invoice{
		Customer: "BigCo",
		Performances: []Performance{
			{PlayID: "hamlet", Audience: 1000},
			{PlayID: "as-like", Audience: 1000},
		},
	}

	expected := `Statement for BigCo
 Hamlet: $10100.00 (1000 seats)
 As You Like It: $8300.00 (1000 seats)
Amount owed is $18400.00
You earned 2140 credits
`

	result := plainTextStatement(invoice, plays)
	if result != expected {
		t.Errorf("Expected:\n%s\nGot:\n%s", expected, result)
	}
}

func statementHandlesSinglePerformance(t *testing.T) {
	plays := map[string]Play{
		"hamlet": {Name: "Hamlet", Type: "tragedy"},
	}

	invoice := Invoice{
		Customer: "BigCo",
		Performances: []Performance{
			{PlayID: "hamlet", Audience: 50},
		},
	}

	expected := `Statement for BigCo
 Hamlet: $600.00 (50 seats)
Amount owed is $600.00
You earned 20 credits
`

	result := plainTextStatement(invoice, plays)
	if result != expected {
		t.Errorf("Expected:\n%s\nGot:\n%s", expected, result)
	}
}
