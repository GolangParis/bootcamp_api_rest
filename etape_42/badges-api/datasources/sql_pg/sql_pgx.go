package sql

import (
	pgx "github.com/jackc/pgx/v4"

	"context"
	//"errors"
	"fmt"
	"log"
	"math/rand"
	
	"badges/types"
)

const (
	DB_DRIVER = "postgres"
)

type UserRecord struct {
	ID      types.ID `db:"id"`
	
	//Name    string   `db:"name" json:"name,omitempty"`
}

type BadgeRecord struct {
	ID      types.ID `db:"id"`
	
	Name    string   `db:"name" json:"name,omitempty"`
	URL     string   `db:"url"  json:"url,omitempty"`

	UserID types.ID  `db:"user_id"  json:"user_id,omitempty"`
}

var (
	conn *pgx.Conn
)

func NewUserID() types.ID {
	id := types.ID(rand.Int31n(100000))
	return id
}

func NewBadgeID() types.ID {
	id := types.ID(rand.Int31n(100000))
	return id
}

type DataSource struct {}

func (ds DataSource) Connect() {
	host := "127.0.0.1"
	user := "postgres"
	password := "postgres"
	dbname := "badges_db"

	connString := fmt.Sprintf("postgres://%s:%s@%s:5432/%s", user, password, host, dbname)

	var err error

	// Création d'une connexion au serveur Postgres

	conn, err = pgx.Connect(context.Background(), connString)
	if err != nil {
		log.Fatalf("Connection error: %v", err)
	}

	if err = conn.Ping(context.Background()); err != nil {
		log.Fatalf("Failed to ping db with error: %v", err)
	}
}

func (ds DataSource) CreateUser(c types.User) (types.ID, error) {
	// if c == nil {
	// 	log.Println("Invalid param attr")
	// 	return types.ID(0), errors.New("Invalid param attr")
	// }

	userRec := UserRecord{
		ID: NewUserID(),
		//Name: c.Name,
	}
	
	if _, err := conn.Exec(context.Background(), Q_INSERT_USER,	userRec.ID /*, userRec.Name*/); err != nil {
		log.Fatal("insert-user", err)
		return 0, err
	}

	return userRec.ID, nil
}


func (ds DataSource) GetUsers() ([]types.User, error) {
	// ([]types.Users, error)

	return nil, nil
}

func (ds DataSource) GetUserBadges(userID types.ID) ([]types.Badge, error) {
	rows, err := conn.Query(context.Background(), Q_SELECT_BADGES_BY_USER_ID, userID)
	if err != nil {
		log.Fatal("get-user-badges", err)
		return nil, err
	}

	defer rows.Close()

	bb := []types.Badge{}
	
	for rows.Next() {
		brec := BadgeRecord{}

		if err = rows.Scan(&brec.ID, &brec.Name, &brec.URL, nil); err != nil {
			if err == pgx.ErrNoRows {
				return nil, nil
			} else {
				return nil, fmt.Errorf("GetUserBadges rows scan err: %v", err)
			}
		}

		b := types.Badge{
			Name: brec.Name,
			URL: brec.URL,
		}

		bb = append(bb, b)
	}

	return bb, nil
}

func (ds DataSource) AddUserBadge(userID types.ID, c types.Badge) (types.ID, error) {
	//if c == nil {
	//	log.Println("Invalid param attr")
	//	return types.ID(0), errors.New("Invalid param attr")
	//}

	badgeRec := BadgeRecord{
		ID: NewBadgeID(),
		Name: c.Name,
		URL: c.URL,
		UserID: userID,
	}
	
	if _, err := conn.Exec(context.Background(), Q_INSERT_BADGE,
		badgeRec.ID, badgeRec.Name, badgeRec.URL, badgeRec.UserID); err != nil {
		log.Fatal("insert-badge", err)
		return 0, err
	}

	return badgeRec.ID, nil
}

func (ds DataSource) UpdateUserBadge(userID types.ID, b types.Badge) error {
	// TODO
	return nil
}

func (ds DataSource) DeleteUserBadge(userID types.ID, badgeName string) error {
	// TODO
	return nil
}
