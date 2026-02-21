package http

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"

	appLogger "github.com/NikolayNam/collabsphere-go/internal/logger"
)

func Register(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID: "create-user",
		Method:      http.MethodPost,
		Path:        "/users",
		Summary:     "Create user",
	}, createUserHandler)
}

func createUserHandler(ctx context.Context, input *CreateUserInput) (*UserResponse, error) {
	log := appLogger.From(ctx)
	log.Info("creating user", "email", input.Body.Email)

	resp := &UserResponse{}
	resp.Body.Email = input.Body.Email
	resp.Body.FirstName = input.Body.FirstName
	resp.Body.LastName = input.Body.LastName
	resp.Body.Phone = input.Body.Phone
	resp.Body.Role = "user"
	resp.Body.IsActive = true

	log.Info("user created", "email", resp.Body.Email)
	return resp, nil
}
