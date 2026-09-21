// web/src/lib/sim/sentence.ts
// The one plain-language line above the results, and the name rewrite the report
// components need under it.
//
// A sim summary's ability, aura and cast names are engine action keys (Task 23), so both
// functions here take the build's name table. namedSummary returns a COPY: the stored
// SimResult keeps the engine's keys, because that is what POST /v1/sims saves and what a
// compare against a later engine version matches on.
import { attackHand, parseActionKey, resolveActionName, type ActionNames } from './action-names';
import { sanitizeAuraTracks } from './aura-rows';
import { attackHandProse, simCopy } from './copy';
import type { Actor, Summary } from '../report/types';

/**
 * The sentence's own phrasing for one raw action key -- "main-hand white hits", not the
 * table's "Main-hand attacks". This reads the SAME parsed tag resolveActionName does,
 * straight off `key`, never off the name that function already resolved: a hand read from
 * a rendered string is exactly the bug this function replaces (Lane W1, copy.ts's own
 * comment on attackHandProse). `resolvedName` is only the fallback for the "other" actions
 * that are not the tagged auto-attack.
 */
function prose(key: string, resolvedName: string): string {
  const parsed = parseActionKey(key);
  if (parsed?.kind === 'other' && parsed.label === 'attack') {
    const hand = attackHand(parsed.tag);
    if (hand !== null) return attackHandProse[hand];
  }
  return simCopy.proseNames[resolvedName] ?? resolvedName;
}

/**
 * The same summary with every ability, aura and cast name resolved. Nothing else changes,
 * and the input is never mutated. Auras additionally pass through sanitizeAuraTracks
 * (aura-rows.ts, Task 4) first: the engine seeds a row for every rank or metric variant a
 * spec could register, whether it ever fired or not, and this is the one place every
 * report component under SimResults.svelte reads auras from, so the cleanup happens once
 * rather than per-tab.
 */
export function namedSummary(summary: Summary, names: ActionNames | null): Summary {
  const name = (key: string): string => resolveActionName(key, names);
  return {
    ...summary,
    damage_done: summary.damage_done.map((actor) => ({
      ...actor,
      abilities: actor.abilities.map((ability) => ({ ...ability, name: name(ability.name) })),
    })),
    auras: sanitizeAuraTracks(summary.auras).map((track) => ({ ...track, name: name(track.name) })),
    casts: summary.casts.map((row) => ({ ...row, spell_name: name(row.spell_name) })),
  };
}

/**
 * Clamped to 100 as a last line of defence (2026-09-21 result-page review round 3, E8: a
 * combined sim result's aura rows read up to 101% here, sim/combine's own bug, now fixed
 * at the source) -- this is the one place every uptime and every damage share in the
 * sentence is rounded, so clamping it once covers both without either caller needing to
 * know why the source disagreed.
 */
function percent(part: number, whole: number): number {
  return whole <= 0 ? 0 : Math.min(Math.round((part / whole) * 100), 100);
}

/**
 * True for an aura the player's OWN class grants (or when there is no class-vs-shared
 * signal to read at all -- `names` not loaded yet, or a hand-built `ActionNames` with no
 * `ownSpell`, which is every fixture written before this preference existed and which
 * means exactly "everything here is the player's own"). False only for an aura
 * `loadActionNames` resolved through the cross-class `simnames/_shared.json` table -- a
 * raid buff from another class's spellbook, or one of the handful no class owns at all
 * (Rallying Cry of the Dragonslayer, Sayge's Fortune...). 2026-09-21 result-page review,
 * Defect 2: the sentence used to pick whichever BUFF aura had the highest uptime, and a
 * raid buff applied once at pull and never dropped always wins that contest against the
 * player's own, shorter-lived cooldowns -- which is how "Thorns is up 100% of the fight"
 * reached the summary sentence of a Frost Mage that never cast Thorns.
 */
function isOwnSpecAura(key: string, names: ActionNames | null): boolean {
  const parsed = parseActionKey(key);
  if (parsed === null || parsed.kind !== 'spell') return true;
  if (names?.ownSpell === undefined) return true;
  return names.ownSpell.has(parsed.label);
}

export function playerActor(summary: Summary): Actor | null {
  if (summary.damage_done.length === 0) return null;
  return summary.damage_done.reduce((best, actor) => (actor.total > best.total ? actor : best));
}

export function summarySentence(summary: Summary, names: ActionNames | null): string {
  const actor = playerActor(summary);
  if (actor === null || actor.total <= 0) return simCopy.noDamage;

  // The raw (unresolved) abilities, not namedSummary's copy: prose() needs each ability's
  // own key to read its tag, and a name already resolved to a display string cannot be
  // read back into a tag (Lane W1's own fix -- see prose's comment above).
  const top = [...actor.abilities].sort((a, b) => b.total - a.total).slice(0, 2);
  const share = percent(
    top.reduce((sum, ability) => sum + ability.total, 0),
    actor.total,
  );
  const phrase = (ability: (typeof top)[number]): string =>
    prose(ability.name, resolveActionName(ability.name, names));
  const damage =
    top.length >= 2
      ? `${phrase(top[0])} and ${phrase(top[1])} are ${share}% of your damage`
      : `${phrase(top[0])} is ${share}% of your damage`;

  // sanitizeAuraTracks, not the raw array: the BUFFS tab reads auras through it (namedSummary
  // above), and the sentence must pick the same row it does -- an engine-internal row
  // (other:move) it would drop, or a tag-variant row it would fold into a bigger one with a
  // summed uptime, is exactly the aura this line then names (final whole-branch review,
  // Finding 1 -- a regression from before the BUFFS/DEBUFFS work, when both read one array).
  //
  // isOwnSpecAura filters to the player's OWN class before the uptime sort: an external
  // raid buff (Thorns, Blessing of Kings, Arcane Brilliance...) is almost always up for the
  // whole fight and would otherwise always win the sort against the player's own, shorter
  // cooldowns. No own-spec buff aura is not "fall back to whichever raid buff is up
  // longest" -- it is nothing to say about uptime at all (2026-09-21 result-page review,
  // Defect 2).
  const aura = sanitizeAuraTracks(summary.auras)
    .filter((track) => track.target_guid === actor.guid && track.type === 'BUFF')
    .filter((track) => isOwnSpecAura(track.name, names))
    .sort((a, b) => b.uptime_ms - a.uptime_ms)[0];
  if (aura === undefined || summary.duration_ms <= 0) return `${damage}.`;

  const auraName = resolveActionName(aura.name, names);
  return `${damage}; ${auraName} is up ${percent(aura.uptime_ms, summary.duration_ms)}% of the fight.`;
}
