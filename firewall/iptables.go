// +build linux

package firewall

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/flexiant/concerto/utils"
)

func driverName() string {
	return "iptables"
}

func apply(policy Policy) error {
	var exitCode int
	utils.RunCmd("/sbin/iptables -w -N CONCERTO")
	utils.RunCmd("/sbin/iptables -w -F CONCERTO")
	utils.RunCmd("/sbin/iptables -w -P INPUT DROP")

	_, exitCode, _, _ = utils.RunCmd("/sbin/iptables -w -C INPUT -i lo -j ACCEPT")
	if exitCode != 0 {
		utils.RunCmd("/sbin/iptables -w -A INPUT -i lo -j ACCEPT")
	}

	_, exitCode, _, _ = utils.RunCmd("/sbin/iptables -w -C INPUT -m state --state ESTABLISHED,RELATED -j ACCEPT")
	if exitCode != 0 {
		utils.RunCmd("/sbin/iptables -w -A INPUT -m state --state ESTABLISHED,RELATED -j ACCEPT")
	}

	for _, rule := range policy.Rules {
		utils.RunCmd(fmt.Sprintf("/sbin/iptables -w -A CONCERTO -s %s -p %s --dport %d:%d -j ACCEPT", rule.Cidr, rule.Protocol, rule.MinPort, rule.MaxPort))
	}

	_, exitCode, _, _ = utils.RunCmd("/sbin/iptables -w -C INPUT -j CONCERTO")
	if exitCode != 0 {
		log.Debugln("Concerto Chain is not existant adding it to INPUT")
		utils.RunCmd("/sbin/iptables -w -A INPUT -j CONCERTO")
	}

	return nil
}

func flush() error {
	utils.RunCmd("/sbin/iptables -w -P INPUT ACCEPT")
	utils.RunCmd("/sbin/iptables -w -F CONCERTO")
	utils.RunCmd("/sbin/iptables -w -D INPUT -j CONCERTO")
	utils.RunCmd("/sbin/iptables -w -X CONCERTO")
	return nil
}
