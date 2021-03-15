# Variables

**Como cualquier lenguaje de programación una variable almacena un valor a un nombre o identificador.**

```
nombre = 'valor'
```

---

## Nombres

A la hora de escribir un nombre para una variable se utiliza el sistema de camelCase y PascalCase.

| camelCase | PascalCase |
| --------- | ---------- |
| miNombre  | MiNombre   |

Cuando utilizamos camelCase, decimos que el acceso es solo para el scope de ese mismo package.
Sin embargo con PascalCase estamos declarandolo de forma global.

[Mas información...](https://golang.org/doc/effective_go#mixed-caps)

## Tipos

Golang es un lenguaje de tipado fuerte, lo cual a la hora a la hora de declarar una variable y definir su tipo no puede cambiar.

Entre ellos tenemos los más importantes como:

- Bool
- String
- Int
- Float32
- Float64

Entre otros para más información [click acá...](https://golang.org/pkg/go/types/)

## Escribiendo variables

Existe distintas formas de declarar, inicializar o ambas una variable. Para ello lo separe en 3 grupos.

- ### **Por lista**

```
var (
	mainCharacterGx string = "Judai"
	mainCharacter5ds string = "Yusei Fudo"
)
```

`En una linea`

```
var (
	mainCharacterGx, mainCharacter5ds string = "Judai", "Yusei Fudo"
	seasonGx, season5ds int = 2, 3
)
```

- ### **Declarados uno por uno**

```
var name string

var age int

var mainCharacterDuelMonsters = "Yugi Muto"
```

- ### **Asignación corta**

```
name := "Yusaku Fujiki"
age := 16
```

---

Si prestaste atención te diste cuenta que en algunas declaraciones se usa el tipo `string` o `int` y en otras no, además de la declaración de tipo `:=`, pero vamos por parte.

- Tipado

Cuando asignamos un tipo a una variable esperamos que el valor declarado sea el correspondido, caso contrario en tiempo de compilación arrojará un error.
Sin embargo cuando no lo declaramos, Go al momento de compilar asume que tipo le asignará dependiendo su tipo de valor.

- Asignación corta

Esta declaración se puede usar dentro de una función sin la especificación de usar `var` y un tipo.
