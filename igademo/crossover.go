package ga

import (
    "math/rand"
    "time"
)

/////// Crossover    /////////////////

func OnePointCrossover(chrom1 []byte, chrom2 []byte) ([]byte, []byte){
    rand.Seed(time.Now().UnixNano())
    ran := rand.Intn(len(chrom1)*8)
    locate := int(float32(ran)/8)
    point := ran%8
    mask := []byte{0x00,0x80,0xc0,0xe0,0xf0,0xf8,0xfc,0xfe}
    next_chrom1 := make([]byte,len(chrom1))
    next_chrom2 := make([]byte,len(chrom2))
    for i:=0;i<locate;i++ {
        next_chrom1[i] = chrom2[i]
        next_chrom2[i] = chrom1[i]
    }
    next_chrom2[locate] = (chrom1[locate] & mask[point]) + (chrom2[locate] &(^mask[point]))
    next_chrom1[locate] = (chrom2[locate] & mask[point]) + (chrom1[locate] &(^mask[point]))
    for i:=locate+1;i<len(chrom1);i++ {
        next_chrom1[i] = chrom1[i]
        next_chrom2[i] = chrom2[i]
    }
    return next_chrom1, next_chrom2
}
