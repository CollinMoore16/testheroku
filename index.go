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
				<h1> Welcome to My website!</h1>
				<p>My name is Collin and i am currently a Junior at Point Park University!</p>
				<img src = "https://yt3.googleusercontent.com/ytc/AIdro_k2JN_PlW7iDX6Dm87p8cmF1PnnDtBY3wT6g9UvNQ=s900-c-k-c0x00ffffff-no-rj", height="250", width="250"
			</body>
		</html>
		
		`)
	})

	//Hobby Page

	http.HandleFunc("/hobbies", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `
		<html>
			<head>
				<title> Sports</title>
			</head>	
			<body>
				<h1> Welcome to the Steelers Country</h1>
				<p>Kordell Stewart Days are back again OH YEA</p>
				<img src = "PointParkPioneers.png", height="250", width="250"
			</body>
		</html>
		
		`)
	})

	//About page

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
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
