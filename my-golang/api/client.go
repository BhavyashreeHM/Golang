package main
import (
    "bufio"
    "fmt"
    "log"
    "net/http"
    "os"
)

func main() {
    resp, err := http.Get("http://localhost:8080/stream")
    if err != nil {
        log.Fatalf("could not get stream: %v", err)
    }
    defer resp.Body.Close()

    scanner := bufio.NewScanner(resp.Body)
    for scanner.Scan() {
        line := scanner.Text()
        fmt.Fprintln(os.Stdout, line)
    }

    if err := scanner.Err(); err != nil {
        log.Fatalf("error reading stream: %v", err)
    }
}