package main

import (
	"cmdLine/cmd"
	"log"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute err:%v", err)
	}
}

// func mian1() {
// 	var name string
// 	flag.StringVar(&name, "name", "changchaofeng", "your name")
// 	flag.Parse()
// 	fmt.Printf("name: %s \n", name)
// }

// func main2() {
// 	flag.Parse()
// 	var name string
// 	goCmd := flag.NewFlagSet("go", flag.ExitOnError)
// 	goCmd.StringVar(&name, "name", "changchaofeng", "your name")
// 	phpCmd := flag.NewFlagSet("php", flag.ExitOnError)
// 	phpCmd.StringVar(&name, "n", "changchaofeng", "your name")

// 	args := flag.Args()
// 	fmt.Println(args)
// 	switch args[0] {
// 	case "go":
// 		goCmd.Parse(args[1:])
// 	case "php":
// 		phpCmd.Parse(args[1:])
// 	}

// 	fmt.Printf("name: %s \n", name)
// }
