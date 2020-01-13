package mouse

type Scroller interface {
	ScrollUp()
	ScrollDown()
}

type Wheel struct {
	ScrollVal   uint8 `json:"scroll_value"`
}