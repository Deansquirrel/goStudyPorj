package global

import "time"

type SysConfig struct {
	Total      total      `toml:"Total"`
	TestConfig testConfig `toml:"TestConfig"`
}

type total struct {
	Title string `toml:"title"`
}

type testConfig struct {
	TimeoutNS  time.Duration `toml:"timeoutNS"`
	Lps        uint32        `toml:"lps"`
	DurationNS time.Duration `toml:"durationNS"`
}
