package magazine

type Subscriber struct {
	Name   string
	Rate   float64
	Active bool
	// HomeAddress Address
	// 익명필드(anonymous field)로 변경
	Address
}

type Employee struct {
	Name   string
	Salary float64
	// HomeAddress Address
	Address
}

type Address struct {
	Street     string
	City       string
	State      string
	PostalCode string
}
