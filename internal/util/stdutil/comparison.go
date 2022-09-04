package stdutil

import "time"

func GetValueOrDefault(value interface{}, defaultValue interface{}) interface{} {
	if value != nil {
		return value
	}
	return defaultValue
}

func GetStringOrDefault(value, def string) string {
	if value == "" {
		return def
	}
	return value
}

func GetIntOrDefault(value, def int) int {
	if value == 0 {
		return def
	}
	return value
}

func GetTimeSecondOrDefault(value, def int) time.Duration {
	t := GetIntOrDefault(value, def)
	if t < 0 {
		t = def
	}
	return time.Duration(t) * time.Second
}
