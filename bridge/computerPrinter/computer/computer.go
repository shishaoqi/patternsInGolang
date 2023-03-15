package computer

import p "bridge/printer"

type Computer interface {
	Print()
	SetPrinter(p.Printer)
}
