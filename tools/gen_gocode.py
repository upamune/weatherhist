import json
import subprocess
t = r"""package weatherhist
// auto generated code
import (
	"fmt"

	"github.com/pkg/errors"
)

type StationID string

func GetStation(id string, groupNumber string) (Station, error) {
	sid := StationID(id + "_" + groupNumber)
	s, ok := allStations[sid]
	if !ok {
		return s, errors.New(fmt.Sprintf("failed to find station: %%s:%%s", id, groupNumber))
	}

	return s, nil
}

func GetAllStations() Stations {
	return allStations
}

const (
    StationTypeS = "s"
    StationTypeA = "a"
)

type Station struct {
    ID StationID
    GroupNumber string
    Name string
    Type string
}


type Stations map[StationID]Station

var allStations = Stations{
%s
}
"""

ot = r""""%04d_%s" : Station{
ID: StationID("%04d"),
GroupNumber: "%s",
Name: "%s",
Type: StationType%s,
},"""

f = open('stations.json')
dict = json.load(f)
f.close()

stations = dict['stations']

o = []

for s in stations:
    ots = ot % (s['station_id'], s['group_id'], s['station_id'], s['group_id'], s['station_name'], s['station_type'])
    o.append(ots)

gocode = t % ('\n'.join(o))

go_code_file_name = 'stations.go'

with open(go_code_file_name, 'w') as o:
    o.write(gocode)

cmd = "gofmt -w %s" % go_code_file_name

subprocess.call(cmd.strip().split(" "))
