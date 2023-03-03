package intf

import rq "CoR/requester"

type Department interface {
	Execute(*rq.Patient)
	SetNext(Department)
}
