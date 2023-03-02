package factory

import pd "abstract_factory/product"

type Adidas struct {
}

type AdidasShoe struct {
	pd.Shoe
}

type AdidasShirt struct {
	pd.Shirt
}

func (a *Adidas) MakeShoe() pd.IShoe {
	return &AdidasShoe{
		// Shoe: pd.Shoe{
		// 	logo: "adidas",
		// 	size: 14,
		// },
		Shoe: *pd.NewShoe("adidas", 14),
	}
}

func (a *Adidas) MakeShirt() pd.IShirt {
	return &AdidasShirt{
		// Shirt: pd.Shirt{
		// 	logo: "adidas",
		// 	size: 14,
		// },
		Shirt: *pd.NewShirt("adidas", 14),
	}
}
