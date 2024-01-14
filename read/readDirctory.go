package read

import (
	"fmt"
	"os"
	"path"
)

var Files []string

func ReadDir(){
  homeDir:= os.ExpandEnv("$HOME")
  workingDir := path.Join(homeDir, ".ankiHome", os.Args[1])
  fmt.Println(workingDir)
  err := os.Walk(workingDir, func(path string, info os.FileInfo, err Error))
  
}
