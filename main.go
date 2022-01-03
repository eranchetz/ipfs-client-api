package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	shell "github.com/ipfs/go-ipfs-api"
)

func main() {
	// Where your local node is running on localhost:5001
	sh := shell.NewShell("localhost:5001")
	cid, err := sh.Add(strings.NewReader("hello world from go"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		os.Exit(1)
	}
	fmt.Printf("added %s\n", cid)
	fmt.Println("*** Catting content START ****")

	//         /ᐠ｡▿｡ᐟ\ﾉ *ᵖᵘʳʳ*
	r, _ := sh.Cat("/ipfs/" + cid)
	buf, err := io.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", buf)
	fmt.Println("*** Catting content END ****")
}
