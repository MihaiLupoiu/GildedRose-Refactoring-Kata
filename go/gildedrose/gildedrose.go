package gildedrose

import (
	"github.com/emilybache/gildedrose-refactoring-kata/gildedrose/items"
)

func UpdateQuality(items []*items.Item) {
	for _, item := range items {
		item.GetUpdateStrategy().Update()
	}
}
