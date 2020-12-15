package navmesh

import (
	"github.com/g3n/engine/math32"
	"gitlab.ferdoran.de/game-dev/go-sro/framework/math"
)

const (
	ObjectTileSize = 100.0
)

type ObjectTile struct {
	*math.Rectangle
	Index         int
	GlobalEdges   []*ObjectGlobalEdge
	Cells         []*ObjectCell
	InternalEdges []*ObjectInternalEdge
}

func NewObjectTile(gridOriginX, gridOriginY float32, index, globalEdgeCount int) *ObjectTile {
	x := float32(index % ObjectTileSize)
	y := float32(index / ObjectTileSize)
	minX := gridOriginX + (x * ObjectTileSize)
	minY := gridOriginY + (y * ObjectTileSize)

	return &ObjectTile{
		Rectangle: &math.Rectangle{
			Min: math32.NewVector2(minX, minY),
			Max: math32.NewVector2(minX+ObjectTileSize, minY+ObjectTileSize),
		},
		Index:         index,
		GlobalEdges:   make([]*ObjectGlobalEdge, globalEdgeCount),
		Cells:         make([]*ObjectCell, 0),
		InternalEdges: make([]*ObjectInternalEdge, 0),
	}
}
