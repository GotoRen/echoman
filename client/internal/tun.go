// Package internal contains the TCP/UDP connection,
// setups TUN/TAP Device, handles DNS packets.
package internal

import (
	"fmt"
	"os/exec"
	"runtime"

	"github.com/GotoRen/echoman/client/internal/logger"
	"github.com/songgao/water"
	"github.com/vishvananda/netlink"
	"golang.org/x/xerrors"
)

// TunInterface manages Tunnel device.
type TunInterface struct {
	Tun     *water.Interface
	address string
}

// NewTunInterface returns Tunnel device.
func NewTunInterface(name string, address string, prefix string) (*TunInterface, error) {
	addr := address + prefix

	switch runtime.GOOS {
	case "linux":
		config := water.Config{
			DeviceType: water.TUN,
		}
		config.Name = name

		ifce, err := water.New(config)
		if err != nil {
			return nil, err
		}

		iface := &TunInterface{
			Tun:     ifce,
			address: addr,
		}

		return iface, nil

	case "darwin":
		ifce, err := water.New(water.Config{
			DeviceType: water.TUN,
		})
		if err != nil {
			return nil, err
		}

		iface := &TunInterface{
			Tun:     ifce,
			address: addr,
		}

		return iface, nil

	case "windows":
		return nil, xerrors.Errorf("Windows is not supported")
	default:
		return nil, xerrors.Errorf("%s is not supported", runtime.GOOS)
	}
}

// Up function ups a virtual interface.
func (iface *TunInterface) Up() error {
	switch runtime.GOOS {
	case "linux":
		out, err := execCmd("ip", []string{"addr", "add", iface.address, "dev", iface.Tun.Name()})
		logger.LogDebug("Add a Virtual Interface", "Virtual Interface", out)

		if err != nil {
			logger.LogErr("ip command add fail", "error", err)
			return err
		}

		set, err := execCmd("ip", []string{"link", "set", "dev", iface.Tun.Name(), "up", "mtu", "1460"})
		logger.LogDebug("Up a Virtual Interface", "Virtual Interface", set)

		if err != nil {
			logger.LogErr("ip command set fail", "error", err)
			return err
		}

	case "darwin":
		out, err := execCmd("ifconfig", []string{iface.Tun.Name(), "up"})
		logger.LogDebug("Add a Virtual Interface", "Virtual Interface", out)

		if err != nil {
			logger.LogErr("ifconfig fail", "error", err)
			return err
		}

		if tun, err := netlink.LinkByName(iface.Tun.Name()); err == nil {
			addr, err := netlink.ParseAddr(iface.address)
			if err != nil {
				logger.LogErr("Unable to parse address", "error", err)
			}

			if err := netlink.AddrAdd(tun, addr); err != nil {
				logger.LogErr("Unable to add IP address to linked device", "error", err)
			}

			// TODO: Change MTU

			logger.LogDebug("Check Virtual Interface Name", "Virtual Interface Name", iface.Tun.Name())
			logger.LogDebug("Check Virtual Interface Address", "Virtual Interface Address", iface.address)
		}

	default:
		logger.LogErr("unsupported", "error", runtime.GOOS)
		logger.LogErr("unsupported", "error", runtime.GOARCH)
		return fmt.Errorf("unsupported: %s %s", runtime.GOOS, runtime.GOARCH)
	}

	return nil
}

// Read function read the virtual interface.
func (iface *TunInterface) Read(buf []byte) (int, error) {
	n, err := iface.Tun.Read(buf)
	// Read Virtual Interface.
	if err != nil {
		logger.LogErr("Failed to read virtual interface", "error", err)
		return 0, err
	}

	return n, nil
}

// Write function write the virtual interface.
func (iface *TunInterface) Write(buf []byte) (int, error) {
	return iface.Tun.Write(buf)
}

// Close function closes the virtual interface.
func (iface *TunInterface) Close() {
	if err := iface.Tun.Close(); err != nil {
		logger.LogErr("Failed to close virtual interface", "error", err)
	}
}

// execCmd executes the given arguments as a command.
func execCmd(cmd string, args []string) (string, error) {
	execCmd := exec.Command(cmd, args...)
	if err := execCmd.Run(); err != nil {
		logger.LogErr("Unable to execute command ", "error", err.Error())
		return execCmd.String(), err
	}

	return execCmd.String(), nil
}
