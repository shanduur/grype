package cli

import (
	"strings"
	"testing"
)

func TestCmd(t *testing.T) {
	tests := []struct {
		name       string
		args       []string
		env        map[string]string
		assertions []traitAssertion
	}{
		{
			name: "no-args-shows-help",
			args: []string{},
			assertions: []traitAssertion{
				assertInOutput("an image/directory argument is required"),                              // specific error that should be shown
				assertInOutput("A vulnerability scanner for container images, filesystems, and SBOMs"), // excerpt from help description
				assertFailingReturnCode,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cmd, stdout, stderr := runGrype(t, test.env, test.args...)
			for _, traitFn := range test.assertions {
				traitFn(t, stdout, stderr, cmd.ProcessState.ExitCode())
			}
			if t.Failed() {
				t.Log("STDOUT:\n", stdout)
				t.Log("STDERR:\n", stderr)
				t.Log("COMMAND:", strings.Join(cmd.Args, " "))
			}
		})
	}
}
