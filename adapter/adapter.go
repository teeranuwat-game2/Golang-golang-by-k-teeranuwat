package adapter

type Adapter struct {
}

func NewAdapter() *Adapter {
	return &Adapter{}
}

func (a *Adapter) GetUsername(id string) string {
	// Simulate a database lookup
	usernames := map[string]string{
		"1": "Alice",
		"2": "Bob",
		"3": "Charlie",
	}

	if username, exists := usernames[id]; exists {
		return username
	}
	return "Unknown User"
}
