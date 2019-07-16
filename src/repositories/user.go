package repositories

import (
	"database/sql"
	"fmt"
	"github/kdswto/webserver_example/src/models"
	"log"
)

type UserRepository struct {
	Database *sql.DB
}

func (r *UserRepository) GetAll() []*models.User {
	if r.Database == nil {
		log.Fatal("UserRepository.GetAll: Database not initialize")
	}

	results, err := r.Database.Query(`
		SELECT u.id, u.name, w.id, w.money 
		FROM user u
		LEFT JOIN wallet w on u.id = w.user_id
	`)
	if err != nil {
		log.Fatal("UserRepository.GetAll, Query: ", err)
	}

	var users []*models.User
	for results.Next() {
		user := &models.User{}
		wallet := &models.Wallet{}
		err = results.Scan(&user.Id, &user.Name, &wallet.Id, &wallet.Money)
		if err != nil {
			log.Fatal("UserRepository.GetAll, Scan: ", err)
		}
		user.Wallet = wallet
		users = append(users, user)
	}

	return users
}

func (r *UserRepository) GetById(id int) *models.User {
	if r.Database == nil {
		log.Fatal("UserRepository.GetById: Database not initialize")
	}

	statement := fmt.Sprintf(`
		SELECT u.id, u.name, w.id, w.money 
		FROM user u 
		LEFT JOIN wallet w ON u.id = w.user_id
		WHERE u.id = %d
	`, id)
	row := r.Database.QueryRow(statement)
	user := &models.User{}
	wallet := &models.Wallet{}

	err := row.Scan(&user.Id, &user.Name, &wallet.Id, &wallet.Money)
	if err != nil {
		log.Println("UserRepository.GetById: ", err)
	}

	return user
}
