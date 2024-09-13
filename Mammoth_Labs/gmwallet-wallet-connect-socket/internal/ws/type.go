package ws

type CreateRoomRes struct {
	Hash   string `json:"hash"`   // ROOM Number
	Expiry string `json:"expiry"` // Expired
	UUID   string `json:"uuid"`
	Jwt    string `json:"jwt"` // JWT
}

type JoinRoomReq struct {
	Hash string `param:"hash"` // ROOM Number
	UUID string `query:"uuid"` // uuid
	Jwt  string `query:"jwt"`  // JWT
}
