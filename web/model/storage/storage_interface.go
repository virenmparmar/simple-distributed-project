package storage

import (
	"github.com/simple-distributed-project/web/model/storage/memory"
)

// Create a user repository if it does not exists
func GetUserRepo() *memory.UserRepo {
	return &memory.UserRepo{}
}
