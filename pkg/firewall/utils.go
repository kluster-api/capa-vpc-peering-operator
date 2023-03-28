package firewall

import (
	"fmt"
	"strings"
)

func GetRouteName(routeTable, destination string) string {
	st := fmt.Sprintf("%s_%s", routeTable, destination)
	st = strings.ReplaceAll(st, ".", "_")
	st = strings.ReplaceAll(st, "/", "_")
	return st
}

func GetSGRuleName(securityGroup, cidr, port string) string {
	st := fmt.Sprintf("%s_%s_%s", securityGroup, cidr, port)
	st = strings.ReplaceAll(st, ".", "_")
	st = strings.ReplaceAll(st, "/", "_")
	return st
}
