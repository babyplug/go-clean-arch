package background

import (
	"log"
	"time"

	"github.com/babyplug/go-clean-arch/internal/core/port"
)

func StartUserCountLogger(repo port.UserRepository, stopCh <-chan struct{}) {
	ticker := time.NewTicker(10 * time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				count, err := repo.Count()
				if err == nil {
					log.Printf("User count: %d", count)
				} else {
					log.Printf("Failed to count users: %v", err)
				}
			case <-stopCh:
				ticker.Stop()
				return
			}
		}
	}()
}
