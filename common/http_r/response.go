package httpr

import (
	"encoding/json"
	"html/template"
	"net/http"
	"time"
)

func Template(
	w http.ResponseWriter,
	statusCode int,
	execute string,
	t *template.Template,
	data any,
) error {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(statusCode)

	if err := t.ExecuteTemplate(w, execute, data); err != nil {
		panic(err)
	}

	return nil
}

// DeleteCookieAndRedirect sets a zero value and expired time to a cookie which will remove it from
// the browser. For this to work, the given cookie has to have a the Path property set to "/"
//
// https://developer.mozilla.org/en-US/docs/Web/HTTP/Cookies
func DeleteCookieAndRedirect(
	w http.ResponseWriter,
	r *http.Request,
	url string,
	cookie string,
	statusCode int,
	now time.Time,
) {
	http.SetCookie(w, &http.Cookie{
		Name:     cookie,
		Value:    "",
		Expires:  now.Add(-time.Hour),
		Secure:   true,
		Path:     "/",
		HttpOnly: true,
	})

	http.Redirect(w, r, url, statusCode)
}

// JSON converts a Go value to JSON and sends it to the client.
func JSON(
	w http.ResponseWriter,
	statusCode int,
	data any,
) error {

	// If there is nothing to marshal then set status code and return.
	if statusCode == http.StatusNoContent || data == nil {
		w.WriteHeader(statusCode)

		return nil
	}

	// Convert the response value to JSON.
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// Set the content type and headers once we know marshaling has succeeded.
	w.Header().Set("Content-Type", "application/json")

	// Write the status code to the response.
	w.WriteHeader(statusCode)

	// Write response data to response body.
	if _, err := w.Write(jsonData); err != nil {
		return err
	}

	return nil
}

func Text(
	w http.ResponseWriter,
	statusCode int,
	data string,
) error {

	// If there is nothing to marshal then set status code and return.
	if statusCode == http.StatusNoContent || data == "" {
		w.WriteHeader(statusCode)
		return nil
	}

	// Set the content type and headers once we know marshaling has succeeded.
	w.Header().Set("Content-Type", "text/plain")

	// Write the status code to the response.
	w.WriteHeader(statusCode)

	// Write response data to response body.
	if _, err := w.Write([]byte(data)); err != nil {
		return err
	}

	return nil
}
