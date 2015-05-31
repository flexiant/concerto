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
	utils.RunCmd("/sbin/iptables -w -N CONCERTO")
	utils.RunCmd("/sbin/iptables -w -F CONCERTO")
	utils.RunCmd("/sbin/iptables -w -P INPUT DROP")
	utils.RunCmd("/sbin/iptables -w -A INPUT -i lo -j ACCEPT")
	utils.RunCmd("/sbin/iptables -w -A INPUT -m state --state ESTABLISHED,RELATED -j ACCEPT")

	for _, rule := range policy.Rules {
		utils.RunCmd(fmt.Sprintf("/sbin/iptables -w -A CONCERTO -s %s -p %s --dport %d:%d -j ACCEPT", rule.Cidr, rule.Protocol, rule.MinPort, rule.MaxPort))
	}

	utils.RunCmd("/sbin/iptables -w -C INPUT -j CONCERTO && /sbin/iptables -w -A INPUT -j CONCERTO")

	return nil
}

func flush() error {
	utils.RunCmd("/sbin/iptables -w -P INPUT ACCEPT")
	utils.RunCmd("/sbin/iptables -w -F CONCERTO")
	utils.RunCmd("/sbin/iptables -w -X CONCERTO")
	return nil
}
