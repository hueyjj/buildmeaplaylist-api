package handlers

//AddUserHandler adds user to database
//func (ps *Postgres) AddUserHandler(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//
//	//parses the variable from the route, maybe not needed?
//	//vars := mux.Vars(r)
//
//	// TODO: move error handling into own package/file
//	if err := r.ParseForm(); err != nil {
//		log.Printf("Could not parse form: %v\n", err)
//
//		var resp *response.ErrorBadFormResponse
//		resp = &response.ErrorBadFormResponse{
//			Status: "Unable parse form",
//		}
//		r, err := json.Marshal(resp)
//		if err != nil {
//			log.Printf("Unable to marshal response: %v\n", err)
//		}
//		w.WriteHeader(http.StatusBadRequest)
//		w.Write(r)
//		return
//	}
//
//	email := r.FormValue("email")
//	firstName := r.FormValue("firstName")
//	lastName := r.FormValue("lastName")
//
//	sqlStatement := `
//	INSERT INTO members(email, first_name, last_name)
//	VALUES('%s', '%s', '%s');`
//	sqlStatement = fmt.Sprintf(sqlStatement, email, firstName, lastName)
//
//	var status string
//	_, err := ps.db.Exec(sqlStatement)
//	if err != nil {
//		log.Printf(sqlStatement)
//		log.Printf("Unable to execute sql statement for adding user: %v\n", err)
//		status = fmt.Sprintf("Unable to execute sql statement for adding user: %v\n", err)
//	} else {
//		log.Printf("Successfully added %s %s (%s) to the database\n", firstName, lastName, email)
//		status = fmt.Sprintf("Successfully added %s %s (%s) to the database\n", firstName, lastName, email)
//	}
//
//	var addUserResponse *response.AddUserResponse
//	addUserResponse = &response.AddUserResponse{
//		Email:     email,
//		FirstName: firstName,
//		LastName:  lastName,
//		Status:    status,
//	}
//	resp, _ := json.Marshal(addUserResponse)
//
//	w.WriteHeader(http.StatusOK)
//	w.Write(resp)
//}
//
////RemoveUserHandler remove user from database
//func (ps *Postgres) RemoveUserHandler(w http.ResponseWriter, r *http.Request) {
//}
//
////UpdateUserHandler udpates user in database
//func (ps *Postgres) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
//}
