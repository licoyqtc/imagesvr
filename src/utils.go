package src

import (
	"github.com/bitly/go-simplejson"
	"fmt"
//	"encoding/json"
	"log"
	"encoding/json"
)

func Simplejson_unmarshal(data interface{}) (*simplejson.Json,error){

	log.Println("data : ", data)

	bdata , err := Json_marshal(data)

	if err == nil {
		return simplejson.NewJson(bdata)
	}
	return  nil,fmt.Errorf("unknown type")
}


func Json_unmarshal(data interface{} , out interface{}) error {

	log.Println("data : ",data)

	bdata , err := Json_marshal(data)

	if err == nil {
		err = json.Unmarshal(bdata , out)
		return err
	}

	return fmt.Errorf("unknown type")
}

func Json_marshal(data interface{}) ( []byte , error){

	if sdata,ok := data.(string) ; ok {
		return []byte(sdata) , nil
	}

	return json.Marshal(data)
}

