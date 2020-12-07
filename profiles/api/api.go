package api

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func RegisterRoutes(router *mux.Router) error {
	router.HandleFunc("/api/profile/{uuid}", getProfile).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/api/profile/{uuid}", updateProfile).Methods(http.MethodPut, http.MethodOptions)

	return nil
}

func getUUID (w http.ResponseWriter, r *http.Request) (uuid string) {
	cookie, err := r.Cookie("access_token")
	if err != nil {
		http.Error(w, errors.New("error obtaining cookie: " + err.Error()).Error(), http.StatusBadRequest)
		log.Print(err.Error())
	}
	//validate the cookie
	claims, err := ValidateToken(cookie.Value)
	if err != nil {
		http.Error(w, errors.New("error validating token: " + err.Error()).Error(), http.StatusUnauthorized)
		log.Print(err.Error())
	}
	log.Println(claims)

	return claims["UserID"].(string)
}

func getProfile(w http.ResponseWriter, r *http.Request) {
	// Obtain the uuid from the url path and store it in a `uuid` variable
	// Hint: mux.Vars()
	// YOUR CODE HERE
	vars := mux.Vars(r)
	userID := vars["uuid"]

	// Initialize a new Profile variable
	//YOUR CODE HERE
	profile := Profile{}

	// Obtain all the information associated with the requested uuid
	// Scan the information into the profile struct's variables
	// Remember to pass in the address!
	err := DB.QueryRow("SELECT * FROM users WHERE uuid = ?;", userID).Scan(&profile.Firstname, &profile.Lastname, &profile.Email, &profile.UUID)

	/*  Check for errors with querying the database
		Return an Internal Server Error if such an error occurs
	*/
	if err != nil {
		http.Error(w, errors.New("issue retrieving user information").Error(), http.StatusInternalServerError)
		log.Print(err.Error())
		return
	}

  	//encode fetched data as json and serve to client
	json.NewEncoder(w).Encode(profile)
	return
}

func updateProfile(w http.ResponseWriter, r *http.Request) {
	// Obtain the requested uuid from the url path and store it in a `uuid` variable
	// YOUR CODE HERE
	vars := mux.Vars(r)
	userID := vars["uuid"]

	// Obtain the userID from the cookie
	// YOUR CODE HERE
	userIDCookie := getUUID(w, r)

	// If the two ID's don't match, return a StatusUnauthorized
	// YOUR CODE HERE
	if userID != userIDCookie {
		http.Error(w, errors.New("cannot delete a post that is not yours").Error(), http.StatusUnauthorized)
		return
	}

	// Decode the Request Body's JSON data into a profile variable
	profile := Profile{}
	err := json.NewDecoder(r.Body).Decode(&profile)

	// Return an InternalServerError if there is an error decoding the request body
	// YOUR CODE HERE
	if err != nil {
		http.Error(w, errors.New("issue updating profile").Error(), http.StatusInternalServerError)
		log.Print(err.Error())
		return
	}

	// Insert the profile data into the users table
	// Check db-server/initdb.sql for the scheme
	// Make sure to use REPLACE INTO (as covered in the SQL homework)
	_, err = DB.Exec("REPLACE INTO users VALUES (?, ?, ?, ?);", profile.Firstname, profile.Lastname, profile.Email, profile.UUID)

	// Return an internal server error if any errors occur when querying the database.
	// YOUR CODE HERE
	if err != nil {
		http.Error(w, errors.New("issue updating profile").Error(), http.StatusInternalServerError)
		log.Print(err.Error())
		return
	}
	return
}
