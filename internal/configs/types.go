package types

type Servers struct {
	PlayersCount        int      `json:"playersCount"`
	MaxPlayersCount     int      `json:"maxPlayersCount"`
	QueueCount          int      `json:"queueCount"`
	Passworded          bool     `json:"passworded"`
	Port                int      `json:"port"`
	Language            string   `json:"language"`
	UseEarlyAuth        bool     `json:"useEarlyAuth"`
	EarlyAuthURL        string   `json:"earlyAuthUrl"`
	UseCDN              bool     `json:"useCdn"`
	CDNURL              string   `json:"cdnUrl"`
	UseVoiceChat        bool     `json:"useVoiceChat"`
	Version             string   `json:"version"`
	Branch              string   `json:"branch"`
	Available           bool     `json:"available"`
	Banned              bool     `json:"banned"`
	Name                string   `json:"name"`
	PublicID            string   `json:"publicId"`
	VanityURL           string   `json:"vanityUrl"`
	Website             string   `json:"website"`
	GameMode            string   `json:"gameMode"`
	Description         string   `json:"description"`
	Tags                []string `json:"tags"`
	LastTimeUpdate      string   `json:"lastTimeUpdate"`
	Verified            bool     `json:"verified"`
	Promoted            bool     `json:"promoted"`
	Visible             bool     `json:"visible"`
	HasCustomSkin       bool     `json:"hasCustomSkin"`
	BannerURL           string   `json:"bannerUrl"`
	Address             string   `json:"address"`
	Group               any      `json:"group"`
	MasterlistIconURL   *string  `json:"masterlist_icon_url"`
	MasterlistBannerURL *string  `json:"masterlist_banner_url"`
}
