package repository

import (
	"database/sql"
	"github.com/muhammadjon1304/jwt-authentication/cmd/models"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

type AuthRepository struct {
	Db *sql.DB
}

func NewAuthRepository(db *sql.DB) AuthRepository {
	return AuthRepository{
		Db: db,
	}
}
func (a *AuthRepository) CreateUser(user models.User) bool {
	stmt, err := a.Db.Prepare("insert into users(firstname,lastname,email,password) values ($1,$2,$3,$4)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	password, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	_, err2 := stmt.Exec(user.Firstname, user.Lastname, user.Email, password)
	if err2 != nil {
		log.Fatal(err)
		return false
	}
	return true
}

func (a *AuthRepository) LoginUser(login_user models.Login_User) models.User {
	query, err := a.Db.Query("SELECT * FROM users where email=$1", login_user.Email)
	if err != nil {
		log.Fatal(err)
		return models.User{}
	}
	var dbUser models.User
	if query != nil {
		for query.Next() {
			var (
				id        int
				firstname string
				lastname  string
				email     string
				password  string
				createdat time.Time
			)
			err := query.Scan(&id, &firstname, &lastname, &email, &password, &createdat)
			if err != nil {
				log.Println(err)
			}
			dbUser = models.User{ID: id, Firstname: firstname, Lastname: lastname, Email: email, Password: password, CreatedAt: createdat}
		}
	}
	err3 := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(login_user.Password))
	if err3 != nil {
		log.Fatal(err)
	}
	return dbUser
}

func (a *AuthRepository) GetByEmail(email string) models.User {
	query, err := a.Db.Query("SELECT * FROM users where email=$1", email)
	if err != nil {
		log.Fatal(err)
		return models.User{}
	}
	var dbUser models.User
	if query != nil {
		for query.Next() {
			var (
				id        int
				firstname string
				lastname  string
				email     string
				password  string
				createdat time.Time
			)
			err := query.Scan(&id, &firstname, &lastname, &email, &password, &createdat)
			if err != nil {
				log.Println(err)
			}
			dbUser = models.User{ID: id, Firstname: firstname, Lastname: lastname, Email: email, Password: password, CreatedAt: createdat}
		}
	}
	return dbUser
}
