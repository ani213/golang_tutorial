package main

import (
	"fmt"
	"strconv"
)

type paymentr interface {
	Pay(amount int) string
}
type payment struct {
	gateWay paymentr
}

type razorpay struct{}

func (r razorpay) Pay(amount int) string {
	return "Paid:-" + strconv.Itoa(amount) + " using Razorpay"
}

func main() {
	razorpayGw := razorpay{}
	paymentObj := payment{gateWay: razorpayGw}
	fmt.Println(paymentObj.gateWay.Pay(1000))
}
