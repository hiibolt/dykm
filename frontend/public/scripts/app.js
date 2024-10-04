console.log("Hello, Alex!");

// JavaScript to handle the form submission
document.getElementById('TallyForm').addEventListener('submit', function(event) {
    event.preventDefault(); // Prevent the form from redirecting the page
    console.log("Form submitted!");

    // Collect the form data
    const form = event.target;
    const formData = new FormData(form);

    // Send the form data using the Fetch API
    fetch('/api/tally', {  // The backend URL
        method: 'POST',
        body: formData
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