package computer

import (
	p "bridge/printer"
	"fmt"
)

type Windows struct {
	printer p.Printer
}

func (w *Windows) Print() {
	fmt.Println("Print request for windows")
	w.printer.PrintFile()
}

func (w *Windows) SetPrinter(p p.Printer) {
	w.printer = p
}
