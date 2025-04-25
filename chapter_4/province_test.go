package chapter_4

import "testing"

func TestProvince(t *testing.T) {
	dataMap := make(map[string]interface{})
	dataMap["name"] = "Asia"
	dataMap["producers"] = []map[string]interface{}{
		{
			"name":       "Byzantium",
			"cost":       10,
			"production": 9,
		},
		{
			"name":       "Attalia",
			"cost":       12,
			"production": 10,
		},
		{
			"name":       "Sinope",
			"cost":       10,
			"production": 6,
		},
	}
	dataMap["demand"] = 30
	dataMap["price"] = 20

	t.Run("GetShortfall", func(t *testing.T) {
		province := NewProvince(dataMap)
		if province.GetShortfall() != 5 {
			t.Errorf("GetShortfall() = %d; want 5", province.GetShortfall())
		}
	})

	t.Run("GetProfit", func(t *testing.T) {
		province := NewProvince(dataMap)
		if province.GetProfit() != 230 {
			t.Errorf("GetProfit() = %d; want 10", province.GetProfit())
		}
	})

	t.Run("ZeroDemand", func(t *testing.T) {
		dataMap := make(map[string]interface{})
		dataMap["name"] = "NoDemand"
		dataMap["producers"] = []map[string]interface{}{
			{
				"name":       "Producer1",
				"cost":       10,
				"production": 10,
			},
		}
		dataMap["demand"] = 0
		dataMap["price"] = 20

		province := NewProvince(dataMap)
		if province.GetShortfall() != -10 {
			t.Errorf("GetShortfall() = %d; want -10", province.GetShortfall())
		}
		if province.GetProfit() != 0 {
			t.Errorf("GetProfit() = %d; want 0", province.GetProfit())
		}
	})

	t.Run("ZeroProducers", func(t *testing.T) {
		dataMap := make(map[string]interface{})
		dataMap["name"] = "NoProducers"
		dataMap["producers"] = []map[string]interface{}{}
		dataMap["demand"] = 30
		dataMap["price"] = 20

		province := NewProvince(dataMap)
		if province.GetShortfall() != 30 {
			t.Errorf("GetShortfall() = %d; want 30", province.GetShortfall())
		}
		if province.GetProfit() != 0 {
			t.Errorf("GetProfit() = %d; want 0", province.GetProfit())
		}
	})

	t.Run("NegativeDemand", func(t *testing.T) {
		dataMap := make(map[string]interface{})
		dataMap["name"] = "NegativeDemand"
		dataMap["producers"] = []map[string]interface{}{
			{
				"name":       "Producer1",
				"cost":       10,
				"production": 10,
			},
		}
		dataMap["demand"] = -5
		dataMap["price"] = 20

		province := NewProvince(dataMap)
		if province.GetShortfall() != -15 {
			t.Errorf("GetShortfall() = %d; want -15", province.GetShortfall())
		}
		if province.GetProfit() != -50 {
			t.Errorf("GetProfit() = %d; want -50", province.GetProfit())
		}
	})
}
