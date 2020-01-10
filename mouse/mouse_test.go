package mouse

import "testing"

type screen struct {
	width, height uint32
}

var width, height uint32 = 1024, 768

var testInfoStr = `Mouse information:
X position - 512
Y position - 384
Sensitivity - 5
Is left button pressed? - true
Is right button pressed? - false
Wheel sensitivity - 4
Scroll value - 9`

func TestMouse_Creating(t *testing.T) {
	myScreen := screen{width, height}
	myMouse := NewMouse(myScreen.width, myScreen.height)
	if myMouse.posX != myScreen.width / 2 || myMouse.posY != myScreen.height / 2 {
		t.Errorf("Initial position of mouse wrong. Expected X - %d, Y - %d; Get X - %d, Y - %d",
			myScreen.width / 2,
			myScreen.height / 2,
			myMouse.posX,
			myMouse.posY)
	}
}

func TestMouse_Move(t *testing.T) {
	myScreen := screen{width, height}
	myMouse := NewMouse(myScreen.width, myScreen.height)
	myMouse.Move(1260, 1024, myScreen.width, myScreen.height)
	if myMouse.posX != myScreen.width || myMouse.posY != myScreen.height {
		t.Errorf("Coords out of screen. Expected X - %d, Y - %d; Get X - %d, Y - %d",
			myScreen.width,
			myScreen.height,
			myMouse.posX,
			myMouse.posY)
	}
}

func TestMouse_Sensitivity(t *testing.T) {
	myScreen := screen{width, height}
	myMouse := NewMouse(myScreen.width, myScreen.height)
	myMouse.Sensitivity(11)
	if myMouse.sensitivity != 10 {
		t.Errorf("Sensitivity out of range. Expected sensitivity - %d, get - %d", 10, myMouse.sensitivity)
	}
}

func TestMouse_WheelSensitivity(t *testing.T) {
	myScreen := screen{width, height}
	myMouse := NewMouse(myScreen.width, myScreen.height)
	myMouse.WheelSensitivity(11)
	if myMouse.Wheel.sensitivity != 10 {
		t.Errorf("Wheel sensitivity out of range. Expected sensitivity - %d, get - %d", 10, myMouse.Wheel.sensitivity)
	}
}

func TestMouse_WheelScrollUp(t *testing.T) {
	myScreen := screen{width, height}
	myMouse := NewMouse(myScreen.width, myScreen.height)
	myMouse.WheelSensitivity(3)
	myMouse.ScrollUp()
	if myMouse.Wheel.scrollVal != 8 {
		t.Errorf("Error scrolling. Scroll value expected - %d, get - %d", 8, myMouse.Wheel.scrollVal)
	}
}

func TestMouse_WheelScrollDown(t *testing.T) {
	myScreen := screen{width, height}
	myMouse := NewMouse(myScreen.width, myScreen.height)
	myMouse.WheelSensitivity(3)
	myMouse.ScrollDown()
	if myMouse.Wheel.scrollVal != 2 {
		t.Errorf("Error scrolling. Scroll value expected - %d, get - %d", 2, myMouse.Wheel.scrollVal)
	}
}
