package access_token

type Maker interface {
	CreateToken(userId int64, expiredAt int64) (string, error)
	VerifyToken(token string) (*Payload, error)
}
