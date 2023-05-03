# Importando o fmt e tipagem

Em Go, para verificar o tipo de uma **variável**, você pode usar o operador `printf` ou `println` junto com o `%T`.

Por exemplo, para imprimir o tipo da variável `x`, você pode usar o seguinte código:

```go
x := 42
fmt.Printf("O tipo de x é %T\n", x)
```

O resultado da execução será: `O tipo de x é int`

Dessa forma, `%T` é um especificador de formato para imprimir o tipo da variável. Quando executado, ele exibe o tipo da **variável x**, que neste caso é int.

Você pode usar essa abordagem para imprimir o tipo de qualquer variável em Go. Basta substituir x pelo nome da variável que deseja verificar.

Aqui estão os principais especificadores de formato:

    - **%v** - representa o valor da variável na sua forma padrão
    - **%+v** - representa o valor da variável, incluindo os campos nomeados da struct
    - **%#v** - representa a representação de valor da variável, em um formato literal Go
    - **%T** - representa o tipo da variável
    - **%t** - representa o valor booleano (true ou false)
    - **%d** - representa um valor inteiro decimal
    - **%b** - representa um valor inteiro em binário
    - **%o** - representa um valor inteiro em octal
    - **%x** - representa um valor inteiro em hexadecimal (com letras minúsculas)
    - **%X** - representa um valor inteiro em hexadecimal (com letras maiúsculas)
    - **%f** - representa um valor de ponto flutuante
    - **%e** - representa um valor de ponto flutuante em notação científica (por exemplo, 1.23e+08)
    - **%E** - representa um valor de ponto flutuante em notação científica (por exemplo, 1.23E+08)
    - **%s** - representa uma string de caracteres
    - **%q** - representa uma string de caracteres, citando-a entre aspas duplas
    - **%p** - representa um ponteiro
Esses são os principais especificadores de formato em Go, mas há mais opções disponíveis. Você pode encontrar mais informações sobre esses especificadores de formato na documentação oficial da linguagem Go.