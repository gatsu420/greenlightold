package main

import "net/http"

func (app *application) homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		// http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		app.notFoundResponse(w, r)
		return
	}

	data := envelope{
		"message": "welcome home",
	}

	err := app.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
