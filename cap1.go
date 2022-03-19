package main
/*package main
É conhecido como declaração de pacote e todo programa Go deve começar com ela.

O package (pacote ou pacotes) são a maneira que o Go organiza e reutiliza código.

*/ 
import "fmt"

/*A palavra reservada "import"
É o modo de incluir código de outros pacotes para que sejam usados em nosso programa */

				/* O fmt é uma abreviatura de format (formatar) que implementa a formatação de entrada e saída. */

func main(){ 
	fmt.Println("Hello, wolrd")
}

/* Agora vemos uma declaração de função
As funções são como blocos de construção de um programa Go.
*/

// Println quer dizer "print line".

/* Exercício do Capitulo 1

1º	Oque é um espaço em branco ?
	-O espaço em branco é ignorado pelo compilador

2º O que é um comentário ? Quais as duas maneira de escrever um comentario ?
	Comentario serve para explicar o código e também é ignorado pelo compilador
	As formas de escrever um comentário são // e /**/

/**/