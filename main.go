package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"strconv"
	"time"
)

// DATABASE GECKO SCHEMES:
type VehicleMongoScheme struct {
	Name   string            `bson:"name" json:"name"`
	Fields map[string]string `bson:"fields" json:"fields"`
}

// DATABASE VEHICLES
type Vehicle struct {
	Id     bson.ObjectId `bson:"_id,omitempty" json:"objectid"`
	Scheme string        `bson:"scheme" json:"data"`
	Data   VehicleData   `bson:"data" json:"data"`
}

// To store data in a scheme, we map the field name to a value using Go maps
type VehicleData map[string]string

func main() {
	var theDbSchemes []VehicleMongoScheme // To get the schemes from the database

	// Open MongoDB server connection
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	c := session.DB("test").C("Vehicle")
	d := session.DB("test").C("VehicleSchemes")

	// First, get all the schemes from the schemes database
	err = d.Find(bson.M{}).All(&theDbSchemes)
	if err != nil {
		fmt.Println("Error occured getting the schemes")
		panic(0)
	}
	fmt.Printf("Fetched %d schemes\n", len(theDbSchemes))

	// Now, insert new vehicles according to the schemes that we just got from the db
	for counter := 0; ; counter++ {
		fmt.Println("Vehicle ", counter)

		// Fill the Data field with a random schema
		theId := bson.NewObjectId()
		theSchemeNr := time.Now().Nanosecond() % len(theDbSchemes)

		theData := make(map[string]string)
		// You could do a switch here for the scheme, but I'm just running over all the fields
		for i := 1; i < len(theDbSchemes[theSchemeNr].Fields)+1; i++ {
			theFieldNumber := strconv.Itoa(i)
			theFieldName := theDbSchemes[theSchemeNr].Fields[theFieldNumber]
			theValue := "Test" + strconv.Itoa(counter)

			theData[theFieldName] = theValue
		}

		theVehicle := Vehicle{Id: theId, Scheme: theDbSchemes[theSchemeNr].Name, Data: theData}

		err = c.Insert(theVehicle)
		if err != nil {
			fmt.Println("Error occured inserting vehicle")
			panic(0)
		}
	}
}
