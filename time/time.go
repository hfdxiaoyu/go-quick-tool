package time

import (
	"time"
)

func TimeFormat(ftime time.Time) string {
	return ftime.Format("2006-01-02 15:04:05")
}

func TimeFileFormat(ftime time.Time) string {
	return ftime.Format("20060102150405")
}

func UnixTimeFormat(ftime int64) string {
	return TimeFormat(time.Unix(ftime, 0))

}

// StringToTime String类型转为时间
// eg: stime: 2024-06-30 11:26:05 format:  2006-01-02 15:04:05
func StringToTime(stime, format string) (time.Time, error) {
	t, err := time.Parse(stime, format)
	if err != nil {
		return t, err
	}
	return t, nil
}
