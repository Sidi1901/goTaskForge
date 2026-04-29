package queue

import (
	"Sidi1901/goTaskForge/shared/constant"
	q "Sidi1901/goTaskForge/shared/queue"
	"context"
	"encoding/json"
)

func TaskConsume(ctx context.Context) (*q.Message, error) {
	result, err := rdb.BRPop(ctx, 0, constant.TaskQueueName).Result()
	if err != nil {
		return nil, err
	}

	var msg q.Message
	if err := json.Unmarshal([]byte(result[1]), &msg); err != nil {
		return nil, err
	}

	return &msg, nil
}
