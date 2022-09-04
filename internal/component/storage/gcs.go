package storage

// BEGIN __INCLUDE_GCS__
import (
	"context"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"cloud.google.com/go/storage"
	"github.com/muhammad-fakhri/archetype-be/internal/config"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
)

type gcsClient struct {
	bucket *storage.BucketHandle
	config *gcsConfig
}

type gcsConfig struct {
	AccessID              string
	PrivateKey            []byte
	DownloadURLExpiryTime time.Duration
	UploadURLExpiryTime   time.Duration
	ReadBatchSize         int64
}

func InitGCSClient(conf *config.GCSConfig) (c Storage) {
	var err error
	defer func() {
		if err != nil {
			log.Fatalf("failed init gcs client. err:%v", err)
		}
	}()

	credentialsJSON, err := ioutil.ReadFile(conf.CredentialFileName)
	if err != nil {
		return nil
	}

	jwtConf, err := google.JWTConfigFromJSON(credentialsJSON)
	if err != nil {
		return nil
	}

	storageClient, err := storage.NewClient(context.Background(), option.WithCredentialsJSON(credentialsJSON))
	if err != nil {
		return nil
	}

	return &gcsClient{
		bucket: storageClient.Bucket(conf.BucketName),
		config: &gcsConfig{
			AccessID:              jwtConf.Email,
			PrivateKey:            jwtConf.PrivateKey,
			ReadBatchSize:         conf.ReadBatchSize,
			DownloadURLExpiryTime: conf.DownloadURLExpiryTime,
			UploadURLExpiryTime:   conf.UploadURLExpiryTime,
		},
	}
}

// GenerateUploadURL generates a temporary upload URL for an object the expiry time.
func (c *gcsClient) GenerateUploadURL(object, contentType string) (string, time.Time, error) {
	expires := time.Now().Add(c.config.UploadURLExpiryTime)
	url, err := c.bucket.SignedURL(object, &storage.SignedURLOptions{
		Scheme:         storage.SigningSchemeV4,
		Method:         http.MethodPut,
		Expires:        expires,
		GoogleAccessID: c.config.AccessID,
		PrivateKey:     c.config.PrivateKey,
		Headers:        []string{contentType},
	})
	if err != nil {
		return "", time.Time{}, err
	}
	return url, expires, err
}

// GenerateDownloadURL generates a temporary download URL for an object and the expiry time.
func (c *gcsClient) GenerateDownloadURL(object string) (string, time.Time, error) {
	expires := time.Now().Add(c.config.DownloadURLExpiryTime)
	url, err := c.bucket.SignedURL(object, &storage.SignedURLOptions{
		Scheme:         storage.SigningSchemeV4,
		Method:         http.MethodGet,
		Expires:        expires,
		GoogleAccessID: c.config.AccessID,
		PrivateKey:     c.config.PrivateKey,
	})
	if err != nil {
		return "", time.Time{}, err
	}
	return url, expires, err
}

// DeleteFile deletes the object from bucket.
func (c *gcsClient) DeleteObject(ctx context.Context, object string) error {
	return c.bucket.Object(object).Delete(ctx)
}

// GetObjectSize returns the object size in bytes.
func (c *gcsClient) GetObjectSize(ctx context.Context, object string) (int64, error) {
	objectHandle := c.bucket.Object(object)
	attrs, err := objectHandle.Attrs(ctx)
	if err != nil {
		return 0, err
	}
	return attrs.Size, nil
}

// ReadRange reads data (bytes) from object from given offset with predetermined size from config.
func (c *gcsClient) ReadRange(ctx context.Context, object string, offset int64) ([]byte, error) {
	objectHandle := c.bucket.Object(object)
	reader, err := objectHandle.NewRangeReader(ctx, offset, c.config.ReadBatchSize)
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *gcsClient) GetObjectWriter(ctx context.Context, object string) (w io.Writer, closeFunc func()) {
	// make sure to close this writer once the related process completed
	writer := c.bucket.Object(object).NewWriter(ctx)

	closeFunc = func() {
		writer.Close()
	}
	return writer, closeFunc
}

func (c *gcsClient) GetObjectReader(ctx context.Context, object string) (r io.Reader, closeFunc func(), err error) {
	// make sure to close this reader once the related process completed
	reader, err := c.bucket.Object(object).NewReader(ctx)
	if err != nil {
		return
	}

	closeFunc = func() {
		reader.Close()
	}
	return reader, closeFunc, nil
}

// END __INCLUDE_GCS__
