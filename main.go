package main

import (
	"./server"
)

// TODO: Document
// TODO: Release
// TODO: Add to Travis
// TODO: Reorder Dockerfile instructions
// TODO: Integration tests
func main() {
	server.New().Execute()
}
