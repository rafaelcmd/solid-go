package main

// Open Closed Principle (OCP)
// Software entities should be open for extension, but closed for modification.

// Adhere to OCP
// The following example adheres to the OCP because the UserService struct is open for
// extension. If we wanted to add a new feature to the UserService struct, we would not
// have to modify the UserService struct itself. We could create a new struct that extends
// the UserService struct and adds the new feature.

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

// AdvancedUserService struct extends UserService with additional features
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
	println("Authenticated:", authenticated)
}
