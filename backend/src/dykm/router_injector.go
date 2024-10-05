package main;

import (
	"os"
	"strings"
)

func inject_component ( file_name string ) Result[string] {
	// Define buffers for the HTML data and any errors that may occur
	var html_data []byte;
	var html_load_err error;

	// Load the original HTML file
	html_data, html_load_err = os.ReadFile("../../../frontend/pages/" + file_name);
	if html_load_err != nil {
		return Err[string]("Error reading file '" + file_name + "'");
	}

	// Loop through all files in the components directory
	var component_dir string = "../../../frontend/components/";

	// Open the directory
	files, err := os.ReadDir(component_dir)
    if err != nil {
		return Err[string]("Error reading directory '" + component_dir + "'");
    }

	var final_html string = string(html_data);

	// Loop through all files in the directory
	var loop_counter int = 0;
	for {
		var made_change = false;
		loop_counter++;

		// Break if the loop has run too many times
		if loop_counter > 20 {
			return Err[string]("20 levels of recursion, potential circular dependency detected!");
			break;
		}

		for _, file := range files {
			var file_name string = file.Name();
			var file_path string = component_dir + file_name;

			var file_data []byte;
			var file_load_err error;

			// Load the file
			file_data, file_load_err = os.ReadFile(file_path);

			// Check for errors
			if file_load_err != nil {
				return Err[string]("Error reading file '" + file_name + "'");
			}

			// Convert the data to a string
			var file_str string = string(file_data);

			// Replace the component tag with the component data
			var final_html_copy string = final_html;
			final_html = strings.ReplaceAll(final_html, "<!--~"+file_name+"~-->", file_str);

			// Check if the file was changed
			if final_html != final_html_copy {
				made_change = true;
			}
		}

		// Break if no changes were made
		if !made_change {
			break;
		}
	}

	return Ok(final_html);
}