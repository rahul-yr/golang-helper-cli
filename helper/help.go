package helper

import "log"

func Help() {
	log.Println("Usage: rr [options]")
	log.Println("Options:")
	log.Println("Create a new golang project: rr --create-go <project-name> or rr -cgo <project-name>")
	log.Println("Enable addons: rr enable-addons <addon-name>")
	log.Println("Enable firebase-auth: rr enable-addons firebase-auth")
	log.Println("Enable pubsub: rr enable-addons pubsub")
	log.Println("Enable redis: rr enable-addons redis")
	log.Println("Enable mongodb: rr enable-addons mongodb")
}
