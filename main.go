package main

import "chag/router"

func main() {
	_ = router.Router().Run(":80")
}
