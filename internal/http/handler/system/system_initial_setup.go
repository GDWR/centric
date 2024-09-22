package system

import (
	"crypto/rand"
	"encoding/json"
	mathRand "math/rand"
	"net/http"

	"github.com/gdwr/centric/internal/database"
	"github.com/gdwr/centric/internal/http/response"
	"github.com/google/uuid"
)

type initialSetupSchema struct {
	AdminAccount struct {
		Username string `json:"username"`
		Password string `json:"password"`
	} `json:"adminAccount"`
	Environment struct {
		Uri  string `json:"uri"`
		Name string `json:"name"`
	} `json:"environment"`
}

var possibleIcons = []string{
	"lucide:mountain",
	"lucide:cloud",
	"lucide:moon",
	"lucide:sun",
	"lucide:star",
	"lucide:cloud-sun",
	"lucide:cloud-lightning",
	"lucide:alarm-clock",
	"lucide:anchor",
	"lucide:apple",
	"lucide:anvil",
	"lucide:candy",
}

func (h Handler) systemInitialSetup(writer http.ResponseWriter, request *http.Request) {
	initialSetupComplete, err := h.databaseService.InitialSetupComplete(request.Context())
	if err != nil {
		response.InternalServerError(err, "Unable to retrieve initial setup status", writer)
		return
	}
	if initialSetupComplete {
		response.BadRequest("Initial setup already completed", writer)
		return
	}

	var data initialSetupSchema
	if err := json.NewDecoder(request.Body).Decode(&data); err != nil {
		response.BadRequest("Invalid request body", writer)
		return
	}

	user, err := h.userService.CreateUser(request.Context(), data.AdminAccount.Username, data.AdminAccount.Password)
	if err != nil {
		response.InternalServerError(err, "Unable to create user", writer)
		return
	}

	secret := make([]byte, 256)
	if _, err := rand.Read(secret); err != nil {
		response.InternalServerError(err, "Unable to generate secret", writer)
		return
	}

	if err = h.databaseService.InitialSetup(request.Context(), database.InitialSetupParams{
		OwnerID:   user.ID,
		JwtSecret: secret,
	}); err != nil {
		response.InternalServerError(err, "Unable to complete initial setup", writer)
		return
	}

	_, err = h.databaseService.CreateEnvironment(request.Context(), database.CreateEnvironmentParams{
		ID:   uuid.New().String(),
		Name: data.Environment.Name,
		Uri:  data.Environment.Uri,
		Icon: possibleIcons[mathRand.Intn(len(possibleIcons))],
	})
	if err != nil {
		response.InternalServerError(err, "Unable to create environment", writer)
		return
	}

	response.Json(data, writer)
}
