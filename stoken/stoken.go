package stoken

import (
	"crypto/sha256"
	"encoding/base64"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Type string

const (
	Hash Type = "hash"
	Time Type = "time"
)

func Authorization(t Type, value string) string {
	switch t {
	case Hash:
		return HashAuthorization(value)
	case Time:
		return TimeAuthorization(value)
	default:
		return ""
	}
}

func HashAuthorization(token string) string {
	if token == "" {
		return ""
	}
	return "Token " + token
}

func TimeAuthorization(secret string) string {
	if secret == "" {
		return ""
	}
	sec := time.Now().Unix()
	head := strconv.FormatInt(sec, 16)
	sum := sha256.Sum256([]byte(head + "." + secret))
	hash := base64.RawStdEncoding.EncodeToString(sum[:])
	return "Token " + head + "." + hash
}

func Handler(t Type, value string) func(http.Handler) http.Handler {
	switch t {
	case Hash:
		return HashHandler(value)
	case Time:
		return TimeHandler(value)
	default:
		return nil
	}
}

func HashHandler(hash string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if hash == "" {
				next.ServeHTTP(w, r)
				return
			}
			token := getToken(r)
			if token == "" {
				unauthorized(w)
				return
			}
			sum := sha256.Sum256([]byte(token))
			val := base64.RawStdEncoding.EncodeToString(sum[:])
			if val != hash {
				unauthorized(w)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

func TimeHandler(secret string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if secret == "" {
				next.ServeHTTP(w, r)
				return
			}
			token := getToken(r)
			parts := strings.SplitN(token, ".", 2)
			if len(parts) != 2 {
				unauthorized(w)
				return
			}
			sec, err := strconv.ParseInt(parts[0], 16, 64)
			if err != nil {
				unauthorized(w)
				return
			}
			sec = time.Now().Unix() - sec
			if sec < -60 || sec > 60 {
				unauthorized(w)
				return
			}
			sum := sha256.Sum256([]byte(parts[0] + "." + secret))
			hash := base64.RawStdEncoding.EncodeToString(sum[:])
			if hash != parts[1] {
				unauthorized(w)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

func unauthorized(w http.ResponseWriter) {
	http.Error(w, "401 unauthorized", http.StatusUnauthorized)
}

func getToken(r *http.Request) string {
	if value := r.URL.Query().Get("token"); value != "" {
		return value
	}
	if values, ok := r.Header["Authorization"]; ok && len(values) > 0 {
		for _, value := range values {
			if strings.HasPrefix(value, "Token ") || strings.HasPrefix(value, "token ") {
				return strings.Split(value, " ")[1]
			}
		}
	}
	return ""
}
