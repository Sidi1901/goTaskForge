package queue

import (
	"Sidi1901/goTaskForge/shared/constant"
	"context"
)

func EnqueueTask(ctx context.Context, taskID string) error {
	return rdb.LPush(ctx, constant.TaskQueueName, taskID).Err()
}
