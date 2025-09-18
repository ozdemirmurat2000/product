package utils

import "time"

func IntValue(p *int) int {
	if p != nil {
		return *p
	}
	return 0
}

func StringValue(p *string) string {
	if p != nil {
		return *p
	}
	return ""
}

func Float64Value(p *float64) float64 {
	if p != nil {
		return *p
	}
	return 0
}

func BoolValue(p *bool) bool {
	if p != nil {
		return *p
	}
	return false
}

func TimeValue(p *time.Time) time.Time {
	if p != nil {
		return *p
	}
	return time.Now()
}
