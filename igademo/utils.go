package ga

import (
    "math/rand"
    "time"
)

//////// Utils /////////////////////////
func CopyPop(p []Gene) []Gene {
    tmpgenes := []Gene{}
    for _, gene := range p {
        tmpchrom := make([]byte,len(gene.chrom))
        copy(tmpchrom,gene.chrom)
        tmpgene := Gene{tmpchrom,gene.fit}
        tmpgenes = append(tmpgenes,tmpgene)
    }
    return tmpgenes
}

func CreateGenes(size int, length int) []Gene {
    rand.Seed(time.Now().UnixNano())
    genes := []Gene{}
    for i:=0;i<size;i++ {
        chrom := []byte{}
        for j:=0;j<length;j++ {
            chrom = append(chrom,byte(rand.Intn(255)))
        }
        gene := Gene{chrom,1}
        genes = append(genes,gene)
    }
    return genes
}

//////// Caluclation Fitness ///////////
func BitCount(buf byte) int{
    b:= (buf & 0x55)+((buf & 0xaa)>>1)
    b = (b & 0x33)+((b & 0xcc)>>2)
    b = (b & 0x0f)+((b & 0xf0)>>4)
    return int(b)
}

func MatchCount(buf1 byte, buf2 byte) int{
    b := buf1 ^ buf2
    return BitCount(^b)
}

func ArrayMatchCount(buf1 []byte, buf2 []byte) int {
    sum := 0
    for i:=0;i<len(buf1);i++ {
        sum = sum + MatchCount(buf1[i],buf2[i])
    }
    return sum
}

/////// Sort Algorithm /////////////////

func BubbleSort(p []Gene) {
    for i:=0;i<len(p);i++{
        for j:=len(p)-1;j>i;j--{
            if(p[j].fit < p[j-1].fit){
                tmpfit := p[j].fit
                tmpchrom := make([]byte,len(p[j].chrom))
                copy(tmpchrom,p[j].chrom)
                p[j].fit = p[j-1].fit
                copy(p[j].chrom, p[j-1].chrom)
                p[j-1].fit = tmpfit
                copy(p[j-1].chrom, tmpchrom)
            }
        }
    }
}
