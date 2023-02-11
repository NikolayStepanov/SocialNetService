package repository

import (
	"SocialNetHTTPService/internal/domain"
	"context"
	"database/sql"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"log"
)

const UsersTable = "users"

type UserDatabase struct {
	db *sqlx.DB
}

func NewUserDatabase(db *sqlx.DB) *UserDatabase {
	return &UserDatabase{db: db}
}

func (ud *UserDatabase) GetUser(ctx context.Context, userID int) (domain.User, error) {
	var (
		id           int
		name         string
		age          int
		friendsArray pq.Int64Array
		friends      []int
		user         domain.User
		row          *sql.Row
	)
	query, args, err := sq.Select("*").
		From(UsersTable).
		Where("id = ?", userID).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if row = ud.db.QueryRow(query, args...); err != nil {
		log.Println(err)
		return user, err
	}
	if err = row.Scan(&id, &name, &age, &friendsArray); err != nil {
		log.Println(err)
		return user, err
	}
	user.SetId(id)
	user.SetName(name)
	user.SetAge(uint8(age))
	for _, value := range friendsArray {
		friends = append(friends, int(value))
	}
	user.SetFriends(friends)
	return user, err
}

func (ud *UserDatabase) CreateUser(ctx context.Context, user domain.User) (int, error) {
	var (
		sql  string
		args []interface{}
		err  error
		id   int
	)
	sql, args, err = sq.Insert(UsersTable).Columns("name", "age", "friends").
		Values(user.Name(), int(user.Age()), pq.Array(user.Friends())).
		Suffix("RETURNING id").
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err = ud.db.GetContext(ctx, &id, sql, args...); err != nil {
		return 0, err
	}

	return id, err
}

func (ud *UserDatabase) CreateUsers(ctx context.Context, users []domain.User) error {
	var (
		err error
		id  int
	)
	for _, user := range users {
		if id, err = ud.CreateUser(ctx, user); err != nil {
			log.Println(err)
			return err
		} else {
			log.Printf("User name = %s created id=%d", user.Name(), id)
		}
	}
	return err
}

func (ud *UserDatabase) DeleteUser(ctx context.Context, user domain.User) error {
	err := ud.DeleteUserById(ctx, user.ID())
	return err
}

func (ud *UserDatabase) DeleteUserById(ctx context.Context, userID int) error {
	var (
		sql  string
		args []interface{}
		err  error
	)
	sql = `DELETE FROM users WHERE id = $1`
	args = append(args, userID)
	if _, err = ud.db.ExecContext(ctx, sql, args...); err != nil {
		log.Println(err)
	}
	return err
}

func (ud *UserDatabase) DeleteUsers(ctx context.Context, users []domain.User) error {
	for index, _ := range users {
		ud.DeleteUser(ctx, users[index])
	}
	return nil
}

func (ud *UserDatabase) ContainsUser(ctx context.Context, user domain.User) bool {
	return ud.ContainsUserById(ctx, user.ID())
}

func (ud *UserDatabase) ContainsUserById(ctx context.Context, userID int) bool {
	var (
		sql  string
		args []interface{}
		err  error
		res  bool
	)
	sql = `SELECT EXISTS (SELECT id FROM users WHERE id = $1)`
	args = append(args, userID)
	if err = ud.db.GetContext(ctx, &res, sql, args...); err != nil {
		log.Println(err)
	}
	return res
}

func (ud *UserDatabase) UpdateAgeUser(ctx context.Context, userID int, age int) error {
	var (
		sql  string
		args []interface{}
		err  error
	)
	sql = `UPDATE users SET age = $1 WHERE id = $2`
	args = append(args, age, userID)
	if _, err = ud.db.ExecContext(ctx, sql, args...); err != nil {
		log.Println(err)
	}
	return err
}

func (ud *UserDatabase) AddFriendUser(ctx context.Context, userID int, friendID int) error {
	var (
		res  sql.Result
		sql  string
		args []interface{}
		err  error
		rows int64
	)
	sql = `SELECT friends FROM users WHERE id = $1 AND $2 = any(friends)`
	args = append(args, userID, friendID)
	if res, err = ud.db.ExecContext(ctx, sql, args...); err != nil {
		log.Println(err)
	}
	if rows, err = res.RowsAffected(); err != nil {
		log.Println(err)
	}
	if rows == 1 {
		err = fmt.Errorf("already friends")
	} else {
		sql = `UPDATE users SET friends = array_append(friends, $1) WHERE id = $2`
		args = nil
		args = append(args, friendID, userID)
		if _, err = ud.db.ExecContext(ctx, sql, args...); err != nil {
			log.Println(err)
		}
	}
	return err
}

func (ud *UserDatabase) DeleteFriend(ctx context.Context, sourceUserID int, targetFriendID int) (int, error) {
	var (
		sql  string
		args []interface{}
		err  error
	)
	sql = `UPDATE users SET friends = array_remove(friends, $1) WHERE id = $2`
	args = append(args, targetFriendID, sourceUserID)
	if _, err = ud.db.ExecContext(ctx, sql, args...); err != nil {
		log.Println(err)
	}
	return sourceUserID, nil
}
