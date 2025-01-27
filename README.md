# protogo-app

**protogo-app** is a mobile front-end for **protogo**.
**protogo** is a career, gig, and business focused decentralized social-media (DeSo) platform, for the Fediverse and Social Web.

## Technology

**protogo-app** is a cross-platform mobile app and progressive-web-app (PWA).

It is written in the Go programming-language (Go) and is compiled to WebAssembly (WASM).
The Go code creates the UI by adding HTML and CSS using DOM manipulation.

## Building

To build **protogo-app** run a command similar to the following:

```
GOOS=js GOARCH=wasm go build -o main.wasm
```

## Author

Software **protogo** was written by [Charles Iliya Krempeaux](http://reiver.link)
