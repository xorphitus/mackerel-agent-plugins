package mpunicorn

import "testing"

type TestWorkersMemoryPipedCommands struct{}

func (r TestWorkersMemoryPipedCommands) Output(commands ...[]string) ([]byte, error) {
	return []byte("3321016320\n"), nil
}

func TestWorkersMemory(t *testing.T) {
	pipedCommands = TestWorkersMemoryPipedCommands{}
	expectedMemory := "3321016320"
	m, _ := workersMemory()
	if m != expectedMemory {
		t.Errorf("workersMemory: expected %s but got %s", expectedMemory, m)
	}
}

type TestMasterMemoryPipedCommands struct{}

func (r TestMasterMemoryPipedCommands) Output(commands ...[]string) ([]byte, error) {
	return []byte("114636\n"), nil
}

func TestMasterMemory(t *testing.T) {
	pipedCommands = TestMasterMemoryPipedCommands{}
	expectedMemory := "114636"
	m, _ := masterMemory()
	if m != expectedMemory {
		t.Errorf("masterMemory: expected %s but got %s", expectedMemory, m)
	}
}

type TestWorkersMemoryAvgPipedCommands struct{}

func (r TestWorkersMemoryAvgPipedCommands) Output(commands ...[]string) ([]byte, error) {
	return []byte("204277504\n"), nil
}

func TestWorkersMemoryAvg(t *testing.T) {
	pipedCommands = TestWorkersMemoryAvgPipedCommands{}
	expectedMemory := "204277504"
	m, _ := workersMemoryAvg()
	if m != expectedMemory {
		t.Errorf("workersMemoryAvg: expected %s but got %s", expectedMemory, m)
	}
}

type TestWorkersMemoryMaxPipedCommands struct{}

func (r TestWorkersMemoryMaxPipedCommands) Output(commands ...[]string) ([]byte, error) {
	return []byte("4194304\n"), nil
}

func TestWorkersMemoryMax(t *testing.T) {
	pipedCommands = TestWorkersMemoryMaxPipedCommands{}
	expectedMemory := "4194304"
	m, _ := workersMemoryMax()
	if m != expectedMemory {
		t.Errorf("workersMemoryMax: expected %s but got %s", expectedMemory, m)
	}
}

type TestWorkersMemoryMinPipedCommands struct{}

func (r TestWorkersMemoryMinPipedCommands) Output(commands ...[]string) ([]byte, error) {
	return []byte("2097152\n"), nil
}

func TestWorkersMemoryMin(t *testing.T) {
	pipedCommands = TestWorkersMemoryMinPipedCommands{}
	expectedMemory := "2097152"
	m, _ := workersMemoryMin()
	if m != expectedMemory {
		t.Errorf("workersMemoryMin: expected %s but got %s", expectedMemory, m)
	}
}
