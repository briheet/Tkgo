package types

type User struct {
	UserName string `json:"userName"`
	UserId   string `json:"userID"`
}

type MapData struct {
	UserInfo  User
	TokenData map[string]int
}

type UserTokenRequest struct {
	UserInfo User
	TokenID  string `json:"tokenID"`
}

type CreateUser struct {
	UserInfo       User
	SimulationTime int `json:"simulationTime"`
	TokenNumbers   int `json:"tokenNumbers"`
}
