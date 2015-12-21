package main

import (
    "encoding/hex"
    "fmt"
)

func main() {

    hexstring := fmt.Sprintf("%x", 12345678)
    fmt.Srintf(hexstring)
    hexbytes, _ := hex.DecodeString(hexstring)
    for {
        if len(hexbytes) >= 4 {
            break
        }
        hexbytes = append(hexbytes, 0)
    }
    fmt.Println(hexbytes)
}
