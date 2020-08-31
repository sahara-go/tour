package store

type Store interface {
	Save()
	SetHost(host string)
}

type BaseStore struct {
	Host string
}

func (s *BaseStore) SetHost(host string) {
	s.Host = host
}
