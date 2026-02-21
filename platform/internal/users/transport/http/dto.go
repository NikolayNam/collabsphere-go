package http

// POST /users (в рамках X-Organization-ID)
type CreateUserInput struct {
	Body struct {
		Email     string `json:"email" required:"true" format:"email"`
		Password  string `json:"password" required:"true" minLength:"6"`
		FirstName string `json:"first_name,omitempty"`
		LastName  string `json:"last_name,omitempty"`
		Phone     string `json:"phone,omitempty"`
		Role      string `json:"role,omitempty" doc:"Role in this organization (optional, default 'member')"`
	}
}

// PUT /users/{user_id} (в рамках X-Organization-ID)
type UpdateUserInput struct {
	UserID uint `path:"user_id" doc:"User ID"`

	Body struct {
		FirstName *string `json:"first_name,omitempty"`
		LastName  *string `json:"last_name,omitempty"`
		Phone     *string `json:"phone,omitempty"`
		Role      *string `json:"role,omitempty" doc:"Update role in this organization"`
		IsActive  *bool   `json:"is_active,omitempty" doc:"Deactivate membership in this organization"`
	}
}

type UserResponse struct {
	Body struct {
		ID        uint   `json:"id"`
		Email     string `json:"email"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Phone     string `json:"phone"`

		OrganizationID string `json:"organization_id"`
		Role           string `json:"role"`
		IsActive       bool   `json:"is_active"`
	}
}
