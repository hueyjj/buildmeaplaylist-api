package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/hueyjj/buildmeaplaylist-api/response"
	"golang.org/x/crypto/bcrypt"
)

//SignUpHandler creates account for user, hashes password with bcrypt
func SignUpHandler(ct *Controller, store *sessions.CookieStore) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		//parses the variable from the route, maybe not needed?
		//vars := mux.Vars(r)

		// TODO: move error handling into own package/file
		if err := r.ParseForm(); err != nil {
			log.Printf("Could not parse form: %v\n", err)

			var resp *response.ErrorBadFormResponse
			resp = &response.ErrorBadFormResponse{
				Status: "Unable parse form",
			}
			r, err := json.Marshal(resp)
			if err != nil {
				log.Printf("Unable to marshal response: %v\n", err)
			}
			w.WriteHeader(http.StatusBadRequest)
			w.Write(r)
			return
		}

		email := r.FormValue("email")
		password := r.FormValue("password")
		firstName := r.FormValue("firstName")
		lastName := r.FormValue("lastName")

		//Hash password with bcrypt
		{
			hashedPass, err := HashPassword(password)
			if err != nil {
				//What do we do when we can't hash the password? what even?!
				//This shouldn't break...
			}
			password = hashedPass
		}

		sqlStatement := `
		INSERT INTO members(email, password, first_name, last_name)
		VALUES('%s', '%s', '%s', '%s');`
		sqlStatement = fmt.Sprintf(sqlStatement, email, password, firstName, lastName)

		var status string
		_, err := ct.db.Exec(sqlStatement)
		if err != nil {
			log.Printf(sqlStatement)
			log.Printf("Unable to execute sql statement for adding user: %v\n", err)
			status = fmt.Sprintf("Unable to execute sql statement for adding user: %v\n", err)
		} else {
			log.Printf("Successfully added %s %s (%s) to the database\n", firstName, lastName, email)
			status = fmt.Sprintf("Successfully added %s %s (%s) to the database\n", firstName, lastName, email)
		}

		var signupResponse *response.SignupResponse
		signupResponse = &response.SignupResponse{
			Email:     email,
			FirstName: firstName,
			LastName:  lastName,
			Status:    status,
		}
		resp, _ := json.Marshal(signupResponse)

		w.WriteHeader(http.StatusOK)
		w.Write(resp)
	}
}

//LoginHandler validate against database, and create and save session in store
func LoginHandler(ct *Controller, store *sessions.CookieStore) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if err := r.ParseForm(); err != nil {
			log.Printf("Could not parse form: %v\n", err)

			var resp *response.ErrorBadFormResponse
			resp = &response.ErrorBadFormResponse{
				Status: "Unable parse form",
			}
			r, err := json.Marshal(resp)
			if err != nil {
				log.Printf("Unable to marshal response: %v\n", err)
			}
			w.WriteHeader(http.StatusBadRequest)
			w.Write(r)
			return
		}

		email := r.FormValue("email")
		password := r.FormValue("password")

		sqlStatement := `
		SELECT password FROM members WHERE email='%s';
		`
		sqlStatement = fmt.Sprintf(sqlStatement, email)

		var hashedPassword string
		err := ct.db.QueryRow(sqlStatement).Scan(&hashedPassword)
		if err != nil {
			log.Printf("Unable to find hashed password: %v\n", err)

			var loginErrorResponse *response.LoginErrorResponse
			loginErrorResponse = &response.LoginErrorResponse{
				Email:  email,
				Status: "Unable to find password",
			}
			resp, _ := json.Marshal(loginErrorResponse)

			w.WriteHeader(http.StatusOK)
			w.Write(resp)
			return
		}

		var loginResponse *response.LoginResponse
		if CheckPassword(password, hashedPassword) == true {
			loginResponse = &response.LoginResponse{
				Email:  email,
				Status: "Successfully logged in",
			}
		} else {
			loginResponse = &response.LoginResponse{
				Email:  email,
				Status: "Failed to login. Password does not match",
			}
		}
		resp, _ := json.Marshal(loginResponse)
		w.WriteHeader(http.StatusOK)
		w.Write(resp)
	}
}

//LogOutHandler removes session and redirects to homepage
func LogOutHandler(ct *Controller, store *sessions.CookieStore) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

//HashPassword hashes password https://gowebexamples.com/password-hashing/d
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

//CheckPassword checks if hash and password are the same https://gowebexamples.com/password-hashing/d
func CheckPassword(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
