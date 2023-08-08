package commands

// Define commands for code-tanks
const (
	NONE                           = 0
	MOVE_FORWARD                   = 1 << 0
	MOVE_BACKWARD                  = 1 << 1
	ROTATE_TANK_CLOCKWISE          = 1 << 2
	ROTATE_TANK_COUNTER_CLOCKWISE  = 1 << 3
	FIRE                           = 1 << 4
	ROTATE_GUN_CLOCKWISE           = 1 << 5
	ROTATE_GUN_COUNTER_CLOCKWISE   = 1 << 6
	ROTATE_RADAR_CLOCKWISE         = 1 << 7
	ROTATE_RADAR_COUNTER_CLOCKWISE = 1 << 8
	LOCK_GUN                       = 1 << 9
	UNLOCK_GUN                     = 1 << 10
	LOCK_RADAR                     = 1 << 11
	UNLOCK_RADAR                   = 1 << 12
)