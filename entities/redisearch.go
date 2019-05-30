package entities

type ClientName int

const (
	Carts ClientName = iota
	CartItems
	Checkouts
)

func (cn ClientName) String() string {
	return []string{
		"carts",
		"cart_items",
		"checkouts",
	}[cn]
}
