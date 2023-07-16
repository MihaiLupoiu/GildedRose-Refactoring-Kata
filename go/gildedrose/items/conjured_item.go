package items

type ConjuredItem struct {
	*Item
}

func (item *ConjuredItem) Update() {
	// “Conjured” items degrade in Quality twice as fast as normal items
	item.decreaseQuality()
	item.decreaseQuality()

	item.SellIn--
	if item.SellIn < 0 {
		item.decreaseQuality()
		item.decreaseQuality()
	}
}
