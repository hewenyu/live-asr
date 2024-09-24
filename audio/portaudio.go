package audio

import (
	"github.com/gordonklaus/portaudio"
)

const (
	DefaultSampleRate = 16000 // DefaultSampleRate is the default sample rate for audio
)

type PortAudio struct {
	InputDevice *portaudio.DeviceInfo
	params      portaudio.StreamParameters
}

func NewPortAudio(device *portaudio.DeviceInfo) *PortAudio {
	param := portaudio.StreamParameters{}
	param.Input.Device = device
	param.Input.Channels = 1
	param.Input.Latency = device.DefaultHighInputLatency

	param.SampleRate = float64(DefaultSampleRate)
	param.FramesPerBuffer = 0
	param.Flags = portaudio.ClipOff
	return &PortAudio{
		InputDevice: device,
		params:      param,
	}
}

func (pa *PortAudio) Stream(audioConfig Audio[portaudio.StreamParameters]) (AudioStream, error) {
	return audioConfig.OpenStream(pa.params)
}
