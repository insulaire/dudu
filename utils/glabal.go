package utils

import (
	"fmt"
)

var logo = `
'   ______     __    __      __      _    _____      __      _      _____   
'  (   __ \    ) )  ( (     /  \    / )  (_   _)    /  \    / )    / ___ \  
'   ) (__) )  ( (    ) )   / /\ \  / /     | |     / /\ \  / /    / /   \_) 
'  (    __/    ) )  ( (    ) ) ) ) ) )     | |     ) ) ) ) ) )   ( (  ____  
'   ) \ \  _  ( (    ) )  ( ( ( ( ( (      | |    ( ( ( ( ( (    ( ( (__  ) 
'  ( ( \ \_))  ) \__/ (   / /  \ \/ /     _| |__  / /  \ \/ /     \ \__/ /  
'   )_) \__/   \______/  (_/    \__/     /_____( (_/    \__/       \____/   
'                                                                           
`
var line_bottom = `-------------------------------RUNING-----------------------------------`
var line_top = `------------------------------------------------------------------------`

func PrintLogo() {
	fmt.Println(line_top)
	fmt.Println(logo)
	fmt.Println(line_bottom)
}

var GlabalObject glabalObject

type glabalObject struct {
	Name    string
	Host    string
	Port    int
	Version string
}

func init() {
	GlabalObject = glabalObject{
		Name:    "SimpleServer",
		Host:    "0.0.0.0",
		Port:    8888,
		Version: "v0.0.1",
	}
	// viper.AddConfigPath("/project/dudu/configs")
	// viper.SetConfigName("server")
	// viper.SetConfigType("yml")
	// err := viper.ReadInConfig()
	// if err != nil {
	// 	panic(err)
	// }
	// v := viper.Get("server")
	// viper.Unmarshal(&GlabalObject)

	PrintLogo()
}
