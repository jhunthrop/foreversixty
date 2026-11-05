// web/src/lib/sim/aura-card.ts
// The Buffs and debuffs card's two columns (spec 2026-11-04-sim-results 3.10, build rulings R1,
// R2): the engine's tracks with its bookkeeping left out, plus every buff the request named that
// has no track, so the card holds the same full list as Settings and files it by the same rule.
// Everything here reads a row's raw key, never its resolved name, except the join to simbuffs,
// which has only names to join on.
import type { AuraTrack } from '../report/types';
import { parseActionKey, resolveActionName, type ActionNames } from './action-names';
import { hasNoDamageEffect, type SpecWeights } from './buff-effect';
import { IMPROVED_SUFFIX } from './buff-id';
import { humanise } from './humanise';
import { aurasCardCopy as copy } from './input-cards-copy';
import type { SimBuffRow } from './sim-buffs';
import type { SimResult } from './types';

export interface AuraEntry {
  key: string;
  name: string;
  /** The track's action key, for its icon and tooltip; null for a buff with no track. */
  actionKey: string | null;
  /** The simbuffs icon stem of the buff this row is, when the join found one. */
  buffIcon: string | null;
  uptimePct: number;
  applications: number;
  /** A stacking debuff's stacks; 0 when it does not stack. */
  stacks: number;
}

export interface AuraColumn {
  /** Rows that come and go, longest uptime first. */
  bars: AuraEntry[];
  upAllFight: AuraEntry[];
  noEffect: AuraEntry[];
}

export interface AuraCard {
  player: AuraColumn;
  target: AuraColumn;
  /** The run kept no target tracks: the column is the request's own list and says so. */
  targetFromRequest: boolean;
}

/** At or above this, an aura leaves the bar list for the Up all fight row. */
const UP_ALL_FIGHT_PCT = 99.9;

/** The extra-attack counter (21919), never a buff the player sees; older results still carry it. */
const THRASH = 21919;
/** Overpower's window lights a button; the client has no such aura (spec 14). */
const OVERPOWER_RANKS: ReadonlySet<number> = new Set([7384, 7887, 11584, 11585]);
/** A Heroic Strike's tagged track is the swing queue, drawn on the timeline. */
const HEROIC_STRIKE_RANKS: ReadonlySet<number> = new Set([78, 284, 285, 1608, 11564, 11565, 11566, 11567, 25286]);
/** Crusader's proc: `/1` is the main hand, `/2` the off hand (confirmed in the engine). */
const HOLY_STRENGTH = 20007;
const HAND_TAGS: Readonly<Record<number, string>> = { 1: copy.mainHand, 2: copy.offHand };

function isShown(track: AuraTrack): boolean {
  if (track.uptime_ms <= 0) return false;
  const parsed = parseActionKey(track.name);
  if (parsed === null || parsed.kind === 'other' || parsed.kind === 'unknown') return false;
  if (parsed.id === THRASH || OVERPOWER_RANKS.has(parsed.id)) return false;
  return !(parsed.tag !== 0 && HEROIC_STRIKE_RANKS.has(parsed.id));
}

function trackName(track: AuraTrack, names: ActionNames | null): string {
  const parsed = parseActionKey(track.name);
  if (parsed === null) return track.name;
  const base = resolveActionName(`${parsed.kind}:${parsed.label}`, names);
  const hand = parsed.id === HOLY_STRENGTH ? HAND_TAGS[parsed.tag] : undefined;
  return hand === undefined ? base : `${base} ${hand}`;
}

const baseId = (id: string): string =>
  id.endsWith(IMPROVED_SUFFIX) ? id.slice(0, -IMPROVED_SUFFIX.length) : id;

/** A debuff of an exclusive group that a lower rank in the same request holds off. */
function isSuppressed(id: string, requested: readonly string[], buffs: Readonly<Record<string, SimBuffRow>>): boolean {
  const own = buffs[baseId(id)]?.exclusive_with;
  if (own === undefined || own === null) return false;
  return requested.some((other) => {
    const theirs = buffs[baseId(other)]?.exclusive_with;
    return theirs !== undefined && theirs !== null && theirs.group === own.group && theirs.rank < own.rank;
  });
}

interface Input {
  result: SimResult;
  buffs: Readonly<Record<string, SimBuffRow>>;
  /** Null until read, or when the build has no row for the spec: nothing is filed apart. */
  weights: SpecWeights | null;
  schools: readonly string[];
  names: ActionNames | null;
  /** The engine's own Debuffs list, for a build whose simbuffs has no `target` flag yet. */
  fallbackTargetIds: ReadonlySet<string>;
}

function emptyColumn(): AuraColumn {
  return { bars: [], upAllFight: [], noEffect: [] };
}

function filed(entries: readonly AuraEntry[], quiet: (entry: AuraEntry) => boolean): AuraColumn {
  const loud = entries.filter((entry) => !quiet(entry));
  const byUptime = (a: AuraEntry, b: AuraEntry): number => b.uptimePct - a.uptimePct;
  return {
    bars: loud.filter((entry) => entry.uptimePct < UP_ALL_FIGHT_PCT).sort(byUptime),
    upAllFight: loud.filter((entry) => entry.uptimePct >= UP_ALL_FIGHT_PCT),
    noEffect: entries.filter(quiet),
  };
}

export function auraCard({ result, buffs, weights, schools, names, fallbackTargetIds }: Input): AuraCard {
  const requested = result.request.character.buffs;
  const idByName = new Map(
    requested.flatMap((id): [string, string][] => {
      const name = buffs[baseId(id)]?.name;
      return name === undefined ? [] : [[name, id]];
    }),
  );
  const seen = new Set<string>();
  const buffIdOf = new Map<string, string>();
  const duration = result.summary.duration_ms;

  const fromTrack = (track: AuraTrack): AuraEntry => {
    const name = trackName(track, names);
    const id = idByName.get(name);
    if (id !== undefined) {
      seen.add(id);
      buffIdOf.set(track.name, id);
    }
    return {
      key: `${track.target_guid}|${track.name}`,
      name,
      actionKey: track.name,
      buffIcon: id === undefined ? null : (buffs[baseId(id)]?.icon ?? null),
      uptimePct: duration <= 0 ? 0 : Math.min(100, (track.uptime_ms / duration) * 100),
      applications: Math.round(track.applications),
      stacks: track.max_stacks > 1 ? track.max_stacks : 0,
    };
  };

  const tracks = result.summary.auras.filter(isShown);
  const player = tracks.filter((track) => track.type === 'BUFF').map(fromTrack);
  const target = tracks.filter((track) => track.type === 'DEBUFF').map(fromTrack);

  const onTarget = (id: string): boolean => buffs[baseId(id)]?.target ?? fallbackTargetIds.has(baseId(id));
  const fromRequest = (id: string): AuraEntry => ({
    key: `request|${id}`,
    name: `${buffs[baseId(id)]?.name ?? humanise(baseId(id))}${id === baseId(id) ? '' : copy.improved}`,
    actionKey: null,
    buffIcon: buffs[baseId(id)]?.icon ?? null,
    uptimePct: 100,
    applications: 0,
    stacks: 0,
  });
  const untracked = requested.filter((id) => !seen.has(id) && !isSuppressed(id, requested, buffs));

  const quiet = (entry: AuraEntry): boolean => {
    if (weights === null) return false;
    const id = entry.actionKey === null ? entry.key.slice('request|'.length) : buffIdOf.get(entry.actionKey);
    return id !== undefined && hasNoDamageEffect(buffs[baseId(id)], weights, schools);
  };

  return {
    player: filed([...player, ...untracked.filter((id) => !onTarget(id)).map(fromRequest)], quiet),
    target:
      target.length + untracked.filter(onTarget).length === 0
        ? emptyColumn()
        : filed([...target, ...untracked.filter(onTarget).map(fromRequest)], quiet),
    targetFromRequest: target.length === 0,
  };
}
