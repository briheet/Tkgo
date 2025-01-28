package types

type User struct {
	UserName string `json:"username"`
	UserId   string `json:"userid"`
}

type MapData struct {
	UserInfo  User
	TokenData map[string]int
}

type UserTokenRequest struct {
	UserInfo User
	TokenID  string `json:"tokenID"`
}

type NonPresistentMap struct {
	TokenMap map[string]map[string]int
}
