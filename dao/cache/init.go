package cache

type Cache struct {
	*User
	*Sign
	*Prize
	*Group
	*Choose
	*Activity
}

func NewCache() *Cache {
	return &Cache{
		newUser(),
		newSign(),
		newPrize(),
		newGroup(),
		newChoose(),
		newActivity(),
	}
}
