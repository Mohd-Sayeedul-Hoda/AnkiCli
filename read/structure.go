package read

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
)

type CardStruct struct{
  Front string
  Back string
  DoubleSided bool
}

type CardsSlice []CardStruct

func ReadFile(){
  for _, file := range FileSlice{
    card := new(CardStruct)
  }
}

func makeStruct(fileString string){

  file, err := os.Open(fileString)
  if err != nil{
      log.Fatal("Cannot open the file")
  } 
  input := bufio.NewScanner(file)

  // define flag for reading input 
  flagIn := false
  var card CardStruct
  structText := ""

  //define flag for completing struct
  flagSt := false


  for input.Scan(){

    text := input.Text()
    if text == ""{
      continue
    }

    text = strings.TrimSpace(text)

    if text[0] == byte('-'){
      flagSt = !flagSt
      continue
    }

    if flagSt{
      if text[0] == byte('*')
    }

    
  }
}

