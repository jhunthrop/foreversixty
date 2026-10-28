// web/src/lib/coach/chips.ts
// The evidence chips under a finding: one per anchor id the answer carries, labelled by the
// copy module's `anchor*` strings and linked through `coachAnchorHref`. An anchor id missing
// from `anchors`, or one that cannot be linked, renders no chip.
import { coachAnchorHref } from './anchors';
import { coachCopy, fill } from './copy';
import type { CoachAnchor } from './types';

export interface Chip {
  id: string;
  href: string;
  label: string;
  spellId: number | null;
}

/** `m:ss` from the pull's start. */
export function clock(ms: number): string {
  const total = Math.max(0, Math.floor(ms / 1000));
  return `${Math.floor(total / 60)}:${String(total % 60).padStart(2, '0')}`;
}

/** The death anchor's own time: the engine's `death-<guid>-<ms>` form ends in the millisecond. */
function deathMs(death: string | undefined): number | null {
  const tail = Number(death?.split('-').pop());
  return Number.isFinite(tail) ? tail : null;
}

/** `m:ss`, or `m:ss–m:ss` when the anchor names a window that ends after it starts. */
function span(anchor: CoachAnchor): string {
  const start = anchor.at_ms ?? 0;
  return anchor.to_ms !== undefined && anchor.to_ms > start
    ? `${clock(start)}–${clock(anchor.to_ms)}`
    : clock(start);
}

export function chipLabel(anchor: CoachAnchor, spellName: (id: number) => string): string | null {
  const spell = anchor.spell_id === undefined ? '' : spellName(anchor.spell_id);
  switch (anchor.kind) {
    case 'cast':
      return fill(coachCopy.anchorCast, { Spell: spell });
    case 'aura':
      return fill(anchor.debuff === true ? coachCopy.anchorDebuffs : coachCopy.anchorBuffs, { Spell: spell });
    case 'buffs':
      return coachCopy.anchorBuffsPull;
    case 'healing':
      return fill(coachCopy.anchorHealing, { Spell: spell });
    case 'time':
      return anchor.at_ms === undefined ? null : fill(coachCopy.anchorTime, { t: span(anchor) });
    case 'resource':
      return anchor.at_ms === undefined ? null : fill(coachCopy.anchorResource, { t: span(anchor) });
    case 'death': {
      const at = deathMs(anchor.death);
      return at === null ? null : fill(coachCopy.anchorDeath, { t: clock(at) });
    }
    default:
      return null;
  }
}

export interface ChipContext {
  anchors: Readonly<Record<string, CoachAnchor>>;
  fightIndex: number;
  firstFight: number;
  playerGuid: string;
  spellName: (id: number) => string;
}

export function chipsFor(ids: readonly string[], context: ChipContext): Chip[] {
  return ids.flatMap((id) => {
    const anchor = context.anchors[id];
    if (anchor === undefined) return [];
    const href = coachAnchorHref(context.fightIndex, context.firstFight, context.playerGuid, anchor);
    const label = chipLabel(anchor, context.spellName);
    return href === null || label === null ? [] : [{ id, href, label, spellId: anchor.spell_id ?? null }];
  });
}

/** The first anchor that carries a spell, for the item's icon. */
export function firstSpellId(
  ids: readonly string[],
  anchors: Readonly<Record<string, CoachAnchor>>,
): number | null {
  for (const id of ids) {
    const spell = anchors[id]?.spell_id;
    if (spell !== undefined) return spell;
  }
  return null;
}
