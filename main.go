package main

import (
    "os"
    "strconv"
    "fmt"
    "bytes"
    "crypto/sha1"
)

func main () {
    if len(os.Args) != 2 {
        println("it must take args")
        os.Exit(1)
    }

    num_hash, err := strconv.Atoi(os.Args[1])
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    alph := "abcdefghijklmnopqrstuvwxyz"
    var strs [100]int

    for i := 0; i < num_hash; i++ {
        up(&strs, alph)
        var buffer bytes.Buffer
        for _, value := range strs {
            buffer.WriteByte(alph[value])
        }

        fmt.Printf("%x\n", sha1.Sum(buffer.Bytes()))
    }
}

func up(strs *[100]int, alph string) {
    for i := 0; i < len(strs); i++ {
        if strs[i] < len(alph) - 1 {
            strs[i]++
            return
        }
        strs[i] = 0
    }
}
