package utils

// import (
// 	"fmt"
// 	"net/http"
// 	"os"

// 	users "github.com/branislavstojkovic70/nft-ticket-verification/domain/users"
// 	"github.com/golang-jwt/jwt/v4"
// )

// func createJWT(user *users.User, jwt_secret string) (string, error) {
// 	claims := &jwt.MapClaims{
// 		"expiresAt":     15000,
// 		"accountNumber": user.ID,
// 	}
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	return token.SignedString([]byte(jwt_secret))
// }

// func permissionDenied(w http.ResponseWriter) {
// 	WriteJSON(w, http.StatusForbidden, "permission denied")
// }

// func withJWTAuth(handlerFunc http.HandlerFunc, user *users.User) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Println("calling JWT auth middleware")

// 		tokenString := r.Header.Get("x-jwt-token")
// 		token, err := validateJWT(tokenString)
// 		if err != nil {
// 			permissionDenied(w)
// 			return
// 		}
// 		if !token.Valid {
// 			permissionDenied(w)
// 			return
// 		}
// 		userID, err := getID(r)
// 		if err != nil {
// 			permissionDenied(w)
// 			return
// 		}
// 		account, err := s.GetAccountByID(userID)
// 		if err != nil {
// 			permissionDenied(w)
// 			return
// 		}

// 		claims := token.Claims.(jwt.MapClaims)
// 		if account.Number != int64(claims["accountNumber"].(float64)) {
// 			permissionDenied(w)
// 			return
// 		}

// 		if err != nil {
// 			WriteJSON(w, http.StatusForbidden, ApiError{Error: "invalid token"})
// 			return
// 		}

// 		handlerFunc(w, r)
// 	}
// }

// func validateJWT(tokenString string) (*jwt.Token, error) {
// 	secret := os.Getenv("JWT_SECRET")

// 	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		// Don't forget to validate the alg is what you expect:
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
// 		}

// 		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
// 		return []byte(secret), nil
// 	})
// }
