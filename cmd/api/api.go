package main

import (
	"fmt"
	"log"
	"os"

	"github.com/FelipeMandelli/ProjetoInterativoV/cmd/api/configs"
	"github.com/spf13/viper"
)

func main() {
	fmt.Println("This is the api appliction!")

	viper.SetConfigFile(configs.BasePath + string(os.PathSeparator) +"configs/configurations.yml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("error reading configs: ", err)
	}

	fmt.Println(viper.GetString("HOST"))
}
