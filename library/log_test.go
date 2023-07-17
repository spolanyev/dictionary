//@author Stanislav Polaniev <spolanyev@gmail.com>

package library

import (
	"os"
	"testing"
)

func TestLogLevelOutput(t *testing.T) {
	cases := []struct {
		testName     string
		logLevel     string
		messageLevel []string
		message      []string
		want         string
	}{
		{
			"1",
			LogLevelInfo,
			[]string{LogLevelInfo},
			[]string{"info"},
			"info\n",
		},
		{
			"2",
			"",
			[]string{LogLevelInfo},
			[]string{"info"},
			"",
		},
		{
			"3",
			LogLevelError,
			[]string{LogLevelDebug, LogLevelInfo, LogLevelWarning, LogLevelError},
			[]string{"debug", "info", "warning", "error"},
			"error\n",
		},
		{
			"4",
			LogLevelWarning,
			[]string{LogLevelDebug, LogLevelInfo, LogLevelWarning, LogLevelError},
			[]string{"debug", "info", "warning", "error"},
			"warning\nerror\n",
		},
		{
			"5",
			LogLevelInfo,
			[]string{LogLevelDebug, LogLevelInfo, LogLevelWarning, LogLevelError},
			[]string{"debug", "info", "warning", "error"},
			"info\nwarning\nerror\n",
		},
		{
			"6",
			LogLevelDebug,
			[]string{LogLevelDebug, LogLevelInfo, LogLevelWarning, LogLevelError},
			[]string{"debug", "info", "warning", "error"},
			"debug\ninfo\nwarning\nerror\n",
		},
		{
			"7",
			LogLevelError,
			[]string{"unknown"},
			[]string{"unknown"},
			"[Using unknown log level] unknown\n",
		},
	}

	for _, theCase := range cases {
		reader, writer, err := os.Pipe()
		if err != nil {
			t.Fatalf("Failed to create pipe: %s", err)
		}

		oldStdout := os.Stdout
		os.Stdout = writer

		err = os.Setenv("LOG_LEVEL", theCase.logLevel)
		if err != nil {
			t.Fatalf("Failed to set log level %s: %s", theCase.logLevel, err)
			return
		}

		for i := range theCase.messageLevel {
			Log(theCase.messageLevel[i], theCase.message[i])
		}

		os.Stdout = oldStdout
		err = writer.Close()
		if err != nil {
			t.Fatalf("Failed to close writer: %s", err)
			return
		}

		var buf []byte
		buf = make([]byte, 1024)
		n, _ := reader.Read(buf)

		got := string(buf[:n])

		if got != theCase.want {
			t.Errorf("Test %v: output == %v, want %v", theCase.testName, got, theCase.want)
		}
	}
}
