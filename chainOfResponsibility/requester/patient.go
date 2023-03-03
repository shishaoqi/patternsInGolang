package requester

type Patient struct {
	Name              string
	registrationDone  bool
	doctorCheckUpDone bool
	medicineDone      bool
	paymentDone       bool
}

func (p Patient) RegistrationDone() bool {
	return p.registrationDone
}

func (p Patient) SetRegistrationDone(result bool) {
	p.registrationDone = result
}

func (p Patient) DoctorCheckUpDone() bool {
	return p.doctorCheckUpDone
}

func (p *Patient) SetDoctorCheckUpDone(result bool) {
	p.doctorCheckUpDone = result
}

func (p Patient) MedicineDone() bool {
	return p.medicineDone
}

func (p *Patient) SetMedicineDone(result bool) {
	p.medicineDone = result
}

func (p Patient) PaymentDone() bool {
	return p.paymentDone
}

func (p *Patient) SetPaymentDone(result bool) {
	p.paymentDone = result
}
