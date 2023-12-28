package connector

import "errors"

type Connector interface {
	GetSignal(input_signal string, sender string) (string, []string, error)
}

type Flip_Flop struct {
	On           bool
	Destinations []string
}

type Conjunction struct {
	Brain        map[string]string
	Destinations []string
}

type Broadcaster struct {
	Destinations []string
}

type Output struct {
}

func (f *Flip_Flop) GetSignal(input_signal string, _ string) (string, []string, error) {
	if input_signal == "hi" {
		return "None", []string{" "}, errors.New("Flip-Flop got High Signal")
	}

	if input_signal == "lo" {
		if f.On {
			f.On = false
			return "lo", f.Destinations, nil
		} else {
			f.On = true
			return "hi", f.Destinations, nil
		}
	}

	return "None", []string{" "}, errors.New("Flip-Flop got invalid signal")
}

func (c *Conjunction) GetSignal(input_signal string, sender string) (string, []string, error) {
	c.Brain[sender] = input_signal

	for _, value := range c.Brain {
		if value == "lo" {
			return "hi", c.Destinations, nil
		}
	}

	return "lo", c.Destinations, nil
}

func (c *Broadcaster) GetSignal(input_signal string, _ string) (string, []string, error) {
	return input_signal, c.Destinations, nil
}

func (o *Output) GetSignal(input_signal string, _ string) (string, []string, error) {
	return input_signal, []string{" "}, errors.New("Output got signal")
}
