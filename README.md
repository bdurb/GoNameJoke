# Go Webservice

Go web service combining uinames api and api.icndb.com
and returning a joke with the username inserted. 

## To run:

from the root of the project run `go run main.go`

This will serve the content to `http://localhost:5000`

This program utalizes a struct to store the name and surname, and a struct with a nested struct to obtain the joke as it's a nested json structure.

requests are made via http, and the json is parsed using json.Unmarshal

Headers are set to `text/html; charset=utf-8` as some of the joke responses have 
quotation.

all imports are from the standard library, and no third party libraries were used.

This application has been tested via:
curl
navigating to localhost:5000 on local machine
PostMan.



