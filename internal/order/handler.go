package order

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/thongsoi/testg/database"
)

func OrderHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/order.html"))
	tmpl.Execute(w, r)

}

func FetchMarketsHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query("SELECT id, en_name FROM markets")
	if err != nil {
		http.Error(w, "Failed to fetch markets", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var marketsHTML string
	for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			http.Error(w, "Error scanning markets", http.StatusInternalServerError)
			return
		}
		// Build the HTML options dynamically
		marketsHTML += fmt.Sprintf(`<option value="%d">%s</option>`, id, name)
	}

	// Return the HTML options
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(marketsHTML))
}
