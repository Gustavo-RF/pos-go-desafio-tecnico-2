package main

import (
	"flag"
	"fmt"
)

// docker run --name desafio-tecnico-2 desafio-tecnico-2 --foo=123 --blau=334
func main() {
	// Define a flag chamada 'foo' com um valor padrão 'default'
	foo := flag.String("foo", "default", "a foo flag")
	blau := flag.String("blau", "default blau", "a blau flag")
	flag.Parse() // Parseia os argumentos de linha de comando

	fmt.Println("FOO:", *foo)
	fmt.Println("BLAU:", *blau)

}
