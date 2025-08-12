# Saludos en GO

Este paquete proporciona una funcionalidad simplre para obtener saludos personalizados

## Instalación
Ejecuta el siguiente comando para instalar el paquete:

```bash
go get -u github.com/KevinAcSpin/greetings
````

## Uso

```go
package main

import (
  "fmt"
  "github.com/KevinAcSpin/greetings"
)

func main() {
  message, err := greetings.Hello("Kevin")
}
```
