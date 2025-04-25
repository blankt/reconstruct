package chapter_4

import (
	"math"
	"sort"
)

type Province struct {
	Name            string
	Producers       []*Producer
	TotalProduction int
	Demand          int
	Price           int
}

func NewProvince(doc map[string]interface{}) *Province {
	province := &Province{
		Name:            doc["name"].(string),
		Producers:       []*Producer{},
		TotalProduction: 0,
		Demand:          doc["demand"].(int),
		Price:           doc["price"].(int),
	}

	for _, p := range doc["producers"].([]map[string]interface{}) {
		province.AddProducer(NewProducer(province, p))
	}

	return province
}

func (p *Province) AddProducer(producer *Producer) {
	p.Producers = append(p.Producers, producer)
	p.TotalProduction += producer.Production
}

func (p *Province) GetShortfall() int {
	return p.Demand - p.TotalProduction
}

func (p *Province) GetProfit() int {
	return p.GetDemandValue() - p.GetDemandCost()
}

func (p *Province) GetDemandCost() int {
	remainingDemand := p.Demand
	result := 0

	sort.Slice(p.Producers, func(i, j int) bool {
		return p.Producers[i].Cost < p.Producers[j].Cost
	})

	for _, producer := range p.Producers {
		contribution := int(math.Min(float64(remainingDemand), float64(producer.Production)))
		remainingDemand -= contribution
		result += contribution * producer.Cost
	}

	return result
}

func (p *Province) GetDemandValue() int {
	return p.GetSatisfiedDemand() * p.Price
}

func (p *Province) GetSatisfiedDemand() int {
	return int(math.Min(float64(p.Demand), float64(p.TotalProduction)))
}
