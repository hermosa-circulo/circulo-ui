package ga

import (
    "math/rand"
    "time"
)

/////// Mutation    /////////////////

func Mutation(chrom []byte) []byte{
    rand.Seed(time.Now().UnixNano())
    onebit := []byte{ 0x80, 0x40, 0x20, 0x10, 0x08, 0x04, 0x02, 0x01 }
    ran1 := rand.Intn(len(chrom))
    ran2 := rand.Intn(8)
    chrom[ran1] ^= onebit[ran2]
    return chrom
}
