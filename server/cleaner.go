package server

import (
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

func RunCleaner(stop chan os.Signal) {
	ticker := time.NewTicker(1 * time.Minute)
	go func() {
		for {
			select {
			case <-ticker.C:
				if err := cleanExpiredSessions(); err != nil {
					log.Errorf("clean expired sessions failed: %v", err)
				}
			case <-stop:
				ticker.Stop()
				return
			}
		}
	}()
}

func cleanExpiredSessions() error {
	return nil
}
