console.log("Hello, Bolt!");

document.getElementById('TallyForm').addEventListener('submit', function(event) {
    event.preventDefault(); // Prevent the form from redirecting the page

    document.getElementById('response').innerText = "Waiting for response...";
    document.getElementById('results').style.display = "none";

    // Collect the form data from the form elements
    const form = event.target;  // This is the form that triggered the event
    const formData = new FormData(form);
    
    // Send the form data using the Fetch API
    fetch('/api/tally', {  // The backend URL
        method: 'POST',
        body: formData
        // No need to manually set headers for FormData
    })
        .then(response => {
            // Verify that the response is OK
            if (!response.ok) {
                throw new Error('Failed to submit form');
            }

            return response.json()
        }) // Parse the response as JSON
        .then(data => {
            console.log(data)
            // 

            document.getElementById('results').style.display = "block";

            // Display the response or handle it as needed
            document.getElementById('response').innerText = "Form submitted successfully!";
            document.getElementById('usernames').innerText = "Usernames: " + data.usernames; 
            document.getElementById('emails').innerText = "Emails: " + data.emails;
            document.getElementById('phones').innerText = "Phones: " + data.phones;
            document.getElementById('hashes').innerText = "Hashes: " + data.hashes;
            document.getElementById('ips').innerText = "Ips: " + data.ips;
            document.getElementById('names').innerText = "Names: " + data.names;
            document.getElementById('passwords').innerText = "Passwords: " + data.passwords;
        })
        .catch(error => {
            // Handle any errors
            document.getElementById('response').innerText = "Error submitting the form.";
            console.error('Error:', error);
        });
});