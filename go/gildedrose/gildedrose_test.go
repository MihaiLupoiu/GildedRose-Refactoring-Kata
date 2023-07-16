package gildedrose_test

import (
	"testing"

	"github.com/emilybache/gildedrose-refactoring-kata/gildedrose"
	"github.com/stretchr/testify/assert"

	"github.com/emilybache/gildedrose-refactoring-kata/gildedrose/items"
)

type testCase struct {
	name     string
	item     *items.Item
	expected *items.Item
}

func Test_Foo(t *testing.T) {
	var items = []*items.Item{
		{Name: "foo", SellIn: 0, Quality: 0},
	}

	gildedrose.UpdateQuality(items)

	assert.Equal(t, items[0].Name, "foo", "Name should not change after updating quality")
}

func Test_StandardItemUpdate(t *testing.T) {
	testCases := []testCase{
		{
			name:     "StandardItem - SellIn and Quality decrease by 1 after each day",
			item:     &items.Item{Name: "StandardItem", SellIn: 5, Quality: 5},
			expected: &items.Item{Name: "StandardItem", SellIn: 4, Quality: 4},
		},
		{
			name:     "StandardItem - Quality degrades twice as fast after the sell-by date has passed",
			item:     &items.Item{Name: "StandardItem", SellIn: 0, Quality: 5},
			expected: &items.Item{Name: "StandardItem", SellIn: -1, Quality: 3},
		},
		{
			name:     "StandardItem - Quality is never negative",
			item:     &items.Item{Name: "StandardItem", SellIn: 5, Quality: 0},
			expected: &items.Item{Name: "StandardItem", SellIn: 4, Quality: 0},
		},
	}

	runUpdateItemTestCases(t, testCases, gildedrose.UpdateQuality)
}

func Test_AgedBrieItemUpdate(t *testing.T) {
	testCases := []testCase{
		{
			name:     "AgedBrieItem - SellIn decrease by 1 and Quality increases by 1 after each day",
			item:     &items.Item{Name: "Aged Brie", SellIn: 5, Quality: 5},
			expected: &items.Item{Name: "Aged Brie", SellIn: 4, Quality: 6},
		},
		{
			name:     "AgedBrieItem - Quality increases  twice as fast after the sell-by date has passed",
			item:     &items.Item{Name: "Aged Brie", SellIn: 0, Quality: 5},
			expected: &items.Item{Name: "Aged Brie", SellIn: -1, Quality: 7},
		},
		{
			name:     "AgedBrieItem - Quality is never more than 50",
			item:     &items.Item{Name: "Aged Brie", SellIn: 5, Quality: 49},
			expected: &items.Item{Name: "Aged Brie", SellIn: 4, Quality: 50},
		},
		{
			name:     "AgedBrieItem - Quality is never more than 50",
			item:     &items.Item{Name: "Aged Brie", SellIn: 5, Quality: 50},
			expected: &items.Item{Name: "Aged Brie", SellIn: 4, Quality: 50},
		},
	}

	runUpdateItemTestCases(t, testCases, gildedrose.UpdateQuality)
}

func Test_SulfurasItemUpdate(t *testing.T) {
	testCases := []testCase{
		{
			name:     "SulfurasItem - SellIn and Quality remain unchanged after each day",
			item:     &items.Item{Name: "Sulfuras, Hand of Ragnaros", SellIn: 5, Quality: 5},
			expected: &items.Item{Name: "Sulfuras, Hand of Ragnaros", SellIn: 5, Quality: 5},
		},
	}

	runUpdateItemTestCases(t, testCases, gildedrose.UpdateQuality)
}

func Test_BackstagePassesItemUpdate(t *testing.T) {
	testCases := []testCase{
		{
			name:     "BackstagePassesItem - SellIn decrease by 1 and Quality increases by 1 after each day when SellIn is more than 10",
			item:     &items.Item{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 15, Quality: 5},
			expected: &items.Item{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 14, Quality: 6},
		},
		{
			name:     "BackstagePassesItem - SellIn decrease by 1 and Quality increases by 1 after each day when SellIn is between 10 and 6",
			item:     &items.Item{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 10, Quality: 5},
			expected: &items.Item{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 9, Quality: 7},
		},
		{
			name:     "BackstagePassesItem - SellIn decrease by 1 and Quality increases by 1 after each day when SellIn is between 10 and 6",
			item:     &items.Item{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 6, Quality: 5},
			expected: &items.Item{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 5, Quality: 7},
		},
		{
			name:     "BackstagePassesItem - SellIn decrease by 1 and Quality increases by 3 after each day when SellIn is between 5 and 1",
			item:     &items.Item{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 5, Quality: 5},
			expected: &items.Item{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 4, Quality: 8},
		},
		{
			name:     "BackstagePassesItem - SellIn decrease by 1 and Quality increases by 3 after each day when SellIn is between 5 and 1",
			item:     &items.Item{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 1, Quality: 5},
			expected: &items.Item{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 0, Quality: 8},
		},
		{
			name:     "BackstagePassesItem - Quality drops to 0 after the concert (when SellIn is negative)",
			item:     &items.Item{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 0, Quality: 5},
			expected: &items.Item{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: -1, Quality: 0},
		},
		{
			name:     "BackstagePassesItem -  Quality is never more than 50",
			item:     &items.Item{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 2, Quality: 49},
			expected: &items.Item{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 1, Quality: 50},
		},
	}

	runUpdateItemTestCases(t, testCases, gildedrose.UpdateQuality)
}

func Test_ConjuredItemUpdate(t *testing.T) {
	testCases := []testCase{
		{
			name:     "ConjuredItem - SellIn decrease by 1 and Quality decrease by 2 after each day",
			item:     &items.Item{Name: "Conjured Item", SellIn: 5, Quality: 5},
			expected: &items.Item{Name: "Conjured Item", SellIn: 4, Quality: 3},
		},
		{
			name:     "ConjuredItem - Quality degrades four times as fast after the sell-by date has passed",
			item:     &items.Item{Name: "Conjured Item", SellIn: 0, Quality: 5},
			expected: &items.Item{Name: "Conjured Item", SellIn: -1, Quality: 1},
		},
		{
			name:     "ConjuredItem - Quality is never negative",
			item:     &items.Item{Name: "Conjured Item", SellIn: 5, Quality: 1},
			expected: &items.Item{Name: "Conjured Item", SellIn: 4, Quality: 0},
		},
	}

	runUpdateItemTestCases(t, testCases, gildedrose.UpdateQuality)
}

func runUpdateItemTestCases(t *testing.T, testCases []testCase, updateFunc func([]*items.Item)) {
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			items := []*items.Item{testCase.item}
			updateFunc(items)
			assert.Equal(t, testCase.expected, testCase.item, "Item update incorrect")
		})
	}
}

func Test_MultipleItems(t *testing.T) {
	items := []*items.Item{
		{
			Name:    "Apple",
			SellIn:  5,
			Quality: 5,
		},
		{
			Name:    "Aged Brie",
			SellIn:  1,
			Quality: 50,
		},
		{
			Name:    "Sulfuras, Hand of Ragnaros",
			SellIn:  10,
			Quality: 30,
		},
		{
			Name:    "Backstage passes to a TAFKAL80ETC concert",
			SellIn:  -1,
			Quality: 20,
		},
		{
			Name:    "Conjured Apple",
			SellIn:  2,
			Quality: 20,
		},
	}

	gildedrose.UpdateQuality(items)

	assert.Equal(t, 4, items[0].Quality, "normal items should have quality tick down")
	assert.Equal(t, 4, items[0].SellIn, "normal items should have sellin tick down")

	assert.Equal(t, 50, items[1].Quality, "Aged Brie should get better as it gets older until it has 50 quality")
	assert.Equal(t, 0, items[1].SellIn, "Aged Brie should have sellin tick down")

	assert.Equal(t, 30, items[2].Quality, "Sulfuras should have a quality of 30 always")
	assert.Equal(t, 10, items[2].SellIn, "sellIn should not change for SulfurasItem")

	assert.Equal(t, 0, items[3].Quality, "Backstage passes should have no quality after the conert")

	assert.Equal(t, 18, items[4].Quality, "Conjured shouldshould have quality tick down by 2")
	assert.Equal(t, 1, items[4].SellIn, "Conjured should have sellin tick down")
}

func Test_ManyTicks(t *testing.T) {
	const BASE = 100
	itemList := []*items.Item{
		{
			Name:    "Apple",
			SellIn:  BASE,
			Quality: 40,
		},
		{
			Name:    "Aged Brie",
			SellIn:  BASE,
			Quality: 1,
		},
		{
			Name:    "Sulfuras, Hand of Ragnaros",
			SellIn:  BASE,
			Quality: 30,
		},
		{
			Name:    "Backstage passes to a TAFKAL80ETC concert",
			SellIn:  BASE,
			Quality: 20,
		},
	}

	for i := 0; i < 12000; i++ {
		gildedrose.UpdateQuality(itemList)
		for _, item := range itemList {
			if item.Name == items.SULFURAS {
				assert.Equal(t, BASE, item.SellIn, "SulfurasItem should not have SellIn decrement")
			} else {
				assert.Equal(t, BASE-i-1, item.SellIn, "All items should have sellin decrement")
			}
		}
	}

	assert.Equal(t, 0, itemList[0].Quality, "normal items should have sellin tick down")
	assert.Equal(t, 50, itemList[1].Quality, "Aged Brie should get better as it gets older until it has 50 quality")
	assert.Equal(t, 30, itemList[2].Quality, "SulfurasItem should have a quality of 30 always")
	assert.Equal(t, 0, itemList[3].Quality, "Backstage passes should have no quality after the conert")
}
