package main  

import(
  "fmt"
  "math"
)

func sqrt(x float 64) string{
  if x < 0 {
    return sqrt(-x) + "i"
  }
  return fmt.Sprint(math.sqrt(x))
}

func main() { 
  fmt.Println(sqrt(2), sqrt(-4))
}
