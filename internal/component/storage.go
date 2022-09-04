package component

import (
	"github.com/muhammad-fakhri/archetype-be/internal/component/storage"
	"github.com/muhammad-fakhri/archetype-be/internal/config"
)

func InitStorageClient() storage.Storage {
	return storage.InitGCSClient(config.Get().GCS)
}
