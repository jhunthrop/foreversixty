// web/src/lib/sim/candidates.ts
// The candidate list Top Gear, talent compare and the Droptimizer all build, and the
// envelope it turns into.
//
// Every function here returns a new list. The rows are the page's own state and a mutation
// in place would leave Svelte's `$state` array holding the same identity with different
// contents, which is exactly the class of bug that makes a checkbox render one tick behind.
import { bulkCopy } from './copy';
import { KEEP_CURRENT_ENCHANT, NO_ENCHANT } from './enchants';
import { slotsForItem } from '../planner/rules';
import { SLOT_ALIASES, type Item, type Slot } from '../planner/types';
import type { BulkMode, BulkSpec, Candidate, GearSet, Precision, TalentLoadout } from './bulk-types';

/** Contract 1.3's origin vocabulary. `drop:` and `set:` carry an id after the colon. */
export type Origin = 'equipped' | 'bag' | 'bank' | 'search' | `drop:${string}` | `set:${string}`;

export interface CandidateRow {
  /** The slot this row is shown under. Rings and trinkets show under the first of the pair. */
  slot: Slot;
  item: Item;
  origin: Origin;
  /**
   * Where it comes from, in words -- "Ragnaros", "Blacksmithing" -- travelling to the
   * engine as `Candidate.SourceName` (contract 10.1 A6) and coming back on the
   * substitution, so the results view never re-joins an id to a name. Empty for anything
   * that is not a drop.
   */
  sourceName: string;
  /** 0 inherits the equipped enchant; -1 is the page's "explicitly none". */
  enchant: number;
  suffix: number;
  checked: boolean;
}

export const CANDIDATE_ROW_SEPARATOR = ':';

/** Row identity: slot, item, enchant, suffix. A copy-and-modify is therefore a new row. */
export function candidateKey(row: Pick<CandidateRow, 'slot' | 'item' | 'enchant' | 'suffix'>): string {
  return [row.slot, row.item.id, row.enchant, row.suffix].join(CANDIDATE_ROW_SEPARATOR);
}

/**
 * What goes on the wire. "" for anything that fits more than one slot -- rings, trinkets
 * and weapons -- because only the item database can decide which slot it lands in, and
 * sim/bulk is the one that has it (contract 1.3, and design 3.2's "rings and trinkets are
 * tried in both slots").
 */
export function envelopeSlotOf(item: Item): string {
  const fits = slotsForItem(item);
  return fits.length === 1 ? fits[0] : '';
}

/**
 * Which slot the grid shows this item under. Rings and trinkets get one row, not two --
 * design 3.1.2, "Rings and trinkets show one list for both slots" -- so an aliased item
 * shows under the first of its pair.
 */
export function uiSlotsOf(item: Item): Slot[] {
  const aliased = SLOT_ALIASES[item.slot];
  if (aliased !== undefined) return [aliased[0]];
  return slotsForItem(item);
}

export function rowFor(item: Item, slot: Slot, origin: Origin, sourceName = ''): CandidateRow {
  return { slot, item, origin, sourceName, enchant: NO_ENCHANT, suffix: 0, checked: false };
}

export function toggleRow(rows: readonly CandidateRow[], key: string): CandidateRow[] {
  return rows.map((row) => (candidateKey(row) === key ? { ...row, checked: !row.checked } : row));
}

/** Adds a row, or ticks the one already there -- adding the same item twice is not two rows. */
export function addRow(rows: readonly CandidateRow[], row: CandidateRow): CandidateRow[] {
  const key = candidateKey(row);
  if (rows.some((existing) => candidateKey(existing) === key)) {
    return rows.map((existing) =>
      candidateKey(existing) === key ? { ...existing, checked: existing.checked || row.checked } : existing,
    );
  }
  return [...rows, row];
}

/**
 * Design 3.1.2's "copy and modify": the same item again with a different enchant or suffix,
 * added beside the original and ticked, never replacing it. Adding a copy that already
 * exists ticks it instead of duplicating it.
 */
export function copyAndModify(
  rows: readonly CandidateRow[],
  key: string,
  patch: { enchant?: number; suffix?: number },
): CandidateRow[] {
  const source = rows.find((row) => candidateKey(row) === key);
  if (source === undefined) return [...rows];
  const copy: CandidateRow = { ...source, ...patch, checked: true };
  if (candidateKey(copy) === key) return rows.map((row) => (row === source ? copy : row));
  const at = rows.indexOf(source);
  const next = [...rows];
  const existing = next.findIndex((row) => candidateKey(row) === candidateKey(copy));
  if (existing >= 0) {
    next[existing] = { ...next[existing], checked: true };
    return next;
  }
  next.splice(at + 1, 0, copy);
  return next;
}

export function removeRow(rows: readonly CandidateRow[], key: string): CandidateRow[] {
  return rows.filter((row) => candidateKey(row) !== key);
}

/**
 * The ticked rows as envelope candidates. A locked slot contributes nothing -- the contract
 * validates "no candidate on a locked slot" and refusing here is what keeps the page from
 * sending a request it knows will be refused.
 */
export function toCandidates(rows: readonly CandidateRow[], locked: readonly string[]): Candidate[] {
  return rows
    .filter((row) => row.checked && !locked.includes(row.slot))
    .map((row) => {
      const candidate: Candidate = {
        slot: envelopeSlotOf(row.item),
        item_id: row.item.id,
        origin: row.origin,
      };
      // NO_ENCHANT (0) is the wire's own "inherit the equipped enchant" and KEEP_CURRENT_ENCHANT
      // (-1) is this page's "explicitly none" -- both mean the same thing on the wire, an
      // omitted field, so neither ever travels as a literal number.
      if (row.enchant !== NO_ENCHANT && row.enchant !== KEEP_CURRENT_ENCHANT) {
        candidate.enchant = row.enchant;
      }
      if (row.suffix > 0) candidate.suffix = row.suffix;
      // Contract 10.1 A6: the page fills the source's name once, here, and reads it back
      // off the substitution rather than joining the id to loot.json a second time.
      if (row.sourceName !== '') candidate.source_name = row.sourceName;
      return candidate;
    });
}

export interface BulkSpecInput {
  mode: BulkMode;
  rows: readonly CandidateRow[];
  locked: readonly string[];
  loadouts: readonly TalentLoadout[];
  sets: readonly GearSet[];
  precision: Precision;
  cap: number;
}

/**
 * The envelope's bulk block. `talents` mode sends no candidates at all, whatever is ticked:
 * the contract validates it and the page's own gear grid is hidden in that mode anyway.
 *
 * `consumables` (Task 1's `BulkSpec.consumables`) is deliberately not read from `input`
 * here: Task 13 adds the control and the plumbing for it. Leaving it unset costs nothing --
 * `BulkSpec.consumables` is optional -- and a later task adds one field to `BulkSpecInput`
 * and one line here rather than reshaping this function.
 */
export function buildBulkSpec(input: BulkSpecInput): BulkSpec {
  return {
    mode: input.mode,
    candidates: input.mode === 'talents' ? [] : toCandidates(input.rows, input.locked),
    talents: [...input.loadouts],
    sets: [...input.sets],
    locked: [...input.locked],
    precision: input.precision,
    cap: input.cap,
  };
}

/**
 * The contract's own validation, before the request is sent, so a refusal is a sentence
 * beside the button rather than an engine error after a wait. Null means valid.
 *
 * All four contract rules, in order: no candidate on a locked slot; `talents` mode needs at
 * least one loadout and no candidates (enforced structurally by `buildBulkSpec` and checked
 * again here for a hand-built spec); `drops` mode's candidates must all be `drop:` origins;
 * `gear` mode needs at least one of candidates, talents or sets.
 */
export function validateBulk(spec: BulkSpec): string | null {
  const locked = spec.locked ?? [];
  if (spec.candidates.some((candidate) => locked.includes(candidate.slot))) {
    return bulkCopy.noCandidates;
  }
  if (spec.mode === 'talents') {
    return (spec.talents ?? []).length > 0 ? null : bulkCopy.noCandidates;
  }
  if (spec.mode === 'drops') {
    if (spec.candidates.length === 0) return bulkCopy.dropsNothing;
    return spec.candidates.every((candidate) => candidate.origin.startsWith('drop:'))
      ? null
      : bulkCopy.dropsNothing;
  }
  const anything =
    spec.candidates.length > 0 || (spec.talents ?? []).length > 0 || (spec.sets ?? []).length > 0;
  return anything ? null : bulkCopy.noCandidates;
}
