package storage

import "errors"

var links = make(map[string]string)

func SaveLink(hash string, body []byte) error {
	links[string(body)] = hash

	return nil
}

func LoadLink(hash []byte) ([]byte, error) {

	h := string(hash)
	if h == "" {
		return nil, errors.New("empty hash")
	}

	for key, value := range links {
		if value == h {
			return []byte(key), nil
		}
	}

	return nil, errors.New("invalid path")
}
