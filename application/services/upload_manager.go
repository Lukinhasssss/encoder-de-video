package services

import (
	"context"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"cloud.google.com/go/storage"
)

type VideoUpload struct {
	Paths        []string
	VideoPath    string
	OutputBucket string
	Errors       []string
}

func NewVideoUpload() *VideoUpload {
  return &VideoUpload{}
}

func (vu *VideoUpload) UploadObject(objectPath string, client *storage.Client, ctx context.Context) error {
  // caminho/x/y/arquivo.mp4
  // split: caminho/x/y/
  // [0] caminho/x/y/arquivo.mp4
  // [1] arquivo.mp4
  path := strings.Split(objectPath, os.Getenv("LOCAL_STORAGE_PATH") + "/")

  f, err := os.Open(objectPath)
  if err != nil {
    return err
  }

  defer f.Close()

  wc := client.Bucket(vu.OutputBucket).Object(path[1]).NewWriter(ctx)

  wc.ACL = []storage.ACLRule{{Entity: storage.AllUsers, Role: storage.RoleReader}}

  if _, err = io.Copy(wc, f); err != nil {
    return err
  }

  if err := wc.Close(); err != nil {
    return err
  }

  return nil
}

func (vu *VideoUpload) loadPaths() error {
  err := filepath.Walk(vu.VideoPath, func(path string, info fs.FileInfo, err error) error {
    if !info.IsDir() {
      vu.Paths = append(vu.Paths, path)
    }

    return nil
  })

  if err != nil {
    return err
  }

  return nil
}

func getClientUpload() (*storage.Client, context.Context, error) {
  ctx := context.Background()

  client, err := storage.NewClient(ctx)
  if err != nil {
    return nil, nil, err
  }

  return client, ctx, nil
}
