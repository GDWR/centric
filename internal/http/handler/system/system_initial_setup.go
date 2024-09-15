package system

import (
	"encoding/json"
	"math/rand"
	"net/http"

	"github.com/gdwr/centric/internal/database"
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
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	if initialSetupComplete {
		http.Error(writer, "Initial setup already completed", http.StatusBadRequest)
		return
	}

	var data initialSetupSchema
	if err := json.NewDecoder(request.Body).Decode(&data); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := h.userService.CreateUser(data.AdminAccount.Username, data.AdminAccount.Password)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = h.databaseService.InitialSetup(request.Context(), user.ID); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = h.databaseService.CreateEnvironment(request.Context(), database.CreateEnvironmentParams{
		ID:   uuid.New().String(),
		Name: data.Environment.Name,
		Uri:  data.Environment.Uri,
		Icon: possibleIcons[rand.Intn(len(possibleIcons))],
	})
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(data)
}
