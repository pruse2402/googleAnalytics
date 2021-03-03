package utils

import (
	"net/http"
	"strings"
)

//MinusIDs get array of IDs
func MinusIDs(from, to []int64) []int64 {
	res := []int64{}

	for _, a := range from {
		find := false
		for _, b := range to {
			if a == b {
				find = true
				break
			}
		}

		if !find {
			res = append(res, a)
		}
	}
	return res
}

// GetClientIP is an utility function for getting actual IP of end user
func GetClientIP(r *http.Request) string {
	// Header X-Forwarded-For
	hdrForwardedFor := http.CanonicalHeaderKey("X-Forwarded-For")
	if fwdFor := strings.TrimSpace(r.Header.Get(hdrForwardedFor)); fwdFor != "" {
		index := strings.Index(fwdFor, ",")
		if index == -1 {
			return fwdFor
		}
		return fwdFor[:index]
	}

	// Header X-Real-Ip
	hdrRealIP := http.CanonicalHeaderKey("X-Real-Ip")
	if realIP := strings.TrimSpace(r.Header.Get(hdrRealIP)); realIP != "" {
		return realIP
	}

	return "10.82.33.161"
}
