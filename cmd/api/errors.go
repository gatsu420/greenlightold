package main

import "net/http"

func (app *application) logError(r *http.Request, err error) {
	app.logger.Println(err)
}

func (app *application) errorResponse(
	w http.ResponseWriter, r *http.Request, status int, message interface{},
) {
	env := envelope{"error": message}

	err := app.writeJSON(w, status, env, nil)
	if err != nil {
		app.logError(r, err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (app *application) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logError(r, err)
	app.errorResponse(
		w,
		r,
		http.StatusInternalServerError,
		http.StatusText(http.StatusInternalServerError),
	)
}

func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(
		w,
		r,
		http.StatusInternalServerError,
		http.StatusText(http.StatusInternalServerError),
	)
}

func (app *application) methodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(
		w,
		r,
		http.StatusMethodNotAllowed,
		http.StatusText(http.StatusMethodNotAllowed),
	)
}
