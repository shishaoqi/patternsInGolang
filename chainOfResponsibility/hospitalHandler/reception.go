package hospitalHandler

import (
	"CoR/intf"
	rq "CoR/requester"
	"fmt"
)

type Reception struct {
	next intf.Department
}

func (r *Reception) Execute(p *rq.Patient) {
	if p.RegistrationDone() {
		fmt.Println("Patient registration already done")
		r.next.Execute(p)
		return
	}
	fmt.Println("Reception registering patient")
	//p.registrationDone = true
	p.SetRegistrationDone(true)
	r.next.Execute(p)
}

func (r *Reception) SetNext(next intf.Department) {
	r.next = next
}
