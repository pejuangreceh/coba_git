package main

import (
	"coba_git/controller"
	"coba_git/model"
	"fmt"
)

func main() {
	controller.TampilkanController()
	model.Create()
	fmt.Println("Ini perubahan lagi")
	fmt.Println("Halo teman")
}
