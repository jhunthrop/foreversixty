// web/tests/e2e/support/tier-list.ts
// What the tier list should show, read from the published BiS files the way the spec states
// it (design/specs/2026-10-09-tier-list.md, section 9) and not from the page's own code.
import { bisBand, readSpecCatalog, type BisBandFile, type SpecCatalogRow } from './bis-file';

export type TierRole = 'dps' | 'tank' | 'healer';
export type TierFaction = 'alliance' | 'horde';

export interface ExpectedTierRow {
  entry: SpecCatalogRow;
  band: BisBandFile;
  /** The sorted figure: damage taken per second for a tank, set DPS otherwise. */
  metric: number;
}

const TIE_PERCENT = 1;

export function expectedRows(role: TierRole, faction: TierFaction): ExpectedTierRow[] {
  const rows = readSpecCatalog()
    .filter((entry) => entry.role === role)
    .map((entry) => {
      const band = bisBand(entry.spec, faction, 60, 'raid');
      const metric = role === 'tank' ? band.metrics!.dtps : band.set_dps!;
      return { entry, band, metric };
    });
  const direction = role === 'tank' ? 1 : -1;
  return rows.sort((a, b) => direction * (a.metric - b.metric) || a.entry.name.localeCompare(b.entry.name));
}

/** How many rows carry a tie mark: those within 1% of a neighbour (DPS and healer only). */
export function expectedTieCount(rows: readonly ExpectedTierRow[], role: TierRole): number {
  if (role === 'tank') return 0;
  const within = (a: number, b: number): boolean => Math.abs(1 - a / b) * 100 <= TIE_PERCENT;
  return rows.filter(
    (row, i) =>
      (i > 0 && within(row.metric, rows[i - 1]!.metric)) ||
      (i < rows.length - 1 && within(rows[i + 1]!.metric, row.metric)),
  ).length;
}

const ONE_DECIMAL = new Intl.NumberFormat('en-US', { minimumFractionDigits: 1, maximumFractionDigits: 1 });
const WHOLE = new Intl.NumberFormat('en-US', { maximumFractionDigits: 0 });
export const oneDecimal = (value: number): string => ONE_DECIMAL.format(value);
/** The sorted figure as the page writes it: whole for a tank, one decimal otherwise. */
export const metricText = (role: TierRole, value: number): string =>
  role === 'tank' ? WHOLE.format(value) : ONE_DECIMAL.format(value);
export const wholeNumber = (value: number): string => WHOLE.format(value);

export function raceName(race: string): string {
  return race
    .split('-')
    .map((part) => part.charAt(0).toUpperCase() + part.slice(1))
    .join(' ');
}
