package agent

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"text/template"
	"time"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"github.com/xaque208/znet/internal/gitwatch"
	"github.com/xaque208/znet/internal/timer"
	"github.com/xaque208/znet/pkg/events"
)

// Agent is an RPC client worker bee.
type Agent struct {
	config Config
	conn   *grpc.ClientConn
}

// NewAgent returns a new *Agent from the given arguments.
func NewAgent(config Config, conn *grpc.ClientConn) *Agent {
	return &Agent{
		config: config,
		conn:   conn,
	}
}

// Subscriptions implements the events.Consumer interface
func (a *Agent) Subscriptions() *events.Subscriptions {
	s := events.NewSubscriptions()

	for _, e := range a.config.Executions {
		for _, x := range e.Events {
			switch x {
			case "NewCommit":
				s.Subscribe(x, a.newCommitHandler)

				b, err := json.Marshal(e.Filter)
				if err != nil {
					log.Errorf("failed to marshal %s filter: %s", x, err)
				}

				f := &gitwatch.GitFilter{}
				err = json.Unmarshal(b, &f)
				if err != nil {
					log.Errorf("failed to unmarshal %s filter into GitFilter: %s", x, err)
				}

				s.Filter(x, f)
			case "NewTag":
				s.Subscribe(x, a.newTagHandler)

				b, err := json.Marshal(e.Filter)
				if err != nil {
					log.Errorf("failed to marshal %s filter: %s", x, err)
				}

				f := &gitwatch.GitFilter{}
				err = json.Unmarshal(b, &f)
				if err != nil {
					log.Errorf("failed to unmarshal %s filter into GitFilter: %s", x, err)
				}

				s.Filter(x, f)
			case "NamedTimer":
				s.Subscribe(x, a.namedTimerHandler)

				f := &timer.TimerFilter{}
				// f.Name = append(f.Name, "ReportFacts")
				s.Filter(x, f)
			default:
				log.WithFields(log.Fields{
					"event": x,
				}).Warn("no execution handler")
			}
		}
	}

	log.WithFields(log.Fields{
		"handlers": s.Handlers,
		"filters":  s.Filters,
	}).Debug("event subscriptions")

	return s
}

func (a *Agent) namedTimerHandler(name string, payload events.Payload) error {
	log.WithFields(log.Fields{
		"name":    name,
		"payload": string(payload),
	}).Warn("TODO")

	return nil
}

func (a *Agent) newExecRequestHandler(name string, payload events.Payload) error {
	var x ExecRequest

	err := json.Unmarshal(payload, &x)
	if err != nil {
		log.Errorf("failed to unmarshal %T: %s", x, err)
	}

	log.Warn("TODO newExecRequestHandler()")

	// returnejnivtfgvbheljfnrilruvlldiennunvcrvbei
	// .executeForEvent(x)

	return nil
}

func (a *Agent) newTagHandler(name string, payload events.Payload) error {
	var x gitwatch.NewTag

	err := json.Unmarshal(payload, &x)
	if err != nil {
		log.Errorf("failed to unmarshal %T: %s", x, err)
	}

	return a.executeForGitEvent(x)
}

func (a *Agent) newCommitHandler(name string, payload events.Payload) error {
	log.Debugf("Agent.newCommitHandler: %+v", string(payload))
	log.Debugf("Agent.newCommitHandler config: %+v", a.config)

	var x gitwatch.NewCommit

	err := json.Unmarshal(payload, &x)
	if err != nil {
		log.Errorf("failed to unmarshal %T: %s", x, err)
	}

	return a.executeForGitEvent(x)
}

func (a *Agent) executeForGitEvent(x interface{}) error {
	log.Tracef("executeForGitEvent %+v", x)

	for _, execution := range a.config.Executions {

		for _, xx := range execution.Events {
			if xx != "" {
				var args []string

				// Render the args as template strings, passing the current x interface.
				for _, v := range execution.Args {
					tmpl, err := template.New("env").Parse(v)
					if err != nil {
						log.Errorf("failed to parse template %s: %s", v, err)
					}

					var buf bytes.Buffer

					err = tmpl.Execute(&buf, x)
					if err != nil {
						log.Error(err)
					}

					args = append(args, buf.String())
				}

				cmd := exec.Command(execution.Command, args...)

				if execution.Dir != "" {
					cmd.Dir = execution.Dir
				}

				var env []string

				// Render the values of the environment variables as templates using
				// the received event.
				for k, v := range execution.Environment {

					tmpl, err := template.New("env").Parse(v)
					if err != nil {
						log.Errorf("failed to parse template %s: %s", v, err)
					}

					var buf bytes.Buffer

					err = tmpl.Execute(&buf, x)
					if err != nil {
						log.Error(err)
					}

					env = append(env, fmt.Sprintf("%s=%s", k, buf.String()))
				}

				if len(env) > 0 {
					cmd.Env = append(os.Environ(), env...)
				}

				start := time.Now()
				// var out bytes.Buffer
				// cmd.Stdout = &out
				output, err := cmd.CombinedOutput()
				if err != nil {
					log.Errorf("command execution failed: %s", err)
				}

				now := time.Now()

				ev := ExecutionResult{
					Time:     &now,
					Command:  execution.Command,
					Args:     args,
					Dir:      execution.Dir,
					Output:   output,
					ExitCode: cmd.ProcessState.ExitCode(),
					Duration: time.Since(start),
				}

				err = events.ProduceEvent(a.conn, ev)
				if err != nil {
					log.Error(err)
				}

			}
		}
	}

	return nil
}
