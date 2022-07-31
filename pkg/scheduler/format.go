package scheduler

const (
	CronFormat        string = "cron"
	CronSecondsFormat string = "cron-seconds"
)

func IsValidFormat(format string) bool {
	return format == CronFormat || format == CronSecondsFormat
}
