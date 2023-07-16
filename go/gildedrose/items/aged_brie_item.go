package items

type AgedBrieItem struct {
	*Item
}

func (item *AgedBrieItem) Update() {
	// “Aged Brie” actually increases in Quality the older it gets
	item.increaseQuality()

	// At the end of each day our system lowers SellIn values for every item
	item.SellIn--

	// Once the sell by date has passed, Quality increases twice as fast
	// The Quality of an item is never negative
	if item.SellIn < 0 {
		item.increaseQuality()
	}
}
