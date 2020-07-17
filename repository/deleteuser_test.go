package repository


import (
	"testing"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestDeleteUser(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()	

	columns := []string{"user_id"}

	mock.ExpectQuery("Select count(.+) from users*").
		WithArgs(uint64(1)).
		WillReturnRows(sqlmock.NewRows(columns).FromCSVString("1"))

	mock.ExpectExec("delete from users*").
		WithArgs(uint64(1)).WillReturnResult(sqlmock.NewResult(1,1))

	msg:=DeleteUser(db,1)

	assert.Equal(t,"Sucessfull",msg)
}

func TestDeleteIDNotExist(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	columns := []string{"count"}

	mock.ExpectQuery("Select count(.+) from users*").
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows(columns).FromCSVString("0"))

	msg:=DeleteUser(db,1)

	assert.Equal(t,"ID doesn't exist",msg)
}