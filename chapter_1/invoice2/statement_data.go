package invoice2

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
第二阶段：为账单打印提供一个html版本
把代码逻辑拆分为账单数据计算与打印两个部分

所做优化：
1.拆分阶段(154): 将逻辑拆分为多个部分，通过中间数据结构进行传递。
2.搬移函数（198）
3.提炼函数（106）
4.管道取代循环（231）？

最后再把两阶段拆分为两个文件
*/
func createStatementData(invoice Invoice, plays map[string]Play) statementData {
	data := statementData{}
	data.Customer = invoice.Customer
	data.Performances = enrichPerformances(invoice.Performances, plays)
	data.TotalAmount = totalAmount2(data.Performances)
	data.TotalVolumeCredits = totalVolumeCredits2(data.Performances)
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
		result[i].Amount = AmountFor2(result[i])
		result[i].Credits = volumeCreditsFor2(result[i])
	}
	return result
}

func totalVolumeCredits2(performances []MiddlePerformance) int {
	total := 0
	for _, perf := range performances {
		total += perf.Credits
	}
	return total
}

func totalAmount2(performances []MiddlePerformance) int {
	total := 0
	for _, perf := range performances {
		total += perf.Amount
	}
	return total
}

func AmountFor2(performance MiddlePerformance) int {
	result := 0
	switch performance.Play.Type {
	case "tragedy":
		result = 40000
		if performance.Audience > 30 {
			result += 1000 * (performance.Audience - 30)
		}
	case "comedy":
		result = 30000
		if performance.Audience > 20 {
			result += 10000 + 500*(performance.Audience-20)
		}
		result += 300 * performance.Audience
	default:
		panic(fmt.Sprintf("unknown type: %s", performance.Play.Type))
	}
	return result
}

func volumeCreditsFor2(performance MiddlePerformance) int {
	result := 0
	result += int(math.Max(float64(performance.Audience-30), 0))
	if performance.Play.Type == "comedy" {
		result += performance.Audience / 5
	}
	return result
}

func PlayFor(performance Performance, plays map[string]Play) Play {
	play, ok := plays[performance.PlayID]
	if !ok {
		panic(fmt.Sprintf("unknown play ID: %s", performance.PlayID))
	}
	return play
}
