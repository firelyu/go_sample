package main

import (
    "bufio"
    "fmt"
    "github.com/firelyu/go_sample/encode"
    "io"
    "os"
)

func usage() {
    fmt.Printf("    %s [<file>]\n", os.Args[0])
    fmt.Printf("        - Decode the file or stdin with base64.\n", os.Args[0])
    os.Exit(1)
}

func main() {
    buf := make([]byte, 4<<20)
    var r *bufio.Reader

    switch len(os.Args) {
    case 1:
        // read from pipe
        r = bufio.NewReader(os.Stdin)
    case 2:
        // read from file
        f, err := os.Open(os.Args[1])
        if err != nil {
            fmt.Println(err)
        }

        r = bufio.NewReader(f)
    default:
        usage()
    }

    n, err := r.Read(buf)
    if err != nil && err != io.EOF {
        fmt.Println(err)
    }

    p, err := encode.DecodeBase64(buf[:n-1])
    if err != nil {
        fmt.Println(err)
    }

    fmt.Println(string(p))
}
