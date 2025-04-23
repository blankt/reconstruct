package chapter_1

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
打印收费详情
代码存在的问题：
1.代码组织不清晰
2.函数职责不清晰
3.如果需要添加新特性，不利于修改，例如需要输出为html格式、增加多种戏剧类型，则要增加更多的判断。
*/
func statement(invoice Invoice, plays map[string]Play) string {
	totalAmount := 0
	volumeCredits := 0
	result := fmt.Sprintf("Statement for %s\n", invoice.Customer)

	format := func(amount int) string {
		return fmt.Sprintf("$%.2f", float64(amount)/100)
	}

	for _, perf := range invoice.Performances {
		play, ok := plays[perf.PlayID]
		if !ok {
			panic(fmt.Sprintf("unknown play ID: %s", perf.PlayID))
		}

		thisAmount := 0
		switch play.Type {
		case "tragedy":
			thisAmount = 40000
			if perf.Audience > 30 {
				thisAmount += 1000 * (perf.Audience - 30)
			}
		case "comedy":
			thisAmount = 30000
			if perf.Audience > 20 {
				thisAmount += 10000 + 500*(perf.Audience-20)
			}
			thisAmount += 300 * perf.Audience
		default:
			panic(fmt.Sprintf("unknown type: %s", play.Type))
		}

		// add volume credits
		volumeCredits += int(math.Max(float64(perf.Audience-30), 0))
		// add extra credit for every ten comedy attendees
		if play.Type == "comedy" {
			volumeCredits += perf.Audience / 5
		}

		// print line for this order
		result += fmt.Sprintf(" %s: %s (%d seats)\n", play.Name, format(thisAmount), perf.Audience)
		totalAmount += thisAmount
	}

	result += fmt.Sprintf("Amount owed is %s\n", format(totalAmount))
	result += fmt.Sprintf("You earned %d credits\n", volumeCredits)
	return result
}
