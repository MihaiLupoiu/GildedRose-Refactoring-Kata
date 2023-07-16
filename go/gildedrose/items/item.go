package items

import (
	"strings"
)

type Item struct {
	Name            string
	SellIn, Quality int
}

const (
	SULFURAS        = "Sulfuras, Hand of Ragnaros"
	AGED_BRIE       = "Aged Brie"
	BACKSTAGE_PASS  = "Backstage passes to a TAFKAL80ETC concert"
	CONJURED_PREFIX = "Conjured"
)

// The Quality of an item is never more than 50
const MAX_QUALITY = 50

// The Quality of an item is never negative
const MIN_QUALITY = 0

type ItemStrategy interface {
	Update()
}

func (item *Item) GetUpdateStrategy() ItemStrategy {
	switch {
	case item.Name == SULFURAS:
		return &SulfurasItem{item}
	case item.Name == AGED_BRIE:
		return &AgedBrieItem{item}
	case item.Name == BACKSTAGE_PASS:
		return &BackstagePassItem{item}
	case strings.Contains(item.Name, CONJURED_PREFIX):
		return &ConjuredItem{item}
	default:
		return &StandardItem{item}
	}
}

func (item *Item) increaseQuality() {
	if item.Quality < MAX_QUALITY {
		item.Quality++
	}
}

func (item *Item) decreaseQuality() {
	if item.Quality > MIN_QUALITY {
		item.Quality--
	}
}
