package twidiscord

import (
	"context"
	"time"

	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/diamondburned/twikit/twipi"
	"github.com/diamondburned/twikit/utils/cfgutil"
	"github.com/pkg/errors"
)

var ErrNotFound = errors.New("not found")

type Store struct {
	SecretStorer
	ChannelStorer
}

type Storer interface {
	SecretStorer
	ChannelStorer
	NumberIsMuted(context.Context, twipi.PhoneNumber) bool
	MuteNumber(context.Context, twipi.PhoneNumber, time.Time) error
	UnmuteNumber(context.Context, twipi.PhoneNumber) error
}

type SecretStorer interface {
	Account(context.Context, twipi.PhoneNumber) (Account, error)
	Accounts(context.Context) ([]Account, error)
	SetAccount(context.Context, Account) error
}

type ChannelStorer interface {
	ChannelToSerial(context.Context, discord.UserID, discord.ChannelID) (int, error)
	SerialToChannel(context.Context, discord.UserID, int) (discord.ChannelID, error)
}

type Account struct {
	UserNumber   twipi.PhoneNumber // key
	TwilioNumber twipi.PhoneNumber
	DiscordToken string
}

type Config struct {
	Discord struct {
		DatabaseURI  cfgutil.EnvString                `toml:"database_uri" json:"database_uri"`
		SecretsDir   cfgutil.EnvString                `toml:"secrets_dir" json:"secrets_dir"`
		KnownNumbers []cfgutil.Env[twipi.PhoneNumber] `toml:"known_numbers" json:"known_numbers"`
		AllowedUsers []cfgutil.Env[twipi.PhoneNumber] `toml:"allowed_users" json:"allowed_users"`
	} `toml:"discord" json:"discord"`
}
