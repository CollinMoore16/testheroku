package main

import (
	"fmt"
	"net/http"
)

func main() {
	//define handler function
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `
		<html>
			<head>
				<title> My Website!</title>
			</head>	
			<body>
				<h1> Welcome to Yizer Country</h1>
				<p>Fun Fact! Did you know that we have 105 microbreweries?</p>
				<img src = "https://www.lonelyplanet.com/usa/pennsylvania/pittsburgh", height="250", width="250"
			</body>
		</html>
		
		`)
	})

	//Sports Page

	http.HandleFunc("/sports", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `
		<html>
			<head>
				<title> Sports</title>
			</head>	
			<body>
				<h1> Welcome to the Steelers Country</h1>
				<p>Kordell Stewart Days are back again OH YEA</p>
				<img src = "https://www.lonelyplanet.com/usa/pennsylvania/pittsburgh", height="250", width="250"
			</body>
		</html>
		
		`)
	})
	//Hobby page

	http.HandleFunc("/hobbies", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `
		<html>
			<head>
				<title> My Website!</title>
			</head>	
			<body>
				<h1> Welcome to Yizer Country</h1>
				<p>Fun Fact! Did you know that we have 105 microbreweries?</p>
				<img src = "https://www.lonelyplanet.com/usa/pennsylvania/pittsburgh", height="250", width="250"
			</body>
		</html>
		
		`)
	})

	//start the service

	fmt.Println("Server listening to port 8080...")
	http.ListenAndServe(":8080", nil)
}
