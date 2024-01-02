package main

import (
	api "rw-file/api"
)

func main() {
	api.HandleJsonW()
	api.HandleJsonR()

	api.HandleXmlW()
	api.HandleXmlR()

	api.HandleGobW()
	api.HandleGobR()

	api.HandleTextW()
	api.HandleTextR()
}
