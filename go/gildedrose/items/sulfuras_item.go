package items

type SulfurasItem struct {
	*Item
}

func (item *SulfurasItem) Update() {
	// “Sulfuras”, being a legendary item, never has to be sold or decreases in Quality
}
