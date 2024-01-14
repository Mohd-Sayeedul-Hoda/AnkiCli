package main

import (
  "fmt"
  "github.com/Mohd-Sayeedul-Hoda/AnkiCli/read"
)

func main(){
  fmt.Println("this is working")
  read.ReadDir()
  read.ReadFile()
  fmt.Println(read.CardsSlice)
  for _, val := range read.CardsSlice{
    fmt.Println("------ this front -----")
    fmt.Println(val.Front)
    fmt.Println("------ this back -----")
    fmt.Println(val.Back)
  }
  fmt.Println("Why this no working")
}
