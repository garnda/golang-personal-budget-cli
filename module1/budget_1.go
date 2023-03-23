package module1

// Budget stores budget information
type Budget struct {
	Year     int
	Income   float64
	Expenses float64
	Max      float32
	Items    []Item
}

// Item stores item information
type Item struct {
	Description string
	Price       float32
}
