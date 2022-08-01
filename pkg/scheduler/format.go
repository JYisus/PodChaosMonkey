package scheduler

const (
	CronFormat        string = "cron"
	CronSecondsFormat string = "cron-seconds"
)

// IsValidFormat checks if the provided format is valid
func IsValidFormat(format string) bool {
	return format == CronFormat || format == CronSecondsFormat
}
