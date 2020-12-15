package main

import (
	"bufio"
	"github.com/sirupsen/logrus"
	"gitlab.ferdoran.de/game-dev/go-sro/framework/logging"
	"os"
)

func main() {
	//logFile, err := os.Create("log.txt")
	//if err != nil {
	//	log.Panicln("Failed to open logfile")
	//}
	logging.Init()
	//log.SetOutput(logFile)
	//r := pk2.NewPk2Reader("/Users/rmueller1/dev/workspaces/private/sro/Media.pk2")
	//r.ReadFile()
	//logrus.Infof("root dir name %s", r.Directory.Name)
	//data, err := r.LoadFile("Media/server_dep/silkroad/textdata/skilldata.txt")
	//if err != nil {
	//	logrus.Error(err)
	//} else {
	//	logrus.Infof("loaded %d bytes", len(data))
	//	fmt.Printf("%s", string(data))
	//}
	//reader.ExtractFiles("E:\\Data")

	//itemData := itemdata.ReadItemData("E:\\Silkroad_TestIn\\Media\\server_dep\\silkroad\\textdata\\itemdata_15000.txt")
	//
	//log.Printf("ItemData contained %d lines\n", len(itemData.Items))
	//itemData := make([]itemdata.ItemData, 0)
	//lines := 0
	//
	//for i := 5000; i < 50000; i += 5000 {
	//	id := itemdata.ReadItemData("E:\\Media\\server_dep\\silkroad\\textdata\\itemdata_" + strconv.FormatInt(int64(i), 10) + ".txt")
	//	lines += len(id.Items)
	//	itemData = append(itemData, id)
	//}
	//
	//log.Printf("Read %d items", lines)
	//dataBasePath := "E:\\Data"
	//loader := navmesh.NewLoader(dataBasePath)
	//loader.LoadNavMeshInfos()
	//loader.LoadNavMeshData()
	//visual.NewRendering(navMeshData)

	reader := bufio.NewReader(os.Stdin)
	logrus.Println("Press Enter to exit...")
	reader.ReadString('\n')

}
