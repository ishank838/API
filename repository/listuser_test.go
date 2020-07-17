package repository


import (
	"testing"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"API/model"
)

func TestListUser(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	var users []model.User = []model.User{
		{uint64(1), "Sara", "Idona Dunlap", "non@iaculisquispede.net"},
		{uint64(2),"Chadwick","Halla Moses","egestas@arcuac.com"},
		{uint64(3),"Tad","Chastity Goodman","elit.Nulla@Vivamuseuismodurna.com"},
		{uint64(4),"Aristotle","Yvette Patrick","feugiat.nec@massarutrummagna.org"},
		{uint64(5),"Summer","Sara Atkinson","dignissim@maurisa.com"}}

	columns := sqlmock.NewRows([]string{"user_id","username","name","email"}).
			AddRow(uint64(1),"Sara","Idona Dunlap","non@iaculisquispede.net").
			AddRow(uint64(2),"Chadwick","Halla Moses","egestas@arcuac.com").
			AddRow(uint64(3),"Tad","Chastity Goodman","elit.Nulla@Vivamuseuismodurna.com").
			AddRow(uint64(4),"Aristotle","Yvette Patrick","feugiat.nec@massarutrummagna.org").
			AddRow(uint64(5),"Summer","Sara Atkinson","dignissim@maurisa.com")

	mock.ExpectQuery("Select . from users ORDER BY*").
		WithArgs(uint64(5),uint64(0)).
		WillReturnRows(columns)

	u:=ListUsers(db,uint64(5),uint64(0))

	for i:=0; i < len(u); i++ {
		assert.Equal(t,users[i],u[i])
	}
}