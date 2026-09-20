// web/src/lib/sim/aura-rows.ts
// Turns the engine's raw AuraTrack rows -- one seeded per rank or metric variant a spec
// could register, whether it ever fired or not -- into the player-facing set the
// BUFFS/DEBUFFS tabs render (Task 4).
//
// Every check here reads a row's own structure -- its type, its raw unresolved key, its
// uptime and application count -- and never its resolved display name. That is a
// controller ruling from this same task: Task 1 changed those strings once already, and
// a filter keyed to a label would break silently the next time a name does.
// sanitizeAuraTracks runs BEFORE resolveActionName touches `name`, in sentence.ts's
// namedSummary.
import { parseActionKey } from './action-names';
import type { AuraTrack } from '../report/types';

/**
 * sim/adapter/adapter.go writes `Appliers: []string{u.Name}` for every sim aura row,
 * unconditionally -- the player's own display NAME where AuraTable.svelte's "from X" line
 * expects a GUID. Left alone, the lookup by that name always misses and the row reads
 * "from an unnamed source" even for a buff the player cast on themselves. When an
 * applier's value is exactly this track's own target_name, it denotes the target itself,
 * so it is rewritten to target_guid -- the same self-applied shape a real fight's own
 * combat log already produces, which AuraTable already renders correctly (no "from" line
 * at all for a self-buff).
 */
function normalizeSelfApplied(track: AuraTrack): AuraTrack {
  const appliers = track.appliers.map((applier) =>
    applier === track.target_name ? track.target_guid : applier,
  );
  return { ...track, appliers };
}

/** other:move -- the engine's own movement bookkeeping, never a player-facing aura. */
function isMovementPseudoAura(track: AuraTrack): boolean {
  const parsed = parseActionKey(track.name);
  return parsed !== null && parsed.kind === 'other' && parsed.label === 'move';
}

/**
 * A row with no uptime AND no applications carried no player-facing signal: it never
 * applied and was never up. This is the structural shape of both defects the brief names
 * -- a rank of a spell that never got cast, and a melee ability the engine mistakenly
 * registered as an aura (it has no duration to be "up" for) -- without reading either
 * row's name. `uptime_ms` alone is not enough to guard on: sim/adapter/adapter.go's
 * UptimeMS and Applications are two independently-rounded averages
 * (UptimeSecondsAvg/ProcsAvg), so a real proc that fires but is consumed almost
 * instantly, or that only procs in a small fraction of iterations, can round to 0ms of
 * uptime while `applications` stays positive -- a genuine "this proc'd N times" signal
 * with no recovery path elsewhere (the Casts tab tracks spell casts, not passive procs).
 * Requiring both fields to be zero keeps that row.
 */
function isInert(track: AuraTrack): boolean {
  return track.uptime_ms === 0 && track.applications === 0;
}

/**
 * Rewrites a spell/item row's key and spell_id to their plain (untagged, unranked) form,
 * recovered from the row's own raw key via parseActionKey -- never from its resolved
 * name. fold.go's own comment on foldAuras says the engine "carries two AuraMetrics for
 * one aura the spec registered twice"; Go folds those only when their derived SpellID
 * (which bakes the tag in) happens to collide. A tag or rank never marks a genuinely
 * different aura, only the same one metered twice, so every spell/item aura row is
 * normalized to its base id here, and rows sharing that id then merge into one.
 * Rows the grammar does not resolve to a spell or item id -- "other:" auras, "unknown" --
 * pass through unchanged; there is no base id to fold them onto.
 */
function normalizeSpellIdentity(track: AuraTrack): AuraTrack {
  const parsed = parseActionKey(track.name);
  if (parsed === null || (parsed.kind !== 'spell' && parsed.kind !== 'item')) return track;
  return { ...track, name: `${parsed.kind}:${parsed.id}`, spell_id: parsed.id };
}

/** Sums the fields a merged group's rows contribute; keeps the rest from the first row. */
function mergeGroup(rows: AuraTrack[]): AuraTrack {
  const [first] = rows;
  const appliers: string[] = [];
  for (const row of rows) {
    for (const applier of row.appliers) {
      if (!appliers.includes(applier)) appliers.push(applier);
    }
  }
  return {
    ...first,
    applications: rows.reduce((sum, row) => sum + row.applications, 0),
    uptime_ms: rows.reduce((sum, row) => sum + row.uptime_ms, 0),
    max_stacks: rows.reduce((max, row) => Math.max(max, row.max_stacks), 0),
    appliers,
  };
}

/**
 * Groups rows sharing one target and one (now-normalized) raw key, in first-seen order,
 * and merges each group.
 */
function foldByIdentity(tracks: AuraTrack[]): AuraTrack[] {
  const groups = new Map<string, AuraTrack[]>();
  for (const track of tracks) {
    const key = `${track.target_guid}|${track.name}`;
    const group = groups.get(key);
    if (group === undefined) groups.set(key, [track]);
    else group.push(track);
  }
  return [...groups.values()].map(mergeGroup);
}

/**
 * The player-facing aura rows: engine-internal noise dropped, rank/tag duplicates of one
 * spell folded into one row, and a self-applied source correctly attributed to the
 * target. `type` (BUFF/DEBUFF) is read, never assigned -- a real debuff row survives as a
 * debuff, and a sim's current BUFF-only data is not reclassified into anything else.
 */
export function sanitizeAuraTracks(tracks: AuraTrack[]): AuraTrack[] {
  const normalized = tracks
    .filter((track) => !isMovementPseudoAura(track))
    .map(normalizeSelfApplied)
    .map(normalizeSpellIdentity);
  return foldByIdentity(normalized).filter((track) => !isInert(track));
}
