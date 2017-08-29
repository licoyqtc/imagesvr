package main

import(
	"fmt"
	"net/http"
	"log"
	"github.com/gorilla/mux"
	"imagesvr/src"
)
	
func main(){
	fmt.Println("Quick Image Server.")
	src.Loadconf()
	r := mux.NewRouter()
	r.HandleFunc("/image/home", src.HomeHandler).Methods("GET")
	r.HandleFunc("/image/upload",src.UploadHandler).Methods("POST")
	r.HandleFunc("/image/{imgid}",src.DownloadHandler).Methods("GET")

    err := http.ListenAndServe(src.Conf.ListenAddr, r)
    if err != nil {
        log.Fatal("ListenAndServe error: ", err)
    }
}
