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
// A key the resolver's own grammar does not recognise -- a real display name from a
// logged fight, `parseActionKey` returning null -- is returned unchanged; that is compare
// mode's own contract, where a logged fight's rows carry real names and must pass through
// untouched. A key the grammar DOES recognise but whose id the build's table does not
// carry -- the names file is still loading, or a build that never published this id --
// is never returned raw either (dps-minmaxer review round 2, D48: a saved run rendered
// "spell:20662" as an ability name): it falls back to humaniseKey (humanise.ts), so the
// screen always reads English, never a wire key.
import { dataUrl, fetchJson, loadOptional } from '../planner/load';
import { attackHandName, simCopy } from './copy';
import { humanise, humaniseKey } from './humanise';

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
  /**
   * Which `spell` ids the player's OWN class's spellconst carries -- before
   * `loadActionNames` folds in `simnames/_shared.json`'s cross-class raid buffs (2026-09-21
   * result-page review, Defect 2). sentence.ts reads this to tell "my own cooldown" from
   * "an external raid buff someone else in the raid granted me" without re-deriving
   * RAID_BUFFS' slug list here; every other reader of an `ActionNames` (resolveActionName
   * itself, compare.ts, sample-log.ts, SpecCard.svelte) ignores it. Optional, and treated as
   * "everything in `spell` is the player's own" when absent, which is what every literal
   * `ActionNames` built by hand (a test, a fixture) already means.
   */
  ownSpell?: ReadonlySet<string>;
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
 * The hand an auto-attack tag names -- wowsims/classic's own AutoAttacks constants
 * (sim/core/attack.go): `tagMainhand = 1`, `tagOffhand = 2`, `tagExtraAttack = 3`. This is
 * the only place either number is read: resolveActionName below and sentence.ts's own
 * prose both call it rather than pattern-matching a rendered name, so a hand can never
 * read one way in a table and a different way in the summary sentence. A tag this build
 * has never seen -- 0, the untagged placeholder a synthetic fixture still uses, or a
 * future engine tag -- returns null, and the caller falls back to a humanised label.
 */
export type AttackHand = 'main' | 'off' | 'extra';

export function attackHand(tag: number): AttackHand | null {
  switch (tag) {
    case 1:
      return 'main';
    case 2:
      return 'off';
    case 3:
      return 'extra';
    default:
      return null;
  }
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
  if (parsed.kind === 'other') {
    if (parsed.label === 'attack') {
      const hand = attackHand(parsed.tag);
      if (hand !== null) return attackHandName[hand];
    }
    return humanise(parsed.label) + simCopy.actionVariant(parsed.tag, parsed.rank);
  }
  const table = parsed.kind === 'spell' ? names?.spell : names?.item;
  const name = table?.[parsed.label] ?? humaniseKey(`${parsed.kind}:${parsed.label}`);
  return name + simCopy.actionVariant(parsed.tag, parsed.rank);
}

/** `simnames/_shared.json`'s empty shape: an older build that scripts/sync-data.mjs has
 *  not regenerated ships none of it, and a missing cross-class table must read as "no
 *  extra names available", never as a load failure the page has to report. */
const EMPTY_SHARED_NAMES: ActionNames = { spell: {}, item: {} };

/**
 * `simnames/_shared.json`: every class's own spell table unioned into one, plus the
 * handful of raid buffs no class's spellbook owns at all (writeSimNames's own doc comment,
 * scripts/sync-data.mjs). A 404 means the build predates this file, not that the build is
 * broken, so it falls back to empty via loadOptional the same way loadSets and loadWeights
 * already treat their own optional build files.
 */
function loadSharedActionNames(build: string): Promise<ActionNames> {
  return loadOptional<ActionNames>(dataUrl(build, 'simnames/_shared.json'), EMPTY_SHARED_NAMES);
}

/**
 * The build's name table for one class, published by scripts/sync-data.mjs -- the class's
 * own spellconst-derived names, with the cross-class `simnames/_shared.json` folded in so
 * a raid buff from another class (a mage raid-buffed with a paladin's Blessing of Kings, a
 * warrior's Battle Shout, Thorns) resolves too (2026-09-21 result-page review, Defect 2).
 * The class's own entry wins a same-id collision -- spell ids are one global namespace, so
 * in practice there is none, but a class's own name is the more specific one to trust if
 * there ever were. `ownSpell` records the pre-merge membership, for sentence.ts's "my own
 * spec's buff, not an external raid buff" preference.
 */
export async function loadActionNames(build: string, classSlug: string): Promise<ActionNames> {
  const [own, shared] = await Promise.all([
    fetchJson<ActionNames>(dataUrl(build, `simnames/${classSlug}.json`)),
    loadSharedActionNames(build),
  ]);
  return {
    spell: { ...shared.spell, ...own.spell },
    item: { ...shared.item, ...own.item },
    ownSpell: new Set(Object.keys(own.spell)),
  };
}
