package main

// Liskov Substitution Principle (LSP)
// Objects in a program should be replaceable with instances of their subtypes without
// altering the correctness of that program.

// Violation of LSP
// The following example violates the LSP because the AdvancedUserService struct is not
// a subtype of the UserService struct. The AdvancedUserService struct does not extend the
// UserService struct, but instead embeds the UserServiceInterface interface. This violates
// the LSP because the AdvancedUserService struct is not a subtype of the UserService struct.

// User struct to hold user data
type User struct {
	Username string
	Password string
}

// UserServiceInterface defines the methods for user services
type UserServiceInterface interface {
	Authenticate(username, password string) bool
	SaveUser(username, password string)
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

// AdvancedUserService struct embeds the UserServiceInterface interface
type AdvancedUserService struct {
	UserServiceInterface
}

// NewAdvancedUserService initializes the AdvancedUserService
func NewAdvancedUserService(service UserServiceInterface) *AdvancedUserService {
	return &AdvancedUserService{
		UserServiceInterface: service,
	}
}

// New feature: LogAuthentication logs the authentication attempt
func (a *AdvancedUserService) LogAuthentication(username string, success bool) {
	println("User:", username, "Authenticated:", success)
}

func main() {
	service := NewUserService()
	service.SaveUser("john_doe", "password123")

	advancedService := NewAdvancedUserService(service)
	authenticated := advancedService.Authenticate("john_doe", "password123")
	advancedService.LogAuthentication("john_doe", authenticated)
}
