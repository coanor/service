package service

import (
	"bytes"
	"testing"
)

func TestUpstartTemplate(t *testing.T) {
	s := &upstart{
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
		Path            string
		HasKillStanza   bool
		HasSetUIDStanza bool
		LogOutput       bool
	}{
		s.Config,
		"/abc/123",
		s.hasKillStanza(),
		s.hasSetUIDStanza(),
		s.Option.bool(optionLogOutput, optionLogOutputDefault),
	}

	var buf bytes.Buffer
	if err := s.template().Execute(&buf, to); err != nil {
		t.Fatal(err)
	}

	t.Logf("%s", buf.String())
}

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
		Path string
	}{
		Config: s.Config,
		Path:   "/abc/123",
	}

	var buf bytes.Buffer
	if err := s.template().Execute(&buf, to); err != nil {
		t.Fatal(err)
	}

	t.Logf("%s", buf.String())
}
