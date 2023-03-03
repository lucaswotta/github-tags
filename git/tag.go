package git

import "fmt"

type BuscadorGit struct {
}

func BuscaGitTag(tag, donoDoRepositorio, repo string) {
	fmt.Println("Buscando commits pertencentes a git tag...")
	fmt.Printf("%s %s %s\n", tag, donoDoRepositorio, repo)
}
