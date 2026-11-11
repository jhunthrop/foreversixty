// web/tests/e2e/support/parity-data-paths.mjs
// The files of data/builds/<active build> that the simulator parity spec serves to the fixture
// build (support/sim-parity.ts) and that design/mocks/ws_sim.py draws its mocks from at test time
// (support/parity-mocks.ts), the files design/mocks/guides_facts.mjs reads for the guides parity
// mocks, the planner's tree art (support/planner-page.ts), and the inputs the derived names read (scripts/sync-data.mjs
// `writeSimNames`). The CI parity job checks out only these
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
  // Read by the simulator tool pages (parity-sim-tools.spec.ts): the items the engine knows,
  // which item search offers, and the random-suffix table the candidate rows read.
  'simitems.json',
  'suffixes.json',
  // Read by writeSimNames for the derived simnames files (spells.json by ws_sim.py too).
  'items.json',
  'spells.json',
  // Read by ws_sim.py alone.
  'racialicons.json',
  'simconsumes.json',
  'spelltips/',
  // Read by design/mocks/guides_facts.mjs (the guides parity spec's test-time mocks) through the
  // guide libs: the rotation lines, the ranked BiS files, the loot, the spell ranks, the trainables.
  'addon-data.json',
  'bis/',
  'loot.json',
  // The race portraits GuideRaceList and GuideRacesLane lead each race with.
  'raceicons.json',
  'spellranks.json',
  'trainables/',
  // Read by design/mocks/planner_facts.mjs and served by support/planner-page.ts (the planner
  // parity spec): the client's tree art behind each talent grid.
  'trees/',
];

const { build } = JSON.parse(
  readFileSync(new URL('../../../src/data/active-build.json', import.meta.url), 'utf8'),
);
console.log(FILES.map((file) => `/data/builds/${build}/${file}`).join('\n'));
