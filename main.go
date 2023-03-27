package main

import "belajar-swagger/routers"

const PORT = ":3000"

func main() {
	routers.Router().Run(PORT)
}
