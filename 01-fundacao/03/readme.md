# Criação de tipos

Em Go, é possível criar um tipo personalizado usando a palavra-chave `type`. Essa funcionalidade permite definir um novo tipo de dados com base em um tipo existente, ou até mesmo sem nenhum tipo subjacente.

Aqui está um exemplo de como criar um tipo personalizado em Go com base em um tipo existente:

```go
type Celsius float64
```

Nesse exemplo, criamos um novo tipo chamado `Celsius`, que é baseado no tipo `float64`. Com isso, podemos criar variáveis do tipo `Celsius` e usá-las em nosso código:

```go
var temperature Celsius = 25.0
fmt.Printf("The temperature is %g degrees Celsius.\n", temperature)
```
Também é possível definir um tipo personalizado sem um tipo subjacente. Por exemplo, vamos criar um novo tipo de dados chamado Point, que representa um ponto no espaço bidimensional:


```go
type Point struct {
    X, Y float64
}
```

Nesse exemplo, criamos um novo tipo Point que possui dois campos `X` e `Y` do tipo `float64`. Com isso, podemos criar variáveis do tipo `Point` e usá-las em nosso código:

```go
var p Point
p.X = 3.0
p.Y = 4.0
fmt.Printf("The point is (%g, %g)\n", p.X, p.Y)
```

Esses são apenas alguns exemplos de como criar tipos personalizados em Go. Com a funcionalidade type, é possível criar tipos mais complexos e estruturados para representar dados específicos em seu código.