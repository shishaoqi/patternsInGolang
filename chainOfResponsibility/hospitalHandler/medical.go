package hospitalHandler

import (
	"CoR/intf"
	rq "CoR/requester"
	"fmt"
)

type Medical struct {
	next intf.Department
}

func (m *Medical) Execute(p *rq.Patient) {
	if p.MedicineDone() {
		fmt.Println("Medicine already given to patient")
		m.next.Execute(p)
		return
	}
	fmt.Println("Medical giving medicine to patient")
	//p.medicineDone = true
	p.SetMedicineDone(true)
	m.next.Execute(p)
}

func (m *Medical) SetNext(next intf.Department) {
	m.next = next
}
