package onlinelabs

type Server struct {
	Bootscript      *NullString        `json:"bootscript"`
	DynamicPublicIP bool               `json:"dynamic_public_ip"`
	ID              string             `json:"id"`
	Image           *Image             `json:"image"`
	Name            string             `json:"name"`
	Organization    string             `json:"organization"`
	PrivateIP       *NullString        `json:"private_ip"`
	PublicIP        *NullString        `json:"public_ip"`
	Running         bool               `json:"running"`
	Tags            []string           `json:"tags"`
	Volumes         map[string]*Volume `json:"volumes"`
}

type Image struct {
	Arch             string      `json:"arch,omitempty"`
	CreationDate     string      `json:"creation_date,omitempty"`
	ExtraVolumes     []*Volume   `json:"extra_volumes,omitempty"`
	FromImage        *NullString `json:"from_image,omitempty"`
	FromServer       *NullString `json:"from_server,omitempty"`
	ID               string      `json:"id"`
	MarketplaceKey   *NullString `json:"marketplace_key,omitempty"`
	ModificationDate string      `json:"modification_date,omitempty"`
	Name             string      `json:"name"`
	Organization     string      `json:"organization,omitempty"`
	Public           bool        `json:"public"`
	RootVolume       *Volume     `json:"root_volume,omitempty"`
}

type Volume struct {
	ExportURI    *NullString        `json:"export_uri"`
	ID           string             `json:"id"`
	Name         string             `json:"name"`
	Organization string             `json:"organization"`
	Server       *AbbreviatedServer `json:"server"`
	Size         uint64             `json:"size"`
	VolumeType   string             `json:"volume_type"`
}

type AbbreviatedServer struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Snapshot struct {
	BaseVolume   *Volume `json:"base_volume"`
	CreationDate string  `json:"creation_date"`
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	Organization string  `json:"organization"`
	Size         uint64  `json:"size"`
	State        string  `json:"state"`
	VolumeType   string  `json:"volume_type"`
}

type NullString struct {
	Value string
}

func (ns *NullString) String() string {
	return ns.Value
}

func (ns *NullString) UnmarshalJSON(j []byte) error {
	if string(j) == "null" {
		ns.Value = ""
	}
	ns.Value = string(j)
	return nil
}

func (ns *NullString) MarshalJSON() ([]byte, error) {
	if ns.Value == "" {
		return []byte("null"), nil
	}

	return []byte(ns.Value), nil
}
