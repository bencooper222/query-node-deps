package yarn

import (
	"encoding/json"
	"os/exec"
	"strings"

	"github.com/bencooper222/query-node-deps/pkg/util"
)

func GetParsedYarnLockfile(loc string) []Package {
	arg := "node"
	file := "./pkg/lockfiles/yarn/index.mjs"

	cmd := exec.Command(arg, file, loc)
	stdout, err := cmd.Output()
	util.CheckErr(err)

	var parsedJson JsonSchema
	json.Unmarshal(stdout, &parsedJson)

	// log.Println(parsedJson.Object[])

	return convertRawYarnLockfileToFilteredFormat(parsedJson)
}

func convertRawYarnLockfileToFilteredFormat(rawYarnLockfile JsonSchema) []Package {
	obj := rawYarnLockfile.Object

	var rtn []Package
	for nameVersionSpec, info := range obj {
		sections := strings.Split(nameVersionSpec, "@")

		name := sections[0]
		versionSpec := sections[1]

		rtn = append(rtn, Package{
			Name:              name,
			SemverVersionSpec: versionSpec,
			ResolvedVersion:   info.Version,
		})
	}

	return rtn
}
