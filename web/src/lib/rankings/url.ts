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

function pick(value: string | null, allowed: readonly string[], fallback: string): string {
  return value !== null && allowed.includes(value) ? value : fallback;
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
  state.phase = pick(
    params.get('phase'),
    PHASES.map((phase) => phase.id),
    '',
  );
  state.region = pick(params.get('region'), REGIONS, '');
  state.ruleset = pick(
    params.get('ruleset'),
    RULESETS.map((ruleset) => ruleset.id),
    '',
  );
  state.faction = pick(params.get('faction'), FACTIONS, '');
  state.since = pick(params.get('since'), SINCE, '');

  // Spec and class are free text: the API owns the per-class spec table, and a spec the
  // site has not heard of yet must still be linkable.
  state.spec = params.get('spec') ?? '';
  state.class = params.get('class') ?? '';

  const page = params.get('page');
  if (page !== null && /^\d+$/.test(page)) state.page = Math.max(1, Number.parseInt(page, 10));

  return state;
}

export function rankingsSearch(state: RankingsState): string {
  const base = defaultRankingsState();
  const params = new URLSearchParams();
  if (state.board !== base.board) params.set('board', state.board);
  if (state.board === 'guild' && state.kind !== base.kind) params.set('kind', state.kind);
  if (state.board === 'character' && state.metric !== base.metric) params.set('metric', state.metric);
  for (const key of ['spec', 'class', 'phase', 'region', 'ruleset', 'faction', 'since'] as const) {
    if (state[key] !== '') params.set(key, state[key]);
  }
  if (state.page !== 1) params.set('page', String(state.page));
  const text = params.toString();
  return text === '' ? '' : `?${text}`;
}
