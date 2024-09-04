package main

import (
    "flag"
    "fmt"
)

func main() {
    // Define flags
    name := flag.String("name", "World", "a name to say hello to")
    age := flag.Int("age", 0, "your age")
    isAdmin := flag.Bool("admin", false, "are you an admin?")

    // Parse the flags
    flag.Parse()

    // Use the flag values
    fmt.Printf("Hello, %s!\n", *name)
    fmt.Printf("Age: %d\n", *age)
    if *isAdmin {
        fmt.Println("You are an admin.")
    } else {
        fmt.Println("You are not an admin.")
    }
}
