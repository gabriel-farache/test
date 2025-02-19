package params

import (
	"os"
	"testing"

	"github.com/AlecAivazis/survey/v2"
)

func TestNewCliOptions(t *testing.T) {
	cliOpts := NewCliOptions()

	if cliOpts == nil {
		t.Fatalf("NewCliOptions() = %v, want non-nil", cliOpts)
	}

	askOpts := &survey.AskOptions{}
	err := cliOpts.AskOpts(askOpts)

	if err != nil {
		t.Errorf("cliOpts.AskOpts() error = %v, want nil", err)
	}

	if askOpts.Stdio.In != os.Stdin {
		t.Errorf("cliOpts.AskOpts() Stdio.In = %v, want %v", askOpts.Stdio.In, os.Stdin)
	}

	if askOpts.Stdio.Out != os.Stdout {
		t.Errorf("cliOpts.AskOpts() Stdio.Out = %v, want %v", askOpts.Stdio.Out, os.Stdout)
	}

	if askOpts.Stdio.Err != os.Stderr {
		t.Errorf("cliOpts.AskOpts() Stdio.Err = %v, want %v", askOpts.Stdio.Err, os.Stderr)
	}
}
