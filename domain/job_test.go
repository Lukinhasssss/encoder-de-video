package domain_test

import (
	"testing"
	"time"

	"github.com/lukinhasssss/encoder-de-video/domain"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestNewJob(t *testing.T) {
  video := domain.NewVideo()

  video.ID = uuid.NewV4().String()
  video.ResourceID = "a"
  video.FilePath = "path"
  video.CreatedAt = time.Now()

  job, err := domain.NewJob("path", "Converted", video)

  require.NotNil(t, job)
  require.Nil(t, err)
}
