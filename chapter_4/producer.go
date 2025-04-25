package chapter_4

import (
	"math"
	"strconv"
)

type Producer struct {
	Province   *Province
	Name       string
	Cost       int
	Production int
}

func NewProducer(province *Province, data map[string]interface{}) *Producer {
	return &Producer{
		Province:   province,
		Name:       data["name"].(string),
		Cost:       data["cost"].(int),
		Production: data["production"].(int),
	}
}

func (p *Producer) SetCost(cost string) {
	parsedCost, _ := strconv.Atoi(cost)
	p.Cost = parsedCost
}

func (p *Producer) SetProduction(amountStr string) {
	amount, _ := strconv.Atoi(amountStr)
	newProduction := 0
	if !math.IsNaN(float64(amount)) {
		newProduction = amount
	}

	p.Province.TotalProduction += newProduction - p.Production
	p.Production = newProduction
}
