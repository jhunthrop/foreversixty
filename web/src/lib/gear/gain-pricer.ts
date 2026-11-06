// web/src/lib/gear/gain-pricer.ts
// The worn paperdoll's gains in the band's own unit (web layout spec 2026-11-05 §4.6, owner
// decision D4): the API says which slots gain (`CharacterGearSlot.gap`, Gear component points);
// this prices each one with the one gain rule the home panels use (lib/home/upgrades.ts
// `slotGainFor`: a listed runner-up's sim delta, else the score difference times the band's
// `reference_dps_per_point`, never a weapon by stat weights). No band, no price: the doll says
// "Not sim-checked", never 0.
import type { BisBand } from '../bis/types';
import { slotScoreUnitFor } from '../bis/tank-view';
import { slotGainFor, weightsRecordFor } from '../home/upgrades';
import type { Item } from '../planner/types';
import type { GainPricer, PaperdollSlotKey } from './paperdoll';

/** Rings and trinkets are one pair each: a pick listed under either slot prices either. */
function pairOf(slot: string): string {
  return slot.replace(/[12]$/, '');
}

/** Prices nothing: the band did not load or the character has no spec. */
export const UNPRICED: GainPricer = { unit: 'dps', price: () => null };

/** The band's pricer; `item` reads the class's items file (a worn id it lacks is unknown). */
export function bandPricer(band: BisBand, item: (id: number) => Item | undefined): GainPricer {
  const weights = weightsRecordFor(band);
  return {
    unit: slotScoreUnitFor(band),
    price: (slot: PaperdollSlotKey, pickId: number, wornId: number | undefined): number | null => {
      const pick =
        band.slots.find((row) => row.slot === slot && row.item_id === pickId) ??
        band.slots.find((row) => pairOf(row.slot) === pairOf(slot) && row.item_id === pickId);
      if (pick === undefined) return null;
      const worn = wornId === undefined ? undefined : item(wornId);
      return slotGainFor(slot, pick, wornId, worn, band, weights).gainDps;
    },
  };
}
