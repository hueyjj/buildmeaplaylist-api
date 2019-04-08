package response

//ErrorBadFormResponse malformed post form
type ErrorBadFormResponse struct {
	Status string `json:"status,omitempty"`
}

//SignupResponse responses tells whether or not a user with email, first name,
//and last name was created in the database
type SignupResponse struct {
	Email     string `json:"email,omitempty"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	Status    string `json:"status,omitempty"`
}

//LoginResponse after logging in, returns a response to the client
type LoginResponse struct {
	Email  string `json:"email,omitempty"`
	Status string `json:"status,omitempty"`
}

//LoginErrorResponse after login fail, returns a response to the client
type LoginErrorResponse struct {
	Email  string `json:"email,omitempty"`
	Status string `json:"status,omitempty"`
}
