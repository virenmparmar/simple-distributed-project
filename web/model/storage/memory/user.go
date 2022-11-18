package memory

// User represents a user in the system.
type User struct {
	ID       int
	Username string
	Password string
}

// UserStore represents a store of users.
type UserStore struct {
	users map[int]*User
}
