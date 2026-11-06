// web/tests/e2e/support/parity-data-paths.mjs
// The files of data/builds/<active build> that the simulator parity spec serves to the fixture
// build (support/sim-parity.ts: the mocks are drawn from them), and the inputs its derived names
// read (scripts/sync-data.mjs `writeSimNames`). The CI parity job checks out only these
// (`git sparse-checkout add $(node tests/e2e/support/parity-data-paths.mjs)`), so the job keeps
// its light checkout and the spec reads the same files it reads locally.
import { readFileSync } from 'node:fs';

const FILES = [
  // Served to the page.
  'classes.json',
  'combos.json',
  'enchants.json',
  'icons/',
  'items/',
  'races.json',
  'sets.json',
  'simbuffs.json',
  'spellconst/',
  'spellicons.json',
  'stat-weights.json',
  'talents/',
  // Read by writeSimNames for the derived simnames files.
  'items.json',
  'spells.json',
];

const { build } = JSON.parse(
  readFileSync(new URL('../../../src/data/active-build.json', import.meta.url), 'utf8'),
);
console.log(FILES.map((file) => `/data/builds/${build}/${file}`).join('\n'));
