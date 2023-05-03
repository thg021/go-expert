# Declaração e Atribuicação

**Declaração e atribuição separadas**: Nessa forma, a variável é declarada primeiro, indicando o tipo de dados, seguido do operador de atribuição = e o valor atribuído à variável. Por exemplo:

```go
var a int
a = 10
```

**Declaração e atribuição combinadas**: Nessa forma, a variável é declarada e atribuída em uma única linha, sem a necessidade de usar o operador de atribuição = separadamente. Por exemplo:

```go
var a int = 10
```
Também é possível omitir o tipo da variável, deixando que o Go deduza o tipo do valor atribuído:

```go
var a = 10
```
***Declaração curta de variáveis***: Essa forma é usada para declarar e atribuir variáveis em uma única linha, sem a necessidade de usar a palavra-chave `var`. Por exemplo:

```go
a := 10
```
Nessa forma, o tipo da variável é deduzido automaticamente pelo Go a partir do valor atribuído. Além disso, é possível declarar e atribuir várias variáveis em uma única linha:

```go
a, b := 10, "hello"
```
Essas são as três formas mais comuns de declarar e atribuir variáveis em Go. Escolha a forma que melhor se adapta ao seu código e estilo de programação.