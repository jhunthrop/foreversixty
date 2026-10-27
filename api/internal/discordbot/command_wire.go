package discordbot

import "encoding/json"

// guildContext and guildInstall are the one context and integration type the
// bot's commands use: server channels, installed to a server.
const (
	guildContext = 0
	guildInstall = 0
)

// Choice is one fixed answer of a command option.
type Choice struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// CommandOption is one argument of a slash command; Type is Discord's option
// type number.
type CommandOption struct {
	Name         string   `json:"name"`
	Description  string   `json:"description"`
	Type         int      `json:"type"`
	Required     bool     `json:"required,omitempty"`
	Autocomplete bool     `json:"autocomplete,omitempty"`
	Choices      []Choice `json:"choices,omitempty"`
}

// Command is a slash command. Released says whether it is registered
// globally yet; it is bot-side state and never goes to Discord.
type Command struct {
	Name        string
	Description string
	Options     []CommandOption
	Released    bool
}

// MarshalJSON writes Discord's application-command object.
func (c Command) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Name             string          `json:"name"`
		Description      string          `json:"description"`
		Options          []CommandOption `json:"options,omitempty"`
		Contexts         []int           `json:"contexts"`
		IntegrationTypes []int           `json:"integration_types"`
	}{c.Name, c.Description, c.Options, []int{guildContext}, []int{guildInstall}})
}
