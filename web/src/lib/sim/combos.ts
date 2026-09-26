// web/src/lib/sim/combos.ts
// Reading a ranked bulk result. Nothing here is a statistic: `group` already says which
// runs are within error of the leader, `delta` already carries its own error, and the
// ordering is the planner's. This turns those into rows, labels and a winning gear list.
import type { BulkResult, Combo, Substitution } from './bulk-types';
import { plannerHrefForSpec } from './character';
import { bulkCopy } from './copy';
import { confidenceBand, formatMargin } from './estimate';
import type { CharacterSpec, Estimate, GearSlot } from './types';
import type { Item, ItemSet } from '../planner/types';

export interface ComboRow {
  /** 1-based, and shared by every member of a within-error group (design 3.3). */
  rank: number;
  combo: Combo;
  /** True for the leader's group, which is the one the page draws a rule under. */
  withinError: boolean;
  percent: number;
}

export function percentOf(delta: number, equippedMean: number): number {
  return equippedMean === 0 ? 0 : (delta / equippedMean) * 100;
}

/**
 * The design system's rule for a negative figure in a table: a minus sign, not a hyphen.
 * Exported so every caller that has to sign its own gain (`deltaLabel` below,
 * `signedGainLabel`'s BY SLOT cell) uses the identical character rather than each typing
 * `'-'` and drifting from the other two (final whole-branch review, Important 1).
 */
export const MINUS = '−';

/**
 * A gain magnitude, formatted for a GAIN column or a headline: one decimal place below 10,
 * whole (thousands-separated) at 10 and above. Rounds to one decimal FIRST, then decides
 * whole-vs-decimal off the ROUNDED value -- 9.95 rounds to 10.0, which is already >= 10, so
 * it renders "10" rather than truncating to "9.9" or keeping a false "10.0". A true zero
 * renders "0", not "0.0": a delta that really is zero should not imply precision.
 *
 * Whole-DPS rounding was hiding real differences (dps D38): two 51-point builds 0.44 DPS
 * apart both read "+0" under the old always-round-to-a-whole-number rule, with nothing on
 * the page to say which was better.
 *
 * Takes `Math.abs` of its own input: this used to only DOCUMENT a non-negative precondition
 * and trust every caller to hold it, and one call site did not -- the BY SLOT column
 * (ComboResults.svelte) passed `SlotSummaryRow.gain` (`combo.delta.mean`, unconstrained in
 * sign) straight through, so `gainLabel(-1234.2)` rendered "-1234.2": a spurious decimal
 * with the thousands separator lost, where the pre-branch code rendered "-1,234" (final
 * whole-branch review, Important 1). Taking the magnitude here, not just at that one call
 * site, means a future caller cannot reintroduce the same bug by forgetting `Math.abs`.
 */
export function gainLabel(magnitude: number): string {
  const abs = Math.abs(magnitude);
  const rounded = Math.round(abs * 10) / 10;
  if (rounded === 0 || rounded >= 10) return Math.round(rounded).toLocaleString('en-US');
  return rounded.toFixed(1);
}

/**
 * "+41 ± 11": the gain and its 95% band, each through `gainLabel`. A minus sign, not a
 * hyphen -- the design system's rule for a negative figure in a table.
 */
export function deltaLabel(delta: Estimate): string {
  const sign = delta.mean < 0 ? MINUS : '+';
  // The gain goes through `gainLabel` (sub-10 keeps a decimal, dps D38), the margin through
  // `formatMargin` -- the one function every other "±" figure on the page uses, and the only
  // one that renders "< 0.1" for a non-zero band that would otherwise round away to "0"
  // (tank-sim review D3). `gainLabel` alone would silently reintroduce that bug here.
  return `${sign}${gainLabel(delta.mean)} ± ${formatMargin(confidenceBand(delta))}`;
}

/**
 * The BY SLOT column's own GAIN cell (`SlotSummaryRow.gain`): an em dash for "we don't
 * know" (no single-slot combination measured it alone), otherwise a signed, formatted
 * magnitude -- MINUS for a slot the winning set would be better off without (a persona
 * reviewer saw a -2 drop) and never mistake for the hyphen `deltaLabel` also avoids. Pulled
 * out of ComboResults.svelte's markup rather than left as an inline expression, per this
 * lane's own rule: the testable decision belongs in a pure function (final whole-branch
 * review, Important 1).
 */
export function signedGainLabel(gain: number | null): string {
  if (gain === null) return '—';
  return `${gain < 0 ? MINUS : '+'}${gainLabel(gain)}`;
}

const SUBSTITUTION_KEY_SEPARATOR = ':';

/**
 * A single substitution's identity, for `comboIdentity` below. An `item` substitution is
 * `item_id` + `enchant` + `suffix` and deliberately NOT `slot` -- design 3.2 tries a ring or
 * trinket in both slots, so the same item at the same enchant/suffix is one candidate
 * whichever slot the engine happened to put it in (newcomer MAJOR, review.md:251-257: ranks
 * 1, 3, 5 and 9 each appeared twice with the same items and the same numbers, differing only
 * by finger1 vs finger2). A `talents` substitution is its `talents` string; a `set` or
 * `consumes` substitution is its `name`.
 */
function substitutionIdentity(sub: Substitution): string {
  if (sub.kind === 'item') {
    return ['item', sub.item_id ?? '', sub.enchant ?? 0, sub.suffix ?? 0].join(SUBSTITUTION_KEY_SEPARATOR);
  }
  if (sub.kind === 'talents') return ['talents', sub.talents ?? ''].join(SUBSTITUTION_KEY_SEPARATOR);
  return [sub.kind, sub.name ?? ''].join(SUBSTITUTION_KEY_SEPARATOR);
}

/**
 * A combination's identity: its substitutions' own identities, sorted so the engine's own
 * emission order (rule 4's dual-wield pair tried both ways round, or any other order it
 * happens to emit two substitutions in) can never make the same combination look like two.
 */
function comboIdentity(combo: Combo): string {
  return combo.substitutions.map(substitutionIdentity).sort().join('|');
}

/**
 * The result's own combos, minus any that duplicate an earlier one's substitution SET
 * (design 3.2's rings-and-trinkets-in-both-slots rule can otherwise emit the same
 * combination twice, at two different slots). The result is already ranked, so the first
 * occurrence of an identity is the best-ranked and every later one is dropped.
 *
 * Keeping the first occurrence is only correct while duplicate placements (the same item at
 * the same enchant/suffix, tried in finger1 vs finger2) carry identical deltas -- the engine
 * treats both fingers as interchangeable today, so it does. If it ever stops (a set bonus or
 * a slot-specific proc that made one finger genuinely better than the other), the surviving
 * row would silently claim whichever number happened to come first, with nothing on the page
 * naming which slot it actually measured (final whole-branch review, Minor 5).
 */
function dedupedCombos(combos: readonly Combo[]): Combo[] {
  const seen = new Set<string>();
  const deduped: Combo[] = [];
  for (const combo of combos) {
    const identity = comboIdentity(combo);
    if (seen.has(identity)) continue;
    seen.add(identity);
    deduped.push(combo);
  }
  return deduped;
}

/**
 * Rank shared across a group, per design 3.3: "Rows in the leader's within-error group
 * carry the same rank." A group's rank is the 1-based position of its first member, so the
 * run after a two-member tie is rank 3.
 *
 * Ranks are computed AFTER the de-dupe, over the de-duplicated list, so dropping a
 * duplicate never opens a hole (a group whose duplicate member was dropped keeps its
 * leader's rank, and every later group's rank shifts down by however many rows were
 * dropped before it) -- ranking the raw list first and filtering the rows afterward would
 * leave exactly such a hole.
 */
export function comboRows(result: BulkResult): ComboRow[] {
  const firstOfGroup = new Map<number, number>();
  return dedupedCombos(result.combos).map((combo, index) => {
    if (!firstOfGroup.has(combo.group)) firstOfGroup.set(combo.group, index + 1);
    return {
      rank: firstOfGroup.get(combo.group)!,
      combo,
      withinError: combo.group === 0,
      percent: percentOf(combo.delta.mean, result.equipped.mean),
    };
  });
}

const COMBO_KEY_PART_SEPARATOR = '|';

/**
 * One substitution's own contribution to `comboKey` below. An `item` substitution is
 * `slot:item_id`, with an `:enchant:<n>:suffix:<n>` tail ONLY when either is non-zero --
 * the common case (neither set) stays the bare `slot:item_id` form every existing
 * `data-testid` and e2e selector already depends on
 * (`web/tests/e2e/sim-drops.spec.ts`'s `sim-drops-pin-finger1:19325` and
 * `sim-drops-pin-head:16963` among them). A non-item substitution (`talents`, `set`,
 * `consumes`) has no slot or item id at all, so its own kind plus whatever names it --
 * `talents` for a `talents` substitution (the loadout's own build, which is what actually
 * distinguishes two loadouts that happen to share a display name), `name` for `set` and
 * `consumes` -- is what makes it stable and distinct from a substitution of a different
 * kind or a different loadout/set/list.
 *
 * Deliberately NOT `substitutionIdentity` above: that function omits `slot` ON PURPOSE
 * (design 3.2 tries a ring or trinket in both slots and wants both tries to collapse to one
 * de-duped row), which is exactly the ambiguity `comboKey` exists to keep apart for a
 * `data-testid` / `{#each}` key -- two rows `comboIdentity` calls "the same candidate" must
 * still never render the same key once dedupe has already run and both survive as
 * genuinely different rows (which happens whenever their deltas are not bit-identical,
 * since `dedupedCombos` only drops an exact repeat of an EARLIER identity).
 */
function comboKeyPart(sub: Substitution): string {
  if (sub.kind === 'item') {
    const base = `${sub.slot ?? ''}:${sub.item_id ?? ''}`;
    const enchant = sub.enchant ?? 0;
    const suffix = sub.suffix ?? 0;
    return enchant === 0 && suffix === 0 ? base : `${base}:enchant:${enchant}:suffix:${suffix}`;
  }
  if (sub.kind === 'talents') return `talents:${sub.talents ?? ''}`;
  return `${sub.kind}:${sub.name ?? ''}`;
}

/**
 * A row's identity for a `{#each}` key or a `data-testid` suffix: every substitution's own
 * `comboKeyPart`, joined -- not just the first. Review fix round 1: keying on the first
 * substitution alone gave the fixture's own `combos[0]` (head+shoulder) and `combos[1]`
 * (head alone) the identical key "head:16963" once this function reached a component whose
 * rows can carry more than one substitution (`ComboResults.svelte`'s gear-mode rows) --
 * harmless in `DropResults.svelte`, whose "drops" mode combos are always
 * single-substitution, but a real collision once shared, and Task 12's e2e selects rows by
 * this id.
 *
 * A single plain `item` substitution still produces exactly `slot:item_id` (join of one
 * part is that part, unchanged), which is what every existing `data-testid` and e2e
 * selector already depends on. The rank is the fallback only when a combo carries no
 * substitution at all (`substitutions.length === 0`) -- a combination this codebase does
 * not otherwise construct, kept only so this function never throws on one.
 *
 * Moved here from `DropResults.svelte` (task 7): `ComboResults.svelte` needs the identical
 * rule for its own "Plan it" test ids, and this lane's own DRY rule is one function, not a
 * second copy typed again in a second component.
 */
export function comboKey(row: ComboRow): string {
  if (row.combo.substitutions.length === 0) return String(row.rank);
  return row.combo.substitutions.map(comboKeyPart).join(COMBO_KEY_PART_SEPARATOR);
}

/**
 * How many of `result.combos` the de-dupe above folded away -- `result.combos.length` minus
 * `comboRows(result).length`.
 *
 * `store.combinations` (BulkRunBar's own count) is the engine's `simCount` over the
 * SUBMITTED request, before this de-dupe ever runs, so a run with any ring, trinket or
 * weapon ticked can read "N valid combinations" above a table of fewer than N rows -- the
 * same complaint class the picker's own counts already drew (dps D34, "the counts in the
 * picker do not match what gets simulated"). Controller ruling (final whole-branch review,
 * Important 2): explain the gap on the results page rather than recompute either number
 * from the other -- `store.combinations` stays the engine's own count, and the results
 * header stays this file's own row count, so neither can drift from what it actually is.
 */
export function collapsedComboCount(result: BulkResult): number {
  return result.combos.length - comboRows(result).length;
}

/**
 * Engine-lane rule 5's sentinel: a two-hander replacing a main-plus-off-hand pair emits a
 * second substitution `{kind:"item", slot:"off_hand", item_id:0, name:"<item removed>"}`.
 * `item_id: 0` is never a real item (simdb has no id 0).
 *
 * Exported because rule 5 says the emptied slot must read as emptied *everywhere* a player
 * can see it -- the chips, the "By slot" panel, the addon string and the planner link --
 * and a second, inline copy of this predicate in a component is how the exported paths came
 * to miss it (final whole-branch review, Important 3).
 */
export function isEmptiedOffHand(sub: Substitution): boolean {
  return sub.kind === 'item' && sub.slot === 'off_hand' && sub.item_id === 0;
}

/**
 * One combination's substitutions written over a base gear list.
 *
 * The rule-5 sentinel REMOVES its slot rather than writing `item_id: 0` over it: this list
 * is what `addon-export.ts` turns into a paste-into-the-game string and what
 * `codeForCharacterSpec` encodes into a planner link, and neither the addon grammar nor
 * FS1 defines 0 as "empty" -- a decoder reads `off_hand=0` as an item that does not exist
 * (final whole-branch review, Important 1).
 *
 * Every step returns a new list: `baseGear` is copied once at the top and never written
 * through -- `winningGear` below and `planItHref`'s own per-row gear both rely on that to
 * leave the caller's own list untouched.
 */
export function gearForCombo(baseGear: readonly GearSlot[], combo: Combo): GearSlot[] {
  let gear: GearSlot[] = baseGear.map((slot) => ({ ...slot }));
  for (const sub of combo.substitutions) {
    if (sub.kind !== 'item' || sub.slot === undefined || sub.item_id === undefined) continue;
    if (isEmptiedOffHand(sub)) {
      gear = gear.filter((slot) => slot.slot !== sub.slot);
      continue;
    }
    const next: GearSlot = { slot: sub.slot, item_id: sub.item_id };
    if (sub.enchant !== undefined && sub.enchant > 0) next.enchant = sub.enchant;
    if (sub.suffix !== undefined && sub.suffix > 0) next.suffix = sub.suffix;
    const at = gear.findIndex((slot) => slot.slot === sub.slot);
    gear = at >= 0 ? gear.map((slot, index) => (index === at ? next : slot)) : [...gear, next];
  }
  return gear;
}

/**
 * The leader's substitutions written over the base character's gear.
 *
 * Reads `result.combos[0]` directly, not the de-duplicated list `comboRows` builds: index 0
 * is always the FIRST occurrence of its own identity (nothing ranked ahead of it could share
 * it), so a de-dupe can never change which combo is at index 0 or what it contains.
 *
 * A result with no combos at all (an empty run) has no leader to write over the base gear
 * with, so this reads back the base character's own gear -- a copy of what is genuinely
 * equipped, never a fabricated `Combo` standing in for "no winner".
 */
export function winningGear(result: BulkResult): GearSlot[] {
  const leader = result.combos[0];
  return leader === undefined
    ? result.request.character.gear.map((slot) => ({ ...slot }))
    : gearForCombo(result.request.character.gear, leader);
}

/**
 * Whether `planItHref` below has anything to open for this row. An `item` substitution
 * changes gear, so `gearForCombo` can write it over the base set. A `talents` substitution
 * changes the loadout instead, so its own `Substitution.talents` string -- the engine's
 * talents string, exactly `CharacterSpec.talents` -- can ride onto the encoded spec in its
 * place, gear untouched.
 *
 * A `set` substitution carries only the set's `name` (contract 10.1's `Substitution` has no
 * gear field -- `sim/bulk`'s own `apply` never puts a named set's items into `subs`, only
 * its label), so a row naming one can never be reconstructed without silently opening the
 * base gear and calling it the set the row actually won with. Excluded whether it is the
 * row's only substitution or mixed with others, since the row's true gear is unknowable
 * either way. A `consumes` substitution changes neither gear nor talents at all.
 */
export function canPlanCombo(combo: Combo): boolean {
  const kinds = new Set(combo.substitutions.map((sub) => sub.kind));
  if (kinds.has('set')) return false;
  return kinds.has('item') || kinds.has('talents');
}

/**
 * "Plan it" (design 1): the base character with this one row's own change opened in the
 * planner, through the same `plannerHrefForSpec`/`codeForCharacterSpec` conversion every
 * other "Open in planner" link uses (`character.ts`) -- one encoder, so no two links can
 * ever disagree about what a code encodes. Null when `canPlanCombo` says the row has
 * nothing a planner link can honestly open (see its own doc comment); the caller draws no
 * link rather than one that opens gear or a loadout this row never actually tried.
 */
export function planItHref(result: BulkResult, combo: Combo, treeVersion: string): string | null {
  if (!canPlanCombo(combo)) return null;
  const talents = combo.substitutions.find((sub) => sub.kind === 'talents')?.talents;
  const spec: CharacterSpec = {
    ...result.request.character,
    gear: gearForCombo(result.request.character.gear, combo),
    ...(talents === undefined ? {} : { talents }),
  };
  return plannerHrefForSpec(spec, treeVersion);
}

export interface SlotSummaryRow {
  slot: string;
  item_id: number;
  /** The item's own name, off the substitution (contract 10.1 A6). Never a re-join. */
  name: string;
  /**
   * What this slot contributed on its own, from the combination that changed only this
   * slot. Null when no such combination exists -- a Droptimizer run has one per slot by
   * construction, a Top Gear run has one whenever the slot's item was also tried alone, and
   * a run where it was not is honest about not knowing rather than apportioning the total.
   */
  gain: number | null;
}

/**
 * Design 3.3's per-slot summary: what the winner uses, and what that slot was worth.
 *
 * Reads `result.combos` directly, not `comboRows`' de-duplicated list: `winner` is index 0
 * (unaffected by a de-dupe, same reasoning as `winningGear`), and the `alone` map below is
 * keyed by `${slot}:${item_id}` -- SLOT included, unlike a combo's own de-dupe identity --
 * so a finger1/finger2 duplicate pair writes two distinct keys and can never collide with
 * the one key `winner`'s own substitutions look up. A literal duplicate (the same slot
 * twice) would overwrite its own key with an identical value, which is a no-op. Either way,
 * de-duplicating first could not change this function's answer.
 */
export function slotSummary(result: BulkResult): SlotSummaryRow[] {
  const winner = result.combos[0];
  if (winner === undefined) return [];
  const alone = new Map<string, number>();
  for (const combo of result.combos) {
    if (combo.substitutions.length !== 1) continue;
    const only = combo.substitutions[0];
    if (only.kind !== 'item' || only.slot === undefined) continue;
    alone.set(`${only.slot}:${only.item_id}`, combo.delta.mean);
  }
  return winner.substitutions
    .filter(
      (sub): sub is Substitution & { slot: string; item_id: number } =>
        sub.kind === 'item' && sub.slot !== undefined && sub.item_id !== undefined,
    )
    .map((sub) => ({
      slot: sub.slot,
      item_id: sub.item_id,
      // Engine-lane rule 5: the emptied off-hand is not an item, and "<item removed>" is
      // not a name a player should ever read verbatim. `SubstitutionChips` reads the same
      // exported predicate, so the chips and this "By slot" panel cannot disagree.
      name: isEmptiedOffHand(sub) ? bulkCopy.offHandEmptied : substitutionLabel(sub),
      gain: alone.get(`${sub.slot}:${sub.item_id}`) ?? null,
    }));
}

/**
 * Design 4.4's results filter: "only combos keeping 4-piece". Tier bonuses are counted by
 * the engine from the gear itself, so this is a filter over what a combination would be
 * wearing, never an input to the run.
 */
export function keepsSetBonus(
  combo: Combo,
  result: BulkResult,
  items: ReadonlyMap<number, Item>,
  sets: readonly ItemSet[],
  pieces: number,
): boolean {
  const gear = new Map(result.request.character.gear.map((slot) => [slot.slot, slot.item_id]));
  for (const sub of combo.substitutions) {
    if (sub.kind !== 'item' || sub.slot === undefined || sub.item_id === undefined) continue;
    gear.set(sub.slot, sub.item_id);
  }
  const counts = new Map<number, number>();
  for (const itemId of gear.values()) {
    const setId = items.get(itemId)?.set_id ?? null;
    if (setId === null) continue;
    counts.set(setId, (counts.get(setId) ?? 0) + 1);
  }
  return sets.some((set) => (counts.get(set.id) ?? 0) >= pieces);
}

/**
 * Whether `keepsSetBonus` would ever be true for this result at all -- dps D24:
 * ComboResults.svelte used to print "Only combinations keeping a 4-piece set bonus" and
 * its checkbox over every result, including a character with no 4-piece set anywhere in
 * play, where ticking it can only ever empty the table. The checkbox means nothing to a
 * player it can never apply to, so the page shows it only when at least one combo in the
 * result reaches the piece count.
 */
export function anyKeepsSetBonus(
  result: BulkResult,
  items: ReadonlyMap<number, Item>,
  sets: readonly ItemSet[],
  pieces: number,
): boolean {
  return result.combos.some((combo) => keepsSetBonus(combo, result, items, sets, pieces));
}

/**
 * Rows whose delta is bit-identical to at least one other row here -- not `combo.group`'s
 * own "within error" test (a statistical closeness the page already draws a rule under),
 * but the same mean AND the same error to the decimal, which only happens when two
 * candidates leave the simulated character in an identical state: neither's stats touch
 * anything this spec's damage depends on (dps D36: three different necks, none of them
 * carrying a melee stat, all landing on the same `-19 ± 2.3`). `dedupedCombos` above only
 * ever drops an exact repeat of the same item id, so distinct items always keep distinct
 * rows here -- these are real ties, grouped so the page can say why once per group instead
 * of leaving unexplained duplicate numbers on screen.
 */
export interface OverlappingRowPair {
  a: ComboRow;
  b: ComboRow;
}

/** Two delta intervals (mean ± 1 SE), the same interval `group` (bulk's own Go ranking)
 *  and the page's own error bars already use, touching or crossing. */
function deltaIntervalsOverlap(a: Estimate, b: Estimate): boolean {
  return a.mean - a.error <= b.mean + b.error && b.mean - b.error <= a.mean + a.error;
}

/**
 * The first adjacent pair of ranked rows whose delta intervals actually overlap, or null
 * when every adjacent gap is clean -- newcomer round 4 (review.md:83-110) / dps D25-
 * pattern: "These runs are too close to separate..." used to print under every result
 * with more than one row, whatever the actual gap, including a −37.4% difference at a
 * ±1.5/±1.6 margin (about 150 times its own error bar). Adjacent pairs, not every pair:
 * two non-adjacent rows both overlapping a shared middle row are not "close to separate"
 * from each other in the ranking a player reads top to bottom, and comparing every pair
 * would flag rows whose own rank already separates them by a mile.
 */
export function closestOverlappingPair(rows: readonly ComboRow[]): OverlappingRowPair | null {
  for (let i = 0; i + 1 < rows.length; i++) {
    if (deltaIntervalsOverlap(rows[i].combo.delta, rows[i + 1].combo.delta)) {
      return { a: rows[i], b: rows[i + 1] };
    }
  }
  return null;
}

/**
 * A row's own name, for the overlap note above: the leading substitution's label, the
 * same one `headlineFor` reads for the leader's own name ("Helm of Wrath", "Build 1"
 * -- a talents-mode substitution's `name` is the loadout's own name). Falls back to the
 * row's rank only for a combo this codebase does not otherwise construct (no
 * substitutions at all), the same edge `comboKey` above already guards.
 */
export function rowName(row: ComboRow): string {
  const first = row.combo.substitutions[0];
  const label = first === undefined ? '' : substitutionLabel(first);
  return label === '' ? `#${row.rank}` : label;
}

export function exactTieGroups(rows: readonly ComboRow[]): ComboRow[][] {
  const byDelta = new Map<string, ComboRow[]>();
  for (const row of rows) {
    const key = `${row.combo.delta.mean}:${row.combo.delta.error}`;
    const group = byDelta.get(key);
    if (group === undefined) byDelta.set(key, [row]);
    else group.push(row);
  }
  return [...byDelta.values()].filter((group) => group.length > 1);
}

/**
 * "Helm of Wrath", "Deep Fury", "AQ set" -- whatever this substitution actually changed.
 *
 * It reads `name`, which contract 10.1 A6 fills for items too, from simdb. There is no
 * item-id join here and no item map parameter: the result already carries every name it
 * needs, and re-deriving one from a per-class item file the reader may not have loaded is
 * how a saved page ends up showing "Item 16963" for something the engine named.
 */
export function substitutionLabel(sub: Substitution): string {
  // Contract 10.8: a `consumes` substitution's name is the ids joined by ", ", which is a
  // list and not a phrase, so it is the one kind that goes through copy on its way out.
  if (sub.kind === 'consumes') return bulkCopy.consumesChip(sub.name ?? '');
  if (sub.name !== undefined && sub.name !== '') return sub.name;
  return sub.kind === 'item' ? `Item ${sub.item_id ?? 0}` : '';
}

/** Where a Droptimizer combination's one substitution came from, in words. */
export function sourceNameOfCombo(combo: Combo): string {
  return combo.substitutions[0]?.source_name ?? '';
}

/**
 * `substitutionLabel`, with the substitution's own source (a boss or vendor name) folded
 * in when it carries one. Task 5 (newcomer MAJOR, review.md:360-363): SubstitutionChips.svelte
 * used to put this exact string in a `title`, which a phone can never hover to read. This is
 * the same "name · source" phrase that title carried, now the chip's own visible text.
 */
export function substitutionChipLabel(sub: Substitution): string {
  const label = substitutionLabel(sub);
  return sub.source_name !== undefined && sub.source_name !== '' ? `${label} · ${sub.source_name}` : label;
}

/**
 * "+41 DPS from Helm of Wrath", for the page's own display. The leader's first substitution
 * is the one named -- the planner orders a combination's substitutions by the slot's own
 * contribution, so the first is the one that carried it.
 *
 * The *stored* headline on a saved sim is not this: contract 10.6 has the API compose it
 * at save time, with its own rules for several substitutions ("… and 2 more") and for an
 * empty result ("no combinations"). `SimListRow.headline` is read, never recomputed.
 *
 * Reads `result.combos[0]` directly, not `comboRows`' de-duplicated list, for the same
 * reason `winningGear` does: index 0 is always the first occurrence of its own identity, so
 * a de-dupe can never change it.
 */
export function headlineFor(result: BulkResult): string {
  const winner = result.combos[0];
  if (winner === undefined) return '';
  const gain = deltaLabel(winner.delta).split(' ')[0];
  const what = substitutionLabel(winner.substitutions[0]);
  return what === '' ? `${gain} DPS` : `${gain} DPS from ${what}`;
}

/** The after-sim sentence's own data shape, off a bulk result's top row -- null when the
 *  result has no genuine upgrade (DropResults.svelte's own `> 0` rule). Exported so
 *  Droptimizer.svelte's write-on-success effect, and its own test, share one function
 *  rather than the effect re-deriving this inline. */
export function topUpgradeOf(
  result: BulkResult,
): { itemName: string; sourceName: string; gain: string } | null {
  const top = comboRows(result)[0];
  if (top === undefined || top.combo.delta.mean <= 0) return null;
  const sourceName = sourceNameOfCombo(top.combo);
  return {
    itemName: substitutionLabel(top.combo.substitutions[0]),
    sourceName: sourceName === '' ? 'an unknown source' : sourceName,
    gain: deltaLabel(top.combo.delta),
  };
}
