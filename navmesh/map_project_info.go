package navmesh

import (
	"encoding/binary"
	"log"
	"os"
)

const (
	RegionsX = 256
	RegionsY = 256

	RegionsTotal = RegionsX * RegionsY

	XSize   = 8
	XOffset = 0

	YSize         = 7
	YOffset       = XOffset + XSize
	DungeonOffset = YOffset + YSize
	DungeonMask   = ((1 << 1) - 1) << DungeonOffset
)

type MapProjectInfo struct {
	MapWidth           uint16
	MapHeight          uint16
	Short2             uint16
	Short3             uint16
	Short4             uint16
	Short5             uint16
	ActiveRegionsCount int
	MapRegions         []byte
	EnabledRegions     []uint16
}

func LoadMapProjectInfo(filename string) MapProjectInfo {
	file, err := os.Open(filename)
	if err != nil {
		log.Panic(err)
	}

	signature := make([]byte, 12)
	_, err = file.Read(signature)
	if err != nil {
		log.Panic("Failed to read signature", err)
	}

	if string(signature) != "JMXVMFO 1000" {
		log.Panicf("Invalid signature: %v\n", signature)
	}

	mapWidthBytes := make([]byte, 2)
	mapHeightBytes := make([]byte, 2)
	short2Bytes := make([]byte, 2)
	short3Bytes := make([]byte, 2)
	short4Bytes := make([]byte, 2)
	short5Bytes := make([]byte, 2)
	_, err = file.Read(mapWidthBytes)
	if err != nil {
		log.Panic("Failed to read map width", err)
	}
	_, err = file.Read(mapHeightBytes)
	if err != nil {
		log.Panic("Failed to read map height", err)
	}
	_, err = file.Read(short2Bytes)
	if err != nil {
		log.Panic("Failed to read short2", err)
	}
	_, err = file.Read(short3Bytes)
	if err != nil {
		log.Panic("Failed to read short3", err)
	}
	_, err = file.Read(short4Bytes)
	if err != nil {
		log.Panic("Failed to read short4", err)
	}
	_, err = file.Read(short5Bytes)
	if err != nil {
		log.Panic("Failed to read short5", err)
	}
	totalRegionBytes := make([]byte, RegionsTotal/8)
	_, err = file.Read(totalRegionBytes)
	if err != nil {
		log.Panic("Failed to read region bytes", err)
	}

	mapProjectInfo := MapProjectInfo{
		MapWidth:           binary.LittleEndian.Uint16(mapWidthBytes),
		MapHeight:          binary.LittleEndian.Uint16(mapHeightBytes),
		Short2:             binary.LittleEndian.Uint16(short2Bytes),
		Short3:             binary.LittleEndian.Uint16(short3Bytes),
		Short4:             binary.LittleEndian.Uint16(short4Bytes),
		Short5:             binary.LittleEndian.Uint16(short5Bytes),
		ActiveRegionsCount: 0,
		MapRegions:         totalRegionBytes,
		EnabledRegions:     make([]uint16, 0),
	}

	for z := 0; z < int(mapProjectInfo.MapHeight); z++ {
		for x := 0; x < int(mapProjectInfo.MapWidth); x++ {
			if IsEnabled(byte(x), byte(z), mapProjectInfo) {
				mapProjectInfo.ActiveRegionsCount++
				mapProjectInfo.EnabledRegions = append(mapProjectInfo.EnabledRegions, binary.LittleEndian.Uint16([]byte{byte(x), byte(z)}))
			}
		}
	}

	return mapProjectInfo
}

func IsEnabled(x, z byte, mapProjectInfo MapProjectInfo) bool {
	regionShort := binary.LittleEndian.Uint16([]byte{x, z})
	if (regionShort&DungeonMask)>>DungeonOffset != 0 {
		// It's a dungeon
		return false
	}

	if int(x) >= int(mapProjectInfo.MapWidth) || int(z) >= int(mapProjectInfo.MapHeight) {
		return false
	}
	return (mapProjectInfo.MapRegions[regionShort>>3] & byte(uint16(128>>(regionShort%8)))) != 0
}
