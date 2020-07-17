package controller

import("net/http"
	"database/sql"
	"API/model"
	"API/repository"
	"encoding/json"
	"github.com/gorilla/mux"
	"strconv"
	"log"
	"io"
	)

func CreateHandler(db *sql.DB) http.HandlerFunc {
	return func (w http.ResponseWriter,r *http.Request){
		w.Header().Set("Content-Type", "application/json")
		var u model.User
		err:=json.NewDecoder(r.Body).Decode(&u)

		if(err==io.EOF){
			respondWithError(w,"Empty Body")
			return
		}

		id,msg := repository.CreateUser(db,&u)
		if(id==0) {
			respondWithError(w,msg)
		} else {
			respondwithJSON(w,u)
		}
	}
}


func ListHandler(db *sql.DB) http.HandlerFunc {
	return func (w http.ResponseWriter,r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		
		params:=mux.Vars(r)

		limit,err:=strconv.ParseUint(params["limit"],10,32)
		if(err!=nil){
			log.Fatal(err)
		}

		offset,err:=strconv.ParseUint(params["offset"],10,32)
		if(err!=nil) {
			log.Fatal(err)
		}

		users:=repository.ListUsers(db,limit,offset)

		respondwithJSON(w,users)
	}
}

func UpdateHandler(db *sql.DB) http.HandlerFunc {
	return func (w http.ResponseWriter,r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var u model.User
		var id uint64
		params:=mux.Vars(r)
		id,err:=strconv.ParseUint(params["id"],10,32)

		if(err!=nil) {
			log.Fatal(err)
		}

		err2:=json.NewDecoder(r.Body).Decode(&u)

		if(err2==io.EOF){
			respondWithError(w,"Empty Body")
			return
		}

		msg,user:=repository.UpdateUser(db,id,u)

		if(user==nil) {
			respondWithError(w,msg)
		} else {
			respondwithJSON(w,*user)
		}
	}
}

func DeleteHandler(db *sql.DB) http.HandlerFunc {
	return func (w http.ResponseWriter,r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var id uint64
		params:=mux.Vars(r)
		id,err:=strconv.ParseUint(params["id"],10,32)

		if(err!=nil) {
			log.Fatal(err)
		}

		msg:=repository.DeleteUser(db,id)
		respondWithError(w,msg)
	}
}

func respondwithJSON(w http.ResponseWriter,u interface{}) {
	response, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func respondWithError(w http.ResponseWriter, msg string) {
	respondwithJSON(w, map[string]string{"message": msg})
}