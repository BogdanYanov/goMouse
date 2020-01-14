package mouse

// Clicker is the interface that wraps mouse buttons functions Up and Down.
type Clicker interface {
	Up()
	Down()
}

// Button implements Up, Down and contains a button pressed state
type Button struct {
	BtnPressed bool `json:"is_pressed"`
}

// Down simulates a button press
func (btn *Button) Down() {
	btn.BtnPressed = true
}

// Up simulates a button release
func (btn *Button) Up() {
	btn.BtnPressed = false
}
