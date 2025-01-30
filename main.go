package main

import (
	"fmt"
	"syscall/js"

	"github.com/reiver/protogo-app/ui/base"
	"github.com/reiver/protogo-app/ui/people"
)

func main() {
	fmt.Println("ProToGo ⚡")

	var document js.Value
	{
		document = js.Global().Get("document")
		if document.IsUndefined() {
			fmt.Println("ERROR: undefined 'document'")
			return
		}
		if document.IsNull() {
			fmt.Println("ERROR: null 'document'")
			return
		}
		if js.TypeObject != document.Type() {
			fmt.Println("ERROR: 'document' not object")
			return
		}
	}

	{
		bodies := document.Call("getElementsByTagName", "body")
		if bodies.IsUndefined() {
			fmt.Println("ERROR: undefined <body>s")
			return
		}
		if bodies.IsNull() {
			fmt.Println("ERROR: null <body>s")
			return
		}
		if js.TypeObject != bodies.Type() {
			fmt.Println("ERROR: <body>s not objects")
			return
		}

		{
			var length int = bodies.Length()
			if length <= 0 {
				fmt.Printf("ERROR: no '<body>' (len=%d)\n", length)
				return
			}
		}
		body := bodies.Index(0)
		if body.IsUndefined() {
			fmt.Println("ERROR: undefined <body>")
			return
		}
		if body.IsNull() {
			fmt.Println("ERROR: null <body>")
			return
		}
		if js.TypeObject != body.Type() {
			fmt.Println("ERROR: <body> not object")
			return
		}

		body.Set("innerHTML", baseui.HTML)
	}
	fmt.Println("base ui created")

	var people []peopleui.Person
	{
//@TODO: replace dummy data with "real" data
			
		people = []peopleui.Person{
			peopleui.Person{
				Name: "Bert Dickson",
				Address: "@bert@soundwave.example",
				AvatarURI: "https://randomuser.me/api/portraits/men/99.jpg",
			},
			peopleui.Person{
				Name: "Bryce McDonald",
				Address: "@farmer@example.com",
				AvatarURI: "https://randomuser.me/api/portraits/men/10.jpg",
			},
			peopleui.Person{
				Name: "Carlos",
				Address: "@bigcarl@example.com",
				AvatarURI: "https://randomuser.me/api/portraits/men/95.jpg",
			},
			peopleui.Person{
				Name: "Charlotte Williams",
				Address: "@jane@smith.example",
				AvatarURI: "https://randomuser.me/api/portraits/women/20.jpg",
			},
			peopleui.Person{
				Name: "Gerry",
				Address: "@gm1990@monkey.example",
				AvatarURI: "https://randomuser.me/api/portraits/men/98.jpg",
			},
			peopleui.Person{
				Name: "Jane Dickson",
				Address: "@bunnyears@example.net",
				AvatarURI: "https://randomuser.me/api/portraits/women/97.jpg",
			},
			peopleui.Person{
				Name: "Jane Doe",
				Address: "@doe@deer.example",
				AvatarURI: "https://randomuser.me/api/portraits/women/17.jpg",
			},
			peopleui.Person{
				Name: "Linda",
				Address: "@linda@hadafarm.example",
				AvatarURI: "https://randomuser.me/api/portraits/women/99.jpg",
			},
			peopleui.Person{
				Name: "Margaret Collins",
				Address: "thedoctor@clinic.example",
				AvatarURI: "https://randomuser.me/api/portraits/women/98.jpg",
			},
		}
			
	}

	{
		const id string = baseui.IDContentPeople

		html, err := peopleui.Render(people)
		if nil != err {
			fmt.Printf("ERROR: problem rendering HTML for element with id=q\n", id)
			return
		}

		element := document.Call("getElementById", id)
		if element.IsUndefined() {
			fmt.Printf("ERROR: undefined element with id=q\n", id)
			return
		}
		if element.IsNull() {
			fmt.Printf("ERROR: null element with id=q\n", id)
			return
		}
		if js.TypeObject != element.Type() {
			fmt.Printf("ERROR: element with id=q not object\n", id)
			return
		}

		element.Set("innerHTML", html)
	}
	fmt.Println("people ui created")

	fmt.Println("now we wait...")
	<-make(chan struct{})
}
