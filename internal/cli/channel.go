package cli

import (
	"fmt"
	"os"
	"strings"
)

// InstallChannel controls whether Gentle AI uses the stable release path or an
// opt-in beta path for product-level dependencies such as core Engram.
type InstallChannel string

const (
	ChannelStable InstallChannel = "stable"
	ChannelBeta   InstallChannel = "beta"

	channelEnvVar = "GENTLE_AI_CHANNEL"
)

// ResolveInstallChannel resolves the channel from the flag value and env var.
// The explicit flag wins over the environment. Empty values default to stable.
func ResolveInstallChannel(flagValue string) (InstallChannel, error) {
	raw := strings.TrimSpace(flagValue)
	if raw == "" {
		raw = strings.TrimSpace(os.Getenv(channelEnvVar))
	}
	if raw == "" {
		return ChannelStable, nil
	}

	switch InstallChannel(strings.ToLower(raw)) {
	case ChannelStable:
		return ChannelStable, nil
	case ChannelBeta, "nightly":
		return ChannelBeta, nil
	default:
		return "", fmt.Errorf("unsupported Gentle AI channel %q (use stable, beta, or nightly)", raw)
	}
}

func (c InstallChannel) IsBeta() bool {
	return c == ChannelBeta
}
