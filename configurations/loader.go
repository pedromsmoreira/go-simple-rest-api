package configurations

// Loader - Defines the common method used to Load a Config file in whichever format you want
type Loader interface {
	// TODO: try to generalize this method, accept a file path and arbitrary filename, it would be fun
	Load() (Configuration, error)
}
