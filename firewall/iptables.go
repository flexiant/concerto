// +build linux

package firewall

import (
	"fmt"
	"github.com/flexiant/concerto/utils"
)

func driverName() string {
	return "iptables"
}

func apply(policy Policy) error {
	utils.RunCmd("/sbin/iptables -F INPUT")
	utils.RunCmd("/sbin/iptables -P INPUT DROP")
	utils.RunCmd("/sbin/iptables -A INPUT -i lo -j ACCEPT")
	utils.RunCmd("/sbin/iptables -A INPUT -m state --state ESTABLISHED,RELATED -j ACCEPT")

	for _, rule := range policy.Rules {
		utils.RunCmd(fmt.Sprintf("/sbin/iptables -A INPUT -s %s -p %s --dport %d:%d -j ACCEPT", rule.Cidr, rule.Protocol, rule.MinPort, rule.MaxPort))
	}

	return nil
}

func flush() error {
	utils.RunCmd("/sbin/iptables -P INPUT DROP")
	utils.RunCmd("/sbin/iptables -F INPUT")
	return nil
}
