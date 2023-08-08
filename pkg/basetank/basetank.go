package basetank

// BaseTank defines the structure of a tank object
type BaseTank interface {
	Run()
	OnEvent(event interface{})
}
