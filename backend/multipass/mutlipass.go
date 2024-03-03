package multipass

import (
	"dockman/backend/cmd"
	"dockman/backend/common"
	"encoding/json"
	"time"
)

type MultipassVM struct {
	ID      string
	Created time.Time
}

func (m *MultipassVM) GetType() common.VMType {
	return common.VMTypeMultipass
}

func (m *MultipassVM) GetID() string {
	return m.ID
}

func (m *MultipassVM) GetCreated() time.Time {
	return m.Created
}

func (m *MultipassVM) GetName() string {
	return m.ID
}

func GetMultipassVMs() []common.VM {
	cmd := cmd.Command{
		Cmd:  "multipass",
		Args: []string{"list", "--format", "json"},
	}

	out, err := cmd.Exec(cmd)
	if err != nil {
		panic(err)
	}

	return parseMultipassVMs(out)
}

type multipassVMInternal struct {
	Name    string   `json:"name"`
	State   string   `json:"state"`
	Ipv4    []string `json:"ipv4"`
	Release string   `json:"release"`
}

type multipassListResponseInternal struct {
	List []multipassVMInternal `json:"list"`
}

func parseMultipassVMs(out string) []common.VM {
	vms := multipassListResponseInternal{}

	err := json.Unmarshal([]byte(out), &vms)
	if err != nil {
		panic(err)
	}

	ret := make([]common.VM, 0)
	for _, vm := range vms.List {
		// cast multipassVM to VM
		ret = append(ret, &MultipassVM{
			ID:      vm.Name,
			Created: time.Now(),
		})
	}

	return ret
}
