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

func TestJobRepositoryDbInsert(t *testing.T) {
	db := database.NewDbTest()

	video := domain.NewVideo()
	video.ID = uuid.New().String()
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	repo := repositories.NewVideoRepository(db)
	repo.Insert(video)

	job, err := domain.NewJob("output_path", "Peding", video)

	require.Nil(t, err)

	repoJob := repositories.NewJobRepository(db)
	repoJob.Insert(job)

	j, err := repoJob.Find(job.ID)
	require.NotEmpty(t, j.ID)
	require.Nil(t, err)
	require.Equal(t, j.ID, job.ID)
	require.Equal(t, j.VideoID, video.ID)
}

func TestJobRepositoryDbUpdate(t *testing.T) {
	db := database.NewDbTest()

	video := domain.NewVideo()
	video.ID = uuid.New().String()
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	repo := repositories.NewVideoRepository(db)
	repo.Insert(video)

	job, err := domain.NewJob("output_path", "Peding", video)

	require.Nil(t, err)

	repoJob := repositories.NewJobRepository(db)
	repoJob.Insert(job)

	job.Status = "Complete"
	repoJob.Update(job)

	j, err := repoJob.Find(job.ID)
	require.NotEmpty(t, j.ID)
	require.Nil(t, err)
	require.Equal(t, j.Status, job.Status)
}
