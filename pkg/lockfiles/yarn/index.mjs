// not worth rewriting this package in golang
import lockfile from "@yarnpkg/lockfile";
import { readFile } from "fs/promises";
import yaml from "js-yaml";

const file = await readFile(process.argv[2], "utf8");
try {
  const res = lockfile.parse(file);
  console.log(JSON.stringify(res));
} catch (err) {
  if (err.message.startsWith('Unknown token')){
    // yarn berry lockfiles can't be parsed by this, but are valid yaml:
    //  https://github.com/yarnpkg/berry/issues/2671
    try {
      const obj = yaml.load(file);
      console.log(JSON.stringify(obj));
    } catch (err) {
      console.log(err);
      process.exit(1);
    }
  } else {
    console.log(err);
    process.exit(1);
  }
}
