package main

import (
	"abstract_factory/factory"
	pd "abstract_factory/product"
	"fmt"
)

func main() {
	nikeFactory, _ := factory.GetSportsFactory("nike")
	nikeShoe := nikeFactory.MakeShoe()
	nikeShirt := nikeFactory.MakeShirt()

	adidasFactory, _ := factory.GetSportsFactory("adidas")
	adidasShoe := adidasFactory.MakeShoe()
	adidasShirt := adidasFactory.MakeShirt()

	printShoeDetails(nikeShoe)
	printShirtDetails(nikeShirt)

	printShoeDetails(adidasShoe)
	printShirtDetails(adidasShirt)
}

func printShoeDetails(s pd.IShoe) {
	fmt.Printf("Logo: %s\n", s.GetLogo())
	fmt.Printf("Size: %d\n", s.GetSize())
}

func printShirtDetails(s pd.IShirt) {
	fmt.Printf("Logo: %s\n", s.GetLogo())
	fmt.Printf("Size: %d\n", s.GetSize())
}
