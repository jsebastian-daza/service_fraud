package models

// IpApiResponse represents the response from the IP API, containing information
// about the IP address, geographical location, and any errors encountered.
type IpApiResponse struct {
	IP             string   `json:"ip"`
	Type           string   `json:"type"`
	ContinentCode  string   `json:"continent_code"`
	ContinentName  string   `json:"continent_name"`
	CountryCode    string   `json:"country_code"`
	CountryName    string   `json:"country_name"`
	RegionCode     string   `json:"region_code"`
	RegionName     string   `json:"region_name"`
	City           string   `json:"city"`
	Zip            string   `json:"zip"`
	Latitude       float64  `json:"latitude"`
	Longitude      float64  `json:"longitude"`
	MSA            string   `json:"msa"`
	DMA            string   `json:"dma"`
	IPRoutingType  string   `json:"ip_routing_type"`
	ConnectionType string   `json:"connection_type"`
	Location       Location `json:"location"`
	Error          IpApiError
}

// HasError checks if the IpApiResponse contains any error.
func (i IpApiResponse) HasError() bool {
	return i.Error.Code != 0 || i.Error.Message != ""
}

// ContainsValidResponse checks if the response includes valid geographical data.
func (i *IpApiResponse) ContainsValidResponse() bool {
	return i.ContinentCode != "" && i.CountryCode != "" && i.CountryName != ""
}

// Location holds detailed geographical information related to the IP address.
type Location struct {
	GeonameID               int64      `json:"geoname_id"`
	Capital                 string     `json:"capital"`
	Languages               []Language `json:"languages"`
	CountryFlag             string     `json:"country_flag"`
	CountryFlagEmoji        string     `json:"country_flag_emoji"`
	CountryFlagEmojiUnicode string     `json:"country_flag_emoji_unicode"`
	CallingCode             string     `json:"calling_code"`
	IsEu                    bool       `json:"is_eu"`
}

// Language represents a language spoken in the country.
type Language struct {
	Code   string `json:"code"`
	Name   string `json:"name"`
	Native string `json:"native"`
}
