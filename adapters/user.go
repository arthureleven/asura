package adapters

import "github.com/uptrace/bun"

type UserModel interface {
}

type UserAdapter struct {
	DB *bun.DB
}
