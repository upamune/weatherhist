import json
t = r"""package weatherhist
// auto generated code
import (
	"fmt"

	"github.com/pkg/errors"
)

type ObservationID string

func GetObservatory(id string, groupNumber string) (Observation, error) {
	oid := ObservationID(id + "_" + groupNumber)
	ob, ok := allObservatories[oid]
	if !ok {
		return ob, errors.New(fmt.Sprintf("failed to find observatory: %%s:%%s", id, groupNumber))
	}

	return ob, nil
}

func GetAllObservatories() Observatries {
	return allObservatories
}

const (
    ObservatoryTypeS = "s"
    ObservatoryTypeA = "a"
)

type Observation struct {
    ID ObservationID
    GroupNumber string
    Name string
    Type string
}


type Observatries map[ObservationID]Observation

var allObservatories = Observatries{
%s
}
"""

ot = r""""%04d_%s" : Observation{
ID: ObservationID("%04d"),
GroupNumber: "%s",
Name: "%s",
Type: ObservatoryType%s,
},"""

f = open('observatories.json')
dict = json.load(f)
f.close()

observatories = dict['observatories']

o = []

for ob in observatories:
    ots = ot % (ob['observatory_id'], ob['group_id'], ob['observatory_id'], ob['group_id'], ob['observatory_name'], ob['observatory_type'])
    o.append(ots)

gocode = t % ('\n'.join(o))
print(gocode)
