package agent

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"

	log "github.com/sirupsen/logrus"

	"github.com/xaque208/znet/internal/events"
	"github.com/xaque208/znet/internal/gitwatch"
)

type Agent struct {
	config Config
}

func NewAgent(config Config) *Agent {
	return &Agent{
		config: config,
	}

}

func (a *Agent) EventNames() []string {
	var names []string

	for _, e := range a.config.Executions {
		for _, x := range e.Events {
			if x != "" {
				names = append(names, x)
			}
		}
	}

	log.Debugf("agent responding to %d event names: %+v", len(names), names)

	return names
}

// Subscriptions implements the events.Consumer interface
func (a *Agent) Subscriptions() map[string][]events.Handler {
	s := events.NewSubscriptions()

	for _, e := range a.config.Executions {
		for _, x := range e.Events {
			switch x {
			case "NewCommit":
				s.Subscribe(x, a.newCommitHandler)
			default:
				log.Errorf("unhandled execution event %s", x)
			}
		}
	}

	log.Debugf("event subscriptions %+v", s.Table)

	return s.Table
}

func (a *Agent) newCommitHandler(name string, payload events.Payload) error {
	log.Debugf("Agent.newCommitHandler: %+v", string(payload))
	log.Debugf("Agent.newCommitHandler config: %+v", a.config)

	var x gitwatch.NewCommit

	err := json.Unmarshal(payload, &x)
	if err != nil {
		log.Errorf("failed to unmarshal %T: %s", x, err)
	}

	for _, e := range a.config.Executions {

		if len(e.Filter) > 0 {
			var key string
			var value interface{}

			if val, ok := e.Filter["key"]; ok {
				key = val.(string)
			}

			if val, ok := e.Filter["value"]; ok {
				value = val
			}

			log.Debugf("filter key %s", key)
			log.Debugf("filter value %+v", value)

			// if val, ok := e.Filter["name"]; ok {
			// 	if val != x.Name {
			// 		continue
			// 	}
			// }
		}

		for _, x := range e.Events {
			if x != "" {
				cmd := exec.Command(e.Command, e.Args...)

				if e.Dir != "" {
					cmd.Dir = e.Dir
				}

				var env []string

				for k, v := range e.Environment {
					env = append(env, fmt.Sprintf("%s=%s", k, v))
				}

				if len(env) > 0 {
					cmd.Env = append(os.Environ(), env...)
				}

				// cmd.Stdin = strings.NewReader("some input")
				var out bytes.Buffer
				cmd.Stdout = &out
				err := cmd.Run()
				if err != nil {
					log.Fatal(err)
				}
			}
		}
	}

	return nil
}