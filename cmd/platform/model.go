package platform

// Platforms - data structure for platforms module types configuration
type Platforms struct {
	Platforms []Modules `yaml:"platform"`
}

// Modules -  modules list
type Modules struct {
	Name    string       `yaml:"name"`
	Modules []Properties `yaml:"modules"`
}

// Properties - properties list
type Properties struct {
	NativeType   string `yaml:"native-type"`
	PlatformType string `yaml:"platform-type"`
}
