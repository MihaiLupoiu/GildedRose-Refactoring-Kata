package items

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestItem struct {
	item      *Item
	strategy  ItemStrategy
	expectErr bool
}

func TestItem_GetUpdateStrategy(t *testing.T) {
	assert := assert.New(t)

	testCases := []TestItem{
		{
			item:      &Item{Name: SULFURAS, SellIn: 10, Quality: 80},
			strategy:  &SulfurasItem{&Item{Name: SULFURAS, SellIn: 10, Quality: 80}},
			expectErr: false,
		},
		{
			item:      &Item{Name: AGED_BRIE, SellIn: 5, Quality: 20},
			strategy:  &AgedBrieItem{&Item{Name: AGED_BRIE, SellIn: 5, Quality: 20}},
			expectErr: false,
		},
		{
			item:      &Item{Name: BACKSTAGE_PASS, SellIn: 15, Quality: 30},
			strategy:  &BackstagePassItem{&Item{Name: BACKSTAGE_PASS, SellIn: 15, Quality: 30}},
			expectErr: false,
		},
		{
			item:      &Item{Name: "Conjured Item", SellIn: 3, Quality: 10},
			strategy:  &ConjuredItem{&Item{Name: "Conjured Item", SellIn: 3, Quality: 10}},
			expectErr: false,
		},
		{
			item:      &Item{Name: "Regular Item", SellIn: 7, Quality: 25},
			strategy:  &StandardItem{&Item{Name: "Regular Item", SellIn: 7, Quality: 25}},
			expectErr: false,
		},
	}

	for _, testCase := range testCases {
		strategy := testCase.item.GetUpdateStrategy()
		assert.Equal(testCase.strategy, strategy, "Mismatched strategy for item "+testCase.item.Name)
	}
}

func TestItem_increaseQuality(t *testing.T) {
	assert := assert.New(t)

	item := &Item{Name: "Test Item", SellIn: 5, Quality: 40}
	item.increaseQuality()
	assert.Equal(41, item.Quality, "Incorrect increased quality")

	// Testing quality doesn't go over 50
	item.Quality = 50
	item.increaseQuality()
	assert.Equal(50, item.Quality, "Quality increased over 50")
}

func TestItem_decreaseQuality(t *testing.T) {
	assert := assert.New(t)

	item := &Item{Name: "Test Item", SellIn: 5, Quality: 5}
	item.decreaseQuality()
	assert.Equal(4, item.Quality, "Incorrect decreased quality")

	// Testing quality doesn't go below 0
	item.Quality = 0
	item.decreaseQuality()
	assert.Equal(0, item.Quality, "Quality decreased below 0")
}
