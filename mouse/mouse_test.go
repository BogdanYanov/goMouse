package mouse

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

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
Scroll value - 9
`

func TestMouse_Creating(t *testing.T) {
	myScreen := screen{width, height}
	myMouse := NewMouse(myScreen.width, myScreen.height)
	if myMouse.PosX != myScreen.width/2 || myMouse.PosY != myScreen.height/2 {
		t.Errorf("Initial position of mouse wrong. Expected X - %d, Y - %d; Get X - %d, Y - %d",
			myScreen.width/2,
			myScreen.height/2,
			myMouse.PosX,
			myMouse.PosY)
	}
}

func TestMouse_Move(t *testing.T) {
	myScreen := screen{width, height}
	myMouse := NewMouse(myScreen.width, myScreen.height)
	myMouse.Move(1260, 1024, myScreen.width, myScreen.height)
	if myMouse.PosX != myScreen.width || myMouse.PosY != myScreen.height {
		t.Errorf("Coords out of screen. Expected X - %d, Y - %d; Get X - %d, Y - %d",
			myScreen.width,
			myScreen.height,
			myMouse.PosX,
			myMouse.PosY)
	}
	myMouse.Sensitivity(10)
	myMouse.Move(968, 743, myScreen.width, myScreen.height)
	if myMouse.PosX != 968 || myMouse.PosY != 743 {
		t.Errorf("Coords out of screen. Expected X - %d, Y - %d; Get X - %d, Y - %d",
			968,
			743,
			myMouse.PosX,
			myMouse.PosY)
	}
	myMouse.Move(1260, 1024, myScreen.width, myScreen.height)
	if myMouse.PosX != myScreen.width || myMouse.PosY != myScreen.height {
		t.Errorf("Coords out of screen. Expected X - %d, Y - %d; Get X - %d, Y - %d",
			myScreen.width,
			myScreen.height,
			myMouse.PosX,
			myMouse.PosY)
	}
}

func TestMouse_Sensitivity(t *testing.T) {
	myScreen := screen{width, height}
	myMouse := NewMouse(myScreen.width, myScreen.height)
	myMouse.Sensitivity(11)
	if myMouse.Sens != 10 {
		t.Errorf("Sensitivity out of range. Expected sensitivity - %d, get - %d", 10, myMouse.Sens)
	}
	myMouse.Sensitivity(0)
	if myMouse.Sens != 1 {
		t.Errorf("Sensitivity out of range. Expected sensitivity - %d, get - %d", 1, myMouse.Sens)
	}
}

func TestMouse_WheelScrollUp(t *testing.T) {
	myScreen := screen{width, height}
	myMouse := NewMouse(myScreen.width, myScreen.height)
	myMouse.ScrollVal = 7
	myMouse.ScrollUp()
	if myMouse.Wheel.ScrollVal != 8 {
		t.Errorf("Error scrolling. Scroll value expected - %d, get - %d", 8, myMouse.Wheel.ScrollVal)
	}
	for i := 0; i < 4; i++ {
		myMouse.ScrollUp()
	}
	if myMouse.Wheel.ScrollVal != 10 {
		t.Errorf("Error scrolling. Scroll value expected - %d, get - %d", 10, myMouse.Wheel.ScrollVal)
	}
}

func TestMouse_WheelScrollDown(t *testing.T) {
	myScreen := screen{width, height}
	myMouse := NewMouse(myScreen.width, myScreen.height)
	myMouse.ScrollVal = 3
	myMouse.ScrollDown()
	if myMouse.Wheel.ScrollVal != 2 {
		t.Errorf("Error scrolling. Scroll value expected - %d, get - %d", 2, myMouse.Wheel.ScrollVal)
	}
	myMouse.ScrollDown()
	myMouse.ScrollDown()
	if myMouse.Wheel.ScrollVal != 1 {
		t.Errorf("Error scrolling. Scroll value expected - %d, get - %d", 1, myMouse.Wheel.ScrollVal)
	}
}

func TestMouse_Click(t *testing.T) {
	myScreen := screen{width, height}
	myMouse := NewMouse(myScreen.width, myScreen.height)
	myMouse.LeftBtnDown()
	if myMouse.LeftBtn.BtnPressed != true {
		t.Errorf("Error button click. Expected - %v, got - %v", true, myMouse.LeftBtn.BtnPressed)
	}
	myMouse.LeftBtnUp()
	if myMouse.LeftBtn.BtnPressed != false {
		t.Errorf("Error button click. Expected - %v, got - %v", false, myMouse.LeftBtn.BtnPressed)
	}
	myMouse.RightBtnDown()
	if myMouse.RightBtn.BtnPressed != true {
		t.Errorf("Error button click. Expected - %v, got - %v", true, myMouse.RightBtn.BtnPressed)
	}
	myMouse.RightBtnUp()
	if myMouse.RightBtn.BtnPressed != false {
		t.Errorf("Error button click. Expected - %v, got - %v", false, myMouse.RightBtn.BtnPressed)
	}

}

func TestMouse_Info(t *testing.T) {
	myScreen := screen{width, height}
	myMouse := NewMouse(myScreen.width, myScreen.height)
	myMouse.Sensitivity(5)
	myMouse.LeftBtnDown()
	myMouse.ScrollVal = 8
	myMouse.ScrollUp()

	oldOutput := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	myMouse.Info()

	w.Close()
	os.Stdout = oldOutput

	var buf bytes.Buffer
	io.Copy(&buf, r)

	if equal := strings.Compare(testInfoStr, buf.String()); equal != 0 {
		t.Errorf("Error info output. Expected:\n%s\nGot:\n%s", testInfoStr, buf.String())
	}

}
