package mouse

// Scroller is the interface that wraps ScrollUp and ScrollDown methods
type Scroller interface {
	ScrollUp()
	ScrollDown()
}

// Wheel contains scroll value of mouse wheel.
type Wheel struct {
	ScrollVal uint8 `json:"scroll_value"`
}
