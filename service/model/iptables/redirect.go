package iptables

import (
	"V2RayA/tools/cmds"
)

type redirect struct{ iptablesSetter }

var Redirect redirect

func (r *redirect) GetSetupCommands() SetupCommands {
	commands := `
iptables -t nat -N V2RAY
iptables -t nat -A V2RAY -d 10.0.0.0/8 -j RETURN
iptables -t nat -A V2RAY -d 100.64.0.0/10 -j RETURN
iptables -t nat -A V2RAY -d 127.0.0.0/8 -j RETURN
iptables -t nat -A V2RAY -d 169.254.0.0/16 -j RETURN
iptables -t nat -A V2RAY -d 172.16.0.0/12 -j RETURN
iptables -t nat -A V2RAY -d 192.0.0.0/24 -j RETURN
iptables -t nat -A V2RAY -d 192.0.2.0/24 -j RETURN
iptables -t nat -A V2RAY -d 192.88.99.0/24 -j RETURN
iptables -t nat -A V2RAY -d 192.168.0.0/16 -j RETURN
iptables -t nat -A V2RAY -d 198.18.0.0/15 -j RETURN
iptables -t nat -A V2RAY -d 198.51.100.0/24 -j RETURN
iptables -t nat -A V2RAY -d 203.0.113.0/24 -j RETURN
iptables -t nat -A V2RAY -d 224.0.0.0/4 -j RETURN
iptables -t nat -A V2RAY -d 240.0.0.0/4 -j RETURN
iptables -t nat -A V2RAY -d 255.255.255.255/32 -j RETURN
iptables -t nat -A V2RAY -m mark --mark 0xff -j RETURN
iptables -t nat -A V2RAY -p tcp -j REDIRECT --to-ports 12345

iptables -t nat -A PREROUTING -p tcp -j V2RAY
iptables -t nat -A OUTPUT -p tcp -j V2RAY
`
	if cmds.IsCommandValid("ip6tables") {
		commands += `
ip6tables -t nat -N V2RAY
#禁用ipv6
ip6tables -t nat -A V2RAY -p tcp -j REDIRECT --to-port 0
ip6tables -t nat -A V2RAY -p udp -j REDIRECT --to-port 0

ip6tables -t nat -A V2RAY -m mark --mark 0xff -j RETURN
ip6tables -t nat -A V2RAY -d ::/128 -j RETURN
ip6tables -t nat -A V2RAY -d ::1/128 -j RETURN
ip6tables -t nat -A V2RAY -d ::ffff:0:0/96 -j RETURN
ip6tables -t nat -A V2RAY -d ::ffff:0:0:0/96 -j RETURN
ip6tables -t nat -A V2RAY -d 64:ff9b::/96 -j RETURN
ip6tables -t nat -A V2RAY -d 100::/64 -j RETURN
ip6tables -t nat -A V2RAY -d 2001::/32 -j RETURN
ip6tables -t nat -A V2RAY -d 2001:20::/28 -j RETURN
ip6tables -t nat -A V2RAY -d 2001:db8::/32 -j RETURN
ip6tables -t nat -A V2RAY -d 2002::/16 -j RETURN
ip6tables -t nat -A V2RAY -d fc00::/7 -j RETURN
ip6tables -t nat -A V2RAY -d fe80::/10 -j RETURN
ip6tables -t nat -A V2RAY -d ff00::/8 -j RETURN
ip6tables -t nat -A V2RAY -p tcp -j REDIRECT --to-ports 12345
ip6tables -t nat -A PREROUTING -p tcp -j V2RAY
ip6tables -t nat -A PREROUTING -p udp -j V2RAY
ip6tables -t nat -A OUTPUT -p tcp -j V2RAY
ip6tables -t nat -A OUTPUT -p udp -j V2RAY
	`
	}
	return SetupCommands(commands)
}

func (r *redirect) GetCleanCommands() CleanCommands {
	commands := `
iptables -t nat -F V2RAY
iptables -t nat -D PREROUTING -p tcp -j V2RAY
iptables -t nat -D PREROUTING -p udp -j V2RAY
iptables -t nat -D OUTPUT -p tcp -j V2RAY
iptables -t nat -D OUTPUT -p udp -j V2RAY
iptables -t nat -X V2RAY
`
	if cmds.IsCommandValid("ip6tables") {
		commands += `
ip6tables -t nat -F V2RAY
ip6tables -t nat -D PREROUTING -p tcp -j V2RAY
ip6tables -t nat -D PREROUTING -p udp -j V2RAY
ip6tables -t nat -D OUTPUT -p tcp -j V2RAY
ip6tables -t nat -D OUTPUT -p udp -j V2RAY
ip6tables -t nat -X V2RAY
`
	}
	return CleanCommands(commands)
}