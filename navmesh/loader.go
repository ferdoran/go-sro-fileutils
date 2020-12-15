package navmesh

import (
	"encoding/gob"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"gitlab.ferdoran.de/game-dev/go-sro/framework/utils"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

type Loader struct {
	DataPath       string
	NavMeshPath    string
	MapProjectInfo MapProjectInfo
	ObjectInfo     ObjectInfo
	NavMeshData    map[string]NavMeshData
}

func NewLoader(dataPath string) *Loader {
	return &Loader{
		DataPath:       dataPath,
		NavMeshPath:    dataPath + string(os.PathSeparator) + "navmesh",
		MapProjectInfo: MapProjectInfo{},
		ObjectInfo:     ObjectInfo{},
	}
}

func (l *Loader) LoadNavMeshInfos() {
	l.MapProjectInfo = LoadMapProjectInfo(l.NavMeshPath + string(os.PathSeparator) + "mapinfo.mfo")
	l.ObjectInfo = LoadObjectInfo(l.NavMeshPath + string(os.PathSeparator) + "object.ifo")
}

func (l *Loader) LoadNavMeshData() map[string]NavMeshData {
	navMeshData := make(map[string]NavMeshData, 0)
	navdataFiles := make(map[string][]byte)
	for x := 0; x < len(l.MapProjectInfo.EnabledRegions); x++ {
		regionShortHex := fmt.Sprintf("%x", l.MapProjectInfo.EnabledRegions[x])
		navMeshHex := fmt.Sprintf("nv_%s.nvm", regionShortHex)
		fmt.Printf("\rReading %s. Finished [%d / %d] files", navMeshHex, x, len(l.MapProjectInfo.EnabledRegions))
		fileContent, _ := ioutil.ReadFile(l.NavMeshPath + string(os.PathSeparator) + navMeshHex)
		navdataFiles[navMeshHex] = fileContent
	}
	fmt.Println()
	fileCounter := 0
	for k, v := range navdataFiles {
		fileCounter++
		navMeshData[k] = ParseNavMeshFile(k, v)
		counter := 0
		logrus.Debugf("Loading file %d/%d", fileCounter, len(navdataFiles))
		for _, o := range navMeshData[k].Objects {
			counter++
			obj := l.ObjectInfo.Objects[o.ID]
			var res *Resource
			var mesh *MeshFile
			if strings.HasSuffix(obj.FilePath, "cpd") {
				cpd := LoadCompoundFile(l.DataPath + string(os.PathSeparator) + obj.FilePath)
				res = LoadResource(l.DataPath + string(os.PathSeparator) + cpd.NavMeshObjPath)
				mesh = LoadMeshFile(l.DataPath + string(os.PathSeparator) + res.NavMeshObjPath)
			} else if strings.HasSuffix(obj.FilePath, "bsr") {
				res = LoadResource(l.DataPath + string(os.PathSeparator) + obj.FilePath)
				mesh = LoadMeshFile(l.DataPath + string(os.PathSeparator) + res.NavMeshObjPath)
			} else if strings.HasSuffix(obj.FilePath, "bms") {
				mesh = LoadMeshFile(l.DataPath + string(os.PathSeparator) + obj.FilePath)
			} else {
				logrus.Panicf("unsupported file: %s\n", obj.FilePath)
			}
			mesh.LoadMeshObject(o)
			utils.PrintProgress(counter, int(navMeshData[k].ObjectCount))
			logrus.Tracef("Loaded file %s. Points to %s", obj.FilePath, res.NavMeshObjPath)
		}
	}
	l.NavMeshData = navMeshData
	return navMeshData
}

func (l *Loader) SaveNavmeshDataAsGOB(filepath string) {
	logrus.Infoln("saving navmesh data as gob")
	//jsonData, err := json.Marshal(l.NavMeshData)
	//if err != nil {
	//	logrus.Errorf("failed to marshal json data. Err = %v\n", err)
	//}
	logrus.Infoln("creating precomuted file")
	f, err := os.Create(filepath)
	if err != nil {
		logrus.Errorf("failed to create file. Err = %v\n", err)
	}

	logrus.Debugf("writing compressed data")
	//w, err := zlib.NewWriterLevel(f, zlib.BestCompression)
	w := gob.NewEncoder(f)
	err = w.Encode(l.NavMeshData)
	if err != nil {
		logrus.Error(err)
	}
	logrus.Debugf("saved navmesh data as gob")
	f.Close()
}

func (l *Loader) SaveNavmeshDataAsJSON(filepath string) {
	logrus.Infoln("saving navmesh data as json")
	//jsonData, err := json.Marshal(l.NavMeshData)
	//if err != nil {
	//	logrus.Errorf("failed to marshal json data. Err = %v\n", err)
	//}
	logrus.Infoln("creating precomuted file")
	f, err := os.Create(filepath)
	if err != nil {
		logrus.Errorf("failed to create file. Err = %v\n", err)
	}

	logrus.Debugf("writing compressed data")
	//w, err := zlib.NewWriterLevel(f, zlib.BestCompression)
	w := json.NewEncoder(f)
	err = w.Encode(l.NavMeshData)
	if err != nil {
		logrus.Error(err)
	}
	logrus.Debugf("saved navmesh data as json")
	f.Close()
}

func (l *Loader) LoadPrecomputedNavmeshDataFromGOB(filepath string) {
	logrus.Debugf("loading precomputed navmeshdata\n")
	f, err := os.Open(filepath)

	if err != nil {
		logrus.Error(err)
	}

	decodeStartTime := time.Now()
	r := gob.NewDecoder(f)
	err = r.Decode(&l.NavMeshData)
	decodeStopTime := time.Now()

	if err != nil {
		logrus.Error(err)
	}

	logrus.Debugf("finished loading navmesh data after %d ms\n", decodeStopTime.Sub(decodeStartTime).Milliseconds())

}
