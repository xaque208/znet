package config

// TimerConfig is the information necessary for Timer to generate timers.
type TimerConfig struct {
	TimeZone       string              `yaml:"timezone"`
	ReloadInterval int                 `yaml:"reload_interval"`
	FutureLimit    int                 `yaml:"future_limit"`
	RepeatEvents   []RepeatEventConfig `yaml:"repeat_events"`
	Events         []EventConfig       `yaml:"events"`
}

type EventConfig struct {
	// Produce is the name of the event to emit.
	Produce string   `yaml:"produce"`
	Time    string   `yaml:"time"`
	Days    []string `yaml:"days"`
}

type RepeatEventConfig struct {
	// Produce is the name of the event to emit.
	Produce string `yaml:"produce"`
	Every   struct {
		Seconds int `yaml:"seconds"`
	} `yaml:"every"`
}
