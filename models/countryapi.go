package models

// CountryResponse represents the response from a country information API,
// including an array of country details and any errors encountered.
type CountryResponse struct {
	ArrayResponse
	Error CountryApiError
}

// HasError checks if the CountryResponse contains any error.
func (i *CountryResponse) HasError() bool {
	return i.Error.Code != 0 || i.Error.Message != ""
}

// ArrayResponse is a slice of CountryResponseElement, representing multiple country responses.
type ArrayResponse []CountryResponseElement

// CountryResponseElement holds detailed information about a specific country.
type CountryResponseElement struct {
	Name         Name                   `json:"name"`
	TLD          []string               `json:"tld"`
	Cca2         string                 `json:"cca2"`
	Ccn3         string                 `json:"ccn3"`
	Cca3         string                 `json:"cca3"`
	Cioc         string                 `json:"cioc"`
	Independent  bool                   `json:"independent"`
	Status       string                 `json:"status"`
	UnMember     bool                   `json:"unMember"`
	Currencies   map[string]Currency    `json:"currencies"`
	Idd          Idd                    `json:"idd"`
	Capital      []string               `json:"capital"`
	AltSpellings []string               `json:"altSpellings"`
	Region       string                 `json:"region"`
	Subregion    string                 `json:"subregion"`
	Languages    Languages              `json:"languages"`
	Translations map[string]Translation `json:"translations"`
	Latlng       []float64              `json:"latlng"`
	Landlocked   bool                   `json:"landlocked"`
	Borders      []string               `json:"borders"`
	Area         float64                `json:"area"`
	Demonyms     Demonyms               `json:"demonyms"`
	Flag         string                 `json:"flag"`
	Maps         Maps                   `json:"maps"`
	Population   int64                  `json:"population"`
	Gini         Gini                   `json:"gini"`
	Fifa         string                 `json:"fifa"`
	Car          Car                    `json:"car"`
	Timezones    []string               `json:"timezones"`
	Continents   []string               `json:"continents"`
	Flags        Flags                  `json:"flags"`
	CoatOfArms   CoatOfArms             `json:"coatOfArms"`
	StartOfWeek  string                 `json:"startOfWeek"`
	CapitalInfo  CapitalInfo            `json:"capitalInfo"`
	Error        CountryApiError
}

// CapitalInfo holds the geographic coordinates of the capital city.
type CapitalInfo struct {
	Latlng []float64 `json:"latlng"`
}

// Car contains information about the country's vehicle regulations.
type Car struct {
	Signs []string `json:"signs"`
	Side  string   `json:"side"`
}

// CoatOfArms holds the URLs for the country's coat of arms in different formats.
type CoatOfArms struct {
	PNG string `json:"png"`
	SVG string `json:"svg"`
}

// Currency represents the details of a currency, including its name and symbol.
type Currency struct {
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

// Demonyms contains English and French demonyms for the country.
type Demonyms struct {
	Eng Eng `json:"eng"`
	Fra Eng `json:"fra"`
}

// Eng holds male and female forms of the demonym.
type Eng struct {
	F string `json:"f"`
	M string `json:"m"`
}

// Flags holds information about the country's flags in different formats.
type Flags struct {
	PNG string `json:"png"`
	SVG string `json:"svg"`
	Alt string `json:"alt"`
}

// Gini holds the Gini index for income inequality for a specific year.
type Gini struct {
	The2019 float64 `json:"2019"`
}

// Idd contains the international dialing code and suffixes for the country.
type Idd struct {
	Root     string   `json:"root"`
	Suffixes []string `json:"suffixes"`
}

type Languages struct {
	SPA string `json:"spa"`
}

// Languages holds the languages spoken in the country.
type Maps struct {
	GoogleMaps     string `json:"googleMaps"`
	OpenStreetMaps string `json:"openStreetMaps"`
}

// Name holds the common, official, and native names of the country.
type Name struct {
	Common     string     `json:"common"`
	Official   string     `json:"official"`
	NativeName NativeName `json:"nativeName"`
}

// NativeName holds the native name in various languages.
type NativeName struct {
	SPA Translation `json:"spa"`
}

// Translation holds the official and common translations of the country name.
type Translation struct {
	Official string `json:"official"`
	Common   string `json:"common"`
}
