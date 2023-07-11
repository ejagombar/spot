package token

import (
	"context"
	"fmt"
	"strings"

	"github.com/spf13/viper"
	"github.com/zmb3/spotify/v2"
	"github.com/zmb3/spotify/v2/auth"
	"golang.org/x/oauth2"
)

func getClinet() *spotify.Client {
	viper.Get("token.timeout")
}
