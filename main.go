package main

import (
	"fmt"

	liqpay "github.com/liqpay/go-sdk"
)

var c = liqpay.New("", "", nil)

func main() {
	r := liqpay.Request{
		"card":           "4000000000003063",
		"card_exp_month": "07",
		"card_exp_year":  "24",
		"action":         "pay",
		"version":        3,
		"amount":         3,
		"currency":       "UAH",
		"description":    "Pay test",
		"order_id":       "4",
		"server_url":     "",
	}
	resp, err := c.Send("request", r)
	if err != nil {
		fmt.Printf("error %v\n", err.Error())
	}
	fmt.Printf("response %#v\n", resp)
}
