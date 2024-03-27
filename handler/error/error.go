package errorHandler

import (
	errorView "github.com/fossyy/cutit/view/error"
	"net/http"
)

func ALL(w http.ResponseWriter, r *http.Request) {
	component := errorView.Main("Not Found")
	component.Render(r.Context(), w)
}
