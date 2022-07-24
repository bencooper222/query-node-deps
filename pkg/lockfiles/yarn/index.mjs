// not worth rewriting this package in golang
import lockfile from "@yarnpkg/lockfile";
import { readFile } from "fs/promises";

try {
  const file = await readFile(process.argv[2], "utf8");
  const res = lockfile.parse(file);
  console.log(JSON.stringify(res));
} catch (err) {
  console.error(err);
  process.exit(1);
}
