package user

import (
	"database/sql"
	"fmt"
	"github.com/muhammadjon1304/jwt-authentication/cmd/types"
	"log"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE email = $1")
	if err != nil {
		return nil, err
	}
	u := new(types.User)
	for rows.Next() {
		u, err = scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}
	if u.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}
	return u, nil
}
func scanRowIntoUser(rows *sql.Rows) (*types.User, error) {
	user := new(types.User)

	err := rows.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Password,
		&user.CreatedAt,
	)
	if err != nil {
		log.Fatal(err)
	}
	return user, err

}

func (s *Store) GetUserByID(id int) (*types.User, error) {
	return nil, nil
}
func (s *Store) CreateUser(types.User) error {
	return nil
}
