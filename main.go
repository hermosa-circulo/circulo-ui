package main

import (
    "fmt"
    "net/http"
    "time"
    "strings"
    "strconv"
    //"encoding/json"

    ga "github.com/hermosa-circulo/circulo-tools/igademo"
    "github.com/hermosa-circulo/circulo-tools/api"
)

func CreatePop(w http.ResponseWriter, r *http.Request) {
    group := "test"
    ga.CreatePopulationDB(group,80,100)
    fmt.Fprintf(w, "Population is Created")
}

func UpdatePop(w http.ResponseWriter, r *http.Request) {
    ans := [][]int{{}}
    for i:=0;i<100;i++ {
        ans[0] = append(ans[0],1)
    }
    group := "test"
    ga.UpdatePopulationDB(group,10,ans)
    fmt.Fprintf(w, "Population is Updated")
}

func PicGene(w http.ResponseWriter, r *http.Request) {
    group := "test"
    pic := []int{1,3,4,5}
    picstr := ga.PicGene(group,pic)
    fmt.Fprintf(w,picstr)
}

type Axis struct {
    X int
    Y int
    Z int
}

func Cube(w http.ResponseWriter, r *http.Request) {
    str_url := strings.Split(r.URL.Path,"cube/")
    asix := strings.Split(str_url[1],"-")
    var ax Axis
    ax.X, _ = strconv.Atoi(asix[0])
    ax.Y, _ = strconv.Atoi(asix[1])
    ax.Z, _ = strconv.Atoi(asix[2])
    /*r.ParseForm()
    fmt.Println(r.Form)
    for key, _ := range r.Form {
        err := json.Unmarshal([]byte(key), &ax)
        if err != nil {
            fmt.Println(err.Error())
        }
    }
    fmt.Println(ax)
    */
    str := "g cube\n"
    cube := [][]int{{0,0,0},{10,0,0},{0,10,0},{10,10,0},{0,0,10},{10,0,10},{0,10,10},{10,10,10}}
    for i:=0;i<len(cube);i++ {
        str = str + "v"
        str = str + " " + strconv.Itoa(cube[i][0]+ax.X)
        str = str + " " + strconv.Itoa(cube[i][1]+ax.Y)
        str = str + " " + strconv.Itoa(cube[i][2]+ax.Z)
        str = str + "\n"
    }
    str = str + "vn 0 0 -1\nvn -1 0 0\nvn 1 0 0\nvn 0 -1 0\nvn 0 1 0\nvn 0 0 1\nf 1//1 3//1 4//1 2//1\nf 1//2 5//2 7//2 3//2\nf 2//3 4//3 8//3 6//3\nf 1//4 2//4 6//4 5//4\nf 3//5 7//5 8//5 4//5\nf 5//6 6//6 8//6 7//6\n"
   fmt.Fprintf(w,str)
}


func main() {
    create := CreatePop
    update := UpdatePop
    picgene := PicGene
    cube := Cube
    router := api.NewRouter()
    router.HandleFunc("/api/createpop", create)
    router.HandleFunc("/api/updatepop", update)
    router.HandleFunc("/api/picgene", picgene)
    router.HandleFunc("/api/cube", cube)
    //fmt.Println(router.routes)
    srv := &http.Server{
            Handler:      router,
            Addr:         "0.0.0.0:8000",
            WriteTimeout: 15 * time.Second,
            ReadTimeout:  15 * time.Second,
    }
    srv.ListenAndServe()
    fmt.Println("ok")
}
