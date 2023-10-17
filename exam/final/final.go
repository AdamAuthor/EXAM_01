package main

import (
    "fmt"
    "github.com/qeesung/asciiplayer"
)

func main() {
    // Create a new asciiplayer instance.
    player := asciiplayer.New()

    // Set the input file.
    err := player.SetInput("ascii.mp4")
    if err != nil {
        fmt.Println(err)
        return
    }

    // Set the output file.
    err = player.SetOutput("output.txt")
    if err != nil {
        fmt.Println(err)
        return
    }

    // Play the ASCII video.
    err = player.Play()
    if err != nil {
        fmt.Println(err)
        return
    }
}
