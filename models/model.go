package models

//(conName string, image string, environment []string, guestPort string, hostPort string, netResp types.NetworkCreateResponse) error {
type Container struct {
	ConName 		string   `json:"conname"`
	Image			string   `json:"images"`
	Environment		[]string `json:"environment"`
	GuestPort		string   `json:"guestport"`
	HostPort		string   `json:"hostport"`
	Driver			string	 `json:"driver"`
}

//(conName []string, images []string, environment [][]string, guestPorts []string, hostPort []string, netResp types.NetworkCreateResponse)
type Compose struct {
	ConName			[]string   `json:"conname"`
	Images			[]string   `json:"images"`
	Environment		[][]string `json:"environment"`
	GuestPorts		[]string   `json:"guestports"`
	HostPorts		[]string   `json:"hostports"`
	ComposeName		string	   `json:"composename`
	Driver			string     `json:"driver"`
}
