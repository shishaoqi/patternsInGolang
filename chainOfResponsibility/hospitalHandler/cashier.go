package hospitalHandler

import (
	"CoR/intf"
	rq "CoR/requester"
	"fmt"
)

type Cashier struct {
	next intf.Department
}

func (c *Cashier) Execute(p *rq.Patient) {
	if p.PaymentDone() {
		fmt.Println("Payment Done")
	}
	fmt.Println("Cashier getting money from patient patient")
}

func (c *Cashier) SetNext(next intf.Department) {
	c.next = next
}
