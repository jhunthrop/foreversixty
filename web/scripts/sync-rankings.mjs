// web/scripts/sync-rankings.mjs
// The rankings snapshot (src/lib/rankings/snapshot.ts) -> src/data/generated/rankings-snapshot.json,
// before dev, check, test and build. seo-fix 2026-11-10: /rankings and /rankings/<encounter> drew no
// row until the island asked the API, so a crawler read 17 words. With the snapshot the pages
// server-render the encounters and each board's first rows, and the real build prerenders a page per
// ranked encounter (src/worker.ts serves it in place of the bare shell).
//
// FOREVER_DATA=fixture (tests, Playwright, Lighthouse) copies src/fixtures/rankings/snapshot.json, so
// every suite reads fixed rows. Otherwise it reads the public API (GET /v1/rankings/encounters, then
// each encounter's every-class DPS board in the current phase, one character per row, and the leader
// encounter's HPS board), the same reads a signed-out visitor's island makes. Production builds on
// every deploy, and the nightly BiS run dispatches one, so the rows are at most a day old.
//
// The API being unreachable never fails a build: the pages then draw what they drew before the
// snapshot existed (the island's skeleton until the API answers), and the build says so loudly. The
// file is gitignored like everything else under src/data/generated; it is never committed.
import { copyFile, mkdir, readFile, writeFile } from 'node:fs/promises';
import path from 'node:path';
import process from 'node:process';
import { fileURLToPath, pathToFileURL } from 'node:url';

const webRoot = path.join(path.dirname(fileURLToPath(import.meta.url)), '..');
export const TARGET = path.join(webRoot, 'src', 'data', 'generated', 'rankings-snapshot.json');
const FIXTURE = path.join(webRoot, 'src', 'fixtures', 'rankings', 'snapshot.json');
const PHASES = path.join(webRoot, 'src', 'data', 'phases.json');
const DEFAULT_API = 'https://api.foreversixty.gg';
/** lib/rankings/snapshot.ts SNAPSHOT_ROWS: a board's first screen. */
export const SNAPSHOT_ROWS = 25;
/** The most encounters read (and prerendered): the busiest, as the API orders them. */
export const MAX_ENCOUNTERS = 40;
/** lib/rankings/landing.ts RAID_SIZE: the leader encounter is the busiest raid boss. */
const RAID_SIZE = 10;
const TIMEOUT_MS = 10_000;

/** The phase `now` falls in (lib/rankings/phases.ts phaseAt, read from the same phases.json). */
export function phaseAt(phases, now) {
  let found = phases[0]?.name ?? '';
  for (const phase of phases) if (Date.parse(phase.start) <= now.getTime()) found = phase.name;
  return found;
}

/** The encounter the landing reads its spec leaders on (lib/rankings/landing.ts leaderEncounter). */
export function leaderSlug(encounters) {
  return (encounters.find((encounter) => (encounter.size ?? 0) >= RAID_SIZE) ?? encounters[0])?.slug ?? null;
}

async function read(apiBase, pathAndQuery, fetchImpl) {
  const response = await fetchImpl(`${apiBase}${pathAndQuery}`, { signal: AbortSignal.timeout(TIMEOUT_MS) });
  if (!response.ok) throw new Error(`${pathAndQuery}: HTTP ${response.status}`);
  const envelope = await response.json();
  if (!envelope?.ok || envelope.data === null)
    throw new Error(`${pathAndQuery}: ${envelope?.error?.message ?? 'no data'}`);
  return envelope.data;
}

function boardQuery(slug, metric, phase) {
  const params = new URLSearchParams({ encounter: slug, metric, phase, distinct: 'character', page: '1' });
  return `/v1/rankings?${params}`;
}

/** The first rows, without a character who did nothing on the board (0 DPS, a healer on the DPS
 *  board): a 0 with a percentile beside it reads as a broken row. */
function trim(page) {
  const active = page.rows.filter((row) => row.value > 0);
  return {
    ...page,
    rows: active.slice(0, SNAPSHOT_ROWS),
    total: page.total - (page.rows.length - active.length),
  };
}

/**
 * Reads the snapshot from the API. A board that fails is left out (its page falls back to the
 * shell); the encounter list failing is the caller's to report.
 * @param {{
 *   apiBase: string,
 *   now: Date,
 *   phases: { name: string, start: string }[],
 *   fetchImpl?: (url: string, init?: RequestInit) => Promise<Response>,
 *   warn?: (line: string) => void,
 * }} options
 * @returns {Promise<{ read_at: string, phase: string, encounters: any[], boards: Record<string, any>, hps: any }>}
 */
export async function readSnapshot({ apiBase, now, phases, fetchImpl = fetch, warn = console.warn }) {
  const phase = phaseAt(phases, now);
  const { encounters: all } = await read(apiBase, '/v1/rankings/encounters', fetchImpl);
  const encounters = all.slice(0, MAX_ENCOUNTERS);
  const boards = {};
  for (const encounter of encounters) {
    try {
      boards[encounter.slug] = trim(await read(apiBase, boardQuery(encounter.slug, 'dps', phase), fetchImpl));
    } catch (error) {
      warn(`sync-rankings: ${encounter.slug} left out (${error.message})`);
    }
  }
  const leader = leaderSlug(encounters);
  let hps = null;
  if (leader !== null) {
    try {
      hps = { slug: leader, page: trim(await read(apiBase, boardQuery(leader, 'hps', phase), fetchImpl)) };
    } catch (error) {
      warn(`sync-rankings: ${leader} healing board left out (${error.message})`);
    }
  }
  return { read_at: now.toISOString(), phase, encounters, boards, hps };
}

export const EMPTY_SNAPSHOT = { read_at: null, phase: '', encounters: [], boards: {}, hps: null };

async function main() {
  await mkdir(path.dirname(TARGET), { recursive: true });
  if (process.env.FOREVER_DATA === 'fixture') {
    await copyFile(FIXTURE, TARGET);
    console.log('sync-rankings: the fixture snapshot');
    return;
  }
  const apiBase = process.env.RANKINGS_SNAPSHOT_API ?? DEFAULT_API;
  const phases = JSON.parse(await readFile(PHASES, 'utf8'));
  let snapshot;
  try {
    snapshot = await readSnapshot({ apiBase, now: new Date(), phases });
  } catch (error) {
    // GitHub's annotation syntax: the deploy log flags it without failing the deploy.
    console.warn(
      `::warning::sync-rankings: ${apiBase} did not answer (${error.message}); the rankings pages ship without server rows`,
    );
    snapshot = EMPTY_SNAPSHOT;
  }
  await writeFile(TARGET, `${JSON.stringify(snapshot)}\n`);
  console.log(
    `sync-rankings: ${snapshot.encounters.length} encounters, ${Object.keys(snapshot.boards).length} boards from ${apiBase}`,
  );
}

if (import.meta.url === pathToFileURL(process.argv[1] ?? '').href) {
  await main();
}
