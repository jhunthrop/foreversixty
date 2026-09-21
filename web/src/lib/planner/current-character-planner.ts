// web/src/lib/planner/current-character-planner.ts
// The planner side of the current-character bridge. `current-character-bridge.ts` is the
// sim side (a SimCharacter becomes a pointer); kept separate because this module depends on
// `PlannerStore`/`BuildRecord` shapes the sim-side module has no reason to import, and
// because the planner's own load sources (a `?code=`, a paste, a saved `record`) are not
// `SimCharacter`s at all -- the planner never builds one.
//
// current-character spec (2026-09-21) section 1: "A page that loads a character from its
// URL or a paste writes the pointer." and "/planner opened bare restore[s] the stored
// character... The planner restores only when its URL is bare and its own build is empty."
// This lane's ruling (task-10-brief.md, superseded only on which sources are chip-visible,
// not on which restore): a 'build' pointer already has its own permalink (`/b/<id>`, not
// `/planner`'s own bare-load path), and 'fight'/'armory' carry no talents or gear the
// planner could show at all, so only a 'code' or 'addon' pointer -- both an FS1 string --
// can ever restore here, fed through the exact path `?code=` already takes.
import { addonCodeFor } from '../addon/build-code';
import { decodeFS1 } from './fs1';
import type { TalentIndex } from './rules';
import type { PlannerStore } from './store.svelte';
import { readCurrent, writeCurrent, type CurrentCharacter } from '../current-character';
import { specOf } from '../sim/character';
import { specLabel } from '../sim/spec-label';

/** True when `/planner`'s own URL carries none of `?code=`, `?class=`, `?race=` -- the only
 *  shape a stored pointer is ever allowed to restore into (spec section 1, "opened bare
 *  restore[s] the stored character"). A `?build=` is not part of this list: `/planner` has
 *  no such query of its own, and a 'build' pointer never restores here regardless (see the
 *  header comment and `decidePlannerLoad` below). */
export function isBarePlannerUrl(search: string): boolean {
  const params = new URLSearchParams(search);
  return params.get('code') === null && params.get('class') === null && params.get('race') === null;
}

export interface PlannerLoadDecision {
  /** The FS1 code to treat exactly as `?code=` would: `decodeFS1` reads it the same way,
   *  whichever source it came from. `null` means nothing to decode -- a fresh, empty build. */
  codeParam: string | null;
  /** True when `codeParam` came from the stored pointer rather than the URL -- the caller's
   *  chip shows "Restored your last character" only then. */
  restored: boolean;
  /** True when a restore was attempted (a bare URL, no `record`, a mount that is allowed to
   *  restore, and a 'code'/'addon' pointer) but that pointer's own code no longer decodes.
   *  The caller forgets it (`clearCurrent()`) and shows no error the visitor did not cause
   *  -- the site's own stored pointer went stale, not a link or a paste just now. Matches
   *  `character-bootstrap.ts`'s `settleRestore` for the /sim* pages. */
  deadPointer: boolean;
  /** The chip's own initial pointer: `stored` on a standalone mount (whatever its source --
   *  the chip still shows a 'build'/'fight'/'armory' pointer, even though only 'code' and
   *  'addon' ever restore into the planner itself), `null` on the inline mount and once
   *  `deadPointer` says the stored one is gone. */
  pointer: CurrentCharacter | null;
}

/**
 * What `/planner`'s own mount should treat as its `?code=`, whether that came from a
 * restored pointer, and the chip's own initial pointer -- one decision, since all three
 * read the same inputs. `urlCode` wins over everything (the existing `?code=` path,
 * unchanged): a bare-load restore is only considered once the URL has offered none.
 *
 * `hasRecord` and `standalone` are two of the spec sentence's three restore gates ("no
 * record, and the planner's own build is empty"); the third is not a runtime read here --
 * before `PlannerStore` exists, the only two ways the build could be non-empty are a
 * `record` or a decoded `?code=`, and `urlCode === null` with `!hasRecord` already rules
 * both out. `standalone` is /planner's/`/b/:id`'s own top-level render; Top Gear's inline
 * "add a build" (`TalentCandidates.svelte`) passes `false` and neither restores nor shows
 * the chip's pointer.
 */
export function decidePlannerLoad(
  urlCode: string | null,
  isBare: boolean,
  hasRecord: boolean,
  standalone: boolean,
  stored: CurrentCharacter | null,
): PlannerLoadDecision {
  const initialPointer = standalone ? stored : null;
  if (urlCode !== null)
    return { codeParam: urlCode, restored: false, deadPointer: false, pointer: initialPointer };
  const eligible = standalone && !hasRecord && isBare && stored !== null;
  const restorable = eligible && (stored.source === 'code' || stored.source === 'addon');
  if (!restorable || stored === null) {
    return { codeParam: null, restored: false, deadPointer: false, pointer: initialPointer };
  }
  const decoded = decodeFS1(stored.ref);
  return decoded.ok
    ? { codeParam: stored.ref, restored: true, deadPointer: false, pointer: initialPointer }
    : { codeParam: null, restored: false, deadPointer: true, pointer: null };
}

/**
 * The label a planner load writes to the pointer: the build's own title when it has one
 * (a saved build, `/b/:id`), else the class's display name, joined with the spec label once
 * the spec is derivable from the loaded talent data (`specOf` then `specLabel`, the same
 * two helpers `current-character-bridge.ts` reads a `SimCharacter`'s spec with) --
 * otherwise the class name alone. `className` is the caller's own `store.classRow.name`:
 * this module has no reference data of its own to look it up from a slug.
 */
export function labelForPlannerLoad(
  buildTitle: string | undefined,
  className: string,
  talentIndex: TalentIndex | null,
  order: readonly number[],
): string {
  if (buildTitle !== undefined && buildTitle.trim() !== '') return buildTitle;
  if (talentIndex === null) return className;
  return `${className} · ${specLabel(specOf(talentIndex, [...order]))}`;
}

/**
 * Writes the current-character pointer for a planner load. The one write path every
 * planner load source shares (a decoded `?code=`/restored code, a successful paste, or this
 * mount's own saved `record`) so the pointer's shape (`current-character.ts`'s own
 * `CurrentCharacter`) is assembled in one place rather than at each of Planner.svelte's
 * three call sites.
 */
export function recordPlannerCharacter(
  source: 'code' | 'addon' | 'build',
  ref: string,
  label: string,
  classSlug: string,
  storage?: Storage,
): void {
  const pointer: CurrentCharacter = { source, ref, label, classSlug, savedAt: new Date().toISOString() };
  writeCurrent(pointer, storage);
}

/**
 * Every planner load source (a decoded/restored code, a paste, this mount's own `record`)
 * calls this, once talent data has loaded and only on a `standalone` mount, to build the
 * label, write the pointer, and hand back the fresh value for the chip -- Planner.svelte's
 * one `writePointer` call site. A no-op (before talent data loads, or inline) hands back
 * whatever is already stored, unchanged, rather than a value the caller has to guess at.
 */
export function writePlannerPointer(
  store: PlannerStore,
  standalone: boolean,
  source: 'code' | 'addon' | 'build',
  ref: string,
  classSlug: string,
  title?: string,
  storage?: Storage,
): CurrentCharacter | null {
  if (standalone && store.talentIndex !== null) {
    const className = store.classRow?.name ?? classSlug;
    const label = labelForPlannerLoad(title, className, store.talentIndex, store.order);
    recordPlannerCharacter(source, ref, label, classSlug, storage);
  }
  return readCurrent(storage);
}

/**
 * The planner's current build as an addon export string, or `''` when there is nothing to
 * export yet (no talent data loaded). `SharePanel.svelte`'s "Copy addon code" button and
 * the current-character chip's own copy-addon link (`CurrentCharacterChip.svelte`'s
 * `addonCode` prop) both read this, so the two can never derive it two different ways.
 */
export function plannerAddonCode(store: PlannerStore): string {
  return store.talentIndex === null
    ? ''
    : addonCodeFor({
        dataBuild: store.treeVersion,
        classSlug: store.classSlug,
        order: store.order,
        gear: store.gear,
        talents: store.talentIndex,
        items: store.itemIndex,
      });
}
