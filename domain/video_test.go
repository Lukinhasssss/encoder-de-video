package domain_test

import (
	"testing"
	"time"

	"github.com/lukinhasssss/encoder-de-video/domain"
	"github.com/stretchr/testify/require"
  uuid "github.com/satori/go.uuid"
)

func TestValidateIfVideoIsEmpty(t *testing.T) {
  video := domain.NewVideo()

  err := video.Validate()

  require.Error(t, err)
}

func TestValidateIfVideoIdIsAValidUuid(t *testing.T) {
  video := domain.NewVideo()

  video.ID = "abc"
  video.ResourceID = "a"
  video.FilePath = "path"
  video.CreatedAt = time.Now()

  err := video.Validate()

  require.Error(t, err)
}

func TestVideoValidation(t *testing.T) {
  video := domain.NewVideo()

  video.ID = uuid.NewV4().String()
  video.ResourceID = "a"
  video.FilePath = "path"
  video.CreatedAt = time.Now()

  err := video.Validate()

  require.Nil(t, err)
}
