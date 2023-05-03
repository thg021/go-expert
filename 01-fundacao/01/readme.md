# Entendendo a primeira linha 

```go
package main
```
Essa linha declara que o arquivo faz parte do pacote principal "main". Em Go, o pacote "main" é especial, pois é o ponto de entrada para o programa.


```go
func main() {
```
Essa linha define a função "main", que é o ponto de entrada do programa. O código dentro dessa função será executado quando o programa for iniciado.

```go
println("Hello, world")
```
Essa linha usa a função "println" para imprimir a string "Hello, world" no console. A função "println" é uma função incorporada em Go que imprime seu argumento na saída padrão com uma nova linha adicionada ao final.

No geral, este código é um programa simples em Go que imprime "Hello, world" na saída do console. Ele é frequentemente usado como um exemplo básico para iniciantes em Go, já que demonstra a estrutura básica de um programa Go e como imprimir na tela.

***Todos os arquivos que estao na mesma pasta tem q ter necessariamente o mesmo nome de pacote***

```bash
go run main.go
```
