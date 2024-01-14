// this is file to get all the file from the respected dir
package read

import (
	"log"
	"os"
	"path"
)

var FileSlice []string

func ReadDir(){
  workingDir := path.Join(os.ExpandEnv("$HOME"), ".ankiHome", os.Args[1])
  files, err := os.ReadDir(workingDir)
  if err != nil{
      log.Fatal(err)
  } 

  for _, f := range files{
    FileSlice = append(FileSlice, path.Join(workingDir, f.Name()))
  }

}
