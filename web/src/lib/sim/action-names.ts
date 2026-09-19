// web/src/lib/sim/action-names.ts
// The engine names nothing. sim/adapter/identity.go turns an engine ActionID into a row id
// and a key -- "spell:<id>", "spell:<id>/<tag>", "spell:<id>+r<rank>", "item:<id>",
// "other:<snake_name>", "unknown", with the tag and rank suffixes on any of them -- and
// says so in its own comment: "the engine's ActionIDs carry no names ... so the web can
// resolve a real name from the build's own spells.json". This is that resolution, and every
// table, sentence and compare row on this lane goes through it. No component renders
// `ability.name` bare.
//
// Two rules this module lives under, both the adapter's:
//
//   * The ROW KEY is `spell_id`, never the name. An untagged spell keeps its client id;
//     everything else takes a derived id at or above 2,000,000, and Summarize refuses to
//     emit two rows sharing one key (ErrDuplicateRow) because Svelte 5 raises a runtime
//     error on a repeated {#each} key. Nothing here re-keys anything.
//   * The NAME carries the client id, which is why the lookup reads the key rather than
//     `spell_id`: "spell:25286/1" has row id 10002025286, and 25286 is what the build's
//     table knows.
//
// A key the resolver does not recognise is returned unchanged. That matters twice: while
// the names file is still loading, and in compare mode, where a logged fight's rows carry
// real display names from the combat log and must pass through untouched.
import { dataUrl, fetchJson } from '../planner/load';
import { simCopy } from './copy';
import { humanise } from './humanise';

/**
 * The first row id sim/adapter allocates for itself (its `syntheticBase`). Every client
 * spell id is below it, so a summary row at or above it is a tagged, ranked, item or
 * "other" action whose id means nothing outside the adapter -- which is what compare mode
 * (Task 16) keys off when it decides whether two rows can be matched by number.
 */
export const SYNTHETIC_ID_BASE = 2_000_000;

export interface ActionNames {
  /** Client spell id as a string to display name. */
  spell: Record<string, string>;
  /** Client item id as a string to display name. */
  item: Record<string, string>;
}

export interface ActionKey {
  kind: 'spell' | 'item' | 'other' | 'unknown';
  /** The client id for a spell or an item; 0 for an other or unknown action. */
  id: number;
  /** The raw label between the colon and the suffixes: "25286", "attack", "". */
  label: string;
  /** The engine's metric tag: a second row of one action. 0 for the plain one. */
  tag: number;
  /** The spell's rank, when the engine recorded one. 0 otherwise. */
  rank: number;
}

// spell:25286 | spell:25286/1 | spell:25286+r3 | spell:25286/1+r3 | item:14554 |
// other:attack | other:rage_gain/2 | unknown | unknown/1
const KEY = /^(?:(spell|item|other):([A-Za-z0-9_]+)|(unknown))(?:\/(\d+))?(?:\+r(\d+))?$/;

export function parseActionKey(key: string): ActionKey | null {
  const match = KEY.exec(key);
  if (match === null) return null;
  const kind = (match[3] === undefined ? match[1] : 'unknown') as ActionKey['kind'];
  const label = match[2] ?? '';
  // Only a spell or an item labels itself with a number; "other" labels itself with the
  // engine's own enum name, and "unknown" with nothing at all.
  const numeric = /^\d+$/.test(label);
  if ((kind === 'spell' || kind === 'item') && !numeric) return null;
  return {
    kind,
    id: numeric ? Number.parseInt(label, 10) : 0,
    label,
    tag: match[4] === undefined ? 0 : Number.parseInt(match[4], 10),
    rank: match[5] === undefined ? 0 : Number.parseInt(match[5], 10),
  };
}

/**
 * The name a player reads. `names` is null until the build's file has loaded, and an id the
 * build does not carry -- a racial from a class file we did not fetch, a proc from an item
 * the player does not own in this build -- keeps its key rather than becoming "Unknown".
 */
export function resolveActionName(key: string, names: ActionNames | null): string {
  const parsed = parseActionKey(key);
  if (parsed === null) return key;
  if (parsed.kind === 'unknown') return key;
  if (parsed.kind === 'other') return humanise(parsed.label) + simCopy.actionVariant(parsed.tag, parsed.rank);
  const table = parsed.kind === 'spell' ? names?.spell : names?.item;
  const name = table?.[parsed.label];
  if (name === undefined) return key;
  return name + simCopy.actionVariant(parsed.tag, parsed.rank);
}

/** The build's name table for one class, published by scripts/sync-data.mjs. */
export async function loadActionNames(build: string, classSlug: string): Promise<ActionNames> {
  return fetchJson<ActionNames>(dataUrl(build, `simnames/${classSlug}.json`));
}
