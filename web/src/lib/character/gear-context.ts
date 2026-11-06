// web/src/lib/character/gear-context.ts
// What the character page's worn paperdoll is read against (design/specs/2026-11-06-character.md
// C2, C4): the class's items file (names' colours, icons, two-handedness, the gain rule's stat
// score) and the spec's BiS band at the character's level and faction. Each half settles on its
// own: a failed items file is null (plain names), a missing band is null (no gain is priced).
import { bandEntryFor, bandForLevel, fetchBisFile } from '../bis/hover';
import type { BisBand, BisSlot } from '../bis/types';
import type { Faction } from '../faction-mark';
import { itemFacts, type GearItemTable } from '../gear/qualities';

export interface GearContext {
  table: GearItemTable | null;
  band: BisBand | null;
}

export interface GearContextInput {
  build: string;
  classSlug: string | null;
  /** `gear_spec` ("priest-holy"). */
  gearSpec: string | null;
  level: number | null;
  faction: Faction | null;
}

/** Both reads at once; never rejects. */
export async function loadGearContext(input: GearContextInput): Promise<GearContext> {
  const { build, classSlug, gearSpec, level, faction } = input;
  const table = classSlug === null ? Promise.resolve(null) : itemFacts(build, classSlug);
  const band =
    gearSpec === null || level === null || faction === null
      ? Promise.resolve(null)
      : fetchBisFile(build, gearSpec)
          .then((file) => (file === null ? null : (bandEntryFor(file, bandForLevel(level), faction) ?? null)))
          .catch(() => null);
  const [loadedTable, loadedBand] = await Promise.all([table, band]);
  return { table: loadedTable, band: loadedBand };
}

/** C4: the band's pick row for a slot (a pair's pick may sit under either ring or trinket). */
export function bandSlotFor(band: BisBand | null, slot: string, pickId: number | null): BisSlot | null {
  if (band === null || pickId === null) return null;
  return band.slots.find((row) => row.slot === slot && row.item_id === pickId) ?? null;
}

/** C4: the inspector's two runners-up: the pick's listed alternatives, the closest first. */
export function runnersUp(
  pick: BisSlot | null,
  count = 2,
): { id: number; name: string; source: string; delta: number }[] {
  return (pick?.alternatives ?? [])
    .filter((alt) => alt.item_name !== undefined)
    .slice(0, count)
    .map((alt) => ({
      id: alt.item_id,
      name: alt.item_name ?? '',
      source: alt.source ?? '',
      delta: alt.dps_delta ?? 0,
    }));
}
