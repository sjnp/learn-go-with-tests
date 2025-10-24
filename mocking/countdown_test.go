package main

import (
	"bytes"
	"slices"
	"testing"
	"time"
)

type SpySleeper struct{
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

type SpyCountdownOperations struct{
	Calls []string
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

type SpyTime struct{
	durationSlept time.Duration
}

func (s *SpyTime) SetDurationSlept(duration time.Duration) {
	s.durationSlept = duration
}

const write = "write"
const sleep = "sleep"

func TestCountdown(t *testing.T) {
	t.Run("sleep with no order checking", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spySleeper := &SpySleeper{}

		Countdown(buffer, spySleeper)

		got := buffer.String()
		expected := "3\n2\n1\nGo!"

		if got != expected {
			t.Errorf("got %q but expected %q", got, expected)
		}

		expectedCalls := 3
		if spySleeper.Calls != 3 {
			t.Errorf("called sleeper %d times but expected %d times", spySleeper.Calls, expectedCalls)
		}
	})

	t.Run("sleep before print", func(t *testing.T) {
		spySleepPrinter := &SpyCountdownOperations{}
		Countdown(spySleepPrinter, spySleepPrinter)

		expected := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !slices.Equal(spySleepPrinter.Calls, expected) {
			t.Errorf("got %v but expected %v", spySleepPrinter.Calls, expected)
		}
		
	})
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{
		duration: sleepTime,
		sleep:    spyTime.SetDurationSlept,
	}

	sleeper.Sleep()
	if spyTime.durationSlept != sleepTime {
		t.Errorf("got %v but expected %v", spyTime.durationSlept, sleepTime)
	}
}
