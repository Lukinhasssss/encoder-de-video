package services_test

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/lukinhasssss/encoder-de-video/application/repositories"
	"github.com/lukinhasssss/encoder-de-video/application/services"
	"github.com/lukinhasssss/encoder-de-video/domain"
	"github.com/lukinhasssss/encoder-de-video/framework/database"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func init() {
  err := godotenv.Load("../../.env")

  if err != nil {
    log.Fatalf("Error loading .env file")
  }
}

func prepare() (*domain.Video, repositories.VideoRepositoryDb) {
	db := database.NewDbTest()
	defer db.Close()

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "sample.mp4"
	video.CreatedAt = time.Now()

	repo := repositories.VideoRepositoryDb{Db: db}

	return video, repo
}

func TestVideoServiceDownload(t *testing.T) {
  video, repo := prepare()

  videoService := services.NewVideoService()
  videoService.Video = video
  videoService.VideoRepository = repo

  err := videoService.Download(os.Getenv("INPUT_BUCKET_NAME"))
  require.Nil(t, err)

  err = videoService.Fragment()
  require.Nil(t, err)

  err = videoService.Encode()
  require.Nil(t, err)

  err = videoService.Finish()
  require.Nil(t, err)
}
