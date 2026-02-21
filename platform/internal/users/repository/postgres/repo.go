package postgres

import (
	"context"
	"errors"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/NikolayNam/collabsphere-go/internal/users/domain"
	"github.com/NikolayNam/collabsphere-go/internal/users/service"
)

type Repo struct {
	pool *pgxpool.Pool
}

func New(pool *pgxpool.Pool) *Repo {
	return &Repo{pool: pool}
}

func (r *Repo) ExistsByEmail(ctx context.Context, organizationID, email string) (bool, error) {
	const q = `
SELECT EXISTS(
  SELECT 1
  FROM users
  WHERE organization_id = $1 AND email = $2
)`
	var exists bool
	err := r.pool.QueryRow(ctx, q, organizationID, email).Scan(&exists)
	return exists, err
}

func (r *Repo) Create(ctx context.Context, u *domain.User) error {
	const q = `
INSERT INTO users (
  organization_id, email, password_hash, first_name, last_name, phone, role, is_active
) VALUES (
  $1,$2,$3,$4,$5,$6,$7,$8
)
RETURNING id
`
	err := r.pool.QueryRow(
		ctx, q,
		u.OrganizationID, u.Email, u.PasswordHash, u.FirstName, u.LastName, u.Phone, u.Role, u.IsActive,
	).Scan(&u.ID)

	if err == nil {
		return nil
	}

	// Unique violation -> conflict
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgErr.Code == "23505" {
		return service.ErrConflict
	}

	return err
}
