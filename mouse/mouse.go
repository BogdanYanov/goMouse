package mouse

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Mouse struct {
	PosX        uint32 `json:"pos_x"`
	PosY        uint32 `json:"pos_y"`
	LeftBtn     Button `json:"l_button"`
	RightBtn    Button `json:"r_button"`
	Sens 		uint8  `json:"sensitivity"`
	Wheel       `json:"wheel"`
}

func NewMouse(monWidth, monHeight uint32) *Mouse {
	mouse := &Mouse{
		monWidth / 2,
		monHeight / 2,
		Button{},
		Button{},
		1,
		Wheel{5},
	}
	return mouse
}

func (m *Mouse) Move(x, y, monWidth, monHeight uint32) {
	if x > monWidth {
		x = monWidth
	}
	if y > monHeight {
		y = monHeight
	}
Loop:
	for {
		if x > m.PosX {
			if m.PosX+uint32(m.Sens) <= x {
				m.PosX += uint32(m.Sens)
			} else {
				m.PosX += x - m.PosX
			}
		} else {
			if m.PosX-uint32(m.Sens) >= x {
				m.PosX -= uint32(m.Sens)
			} else {
				m.PosX -= m.PosX - x
			}
		}

		if y > m.PosY {
			if m.PosY+uint32(m.Sens) <= y {
				m.PosY += uint32(m.Sens)
			} else {
				m.PosY += y - m.PosY
			}
		} else {
			if m.PosY-uint32(m.Sens) >= y {
				m.PosY -= uint32(m.Sens)
			} else {
				m.PosY -= m.PosY - y
			}
		}

		if m.PosX == x && m.PosY == y {
			break Loop
		}
	}
	m.WriteJSON()
}

func (m *Mouse) Sensitivity(val uint8) {
	if val > 10 {
		m.Sens = 10
	} else if val == 0 {
		m.Sens = 1
	} else {
		m.Sens = val
	}
	m.WriteJSON()
}

func (m *Mouse) LeftBtnDown() {
	m.LeftBtn.Down()
	m.WriteJSON()
}

func (m *Mouse) RightBtnDown() {
	m.RightBtn.Down()
	m.WriteJSON()
}

func (m *Mouse) LeftBtnUp() {
	m.LeftBtn.Up()
	m.WriteJSON()
}

func (m *Mouse) RightBtnUp() {
	m.RightBtn.Up()
	m.WriteJSON()
}

func (m *Mouse) ScrollUp() {
	if m.Wheel.ScrollVal != 10 {
		m.Wheel.ScrollVal++
		m.WriteJSON()
	}
}

func (m *Mouse) ScrollDown() {
	if m.Wheel.ScrollVal != 1 {
		m.Wheel.ScrollVal--
		m.WriteJSON()
	}
}

func (m *Mouse) Info() {
	fmt.Printf("Mouse information:\n"+
		"X position - %d\n"+
		"Y position - %d\n"+
		"Sensitivity - %d\n"+
		"Is left button pressed? - %v\n"+
		"Is right button pressed? - %v\n"+
		"Scroll value - %d\n",
		m.PosX, m.PosY, m.Sens, m.LeftBtn.BtnPressed, m.RightBtn.BtnPressed, m.Wheel.ScrollVal)
}

func (m *Mouse) Reset(monWidth, monHeight uint32) {
	m.PosX = monWidth / 2
	m.PosY = monHeight / 2
	m.LeftBtnUp()
	m.RightBtnUp()
	m.Wheel.ScrollVal = 5
	m.Sensitivity(1)
	m.WriteJSON()
}

func GetMouse(m *Mouse) {
	file, _ := os.OpenFile("mouse.json", os.O_RDWR, 0755)
	defer file.Close()
	jsonMouse, _ := ioutil.ReadAll(file)
	json.Unmarshal(jsonMouse, m)
}

func (m *Mouse) WriteJSON() {
	var flags int
	if FileExists() {
		os.Truncate("mouse.json", 0)
		flags = os.O_WRONLY
	} else {
		flags = os.O_WRONLY|os.O_CREATE
	}
	file, _ := os.OpenFile("mouse.json", flags, 0755)
	defer file.Close()
	json.NewEncoder(file).Encode(m)
}

func FileExists() bool {
	info, err := os.Stat("mouse.json")
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}


