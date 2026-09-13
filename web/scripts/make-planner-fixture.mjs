// web/scripts/make-planner-fixture.mjs
// Regenerates web/src/fixtures/planner/ — a small, contract-shaped stand-in for one
// data/builds/<build>/ directory, used until the data/ plan emits the real per-class files.
// Its contents are NOT a source of truth about the game: the talent names and numbers are
// invented, and every fact the site publishes comes from the pipeline, not from here.
// Run with: node scripts/make-planner-fixture.mjs   (from web/)
import { createHash } from 'node:crypto';
import { mkdir, rm, writeFile } from 'node:fs/promises';
import path from 'node:path';
import { fileURLToPath } from 'node:url';

const HERE = path.dirname(fileURLToPath(import.meta.url));
const OUT = path.resolve(HERE, '../src/fixtures/planner');
const BUILD = '1.15.9.69722';
const PRODUCT = 'wow_classic_era';

// A 1x1 fully transparent WebP. The fixture only needs the icon files to exist and decode;
// the real pipeline writes 64x64 icons to the same paths.
const BLANK_WEBP = Buffer.from('UklGRhoAAABXRUJQVlA4TA0AAAAvAAAAEAcQERGIiP4HAA==', 'base64');

const SOURCES = {
  blizzcon: {
    label: 'Blizzard, World of Warcraft at BlizzCon 2026',
    url: 'https://news.blizzard.com/en-us/article/24301145/world-of-warcraft-at-blizzcon-2026-discover-whats-next',
    kind: 'blizzard',
  },
  skyborne: {
    label: 'Wowhead, Skyborne First Look',
    url: 'https://www.wowhead.com/forever/news/skyborne-first-look-new-neutral-race-in-world-of-warcraft-forever-382829',
    kind: 'community',
  },
};

const CLASSES = [
  [1, 'Warrior', 'warrior', '#c69b6d', 'Skyborne can be Warriors.'],
  [2, 'Paladin', 'paladin', '#f48cba', 'Undead can be Paladins.'],
  [3, 'Hunter', 'hunter', '#aad372', 'Skyborne can be Hunters.'],
  [4, 'Rogue', 'rogue', '#fff468', 'Skyborne can be Rogues.'],
  [5, 'Priest', 'priest', '#ffffff', 'Race-specific Priest spells are unchanged.'],
  [7, 'Shaman', 'shaman', '#3f8fe0', 'Shaman remains Horde-only.'],
  [8, 'Mage', 'mage', '#3fc7eb', 'Mage race list is unchanged.'],
  [9, 'Warlock', 'warlock', '#8788ee', 'Warlock race list is unchanged.'],
  [11, 'Druid', 'druid', '#ff7c0a', 'Skyborne can be Druids, with sky-blue forms.'],
];

const RACES = [
  [1, 'Human', 'human', 'alliance', 'Human racials are unchanged.', false],
  [2, 'Orc', 'orc', 'horde', 'Orc racials are unchanged.', false],
  [3, 'Dwarf', 'dwarf', 'alliance', 'Dwarf racials are unchanged.', false],
  [4, 'Night Elf', 'night-elf', 'alliance', 'Night Elf racials are unchanged.', false],
  [5, 'Undead', 'undead', 'horde', 'Undead gain access to Paladin.', false],
  [6, 'Tauren', 'tauren', 'horde', 'Tauren racials are unchanged.', false],
  [7, 'Gnome', 'gnome', 'alliance', 'Gnome racials are unchanged.', false],
  [8, 'Troll', 'troll', 'horde', 'Troll racials are unchanged.', false],
  [11, 'Skyborne', 'skyborne', 'neutral', 'New neutral race; faction is chosen at character creation.', true],
];

// Classic Era combinations, by class slug.
const ERA_COMBOS = {
  warrior: ['human', 'orc', 'dwarf', 'night-elf', 'undead', 'tauren', 'gnome', 'troll'],
  paladin: ['human', 'dwarf'],
  hunter: ['orc', 'dwarf', 'night-elf', 'tauren', 'troll'],
  rogue: ['human', 'orc', 'dwarf', 'night-elf', 'undead', 'gnome', 'troll'],
  priest: ['human', 'dwarf', 'night-elf', 'undead', 'troll'],
  shaman: ['orc', 'tauren', 'troll'],
  mage: ['human', 'undead', 'gnome', 'troll'],
  warlock: ['human', 'orc', 'undead', 'gnome'],
  druid: ['night-elf', 'tauren'],
};

// Added by Forever, by class slug.
const NEW_COMBOS = {
  paladin: ['undead'],
  warrior: ['skyborne'],
  hunter: ['skyborne'],
  rogue: ['skyborne'],
  druid: ['skyborne'],
};

// [name, max_rank, tier, column, prereq_talent_id, prereq_rank, description template, step]
const TREES = [
  [
    161,
    'Arms',
    0,
    [
      [
        1001,
        'Improved Heroic Strike',
        3,
        0,
        0,
        null,
        null,
        'Reduces the rage cost of Heroic Strike by {n}.',
        1,
      ],
      [1002, 'Deflection', 5, 0, 1, null, null, 'Increases your chance to parry by {n}%.', 1],
      [1003, 'Improved Rend', 3, 1, 0, null, null, 'Increases the damage of Rend by {n}%.', 5],
      [1004, 'Tactical Mastery', 5, 1, 1, 1002, 2, 'Retains up to {n} rage when you change stances.', 5],
      [
        1005,
        'Sweeping Strikes',
        1,
        2,
        0,
        1003,
        3,
        'Your next {n} melee attacks strike an additional nearby target.',
        5,
      ],
      [1006, 'Impale', 2, 2, 1, null, null, 'Increases critical strike damage by {n}%.', 10],
      [
        1007,
        'Axe Specialization',
        5,
        3,
        0,
        null,
        null,
        'Increases your chance to critically hit with axes by {n}%.',
        1,
      ],
    ],
  ],
  [
    164,
    'Fury',
    1,
    [
      [2001, 'Booming Voice', 5, 0, 0, null, null, 'Increases the duration of your shouts by {n}%.', 10],
      [2002, 'Cruelty', 5, 0, 1, null, null, 'Increases your chance to critically hit by {n}%.', 1],
      [
        2003,
        'Unbridled Wrath',
        5,
        1,
        0,
        null,
        null,
        'Gives you a {n}% chance to generate an extra rage point on a melee hit.',
        8,
      ],
      [
        2004,
        'Improved Battle Shout',
        5,
        1,
        1,
        2001,
        1,
        'Increases the attack power of Battle Shout by {n}%.',
        5,
      ],
      [
        2005,
        'Enrage',
        5,
        2,
        0,
        null,
        null,
        'Increases your melee damage by {n}% after being critically hit.',
        3,
      ],
      [
        2006,
        'Death Wish',
        1,
        2,
        1,
        2003,
        3,
        'Increases your physical damage by {n}% and raises damage taken by 20%.',
        20,
      ],
      [
        2007,
        'Flurry',
        5,
        3,
        0,
        null,
        null,
        'Increases your attack speed by {n}% after a critical strike.',
        5,
      ],
    ],
  ],
];

const ITEMS = [
  [
    12640,
    'Lionheart Helm',
    'fixture_item_lionheart_helm',
    'head',
    4,
    60,
    65,
    565,
    { strength: 18, crit: 2, hit: 2 },
    null,
    false,
  ],
  [
    16963,
    'Helm of Wrath',
    'fixture_item_helm_of_wrath',
    'head',
    4,
    60,
    76,
    610,
    { strength: 25, stamina: 20 },
    550,
    false,
  ],
  [
    16966,
    'Spaulders of Wrath',
    'fixture_item_spaulders_of_wrath',
    'shoulder',
    4,
    60,
    76,
    500,
    { strength: 20, stamina: 18 },
    550,
    false,
  ],
  [
    12784,
    'Arcanite Reaper',
    'fixture_item_arcanite_reaper',
    'main_hand',
    4,
    60,
    63,
    0,
    { attack_power: 62 },
    null,
    false,
  ],
  [
    19325,
    'Band of Accuria',
    'fixture_item_band_of_accuria',
    'finger',
    4,
    60,
    78,
    0,
    { strength: 19, hit: 2 },
    null,
    true,
  ],
  [
    13968,
    'Snakestone Charm',
    'fixture_item_snakestone_charm',
    'trinket',
    3,
    55,
    60,
    0,
    { defense: 7 },
    null,
    false,
  ],
];

const SETS = [
  {
    id: 550,
    name: 'Battlegear of Wrath',
    item_ids: [16963, 16966],
    bonuses: [{ pieces: 2, description: 'Increases your chance to parry an attack by 1%.' }],
  },
];

const iconFor = (name) =>
  `fixture_${name
    .toLowerCase()
    .replace(/[^a-z0-9]+/g, '_')
    .replace(/^_|_$/g, '')}`;

function buildTalents() {
  return {
    build: BUILD,
    class_id: 1,
    class_slug: 'warrior',
    trees: TREES.map(([id, name, position, rows]) => ({
      id,
      name,
      position,
      talents: rows.map(([tid, tname, maxRank, tier, column, prereqId, prereqRank, template, step]) => ({
        id: tid,
        name: tname,
        icon: iconFor(tname),
        max_rank: maxRank,
        tier,
        column,
        prereq_talent_id: prereqId,
        prereq_rank: prereqRank,
        ranks: Array.from({ length: maxRank }, (_, r) => ({
          spell_id: tid * 10 + r + 1,
          description: template.replace('{n}', String(step * (r + 1))),
        })),
      })),
    })),
  };
}

function buildReference() {
  const classes = CLASSES.map(([id, name, slug, color, text]) => ({
    id,
    name,
    slug,
    color,
    forever_changes: [{ text, sources: [slug === 'paladin' ? SOURCES.blizzcon : SOURCES.skyborne] }],
  }));
  const races = RACES.map(([id, name, slug, faction, text, placeholder]) => {
    const row = {
      id,
      name,
      slug,
      faction,
      forever_changes: [{ text, sources: [placeholder ? SOURCES.skyborne : SOURCES.blizzcon] }],
    };
    return placeholder ? { ...row, placeholder: true } : row;
  });
  const classId = Object.fromEntries(classes.map((c) => [c.slug, c.id]));
  const raceId = Object.fromEntries(races.map((r) => [r.slug, r.id]));
  const combos = [];
  for (const [slug, raceSlugs] of Object.entries(ERA_COMBOS)) {
    for (const race of raceSlugs) {
      combos.push({ race_id: raceId[race], class_id: classId[slug], new_in_forever: false });
    }
  }
  for (const [slug, raceSlugs] of Object.entries(NEW_COMBOS)) {
    for (const race of raceSlugs) {
      combos.push({ race_id: raceId[race], class_id: classId[slug], new_in_forever: true });
    }
  }
  combos.sort((a, b) => a.class_id - b.class_id || a.race_id - b.race_id);
  return { classes, races, combos };
}

function buildItems() {
  return {
    build: BUILD,
    class_slug: 'warrior',
    items: ITEMS.map(
      ([id, name, icon, slot, quality, requiredLevel, itemLevel, armor, stats, setId, isUnique]) => ({
        id,
        name,
        icon,
        slot,
        quality,
        required_level: requiredLevel,
        item_level: itemLevel,
        armor,
        stats,
        set_id: setId,
        unique: isUnique,
      }),
    ),
  };
}

async function main() {
  const talents = buildTalents();
  const { classes, races, combos } = buildReference();
  const items = buildItems();

  await rm(OUT, { recursive: true, force: true });
  await mkdir(path.join(OUT, 'talents'), { recursive: true });
  await mkdir(path.join(OUT, 'items'), { recursive: true });
  await mkdir(path.join(OUT, 'icons'), { recursive: true });

  const written = new Map();
  const writeJson = async (rel, value) => {
    const body = `${JSON.stringify(value, null, 2)}\n`;
    await writeFile(path.join(OUT, rel), body, 'utf8');
    written.set(rel, createHash('sha256').update(body).digest('hex'));
  };

  await writeJson('talents/warrior.json', talents);
  await writeJson('items/warrior.json', items);
  await writeJson('sets.json', SETS);
  await writeJson('classes.json', classes);
  await writeJson('races.json', races);
  await writeJson('combos.json', combos);

  const icons = new Set([
    ...talents.trees.flatMap((t) => t.talents.map((x) => x.icon)),
    ...items.items.map((i) => i.icon),
  ]);
  for (const icon of [...icons].sort()) {
    const rel = `icons/${icon}.webp`;
    await writeFile(path.join(OUT, rel), BLANK_WEBP);
    written.set(rel, createHash('sha256').update(BLANK_WEBP).digest('hex'));
  }

  const manifest = {
    build: BUILD,
    product: PRODUCT,
    fetched_at: '2026-09-13T00:00:00Z',
    fixture: true,
    files: Object.fromEntries([...written.entries()].sort(([a], [b]) => a.localeCompare(b))),
  };
  await writeFile(path.join(OUT, 'manifest.json'), `${JSON.stringify(manifest, null, 2)}\n`, 'utf8');

  console.log(`wrote ${written.size + 1} fixture files to ${path.relative(process.cwd(), OUT)}`);
}

await main();
