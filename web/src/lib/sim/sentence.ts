// web/src/lib/sim/sentence.ts
// The one plain-language line above the results, and the name rewrite the report
// components need under it.
//
// A sim summary's ability, aura and cast names are engine action keys (Task 23), so both
// functions here take the build's name table. namedSummary returns a COPY: the stored
// SimResult keeps the engine's keys, because that is what POST /v1/sims saves and what a
// compare against a later engine version matches on.
import { resolveActionName, type ActionNames } from './action-names';
import { simCopy } from './copy';
import type { Actor, Summary } from '../report/types';

/** A resolved name the sentence says differently from a table. Both sets live in copy.ts. */
function prose(name: string): string {
  return simCopy.proseNames[name] ?? name;
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

  const named = namedSummary(summary, names);
  const namedActor = named.damage_done.find((row) => row.guid === actor.guid) ?? actor;
  const top = [...namedActor.abilities].sort((a, b) => b.total - a.total).slice(0, 2);
  const share = percent(
    top.reduce((sum, ability) => sum + ability.total, 0),
    actor.total,
  );
  const damage =
    top.length >= 2
      ? `${prose(top[0].name)} and ${prose(top[1].name)} are ${share}% of your damage`
      : `${prose(top[0].name)} is ${share}% of your damage`;

  const aura = [...named.auras]
    .filter((track) => track.target_guid === actor.guid && track.type === 'BUFF')
    .sort((a, b) => b.uptime_ms - a.uptime_ms)[0];
  if (aura === undefined || summary.duration_ms <= 0) return `${damage}.`;

  return `${damage}; ${aura.name} is up ${percent(aura.uptime_ms, summary.duration_ms)}% of the fight.`;
}
