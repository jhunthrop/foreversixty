// web/src/lib/sim/sentence.ts
// The one plain-language line above the results, and the name rewrite the report
// components need under it.
//
// A sim summary's ability, aura and cast names are engine action keys (Task 23), so both
// functions here take the build's name table. namedSummary returns a COPY: the stored
// SimResult keeps the engine's keys, because that is what POST /v1/sims saves and what a
// compare against a later engine version matches on.
import { attackHand, parseActionKey, resolveActionName, type ActionNames } from './action-names';
import { simCopy } from './copy';
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
    if (hand !== null) return simCopy.attackHandProse[hand];
  }
  return simCopy.proseNames[resolvedName] ?? resolvedName;
}

/**
 * The same summary with every ability, aura and cast name resolved. Nothing else changes,
 * and the input is never mutated.
 */
export function namedSummary(summary: Summary, names: ActionNames | null): Summary {
  const name = (key: string): string => resolveActionName(key, names);
  return {
    ...summary,
    damage_done: summary.damage_done.map((actor) => ({
      ...actor,
      abilities: actor.abilities.map((ability) => ({ ...ability, name: name(ability.name) })),
    })),
    auras: summary.auras.map((track) => ({ ...track, name: name(track.name) })),
    casts: summary.casts.map((row) => ({ ...row, spell_name: name(row.spell_name) })),
  };
}

function percent(part: number, whole: number): number {
  return whole <= 0 ? 0 : Math.round((part / whole) * 100);
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

  const aura = [...summary.auras]
    .filter((track) => track.target_guid === actor.guid && track.type === 'BUFF')
    .sort((a, b) => b.uptime_ms - a.uptime_ms)[0];
  if (aura === undefined || summary.duration_ms <= 0) return `${damage}.`;

  const auraName = resolveActionName(aura.name, names);
  return `${damage}; ${auraName} is up ${percent(aura.uptime_ms, summary.duration_ms)}% of the fight.`;
}
