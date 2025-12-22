package main


import "os"
import "cmd/envv/envv"
import "github.com/spf13/cobra"


func main() { 
	if err := envv.Execute(); err!=nil{
		// fmt.Println("program exited with error")
		os.Exit(1)
}
}

