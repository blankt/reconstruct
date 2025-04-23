package invoice3

import (
	"fmt"
	"math"
)

type Play struct {
	Name string
	Type string
}

type Performance struct {
	PlayID   string
	Audience int
}

type Invoice struct {
	Customer     string
	Performances []Performance
}

type statementData struct {
	Customer           string
	Performances       []MiddlePerformance
	TotalAmount        int
	TotalVolumeCredits int
}

type MiddlePerformance struct {
	Performance
	Play    Play
	Amount  int
	Credits int
}

/*
第三阶段：按照戏剧类型重组计算过程 引入多态来处理计算过程
戏剧类型在计算分支的选择上起着重要的作用，但是这样的分支选择会随着代码的堆积而腐坏，且不容易增加新的戏剧类型。

所做优化：
1.以多态取代条件表达式（272）：不同类型的戏剧的计算各子集中到一个地方，对大多数只是修改特定的计算方式很有意义。如果增加新的戏剧类型，只需要增加一个新的计算类，而不需要修改原有的代码。
*/
func createStatementData(invoice Invoice, plays map[string]Play) statementData {
	data := statementData{}
	data.Customer = invoice.Customer
	data.Performances = enrichPerformances(invoice.Performances, plays)
	data.TotalAmount = totalAmount(data.Performances)
	data.TotalVolumeCredits = totalVolumeCredits(data.Performances)
	return data
}

func enrichPerformances(performances []Performance, plays map[string]Play) []MiddlePerformance {
	result := make([]MiddlePerformance, len(performances))
	for i, perf := range performances {
		play := PlayFor(perf, plays)
		result[i] = MiddlePerformance{
			Performance: perf,
			Play:        play,
		}
		calculator := CreatePerformanceCalculator(play, perf)
		result[i].Amount = calculator.amount()
		result[i].Credits = calculator.volumeCredits()
	}
	return result
}

func totalVolumeCredits(performances []MiddlePerformance) int {
	total := 0
	for _, perf := range performances {
		total += perf.Credits
	}
	return total
}

func totalAmount(performances []MiddlePerformance) int {
	total := 0
	for _, perf := range performances {
		total += perf.Amount
	}
	return total
}

func PlayFor(performance Performance, plays map[string]Play) Play {
	play, ok := plays[performance.PlayID]
	if !ok {
		panic(fmt.Sprintf("unknown play ID: %s", performance.PlayID))
	}
	return play
}

type PerformanceCalculatorI interface {
	amount() int
	volumeCredits() int
}

func CreatePerformanceCalculator(play Play, perf Performance) PerformanceCalculatorI {
	switch play.Type {
	case "tragedy":
		return TragedyCalculator{Play: play, Perf: perf}
	case "comedy":
		return ComedyCalculator{Play: play, Perf: perf}
	default:
		panic(fmt.Sprintf("unknown type: %s", play.Type))
	}
}

type ComedyCalculator struct {
	Play Play
	Perf Performance
}

func (c ComedyCalculator) amount() int {
	result := 30000
	if c.Perf.Audience > 20 {
		result += 10000 + 500*(c.Perf.Audience-20)
	}
	result += 300 * c.Perf.Audience
	return result
}

func (c ComedyCalculator) volumeCredits() int {
	result := int(math.Max(float64(c.Perf.Audience-30), 0))
	if c.Play.Type == "comedy" {
		result += c.Perf.Audience / 5
	}
	return result
}

type TragedyCalculator struct {
	Play Play
	Perf Performance
}

func (t TragedyCalculator) amount() int {
	result := 40000
	if t.Perf.Audience > 30 {
		result += 1000 * (t.Perf.Audience - 30)
	}
	return result
}

func (t TragedyCalculator) volumeCredits() int {
	result := int(math.Max(float64(t.Perf.Audience-30), 0))
	return result
}
