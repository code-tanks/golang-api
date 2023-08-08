package tanks

import (
	"codetanks/pkg/commands"
	"sync"
)

// Tank implements the BaseTank interface
type Tank struct {
	Mu       sync.Mutex
	Commands []int
}

func (t *Tank) Run() {
	t.Mu.Lock()
	defer t.Mu.Unlock()

	t.Commands = append(t.Commands, commands.MOVE_FORWARD)
	// Add other commands as needed
}

func (t *Tank) OnEvent(event interface{}) {
	t.Mu.Lock()
	defer t.Mu.Unlock()

	t.Commands = append(t.Commands, commands.MOVE_BACKWARD)
	// Handle the event and add appropriate commands
}

func CreateTank() *Tank {
	// Initialize and return a new Tank instance
	return &Tank{
		Commands: make([]int, 0),
	}
}