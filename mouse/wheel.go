package mouse

type Scroller interface {
	ScrollUp()
	ScrollDown()
}

type Wheel struct {
	scrollVal uint8
	sensitivity uint8
}

func (w *Wheel) WheelSensitivity(val uint8){
	if val > 10 {
		w.sensitivity = 10
	} else if val > 0 {
		w.sensitivity = val
	}
}

func (w *Wheel) ScrollUp() {
	if w.scrollVal + w.sensitivity > 10 {
		w.scrollVal = 10
	} else {
		w.scrollVal += w.sensitivity
	}
}

func (w *Wheel) ScrollDown() {
	if w.scrollVal - w.sensitivity < 1 {
		w.scrollVal = 1
	} else {
		w.scrollVal -= w.sensitivity
	}
}