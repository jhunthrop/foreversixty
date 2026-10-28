// web/tests/captures/premium-examples/capture.spec.ts
// Not a test of behaviour: it produces the premium example's assets from one real run. It opens
// /sim/{SIM_ID} at 440 x 900 CSS px (DPR 2), clips a 440 x 330 box from the saved title down to
// web/public/premium/example-server-run.{webp,avif} (880 x 660), and writes
// web/src/data/premium-examples.json from the sim response. It fails unless the result answers 200
// signed out and is a finished server Top Gear run, and unless the written file parses.
// SIM_ID and SIM_RAN_AT (the owner's /mine row finished_at) are required; API_BASE defaults to
// production; PRESET_BUILD names the site build when the run used the preset gear.
// The crop shows what a player recognises as a real result: the saved title band and the top of
// the character strip (name and descriptor: spec, level), never the gear grid or the swaps table.
// The numbers (combinations, ranAt, preset) are not in the image: they go into the data file and
// the card's caption (serverRunCaption) renders them.
// Depends on these test ids (src/components/sim): sim-saved-title (SavedSim.svelte, the clip's top
// edge); sim-character, sim-character-descriptor and sim-character-name (CharacterStrip.svelte;
// the strip's top, the descriptor and, when present, the name must lie inside the clip).
import { mkdirSync, writeFileSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { expect, test, type Page } from '@playwright/test';
import sharp from 'sharp';
import { parseServerRunExample } from '../../../src/lib/billing/premium-examples';
import { specLabel } from '../../../src/lib/sim/spec-label';
import type { SimResult } from '../../../src/lib/sim/types';

interface Rect {
  x: number;
  y: number;
  width: number;
  height: number;
}

const HERE = path.dirname(fileURLToPath(import.meta.url));
const WEB_ROOT = path.resolve(HERE, '../../..');
const IMAGE_BASE = path.join(WEB_ROOT, 'public/premium/example-server-run');
const DATA_FILE = path.join(WEB_ROOT, 'src/data/premium-examples.json');

const VIEWPORT = { width: 440, height: 900 };
const CLIP = { x: 0, y: 0, width: 440, height: 330 };
const PIXEL_RATIO = 2;
const AVIF_MAX_BYTES = 45 * 1024;
const AVIF_START_QUALITY = 60;
const AVIF_MIN_QUALITY = 20;
const AVIF_QUALITY_STEP = 5;
const WEBP_QUALITY = 82;
const SAVED_TITLE_TEST_ID = 'sim-saved-title';
const CHARACTER_TEST_ID = 'sim-character';
const NAME_TEST_ID = 'sim-character-name';
const DESCRIPTOR_TEST_ID = 'sim-character-descriptor';
const SETTLE_MS = 800;
const TOP_GEAR_MODE = 'gear';
const DEFAULT_API_BASE = 'https://api.foreversixty.gg';
/** Gear sources that are the player's own character; a site build is the preset. */
const OWN_GEAR_SOURCES: readonly string[] = ['armory', 'addon', 'blizzard', 'manual', 'fight'];
const PRESET_SOURCE = 'build';

function requireEnv(name: string): string {
  const value = process.env[name];
  if (value === undefined || value === '') throw new Error(`${name} is required`);
  return value;
}

async function fetchRun(simId: string, apiBase: string): Promise<SimResult> {
  const res = await fetch(`${apiBase}/v1/sims/${encodeURIComponent(simId)}`);
  if (res.status !== 200) throw new Error(`GET /v1/sims/${simId} answered ${res.status}, expected 200`);
  const body = (await res.json()) as { data?: SimResult };
  if (body.data === undefined) throw new Error('the sim response has no `data` envelope');
  return body.data;
}

function assertFinishedServerTopGear(result: SimResult): void {
  const problems: string[] = [];
  if (result.lane !== 'server') problems.push(`lane is "${result.lane}", not server`);
  if (result.request?.bulk?.mode !== TOP_GEAR_MODE) problems.push('not a Top Gear (gear mode) run');
  if (result.error !== undefined && result.error !== '') problems.push(`run failed: ${result.error}`);
  if (result.aborted === true) problems.push('run was aborted');
  if ((result.combos ?? []).length === 0) problems.push('no ranked combinations: not finished');
  if (problems.length > 0)
    throw new Error(`SIM_ID is not a finished server Top Gear run: ${problems.join('; ')}`);
}

/** With PRESET_BUILD set the run must be that site build; without it the run must be own gear. */
function isPresetGear(result: SimResult, presetBuild: string | undefined): boolean {
  const { kind, ref } = result.request.source;
  if (presetBuild !== undefined) {
    if (kind !== PRESET_SOURCE || ref !== presetBuild) {
      throw new Error(`the run's gear source is ${kind}:${ref}, expected ${PRESET_SOURCE}:${presetBuild}`);
    }
    return true;
  }
  if (OWN_GEAR_SOURCES.includes(kind)) return false;
  throw new Error(`gear source "${kind}" needs PRESET_BUILD to be named as the site's preset`);
}

/** The first stage ranks every combination, so its count is the run's combinations. */
function combinationCount(result: SimResult): number {
  const first = result.stages?.[0]?.combos;
  if (first === undefined || first < 1) throw new Error('the run has no stage with a combination count');
  return first;
}

/** Fails unless the element (or, with `top`, just its top edge) lies inside the clip. */
async function assertInside(
  page: Page,
  testId: string,
  clip: Rect,
  options: { top?: boolean } = {},
): Promise<void> {
  const box = await page.getByTestId(testId).boundingBox();
  if (box === null) throw new Error(`${testId} has no box`);
  const bottom = options.top === true ? box.y : box.y + box.height;
  if (box.y < clip.y || bottom > clip.y + clip.height) {
    throw new Error(`${testId} does not fit inside the ${clip.width} x ${clip.height} clip`);
  }
}

async function encodeAvifUnderBudget(png: Buffer): Promise<Buffer> {
  for (let quality = AVIF_START_QUALITY; quality >= AVIF_MIN_QUALITY; quality -= AVIF_QUALITY_STEP) {
    const avif = await sharp(png).avif({ quality }).toBuffer();
    if (avif.length <= AVIF_MAX_BYTES) return avif;
  }
  throw new Error(`AVIF stays over ${AVIF_MAX_BYTES} bytes down to quality ${AVIF_MIN_QUALITY}`);
}

test('capture the premium server-run example', async ({ browser }) => {
  const simId = requireEnv('SIM_ID');
  const apiBase = process.env.API_BASE ?? DEFAULT_API_BASE;
  const ranAt = requireEnv('SIM_RAN_AT');
  const result = await fetchRun(simId, apiBase);
  assertFinishedServerTopGear(result);
  const preset = isPresetGear(result, process.env.PRESET_BUILD || undefined);

  const context = await browser.newContext({ viewport: VIEWPORT, deviceScaleFactor: PIXEL_RATIO });
  const page = await context.newPage();
  const response = await page.goto(`/sim/${encodeURIComponent(simId)}`, { waitUntil: 'networkidle' });
  expect(response?.status()).toBe(200);
  const title = page.getByTestId(SAVED_TITLE_TEST_ID);
  await title.waitFor();
  await page.getByTestId(CHARACTER_TEST_ID).waitFor();
  await page.waitForTimeout(SETTLE_MS);
  const top = await title.boundingBox();
  if (top === null) throw new Error('the saved result title has no box');
  const clip = { ...CLIP, x: top.x, y: top.y };
  await assertInside(page, CHARACTER_TEST_ID, clip, { top: true });
  await assertInside(page, DESCRIPTOR_TEST_ID, clip);
  if ((await page.getByTestId(NAME_TEST_ID).count()) > 0) await assertInside(page, NAME_TEST_ID, clip);
  const png = await page.screenshot({ clip, type: 'png' });
  await context.close();

  const size = { width: CLIP.width * PIXEL_RATIO, height: CLIP.height * PIXEL_RATIO };
  const meta = await sharp(png).metadata();
  expect({ width: meta.width, height: meta.height }).toEqual(size);

  mkdirSync(path.dirname(IMAGE_BASE), { recursive: true });
  writeFileSync(`${IMAGE_BASE}.webp`, await sharp(png).webp({ quality: WEBP_QUALITY }).toBuffer());
  writeFileSync(`${IMAGE_BASE}.avif`, await encodeAvifUnderBudget(png));

  const serverRun = {
    simId,
    characterName: preset ? '' : result.request.character.name,
    level: result.request.character.level,
    specLabel: specLabel(result.request.spec),
    preset,
    placeholder: false,
    combinations: combinationCount(result),
    ranAt,
  };
  const file = { serverRun };
  expect(() => parseServerRunExample(file)).not.toThrow();
  writeFileSync(DATA_FILE, `${JSON.stringify(file, null, 2)}\n`);
});
