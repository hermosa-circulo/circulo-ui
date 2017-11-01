package main

import (
    "fmt"
    "net/http"
    "time"
    "strings"
    "strconv"

    ga "github.com/hermosa-circulo/circulo-ui/igademo"
    "github.com/hermosa-circulo/circulo-ui/api"
)

type Axis struct {
    X []int
    Y []int
    Z []int
}

func Cube(w http.ResponseWriter, r *http.Request) {
    str_url := strings.Split(r.URL.Path,"cube/")
    three_asixs := strings.Split(str_url[1],"/")
    var ax Axis
    for _, three_asix := range three_asixs {
	asix := strings.Split(three_asix,"~")
        x, _ := strconv.Atoi(asix[0])
	ax.X = append(ax.X,x)
        y, _ := strconv.Atoi(asix[1])
	ax.Y = append(ax.Y,y)
        z, _ := strconv.Atoi(asix[2])
	ax.Z = append(ax.Z,z)
    }

    str := "g cube\n"
    cube := [][]int{{0,0,0},{1,0,0},{0,1,0},{1,1,0},{0,0,1},{1,0,1},{0,1,1},{1,1,1}}
    for i:=0;i<len(ax.X);i++ {
        for j:=0;j<len(cube);j++ {
            str = str + "v"
            str = str + " " + strconv.Itoa(cube[j][0]+ax.X[i])
            str = str + " " + strconv.Itoa(cube[j][1]+ax.Y[i])
            str = str + " " + strconv.Itoa(cube[j][2]+ax.Z[i])
            str = str + "\n"
        }
    }
    str = str + "vn 0 0 -1\nvn -1 0 0\nvn 1 0 0\nvn 0 -1 0\nvn 0 1 0\nvn 0 0 1\n"
    f:=""
    for i:=0;i<len(ax.X);i++ {
        f = f + "f "+strconv.Itoa((i*8)+1)+"//1 "+strconv.Itoa((8*i)+3)+"//1 "+strconv.Itoa((8*i)+4)+"//1 "+strconv.Itoa((8*i)+2)+"//1\n"
        f = f + "f "+strconv.Itoa((i*8)+1)+"//2 "+strconv.Itoa((i*8)+5)+"//2 "+strconv.Itoa((i*8)+7)+"//2 "+strconv.Itoa((i*8)+3)+"//2\n"
        f = f + "f "+strconv.Itoa((i*8)+2)+"//3 "+strconv.Itoa((i*8)+4)+"//3 "+strconv.Itoa((i*8)+8)+"//3 "+strconv.Itoa((i*8)+6)+"//3\n"
        f = f + "f "+strconv.Itoa((i*8)+1)+"//4 "+strconv.Itoa((i*8)+2)+"//4 "+strconv.Itoa((i*8)+6)+"//4 "+strconv.Itoa((i*8)+5)+"//4\n"
        f = f + "f "+strconv.Itoa((i*8)+3)+"//5 "+strconv.Itoa((i*8)+7)+"//5 "+strconv.Itoa((i*8)+8)+"//5 "+strconv.Itoa((i*8)+4)+"//5\n"
        f = f + "f "+strconv.Itoa((i*8)+5)+"//6 "+strconv.Itoa((i*8)+6)+"//6 "+strconv.Itoa((i*8)+8)+"//6 "+strconv.Itoa((i*8)+7)+"//6\n"
   }
   str = str + f
   fmt.Fprintf(w,str)
}


func main() {
    create := ga.CreatePopAPI
    update := ga.UpdatePopAPI
    picgene := ga.PicGeneAPI
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
