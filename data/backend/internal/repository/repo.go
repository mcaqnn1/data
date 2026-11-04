package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type Repo struct {
	db *sqlx.DB
}

func NewRepo(db *sqlx.DB) *Repo {
	return &Repo{db: db}
}

type LoginData struct {
	ID       int    `json:"id"`
	UName    string `json:"user_name"`
	Password string `json:"password"`
}

func (r *Repo) Login(ctx context.Context,d LoginData) error  {
	query := `
		INSERT INTO users (user_name,password) 
		VALUES ($1,$2);
	`

	_,err := r.db.ExecContext(ctx,query,d.UName,d.Password)
	if err != nil {
		return err
	}

	return nil
}