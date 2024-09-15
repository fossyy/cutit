package alias

import (
	"github.com/fossyy/cutit/app"
	"github.com/gorilla/mux"
	"net/http"
)

func ALL(w http.ResponseWriter, r *http.Request) {
	alias := mux.Vars(r)
	link, err := app.Server.Database.GetLink(alias["alias"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, link.URL, http.StatusSeeOther)
}
