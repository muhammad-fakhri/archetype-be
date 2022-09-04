package storage

import (
	"context"
	"io"
	"time"
)

type Storage interface {
	GenerateUploadURL(object, contentType string) (string, time.Time, error)
	GenerateDownloadURL(object string) (string, time.Time, error)
	DeleteObject(ctx context.Context, object string) error
	GetObjectSize(ctx context.Context, object string) (int64, error)
	ReadRange(ctx context.Context, object string, offset int64) ([]byte, error)
	GetObjectWriter(ctx context.Context, object string) (w io.Writer, closeFunc func())
	GetObjectReader(ctx context.Context, object string) (r io.Reader, closeFunc func(), err error)
}
