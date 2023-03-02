package main

import "fmt"

func buscaGitTag() {
	fmt.Println("Buscando commits pertencentes a git tag...")
}

var donoDoRepositorio string

func main() {
	fmt.Println("Hello World!")
	tag := "v1.26.2"
	donoDoRepositorio = "kubernetes"
	repo := "kubernetes"

	fmt.Printf("%s %s %s", tag, donoDoRepositorio, repo)
	buscaGitTag()
}
