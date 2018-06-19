package config

import (
	"fmt"
	"log"
	"os"
)

func init() {
	fmt.Println("config init...")
	log.SetOutput(os.Stdout)
}
