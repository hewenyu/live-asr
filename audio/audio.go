package audio

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"runtime"
)

// SaveAudioToFile saves audio data to a file.
func SaveAudioToFile(filename string, audioData []byte) error {
	err := os.WriteFile(filename, audioData, 0644)
	if err != nil {
		return fmt.Errorf("failed to save audio to file: %v", err)
	}
	return nil
}

// CaptureAudioOutputCrossPlatform captures audio output from the system.
func CaptureAudioOutputCrossPlatform() ([]byte, error) {
	var cmd *exec.Cmd
	switch os := runtime.GOOS; os {
	case "windows":
		cmd = exec.Command("powershell", "-c", "Add-Type -TypeDefinition @\"using System;using System.Runtime.InteropServices;public class AudioCapture { [DllImport(\\\"winmm.dll\\\")] public static extern int waveInGetNumDevs(); }\"@; [AudioCapture]::waveInGetNumDevs()")
	case "darwin":
		cmd = exec.Command("sh", "-c", "rec -c 1 -r 44100 -b 16 -e signed-integer -t wav -")
	case "linux":
		cmd = exec.Command("arecord", "-f", "cd", "-t", "wav", "-d", "10", "-q", "-")
	default:
		return nil, fmt.Errorf("unsupported platform")
	}

	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("failed to capture audio: %v", err)
	}
	return out.Bytes(), nil
}

func CaptureAudioOutputCrossPlatformStream() (io.ReadCloser, error) {
	var cmd *exec.Cmd
	switch os := runtime.GOOS; os {
	case "windows":
		cmd = exec.Command("powershell", "-c", "Add-Type -TypeDefinition @\"using System;using System.Runtime.InteropServices;public class AudioCapture { [DllImport(\\\"winmm.dll\\\")] public static extern int waveInGetNumDevs(); }\"@; [AudioCapture]::waveInGetNumDevs()")
	case "darwin":
		cmd = exec.Command("sh", "-c", "rec -c 1 -r 44100 -b 16 -e signed-integer -t wav -")
	case "linux":
		cmd = exec.Command("arecord", "-f", "cd", "-t", "wav", "-d", "10", "-q", "-")
	default:
		return nil, fmt.Errorf("unsupported platform")
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, fmt.Errorf("failed to get stdout pipe: %v", err)
	}

	if err := cmd.Start(); err != nil {
		return nil, fmt.Errorf("failed to start command: %v", err)
	}

	return stdout, nil
}
