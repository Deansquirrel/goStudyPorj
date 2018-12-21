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
	//相应超时时间
	TimeoutNS time.Duration `toml:"timeoutNS"`
	//每秒载荷量
	Lps uint32 `toml:"lps"`
	//负载持续时间
	DurationNS time.Duration `toml:"durationNS"`
}
