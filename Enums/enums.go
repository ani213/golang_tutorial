package main

import "fmt"

// Enumarated types

type OrderStatus int

const (
	Pending OrderStatus = iota
	Processing
	Shipped
	Delivered
	Cancelled
)

type PaymentStatus string

const (
	Paid     PaymentStatus = "Paid"
	Unpaid   PaymentStatus = "Unpaid"
	Refunded PaymentStatus = "Refunded"
)

func getPaymentStatus(paymentStatus PaymentStatus) string {
	fmt.Println("Payment Status:-", paymentStatus)
	switch paymentStatus {
	case Paid:
		return "Payment has been made."
	case Unpaid:
		return "Payment is pending."
	case Refunded:
		return "Payment has been refunded."
	default:
		return "Unknown payment status."
	}

}

func changeStatus(status OrderStatus) string {
	fmt.Println("status", status)
	// Switch statement to handle different order statuses
	switch status {
	case Pending:
		return "Order is pending."
	case Processing:
		return "Order is being processed."
	case Shipped:
		return "Order has been shipped."
	case Delivered:
		return "Order has been delivered."
	case Cancelled:
		return "Order has been cancelled."
	default:
		return "Unknown order status."
	}
}

func main() {
	fmt.Println(changeStatus(Processing))
	fmt.Println(getPaymentStatus(Paid))
	fmt.Println(getPaymentStatus(Unpaid))
	fmt.Println(getPaymentStatus(Refunded))
	fmt.Println(getPaymentStatus("Unknown")) // This will not match any defined PaymentStatus
}
