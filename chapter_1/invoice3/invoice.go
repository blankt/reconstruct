package invoice3

import (
	"fmt"
	"reconstruct/chapter_1/invoice1"
)

func plainTextStatement(invoice Invoice, plays map[string]Play) string {
	data := createStatementData(invoice, plays)
	return RenderPlainText(data)
}

func RenderPlainText(data statementData) string {
	result := fmt.Sprintf("Statement for %s\n", data.Customer)
	for _, perf := range data.Performances {
		// print line for this order
		result += fmt.Sprintf(" %s: %s (%d seats)\n", perf.Play.Name, invoice1.Usd(perf.Amount), perf.Audience)
	}

	result += fmt.Sprintf("Amount owed is %s\n", invoice1.Usd(data.TotalAmount))
	result += fmt.Sprintf("You earned %d credits\n", data.TotalVolumeCredits)
	return result
}

func htmlStatement(invoice Invoice, plays map[string]Play) string {
	data := createStatementData(invoice, plays)
	return RenderHtml(data)
}

func RenderHtml(data statementData) string {
	result := fmt.Sprintf("<h1>Statement for %s</h1>\n", data.Customer)
	result += "<table>\n"
	result += "<tr><th>play</th><th>seats</th><th>cost</th></tr>"
	for _, perf := range data.Performances {
		result += fmt.Sprintf(" <tr><td>%s</td><td>%d</td>", perf.Play.Name, perf.Audience)
		result += fmt.Sprintf("<td>%s</td></tr>\n", invoice1.Usd(perf.Amount))
	}
	result += "</table>\n"
	result += fmt.Sprintf("<p>Amount owed is <em>%s</em></p>\n", invoice1.Usd(data.TotalAmount))
	result += fmt.Sprintf("<p>You earned <em>%d</em> credits</p>\n", data.TotalVolumeCredits)
	return result
}
