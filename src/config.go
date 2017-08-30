package src

import (
    "encoding/json"
    "log"
    "os"
)

type Config struct {
    ListenAddr string
    Storage string
    Domain string
}

var Conf Config

func Loadconf(){
	r, err := os.Open("conf/config.json")
    if err != nil {
        log.Fatalln(err)
    }
    decoder := json.NewDecoder(r)
    err = decoder.Decode(&Conf)
    if err != nil {
        log.Fatalln(err)
    }
}