# Percorrendo arrays

> Em resumo, um array é uma estrutura de dados em que os elementos são armazenados em uma sequência contígua na memória. Cada elemento em um array é identificado por um índice inteiro, que começa em zero. Os arrays têm um tamanho fixo e são usados para armazenar coleções de valores do mesmo tipo, como números inteiros, pontos flutuantes ou caracteres. Em Go, os arrays são declarados usando a sintaxe [tamanho]tipo, onde tamanho é o número de elementos que o array pode armazenar e tipo é o tipo de cada elemento no array.

Em Go, para percorrer um array, podemos usar um loop `for`. Existem duas formas principais de percorrer um array em Go:

Usando o `for` com o índice:
Podemos usar o loop for com um índice para percorrer todos os elementos do array:

```go
arr := [5]int{1, 2, 3, 4, 5}

for i := 0; i < len(arr); i++ {
    fmt.Printf("O valor do elemento %d é %d\n", i, arr[i])
}
```

Nesse exemplo, usamos o loop `for` com uma variável de índice `i`, que é incrementada a cada iteração. Em cada iteração, imprimimos o índice e o valor do elemento correspondente no array.

Usando o `for` com o operador `range`:
Também podemos usar o operador range para percorrer um array em Go. Essa abordagem é mais concisa e mais segura do que usar um índice:

```go
arr := [5]int{1, 2, 3, 4, 5}

for i, v := range arr {
    fmt.Printf("O valor do elemento %d é %d\n", i, v)
}
```

Nesse exemplo, usamos o operador `range` para iterar sobre o array `arr`. A cada iteração, a variável `i` contém o índice atual, e a variável `v` contém o valor correspondente no array.

Essas são as duas formas ***mais comuns*** de percorrer um array em Go. Dependendo do caso de uso específico, uma abordagem pode ser mais adequada do que a outra.

## Tipos de for

### Loop for básico:
```go
for i := 0; i < 10; i++ {
    // corpo do loop
}
```

### Loop for sem condições:

```go
for {
    // corpo do loop
}
```

### Loop for com o operador range:

```go
for _, value := range someArray {
    // corpo do loop
}
``` 

### Loop for com continue e break:
```go
for i := 0; i < 10; i++ {
    if i%2 == 0 {
        continue // pula para a próxima iteração
    }
    if i == 7 {
        break // interrompe o loop
    }
    // corpo do loop
}
```

### Loop for com vários contadores:

```go
for i, j := 0, 10; i < 10 && j > 0; i, j = i+1, j-1 {
    // corpo do loop
}
```

### Loop for com range em um mapa:

```go
someMap := map[string]int{"a": 1, "b": 2, "c": 3}
for key, value := range someMap {
    // corpo do loop
}
```

Essas são as principais formas de trabalhar com o loop for em Go. Cada uma dessas abordagens pode ser útil em diferentes situações.