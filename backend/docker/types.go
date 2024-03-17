package docker

import "time"

type DockerContainer struct {
	ID              string          `json:"Id"`
	Created         time.Time       `json:"Created"`
	Path            string          `json:"Path"`
	Args            []string        `json:"Args"`
	State           State           `json:"State"`
	Image           string          `json:"Image"`
	ResolvConfPath  string          `json:"ResolvConfPath"`
	HostnamePath    string          `json:"HostnamePath"`
	HostsPath       string          `json:"HostsPath"`
	LogPath         string          `json:"LogPath"`
	Name            string          `json:"Name"`
	RestartCount    int             `json:"RestartCount"`
	Driver          string          `json:"Driver"`
	Platform        string          `json:"Platform"`
	MountLabel      string          `json:"MountLabel"`
	ProcessLabel    string          `json:"ProcessLabel"`
	AppArmorProfile string          `json:"AppArmorProfile"`
	ExecIDs         []string        `json:"ExecIDs"`
	HostConfig      HostConfig      `json:"HostConfig"`
	GraphDriver     GraphDriver     `json:"GraphDriver"`
	Mounts          []Mount         `json:"Mounts"`
	Config          Config          `json:"Config"`
	NetworkSettings NetworkSettings `json:"NetworkSettings"`
}

type State struct {
	Status     string    `json:"Status"`
	Running    bool      `json:"Running"`
	Paused     bool      `json:"Paused"`
	Restarting bool      `json:"Restarting"`
	OOMKilled  bool      `json:"OOMKilled"`
	Dead       bool      `json:"Dead"`
	Pid        int       `json:"Pid"`
	ExitCode   int       `json:"ExitCode"`
	Error      string    `json:"Error"`
	StartedAt  time.Time `json:"StartedAt"`
	FinishedAt time.Time `json:"FinishedAt"`
}

type HostConfig struct {
	Binds             []string               `json:"Binds"`
	ContainerIDFile   string                 `json:"ContainerIDFile"`
	LogConfig         LogConfig              `json:"LogConfig"`
	NetworkMode       string                 `json:"NetworkMode"`
	PortBindings      map[string]interface{} `json:"PortBindings"`
	RestartPolicy     RestartPolicy          `json:"RestartPolicy"`
	AutoRemove        bool                   `json:"AutoRemove"`
	VolumeDriver      string                 `json:"VolumeDriver"`
	VolumesFrom       []string               `json:"VolumesFrom"`
	ConsoleSize       []int                  `json:"ConsoleSize"`
	CapAdd            []string               `json:"CapAdd"`
	CapDrop           []string               `json:"CapDrop"`
	CgroupnsMode      string                 `json:"CgroupnsMode"`
	Dns               []string               `json:"Dns"`
	DnsOptions        []string               `json:"DnsOptions"`
	DnsSearch         []string               `json:"DnsSearch"`
	ExtraHosts        []string               `json:"ExtraHosts"`
	GroupAdd          []string               `json:"GroupAdd"`
	IpcMode           string                 `json:"IpcMode"`
	Cgroup            string                 `json:"Cgroup"`
	Links             []string               `json:"Links"`
	OomScoreAdj       int                    `json:"OomScoreAdj"`
	PidMode           string                 `json:"PidMode"`
	Privileged        bool                   `json:"Privileged"`
	PublishAllPorts   bool                   `json:"PublishAllPorts"`
	ReadonlyRootfs    bool                   `json:"ReadonlyRootfs"`
	SecurityOpt       []string               `json:"SecurityOpt"`
	UTSMode           string                 `json:"UTSMode"`
	UsernsMode        string                 `json:"UsernsMode"`
	ShmSize           int                    `json:"ShmSize"`
	Runtime           string                 `json:"Runtime"`
	Isolation         string                 `json:"Isolation"`
	CpuShares         int                    `json:"CpuShares"`
	Memory            int                    `json:"Memory"`
	NanoCpus          int                    `json:"NanoCpus"`
	CgroupParent      string                 `json:"CgroupParent"`
	BlkioWeight       int                    `json:"BlkioWeight"`
	BlkioWeightDevice []struct {
		Path   string `json:"Path"`
		Weight int    `json:"Weight"`
	} `json:"BlkioWeightDevice"`
	BlkioDeviceReadBps   []interface{} `json:"BlkioDeviceReadBps"`
	BlkioDeviceWriteBps  []interface{} `json:"BlkioDeviceWriteBps"`
	BlkioDeviceReadIOps  []interface{} `json:"BlkioDeviceReadIOps"`
	BlkioDeviceWriteIOps []interface{} `json:"BlkioDeviceWriteIOps"`
	CpuPeriod            int           `json:"CpuPeriod"`
	CpuQuota             int           `json:"CpuQuota"`
	CpuRealtimePeriod    int           `json:"CpuRealtimePeriod"`
	CpuRealtimeRuntime   int           `json:"CpuRealtimeRuntime"`
	CpusetCpus           string        `json:"CpusetCpus"`
	CpusetMems           string        `json:"CpusetMems"`
	Devices              []interface{} `json:"Devices"`
	DeviceCgroupRules    []interface{} `json:"DeviceCgroupRules"`
	DeviceRequests       []interface{} `json:"DeviceRequests"`
	MemoryReservation    int           `json:"MemoryReservation"`
	MemorySwap           int           `json:"MemorySwap"`
	MemorySwappiness     interface{}   `json:"MemorySwappiness"`
	OomKillDisable       bool          `json:"OomKillDisable"`
	PidsLimit            interface{}   `json:"PidsLimit"`
	Ulimits              []interface{} `json:"Ulimits"`
	CpuCount             int           `json:"CpuCount"`
	CpuPercent           int           `json:"CpuPercent"`
	IOMaximumIOps        int           `json:"IOMaximumIOps"`
	IOMaximumBandwidth   int           `json:"IOMaximumBandwidth"`
	MaskedPaths          []string      `json:"MaskedPaths"`
	ReadonlyPaths        []string      `json:"ReadonlyPaths"`
}

type LogConfig struct {
	Type   string   `json:"Type"`
	Config struct{} `json:"Config"`
}

type RestartPolicy struct {
	Name              string `json:"Name"`
	MaximumRetryCount int    `json:"MaximumRetryCount"`
}

type GraphDriver struct {
	Data Data   `json:"Data"`
	Name string `json:"Name"`
}

type Data struct {
	LowerDir  string `json:"LowerDir"`
	MergedDir string `json:"MergedDir"`
	UpperDir  string `json:"UpperDir"`
	WorkDir   string `json:"WorkDir"`
}

type Mount struct {
	Source      string `json:"Source"`
	Destination string `json:"Destination"`
	Mode        string `json:"Mode"`
	RW          bool   `json:"RW"`
	Propagation string `json:"Propagation"`
}

type Config struct {
	Hostname     string            `json:"Hostname"`
	Domainname   string            `json:"Domainname"`
	User         string            `json:"User"`
	AttachStdin  bool              `json:"AttachStdin"`
	AttachStdout bool              `json:"AttachStdout"`
	AttachStderr bool              `json:"AttachStderr"`
	Tty          bool              `json:"Tty"`
	OpenStdin    bool              `json:"OpenStdin"`
	StdinOnce    bool              `json:"StdinOnce"`
	Env          []string          `json:"Env"`
	Cmd          []string          `json:"Cmd"`
	Image        string            `json:"Image"`
	Volumes      map[string]string `json:"Volumes"`
	WorkingDir   string            `json:"WorkingDir"`
	Entrypoint   []string          `json:"Entrypoint"`
	OnBuild      []string          `json:"OnBuild"`
	Labels       map[string]string `json:"Labels"`
}

type NetworkSettings struct {
	Bridge                 string                 `json:"Bridge"`
	SandboxID              string                 `json:"SandboxID"`
	SandboxKey             string                 `json:"SandboxKey"`
	Ports                  map[string]interface{} `json:"Ports"`
	HairpinMode            bool                   `json:"HairpinMode"`
	LinkLocalIPv6Address   string                 `json:"LinkLocalIPv6Address"`
	LinkLocalIPv6PrefixLen int                    `json:"LinkLocalIPv6PrefixLen"`
	SecondaryIPAddresses   []interface{}          `json:"SecondaryIPAddresses"`
	SecondaryIPv6Addresses []interface{}          `json:"SecondaryIPv6Addresses"`
	EndpointID             string                 `json:"EndpointID"`
	Gateway                string                 `json:"Gateway"`
	GlobalIPv6Address      string                 `json:"GlobalIPv6Address"`
	GlobalIPv6PrefixLen    int                    `json:"GlobalIPv6PrefixLen"`
	IPAddress              string                 `json:"IPAddress"`
	IPPrefixLen            int                    `json:"IPPrefixLen"`
	IPv6Gateway            string                 `json:"IPv6Gateway"`
	MacAddress             string                 `json:"MacAddress"`
	Networks               map[string]Network     `json:"Networks"`
}

type Network struct {
	IPAMConfig          interface{} `json:"IPAMConfig"`
	Links               interface{} `json:"Links"`
	Aliases             interface{} `json:"Aliases"`
	MacAddress          string      `json:"MacAddress"`
	NetworkID           string      `json:"NetworkID"`
	EndpointID          string      `json:"EndpointID"`
	Gateway             string      `json:"Gateway"`
	IPAddress           string      `json:"IPAddress"`
	IPPrefixLen         int         `json:"IPPrefixLen"`
	IPv6Gateway         string      `json:"IPv6Gateway"`
	GlobalIPv6Address   string      `json:"GlobalIPv6Address"`
	GlobalIPv6PrefixLen int         `json:"GlobalIPv6PrefixLen"`
	DriverOpts          interface{} `json:"DriverOpts"`
	DNSNames            interface{} `json:"DNSNames"`
}
