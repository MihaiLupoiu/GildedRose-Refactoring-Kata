package items

type StandardItem struct {
	*Item
}

func (item *StandardItem) Update() {
	// At the end of each day our system lowers Quality values for every item
	item.decreaseQuality()

	// At the end of each day our system lowers SellIn values for every item
	item.SellIn--

	// Once the sell by date has passed, Quality degrades twice as fast
	// The Quality of an item is never negative
	if item.SellIn < 0 {
		item.decreaseQuality()
	}
}
