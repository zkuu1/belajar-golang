package main
import ("fmt")

func loopExample() {
  for i := 0; i < 5; i++ {
    if i == 3 {
      continue
    }
   fmt.Println(i)
  }
}