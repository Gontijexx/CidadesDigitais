package util

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
)

//BodyToBytes converts raw json to slice of bytes
func BodyToBytes(body io.Reader) (bytes []byte, err error) {

	bytes, err = ioutil.ReadAll(body)
	if err != nil {
		log.Printf("[WARN] problem converting body, because, %v\n", err)
		return
	}
	return
}

//BytesToStruct parse a json into a struct
func BytesToStruct(bytes []byte, entity interface{}) (err error) {

	err = json.Unmarshal(bytes, &entity)
	if err != nil {
		log.Printf("[WARN] problem parsing json body to struct, because, %v\n", err)
		return
	}
	return
}

//StructToBytes parse a struct to a json
func StructToBytes(entity interface{}) (bytes []byte, err error) {

	var b []byte

	b, err = json.Marshal(&entity)
	if err != nil {
		log.Printf("[WARN] problem parsing struct to json body, because, %v\n", err)
		return
	}
	return b, err

}

//teste de commit
