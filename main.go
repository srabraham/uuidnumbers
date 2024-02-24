package main

import (
	"github.com/google/uuid"
	"log"
	"math/big"
	"strings"
)

func toInt128(uu uuid.UUID) *big.Int {
	myUUID := uu.String()
	// remove the version
	myUUID = myUUID[:14] + myUUID[15:]
	//log.Printf("uuid without the version in it: %v", myUUID)
	// remove the dashes
	myUUID = strings.Replace(myUUID, "-", "", 4)
	//log.Printf("uuid as hex: %v", myUUID)
	var i big.Int
	i.SetString(myUUID, 16)
	return &i
}

func main() {
	for range 100 {
		myUUID := uuid.Must(uuid.NewV7())
		log.Printf("got uuid: %v", myUUID)
		log.Printf("uuid as int: %v", toInt128(myUUID).String())
		log.Println("------------")
	}
}
