package mypkg 

import "fmt"

func Malayalam() {
    greeting := "നമസ്കാരം" // "Namaskaram" in Malayalam

    fmt.Println("String:", greeting)
    fmt.Println("Length in bytes:", len(greeting))

    fmt.Println("\nCharacters (runes):")
    for i, r := range greeting {
        fmt.Printf("Index %d: %c (Unicode: U+%04X)\n", i, r, r)
    }
}
