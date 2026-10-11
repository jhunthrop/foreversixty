// web/src/lib/rankings/url.ts
// Spec section 6: "Filters: boss, region, ruleset, faction, class and spec, phase, item
// level, date; all in the URL." The boss is the path; everything else is here. Parsing is
// total and serialising omits defaults, exactly as the report's own URL state does.
import { REGIONS, RULESETS } from '../characters';
import { PHASES } from './phases';

export const RANKING_METRICS = [
  { id: 'dps', label: 'Damage' },
  { id: 'hps', label: 'Healing' },
  { id: 'damage_taken', label: 'Damage taken' },
  { id: 'execution', label: 'Execution' },
] as const;

export const GUILD_KINDS = [
  { id: 'progress', label: 'Progress' },
  { id: 'speed', label: 'Speed' },
  { id: 'execution', label: 'Execution' },
] as const;

export const FACTIONS = ['alliance', 'horde'] as const;
export const SINCE = ['today'] as const;

export interface RankingsState {
  board: 'character' | 'guild';
  metric: string;
  kind: string;
  spec: string;
  class: string;
  phase: string;
  region: string;
  ruleset: string;
  faction: string;
  /** Empty is all time; `today` is spec section 6's "today versus all-time". */
  since: string;
  page: number;
}

export function defaultRankingsState(): RankingsState {
  return {
    board: 'character',
    metric: 'dps',
    kind: 'progress',
    spec: '',
    class: '',
    phase: '',
    region: '',
    ruleset: '',
    faction: '',
    since: '',
    page: 1,
  };
}

/** `?phase=all`: an explicit "every phase" (and likewise for ruleset, class and spec), so a
 *  reload keeps the choice instead of falling back to the viewer's own default. */
export const EXPLICIT_ALL = 'all';

/** The filters whose default depends on who is looking, and so need an explicit "all". */
export const AIMED_KEYS = ['class', 'spec', 'phase', 'ruleset'] as const;

/** What the page starts from for this viewer: the aimed filters plus the metric. */
export type RankingsDefaults = Pick<RankingsState, (typeof AIMED_KEYS)[number] | 'metric'>;

/**
 * Whether a board needs an encounter to answer at all: the Characters board always ranks
 * one encounter, and so does the Guilds board's speed and execution kinds (`Store.
 * GuildRankings` in the API 400s without one), but the Guilds board's progress kind
 * ranks every encounter a guild has cleared and needs none. Read by Rankings.svelte to
 * decide, with no encounter to default to, whether to show the encounter picker instead
 * of asking the API for a ranking it cannot answer.
 */
export function requiresEncounter(state: RankingsState): boolean {
  return state.board === 'character' || state.kind !== 'progress';
}

function pick(value: string | null, allowed: readonly string[], fallback: string): string {
  return value !== null && allowed.includes(value) ? value : fallback;
}

function pickOrAll(value: string | null, allowed: readonly string[]): string {
  return value === EXPLICIT_ALL ? '' : pick(value, allowed, '');
}

/** Free text, except that `all` means none: the API owns the per-class spec table, and a spec
 *  the site has not heard of yet must still be linkable. */
function freeTextOrAll(value: string | null): string {
  return value === null || value === EXPLICIT_ALL ? '' : value;
}

export function parseRankingsState(search: string): RankingsState {
  const params = new URLSearchParams(search);
  const state = defaultRankingsState();

  state.board = pick(params.get('board'), ['character', 'guild'], 'character') as RankingsState['board'];
  state.metric = pick(
    params.get('metric'),
    RANKING_METRICS.map((metric) => metric.id),
    state.metric,
  );
  state.kind = pick(
    params.get('kind'),
    GUILD_KINDS.map((kind) => kind.id),
    state.kind,
  );
  state.phase = pickOrAll(
    params.get('phase'),
    PHASES.map((phase) => phase.id),
  );
  state.region = pick(params.get('region'), REGIONS, '');
  state.ruleset = pickOrAll(
    params.get('ruleset'),
    RULESETS.map((ruleset) => ruleset.id),
  );
  state.faction = pick(params.get('faction'), FACTIONS, '');
  state.since = pick(params.get('since'), SINCE, '');
  state.spec = freeTextOrAll(params.get('spec'));
  state.class = freeTextOrAll(params.get('class'));

  const page = params.get('page');
  if (page !== null && /^\d+$/.test(page)) state.page = Math.max(1, Number.parseInt(page, 10));

  return state;
}

/**
 * The URL for a state. Only what differs from `defaults` is written, so a viewer's own board
 * stays a short link; an aimed filter the viewer set to "all" is written as `all`, because
 * leaving it out would mean "your default".
 */
export function rankingsSearch(
  state: RankingsState,
  defaults: RankingsDefaults = defaultRankingsState(),
): string {
  const base = defaultRankingsState();
  const params = new URLSearchParams();
  if (state.board !== base.board) params.set('board', state.board);
  if (state.board === 'guild' && state.kind !== base.kind) params.set('kind', state.kind);
  if (state.board === 'character' && state.metric !== defaults.metric) params.set('metric', state.metric);
  for (const key of AIMED_KEYS) {
    if (state[key] === defaults[key]) continue;
    params.set(key, state[key] === '' ? EXPLICIT_ALL : state[key]);
  }
  for (const key of ['region', 'faction', 'since'] as const) {
    if (state[key] !== '') params.set(key, state[key]);
  }
  if (state.page !== 1) params.set('page', String(state.page));
  const text = params.toString();
  return text === '' ? '' : `?${text}`;
}
