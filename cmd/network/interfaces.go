package network

import (
	"fmt"
	"git.f-i-ts.de/cloud-native/metal/metal-hammer/metal-core/models"
	"git.f-i-ts.de/cloud-native/metal/metal-hammer/pkg/lldp"
	"git.f-i-ts.de/cloud-native/metallib/version"
	log "github.com/inconshreveable/log15"
	"github.com/vishvananda/netlink"
	"net"
	"strings"
	"time"
)

// Network provides networking operations.
type Network struct {
	IPAddress  string
	Started    time.Time
	DeviceUUID string
	LLDPClient *LLDPClient
}

// UpAllInterfaces set all available eth* interfaces up
// to ensure they do ipv6 link local autoconfiguration and
// therefore neighbor discovery,
// which is required to make all local mac's visible on the switch side.
func (n *Network) UpAllInterfaces() error {
	description := fmt.Sprintf("metal-hammer IP:%s version:%s waiting since %s for installation", n.IPAddress, version.V, n.Started)
	interfaces := make([]string, 0)
	for _, name := range Interfaces() {
		if !strings.HasPrefix(name, "eth") {
			continue
		}
		interfaces = append(interfaces, name)

		err := linkSetUp(name)
		if err != nil {
			return fmt.Errorf("Error set link %s up: %v", name, err)
		}

		lldpd, err := lldp.NewDaemon(n.DeviceUUID, description, name, 5*time.Second)

		if err != nil {
			return fmt.Errorf("Error start lldpd on %s info: %v", name, err)
		}
		lldpd.Start()
	}

	lc := NewLLDPClient(interfaces, 2, 2, 0)
	n.LLDPClient = lc
	go lc.Start()

	return nil
}

func linkSetUp(name string) error {
	iface, err := netlink.LinkByName(name)
	if err != nil {
		return err
	}
	err = netlink.LinkSetUp(iface)
	if err != nil {
		return err
	}
	return nil
}

// Neighbors of a interface, detected via ip neighbor detection
func (n *Network) Neighbors(name string) ([]*models.ModelsMetalNic, error) {
	neighbors := make([]*models.ModelsMetalNic, 0)

	host := n.LLDPClient.Host

	for !host.done {
		log.Info("not all lldp pdu's received, waiting...", "interface", name)
		time.Sleep(1 * time.Second)

		duration := time.Now().Sub(host.start)
		if duration > host.timeout {
			return nil, fmt.Errorf("not all neighbor requirements where met within: %s, exiting", host.timeout)
		}
	}
	log.Info("all lldp pdu's received", "interface", name)

	neighs, _ := host.neighbors[name]
	for _, neigh := range neighs {
		if neigh.Port.Type != lldp.Mac {
			continue
		}
		macAddress := neigh.Port.Value
		neighbors = append(neighbors, &models.ModelsMetalNic{Mac: &macAddress})
	}
	return neighbors, nil
}

// InternalIP returns the first ipv4 ip of a eth* interface.
func InternalIP() string {
	for _, name := range Interfaces() {
		if !strings.HasPrefix(name, "eth") {
			continue
		}
		itf, _ := net.InterfaceByName(name)
		item, _ := itf.Addrs()
		for _, addr := range item {
			switch v := addr.(type) {
			case *net.IPNet:
				if !v.IP.IsLoopback() {
					if v.IP.To4() != nil {
						return v.IP.String()
					}
				}
			}
		}
	}
	return ""
}

// Interfaces return a list of all known interfaces.
func Interfaces() []string {
	var interfaces []string
	links, err := netlink.LinkList()
	if err != nil {
		return interfaces
	}
	for _, nic := range links {
		name := nic.Attrs().Name
		interfaces = append(interfaces, name)
	}
	return interfaces
}
