package http_handler

import (
	"encoding/base64"
	"net/http"
)

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func redirectPolicyFunc(req *http.Request, via []*http.Request) error {
	//	creds := configure.GetCredentials()
	//	req.Header.Add("Authorization", "Basic "+basicAuth(creds.User, creds.Password))
	return nil
}
