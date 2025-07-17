// cmd/blockchainexplorersdk/main.go
package main

import (
"flag"
"log"
"os"

"blockchainexplorersdk/internal/blockchainexplorersdk"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := blockchainexplorersdk.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
