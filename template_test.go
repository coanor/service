package service

import (
	"bytes"
	"testing"
)

func TestSysvTemplate(t *testing.T) {

	s := &sysv{
		Config: &Config{
			Name:       "abc",
			Executable: "/abc/123",
			Envs: map[string]string{
				"ENV_1": "1",
				"ENV_2": "2",
			},
		},
	}

	var to = &struct {
		*Config
	}{
		s.Config,
	}

	var buf bytes.Buffer
	if err := s.template().Execute(&buf, to); err != nil {
		t.Error()
	}

	t.Logf("%s", buf.String())
}

func TestBuildTemplate2(t *testing.T) {

	s := &darwinLaunchdService{
		Config: &Config{
			Name:       "abc",
			Executable: "/abc/123",
			Envs: map[string]string{
				"ENV_1": "1",
				"ENV_2": "2",
			},
		},
	}

	var to = &struct {
		*Config
		Path string

		KeepAlive, RunAtLoad bool
		SessionCreate        bool
		StandardOut          bool
		StandardError        bool
	}{
		Config:        s.Config,
		Path:          "some/path",
		KeepAlive:     s.Option.bool(optionKeepAlive, optionKeepAliveDefault),
		RunAtLoad:     s.Option.bool(optionRunAtLoad, optionRunAtLoadDefault),
		SessionCreate: s.Option.bool(optionSessionCreate, optionSessionCreateDefault),
	}

	var buf bytes.Buffer
	if err := s.template().Execute(&buf, to); err != nil {
		t.Error()
	}

	t.Logf("%s", buf.String())
}
