package main

import (
	cpt "bridge/computer"
	p "bridge/printer"
	"fmt"
)

func main() {

	hpPrinter := &p.Hp{}
	epsonPrinter := &p.Epson{}

	macComputer := &cpt.Mac{}
	// mac 组合 hp
	macComputer.SetPrinter(hpPrinter)
	macComputer.Print()
	fmt.Println()
	// mac 组合 epson
	macComputer.SetPrinter(epsonPrinter)
	macComputer.Print()
	fmt.Println()

	winComputer := &cpt.Windows{}
	// windows 组合 hp
	winComputer.SetPrinter(hpPrinter)
	winComputer.Print()
	fmt.Println()
	// windows 组合 epson
	winComputer.SetPrinter(epsonPrinter)
	winComputer.Print()
	fmt.Println()
}
