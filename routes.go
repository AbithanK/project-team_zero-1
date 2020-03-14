package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

const clientID = "ecd70f356de063418ef0"
const clientSecret = "0090c648260f30e8da7a38d252b8b6a17f1576c1"

func getToken(r *http.Request) string {
	for _, cookie := range r.Cookies() {
		if cookie.Name == "token" {
			return cookie.Value
		}
	}
	return ""
}

func handleRoutes(router *mux.Router) {

	// EXAMPLE OF API SUPER BASIC API HANDLER
	// CAN ALSO USE MIDDLEWARE TO DICTATE WHAT
	// TYPE OF CALLS THE ENDPOINT ACCEPTS
	// WITH .Methods("GET", "POST", etc...)
	// SEE https://github.com/gorilla/mux#examples
	router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		// an example API handler
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})

	// ADD OTHER ROUTES HERE
	router.HandleFunc("/api/boards/", createBoard).Methods("POST")
	router.HandleFunc("/api/boards/{boardID}/columns/", createColumn).Methods("POST")
	router.HandleFunc("/api/boards/{boardID}/columns/{columnID}/tickets/", createTicket).Methods("POST")
	router.HandleFunc("/api/boards/{boardID}/", updateBoard).Methods("PATCH")
	router.HandleFunc("/api/boards/{boardID}/", getBoard).Methods("GET")
	router.HandleFunc("/api/boards/{boardID}/", deleteBoard).Methods("DELETE")
	router.HandleFunc("/api/boards/{boardID}/columns/{columnID}/", deleteColumn).Methods("DELETE")
	router.HandleFunc("/api/boards/{boardID}/columns/{columnID}/tickets/{ticketID}", deleteTicket).Methods("DELETE")
	// Handle GitHub OAuth authorization
	router.HandleFunc("/oauth/redirect", func(w http.ResponseWriter, r *http.Request) {
		githubAuthorize(w, r, clientID, clientSecret)
	})

	// POST to create a branch
	router.HandleFunc("/api/branches", func(w http.ResponseWriter, r *http.Request) {
		createBranch(w, r, getToken(r))
	})

	// POST to retrieve branches for given repo
	router.HandleFunc("/api/repos/branches", func(w http.ResponseWriter, r *http.Request) {
		getBranches(w, r, getToken(r))
	})

	//GET
	router.HandleFunc("/api/repos", func(w http.ResponseWriter, r *http.Request) {
		getRepositories(w, r, getToken(r))
	})

	//POST to get a list of collaborators
	router.HandleFunc("/api/repos/users", func(w http.ResponseWriter, r *http.Request) {
		getRepositoryCollaborators(w, r, getToken(r))
	})
}
