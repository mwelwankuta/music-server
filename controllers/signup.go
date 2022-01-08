package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mwelwankuta/music-server/controllers/utils"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20)
	form := r.MultipartForm.Value

	email := form["email"]

	if err := utils.SendEmail(email, fmt.Sprintf("hello %v", email)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"email": email[0],
	})
}
