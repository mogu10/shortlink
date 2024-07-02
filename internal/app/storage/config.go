package storage

type Storage interface {
	SaveLinkToStge(hash string, body []byte) error
	GetLinkFromStge(hash []byte) ([]byte, error)
}

func InitDefaultStorage() (*DefaultStorage, error) {
	return &DefaultStorage{}, nil
}

func InitFileStorage(flag string) (*FileStorage, error) {
	return newFileStorage(flag)
}
