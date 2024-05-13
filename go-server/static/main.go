package main		

import ("fmt"
"log"
"net/http")


func HelloHandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/hello"{
		httpe.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.method != "GET"{
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(w, "Hello!")
}

func formHandler(w http.ResponseWriter, r *http.Request)
	if err := r.ParseForm(); err != nil{
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	name := r.FormValue("name")
	address:= r.FormValue("address")
	fmt.Fprintf(w, "Hello %s!", name)
	fmt.Fprintf(w, "Your address is %s", address)
}

func main(){
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", HelloHandler)

	fmt.print("Starting server port at 8000\n")
	if err := http.listenAndServe(":8000", nil); err != nil{
		log.Fatal(err)
	}
}