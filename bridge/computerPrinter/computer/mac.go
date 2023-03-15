package computer

import (
	p "bridge/printer"
	"fmt"
)

type Mac struct {
	printer p.Printer
}

func (m *Mac) Print() {
	fmt.Println("Print request for mac")
	m.printer.PrintFile()
}

func (m *Mac) SetPrinter(p p.Printer) {
	m.printer = p
}
