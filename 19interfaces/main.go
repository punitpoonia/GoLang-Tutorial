package main

import "fmt"

type payment struct{}

func (p payment) makePayment(amount float32) {
	myRazorPayGw := razorpay{}
	myRazorPayGw.pay(amount)
}

type razorpay struct{}

func (r razorpay) pay(amount float32) {

	fmt.Println("Payment is done through razorpay : ", amount)

}

func main() {
	newPayment := payment{}
	newPayment.makePayment(100)
}
