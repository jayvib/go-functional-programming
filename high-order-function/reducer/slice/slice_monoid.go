package main

type LineItemMonoid interface {
	Zero() LineItem
	Append(lineitems ...LineItem) LineItemMonoid
	Reduce() LineItem
}

type LineItem struct {
	Quantity int
	Price int
	ListPrice int
}

type LineItemContainer struct {
	lineItems []LineItem
}

func (l LineItemContainer) Zero() []LineItem {
	return nil
}

func (l LineItemContainer) Append(lineitems ...LineItem) LineItemContainer {
	l.lineItems = append(l.lineItems, lineitems...)
	return l
}

func (l LineItemContainer) Reduce() LineItem {
	totalQuantity := 0
	totalPrice := 0
	totalListPrice := 0
	for _, li := range l.lineItems {
		totalQuantity += li.Quantity
		totalPrice += li.Price
		totalListPrice += li.ListPrice
	}
	return LineItem{
		Quantity: totalQuantity,
		Price: totalPrice,
		ListPrice: totalListPrice,
	}
}
