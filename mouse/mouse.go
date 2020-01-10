package mouse

import "fmt"

type Mouse struct {
	posX, posY        uint32
	leftBtn, rightBtn Button
	sensitivity       uint8
	Wheel
}

func NewMouse(monWidth, monHeight uint32) *Mouse {
	return &Mouse {
		monWidth / 2,
		monHeight / 2,
		Button{},
		Button{},
		1,
		Wheel{5, 1},
	}
}

func (m *Mouse) Move(x, y, monWidth, monHeight uint32) {
	switch {
	case x > monWidth:
		x = monWidth
	case y > monHeight:
		y = monHeight
	}
	Loop:
		for {
			if x > m.posX {
				if m.posX + uint32(m.sensitivity) <= x {
					m.posX += uint32(m.sensitivity)
				} else {
					m.posX += x - m.posX
				}
			} else {
				if m.posX - uint32(m.sensitivity) >= x {
					m.posX -= uint32(m.sensitivity)
				} else {
					m.posX -= m.posX - x
				}
			}

			if y > m.posY {
				if m.posY + uint32(m.sensitivity) <= y {
					m.posY += uint32(m.sensitivity)
				} else {
					m.posY += y - m.posY
				}
			} else {
				if m.posY - uint32(m.sensitivity) >= y {
					m.posY -= uint32(m.sensitivity)
				} else {
					m.posY -= m.posY - y
				}
			}

			if x == m.posX && y == m.posY {
				break Loop
			}
		}
}

func (m *Mouse) Sensitivity(val uint8) {
	if val > 10 {
		m.sensitivity = 10
	} else if val > 0 {
		m.sensitivity = val
	}
}

func (m *Mouse) LeftBtnDown() {
	m.leftBtn.Down()
}

func (m *Mouse) RightBtnDown() {
	m.rightBtn.Down()
}

func (m *Mouse) LeftBtnUp() {
	m.leftBtn.Up()
}

func (m *Mouse) RightBtnUp() {
	m.rightBtn.Up()
}

func (m *Mouse) Info() {
	fmt.Printf("Mouse information:\n" +
		"X position - %d\n" +
		"Y position - %d\n" +
		"Sensitivity - %d\n" +
		"Is left button pressed? - %v\n" +
		"Is right button pressed? - %v\n" +
		"Wheel sensitivity - %d\n" +
		"Scroll value - %d",
		m.posX, m.posY, m.sensitivity, m.leftBtn.BtnPressed, m.rightBtn.BtnPressed, m.Wheel.sensitivity, m.Wheel.scrollVal)
}


