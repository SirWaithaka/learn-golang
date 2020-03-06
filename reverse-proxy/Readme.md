# Reverse Proxy for Go

I was once required to render html from a next server located remotely, from an iot device
that can be located anywhere.

The idea here was to make the next server decoupled from the
iot devices, such that updates to the Next React application can be done instantly without
and the iot devices can see the changes at the same time in an instant.

Solution was to create a reverse proxy on the IOT device server. Requests that match
a certain url are served by the proxy server.

## Fasthttp Reverse Proxy
The application uses gofiber/fiber as the server which has fasthttp as a dependency.
Implementing a reverse server in go standard libraries is really easy. Look at main.go file.

An implementation of the reverse proxy in fasthttp can be run and tested. Make sure you
have a server running on localhost:3000. Then run ...

```bash

$ go run reverse.go

```

The server starts on port `2720`

Open your browser and visit any url hosted on `localhost:3000/some-url` by entering
`localhost:2720/some-url`.


## Updates
Any updates to the reverse proxy will be added here in case of features.

