package queue

import (
	"notification-service/services"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

// NotificationQueue holds the queue of notification tasks.
type NotificationQueue struct {
	queue     []NotificationTask
	rateLimit int
	interval  time.Duration
	mu        sync.Mutex
}

// NotificationTask represents a task to send a notification.
type NotificationTask struct {
	Notification services.Notification
	Recipient    string
	Message      string
	MaxRetries   int
	RetryDelay   time.Duration
	RetryCount   int
}

// NewNotificationQueue creates a new NotificationQueue with the given rate limit and interval.
func NewNotificationQueue(rateLimit int, interval time.Duration) *NotificationQueue {
	return &NotificationQueue{
		queue:     make([]NotificationTask, 0),
		rateLimit: rateLimit,
		interval:  interval,
	}
}

// AddTask adds a new notification task to the queue.
func (q *NotificationQueue) AddTask(task NotificationTask) {
	q.mu.Lock()
	q.queue = append(q.queue, task)
	q.mu.Unlock()
}

// StartWorker starts processing the queue at the defined rate.
func (q *NotificationQueue) StartWorker() {
	ticker := time.NewTicker((q.interval))
	defer ticker.Stop()

	for {
		<-ticker.C
		q.processTasks()
	}
}

// processTasks processes up to rateLimit tasks from the queue.
func (q *NotificationQueue) processTasks() {
	q.mu.Lock()
	defer q.mu.Unlock()

	limit := q.rateLimit
	if len(q.queue) < limit {
		limit = len(q.queue)
	}

	for i := 0; i < limit; i++ {
		task := q.queue[i]
		go q.processTaskWithRetry(task)
	}

	// Remove processed tasks from the queue
	q.queue = q.queue[limit:]
}

// processTaskWithRetry processes a single notification task with retry logic.
func (q *NotificationQueue) processTaskWithRetry(task NotificationTask) {
	for {
		err := task.Notification.Send(task.Recipient, task.Message)
		if err != nil {
			logrus.Errorf("Error sending notification: %v\n", err)
			task.RetryCount++
			if task.RetryCount > task.MaxRetries {
				logrus.Fatalf("Max retries reached for recipient %s. Giving up.\n", task.Recipient)
				return
			}
			logrus.Infof("Retrying... (%d/%d)\n", task.RetryCount, task.MaxRetries)
			time.Sleep(task.RetryDelay)
		} else {
			logrus.Infof("Notification sent successfully to %s\n", task.Recipient)
			return
		}
	}
}
