// web/tests/e2e/support/tier-list.ts
// What the tier list should show, read from the published BiS files the way the spec states
// it (design/specs/2026-10-09-tier-list.md, section 9) and not from the page's own code.
import { bisBand, readSpecCatalog, type BisBandFile, type SpecCatalogRow } from './bis-file';

export type TierRole = 'dps' | 'tank' | 'healer';
export type TierFaction = 'alliance' | 'horde';

export interface ExpectedTierRow {
  entry: SpecCatalogRow;
  band: BisBandFile;
  /** The sorted figure: damage taken per second for a tank, the raid fight's HPS for a healer
   *  once every healer band publishes it (set DPS until then), set DPS for damage. */
  metric: number;
}

/** A healer band's raid-fight figure (metrics.raid_fight), absent until the ranker publishes it. */
const raidFightOf = (band: BisBandFile) =>
  (band.metrics as { raid_fight?: { hps: number } } | null | undefined)?.raid_fight;

function metricOf(role: TierRole, band: BisBandFile, healerOnRaid: boolean): number {
  if (role === 'tank') return band.metrics!.dtps;
  if (role === 'healer' && healerOnRaid) return raidFightOf(band)!.hps;
  return band.set_dps!;
}

export function expectedRows(role: TierRole, faction: TierFaction): ExpectedTierRow[] {
  const bands = readSpecCatalog()
    .filter((entry) => entry.role === role)
    .map((entry) => ({ entry, band: bisBand(entry.spec, faction, 60, 'raid') }));
  // A healer list ranks on the raid fight only when every healer band carries it (bbc07527e).
  const healerOnRaid = bands.every(({ band }) => raidFightOf(band) !== undefined);
  const rows = bands.map(({ entry, band }) => ({ entry, band, metric: metricOf(role, band, healerOnRaid) }));
  const direction = role === 'tank' ? 1 : -1;
  return rows.sort((a, b) => direction * (a.metric - b.metric) || a.entry.name.localeCompare(b.entry.name));
}

const ONE_DECIMAL = new Intl.NumberFormat('en-US', { minimumFractionDigits: 1, maximumFractionDigits: 1 });
const WHOLE = new Intl.NumberFormat('en-US', { maximumFractionDigits: 0 });
export const oneDecimal = (value: number): string => ONE_DECIMAL.format(value);
/** The sorted figure as the page writes it: whole for a tank, one decimal otherwise. */
export const metricText = (role: TierRole, value: number): string =>
  role === 'tank' ? WHOLE.format(value) : ONE_DECIMAL.format(value);
export const wholeNumber = (value: number): string => WHOLE.format(value);
