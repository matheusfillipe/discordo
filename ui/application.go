package ui

import (
	"context"
	"strings"

	"github.com/ayntgl/discordo/config"
	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/gateway"
	"github.com/diamondburned/arikawa/v3/state"
)

// Application amalgamates the management of the terminal user interface, state (including underlying session) and configuration.
type Application struct {
	config *config.Config
	state  *state.State
}

func NewApplication(token string, c *config.Config) *Application {
	// If the client authentication token belongs to a user account, mimic the connection properties identical to what the official client assigns.
	id := gateway.DefaultIdentifier(token)
	if !strings.HasPrefix(token, "Bot ") {
		api.UserAgent = c.Identify.UserAgent

		id.Compress = false
		id.LargeThreshold = 0
		id.Properties = gateway.IdentifyProperties{
			OS:      c.Identify.OS,
			Browser: c.Identify.Browser,
			// The device property is assigned as an empty string by the official client.
			Device: "",
		}
	}

	return &Application{
		state:  state.NewWithIdentifier(id),
		config: c,
	}
}

func (a *Application) Start() error {
	err := a.state.Open(context.Background())
	if err != nil {
		return err
	}

	return err
}
