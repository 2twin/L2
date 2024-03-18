Что выведет программа? Объяснить вывод программы. Объяснить как работают defer’ы и их порядок вызовов.

```go
package main

import (
	"fmt"
)


func test() (x int) {
	defer func() {
		x++
	}()
	x = 1
	return
}


func anotherTest() int {
	var x int
	defer func() {
		x++
	}()
	x = 1
	return x
}


func main() {
	fmt.Println(test())
	fmt.Println(anotherTest())
}
```

Ответ:
```
2
1

```
В первом случае манипуляции производятся с именованной переменной `x`, возвращаемой функцией `test`, в этом случае `defer` увелечит значение `x` на 1 перед `return`, в результате `x` окажется равен 2.

Во втором случае конструкция `defer` никак не повлияет на `x`, в итоге `x` будет хранить значение 1.