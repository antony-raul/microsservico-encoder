package repositories_test

import (
	"testing"
	"time"

	"github.com/antony-raul/microsservico-encoder/application/repositories"
	"github.com/antony-raul/microsservico-encoder/domain"
	"github.com/antony-raul/microsservico-encoder/framework/database"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestVideoRepositoryDbInsert(t *testing.T) {
	db := database.NewDbTest()

	video := domain.NewVideo()
	video.ID = uuid.New().String()
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	repo := repositories.NewVideoRepository(db)

	repo.Insert(video)

	v, err := repo.Find(video.ID)

	require.NotEmpty(t, v.ID)
	require.Nil(t, err)
	require.Equal(t, v.ID, video.ID)
}
