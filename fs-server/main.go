package main

import (
	"crypto/sha256"
	"crypto/subtle"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	if len(os.Args) != 3 {
		fmt.Println("Correct usage: fs-server <dir> <port>")
		os.Exit(1)
	}

	targetDir := os.Args[1]
	targetPort := os.Args[2]

	// authenticator := auth.NewBasicAuthenticator("myserver.com", getSecret)

	// http.HandleFunc("/", authenticator.Wrap(func(w http.ResponseWriter, ar *auth.AuthenticatedRequest) {
	// http.FileServer(http.Dir(targetDir)).ServeHTTP(w, &ar.Request)
	// }))
	http.HandleFunc("/", basicAuthSHA256())

	fmt.Printf("Running FS Server on port %s", targetPort)
	log.Fatal(http.ListenAndServe(":"+targetPort, nil))
}

func basicAuthSHA256(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()

		if ok {

			userHash := sha256.Sum256([]byte(username))
			passwordHash := sha256.Sum256([]byte(password))

			expectedUserHash := sha256.Sum256([]byte("gabrielgusn"))
			expectedPasswordHash := sha256.Sum256([]byte("super-strong-password"))

			userMatch := subtle.ConstantTimeCompare(userHash[:], expectedUserHash[:]) == 1
			passwordMatch := subtle.ConstantTimeCompare(passwordHash[:], expectedPasswordHash[:]) == 1

			if userMatch && passwordMatch {
				next.ServeHTTP(w, r)
				return
			}

			w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset=UTF-8`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
		}
	})
}

func protectedHandler(w http.ResponseWriter, r *http.Request) {
	return http.FileServer(http.Dir(targetDir))
}

func getSecret(user, realm string) string {
	if user == "gabrielgusn" {
		return fmt.Sprintf("%x", sha256.Sum256([]byte("super-secret-password")))
	}

	return ""
}
