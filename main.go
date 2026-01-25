package main

//“It’s a small Go HTTP service with in-memory state, background order processing using goroutines, context-based cancellation, explicit error handling, and graceful shutdown.”

import "fmt"

import (
	"flag"
	"net/http"
)

func handleGetUser(w http.ResponseWriter, r *http.Request) {}
func handleGetAccount(w http.ResponseWriter, r *http.Request) {}

func main() {

	listenAddr:= flag.String("listenaddr", ":49999", "ice cream")
	flag.Parse()

	http.HandleFunc("/user", handleGetUser)
	http.HandleFunc("/account", handleGetAccount)

	http.ListenAndServe(*listenAddr, nil);

	fmt.Println("--- Pupu opened the shop ---")
	fmt.Println("--- Pupu have seen your order ---")
	fmt.Println("--- Pupu is processing your ice cream orders ---")
	fmt.Println("--- You chose to pay by cash! Baobao will deliver to you. Prepare your money! ---")
	fmt.Println("--- Delivering... Baobao is on his way! ---")
	fmt.Println("--- You bought 5 boxes of ice cream. You get 5 minutes hug with Baobao ---")
	fmt.Println("--- Hope you enjoyed our service. See you again! ---")

}