package main

// Liskov Substitution Principle (LSP)
// Objects in a program should be replaceable with instances of their subtypes without
// altering the correctness of that program.

// Adhere to LSP
// The following example adheres to the LSP because the AdvancedUserService struct extends
// the UserService struct. The AdvancedUserService struct is a subtype of the UserService
// struct, which means that the AdvancedUserService struct can be used in place of the
// UserService struct without altering the correctness of the program.
// The AdvancedUserService struct extends the UserService struct by adding a new method,
// ResetPassword, which allows users to reset their password.
// The AdvancedUserService struct also embeds the UserService struct, which allows the
// AdvancedUserService struct to use the methods from the UserService struct.

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

// Authenticate handles user authentication
func (s *UserService) Authenticate(username, password string) bool {
	if pass, ok := s.users[username]; ok {
		return pass == password
	}
	return false
}

// SaveUser handles saving user data
func (s *UserService) SaveUser(username, password string) {
	s.users[username] = password
}

// AdvancedUserService struct extends the UserService struct with additional features
type AdvancedUserService struct {
	*UserService
}

// NewAdvancedUserService initializes the AdvancedUserService
func NewAdvancedUserService() *AdvancedUserService {
	return &AdvancedUserService{
		UserService: NewUserService(),
	}
}

// ResetPassword handles resetting a user's password
func (s *AdvancedUserService) ResetPassword(username, password string) {
	s.SaveUser(username, password)
}

func main() {
	service := NewAdvancedUserService()

	service.SaveUser("john_doe", "password123")

	service.ResetPassword("john_doe", "new_password")

	authenticated := service.Authenticate("john_doe", "new_password")

	println("Authenticated:", authenticated)
}
