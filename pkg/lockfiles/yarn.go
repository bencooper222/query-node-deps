package yarn

import (
	"encoding/json"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/bencooper222/query-node-deps/pkg/util"
	"github.com/google/uuid"
)

func GetParsedYarnLockFileFromLocalFile(loc string) []Package {
	arg := "node"
	file := "./pkg/lockfiles/yarn/index.mjs"

	cmd := exec.Command(arg, file, loc)
	stdout, err := cmd.Output()
	if err != nil {
		log.Fatalf("node command failed with stdout:\n%s\nerror: %v", stdout, err)
	}

	var parsedJson YarnLockJsonSchema
	err = json.Unmarshal(stdout, &parsedJson)
	if err != nil {
		log.Fatalf("failed to unmarshal json from stdout: %v", err)
	}
	if parsedJson.Object == nil {
		// try berry format
		var parsedBerryJson interface{}
		err = json.Unmarshal(stdout, &parsedBerryJson)
		if err != nil {
			log.Fatalf("failed to unmarshal json from stdout: %v", err)
		}
		return convertRawYarnBerryLockfileToFilteredFormat(parsedBerryJson.(map[string]interface{}))
	}
	return convertRawYarnLockfileToFilteredFormat(parsedJson)

}

type packagejson struct {
	Dependencies map[string]string `json:"dependencies"`
}

func GetPackageJSONDependencies(contents string) map[string]string {
	var parsedJson packagejson
	err := json.Unmarshal([]byte(contents), &parsedJson)
	if err != nil {
		log.Fatal("Problem unmarshalling package.json", err)
	}
	return parsedJson.Dependencies
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

	return GetParsedYarnLockFileFromLocalFile("tmp/" + genUuid)
}

func convertRawYarnLockfileToFilteredFormat(rawYarnLockfile YarnLockJsonSchema) []Package {
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

func convertRawYarnBerryLockfileToFilteredFormat(rawYarnLockfile map[string]interface{}) []Package {
	var rtn []Package
	for nameVersionSpec, info := range rawYarnLockfile {
		// deals with yarn spec lines that look like @thing/package@npm:1.0.0
		// TODO: this will fail for resolve@patch:resolve@1.1.7#~builtin<compat/resolve>
		name, versionSpec, matched := strings.Cut(nameVersionSpec, "@npm:")
		if !matched {
			log.Fatalf("could not parse name and version spec from %s", nameVersionSpec)
		}

		rtn = append(rtn, Package{
			Name:              name,
			SemverVersionSpec: versionSpec,
			ResolvedVersion:   info.(map[string]interface{})["version"].(string),
		})
	}

	return rtn
}
