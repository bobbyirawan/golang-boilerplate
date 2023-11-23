package utils

import (
	"time"
)

type (
	TimeService interface {
		Indonesia() time.Time
		Singapore() time.Time
		Istanbul() time.Time
	}

	// WaktuServiceDefault adalah implementasi default dari antarmuka WaktuService.
	timeServiceImpl struct{}
)

// new service
func NewTimeService() TimeService {
	return &timeServiceImpl{}
}

// indonesia
func (ts *timeServiceImpl) Indonesia() time.Time {
	lokasi, err := time.LoadLocation("Asia/Jakara")
	if err != nil {
		return time.Now()
	}

	return time.Now().In(lokasi)
}

func (ts *timeServiceImpl) Singapore() time.Time {
	lokasi, err := time.LoadLocation("Asia/Singapore")
	if err != nil {
		return time.Now()
	}

	return time.Now().In(lokasi)
}

func (ts *timeServiceImpl) Istanbul() time.Time {
	lokasi, err := time.LoadLocation("Europe/Istanbul")
	if err != nil {
		return time.Now()
	}

	return time.Now().In(lokasi)
}
