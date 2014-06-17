package main

import (
    "fmt"
    "github.com/joliv/spark"
)

func main() {
    boring_data := []float64{1, 2, 3, 4.623, 5, 6, 8, 7}
    sparkline := spark.Line(boring_data)
    fmt.Println(sparkline)
}
