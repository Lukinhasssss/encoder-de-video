package services

import (
	"encoding/json"
	"os"
	"sync"
	"time"

	"github.com/lukinhasssss/encoder-de-video/domain"
	"github.com/lukinhasssss/encoder-de-video/framework/utils"
	uuid "github.com/satori/go.uuid"
	"github.com/streadway/amqp"
)

type JobWorkerResult struct {
	Job     domain.Job
	Message *amqp.Delivery
	Error   error
}

var Mutex = &sync.Mutex{}

func JobWorker(messageChannel chan amqp.Delivery, returnChannel chan JobWorkerResult, jobService JobService, job domain.Job, workerID int) {
	// Pega body do JSON
	// Valida se o json é válido
	// Valida o vídeo
	// Insere o vídeo no banco de dados
	// Start

	for message := range messageChannel {
		err := utils.IsJson(string(message.Body))

		if err != nil {
			returnChannel <- returnJobResult(job, message, err)
			continue
		}

		Mutex.Lock()
		err = json.Unmarshal(message.Body, &jobService.VideoService.Video)
		jobService.VideoService.Video.ID = uuid.NewV4().String()
		Mutex.Unlock()

		if err != nil {
			returnChannel <- returnJobResult(job, message, err)
			continue
		}

		err = jobService.VideoService.Video.Validate()

		if err != nil {
			returnChannel <- returnJobResult(job, message, err)
			continue
		}

		Mutex.Lock()
		err = jobService.VideoService.InsertVideo()
		Mutex.Unlock()

		if err != nil {
			returnChannel <- returnJobResult(job, message, err)
			continue
		}

		job.ID = uuid.NewV4().String()
		job.Video = jobService.VideoService.Video
		job.OutputBucketPath = os.Getenv("OUTPUT_BUCKET_NAME")
		job.Status = "STARTING"
		job.CreatedAt = time.Now()

		Mutex.Lock()
		_, err = jobService.JobRepository.Insert(&job)
		Mutex.Unlock()

		if err != nil {
			returnChannel <- returnJobResult(job, message, err)
			continue
		}

		jobService.Job = &job

		err = jobService.Start()

		if err != nil {
			returnChannel <- returnJobResult(job, message, err)
			continue
		}

		returnChannel <- returnJobResult(job, message, nil)
	}
}

func returnJobResult(job domain.Job, message amqp.Delivery, err error) JobWorkerResult {
	result := JobWorkerResult{
		Job:     job,
		Message: &message,
		Error:   err,
	}

	return result
}
