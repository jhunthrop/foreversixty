// web/src/lib/rankings/snapshot.ts
// The rankings' build-time snapshot (seo-fix 2026-11-10: the boards were API-only, so /rankings and
// /rankings/<encounter> server-rendered no row at all). scripts/sync-rankings.mjs writes
// src/data/generated/rankings-snapshot.json on every build from the public API: the ranked
// encounters, and the first rows of each one's every-class DPS board in the current phase, the
// board a signed-out visitor opens. The deploy the nightly BiS run dispatches rebuilds it, so the
// rows are at most a day old. The pages hand the island only the part it draws (`seedFor`), and the
// island shows those rows until the live board answers.
import type { RankedEncounter, RankingsPage } from './api';
import type { RankingsState } from './url';

/** The rows kept per board: a board's first screen, and every spec's best on today's boards. */
export const SNAPSHOT_ROWS = 25;

export interface RankingsSnapshot {
  /** When the snapshot was read; null for the empty snapshot. */
  read_at: string | null;
  /** The phase the boards were read in. */
  phase: string;
  /** Ranked encounters, most kills first. */
  encounters: RankedEncounter[];
  /** Each encounter's every-class DPS board, its first rows, by slug. */
  boards: Record<string, RankingsPage>;
  /** The landing's healer leaders: the leader encounter's every-class HPS board. */
  hps: { slug: string; page: RankingsPage } | null;
}

/** What the island draws from the snapshot before the live boards answer. */
export interface RankingsSeed {
  phase: string;
  encounters: RankedEncounter[];
  boards: Record<string, RankingsPage>;
  hps: { slug: string; page: RankingsPage } | null;
}

export const EMPTY_SNAPSHOT: RankingsSnapshot = {
  read_at: null,
  phase: '',
  encounters: [],
  boards: {},
  hps: null,
};

/** A board cut to its first `rows` rows; the total stays the board's own. */
export function trimBoard(page: RankingsPage, rows: number = SNAPSHOT_ROWS): RankingsPage {
  return { ...page, rows: page.rows.slice(0, rows) };
}

function boardsOf(
  snapshot: RankingsSnapshot,
  slugs: readonly string[],
  rows?: number,
): Record<string, RankingsPage> {
  return Object.fromEntries(
    slugs.flatMap((slug) => {
      const board = snapshot.boards[slug];
      return board === undefined ? [] : [[slug, rows === undefined ? board : trimBoard(board, rows)]];
    }),
  );
}

/**
 * The part of the snapshot one page draws: for the landing (`slug` empty) the encounters, the
 * top rows of the encounters it shows and the leader encounter's whole boards; for an encounter's
 * page that encounter's board. Null when the snapshot has nothing for the page.
 */
export function seedFor(
  snapshot: RankingsSnapshot,
  slug: string,
  landing: { shown: number; topRows: number; leader: string | null },
): RankingsSeed | null {
  if (snapshot.encounters.length === 0) return null;
  const base = { phase: snapshot.phase, encounters: snapshot.encounters, hps: null };
  if (slug !== '') {
    const board = snapshot.boards[slug];
    return board === undefined ? null : { ...base, boards: { [slug]: board } };
  }
  const shown = snapshot.encounters.slice(0, landing.shown).map((encounter) => encounter.slug);
  const leader = landing.leader;
  return {
    ...base,
    boards: {
      ...boardsOf(snapshot, shown, landing.topRows),
      ...(leader === null ? {} : boardsOf(snapshot, [leader])),
    },
    hps: leader !== null && snapshot.hps?.slug === leader ? snapshot.hps : null,
  };
}

/** The encounter slugs a real build prerenders a page for: every one the snapshot holds a board for. */
export function snapshotSlugs(snapshot: RankingsSnapshot): string[] {
  return snapshot.encounters
    .map((encounter) => encounter.slug)
    .filter((slug) => snapshot.boards[slug] !== undefined);
}

/**
 * The seed's board for `slug` when the board on screen is the one the snapshot read: every class
 * and spec, DPS, page one, no other filter, in the snapshot's phase (or before the island has set
 * the phase at all). Any other board waits for the API.
 */
export function seedBoard(
  seed: RankingsSeed | null,
  slug: string,
  state: RankingsState,
): RankingsPage | null {
  if (seed === null) return null;
  const unfiltered =
    state.board === 'character' &&
    state.metric === 'dps' &&
    state.class === '' &&
    state.spec === '' &&
    state.region === '' &&
    state.ruleset === '' &&
    state.faction === '' &&
    state.since === '' &&
    state.page === 1 &&
    (state.phase === '' || state.phase === seed.phase);
  return unfiltered ? (seed.boards[slug] ?? null) : null;
}
