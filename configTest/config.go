package config

// Cart represents the state of a buyer's shopping cart
type Config struct {
	lineItems map[string]LineItem
}

// Item represents any item available for sale
type LineItem struct {
	ID    string
	Name  string
	Price float64
	Qty   int
}

func (c *Config) init() {
	if c.lineItems == nil {
		c.lineItems = map[string]LineItem{}
	}
}

// AddItem adds an item to the cart
func (c *Config) AddLineItem(i LineItem) {
	c.init()
	if existingLineItem, ok := c.lineItems[i.ID]; ok {
		existingLineItem.Qty++
		c.lineItems[i.ID] = existingLineItem
	} else {
		i.Qty = 1
		c.lineItems[i.ID] = i
	}
}

// RemoveItem removes n number of items with give id from the cart
func (c *Config) RemoveLineItem(id string, n int) {
	c.init()
	if existingLineItem, ok := c.lineItems[id]; ok {
		if existingLineItem.Qty <= n {
			delete(c.lineItems, id)
		} else {
			existingLineItem.Qty -= n
			c.lineItems[id] = existingLineItem
		}
	}
}

// TotalAmount returns the total amount of the cart
func (c *Config) TotalAmount() float64 {
	c.init()
	totalAmount := 0.0
	for _, i := range c.lineItems {
		totalAmount += i.Price * float64(i.Qty)
	}
	return totalAmount
}

// TotalUnits returns the total number of units across all items in the cart
func (c *Config) TotalUnits() int {
	c.init()
	totalUnits := 0
	for _, i := range c.lineItems {
		totalUnits += i.Qty
	}
	return totalUnits
}

// TotalUniqueItems returns the number of unique items in the cart
func (c *Config) TotalUniqueLineItems() int {
	return len(c.lineItems)
}
