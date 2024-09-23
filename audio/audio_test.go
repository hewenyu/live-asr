package audio

import (
	"bytes"
	"io"
	"os"
	"os/exec"
	"testing"
)

var execCommand = exec.Command

func TestCaptureAudioOutput(t *testing.T) {
	// Mock exec.Command to avoid actually running the arecord command
	execCommand = func(name string, arg ...string) *exec.Cmd {
		cmd := exec.Command("echo", "mocked audio data")
		var out bytes.Buffer
		out.WriteString("mocked audio data")
		cmd.Stdout = &out
		return cmd
	}

	defer func() { execCommand = exec.Command }() // Restore original exec.Command after test

	output, err := CaptureAudioOutputCrossPlatform()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	expectedOutput := "mocked audio data"
	if string(output) != expectedOutput {
		t.Errorf("expected %q, got %q", expectedOutput, string(output))
	}
}

func TestSaveAudioToFile(t *testing.T) {
	// Mock exec.Command to avoid actually running the arecord command
	execCommand = func(name string, arg ...string) *exec.Cmd {
		cmd := exec.Command("echo", "mocked audio data")
		var out bytes.Buffer
		out.WriteString("mocked audio data")
		cmd.Stdout = &out
		return cmd
	}

	audioData, err := CaptureAudioOutputCrossPlatform()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	filename := "test_audio.wav"

	err = SaveAudioToFile(filename, audioData)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	defer os.Remove(filename) // Clean up the test file after the test

	// Verify the file was written correctly
	writtenData, err := os.ReadFile(filename)
	if err != nil {
		t.Fatalf("expected no error reading file, got %v", err)
	}

	if string(writtenData) != string(audioData) {
		t.Errorf("expected file content %q, got %q", string(audioData), string(writtenData))
	}
}

func TestCaptureAudioOutputCrossPlatformStream(t *testing.T) {
	// Mock exec.Command to avoid actually running the arecord command
	execCommand = func(name string, arg ...string) *exec.Cmd {
		cmd := exec.Command("echo", "mocked audio data")
		var out bytes.Buffer
		out.WriteString("mocked audio data")
		cmd.Stdout = &out
		return cmd
	}

	defer func() { execCommand = exec.Command }() // Restore original exec.Command after test

	stream, err := CaptureAudioOutputCrossPlatformStream()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	defer stream.Close()

	var buf bytes.Buffer
	_, err = io.Copy(&buf, stream)
	if err != nil {
		t.Fatalf("expected no error copying stream, got %v", err)
	}

	expectedOutput := "mocked audio data\n"
	if buf.String() != expectedOutput {
		t.Errorf("expected %q, got %q", expectedOutput, buf.String())
	}
}
