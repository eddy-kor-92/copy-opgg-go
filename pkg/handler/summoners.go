package handler

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type SummonerHandler struct{}

func (h SummonerHandler) Init(rootRouter *mux.Router) {
	rootRouter.HandleFunc("/summoners/profile-info/{searchWord}", h.getSummoner).Methods("GET")
}

// get mapping /summoners/profile-info/{searchWord}
func (h SummonerHandler) getSummoner(w http.ResponseWriter, r *http.Request) {
	searchWord := mux.Vars(r)["searchWord"]
	fmt.Fprintf(w, "get summoner with name %s called!\n", searchWord)
}
