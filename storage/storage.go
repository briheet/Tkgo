package storage

import "github.com/briheet/tkgo/types"

func NewStorage() *types.NonPresistentMap {
	return &types.NonPresistentMap{
		Map: make(map[string]types.UserData),
	}
}
