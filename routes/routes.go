package routes

import (
	"github.com/gorilla/mux"
	"github.com/mwelwankuta/music-server/controllers"
)

func Routes(r *mux.Router) {
	r.HandleFunc("/api/signup", controllers.SignUp)
}
