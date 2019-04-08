package response

//ErrorBadFormResponse malformed post form
type ErrorBadFormResponse struct {
	Status string `json:"status,omitempty"`
}

//AddUserResponse responses tells whether or not a user with email, first name,
//and last name was created in the database
type AddUserResponse struct {
	Email     string `json:"email,omitempty"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	Status    string `json:"status,omitempty"`
}
