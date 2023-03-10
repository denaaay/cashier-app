package api

import (
	"a21hc3NpZ25tZW50/model"
	"context"
	"encoding/json"
	"net/http"
)

func (api *API) Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("session_token")
		if err != nil {
			if err == http.ErrNoCookie {
				w.WriteHeader(401)
				error := model.ErrorResponse{
					Error: "http: named cookie not present",
				}

				jsonError, err := json.Marshal(error)
				if err != nil {
					return
				}
				w.Write(jsonError)
				return
			}
		}

		sessionToken := c.Value // TODO: replace this

		sessionFound, err := api.sessionsRepo.CheckExpireToken(sessionToken)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
			return
		}

		ctx := context.WithValue(r.Context(), "username", sessionFound.Username)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (api *API) Get(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.WriteHeader(405)
			error := model.ErrorResponse{
				Error: "Method is not allowed!",
			}
			jsonErorr, err := json.Marshal(error)
			if err != nil {
				panic(err)
			}
			w.Write([]byte(jsonErorr))
			return
		}
		// TODO: answer here
		next.ServeHTTP(w, r)
	})
}

func (api *API) Post(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			w.WriteHeader(405)
			error := model.ErrorResponse{
				Error: "Method is not allowed!",
			}
			jsonErorr, err := json.Marshal(error)
			if err != nil {
				panic(err)
			}
			w.Write([]byte(jsonErorr))
			return
		}
		// TODO: answer here
		next.ServeHTTP(w, r)
	})
}
