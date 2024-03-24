package wrapper

import (
	"encoding/json"

	"github.com/sharki13/dockman/backend/cmd"
)

type MultipassWrapper struct{}

func MakeMultipassWrapper() Wrapper {
	return &MultipassWrapper{}
}

func (w *MultipassWrapper) GetType() WarpperType {
	return MultipassWrapperType
}

var multipassListCommand = cmd.Command{
	Cmd:  "multipass",
	Args: []string{"list", "--format", "json"},
}

type multipassList struct {
	List []struct {
		Name    string   `json:"name"`
		State   string   `json:"state"`
		Release string   `json:"release"`
		Ipv4    []string `json:"ipv4"`
	} `json:"list"`
}

func getUniversalState(state string) StateType {
	if state == "Running" {
		return StateRunning
	}
	return StateStopped
}

func (w *MultipassWrapper) GetInstances() ([]Instance, error) {
	out, ret := multipassListCommand.Exec()
	if ret != nil {
		return nil, ret
	}

	var list multipassList

	err := json.Unmarshal([]byte(out), &list)
	if err != nil {
		return nil, err
	}

	instances := []Instance{}
	for _, i := range list.List {
		instances = append(instances, Instance{
			ID:    i.Name,
			Name:  i.Name,
			State: getUniversalState(i.State),
			Type:  MultipassInstance,
		})
	}

	return instances, nil
}
