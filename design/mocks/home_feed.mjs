// The signed-in home's "Latest on Forever" rows, from the build's own models, for ws_home.py:
//   node design/mocks/home_feed.mjs <request.json>   (prints the rows as JSON)
// request.json: { rows, spec, ownGuild, raidOnPlaque }, the feed fixture's rows and the viewer.
// It runs lib/home/floor-load.ts floorInputs and lib/home/floor.ts buildFloor on the guide and
// changelog frontmatter (index.astro's own calls, at today's date as a build made now has them),
// then lib/home/feed-mix.ts signedInRows (HomeFeed's signed-in call): the viewer's guild left off
// its sims, one row per character, at most three sims, the viewer's spec's site events mixed in.
import { createRequire } from "node:module";
import { readFileSync, readdirSync } from "node:fs";
import { join } from "node:path";

const WEB = new URL("../../web/", import.meta.url).pathname;
const require = createRequire(join(WEB, "package.json"));
const { createJiti } = require("jiti");
const YAML = require("yaml");
const jiti = createJiti(import.meta.url);

const frontmatter = (file) =>
  YAML.parse(/^---\n([\s\S]*?)\n---/.exec(readFileSync(file, "utf8"))[1]);

const GUIDES = join(WEB, "src/content/guides");
const guides = readdirSync(GUIDES, { withFileTypes: true })
  .filter((dir) => dir.isDirectory())
  .flatMap((dir) =>
    readdirSync(join(GUIDES, dir.name))
      .filter((file) => file.endsWith(".md"))
      .map((file) => {
        const data = frontmatter(join(GUIDES, dir.name, file));
        return {
          classSlug: data.classSlug,
          spec: data.spec,
          updated: new Date(data.updated),
        };
      }),
  );
const CHANGELOG = join(WEB, "src/content/changelog");
const changelog = readdirSync(CHANGELOG)
  .filter((file) => file.endsWith(".md"))
  .map((file) => {
    const data = frontmatter(join(CHANGELOG, file));
    return {
      id: file.replace(/\.md$/, ""),
      title: data.title,
      date: new Date(data.date),
    };
  });

const { floorInputs } = await jiti.import(
  join(WEB, "src/lib/home/floor-load.ts"),
);
const { buildFloor } = await jiti.import(join(WEB, "src/lib/home/floor.ts"));
const { signedInRows } = await jiti.import(
  join(WEB, "src/lib/home/feed-mix.ts"),
);

const request = JSON.parse(readFileSync(process.argv[2], "utf8"));
const floor = buildFloor(floorInputs(guides, changelog, new Date()));
const rows = signedInRows({
  rows: request.rows,
  floor,
  spec: request.spec ?? undefined,
  ownGuild: request.ownGuild ?? undefined,
  raidOnPlaque: request.raidOnPlaque,
});
process.stdout.write(JSON.stringify(rows));
