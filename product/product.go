// methods/example-project/product/product.go
package product

import money "github.com/Rhymond/go-money"

type Product struct {
	ID    string
	Name  string
	Price *money.Money
}
