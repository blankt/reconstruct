package invoice1

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

/*
第一阶段：把复杂的代码拆分为更小的单元

打印收费详情优化版本1
所做的优化：
1.提炼函数（106）
2.以查询取代临时变量（178）：分解长函数的时候，只具有局部作用域的临时变量会试提炼函数变得更复杂
3.改变函数申明（124）：好的函数名十分重要
4.内联变量（123）：去除临时变量，直接用函数返回值
5.应用拆分循环（227）
6.移动语句（223）: 将变量和逻辑惯量强的代码移动到一起

担忧的问题：
重复循环带来的性能消耗问题大吗？
大多数时候软件的性能只和一小部分代码有关，改变的一些部分是对性能的影响甚微。
大多数情况下可以忽略它。如果重构引入了性能损耗，先完成重构，再做性能优化。
*/
func statement(invoice Invoice, plays map[string]Play) string {
	result := fmt.Sprintf("Statement for %s\n", invoice.Customer)
	for _, perf := range invoice.Performances {
		// print line for this order
		result += fmt.Sprintf(" %s: %s (%d seats)\n", PlayFor(perf, plays).Name, Usd(AmountFor(perf, plays)), perf.Audience)
	}

	result += fmt.Sprintf("Amount owed is %s\n", Usd(totalAmount(invoice.Performances, plays)))
	result += fmt.Sprintf("You earned %d credits\n", totalVolumeCredits(invoice.Performances, plays))
	return result
}

func PlayFor(performance Performance, plays map[string]Play) Play {
	play, ok := plays[performance.PlayID]
	if !ok {
		panic(fmt.Sprintf("unknown play ID: %s", performance.PlayID))
	}
	return play
}

func AmountFor(performance Performance, plays map[string]Play) int {
	result := 0
	switch PlayFor(performance, plays).Type {
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
		panic(fmt.Sprintf("unknown type: %s", PlayFor(performance, plays).Type))
	}
	return result
}

func volumeCreditsFor(performance Performance, plays map[string]Play) int {
	result := 0
	result += int(math.Max(float64(performance.Audience-30), 0))
	if PlayFor(performance, plays).Type == "comedy" {
		result += performance.Audience / 5
	}
	return result
}

func totalVolumeCredits(performances []Performance, plays map[string]Play) int {
	total := 0
	for _, perf := range performances {
		total += volumeCreditsFor(perf, plays)
	}
	return total
}

func totalAmount(performances []Performance, plays map[string]Play) int {
	total := 0
	for _, perf := range performances {
		total += AmountFor(perf, plays)
	}
	return total
}

func Usd(amount int) string {
	return fmt.Sprintf("$%.2f", float64(amount)/100)
}
