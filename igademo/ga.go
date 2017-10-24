package ga

import (
    "fmt"
    "math/rand"
    "time"
    "strconv"

    mgo "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
)


type Person struct {
        Group string
        Gene string
}

type Gene struct {
    dna []int
    fitness float64
}

type Population struct {
    genes []Gene
}

func (g *Gene) CalcFitness(ans [][]int){
    g.fitness = 0
    for i:=0;i<len(ans);i++{
        for j:=0;j<len(g.dna);j++{
            if ans[i][j] == g.dna[j] {
               g.fitness += 1
            }
        }
    }
}

func (p *Population) CalcFitness(ans [][]int) {
    for i:=0; i< len(p.genes); i++ {
        p.genes[i].CalcFitness(ans)
    }
}

func Addgene(g []Gene, g_add Gene) []Gene{
    g = append(g , g_add)
    return g
}

func CreatePopulation( size int, leng int ) Population {
    population := Population{}
    for i:=0; i< size ; i++ {
        var dna_tmp []int
        for j := 0; j < leng ; j++ {
            rand.Seed(time.Now().UnixNano())
            ran := rand.Intn(2)
            dna_tmp = append(dna_tmp, ran)
        }
        gene := Gene{dna_tmp, 1}
        population.genes = Addgene(population.genes, gene)
    }
    return population
}

func (p *Population) SelectRoulette() {
    rand.Seed(time.Now().UnixNano())
    fitnesssum :=  0.0
    for i:=0;i<len(p.genes);i++{
        fitnesssum = fitnesssum + p.genes[i].fitness
    }
    temp_genes := []Gene{}
    for i:=0;i<len(p.genes);i++{
        ran := rand.Float64()*fitnesssum
        sum := 0.0
        for j:=0;j<len(p.genes);j++{
            sum = sum + p.genes[j].fitness
            if ran <= sum {
                temp_genes = Addgene(temp_genes, p.genes[i])
                break
            }
        }
    }
    for i:=0;i<len(p.genes);i++{
        p.genes[i] = temp_genes[i]
    }
}

func (p *Population) Shufflegenes(){
    n :=  len(p.genes)
    rand.Seed(time.Now().UnixNano())
    for i:=n -1; i>=0;i-- {
        ran := rand.Intn(i+1)
        temp := []int{}
        tempfit := p.genes[i].fitness
        for j:=0;j<len(p.genes[i].dna);j++ {
            temp = append(temp,p.genes[i].dna[j])
        }
        for j:=0;j<len(p.genes[i].dna);j++ {
            p.genes[i].dna[j] = p.genes[ran].dna[j]
            p.genes[ran].dna[j] = temp[j]
        }
        p.genes[i].fitness = p.genes[ran].fitness
        p.genes[ran].fitness = tempfit
    }
}

func (p *Population) Maxfitness() float64{
    max := 0.0
    for i:=0;i<len(p.genes);i++ {
        if max < p.genes[i].fitness {
           max = p.genes[i].fitness
        }
    }
    return max
}

func (p *Population) Print(){
    for i := 0; i < len(p.genes); i++{
        fmt.Println(p.genes[i])
    }
}

func (p *Population) Crossover(g1 int, g2 int) []Gene{
    next_genes := []Gene{}
    c1 := make([]int,len(p.genes[g1].dna))
    c2 := make([]int,len(p.genes[g2].dna))
    c_temp := make([]int,len(c1))
    copy(c1,p.genes[g1].dna)
    copy(c2,p.genes[g2].dna)
    copy(c_temp,c1)
    temp := []int{}
    rand.Seed(time.Now().UnixNano())
    crosspoint := rand.Intn(len(p.genes[g1].dna))
    temp = append(c1[:crosspoint],c2[crosspoint:]...)
    temp_gene := Gene{temp, 0}
    next_genes = append(next_genes, temp_gene)
    temp = append(c2[:crosspoint],c_temp[crosspoint:]...)
    temp_gene = Gene{temp, 0}
    next_genes = append(next_genes, temp_gene)
    return next_genes
}

func (p *Population) Mutation(per float64){
    rand.Seed(time.Now().UnixNano())
    for i := 0; i < len(p.genes); i++{
        for j:=0;j<len(p.genes[i].dna);j++{
            ran := rand.Float64()
            if ran < per {
                if p.genes[i].dna[j] == 0 {
                    p.genes[i].dna[j] = 1
                } else {
                    p.genes[i].dna[j] = 0
                }
            }
        }
    }
}

func Sort(genes []Gene) {
    for i:=0; i< len(genes)-1;i++ {
        for j:=len(genes)-1;j>i;j-- {
            if genes[j-1].fitness > genes[j].fitness {
                temp_f := genes[j-1].fitness
                temp_dna := make([]int,len(genes[j-1].dna))
                copy(temp_dna,genes[j-1].dna)
                genes[j-1].fitness = genes[j].fitness
                for k:=0;k<len(genes[j-1].dna);k++ {
                    genes[j-1].dna[k] = genes[j].dna[k]
                }
                genes[j].fitness = temp_f
                for k:=0;k<len(genes[j-1].dna);k++ {
                    genes[j].dna[k] = temp_dna[k]
                }
            }
        }
    }
}

func (p *Population) Select( elite_length int , elite_flag int) Population{
    temp_genes := []Gene{}
    for i:=0;i<len(p.genes);i++{
        temp_f := p.genes[i].fitness
        temp_dna := make([]int, len(p.genes[i].dna))
        copy(temp_dna, p.genes[i].dna)
        temp_gene := Gene{temp_dna, temp_f}
        temp_genes = Addgene(temp_genes,temp_gene)
    }
    Sort(temp_genes)
    var result []Gene
    if elite_flag == 1{
        result = temp_genes[len(p.genes)-elite_length:]
    } else {
        result = temp_genes[:elite_length]
    }
    result_population := Population{result}
    return result_population
}

func (p Population) Cycle(ans [][]int, cycle int, elite_per float64,cross_per float64, mutate_per float64) Population{
    length := len(p.genes)
    rand.Seed(time.Now().UnixNano())
    elite_num := int(float64(length)*elite_per)
    cross_num := int(float64(elite_num)*cross_per)
    for i:=0;i<cycle;i++{
        elite := p.Select(elite_num,1)
        for j:=0; j< cross_num; j++ {
           ran1 := rand.Intn(len(elite.genes))
           ran2 := rand.Intn(len(elite.genes))
           addgenes := elite.Crossover(ran1,ran2)
           p.genes = Addgene(p.genes,addgenes[0])
           p.genes = Addgene(p.genes,addgenes[1])
        }
        p.CalcFitness(ans)
        p.Mutation(mutate_per)
        p = p.Select(length,1)
    }
    return p
}

func ArrayString(array []int) string{
    str := ""
    for i:=0;i<len(array);i++{
        str = str + strconv.Itoa(array[i])
    }
    return str
}

func StringArray(str string) []int{
    array := []int{}
    for _, c := range str {
        int_c, _ := strconv.Atoi(string(c))
        array = append(array,int_c)
    }
    return array
}

func PicPopulation(c *mgo.Collection, group string) Population{
    population := Population{}
    result := []Person{}
    err := c.Find(bson.M{"group": group}).All(&result)
    if err != nil {
        fmt.Println(err)
    }
    for i:=0;i<len(result);i++ {
        dna_tmp := StringArray(result[i].Gene)
        gene := Gene{dna_tmp, 1}
        population.genes = Addgene(population.genes, gene)
    }
    return population
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

func (p *Population)InsertPopulation(c *mgo.Collection, group string) {
    for i:=0;i<len(p.genes);i++{
        gene_str := ArrayString(p.genes[i].dna)
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
    DeletePopulation(c, group)
    population := CreatePopulation( size,length )
    population.InsertPopulation(c,group)
}

func UpdatePopulationDB(group string,cycle int,ans [][]int) {
    session, _ := mgo.Dial("mongodb://db:27017/test")
    c := session.DB("test").C("Gene")
    defer session.Close()
    population := PicPopulation(c, group)
    population = population.Cycle(ans, cycle,0.3,0.9,0.01)
    DeletePopulation(c, group)
    population.InsertPopulation(c,group)
}

func PicGene(group string, pic []int) string{
    session, _ := mgo.Dial("mongodb://db:27017/test")
    c := session.DB("test").C("Gene")
    defer session.Close()
    population := PicPopulation(c, group)
    elite := population.Select(len(population.genes),1)
    return_str :="{"
    for _, num := range pic {
        return_str = return_str+ "{"+ArrayString(elite.genes[len(population.genes)-num].dna)+"},"
    }
    return_str = return_str + "}"
    return return_str
}


/*func main(){
    ans := [][]int{{}}
    for i:=0;i<100;i++ {
        ans[0] = append(ans[0],1)
    }
    group := "test"
    CreatePopulationDB(group,80,100)
    UpdatePopulationDB(group,10,ans)
}*/
