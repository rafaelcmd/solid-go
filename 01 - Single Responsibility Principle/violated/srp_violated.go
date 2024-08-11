package main

import (
	"fmt"
)

// Single Responsibility Principle (SRP)
// A class should have only one reason to change, meaning that a class should have
// only one job.

// Violation of SRP
// The following example violates the SRP because the UserService struct is responsible
// for both user authentication and data storage. This violates the SRP because the
// UserService struct has two reasons to change: if the authentication logic changes or
// if the data storage logic changes. It would be better to separate the authentication
// logic from the data storage logic.

// User struct to hold user data
type User struct {
	Username string
	Password string
}

// UserService struct handles both authentication and data storage
type UserService struct {
	users map[string]string
}

// NewUserService initializes the UserService
func NewUserService() *UserService {
	return &UserService{
		users: make(map[string]string),
	}
}

// Authenticate handles user authentication (violates SRP)
func (s *UserService) Authenticate(username, password string) bool {
	if pass, ok := s.users[username]; ok {
		return pass == password
	}
	return false
}

// SaveUser handles saving user data (violates SRP)
func (s *UserService) SaveUser(username, password string) {
	s.users[username] = password
}

func main() {
	service := NewUserService()
	service.SaveUser("john_doe", "password123")

	authenticated := service.Authenticate("john_doe", "password123")
	fmt.Println("Authenticated:", authenticated)
}
