package ga

import (
    "math/rand"
    "time"
)

type Gene struct {
    chrom []byte
    fit   int
}

////// Loop GA ////////////
func Cycle(p []Gene,ans []byte, cycle int, elete_rate float32, cross_rate float32, mutate_rate float32) []Gene{
    rand.Seed(time.Now().UnixNano())
    p_len := len(p)
    elete_num := int(elete_rate*float32(len(p)))
    cross_num := int(cross_rate*float32(len(p)))
    for i:=0;i<cycle;i++ {
        //Calculate Fitness
        for j:=0;j<len(p);j++ {
            p[j].fit = ArrayMatchCount(ans, p[j].chrom)
        }
        elete := EleteSelection(p,elete_num)
        for j:=0;j<cross_num;j++ {
            ran1 := rand.Intn(len(elete))
            ran2 := rand.Intn(len(elete))
            next_chrom1, next_chrom2 := OnePointCrossover(elete[ran1].chrom,elete[ran2].chrom)
            p = append(p, Gene{next_chrom1,1})
            p = append(p, Gene{next_chrom2,1})
        }
        for j:=0;j<len(p);j++ {
            p[j].fit = ArrayMatchCount(ans, p[j].chrom)
        }
        for j:=0;j<len(p);j++ {
            ranmute := rand.Float32()
            if ranmute < mutate_rate {
                Mutation(p[j].chrom)
            }
        }
        p = EleteSelection(p,p_len)
    }
    return p
}
