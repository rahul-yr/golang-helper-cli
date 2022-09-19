package helper

func EnableAddons(addonName string) {
	switch addonName {
	case SA_FIREBASE_AUTH[0]:
		EnableFirebaseAuth()
	}
}
