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
				<title> My Hobbies </title>
			</head>	
			<body>
				<center>
				<h1>
				On this page i will tell you about my hobbies</h1>
				<p>I like to play guitar my favorite kind of guitar is a Fender Telecaster<br>
				<center>
				<img src = "https://i.ebayimg.com/images/g/gggAAOSw~ytjIeBY/s-l1200.webp", height="450", width="450" <br> <br> <br>
				I like to watch sports like the steelers, penguins, and the buccs sometimes<br>
				I build computers for my friends and basically do support for them<br>
				I love to listen to music<br>
				I also love to play video games <br>
				Last but not least I like to go on long walks and hangout with friends outside.
				</p>
			</body>
		</html>
		
		`)
	})

	//About page

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `
		<html>
			<head>
				<title> About Me </title>
			</head>	
			<center>
			<body>
				<h1> This page is about me! </h1>
				<p>
				I am from a small town with less than 3000 people called Sharpsville! <br> 
				and i grew up playing sports but due to my knees and ankles being made of jelly <br>
				Ive since had to stop playing. I took a huge interest in computers which brought me <br>
				down the career path of Computer Science. I have a cat named munchkin she is a orange <br> 
				Tabby and i got her when i was in 8th grade i love that cat so much! she looks like the cat <br> 
				pictured below!
				</p>
				<img src = "https://upload.wikimedia.org/wikipedia/commons/thumb/4/4b/Dilute_orange_tabby_cat.jpg/220px-Dilute_orange_tabby_cat.jpg", height="250", width="250"
			</body>
		</html>
		
		`)
	})

	//start the service

	fmt.Println("Server listening to port 8080...")
	http.ListenAndServe(":8080", nil)
}
