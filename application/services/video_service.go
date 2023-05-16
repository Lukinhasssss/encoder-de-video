package services

import (
	"context"
	"io/ioutil"
	"log"
	"os"

	"cloud.google.com/go/storage"
	"github.com/lukinhasssss/encoder-de-video/application/repositories"
	"github.com/lukinhasssss/encoder-de-video/domain"
)

type VideoService struct {
  Video *domain.Video
  VideoRepository repositories.VideoRepository
}

func NewVideoService() VideoService {
  return VideoService{}
}

func (v *VideoService) Download(bucketName string) error {
  ctx := context.Background()

  client, err := storage.NewClient(ctx)

  if err != nil {
    return err
  }

  bkt := client.Bucket(bucketName)
  obj := bkt.Object(v.Video.FilePath)

  r, err := obj.NewReader(ctx)

  if err != nil {
    return err
  }

  defer r.Close()

  body, err := ioutil.ReadAll(r)

  if err != nil {
    return err
  }

  f, err := os.Create(os.Getenv("LOCAL_STORAGE_PATH") + "/" + v.Video.ID + ".mp4")

  if err != nil {
    return err
  }

  _, err = f.Write(body)

  if err != nil {
    return err
  }

  defer f.Close()

  log.Printf("video %v has been stored", v.Video.ID)

  return nil
}
