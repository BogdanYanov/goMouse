package mouse

type Clicker interface {
	Up()
	Down()
}

type Button struct {
	BtnPressed bool
}

func (btn *Button) Down() {
	btn.BtnPressed = true
}

func (btn *Button) Up() {
	btn.BtnPressed = false
}


