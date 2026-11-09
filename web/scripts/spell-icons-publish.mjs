// web/scripts/spell-icons-publish.mjs
// The report's spell icons beyond the build's id table (sync-data.mjs calls it for every build).
// The table (`spellicons.json`) carries a player spell's own id only, so most of a log's spells
// (another rank, a Forever spell, a consumable) and every melee swing would draw an empty socket.
// This writes `spellicons-by-name.json` (a name to the icon of a same-named spell in the table)
// beside it, and for the fixture source, which publishes neither the table nor the full icon set,
// the table itself and the icon files the bundled sample and the report fixtures name, plus the
// fallbacks (web/src/data/spell-icon-fallbacks.json: melee, the client's own overrides, each
// school's base spell). Icons are the pipeline's own files, byte for byte.
import { access, cp, mkdir, readFile, readdir, writeFile } from 'node:fs/promises';
import path from 'node:path';

const exists = (file) =>
  access(file).then(
    () => true,
    () => false,
  );

const readJson = async (file) => JSON.parse(await readFile(file, 'utf8'));

/** Every `spell_id` and every spell name (`name`, `spell_name`) in a parsed report file. */
export function reportSpells(value, out = { ids: new Set(), names: new Set() }) {
  if (Array.isArray(value)) {
    for (const entry of value) reportSpells(entry, out);
  } else if (value !== null && typeof value === 'object') {
    for (const [key, entry] of Object.entries(value)) {
      if (key === 'spell_id' && typeof entry === 'number') out.ids.add(entry);
      else if ((key === 'spell_name' || key === 'name') && typeof entry === 'string') out.names.add(entry);
      else reportSpells(entry, out);
    }
  }
  return out;
}

/** A name to the icon of the first spell of that name the table carries an icon for. */
export function iconsByName(spells, table) {
  const byName = {};
  for (const spell of spells) {
    const stem = table[String(spell.id)];
    if (stem !== undefined && byName[spell.name] === undefined) byName[spell.name] = stem;
  }
  return byName;
}

/** The report fixture files a fixture build serves: the bundled sample, src/fixtures/report and the coach's. */
async function reportFiles(webRoot) {
  const out = [];
  const walk = async (dir) => {
    let entries = [];
    try {
      entries = await readdir(dir, { withFileTypes: true });
    } catch {
      return;
    }
    for (const entry of entries) {
      const full = path.join(dir, entry.name);
      if (entry.isDirectory()) await walk(full);
      else if (entry.name.endsWith('.json')) out.push(full);
    }
  };
  await walk(path.join(webRoot, 'public/sample-report'));
  await walk(path.join(webRoot, 'src/fixtures/report'));
  await walk(path.join(webRoot, 'src/fixtures/coach'));
  return out;
}

/**
 * @param {{ repoRoot: string, webRoot: string, build: string, publicDir: string, fixture: boolean }} options
 * @returns {Promise<string[]>} what it wrote
 */
export async function publishSpellIcons({ repoRoot, webRoot, build, publicDir, fixture }) {
  const buildDir = path.join(repoRoot, 'data/builds', build);
  const tableFile = path.join(buildDir, 'spellicons.json');
  const spellsFile = path.join(buildDir, 'spells.json');
  if (!(await exists(tableFile)) || !(await exists(spellsFile))) return [];
  const table = await readJson(tableFile);
  const byName = iconsByName(await readJson(spellsFile), table);
  await writeFile(path.join(publicDir, 'spellicons-by-name.json'), `${JSON.stringify(byName)}\n`);
  if (!fixture) return ['spellicons-by-name.json'];

  await cp(tableFile, path.join(publicDir, 'spellicons.json'));
  const fallbacks = await readJson(path.join(webRoot, 'src/data/spell-icon-fallbacks.json'));
  const wanted = reportSpells({});
  for (const file of await reportFiles(webRoot)) reportSpells(await readJson(file), wanted);
  const stems = new Set([
    fallbacks.melee,
    ...Object.values(fallbacks.roles),
    ...Object.values(fallbacks.client),
    ...Object.values(fallbacks.schools),
  ]);
  for (const id of wanted.ids) if (table[String(id)] !== undefined) stems.add(table[String(id)]);
  for (const name of wanted.names) if (byName[name] !== undefined) stems.add(byName[name]);

  const iconsDir = path.join(publicDir, 'icons');
  await mkdir(iconsDir, { recursive: true });
  let added = 0;
  for (const stem of stems) {
    const from = path.join(buildDir, 'icons', `${stem}.webp`);
    // Over the fixture's own file too: its icon set holds 1 px placeholders for some of these stems.
    if (!(await exists(from))) continue;
    const target = path.join(iconsDir, `${stem}.webp`);
    await cp(from, target);
    added += 1;
  }
  return ['spellicons-by-name.json', 'spellicons.json', ...(added === 0 ? [] : [`spell icons (${added})`])];
}
