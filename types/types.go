package types

type User struct {
	UserName string `json:"userName"`
	UserId   string `json:"userId"`
}

type MapData struct {
	UserInfo  User
	TokenData map[string]int
}

type CreateUser struct {
	UserInfo       User
	SimulationTime int `json:"simulationTime"`
	TokenNumbers   int `json:"tokenNumbers"`
}
