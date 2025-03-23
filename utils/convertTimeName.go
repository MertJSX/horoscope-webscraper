package utils

func ConvertTimeName(time string) string {
	switch time {
	case "daily":
		return "dneven"
	case "weekly":
		return "sedmichen"
	case "monthly":
		return "mesechen"
	default:
		return "godishen"
	}
}
