package factory

import pd "abstract_factory/product"

type Nike struct {
}

type NikeShirt struct {
	pd.Shirt
}

type NikeShoe struct {
	pd.Shoe
}

func (n *Nike) MakeShoe() pd.IShoe {
	return &NikeShoe{
		// Shoe: pd.Shoe{
		// 	logo: "nike",
		// 	size: 14,
		// },
		Shoe: *pd.NewShoe("nike", 14),
	}
}

func (n *Nike) MakeShirt() pd.IShirt {
	return &NikeShirt{
		// Shirt: pd.Shirt{
		// 	logo: "nike",
		// 	size: 14,
		// },
		Shirt: *pd.NewShirt("nike", 14),
	}
}
