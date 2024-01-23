package database

type UserStore struct {
	db *Service
}

func NewUserStore(db *Service) *UserStore {
	return &UserStore{
		db: db,
	}
}

// func (s *UserStore) InsertUser() {
// 	s.db.DB.Query("")
// }
