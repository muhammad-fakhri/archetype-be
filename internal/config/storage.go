package config

// BEGIN __INCLUDE_GCS__
import "time"

type GCSConfig struct {
	CredentialFileName    string
	BucketName            string
	DownloadURLExpiryTime time.Duration
	UploadURLExpiryTime   time.Duration
	ReadBatchSize         int64
}

func (c *Config) initGCSConfig(cfg *configIni) {
	c.GCS = &GCSConfig{
		CredentialFileName:    DefaultConfigPath + cfg.GCSCredentialFileName,
		BucketName:            cfg.GCSBucketName,
		DownloadURLExpiryTime: time.Duration(cfg.GCSDownloadURLExpiryTimeSec) * time.Second,
		UploadURLExpiryTime:   time.Duration(cfg.GCSUploadURLExpiryTimeSec) * time.Second,
		ReadBatchSize:         cfg.GCSReadBatchSize,
	}
}

// END __INCLUDE_GCS__
