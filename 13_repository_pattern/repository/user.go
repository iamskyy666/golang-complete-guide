package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/iamskyy666/db_programming/13_repository_pattern/models"
	"golang.org/x/crypto/bcrypt"
)

type UserRepo interface {
	CreateUser(name, email, hashed_password, avatar string) (int64, error)
	GetUserByEmail(demail string) (*models.User, error)
	GetAllUsers() ([]models.User, error)
}

type SQLUserRepository struct {
	db *sql.DB
}

// CreateUser implements [UserRepo].
func (s *SQLUserRepository) CreateUser(name string, email string, hashed_password string, avatar string) (int64, error) {
	ctx:=context.Background()

	// begin the tx
	tx,err:=s.db.BeginTx(ctx,nil)
	if err!=nil{
		return 0,err
	}

	defer tx.Rollback()


	//query/stmt
	stmt,err:=tx.PrepareContext(ctx,`INSERT INTO users (name,email,hashed_password)  VALUES (?, ?, ?)`)
	if err!=nil{
		return 0,err
	}

	defer stmt.Close()

	//! hash the password first
	hp,err:=bcrypt.GenerateFromPassword([]byte(hashed_password),bcrypt.DefaultCost)
	if err!=nil{
		return 0,err
	}

	res,err:= stmt.Exec(name,email,string(hp))
	if err!=nil{
		return 0,err
	}

	userId,err:= res.LastInsertId()
	if err!=nil{
		return 0,err
	}

	// Now, create the profile

	profileStmt,err:=tx.PrepareContext(ctx, `INSERT INTO profile (user_id,avatar) VALUES(?, ?)`)

	if err!=nil{
		return 0,err
	}

	defer profileStmt.Close()
	_,err=profileStmt.Exec(userId,avatar)
	if err!=nil{
		return 0,err
	}

	// At the end, commit the tx.
	err = tx.Commit()
	if err!=nil{
		return 0,err
	}

	return userId,nil
}

// GetAllUsers implements [UserRepo].
func (s *SQLUserRepository) GetAllUsers() ([]models.User, error) {
	stmt:=`SELECT id,name,email,hashed_password,created_at FROM users`

	rows,err:=s.db.Query(stmt)

	if err!=nil{
		return nil,err
	}

	defer rows.Close()

	var users []models.User
		for rows.Next(){
			var user models.User
			if err:=rows.Scan(&user.ID,&user.Name,&user.Email,&user.HashedPassword, &user.CreatedAt); err!=nil{
			return nil,err
			}
			users = append(users, user)
		
		}
		if err:=rows.Err();err!=nil{
			return nil,err
		}
		return users,nil

}

// GetUserByEmail implements [UserRepo].
func (s *SQLUserRepository) GetUserByEmail(email string) (*models.User, error) {
	stmt:=`SELECT u.id,u.name,u.email,u.hashed_password,u.created_at, p.avatar FROM users u INNER JOIN profile p ON u.id = p.user_id WHERE u.email = ?`

	row:=s.db.QueryRow(stmt,email)

	// scan the data back into the struct
	var user models.User

	err:=row.Scan(&user.ID,&user.Name,&user.Email,&user.HashedPassword, &user.CreatedAt, &user.Profile.Avatar)

	if err!=nil{
		if err==sql.ErrNoRows{
			log.Fatal("No records found:",err.Error())
		}
		return nil,err
	}

	user.Profile.UserId=user.ID
	return &user, nil
}

// Create new UserRepository type
func NewSQLUserRepository(db *sql.DB) UserRepo {
	return &SQLUserRepository{
		db: db,
	}
}
