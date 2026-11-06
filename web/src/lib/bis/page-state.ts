// web/src/lib/bis/page-state.ts
// The BiS spec page's state, with the URL as the truth (web layout spec 2026-11-05 §3.4, R1):
// `?band=60&faction=horde&race=orc&view=where&sources=dungeon,quest&preset=bare`. A parameter at
// its default is left out of the URL. The legacy `#band-<faction>-<band>` hash (every link the
// site and the header selector wrote before this page) still opens that band and faction.
// Defaults follow the band (§5.10 item 10, §5.11): below 60 the page opens on Where to go with
// Reputation off (standings up to Friendly kept); at 60 on By slot with every source on.
import { bandForLevel } from './hover';
import { SOURCE_FILTER_KINDS } from './source-filter';
import type { BisPresetId, Faction, SourceFilterKind } from './types';

export type BisView = 'slot' | 'where';

export interface BisPageState {
  faction: Faction;
  band: number;
  preset: BisPresetId;
  race: string;
  view: BisView;
  /** The source kinds that are on. */
  sources: readonly SourceFilterKind[];
}

/** What the page offers: the bands, the presets per band, the races per faction. */
export interface BisStateContext {
  bands: readonly number[];
  presetsFor: (faction: Faction, band: number) => readonly BisPresetId[];
  racesFor: (faction: Faction) => readonly string[];
  /** The race a band was simmed as, the race control's default. */
  simRaceFor: (faction: Faction, band: number, preset: BisPresetId) => string;
}

export const MAX_BAND = 60;

/** The query params the page's state reads (stateSearch writes the same names). */
export const STATE_PARAMS = ['band', 'faction', 'preset', 'race', 'view', 'sources'] as const;
const LEGACY_HASH = /^#band-(alliance|horde)-(\d+)$/;

export function defaultView(band: number): BisView {
  return band < MAX_BAND ? 'where' : 'slot';
}

/** Below 60 Reputation is off by default (most levelers stand below Honored); at 60 all are on. */
export function defaultSources(band: number): SourceFilterKind[] {
  return band < MAX_BAND ? SOURCE_FILTER_KINDS.filter((kind) => kind !== 'rep') : [...SOURCE_FILTER_KINDS];
}

/** The default preset: raid-ready where the band has it. */
export function defaultPreset(presets: readonly BisPresetId[]): BisPresetId {
  return presets.includes('raid') ? 'raid' : (presets[0] ?? 'bare');
}

function sameKinds(a: readonly string[], b: readonly string[]): boolean {
  return a.length === b.length && a.every((kind) => b.includes(kind));
}

function parseSources(value: string | null, band: number): SourceFilterKind[] {
  if (value === null) return defaultSources(band);
  if (value === 'all') return [...SOURCE_FILTER_KINDS];
  const picked = value.split(',').filter((kind): kind is SourceFilterKind =>
    (SOURCE_FILTER_KINDS as readonly string[]).includes(kind),
  );
  return SOURCE_FILTER_KINDS.filter((kind) => picked.includes(kind));
}

function nearestBand(bands: readonly number[], level: number): number {
  const band = bandForLevel(level);
  return bands.includes(band) ? band : (bands[0] ?? band);
}

/** The state a URL names, every missing or unknown value at its default. */
export function parseState(search: string, hash: string, context: BisStateContext): BisPageState {
  const params = new URLSearchParams(search);
  const legacy = LEGACY_HASH.exec(hash);
  const factionParam = legacy?.[1] ?? params.get('faction');
  const faction: Faction = factionParam === 'horde' ? 'horde' : 'alliance';
  const bandParam = legacy?.[2] ?? params.get('band');
  const band =
    bandParam !== null && /^\d{1,2}$/.test(bandParam)
      ? nearestBand(context.bands, Number(bandParam))
      : (context.bands[0] ?? MAX_BAND);
  const presets = context.presetsFor(faction, band);
  const presetParam = params.get('preset');
  const preset =
    presetParam !== null && (presets as readonly string[]).includes(presetParam)
      ? (presetParam as BisPresetId)
      : defaultPreset(presets);
  const races = context.racesFor(faction);
  const raceParam = params.get('race');
  const race = raceParam !== null && races.includes(raceParam) ? raceParam : context.simRaceFor(faction, band, preset);
  const viewParam = params.get('view');
  const view: BisView = viewParam === 'slot' || viewParam === 'where' ? viewParam : defaultView(band);
  return { faction, band, preset, race, view, sources: parseSources(params.get('sources'), band) };
}

/** The query string for a state: only what differs from its defaults, in a fixed order. */
export function stateSearch(state: BisPageState, context: BisStateContext): string {
  const params = new URLSearchParams();
  if (state.band !== context.bands[0]) params.set('band', String(state.band));
  if (state.faction !== 'alliance') params.set('faction', state.faction);
  if (state.preset !== defaultPreset(context.presetsFor(state.faction, state.band))) params.set('preset', state.preset);
  if (state.race !== context.simRaceFor(state.faction, state.band, state.preset)) params.set('race', state.race);
  if (state.view !== defaultView(state.band)) params.set('view', state.view);
  if (!sameKinds(state.sources, defaultSources(state.band))) {
    params.set('sources', sameKinds(state.sources, SOURCE_FILTER_KINDS) ? 'all' : state.sources.join(','));
  }
  const text = params.toString().replace(/%2C/g, ',');
  return text === '' ? '' : `?${text}`;
}

/** A change of band or faction keeps what still applies and resets what belongs to the band:
 *  the preset (a band without raid falls back), the race (a faction's own), the defaults. */
export function withChange(
  state: BisPageState,
  change: Partial<BisPageState>,
  context: BisStateContext,
): BisPageState {
  const next = { ...state, ...change };
  const bandMoved = change.band !== undefined && change.band !== state.band;
  const factionMoved = change.faction !== undefined && change.faction !== state.faction;
  const presets = context.presetsFor(next.faction, next.band);
  const preset = presets.includes(next.preset) && !bandMoved ? next.preset : defaultPreset(presets);
  const races = context.racesFor(next.faction);
  const race =
    factionMoved || !races.includes(next.race) ? context.simRaceFor(next.faction, next.band, preset) : next.race;
  const crossesSixty = bandMoved && (state.band < MAX_BAND) !== (next.band < MAX_BAND);
  return {
    ...next,
    preset,
    race,
    view: crossesSixty && change.view === undefined ? defaultView(next.band) : next.view,
    sources: crossesSixty && change.sources === undefined ? defaultSources(next.band) : next.sources,
  };
}

/** Whether the URL names anything the server-rendered default did not show. */
export function isDefaultState(state: BisPageState, context: BisStateContext): boolean {
  return stateSearch(state, context) === '';
}
