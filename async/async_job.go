package async

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/go-redis/redis/v7"
)

// Param structure defines the job parameters.
type Param struct {
	Payload string
}

// Job structure will be used to handle every jobs.
// Every method (except "Schedule" and "Execute") accept only one string parameter.
// This param will be decoded into a specific structure (manually defined into your job).
// Redis and MySQL is mandatory to use this feature.
// Redis will be used to handle every queue and MySQL to store failed jobs.
type Job struct {
	Name       string
	MethodName string
	Params     Param
	Queue      string
}

// Schedule specific job
func (j *Job) Schedule(queueName string, redis *redis.Client) error {
	jobStr, err := json.Marshal(j)
	if err != nil {
		return err
	}

	queue := fmt.Sprintf("queue:%s", queueName)
	redis.RPush(queue, jobStr)

	return nil
}

// Execute specific job
func (j *Job) Execute() (bool, error) {
	r := reflect.ValueOf(Job{})
	method := r.MethodByName(j.MethodName)
	result := method.Call([]reflect.Value{reflect.ValueOf(j.Params.Payload)})

	// Put job on failed table
	if result[1].Interface() != nil {
		err := result[1].Interface().(error)

		return false, err
	}

	return true, nil
}
