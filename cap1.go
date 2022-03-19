package main
/*package main
É conhecido como declaração de pacote e todo programa Go deve começar com ela.

O package (pacote ou pacotes) são a maneira que o Go organiza e reutiliza código.

*/ 
import "fmt"

/*A palavra reservada "import"
É o modo de incluir código de outros pacotes para que sejam usados em nosso programa */

// O fmt é uma abreviatura de format (formatar) que implementa a formatação de entrada e saída.

func main(){ 
	fmt.Println("Hello, wolrd")
}

/* Agora vemos uma declaração de função
As funções são como blocos de construção de um programa Go.
*/
