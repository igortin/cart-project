package cart

import (
	"errors"
	"os/user"
	"time"

	money "github.com/Rhymond/go-money"
	"github.com/igortin/cart-project/product"
)

type Cart struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	lockedAt  time.Time
	user.User
	Items        []Item
	CurrencyCode string
	isLocked     bool
}

type Item struct {
	product.Product
	Quantity uint8
}

// Public method
func (c *Cart) TotalPrice() (*money.Money, error) {
	sum := money.New(0, c.CurrencyCode)
	var err error
	c.Lock()
	for _, v := range c.Items {
		subItemSum := v.Product.Price.Multiply(int64(v.Quantity))

		sum, err = sum.Add(subItemSum)
		if err != nil {
			return nil, err
		}
	}
	c.delete()
	return sum, nil
}

// Public method Visible
func (c *Cart) Lock() error {

	if c.isLocked {
		return errors.New("already locked")
	}
	err := c.set()
	if err != nil {
		return err
	}
	return nil
}

// Private method Invisible
func (c *Cart) delete() error {
	c.isLocked = false
	c.lockedAt = time.Unix(0, 0)
	return nil
}

// Private method Invisible
func (c *Cart) set() error {
	c.isLocked = true
	c.lockedAt = time.Now()
	return nil
}
