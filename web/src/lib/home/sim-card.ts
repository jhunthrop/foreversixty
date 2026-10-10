// web/src/lib/home/sim-card.ts
// What the home hero's Simulator card may say, as one pure decision (tenet 8). The card pairs the
// player's own latest saved sim ("26 DPS now") with the band's best-in-slot set DPS ("37.5 DPS at
// band best in slot"). The two figures are only comparable when the saved sim ran the same spec,
// band, faction and preset as the band entry the second figure comes from; otherwise the first
// figure is withheld and the band figure stands alone.
import { presetOf } from '../bis/presets';
import { headlineDpsOf } from '../bis/tank-view';
import type { BisBand } from '../bis/types';
import type { SimListRow } from '../sim/types';

export interface SimCardTarget {
  /** The hero's canonical spec key ("hunter-marksmanship"), undefined with no spec learned. */
  specKey: string | undefined;
  /** The hero's own band entry, undefined when none is published. */
  band: BisBand | undefined;
}

export interface SimCardFigures {
  /** The saved sim's DPS, only when it is this hero's spec and (with a band) the band's own setup. */
  nowDps: number | null;
  /** The band's best-in-slot DPS, only for a DPS band (a tank band has no headline DPS). */
  bandDps: number | null;
}

/** True when `row` was run on exactly the setup `band` describes. A row that does not say which
 *  band, faction and preset it ran with cannot be shown to match, so it does not. */
export function simMatchesBand(row: SimListRow, specKey: string | undefined, band: BisBand): boolean {
  return (
    specKey !== undefined &&
    row.spec === specKey &&
    row.band === band.band &&
    row.faction === band.faction &&
    row.preset === presetOf(band)
  );
}

export function simCardFigures(row: SimListRow, target: SimCardTarget): SimCardFigures {
  const { specKey, band } = target;
  if (band === undefined) {
    return { nowDps: specKey !== undefined && row.spec === specKey ? row.dps : null, bandDps: null };
  }
  return {
    nowDps: simMatchesBand(row, specKey, band) ? row.dps : null,
    bandDps: headlineDpsOf(band) ?? null,
  };
}
