package main;

import (
	"fmt"
	"net/http"
	"strings"
)

func page_handler(w http.ResponseWriter, r *http.Request) {
	// Split the path into segments
	path_segments := strings.Split(r.URL.Path, "/");

	// Home Directory
	if len(path_segments) < 2 || path_segments[1] == "" {
		var final_html_result Result[string] = inject_component("index.html");

		// Check for errors
		if final_html_result.IsErr() {
			ReturnErr(w, 400, "Bad Request", final_html_result.UnwrapErr());
			return;
		}

		var final_html string = final_html_result.UnwrapOk();

		// Serve the respective page
		fmt.Fprintf(w, final_html);

		return
	}

	switch path_segments[1] {
	case "about", "app":
		var final_html_result Result[string] = inject_component(path_segments[1]+".html");

		// Check for errors
		if final_html_result.IsErr() {
			ReturnErr(w, 400, "Bad Request", final_html_result.UnwrapErr());
			return;
		}

		var final_html string = final_html_result.UnwrapOk();

		// Serve the respective page
		fmt.Fprintf(w, final_html);
	default:
		var final_html_result Result[string] = inject_component("404.html");

		// Check for errors
		if final_html_result.IsErr() {
			ReturnErr(w, 400, "Bad Request", final_html_result.UnwrapErr());
			return;
		}

		var final_html string = final_html_result.UnwrapOk();

		// Serve the 404 page
		fmt.Fprintf(w, final_html);
	}
}