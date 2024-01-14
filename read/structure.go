package read

import (
	"bufio"
	//"fmt"
	"log"
	"os"
	"strings"
)

type CardStruct struct{
  Front string
  Back string
  //DoubleSided bool
}

var CardsSlice []*CardStruct

func ReadFile(){
  CardsSlice = make([]*CardStruct, 0)
  for _, file := range FileSlice{
    makeStruct(file)
  }
}

func makeStruct(fileString string){

  file, err := os.Open(fileString)
  if err != nil{
      log.Fatal(err)
  } 
  input := bufio.NewScanner(file)

  // define flag for reading input 
  flagIn := false
  var card *CardStruct

  //define flag for completing struct
  flagSt := false


  for input.Scan(){

    text := input.Text()
    text = strings.TrimSpace(text)
    //fmt.Println(text)

    if text == ""{
      continue
    }


    if text[0] == byte('-'){
      flagSt = !flagSt
      if flagSt{
	card = new(CardStruct)
      }else{
	CardsSlice = append(CardsSlice, card)
	card = nil
      }
      continue
    }

    if flagSt{
      if text[0] == byte('*'){
	flagIn = !flagIn
	text = strings.TrimSpace(text[1:])
	if flagIn{
	  card.Front =text
	}else{
	  card.Back = text
	}
      }
    }

    
  }
}

