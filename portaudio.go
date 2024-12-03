package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gordonklaus/portaudio"
)

func main() {
	// Initialize PortAudio
	if err := portaudio.Initialize(); err != nil {
		log.Fatal("Error initializing PortAudio:", err)
	}
	defer portaudio.Terminate()

	// Define the audio stream parameters
	const sampleRate = 44100
	const framesPerBuffer = 64

	// Create a buffer to hold audio data
	buffer := make([]float32, framesPerBuffer)

	// Open a stream for input and output
	stream, err := portaudio.OpenStream(portaudio.StreamParameters{
		Device:   portaudio.DefaultInputDevice(),
		Channels: 1,
		SampleRate: sampleRate,
		FramesPerBuffer: framesPerBuffer,
	}, portaudio.StreamParameters{
		Device:   portaudio.DefaultOutputDevice(),
		Channels: 1,
		SampleRate: sampleRate,
		FramesPerBuffer: framesPerBuffer,
	}, buffer)
	if err != nil {
		log.Fatal("Error opening stream:", err)
	}
	defer stream.Close()

	// Start the audio stream
	if err := stream.Start(); err != nil {
		log.Fatal("Error starting stream:", err)
	}
	defer stream.Stop()

	fmt.Println("Recording and playing audio. Press Ctrl+C to stop.")

	// Continuously read from the input and write to the output
	for {
		if err := stream.Read(); err != nil {
			log.Fatal("Error reading from stream:", err)
		}

		// Output the audio data
		if err := stream.Write(); err != nil {
			log.Fatal("Error writing to stream:", err)
		}

		// Sleep briefly to avoid busy-waiting
		time.Sleep(10 * time.Millisecond)
	}
}