package worker

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/abhiraj-ku/health_app/internal/model"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

type EmailWorker struct {
	RedisClient *redis.Client
}

func NewEmailWorker(redisClient *redis.Client) *EmailWorker {
	return &EmailWorker{RedisClient: redisClient}
}

func (w *EmailWorker) EnqueueEmail(user *model.User) {
	emailTask := map[string]interface{}{
		"UserID": user.ID,
		"role":   user.Role,
		"emaiil": user.Name,
	}
	data, err := json.Marshal(emailTask)
	if err != nil {
		log.Printf("failed to marshal the email task %v", err)
		return

	}

	w.RedisClient.LPush(ctx, "emailQueue", data)
}

func (w *EmailWorker) ProcessEmailQueue() {
	for {
		result, err := w.RedisClient.BRPop(ctx, 0*time.Second, "emailQueue").Result()
		if err != nil {
			log.Printf("error popping from email queue %v", err)
			return
		}

		// in case redis return only key or what
		if len(result) < 2 {
			continue
		}

		emailData := result[1]
		// Let's simulate an event of sending email
		log.Printf("seding email to : %s\n", emailData)
		time.Sleep(2 * time.Second)
		log.Println("email sent successfully")
	}
}
