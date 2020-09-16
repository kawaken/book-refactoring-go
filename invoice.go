package main

import (
	"fmt"

	"github.com/leekchan/accounting"
)

type Invoice struct {
	Customer     string `json:"customer"`
	Performances []*struct {
		PlayID   string `json:"playID"`
		Audience int    `json:"audience"`
	} `json:"performances"`
}

type Play struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func statement(invoice *Invoice, plays map[string]*Play) (string, error) {
	totalAmount := 0
	volumeCredit := 0
	result := fmt.Sprintf("Statement for %s\n", invoice.Customer)

	format := accounting.DefaultAccounting("$", 2).FormatMoney

	for _, perf := range invoice.Performances {
		play := plays[perf.PlayID]
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
			return "", fmt.Errorf("unknown type: %s", play.Type)
		}

		// ボリューム特典のポイントを加算
		volumeCredit += Max(perf.Audience-30, 0)
		// 喜劇の時は10人につき、さらにポイントを加算
		if play.Type == "comedy" {
			volumeCredit += perf.Audience / 5 // math.Floorいるかな？
		}
		// 注文の内訳を出力
		result += fmt.Sprintf("  %s: %s (%d seats)\n", play.Name, format(thisAmount/100), perf.Audience)
		totalAmount += thisAmount
	}
	result += fmt.Sprintf("Amount owed is %s\n", format(totalAmount/100))
	result += fmt.Sprintf("You earned %d credits\n", volumeCredit)
	return result, nil
}
