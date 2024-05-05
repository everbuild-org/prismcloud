package client

import "fmt"

var (
	MB_MULTIPLIER int64 = 1000000
	GB_MULTIPLIER int64 = 1000 * MB_MULTIPLIER
	Mi_MULTIPLIER int64 = 1024 * 1024
	Gi_MULTIPLIER int64 = 1024 * Mi_MULTIPLIER
)

func ParseRam(ram string) (int64, error) {
	// Suffixed with 'Gi', 'Mi', 'GB', 'MB'

	if len(ram) < 3 {
		return 0, fmt.Errorf("invalid RAM format, must be suffixed with 'Gi', 'Mi', 'GB', 'MB'")
	}

	multiplier := int64(0)

	if ram[len(ram)-2:] == "Gi" {
		multiplier = Gi_MULTIPLIER
	}

	if ram[len(ram)-2:] == "Mi" {
		multiplier = Mi_MULTIPLIER
	}

	if ram[len(ram)-2:] == "GB" {
		multiplier = GB_MULTIPLIER
	}

	if ram[len(ram)-2:] == "MB" {
		multiplier = MB_MULTIPLIER
	}

	if multiplier == 0 {
		return 0, fmt.Errorf("invalid RAM format, must be suffixed with 'Gi', 'Mi', 'GB', 'MB'")
	}

	ram = ram[:len(ram)-2]

	var ramInt int64
	_, err := fmt.Sscanf(ram, "%d", &ramInt)
	if err != nil {
		return 0, fmt.Errorf("invalid RAM format, must be a number")
	}

	return ramInt * multiplier, nil
}

func FmtRam(ram int64) string {
	if ram >= GB_MULTIPLIER {
		return fmt.Sprintf("%dGB", ram/int64(GB_MULTIPLIER))
	}

	return fmt.Sprintf("%dMB", ram/int64(MB_MULTIPLIER))
}
