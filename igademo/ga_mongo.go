package ga

import (
    "fmt"
    "strings"
    "strconv"

    mgo "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
)

type Person struct {
        Group string
        Gene string
}
/////// Mongo///

func PicPopulation(c *mgo.Collection, group string) []Gene{
    pop := []Gene{}
    result := []Person{}
    err := c.Find(bson.M{"group": group}).All(&result)
    if err != nil {
        fmt.Println(err)
    }
    for i:=0;i<len(result);i++ {
        split_str := strings.Split(result[i].Gene,",")
        chrom_tmp := []byte{}
        for _, str := range split_str {
            str_int, _ := strconv.Atoi(str)
            chrom_tmp = append(chrom_tmp,byte(str_int))
        }
        gene := Gene{chrom_tmp, 1}
        pop = append(pop, gene)
    }
    return pop
}

func DeletePopulation(c *mgo.Collection, group string) {
    num, err := c.Find(bson.M{"group": group}).Count()
    if err != nil{
        fmt.Println(err)
    }
    for i:=0;i<num;i++ {
        err = c.Remove(bson.M{"group": group})
        if err != nil {
            fmt.Println(err)
        }
    }
}

func InsertPopulation(c *mgo.Collection, group string, pop []Gene) {
    for i:=0;i<len(pop);i++{
        gene_str := ""
        for _,value := range pop[i].chrom {
            gene_str = gene_str + strconv.Itoa(int(value)) +","
        }
        gene_str = strings.TrimRight(gene_str,",")
        err := c.Insert(&Person{group, gene_str})
        if err != nil {
            fmt.Println(err)
        }
    }
}

func PrintPopulation(c *mgo.Collection, group string) {
    result := []Person{}
    err := c.Find(bson.M{"group": group}).All(&result)
    if err != nil {
        fmt.Println(err)
    }
    for i:=0;i<len(result);i++ {
        fmt.Println("id:",i," Group:",group," Gene:",result[i].Gene)
    }
}

func CreatePopulationDB(group string,size int, length int) {
    session, _ := mgo.Dial("mongodb://db:27017/test")
    c := session.DB("test").C("Gene")
    defer session.Close()
    pop := CreateGenes( size,length )
    DeletePopulation(c, group)
    InsertPopulation(c,group,pop)
}

func UpdatePopulationDB(group string,cycle int,ans []byte) {
    session, _ := mgo.Dial("mongodb://db:27017/test")
    c := session.DB("test").C("Gene")
    defer session.Close()
    pop := PicPopulation(c, group)
    pop = Cycle(pop, ans, cycle, 0.3, 0.8, 0.01)
    DeletePopulation(c, group)
    InsertPopulation(c,group,pop)
}

func PicGene( group string, pic int) string{
    session, _ := mgo.Dial("mongodb://db:27017/test")
    c := session.DB("test").C("Gene")
    defer session.Close()
    pop := PicPopulation(c, group)
    str := ""
    for j:=0;j<len(pop[pic].chrom);j++{
        str = str + fmt.Sprintf("%b",pop[pic].chrom[j])
    }
    return str
}
