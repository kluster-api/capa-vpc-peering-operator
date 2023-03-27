package firewall

import "fmt"

func GetRouteName(routeTable, destination string) string {
	return fmt.Sprintf("%s_%s", routeTable, destination)
}
