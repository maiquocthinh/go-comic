package utils

import "net/url"

func ChangeDomain(originalURL *string, newDomain string) error {
	parsedURL, err := url.Parse(*originalURL)
	if err != nil {
		return err
	}

	parsedURL.Host = newDomain

	*originalURL = parsedURL.String()

	return nil
}
