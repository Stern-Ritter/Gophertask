package utils

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func Coalesce[T ~string | ~int | ~float64 | ~bool | time.Time | time.Duration](firstValue, secondValue T) T {
	var zeroValue T
	if firstValue != zeroValue {
		return firstValue
	}
	return secondValue
}

func ParseDuration(durationStr string) (time.Duration, error) {
	parts := strings.Split(durationStr, ":")
	if len(parts) != 3 {
		return 0, fmt.Errorf("invalid duration format: %s", durationStr)
	}

	hours, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, fmt.Errorf("invalid hours: %v", err)
	}

	minutes, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, fmt.Errorf("invalid minutes: %v", err)
	}

	seconds, err := strconv.Atoi(parts[2])
	if err != nil {
		return 0, fmt.Errorf("invalid seconds: %v", err)
	}

	duration := time.Duration(hours)*time.Hour +
		time.Duration(minutes)*time.Minute +
		time.Duration(seconds)*time.Second

	return duration, nil
}

func FormatDuration(d *durationpb.Duration) string {
	duration := d.AsDuration()

	hours := int(duration.Hours())
	minutes := int(duration.Minutes()) % 60
	seconds := int(duration.Seconds()) % 60

	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
}

func FormatTimestamp(ts *timestamppb.Timestamp) string {
	t := ts.AsTime()
	return t.Format("02.01.2006")
}

func ParseDate(date string) (time.Time, error) {
	return time.Parse("02.01.2006", date)
}
