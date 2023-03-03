package main

import (
	"fmt"

	"github.com/lucaswotta/github-tags/git"
)

var donoDoRepositorio string

func main() {
	fmt.Println("Hello World!")
	tag := "v1.26.2"
	donoDoRepositorio := "kubernetes"
	repo := "kubernetes"

	go git.BuscaGitTag(tag, donoDoRepositorio, repo)

	go git.BuscaGitTag(tag, donoDoRepositorio, repo)
	fmt.Println("Sair do programa")
}
