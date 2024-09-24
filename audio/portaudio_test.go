package audio

import (
	"testing"

	"github.com/gordonklaus/portaudio"
	"github.com/stretchr/testify/assert"
)

func TestNewPortAudio(t *testing.T) {
	assert.Nil(t, portaudio.Initialize())
	defer portaudio.Terminate()
	device, err := portaudio.DefaultOutputDevice()
	assert.NoError(t, err)

	pa := NewPortAudio(device)
	_, err = pa.Stream(PortAudioStream{})
	assert.NoError(t, err)

}
