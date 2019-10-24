package dragonfly

// Config is the configuration of a Dragonfly server. It holds settings that affect different aspects of the
// server, such as its name and maximum players.
type Config struct {
	// Network holds settings related to network aspects of the server.
	Network struct {
		// Address is the address on which the server should listen. Players may connect to this address in
		// order to join.
		Address string
		// EnableEndpoints enables HTTP endpoints which may be hit to obtain information about the server.
		// When set to true, the endpoints will be served on the Address above.
		EnableEndpoints bool
	}
	Server struct {
		// Name is the name of the server as it shows up in the server list.
		Name string
		// MaximumPlayers is the maximum amount of players allowed to join the server at the same time. If set
		// to 0, the amount of maximum players will grow every time a player joins.
		MaximumPlayers int
		// ShutdownMessage is the message shown to players when the server shuts down. If empty, players will
		// be directed to the menu screen right away.
		ShutdownMessage string
	}
	World struct {
		// Name is the name of the world that the server holds. A world with this name will be loaded and
		// the name will be displayed at the top of the player list in the in-game pause menu.
		Name string
		// Folder is the folder that the data of the world resides in.
		Folder string
		// MaximumChunkRadius is the maximum chunk radius that players may set in their settings. If they try
		// to set it above this number, it will be capped and set to the max.
		MaximumChunkRadius int
	}
}

// DefaultConfig returns a configuration with the default values filled out.
func DefaultConfig() Config {
	c := Config{}
	c.Network.Address = ":19132"
	c.Network.EnableEndpoints = true
	c.Server.Name = "Dragonfly Server"
	c.World.Name = "World"
	c.World.Folder = "world"
	c.World.MaximumChunkRadius = 32
	return c
}
