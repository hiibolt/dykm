console.log("Hello, Burger!");

document.getElementById('TallyForm').addEventListener('submit', function(event) {
    event.preventDefault(); // Prevent the form from redirecting the page
    console.log("Form submitted!");

    // Collect the form data from the form elements
    const form = event.target;  // This is the form that triggered the event
    const formData = new FormData(form);  // Automatically picks up all form inputs

    // Logging form data for debugging
    for (const entry of formData.entries()) {
        console.log(entry[0], entry[1]);
    }
    
    // Send the form data using the Fetch API
    fetch('/api/tally', {  // The backend URL
        method: 'POST',
        body: formData,
        headers: {
            "Content-Type": "application/x-www-form-urlencoded"
        }
        // No need to manually set headers for FormData
    })
    .then(response => response.json()) // Parse the response as JSON
    .then(data => {
        // Display the response or handle it as needed
        document.getElementById('response').innerText = "Form submitted successfully!";
    })
    .catch(error => {
        // Handle any errors
        document.getElementById('response').innerText = "Error submitting the form.";
        console.error('Error:', error);
    });
});