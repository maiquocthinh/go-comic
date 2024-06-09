package utils

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"io"
	"net/url"
	"strconv"
)

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func ChangeDomain(originalURL *string, newDomain string) error {
	parsedURL, err := url.Parse(*originalURL)
	if err != nil {
		return err
	}

	parsedURL.Host = newDomain

	*originalURL = parsedURL.String()

	return nil
}

var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

func GenerateOTPCode(max int) string {
	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}

func SecondsToMinutes(seconds int) string {
	minutes := seconds / 60
	remainderSeconds := seconds % 60

	result := strconv.Itoa(minutes) + " minutes"
	if remainderSeconds > 0 {
		result += " " + strconv.Itoa(remainderSeconds) + " seconds"
	}

	return result
}
