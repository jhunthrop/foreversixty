// web/src/lib/addon/import.ts
// An addon export, turned into the draft the planner edits.
//
// The export carries a finished tree and no order, because nothing in the game records
// the order a build was spent in. `orderFromRanks` reconstructs a legal one with the
// planner's own gates -- lowest tier first, left to right -- so an imported build is one
// the planner and the API both accept by construction, and the box says the order is
// approximated rather than implying the addon knew it.
import { addonCopy } from './copy';
import { decodeFS1, orderFromRanks } from '../planner/fs1';
import type { TalentIndex } from '../planner/rules';
import type { Gear } from '../planner/types';

export type ImportOutcome =
  | { ok: true; classSlug: string; raceSlug: string; order: number[]; gear: Gear; notes: string[] }
  | { ok: false; message: string };

export function importFromAddon(code: string, talents: TalentIndex, activeBuild: string): ImportOutcome {
  const decoded = decodeFS1(code);
  if (!decoded.ok) return { ok: false, message: decoded.message };

  // Refused, not reconstructed: `orderFromRanks` below has no idea which class its ranks
  // came from, and reconstructing anyway would spend `talents`' class's talent ids to match
  // rank counts read positionally off a tree that is not the export's own -- a plausible-
  // looking build that is not the one the export names. `talents.file.class_slug` is the
  // class the planner is actually showing (`talents` is always `store.talentIndex`, which
  // only ever holds the loaded class's own file), so no extra parameter is needed to know
  // it. Checked first, ahead of the older-build note below: there is no order to annotate
  // a build note onto when nothing is being imported.
  if (decoded.build.classSlug !== talents.file.class_slug) {
    return {
      ok: false,
      message: addonCopy.importWrongClass(decoded.build.classSlug, talents.file.class_slug),
    };
  }

  const { order, dropped } = orderFromRanks(talents, decoded.build.treeRanks);
  // Annotated explicitly: `addonCopy`'s `as const` narrows `importOrderApproximated` to its
  // own literal type, and without this the array would infer that single literal type
  // instead of `string[]`, refusing the general `string` the two pushes below return.
  const notes: string[] = [addonCopy.importOrderApproximated];
  if (dropped.length > 0) {
    const names = dropped.map((id) => talents.byId.get(id)?.name ?? String(id)).join(', ');
    notes.push(addonCopy.importDropped(names));
  }
  if (decoded.build.dataBuild !== activeBuild) {
    // A note, never a refusal. An addon a release behind still describes the
    // player's real character, and refusing it would make the feature useless
    // for the week after every data build.
    notes.push(addonCopy.importOlderBuild(decoded.build.dataBuild, activeBuild));
  }

  return {
    ok: true,
    classSlug: decoded.build.classSlug,
    raceSlug: decoded.build.raceSlug,
    order,
    // `gear` rather than `gearSlots`: the planner's map has no room for an
    // enchant or a suffix, and the decoder already reduced them for us.
    gear: decoded.build.gear,
    notes,
  };
}
