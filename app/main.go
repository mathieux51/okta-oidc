package main

import (
	"log"
	"net/http"
)

// Run next to envoy
// accept config with well-known endpoint
// check if bearer is valid against Okta
// returns Okta response if not good
// let request go through

// func middleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		// authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
// 		// if len(authHeader) != 2 {
// 		// 	fmt.Println("Malformed token")
// 		// 	w.WriteHeader(http.StatusUnauthorized)
// 		// 	w.Write([]byte("Malformed Token"))
// 		// } else {
// 		// 	jwtToken := authHeader[1]
// 		// 	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
// 		// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 		// 			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
// 		// 		}
// 		// 		return []byte(SECRETKEY), nil
// 		// 	})
//
// 		// 	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
// 		// 		ctx := context.WithValue(r.Context(), "props", claims)
// 		// 		// Access context values in handlers like this
// 		// 		// props, _ := r.Context().Value("props").(jwt.MapClaims)
// 		// 		next.ServeHTTP(w, r.WithContext(ctx))
// 		// 	} else {
// 		// 		fmt.Println(err)
// 		// 		w.WriteHeader(http.StatusUnauthorized)
// 		// 		w.Write([]byte("Unauthorized"))
// 		// 	}
// 		// }
// 	})
// }

// func handle(w http.ResponseWriter, r *http.Request) {
// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte("handle"))
// }

func main() {
	// http.Handle("/", middleware(http.HandlerFunc(handle)))
	log.Println("Hello test")
	log.Fatal(http.ListenAndServe(":10003", nil))
	// log.Println("Hello world")
}
