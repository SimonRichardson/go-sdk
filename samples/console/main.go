package main

import (
	"fmt"

	configcat "github.com/configcat/go-sdk/v4"
)

func main() {
	// Setting log level to Info to show detailed feature flag evaluation
	logger := configcat.DefaultLogger(configcat.LogLevelInfo)

	client := configcat.NewCustomClient("PKDVCLf-Hq-h-kCzMp-L7Q/HhOWfwVtZ0mb30i9wi17GQ", configcat.ClientConfig{Logger: logger})

	// create a user object to identify your user (optional)
	user := configcat.NewUserWithAdditionalAttributes("##SOME-USER-IDENTIFICATION##", "configcat@example.com", "", nil)

	// get individual config values identified by a key and a user object
	value := client.GetValueForUser("isPOCFeatureEnabled", false, user)

	fmt.Println("isPOCFeatureEnabled: ", value)
}
