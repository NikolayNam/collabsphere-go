package service

import (
	"context"
	"strings"

	"golang.org/x/crypto/bcrypt"

	"github.com/NikolayNam/collabsphere-go/internal/users/domain"
)

type Repository interface {
	Create(ctx context.Context, u *domain.User) error
	ExistsByEmail(ctx context.Context, organizationID, email string) (bool, error)
}

type Service struct {
	repo Repository
}

func New(repo Repository) *Service {
	return &Service{repo: repo}
}

type CreateUserCmd struct {
	OrganizationID string
	Email          string
	Password       string
	FirstName      string
	LastName       string
	Phone          string
	Role           string
}

func (s *Service) CreateUser(ctx context.Context, cmd CreateUserCmd) (*domain.User, error) {
	orgID := strings.TrimSpace(cmd.OrganizationID)
	email := strings.ToLower(strings.TrimSpace(cmd.Email))
	role := strings.TrimSpace(cmd.Role)
	if role == "" {
		role = "member"
	}

	if orgID == "" || email == "" || len(cmd.Password) < 6 {
		return nil, ErrValidation
	}
	if !strings.Contains(email, "@") {
		return nil, ErrValidation
	}

	exists, err := s.repo.ExistsByEmail(ctx, orgID, email)
	if err != nil {
		return nil, ErrInternal
	}
	if exists {
		return nil, ErrConflict
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(cmd.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, ErrInternal
	}

	u := &domain.User{
		OrganizationID: orgID,
		Email:          email,
		PasswordHash:   string(hash),
		FirstName:      strings.TrimSpace(cmd.FirstName),
		LastName:       strings.TrimSpace(cmd.LastName),
		Phone:          strings.TrimSpace(cmd.Phone),
		Role:           role,
		IsActive:       true,
	}

	if err := s.repo.Create(ctx, u); err != nil {
		// repo может вернуть ErrConflict (например, unique violation)
		if err == ErrConflict {
			return nil, ErrConflict
		}
		return nil, ErrInternal
	}

	return u, nil
}
