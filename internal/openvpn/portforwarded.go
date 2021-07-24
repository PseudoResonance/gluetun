package openvpn

import (
	"context"
	"net"
	"net/http"
	"os"
	"strings"

	"github.com/qdm12/gluetun/internal/provider"
)

func (l *looper) GetPortForwarded() (port uint16) {
	return l.state.GetPortForwarded()
}

func (l *looper) PortForward(vpnGateway net.IP) { l.portForwardSignals <- vpnGateway }

// portForward is a blocking operation which may or may not be infinite.
// You should therefore always call it in a goroutine.
func (l *looper) portForward(ctx context.Context,
	providerConf provider.Provider, client *http.Client, gateway net.IP) {
	settings := l.state.GetSettings()
	if !settings.Provider.PortForwarding.Enabled {
		return
	}
	syncState := func(port uint16) (pfFilepath string) {
		l.state.SetPortForwarded(port)
		settings := l.state.GetSettings()
		return settings.Provider.PortForwarding.Filepath
	}
	providerConf.PortForward(ctx, client, l.pfLogger,
		gateway, l.fw, syncState)
}

func (l *looper) writeOpenvpnConf(lines []string) error {
	file, err := os.OpenFile(l.targetConfPath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	_, err = file.WriteString(strings.Join(lines, "\n"))
	if err != nil {
		return err
	}
	return file.Close()
}