package main

import (
	hl "CoR/hospitalHandler"
	rq "CoR/requester"
)

func main() {

	cashier := &hl.Cashier{}

	//Set next for medical department
	medical := &hl.Medical{}
	medical.SetNext(cashier)

	//Set next for doctor department
	doctor := &hl.Doctor{}
	doctor.SetNext(medical)

	//Set next for reception department
	reception := &hl.Reception{}
	reception.SetNext(doctor)

	patient := &rq.Patient{Name: "abc"}
	//Patient visiting
	reception.Execute(patient)
}
