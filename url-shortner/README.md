# Endpoints

##### /shorten
- A POST endpoint for shortening URLs. You can send a POST request to this endpoint with a form parameter named "url" containing the original URL.

##### /:shortKey 
- An endpoint for redirecting to the original URL. If you access this endpoint with a short key, it will redirect to the corresponding original URL.

# Testing with Postman
-Set the request method to POST.
-Enter the URL of your URL shortener endpoint (e.g., http://localhost:8080/shorten).
-Go to the "Body" tab and choose "x-www-form-urlencoded" as the body type.
-Add a key-value pair with the key "url" and the value being the URL you want to shorten.

# Additional Features
Feel free to add additional features such as persistence, error handling, and security considerations.