package yarn

import (
	"encoding/json"
	"log"
	"os"
	"os/exec"

	"github.com/bencooper222/query-node-deps/pkg/util"
	"github.com/google/uuid"
)

func GetParsedYarnLockfileFromLocalFile(loc string) []Package {
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

func GetParsedYarnLockfileFromAlreadyStringifiedLockfile(contents string) []Package {
	if err := os.MkdirAll("tmp", os.ModePerm); err != nil {
		log.Fatal(err)
	}

	genUuid := uuid.New().String()
	f, _ := os.Create("tmp/" + genUuid)
	defer f.Close()

	f.WriteString(contents)
	f.Sync()

	return GetParsedYarnLockfileFromLocalFile("tmp/" + genUuid)
}

func convertRawYarnLockfileToFilteredFormat(rawYarnLockfile JsonSchema) []Package {
	obj := rawYarnLockfile.Object

	var rtn []Package
	for nameVersionSpec, info := range obj {
		// deals with yarn spec lines that look like @thing/package@1.0.0
		name, versionSpec := util.SplitOnLastAppearanceOfDelimiter(nameVersionSpec, "@")

		rtn = append(rtn, Package{
			Name:              name,
			SemverVersionSpec: versionSpec,
			ResolvedVersion:   info.Version,
		})
	}

	return rtn
}
