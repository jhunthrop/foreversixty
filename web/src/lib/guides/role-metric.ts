// web/src/lib/guides/role-metric.ts
// The role a band was ranked for (2026-11-02 launch defect: every role read "DPS"). A damage
// band's figures are damage; a healer band's `*_dps*` keys carry effective healing (types.ts
// `BisBand.role`); a tank band's per-point figures are tank-score points
// (`score_unit: "tank_score"`, `tankCopy`). Never relabels one role's number as another's.
import type { BisBand, BisRole } from '../bis/types';

/** A band's role, `'dps'` for a file published before roles (`normaliseBisFile`'s default). */
export function roleOf(band: Pick<BisBand, 'role'>): BisRole {
  return band.role ?? 'dps';
}
