package helper

func EnableAddons(addonName string) {
	switch addonName {
	case SA_FIREBASE_AUTH[0]:
		EnableFirebaseAuth()
	case SA_PUBSUB[0]:
		EnablePubSub()
	case SA_MONOGODB[0]:
		EnableMongoDB()
	case SA_REDIS[0]:
		EnableRedis()
	}
}
