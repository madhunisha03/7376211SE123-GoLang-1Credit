package main_3

import (
	"fmt"
	"net/http"
)

func main() {
	routes := map[string]func(http.ResponseWriter, *http.Request){
		"/":         home,
		"/products": product,
		"/services": service,
	}
	for path, handler := range routes {
		http.HandleFunc(path, handler)
	}
	go func() {
		fmt.Println("Server is listening on port 8080...")
		if err := http.ListenAndServe(":8080", nil); err != nil {
			fmt.Println("Error:", err)
		}
	}()
	fmt.Println("Press Ctrl+C to stop the server.")
	select {}
}
func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Our Website! Explore our products and services.")
}
func product(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Products: [Laptop, Keyboard, Notepad]")
}
func service(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Our Services: [Service A, Service B, Service C]")
}
