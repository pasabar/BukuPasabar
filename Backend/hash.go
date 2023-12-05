package pasabarbackend

import "golang.org/x/crypto/bcrypt"

func HashPass(password string) (string, error) {
	bytess, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(bytess), err
}

func CompareHashPass(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

$ git push
Total 0 (delta 0), reused 0 (delta 0), pack-reused 0
To github.com:pasabar/pasabar-backend.git
21b7e22..abs1326 main -> main