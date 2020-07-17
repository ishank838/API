package repository


import (
	"testing"
	"API/model"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func Testupdateuser(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()	

	user:= model.User{
			User_id: uint64(1),
			User_name : "Raja",
			Name : "Irma Frank",
			Email : "Vestibulum.accumsan@est.edu"}

	columns := []string{"user_id"}

	mock.ExpectQuery("select count(.+) from users*").
		WithArgs(user.User_name,user.Email).
		WillReturnRows(sqlmock.NewRows(columns).FromCSVString("1"))

	mock.ExpectQuery("update users set").
		WithArgs(user.User_name,user.Name,user.Email,user.User_id)

	msg,u:=UpdateUser(db,user.User_id,user)

	assert.Equal(t,"Sucessfull",msg)
	assert.Equal(t,user,*u)
}

func TestUserNotExist(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()	

	user:= model.User{
			User_id: uint64(1),
			User_name : "Raja",
			Name : "Irma Frank",
			Email : "Vestibulum.accumsan@est.edu"}

	columns := []string{"count"}

	mock.ExpectQuery("select count(.+) from users*").
		WithArgs(user.User_id).
		WillReturnRows(sqlmock.NewRows(columns).FromCSVString("0"))

	msg,u:=UpdateUser(db,user.User_id,user)

	assert.Equal(t,"ID doesn't exist",msg)
	
	if u!=nil {
		t.Errorf("nil value expected")
	}

}