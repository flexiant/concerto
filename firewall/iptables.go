// +build linux
package firewall

import (
	log "github.com/Sirupsen/logrus"
	"github.com/flexiant/krane/utils"
	"os"
	"os/exec"
	"strings"
)

func driverName() string {
	return "iptables"
}

func apply(policy Policy) error {
	utils.RunCmd("iptables -A INPUT -i lo -j ACCEPT")
	utils.RunCmd("iptables -A INPUT -m state --state ESTABLISHED,RELATED -j ACCEPT")

	for _, rule := range policy.Rules {
		utils.RunCmd("iptables -A INPUT -s %s -p %s --dport %d:%d -j ACCEPT\n", rule.Cidr, rule.Protocol, rule.MinPort, rule.MaxPort)
	}
	utils.RunCmd("iptables -P INPUT ACCEPT")
	utils.RunCmd("iptables -F INPUT")
	return nil
}

func flush() error {
	utils.RunCmd("iptables -F INPUT")
	utils.RunCmd("iptables -P INPUT DROP")
	return nil
}
