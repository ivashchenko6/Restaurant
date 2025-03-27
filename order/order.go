package order

type Order struct {
	TableID int
	Dishes  map[string]int
	Check   float64
	Status  string
}
