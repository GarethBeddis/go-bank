package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	listenAddr string
	store      Storage
}

func NewApiServer(listenAddr string, store Storage) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		store:      store,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/account", makeHTTPHandleFunc(s.handleAccount))

	router.HandleFunc("/account/{id}", makeHTTPHandleFunc(s.handleGetAccount))

	log.Println("API server listening on", s.listenAddr)

	http.ListenAndServe(s.listenAddr, router)
}

func (s *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case http.MethodGet:
		return s.handleGetAccount(w, r)
	case http.MethodPost:
		return s.handleCreateAccount(w, r)
	case http.MethodDelete:
		return s.handleDeleteAccount(w, r)
	case http.MethodPut:
		return s.handleTransfer(w, r)
	default:
		return fmt.Errorf("method %s not allowed", r.Method)
	}
}

func (s *APIServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	id := mux.Vars(r)["id"]

	fmt.Println(id)

	return WriteJSON(w, http.StatusOK, &Account{})
}

func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	createAccountRequest := CreateAccountRequest{}

	if err := json.NewDecoder(r.Body).Decode(&createAccountRequest); err != nil {
		return err
	}

	account := CreateAccount(createAccountRequest.Username)

	if err := s.store.CreateAccount(account); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, account)
}

func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleTransfer(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// Helpers
func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

type apiFunc func(w http.ResponseWriter, r *http.Request) error

type ApiError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func makeHTTPHandleFunc(fn apiFunc) http.HandlerFunc {
	// Wrap the apiFunc in a http.HandlerFunc
	return func(w http.ResponseWriter, r *http.Request) {

		// If the function returns an error, write it to the response
		if err := fn(w, r); err != nil {
			var apiErr *ApiError

			// As() will check if the error is an ApiError, and if so, assign it to apiErr
			if ok := errors.As(err, &apiErr); ok {
				// If it is, write the ApiError to the response
				WriteJSON(w, apiErr.Code, apiErr)
			} else {
				// Otherwise, return a generic 500 Internal Server Error
				WriteJSON(w, http.StatusInternalServerError, &ApiError{
					Code:    http.StatusInternalServerError,
					Message: "Internal Server Error",
				})
			}
		}
	}
}
