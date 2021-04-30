/**
 * @EditTime: 2021-04-30
 * @Developer: zstk
 */

package main

import (
	"lib_new_website_server/app/core"
	"lib_new_website_server/app/utils"
	"log"
)

func main() {
	engine := core.GetDefaultEngine()
	port := utils.GetConfig("server", "addr")
	if err := engine.Run(port); err != nil {
		log.Fatal("Run Error")
	}
}
