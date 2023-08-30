package appconfig

import (
	"fmt"
	"github.com/codingconcepts/env"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var environmentVarKeys = []string{"MY_HOME_HOST_NAME", "MY_HOME_HOST_PORT"}

func TestAppConfig_Init(t *testing.T) {
	tests := []struct {
		name       string
		envs       map[string]string
		flags      []string
		funcEnvSet func(i interface{}) error
		wantErr    error
		wantData   AppConfig
	}{
		{
			name:       "nothing passed",
			envs:       map[string]string{},
			flags:      []string{},
			funcEnvSet: env.Set,
			wantErr:    nil,
			wantData: AppConfig{
				DebugMode: false,
				Host:      "localhost",
				Port:      "8080",
			},
		},
		{
			name: "all flags passed",
			envs: map[string]string{
				"MY_HOME_HOST_NAME": "new_host_name",
				"MY_HOME_HOST_PORT": "2020",
			},
			flags:      []string{},
			funcEnvSet: env.Set,
			wantErr:    nil,
			wantData: AppConfig{
				DebugMode: false,
				Host:      "new_host_name",
				Port:      "2020",
			},
		},
		{
			name: "all flags passed",
			envs: map[string]string{
				"MY_HOME_HOST_NAME": "new_host_name",
				"MY_HOME_HOST_PORT": "2020",
			},
			flags:      []string{"-debug", "-host", "flag_host", "-port", "1010"},
			funcEnvSet: env.Set,
			wantErr:    nil,
			wantData: AppConfig{
				DebugMode: true,
				Host:      "flag_host",
				Port:      "1010",
			},
		},
		{
			name:       "bad flags passed",
			envs:       map[string]string{},
			flags:      []string{"-joe"},
			funcEnvSet: env.Set,
			wantErr:    fmt.Errorf("flag provided but not defined: -joe"),
			wantData:   AppConfig{},
		},
		{
			name:  "env failed",
			envs:  map[string]string{},
			flags: []string{},
			funcEnvSet: func(i interface{}) error {
				return fmt.Errorf("env has faild you")
			},
			wantErr:  fmt.Errorf("env has faild you"),
			wantData: AppConfig{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			envSet = tt.funcEnvSet
			baseDir := t.TempDir()
			assert.NoError(t, os.Chdir(baseDir))
			setEnvironmentVars(t, tt.envs)

			// setup flags for test
			os.Args = append([]string{"unittest"}, tt.flags...)

			// Make Call
			ac := &AppConfig{}
			err := ac.Init()

			validateError(t, tt.wantErr, err)
			if tt.wantErr == nil {
				assert.EqualValues(t, &tt.wantData, ac)
			}

			// Clear for good measurement
			setEnvironmentVars(t, map[string]string{})
		})
	}
}

func validateError(t *testing.T, want error, have error) {
	if want != nil {
		assert.Equal(t, want.Error(), have.Error())
	} else {
		assert.NoError(t, have)
	}
}

func setEnvironmentVars(t *testing.T, v map[string]string) {
	// clear any previous vars
	for _, key := range environmentVarKeys {
		if err := os.Unsetenv(key); err != nil {
			t.Errorf("failed to clear env vars: %s", err.Error())
			return
		}
	}

	// Set values
	for key, value := range v {
		if err := os.Setenv(key, value); err != nil {
			t.Errorf("failed to set env var [%s](%s): %s", key, value, err.Error())
			return
		}
	}
}
