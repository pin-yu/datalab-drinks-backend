package entities

// Series represents series of cama's drink
type Series struct {
	Series string `yaml:"series" json:"series"`
	Items  []Item `yaml:"items" json:"items"`
}
