package domain

type User struct {
	ID             uint
	OrganizationID string

	Email        string
	PasswordHash string

	FirstName string
	LastName  string
	Phone     string

	Role     string
	IsActive bool
}
