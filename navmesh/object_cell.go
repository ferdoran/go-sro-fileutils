package navmesh

import "gitlab.ferdoran.de/game-dev/go-sro/framework/math"

type ObjectCell struct {
	*math.Triangle
	Index int
	Flag  uint16
}
