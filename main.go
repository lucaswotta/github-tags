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

	b := git.BuscadorGit{
		DonoDoRepositorio: donoDoRepositorio,
		Repo:              repo,
	}
	b.BuscaGitTag(tag)

	fmt.Println("Sair do programa")
}
