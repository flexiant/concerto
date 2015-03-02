// +build windows

package firewall

import (
	"fmt"
	"github.com/flexiant/concerto/utils"
)

func driverName() string {
	return "windows"
}

func apply(policy Policy) error {
	utils.RunCmd("netsh advfirewall set allprofiles state off")
	utils.RunCmd("netsh advfirewall set allprofiles firewallpolicy blockinbound,allowoutbound")
	utils.RunCmd("netsh advfirewall firewall delete rule name=all")

	for _, rule := range policy.Rules {
		utils.RunCmd(fmt.Sprintf("netsh advfirewall firewall add rule name=\"Concerto firewall\" dir=in action=allow remoteip=#{%s} protocol=#{%s} localport=#{%d}-#{%s}", rule.Cidr, rule.Protocol, rule.MinPort, rule.MaxPort))
	}

	utils.RunCmd("netsh advfirewall set allprofiles state on")
	return nil
}

func flush() error {
	utils.RunCmd("netsh advfirewall set allprofiles state off")
	utils.RunCmd("netsh advfirewall set allprofiles firewallpolicy allowinbound,allowoutbound")
	utils.RunCmd("netsh advfirewall firewall delete rule name=all")
	return nil
}
