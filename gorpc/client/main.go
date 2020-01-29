package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/sirwaithaka/gorpc/proto"
	"google.golang.org/grpc"
)

func main() {
	// Create a connection to the grpc server and handle errors
	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	// build a client out of the grpc connection
	client := proto.NewAddServiceClient(conn)
	ctx := context.Background()

	// create a http router to handle http Requests
	router := httprouter.New()
	router.GET("/add/:a/:b", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		// parse the parameter into uint64 type
		a, err := strconv.ParseUint(p.ByName("a"), 10, 64)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(struct{ message string }{"Could not get parameter A"})

		}

		// parse the parameter into uint64 type
		b, err := strconv.ParseUint(p.ByName("b"), 10, 64)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(struct{ Message string }{"Could not get parameter B"})
		}

		// Build a protobuf request object with the params
		req := &proto.Request{A: int64(a), B: int64(b)}
		// Call the remote Add procedure and get the reponse
		if response, err := client.Add(ctx, req); err == nil {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(struct{ Result string }{fmt.Sprint(response.Result)})
		} else {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(struct{ Message string }{err.Error()})
		}
	})

	router.GET("/mult/:a/:b", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		a, err := strconv.ParseUint(p.ByName("a"), 10, 64)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(struct{ message string }{"Could not get parameter A"})

		}

		b, err := strconv.ParseUint(p.ByName("b"), 10, 64)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(struct{ Message string }{"Could not get parameter B"})
		}

		req := &proto.Request{A: int64(a), B: int64(b)}
		if response, err := client.Multiply(ctx, req); err == nil {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(struct{ Result string }{fmt.Sprint(response.Result)})
		} else {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(struct{ Message string }{err.Error()})
		}
	})

	log.Fatal(http.ListenAndServe(":8081", router))
}
