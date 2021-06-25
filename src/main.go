package main

import (
	"fmt"
	"mygin/src/routers"
)

func main() {
	r := routers.SetupRouter()
	if err := r.Run("localhost:8080"); err != nil {
		fmt.Println("startup service failed, err:%v\n", err)
	}



}

