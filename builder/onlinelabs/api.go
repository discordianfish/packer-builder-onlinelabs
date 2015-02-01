package onlinelabs

type Server struct {
	Bootscript      *Bootscript        `json:"bootscript,omitempty"`
	DynamicPublicIP bool               `json:"dynamic_public_ip"`
	ID              string             `json:"id"`
	Image           *Image             `json:"image"`
	Name            string             `json:"name"`
	Organization    string             `json:"organization"`
	PrivateIP       *NullString        `json:"private_ip"`
	PublicIP        *NullString        `json:"public_ip"`
	State           string             `json:"state"`
	Tags            []string           `json:"tags"`
	Volumes         map[string]*Volume `json:"volumes"`
}

type Image struct {
	DefaultBootscript *Bootscript `json:"default_bootscript,omitempty"`
	Arch              string      `json:"arch,omitempty"`
	CreationDate      string      `json:"creation_date,omitempty"`
	ExtraVolumes      string      `json:"extra_volumes,omitempty"` // is this a bug?? e.g.: "[]"
	FromImage         *NullString `json:"from_image,omitempty"`
	FromServer        *NullString `json:"from_server,omitempty"`
	ID                string      `json:"id"`
	MarketplaceKey    *NullString `json:"marketplace_key,omitempty"`
	ModificationDate  string      `json:"modification_date,omitempty"`
	Name              string      `json:"name"`
	Organization      string      `json:"organization,omitempty"`
	Public            bool        `json:"public"`
	RootVolume        *Volume     `json:"root_volume,omitempty"`
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

type Bootscript struct {
	Kernel       *Kernel      `json:"kernel"`
	Title        string       `json:"title"`
	Public       bool         `json:"public"`
	Initrd       *Initrd      `json:"initrd"`
	BootCmdArgs  *BootCmdArgs `json:"bootcmdargs"`
	Organization string       `json:"organization"`
	ID           string       `json:"id"`
}

type Kernel struct {
	Dtb   string `json:"dtb"`
	Path  string `json:"path"`
	ID    string `json:"id"`
	Title string `json:"title"`
}

type Initrd struct {
	Path  string `json:"path"`
	ID    string `json:"id"`
	Title string `json:"title"`
}

type BootCmdArgs struct {
	ID    string `json:"id"`
	Value string `json:"value"`
}

type createServerParams struct {
	Organization string             `json:"organization"`
	Name         string             `json:"name"`
	Image        string             `json:"image"`
	Tags         []string           `json:"tags,omitempty"`
	Volumes      map[string]*Volume `json:"volumes"`
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
