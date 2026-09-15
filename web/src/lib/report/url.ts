// web/src/lib/report/url.ts
// The report page's whole state lives in the query string, exactly as
// docs/superpowers/specs/2026-09-14-phase-3-interfaces.md specifies:
//   ?fight=<n>&mode=analyze&view=tables&tab=damage-done&source=<guid|friendlies|enemies>&start=<ms>&end=<ms>
// Pasting a URL has to reproduce what the sender was looking at, so parsing is total: an
// unrecognised value falls back to its default rather than throwing, and serialising omits
// anything still at its default so the common link stays short.
//
// Fight numbers come from the engine (fight.Fight.Index), which starts at 1, so the
// default fight is the first entry of report.json rather than a hard-coded 0.

export type Mode = 'analyze' | 'compare' | 'rankings';
export type View = 'tables' | 'timelines' | 'events' | 'queries';
export type Tab =
  | 'summary'
  | 'damage-done'
  | 'damage-taken'
  | 'healing'
  | 'threat'
  | 'buffs'
  | 'debuffs'
  | 'deaths'
  | 'interrupts'
  | 'dispels'
  | 'resources'
  | 'casts';

export interface ModeOption {
  id: Mode | 'mechanics' | 'replay';
  label: string;
  enabled: boolean;
  /** Shown beside a disabled mode. The spec defers these two, it does not drop them. */
  note?: string;
}

export const MODES: readonly ModeOption[] = [
  { id: 'analyze', label: 'Analyze', enabled: true },
  { id: 'compare', label: 'Compare', enabled: true },
  { id: 'rankings', label: 'Rankings', enabled: true },
  { id: 'mechanics', label: 'Mechanics', enabled: false, note: 'later' },
  { id: 'replay', label: 'Replay', enabled: false, note: 'later' },
];

export const VIEWS: readonly { id: View; label: string }[] = [
  { id: 'tables', label: 'Tables' },
  { id: 'timelines', label: 'Timelines' },
  { id: 'events', label: 'Events' },
  { id: 'queries', label: 'Queries' },
];

export const TABS: readonly { id: Tab; label: string }[] = [
  { id: 'summary', label: 'Summary' },
  { id: 'damage-done', label: 'Damage Done' },
  { id: 'damage-taken', label: 'Damage Taken' },
  { id: 'healing', label: 'Healing' },
  { id: 'threat', label: 'Threat' },
  { id: 'buffs', label: 'Buffs' },
  { id: 'debuffs', label: 'Debuffs' },
  { id: 'deaths', label: 'Deaths' },
  { id: 'interrupts', label: 'Interrupts' },
  { id: 'dispels', label: 'Dispels' },
  { id: 'resources', label: 'Resources' },
  { id: 'casts', label: 'Casts' },
];

/** The two scopes that are not a single unit. Anything else is a GUID. */
export const SOURCE_FRIENDLIES = 'friendlies';
export const SOURCE_ENEMIES = 'enemies';

/** The fight number that means every boss pull of the report, and its url spelling. */
export const ALL_FIGHTS = 0;
export const ALL_FIGHTS_PARAM = 'all';

export interface ReportState {
  /** A fight index from report.json, or ALL_FIGHTS for the whole night. */
  fight: number;
  mode: Mode;
  view: View;
  tab: Tab;
  /** `friendlies`, `enemies`, or one unit's GUID. */
  source: string;
  /** Milliseconds from the fight's start; null means the whole fight. */
  start: number | null;
  end: number | null;
}

const ENABLED_MODES = MODES.filter((m) => m.enabled).map((m) => m.id) as Mode[];
const VIEW_IDS = VIEWS.map((v) => v.id);
const TAB_IDS = TABS.map((t) => t.id);

export function defaultState(firstFight: number): ReportState {
  return {
    fight: firstFight,
    mode: 'analyze',
    view: 'tables',
    tab: 'summary',
    source: SOURCE_FRIENDLIES,
    start: null,
    end: null,
  };
}

/** A non-negative integer, or null for anything else. */
function readMs(value: string | null): number | null {
  if (value === null) return null;
  if (!/^\d+$/.test(value)) return null;
  return Number.parseInt(value, 10);
}

function readOne<T extends string>(value: string | null, allowed: readonly T[], fallback: T): T {
  return value !== null && (allowed as readonly string[]).includes(value) ? (value as T) : fallback;
}

export function parseReportState(search: string, firstFight: number): ReportState {
  const params = new URLSearchParams(search);
  const state = defaultState(firstFight);

  const fight = params.get('fight');
  if (fight === ALL_FIGHTS_PARAM) state.fight = ALL_FIGHTS;
  else if (fight !== null && /^\d+$/.test(fight)) state.fight = Number.parseInt(fight, 10);

  state.mode = readOne(params.get('mode'), ENABLED_MODES, state.mode);
  state.view = readOne(params.get('view'), VIEW_IDS, state.view);
  state.tab = readOne(params.get('tab'), TAB_IDS, state.tab);

  const source = params.get('source');
  // A GUID is the engine's own format: a kind, then hyphen-separated numeric fields.
  if (source !== null && /^[A-Za-z0-9-]{1,64}$/.test(source)) state.source = source;

  const start = readMs(params.get('start'));
  const end = readMs(params.get('end'));
  if (start !== null && end !== null && end > start) {
    state.start = start;
    state.end = end;
  }
  return state;
}

/** The contract's field order, so two links to the same view are the same string. */
export function reportSearch(state: ReportState, firstFight: number): string {
  const base = defaultState(firstFight);
  const params = new URLSearchParams();
  if (state.fight !== base.fight)
    params.set('fight', state.fight === ALL_FIGHTS ? ALL_FIGHTS_PARAM : String(state.fight));
  if (state.mode !== base.mode) params.set('mode', state.mode);
  if (state.view !== base.view) params.set('view', state.view);
  if (state.tab !== base.tab) params.set('tab', state.tab);
  if (state.source !== base.source) params.set('source', state.source);
  if (state.start !== null && state.end !== null) {
    params.set('start', String(state.start));
    params.set('end', String(state.end));
  }
  const query = params.toString();
  return query === '' ? '' : `?${query}`;
}

/** State is replaced, never mutated: the island holds it in one `$state` object. */
export function withState(state: ReportState, patch: Partial<ReportState>): ReportState {
  return { ...state, ...patch };
}
