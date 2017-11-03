package ga

import (
    "fmt"
    "net/http"
    "encoding/json"
    "io/ioutil"
)

func CreatePopAPI(w http.ResponseWriter, r *http.Request) {
    group := "test"
    CreatePopulationDB(group,80,125)
    fmt.Fprintf(w, "Population is Created")
}

func UpdatePopAPI(w http.ResponseWriter, r *http.Request) {
    ans := []byte{}
    for i:=0;i<125;i++{
        ans = append(ans,0xff)
    }
    group := "test"
    UpdatePopulationDB(group,10,ans)
    fmt.Fprintf(w, "Population is Updated")
}

type Pic struct {
    Point int
}

func PicGeneAPI(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
        fmt.Fprintf(w,"Not POST ..")
        return
    }
    body, e := ioutil.ReadAll(r.Body)
    if e != nil {
        fmt.Fprintf(w, e.Error())
        fmt.Println(e.Error())
        return
    }
    var pic_point Pic
    e = json.Unmarshal(body, &pic_point)
    if e != nil {
        fmt.Fprintf(w, e.Error())
        fmt.Println(e.Error())
        return
    }
    group := "test"
    fmt.Println(pic_point.Point)
    fmt.Println(r.Form)
    picstr := PicGene(group,pic_point.Point)
    fmt.Fprintf(w,picstr)
}
