package firewall

import (
	"fmt"
	"strings"
)

func GetRouteName(routeTable, destination string) string {
	st := fmt.Sprintf("%s-%s", routeTable, destination)
	st = strings.ReplaceAll(st, ".", "-")
	st = strings.ReplaceAll(st, "/", "-")
	return st
}

func GetSGRuleName(securityGroup, cidr, port string) string {
	st := fmt.Sprintf("%s-%s-%s", securityGroup, cidr, port)
	st = strings.ReplaceAll(st, ".", "-")
	st = strings.ReplaceAll(st, "/", "-")
	return st
}
