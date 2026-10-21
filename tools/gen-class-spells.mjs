// tools/gen-class-spells.mjs
//
// Writes api/internal/engine/classspells.json: which class owns a spell id, and which spec a
// spell id defines, read from one client build's data. The parse pipeline uses the table to
// read a class and a spec off the spells a player cast when the log carries no COMBATANT_INFO.
//
//   node tools/gen-class-spells.mjs [data/builds/<build>]
//
// A spell id is a class spell only when exactly one class lists it. A spec spell is a talent
// spell (any rank) that belongs to exactly one talent tree and one class.
import { readFileSync, readdirSync, writeFileSync } from 'node:fs';
import { dirname, join, resolve } from 'node:path';
import { fileURLToPath } from 'node:url';

const root = resolve(dirname(fileURLToPath(import.meta.url)), '..');
const build = resolve(process.argv[2] ?? join(root, 'data/builds/1.60.1.70450'));
const read = (p) => JSON.parse(readFileSync(p, 'utf8'));
const title = (slug) => slug.replace(/(^|-)([a-z])/g, (_, s, c) => (s ? ' ' : '') + c.toUpperCase());

const specs = read(join(root, 'data/curated/specs.json'));
const specName = (slug, index) => specs.find((s) => s.class_slug === slug && s.tree_index === index)?.name;

const owners = new Map(); // id -> Set(class slug)
const own = (id, slug) => {
  if (!id) return;
  if (!owners.has(id)) owners.set(id, new Set());
  owners.get(id).add(slug);
};
for (const [slug, byName] of Object.entries(read(join(build, 'spellranks.json')).classes))
  for (const ranks of Object.values(byName)) for (const r of ranks) own(r.id, slug);

const treeOf = new Map(); // id -> Set("slug/spec")
for (const f of readdirSync(join(build, 'talents'))) {
  const t = read(join(build, 'talents', f));
  t.trees.forEach((tree, i) => {
    const name = specName(t.class_slug, tree.position ?? i) ?? tree.name;
    for (const ta of tree.talents)
      for (const id of new Set([ta.spell_id, ...ta.ranks.map((r) => r.spell_id)])) {
        if (!id) continue;
        own(id, t.class_slug);
        if (!treeOf.has(id)) treeOf.set(id, new Set());
        treeOf.get(id).add(`${t.class_slug}/${name}`);
      }
  });
}

const classes = {};
const spec = {};
for (const [id, set] of [...owners].sort((a, b) => a[0] - b[0])) {
  if (set.size !== 1) continue;
  const slug = [...set][0];
  (classes[title(slug)] ??= []).push(id);
  const trees = treeOf.get(id);
  if (trees && trees.size === 1) {
    const [, name] = [...trees][0].split('/');
    ((spec[title(slug)] ??= {})[name] ??= []).push(id);
  }
}
const out = { build: build.split(/[\/]/).pop(), classes, spec };
writeFileSync(join(root, 'api/internal/engine/classspells.json'), JSON.stringify(out) + '\n');
console.log(
  `${Object.values(classes).flat().length} class spells, ${Object.values(spec).flatMap((m) => Object.values(m)).flat().length} spec spells`,
);
