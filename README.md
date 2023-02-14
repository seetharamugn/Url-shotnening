   ## Documentation

The documentation for the URL shortener service should cover the following topics:

   ## Overview of the service and its purpose.
  
    How to run the service.
    How to use the service.
    API documentation.
    Test cases.
    Here's an example of how you can write the documentation:


  ## URL Shortener Service

The URL shortener service is a web service that allows you to shorten URLs. The service uses a database to store the original URLs and their corresponding shortened URLs.
    
    
  ## To run the URL shortener service, follow these steps: 

      Clone the repository.
      Install the required dependencies.
      Configure the database connection.
      Run the service using the go run command.
      How to Use the Service
      
 ## To use the URL shortener service, follow these steps:

Send a POST request to the /shorten endpoint with the original parameter set to the URL you want to shorten.
The service will return a JSON object with the shortened URL.
To use the shortened URL, send a GET request to the endpoint with the shortened URL as the path parameter.
The service will redirect you to the original URL.

 ## API Documentation
**POST /shorten**

Shorten a URL.

**Request Body**

| Parameter | Type | Description |
| --- | --- | --- |
| original | string | The URL to be shortened. |



**Response Body**

| Parameter | Type | Description |
| --- | --- | --- |
| shortened	| string	| The shortened URL. |
| OriginalUrl | String | The Original URL |
| ExpiredAt | Integer | The shortened URL is expired in specific time |

**GET /:shortened**

Redirect to the original URL.

## Test Cases

The URL shortener service includes the following test cases:

      Test that a valid URL can be successfully shortened.
      Test that an invalid URL returns an error.
      Test that a shortened URL redirects to the original URL.
      Test that an expired shortened URL returns a 410 Gone error.
      Test that a maximum of 20000 shortened URLs can be stored in the database.
