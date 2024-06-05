package interface_hash_notifSvc

type IhashPassword interface {
	HashPassword(password string) string
	CompairPassword(hashedPassword string, plainPassword string) error
}
