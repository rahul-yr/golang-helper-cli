package helper

import "log"

func Help() {
	log.Println("Usage: rr [options]")
	log.Println("Options:")
	log.Panicln("Create a new golang project: rr --create-go <project-name> or rr -cgo <project-name>")
	log.Panicln("Enable addons: rr enable-addons <addon-name>")
	log.Panicln("Enable firebase-auth: rr enable-addons firebase-auth")
	log.Panicln("Enable pubsub: rr enable-addons pubsub")
	log.Panicln("Enable redis: rr enable-addons redis")
	log.Panicln("Enable mongodb: rr enable-addons mongodb")
}
