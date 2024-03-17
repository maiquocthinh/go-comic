package uploadprovider

import "context"

type UploadProvider interface {
	UploadImage(ctx context.Context, data []byte, path string) error
	GetShareLink(ctx context.Context, path string) (string, error)
}
