package main

import (
	"fmt"
	hgt "map-dem/nasadem_hgt"
)

func main() {
	elevationFileStorage := hgt.NewMemcacheNasaElevationFileStorage("/nasadem_hgt/tmp", "Bearer eyJ0eXAiOiJKV1QiLCJvcmlnaW4iOiJFYXJ0aGRhdGEgTG9naW4iLCJzaWciOiJlZGxqd3RwdWJrZXlfb3BzIiwiYWxnIjoiUlMyNTYifQ.eyJ0eXBlIjoiVXNlciIsInVpZCI6ImRvbmdsaWFuZ2RlIiwiZXhwIjoxNjY5MjY5Njk2LCJpYXQiOjE2NjQwODU2OTYsImlzcyI6IkVhcnRoZGF0YSBMb2dpbiJ9.UhIB0N5Tm3H0PoyQnmo-5bvWslbyLMraCjObAvrcgEX_g3QXaCWaCIrJIg4yp27nIaMUhvvD-8gD7MlGHOHnCFRmxYCrGFqn4Q_FNBxfDMZCUUJYW7GY_rUhvInEHbPoCn7okt2rn_hctQmQxz8K_mCPaiN4RdwI7ps5ARu7qQ41osg_uf-UYwiyxGGlulWQhRtPQtj5iLYLXPgoCL4E8ThhR-c0qckVRhRi8m6Exji-A21Rsrq-wJAbXBo3HwskDZ3rquTbNAtGObS9pVMEaeWtn5CKcRlii0mYdgzgKsSHJ-jMYOupz23IpF7YAIlZUpLNQjDZc4hD-1dsbsQozw")
	elevationWayGetterCloser := hgt.NewHgt(elevationFileStorage)
	points := []hgt.Point{{28.3525, 85.77917}}
	elevation, err := elevationWayGetterCloser.Get(points)
	if err != nil {
		panic(err)
	}
	fmt.Println(elevation)
}
