package main

import "github.com/branislavstojkovic70/nft-ticket-verification/bootstrap"

func main() {
	_ = bootstrap.Run()

	select {}
}
