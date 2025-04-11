package in

type Endpoint struct {
	Value       string `json:"Value"`
	Overridable bool   `json:"Overridable"`
}

type AllowedIPs struct {
	Value       []string `json:"Value"`
	Overridable bool     `json:"Overridable"`
}

type ConfigResponse struct {
	Identifier          string     `json:"Identifier"`
	DisplayName         string     `json:"DisplayName"`
	UserIdentifier      string     `json:"UserIdentifier"`
	InterfaceIdentifier string     `json:"InterfaceIdentifier"`
	Disabled            bool       `json:"Disabled"`
	DisabledReason      string     `json:"DisabledReason"`
	Notes               string     `json:"Notes"`
	Endpoint            Endpoint   `json:"Endpoint"`
	EndpointPublicKey   Endpoint   `json:"EndpointPublicKey"`
	AllowedIPs          AllowedIPs `json:"AllowedIPs"`
	ExtraAllowedIPs     []string   `json:"ExtraAllowedIPs"`
	PresharedKey        string     `json:"PresharedKey"`
	PersistentKeepalive Endpoint   `json:"PersistentKeepalive"`
	PrivateKey          string     `json:"PrivateKey"`
	PublicKey           string     `json:"PublicKey"`
	Mode                string     `json:"Mode"`
	Addresses           []string   `json:"Addresses"`
	CheckAliveAddress   string     `json:"CheckAliveAddress"`
	Dns                 AllowedIPs `json:"Dns"`
	DnsSearch           AllowedIPs `json:"DnsSearch"`
	Mtu                 Endpoint   `json:"Mtu"`
	FirewallMark        Endpoint   `json:"FirewallMark"`
	RoutingTable        Endpoint   `json:"RoutingTable"`
	PreUp               Endpoint   `json:"PreUp"`
	PostUp              Endpoint   `json:"PostUp"`
	PreDown             Endpoint   `json:"PreDown"`
	PostDown            Endpoint   `json:"PostDown"`
}
