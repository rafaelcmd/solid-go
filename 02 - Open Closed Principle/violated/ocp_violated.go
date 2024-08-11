package main

// Open Closed Principle (OCP)
// Software entities should be open for extension, but closed for modification.

// Violation of OCP
// The following example violates the OCP because the UserService struct is not open for
// extension. If we wanted to add a new feature to the UserService struct, we would have
// to modify the UserService struct itself. This violates the OCP because we should be able
// to add new features to the UserService struct without modifying the UserService struct.
// It would be better to create a new struct that extends the UserService struct and adds
// the new feature.

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

// UserAuthenticator struct extends the UserService struct and adds user authentication
type UserAuthenticator struct {
	*UserService
}

// NewUserAuthenticator initializes the UserAuthenticator
func NewUserAuthenticator(service *UserService) *UserAuthenticator {
	return &UserAuthenticator{
		UserService: service,
	}
}

// Authenticate handles user authentication
func (a *UserAuthenticator) Authenticate(username, password string) bool {
	return a.UserService.Authenticate(username, password)
}

func main() {
	service := NewUserService()
	service.SaveUser("john_doe", "password123")

	authenticator := NewUserAuthenticator(service)
	authenticated := authenticator.Authenticate("john_doe", "password123")
	println("Authenticated:", authenticated)
}
