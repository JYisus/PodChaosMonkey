package scheduler

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScheduler_Constructor(t *testing.T) {
	tests := []struct {
		name           string
		scheduleFormat string
		wantError      bool
		errorString    string
	}{
		{
			name:           "Scheduler accepts cron format",
			scheduleFormat: CronFormat,
			wantError:      false,
			errorString:    "",
		},
		{
			name:           "Scheduler accepts accept cron-seconds format",
			scheduleFormat: CronSecondsFormat,
			wantError:      false,
			errorString:    "",
		},
		{
			name:           "Scheduler constructor fails on invalid format",
			scheduleFormat: "InvalidFormat",
			wantError:      true,
			errorString:    "unknown schedule format: InvalidFormat",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			cronScheduler, err := NewCronScheduler(tc.scheduleFormat)
			if tc.wantError {
				assert.Error(t, err, fmt.Errorf(tc.errorString))
				assert.Nil(t, cronScheduler)
				return
			}
			assert.NoError(t, err)
			assert.NotNil(t, cronScheduler)
		})
	}
}
