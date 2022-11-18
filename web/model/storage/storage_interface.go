package storage

import (
	"github.com/virenmparmar/simple-distributed-project/web/model/storage/memory/user"
)

// Create a user repository if it does not exists
func GetUserRepo() *user.UserRepo {
	return &user.UserRepo{}
}
