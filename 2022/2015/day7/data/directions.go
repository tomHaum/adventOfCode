package data

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

type Circuit struct {
	Gates     map[Wire]Gate
	GateCache map[Wire]Signal
}

type Operation func(circuit *Circuit, input1 Input, input2 Input) Signal

type Gate struct {
	Input1        Input
	Input2        Input
	Operation     Operation
	operationName string
	Output        Wire
}

// Input either Wire or Signal is populated, not both
type Input struct {
	Wire   Wire
	Signal Signal
}
type Wire string
type Signal uint16

const noSignal Signal = 0
const noWire Wire = ""

var noInput = Input{noWire, noSignal}

// GetData true for up false for down
func GetData(fileName string) Circuit {
	file, err := os.Open(fileName)
	Check(err)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	circuit := Circuit{make(map[Wire]Gate), make(map[Wire]Signal)}
	notRegex := regexp.MustCompile(`^NOT ((\d+)|([a-z]+)) -> (\w+)$`)
	signalRegex := regexp.MustCompile(`^((\d+)|([a-z]+)) -> (\w+)$`)
	andRegex := regexp.MustCompile(`^((\d+)|([a-z]+)) AND ((\d+)|([a-z]+)) -> (\w+)$`)
	orRegex := regexp.MustCompile(`^((\d+)|([a-z]+)) OR ((\d+)|([a-z]+)) -> (\w+)$`)
	leftShiftRegex := regexp.MustCompile(`^((\d+)|([a-z]+)) LSHIFT ((\d+)|([a-z]+)) -> (\w+)$`)
	rightShiftRegex := regexp.MustCompile(`^((\d+)|([a-z]+)) RSHIFT ((\d+)|([a-z]+)) -> (\w+)$`)

	for scanner.Scan() {
		line := scanner.Text()

		matches := notRegex.FindStringSubmatch(line)
		if matches != nil {
			input1 := parseInput(matches[2:4])
			input2 := noInput
			output := Wire(matches[4])

			gate := Gate{input1, input2, not, "not", output}

			circuit.Gates[output] = gate
			continue
		}

		matches = signalRegex.FindStringSubmatch(line)
		if matches != nil {
			input1 := parseInput(matches[2:4])
			input2 := noInput
			output := Wire(matches[4])

			gate := Gate{input1, input2, signal, "signal", output}

			circuit.Gates[output] = gate
			continue
		}

		matches = andRegex.FindStringSubmatch(line)
		if matches != nil {
			input1 := parseInput(matches[2:4])
			input2 := parseInput(matches[5:7])
			output := Wire(matches[7])

			gate := Gate{input1, input2, and, "and", output}

			circuit.Gates[output] = gate
			continue
		}

		matches = orRegex.FindStringSubmatch(line)
		if matches != nil {
			input1 := parseInput(matches[2:4])
			input2 := parseInput(matches[5:7])
			output := Wire(matches[7])

			gate := Gate{input1, input2, or, "or", output}

			circuit.Gates[output] = gate
			continue
		}

		matches = leftShiftRegex.FindStringSubmatch(line)
		if matches != nil {
			input1 := parseInput(matches[2:4])
			input2 := parseInput(matches[5:7])
			output := Wire(matches[7])

			gate := Gate{input1, input2, leftShift, "leftShift", output}

			circuit.Gates[output] = gate
			continue
		}

		matches = rightShiftRegex.FindStringSubmatch(line)
		if matches != nil {
			input1 := parseInput(matches[2:4])
			input2 := parseInput(matches[5:7])
			output := Wire(matches[7])

			gate := Gate{input1, input2, rightShift, "rightShift", output}

			circuit.Gates[output] = gate
			continue
		}
	}

	return circuit
}
func parseInput(strs []string) Input {
	if len(strs) != 2 {
		panic("invalid input")
	}
	if strs[0] != "" {
		signal, err := strconv.Atoi(strs[0])
		Check(err)
		return Input{Signal: Signal(uint16(signal)), Wire: noWire}
	}
	if strs[1] != "" {
		return Input{Signal: noSignal, Wire: Wire(strs[1])}
	}
	panic("invalid input")
}
func signal(circuit *Circuit, input1 Input, input2 Input) Signal {
	return eval(circuit, input1)
}
func not(circuit *Circuit, input1 Input, input2 Input) Signal {
	return ^eval(circuit, input1)
}

func and(circuit *Circuit, input1 Input, input2 Input) Signal {
	return eval(circuit, input1) & eval(circuit, input2)
}
func or(circuit *Circuit, input1 Input, input2 Input) Signal {
	return eval(circuit, input1) | eval(circuit, input2)
}
func leftShift(circuit *Circuit, input1 Input, input2 Input) Signal {
	return eval(circuit, input1) << eval(circuit, input2)
}

func rightShift(circuit *Circuit, input1 Input, input2 Input) Signal {
	return eval(circuit, input1) >> eval(circuit, input2)
}

func eval(circuit *Circuit, input Input) Signal {
	if input.Wire == noWire {
		return input.Signal
	}

	return (*circuit).Gates[input.Wire].Evaluate(circuit)
}

var i = 0

func (gate Gate) Evaluate(circuit *Circuit) Signal {
	i++
	localCiruit := *circuit

	if val, ok := localCiruit.GateCache[gate.Output]; ok {
		return val
	}
	valueToCache := gate.Operation(circuit, gate.Input1, gate.Input2)
	localCiruit.GateCache[gate.Output] = valueToCache
	return valueToCache
}
func Check(e error) {
	if e != nil {
		panic(e)
	}
}
