package entities

type Employee struct {
	Name     string
	Age      byte
	Location Location
	Gender   string
}

func (employee Employee) GetCity() string {
	return employee.Location.GetCity()
}
