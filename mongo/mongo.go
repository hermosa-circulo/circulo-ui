package main

import (
    "fmt"
    mgo "gopkg.in/mgo.v2"
    //"gopkg.in/mgo.v2/bson"
    "github.com/hermosa-circulo/circulo-tools/igademo"
    //"../igademo"
)

type Person struct {
        Group string
        Gene string
}

func CreatePopulationGroup( group string, size int, length int) {
    population := igademo.CreatePopulation(size, length)
    fmt.Println(population)
}

func main() {
    test := "test"
    CreatePopulationGroup(test, 10, 10 )
    session, _ := mgo.Dial("mongodb://localhost:28001/test")
    db := session.DB("test").C("Gene")
    defer session.Close()
    /*

    err := db.Insert(&Person{"Ale", "+55 53 8116 9639"},
                       &Person{"Cla", "+55 53 8402 8510"})
    if err != nil {
        fmt.Println(err)
    }
    result := Person{}
    err = db.Find(bson.M{"group": "Ale"}).One(&result)
    if err != nil {
        fmt.Println(err)
    }

    fmt.Println("Phone:", result.Gene)
    */
}
