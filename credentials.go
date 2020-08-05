package vcodeHMAC

import (
	"errors"
	"os"
)

func getCredentials(veracodeAPIKeyID, veracodeAPIKeySecret string) ([2]string, error) {
	var credentials [2]string

	// Read Veracode API Key ID from arg, if not provided as an arg
	if veracodeAPIKeyID == "" {
		veracodeAPIKeyID = os.Getenv("VERACODE_API_KEY_ID")
	}

	// Read Veracode API Key Secret from arg, if not provided as an arg
	if veracodeAPIKeySecret == "" {
		veracodeAPIKeySecret = os.Getenv("VERACODE_API_KEY_SECRET")
	}

	// If still empty, then generate an error for missing credentials
	var err error
	if veracodeAPIKeyID == "" {
		err = errors.New("No veracodeAPIKeyID found in env var")
	}
	if veracodeAPIKeySecret == "" {
		err = errors.New("No veracodeAPIKeySecret found in env var")
	}

	// Setup the credentials arr
	credentials[0] = veracodeAPIKeyID
	credentials[1] = veracodeAPIKeySecret

	return credentials, err
}
