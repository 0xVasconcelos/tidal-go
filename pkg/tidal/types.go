package tidal

type Client struct {
	Token        string
	SoundQuality string
	Username     string
	Password     string
	User         *User
}

type User struct {
	UserID      int    `json:"userId"`
	SessionID   string `json:"sessionId"`
	CountryCode string `json:"countryCode"`
}

type Media struct {
	Media         []byte
	EncryptionKey []byte
	Key           []byte
	Nonce         []byte
}

type LoginError struct {
	Status      int    `json:"status"`
	SubStatus   int    `json:"subStatus"`
	UserMessage string `json:"userMessage"`
}

type Search struct {
	Artists struct {
		Limit              int       `json:"limit"`
		Offset             int       `json:"offset"`
		TotalNumberOfItems int       `json:"totalNumberOfItems"`
		Items              []*Artist `json:"items"`
	} `json:"artists"`
	Albums struct {
		Limit              int      `json:"limit"`
		Offset             int      `json:"offset"`
		TotalNumberOfItems int      `json:"totalNumberOfItems"`
		Items              []*Album `json:"items"`
	} `json:"albums"`
	Playlists struct {
		Limit              int         `json:"limit"`
		Offset             int         `json:"offset"`
		TotalNumberOfItems int         `json:"totalNumberOfItems"`
		Items              []*Playlist `json:"items"`
	} `json:"playlists"`
	Tracks struct {
		Limit              int      `json:"limit"`
		Offset             int      `json:"offset"`
		TotalNumberOfItems int      `json:"totalNumberOfItems"`
		Items              []*Track `json:"items"`
	} `json:"tracks"`
}

type Artist struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	ArtistTypes []string `json:"artistTypes"`
	Picture     string   `json:"picture"`
	Popularity  int      `json:"popularity"`
}

type Album struct {
	ID                   int      `json:"id"`
	Title                string   `json:"title"`
	Duration             int      `json:"duration"`
	StreamReady          bool     `json:"streamReady"`
	StreamStartDate      string   `json:"streamStartDate"`
	AllowStreaming       bool     `json:"allowStreaming"`
	PremiumStreamingOnly bool     `json:"premiumStreamingOnly"`
	NumberOfTracks       int      `json:"numberOfTracks"`
	NumberOfVideos       int      `json:"numberOfVideos"`
	NumberOfVolumes      int      `json:"numberOfVolumes"`
	ReleaseDate          string   `json:"releaseDate"`
	Copyright            string   `json:"copyright"`
	Type                 string   `json:"type"`
	URL                  string   `json:"url"`
	Cover                string   `json:"cover"`
	Explicit             bool     `json:"explicit"`
	Upc                  string   `json:"upc"`
	Popularity           int      `json:"popularity"`
	AudioQuality         string   `json:"audioQuality"`
	AudioModes           []string `json:"audioModes"`
	Artists              []*Artist
}

type Playlist struct {
	UUID           string `json:"uuid"`
	Title          string `json:"title"`
	NumberOfTracks int    `json:"numberOfTracks"`
	NumberOfVideos int    `json:"numberOfVideos"`
	Creator        struct {
		ID         int    `json:"id"`
		Name       string `json:"name"`
		Popularity int    `json:"popularity"`
	} `json:"creator"`
}

type Track struct {
	ID                   int       `json:"id"`
	Title                string    `json:"title"`
	Duration             int       `json:"duration"`
	ReplayGain           float64   `json:"replayGain"`
	Peak                 float64   `json:"peak"`
	AllowStreaming       bool      `json:"allowStreaming"`
	StreamReady          bool      `json:"streamReady"`
	StreamStartDate      string    `json:"streamStartDate"`
	PremiumStreamingOnly bool      `json:"premiumStreamingOnly"`
	TrackNumber          int       `json:"trackNumber"`
	VolumeNumber         int       `json:"volumeNumber"`
	Popularity           int       `json:"popularity"`
	Copyright            string    `json:"copyright"`
	URL                  string    `json:"url"`
	Isrc                 string    `json:"isrc"`
	Editable             bool      `json:"editable"`
	Explicit             bool      `json:"explicit"`
	AudioQuality         string    `json:"audioQuality"`
	AudioModes           []string  `json:"audioModes"`
	Artists              []*Artist `json:"artists"`
	Album                *Album    `json:"album"`
}

type Stream struct {
	URL                   string `json:"url"`
	TrackID               int    `json:"trackId"`
	PlayTimeLeftInMinutes int    `json:"playTimeLeftInMinutes"`
	SoundQuality          string `json:"soundQuality"`
	EncryptionKey         string `json:"encryptionKey"`
	Codec                 string `json:"codec"`
}
