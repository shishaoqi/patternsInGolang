package hospitalHandler

import (
	"CoR/intf"
	rq "CoR/requester"
	"fmt"
)

type Doctor struct {
	next intf.Department
}

func (d *Doctor) Execute(p *rq.Patient) {
	if p.DoctorCheckUpDone() {
		fmt.Println("Doctor checkup already done")
		d.next.Execute(p)
		return
	}
	fmt.Println("Doctor checking patient")
	//p.doctorCheckUpDone = true
	p.SetDoctorCheckUpDone(true)
	d.next.Execute(p)
}

func (d *Doctor) SetNext(next intf.Department) {
	d.next = next
}
