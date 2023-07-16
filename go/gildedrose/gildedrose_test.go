package gildedrose_test

import (
	"testing"

	"github.com/emilybache/gildedrose-refactoring-kata/gildedrose"
	"github.com/stretchr/testify/assert"

	"github.com/emilybache/gildedrose-refactoring-kata/gildedrose/items"
)

func Test_Foo(t *testing.T) {
	var items = []*items.Item{
		{Name: "foo", SellIn: 0, Quality: 0},
	}

	gildedrose.UpdateQuality(items)

	assert.Equal(t, items[0].Name, "foo", "Name should not change after updating quality")
}

func Test_StandardItem(t *testing.T) {
	var items = []*items.Item{
		{
			Name:    "StandardItem",
			SellIn:  5,
			Quality: 5,
		},
		{
			Name:    "StandardItem",
			SellIn:  0,
			Quality: 5,
		},
		{
			Name:    "StandardItem",
			SellIn:  5,
			Quality: 0,
		},
	}

	gildedrose.UpdateQuality(items)

	// At the end of each day our system lowers SellIn values for every item
	assert.Equal(t, items[0].SellIn, 4, "SellIn should be one less than previusly.")

	// At the end of each day our system lowers Quality values for every item
	assert.Equal(t, items[0].Quality, 4, "Quality should be one less than previusly.")

	// Once the sell by date has passed, Quality degrades twice as fast
	assert.Equal(t, items[1].Quality, 3, "Quality should be two less than previusly.")
	assert.Equal(t, items[1].SellIn, -1, "SellIn should be one less than previusly.")

	// The Quality of an item is never negative
	assert.Equal(t, items[2].Quality, 0, "Quality should be 0.")
	assert.Equal(t, items[2].SellIn, 4, "SellIn should be one less than previusly.")
}

func Test_AgedBrieItem(t *testing.T) {
	var items = []*items.Item{
		{
			Name:    "Aged Brie",
			SellIn:  5,
			Quality: 5,
		},
		{
			Name:    "Aged Brie",
			SellIn:  0,
			Quality: 5,
		},
		{
			Name:    "Aged Brie",
			SellIn:  5,
			Quality: 0,
		},
		{
			Name:    "Aged Brie",
			SellIn:  5,
			Quality: 49,
		},
		{
			Name:    "Aged Brie",
			SellIn:  5,
			Quality: 50,
		},
	}

	gildedrose.UpdateQuality(items)

	// At the end of each day our system lowers SellIn values for every item
	assert.Equal(t, items[0].SellIn, 4, "SellIn should be one less than previusly.")

	// “Aged Brie” actually increases in Quality the older it gets
	assert.Equal(t, items[0].Quality, 6, "Quality should be one more than previusly.")

	// Once the sell by date has passed, Quality increases twice as fast
	assert.Equal(t, items[1].Quality, 7, "Quality should be two more than previusly.")
	assert.Equal(t, items[1].SellIn, -1, "SellIn should be one less than previusly.")

	// The Quality of an item is never negative
	assert.Equal(t, items[2].Quality, 1, "Quality should be 1.")
	assert.Equal(t, items[2].SellIn, 4, "SellIn should be one less than previusly.")

	// The Quality of an item is never more than 50
	assert.Equal(t, items[3].Quality, 50, "Quality should max 50.")
	assert.Equal(t, items[3].SellIn, 4, "SellIn should be one less than previusly.")

	assert.Equal(t, items[4].Quality, 50, "Quality should max 50.")
	assert.Equal(t, items[4].SellIn, 4, "SellIn should be one less than previusly.")
}

func Test_SulfurasItem(t *testing.T) {
	var items = []*items.Item{
		{
			Name:    "Sulfuras, Hand of Ragnaros",
			SellIn:  5,
			Quality: 5,
		},
	}

	gildedrose.UpdateQuality(items)

	// “Sulfuras”, being a legendary item, never has to be sold or decreases in Quality
	assert.Equal(t, items[0].SellIn, 5, "SellIn should be not have changed.")
	assert.Equal(t, items[0].Quality, 5, "Quality should be not have changed.")
}

func Test_BackstagePassesItem(t *testing.T) {
	var items = []*items.Item{
		{
			Name:    "Backstage passes to a TAFKAL80ETC concert",
			SellIn:  15,
			Quality: 5,
		},
		{
			Name:    "Backstage passes to a TAFKAL80ETC concert",
			SellIn:  9,
			Quality: 5,
		},
		{
			Name:    "Backstage passes to a TAFKAL80ETC concert",
			SellIn:  4,
			Quality: 5,
		},
		{
			Name:    "Backstage passes to a TAFKAL80ETC concert",
			SellIn:  0,
			Quality: 5,
		},
		{
			Name:    "Backstage passes to a TAFKAL80ETC concert",
			SellIn:  3,
			Quality: 49,
		},
		{
			Name:    "Backstage passes to a TAFKAL80ETC concert",
			SellIn:  9,
			Quality: 49,
		},

		{
			Name:    "Backstage passes to a TAFKAL80ETC concert",
			SellIn:  15,
			Quality: 50,
		},
	}

	gildedrose.UpdateQuality(items)

	// “Backstage passes”, like aged brie, increases in Quality as its SellIn value approaches;
	assert.Equal(t, items[0].SellIn, 14, "SellIn should be one less than previusly.")
	assert.Equal(t, items[0].Quality, 6, "Quality should be one more than previusly.")

	// Quality increases by 2 when there are 10 days or less and by 3 when there are 5 days or less but
	assert.Equal(t, items[1].SellIn, 8, "SellIn should be one less than previusly.")
	assert.Equal(t, items[1].Quality, 7, "Quality should be two more than previusly.")
	assert.Equal(t, items[2].SellIn, 3, "SellIn should be one less than previusly.")
	assert.Equal(t, items[2].Quality, 8, "Quality should be three more than previusly.")

	// Quality drops to 0 after the concert
	assert.Equal(t, items[3].SellIn, -1, "SellIn should be one less than previusly.")
	assert.Equal(t, items[3].Quality, 0, "Quality should be 0.")

	assert.Equal(t, items[4].Quality, 50, "Quality should max 50.")
	assert.Equal(t, items[4].SellIn, 2, "SellIn should be one less than previusly.")

	assert.Equal(t, items[5].Quality, 50, "Quality should max 50.")
	assert.Equal(t, items[5].SellIn, 8, "SellIn should be one less than previusly.")

	assert.Equal(t, items[6].Quality, 50, "Quality should max 50.")
	assert.Equal(t, items[6].SellIn, 14, "SellIn should be one less than previusly.")
}

func Test_ConjuredItem(t *testing.T) {
	var items = []*items.Item{
		{
			Name:    "Conjured",
			SellIn:  5,
			Quality: 5,
		},
		{
			Name:    "Conjured",
			SellIn:  0,
			Quality: 5,
		},
	}

	gildedrose.UpdateQuality(items)

	// “Conjured” items degrade in Quality twice as fast as normal items
	assert.Equal(t, items[0].SellIn, 4, "SellIn should be one less than previusly.")
	assert.Equal(t, items[0].Quality, 3, "Quality should be two less than previusly.")
	assert.Equal(t, items[1].SellIn, -1, "SellIn should be one less than previusly.")
	assert.Equal(t, items[1].Quality, 1, "Quality should be four times less than previusly.")
}

// TODO:

// Test multiple items.
// Test multiple ticks of time.
// Todo make test in table format.
