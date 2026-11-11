// web/src/lib/coach/chips.ts
// The evidence chips under a finding: one per anchor id the answer carries, labelled by the
// copy module's `anchor*` strings and linked through `coachAnchorHref`. An anchor id missing
// from `anchors`, or one that cannot be linked, renders no chip.
import { coachAnchorHref } from './anchors';
import { coachCopy, fill } from './copy';
import type { CoachAnchor } from './types';
import { clockSeconds, SECOND_MS } from '../format/duration';

export interface Chip {
  id: string;
  href: string;
  label: string;
  spellId: number | null;
}

/** `m:ss` from the pull's start, to the whole second (lib/format/duration.ts). */
export function pullClock(ms: number): string {
  return clockSeconds(Math.floor(Math.max(0, ms) / SECOND_MS));
}

/** A legacy death anchor's own time: the engine's `death-<guid>-<ms>` form ends in the millisecond. */
function deathMs(death: string | undefined): number | null {
  const tail = Number(death?.split('-').pop());
  return Number.isFinite(tail) ? tail : null;
}

/** `m:ss`, or `m:ss–m:ss` when the anchor names a window that ends after it starts. */
function anchorSpan(anchor: CoachAnchor): string {
  const start = anchor.at_ms ?? 0;
  return anchor.to_ms !== undefined && anchor.to_ms > start
    ? `${pullClock(start)}–${pullClock(anchor.to_ms)}`
    : pullClock(start);
}

const SPELL_LEAD = '{Spell} · ';

/** `{Spell} · Casts` with the name, or the kind word alone (`Casts`) when the name is unknown. */
function spellChip(template: string, spell: string): string {
  return spell === '' ? template.replace(SPELL_LEAD, '') : fill(template, { Spell: spell });
}

export function chipLabel(anchor: CoachAnchor, spellName: (id: number) => string): string | null {
  const spell = anchor.spell_name ?? (anchor.spell_id === undefined ? '' : spellName(anchor.spell_id));
  switch (anchor.kind) {
    case 'cast':
      return spellChip(coachCopy.anchorCast, spell);
    case 'aura':
      return spellChip(anchor.debuff === true ? coachCopy.anchorDebuffs : coachCopy.anchorBuffs, spell);
    case 'buffs':
      return coachCopy.anchorBuffsPull;
    case 'healing':
      return spellChip(coachCopy.anchorHealing, spell);
    case 'time':
      return anchor.at_ms === undefined ? null : fill(coachCopy.anchorTime, { t: anchorSpan(anchor) });
    case 'resource':
      return anchor.at_ms === undefined ? null : fill(coachCopy.anchorResource, { t: anchorSpan(anchor) });
    case 'death': {
      const at = anchor.at_ms ?? deathMs(anchor.death);
      return at === null ? null : fill(coachCopy.anchorDeath, { t: pullClock(at) });
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
