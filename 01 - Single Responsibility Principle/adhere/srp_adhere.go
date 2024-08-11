package main

// Single Responsibility Principle (SRP)
// A class should have only one reason to change, meaning that a class should have
// only one job.

// Adhere to SRP
// The following example adheres to the SRP because the UserService struct is responsible
// for user authentication and the UserStorage struct is responsible for data storage.
// This adheres to the SRP because the UserService struct only has one reason to change:
// if the authentication logic changes. The UserStorage struct only has one reason to
// change: if the data storage logic changes.

// User struct to hold user data
type User struct {
	Username string
	Password string
}

// UserService struct handles user authentication
type UserService struct {
	users map[string]string
}

// NewUserService initializes the UserService
func NewUserService() *UserService {
	return &UserService{
		users: make(map[string]string),
	}
}

// Authenticate handles user authentication
func (s *UserService) Authenticate(username, password string) bool {
	if pass, ok := s.users[username]; ok {
		return pass == password
	}
	return false
}

// UserStorage struct handles user data storage
type UserStorage struct {
	users map[string]string
}

// NewUserStorage initializes the UserStorage
func NewUserStorage() *UserStorage {
	return &UserStorage{
		users: make(map[string]string),
	}
}

// SaveUser handles saving user data
func (s *UserStorage) SaveUser(username, password string) {
	s.users[username] = password
}

func main() {
	userStorage := NewUserStorage()

	userStorage.SaveUser("john_doe", "password123")

	userService := NewUserService()

	authenticated := userService.Authenticate("john_doe", "password123")

	println("Authenticated:", authenticated)

}
