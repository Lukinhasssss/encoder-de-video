package repositories_test

import (
	"testing"
	"time"

	"github.com/lukinhasssss/encoder-de-video/application/repositories"
	"github.com/lukinhasssss/encoder-de-video/domain"
	"github.com/lukinhasssss/encoder-de-video/framework/database"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestVideoRepositoryDbInsert(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	repo := repositories.VideoRepositoryDb{Db: db}

	repo.Insert(video)

	v, err := repo.Find(video.ID)

	require.Nil(t, err)
	require.NotEmpty(t, v.ID)
	require.Equal(t, v.ID, video.ID)
}
