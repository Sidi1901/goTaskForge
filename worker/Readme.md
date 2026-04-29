
Processor (core brain)

Design decision
Task Loss solution - Worker pops the task from task_queue -> worker crashes -> task is lost.

Approach 
By using processing queue - Worker pops the task from task queue -> Worker pushes task ID into processing queue -> worker executes/fails -> pops task ID from processing queue only if task execution is completed.

worker/
├── cmd/worker/main.go
└── internal/
    ├── queue/
    │   └── redis.go
    ├── processor/
    │   └── processor.go
    ├── executor/
    │   └── executor.go
    └── database/