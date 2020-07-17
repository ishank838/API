package repository

import (
	"testing"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"API/model"
)

func TestCreateuser(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()	

	user:= model.User{
			User_name : "Raja",
			Name : "Irma Frank",
			Email : "Vestibulum.accumsan@est.edu"}

	columns := []string{"user_id"}

	mock.ExpectQuery("select count(.+) from users*").
		WithArgs(user.User_name,user.Email).
		WillReturnRows(sqlmock.NewRows(columns).FromCSVString("0"))

	mock.ExpectQuery("insert into users values*").
		WithArgs(user.User_name,user.Name,user.Email).
		WillReturnRows(sqlmock.NewRows(columns).FromCSVString("1"))

	id,msg:=CreateUser(db,&user)

	assert.Equal(t,uint64(1),id)
	assert.Equal(t,"Sucessfull",msg)
}

func TestUserNoUsername(t *testing.T) {
	db,_,_:=sqlmock.New()

	user:= model.User{
			User_name : "",
			Name : "Irma Frank",
			Email : "Vestibulum.accumsan@est.edu"}

	id,msg:=CreateUser(db,&user)

	assert.Equal(t,uint64(0),id)
	assert.Equal(t,"Username can't Be nil",msg)
}

func TestUserNoEmail(t *testing.T) {
	db,_,_:=sqlmock.New()

	user:= model.User{
			User_name : "Raja",
			Name : "Irma Frank",
			Email : ""}

	id,msg:=CreateUser(db,&user)

	assert.Equal(t,uint64(0),id)
	assert.Equal(t,"Email can't be nil",msg)
}

func TestCreateUserExists(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()	

	user:= model.User{
			User_name : "Raja",
			Name : "Irma Frank",
			Email : "Vestibulum.accumsan@est.edu"}

	columns := []string{"user_id"}

	mock.ExpectQuery("select count(.+) from users*").
		WithArgs(user.User_name,user.Email).
		WillReturnRows(sqlmock.NewRows(columns).FromCSVString("1"))

	id,msg:=CreateUser(db,&user)

	assert.Equal(t,uint64(0),id)
	assert.Equal(t,"User already exists",msg)

}