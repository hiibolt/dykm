/*

<div id="not_logged_in">
    <p>Enter API Key:</p>
    <input type="text" id="api_key_input" />
    <button id="set_key_button">Set Key</button>
    <button id="get_key_button" onclick="location.href='/pricing'">Get Key</button>
</div>
<div id="logged_in" style="display:none">
    <p>API Key: <span id="api_key_display">************</span></p>
    <p>Balance: <span id="api_key_balance"></span></p>
    <button id="reset_key_button">Reset Key</button>
</div>


*/

// Checks if there is an API key stored in the browser's local storage
function hasApiKey() {
    return localStorage.getItem('User-API-Key') !== null;
}

// Sets the API key in the browser's local storage
function setApiKey(apiKey) {
    localStorage.setItem('User-API-Key', apiKey);
}

// Gets the API key from the browser's local storage
function getApiKey() {
    return localStorage.getItem('User-API-Key');
}

async function checkApiKey(apiKey) {
    try {
        const formData = new FormData();
    
        // Add the API key to the form data
        formData.append('User-API-Key', apiKey);
    
        let res = await fetch('/api/user/get', {  // The backend URL
            method: 'POST',
            body: formData
        });

        if (!res.ok) {
            return false;
        }
        
        return true;
    } catch (e) {
        return false;
    }
}

// Gets the balance from the backend
async function getBalance() {
    const formData = new FormData();

    // Add the API key to the form data
    formData.append('User-API-Key', getApiKey());

    let res = await fetch('/api/user/get', {  // The backend URL
        method: 'POST',
        body: formData
    });
    
    // Verify that the response is OK
    if (!res.ok) {
        throw new Error('Failed to submit form');
    }

    let data = await res.json();

    console.log(data)
    document.getElementById('api_key_balance').innerText = data.balance;

    return data.balance;
}

// Refreshes the state of the page based on the API key stored in the browser's local storage
async function refresh_state(){
    // If there is an API key stored in the browser's local storage, hide the "not_logged_in" div
    if (hasApiKey()) {
        document.getElementById('not_logged_in').style.display = "none";
        document.getElementById('logged_in').style.display = "block";
    
        // Set the field in the form to the API key, but
        //  only the first 4 characters followed by asterisks
        document.getElementById('api_key_display')
            .innerText = getApiKey().substring(0, 4) + "*".repeat(getApiKey().length - 4);

        // Fetch the balance from the backend
        if ( !checkApiKey(getApiKey()) ) {
            document.getElementById('error_message').innerText = "`" + getApiKey() + "` is an invalid API Key!";
            return;
        }

        try {
            await getBalance();
        } catch (error) {
            console.log(error);
            return;
        }

        document.getElementById('error_message').innerText = "";
    } else {
        document.getElementById('not_logged_in').style.display = "block";
        document.getElementById('logged_in').style.display = "none";
    }
}

// When the "set_key_button" is clicked, set the API key in the browser's local storage
document.getElementById('set_key_button').addEventListener('click', function() {
    const apiKey = document.getElementById('api_key_input').value;

    if (apiKey === "") {
        return;
    }

    checkApiKey(apiKey)
        .then((res) => {
            console.log("API Key set: " + res);

            if (!res) {
                document.getElementById('error_message').innerText = "`" + apiKey + "` is an invalid API Key!";
            } else {
                document.getElementById('error_message').innerText = "";
                setApiKey(apiKey);
                refresh_state();
            }
        })
        .catch((e) => {
            console.log(e);
        });
});

// When the "reset_key_button" is clicked, remove the API key from the browser's local storage
//  and 
document.getElementById('reset_key_button').addEventListener('click', function() {
    localStorage.removeItem('User-API-Key');

    console.log("API Key removed.");

    refresh_state();
});

console.log("Hello, Bolt, again!! Logged in: " + hasApiKey());


refresh_state();