package manager

type Manager struct {
	managerID string
	location  string
	resources map[string]int
}

// Initialize with Values
func NewManager(managerID string, location string) *Manager {
	instance := &Manager{
		managerID: managerID,
		location:  location,
		resources: make(map[string]int),
	}
	return instance
}

// Getters
func (m Manager) GetManagerID() string {
	return m.managerID
}

func (m Manager) GetLocation() string {
	//modify the value of location
	return m.location
}

func (m Manager) GetNumebrOfProducts(NameofProduct string) int {
	if quantity, ok := m.resources[NameofProduct]; ok {
		return quantity
	}
	//return -1 if the product is not found
	return -1
}
