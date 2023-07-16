package items

type BackstagePassItem struct {
	*Item
}

func (item *BackstagePassItem) Update() {
	// “Backstage passes”, like aged brie, increases in Quality as its SellIn value approaches;
	item.increaseQuality()

	// Quality increases by 2 when there are 10 days or less and by 3 when there are 5 days or less but
	if item.SellIn <= 10 {
		item.increaseQuality()
	}

	if item.SellIn <= 5 {
		item.increaseQuality()
	}

	item.SellIn--
	// Quality drops to 0 after the concert
	if item.SellIn < 0 {
		item.Quality = 0
	}
}
