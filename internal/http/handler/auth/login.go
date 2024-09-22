package auth

import (
	"encoding/json"
	"net/http"

	"github.com/gdwr/centric/internal/http/response"
)

type loginSchema struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type loginResponse struct {
	Token string `json:"token"`
}

func (h Handler) login(writer http.ResponseWriter, request *http.Request) {
	loginSchema := loginSchema{}
	if err := json.NewDecoder(request.Body).Decode(&loginSchema); err != nil {
		response.BadRequest("invalid request body", writer)
		return
	}

	if loginSchema.Username == "" || loginSchema.Password == "" {
		response.BadRequest("username and password are required", writer)
		return
	}

	user, err := h.databaseService.GetUserByName(request.Context(), loginSchema.Username)
	if err != nil {
		response.Unauthorized("invalid username or password", writer)
		return
	}

	success, err := h.authService.VerifyPassword(loginSchema.Password, user.Password)
	if err != nil || !success {
		response.Unauthorized("invalid username or password", writer)
		return
	}

	token, err := h.authService.GenerateToken(request.Context(), user.ID)
	if err != nil {
		response.InternalServerError(err, "unable to log you in at this time", writer)
		return
	}

	response.Json(loginResponse{Token: token}, writer)
}
