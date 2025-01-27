package main

import (
	"fmt"
	"syscall/js"

	"github.com/reiver/protogo-app/ui"
)

func main() {
	fmt.Println("ProToGo ⚡")

	{
		document := js.Global().Get("document")
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

		body.Set("innerHTML", ui.Main)
	}
	fmt.Println("ui created")

	fmt.Println("now we wait...")
	<-make(chan struct{})
}
