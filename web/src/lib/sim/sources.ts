// web/src/lib/sim/sources.ts
// The five ways a character reaches the simulator, and the pill that says which. Every one
// of them ends in a SimCharacter, so nothing downstream -- the strip, the request builder,
// the planner link -- knows or cares where it came from. Every one of them also writes the
// site's current-character pointer (current-character-bridge.ts) on success, so a bare
// /sim or /planner load can restore whatever was last loaded, from wherever it came from.
//
//   stored   GET /v1/characters/<region>/<ruleset>/<name>/sim-input -- whatever the API's
//            newest source for that character is. Armory is not one of them yet (simulator
//            contract, sim-input row): today the API answers from the newest addon export or
//            the last ranked fight and says which, so this function never names the source
//            itself -- it copies the one the API returned. The day Armory lands, the pill
//            starts saying "Armory" with no change here.
//   addon    an FS1 string, pasted or pushed by the companion
//   build    GET /v1/builds/<id>, the planner's own record
//   fight    <report_id>:<fight_index>, read out of the fight's COMBATANT_INFO row
//   manual   an unsaved planner build's own FS1 code (a "Sim this build" link)
//
// Each returns a result rather than throwing: every one of these is something a player
// typed or pasted, and the page shows the reason inline beside the field.
import { API_BASE_URL } from '../planner/config';
import { FS1_PREFIX, decodeFS1 } from '../planner/fs1';
import { loadReference, loadTalents } from '../planner/load';
import type { BuildRecord } from '../planner/types';
import { fetchReportMeta, fetchSummary } from '../report/load';
import { gearFromCombatant, treeRanksFromTalents } from '../report/planner-link';
import { classSlugFromName } from '../report/tree-sizes';
import type { RosterRow } from '../report/types';
import { requestEnvelope } from '../account/api';
import type { CharacterPath } from '../characters';
import {
  PENDING_RACE,
  characterFromFs1,
  fromBuildDraft,
  gearSlots,
  specForSplit,
  talentLevel,
  type SimCharacter,
} from './character';
import { fetchSimInput } from './api';
import { simCopy } from './copy';
import { recordCurrentCharacter } from './current-character-bridge';
import type { CharacterSource } from './types';

export interface LoadContext {
  /** The active data build; every /data/<build>/ fetch and every character carries it. */
  treeVersion: string;
  apiBase?: string;
}

export type SourceResult = { ok: true; character: SimCharacter } | { ok: false; message: string };

// A logged-fight ref's optional third part names the exact combatant: a WoW player guid,
// "Player-<realm id>-<hex spawn id>" (logs/engine/units/units.go's own Parse comment: "the
// two shapes are Player-<realm>-<hex uid> and …", confirmed against the report fixtures --
// e.g. "Player-4184-000000A1" in src/fixtures/report/report.json). It arrives through the
// query string -- attacker-controlled input bounded only by url.ts's MAX_REF=128 -- so it is
// matched against this exact shape rather than accepted as `(.+)`: a ref whose third part
// does not fit is invalid, not silently passed through.
//
// REALM_ID_MAX_DIGITS and SPAWN_ID_MAX_HEX_CHARS are generous rather than exact (a real spawn
// id is 8 hex characters) so a longer one is still accepted; the longest ref this can produce
// -- 12 + 1 + 6 + 1 + "Player-".length + 6 + 1 + 16 = 50 characters -- comfortably fits
// under MAX_REF.
const REALM_ID_MAX_DIGITS = 6;
const SPAWN_ID_MAX_HEX_CHARS = 16;
const GUID_PATTERN = `Player-\\d{1,${REALM_ID_MAX_DIGITS}}-[0-9A-Fa-f]{1,${SPAWN_ID_MAX_HEX_CHARS}}`;
const FIGHT_REF = new RegExp(`^([a-z2-7]{12}):(\\d{1,6})(?::(${GUID_PATTERN}))?$`);

export function parseFightRef(ref: string): { reportId: string; fightIndex: number; guid?: string } | null {
  const match = FIGHT_REF.exec(ref.trim());
  if (match === null) return null;
  return { reportId: match[1], fightIndex: Number.parseInt(match[2], 10), guid: match[3] };
}

/**
 * Which roster row a fight ref resolves to. A named guid wins outright when the roster has
 * them, whatever their role -- a healer or tank named by the link loads as that combatant;
 * whether the sim can run that spec is the sim's own business, not this function's. Absent,
 * or not on the roster, falls back to the first dps row, unchanged from before a guid
 * existed. Both branches require `class !== undefined`: a row the report engine inferred
 * nothing about is one neither branch can build a character from.
 */
export function selectRosterRow(roster: RosterRow[], guid: string | undefined): RosterRow | undefined {
  const named =
    guid === undefined ? undefined : roster.find((row) => row.class !== undefined && row.guid === guid);
  return named ?? roster.find((row) => row.class !== undefined && row.role === 'dps');
}

const MINUTE = 60_000;
const HOUR = 60 * MINUTE;
const DAY = 24 * HOUR;

function plural(count: number, one: string): string {
  return `${count} ${count === 1 ? one : `${one}s`} ago`;
}

export function relativeTime(iso: string, now: Date = new Date()): string {
  const at = new Date(iso);
  if (Number.isNaN(at.getTime())) return 'at an unknown time';
  const elapsed = Math.max(0, now.getTime() - at.getTime());
  if (elapsed < MINUTE) return 'just now';
  if (elapsed < HOUR) return plural(Math.floor(elapsed / MINUTE), 'minute');
  if (elapsed < DAY) return plural(Math.floor(elapsed / HOUR), 'hour');
  return plural(Math.floor(elapsed / DAY), 'day');
}

/**
 * The strip's source pill. Armory and addon carry a time because freshness is the question
 * a player has about them; a build and a fight are fixed things, so naming them is enough.
 */
export function sourcePill(source: CharacterSource, now: Date = new Date()): string {
  switch (source.kind) {
    case 'armory':
      return `Armory, ${relativeTime(source.captured_at, now)}`;
    case 'blizzard':
      return `Battle.net, ${relativeTime(source.captured_at, now)}`;
    case 'addon':
      return `Addon export, ${relativeTime(source.captured_at, now)}`;
    case 'build':
      return 'Build from planner';
    case 'fight':
      return 'From this fight';
    default:
      return 'Entered by hand';
  }
}

async function reference(ctx: LoadContext) {
  return loadReference(ctx.treeVersion);
}

export async function fromStoredCharacter(
  path: CharacterPath,
  ctx: LoadContext,
  storage?: Storage,
): Promise<SourceResult> {
  let input;
  try {
    input = await fetchSimInput(path, ctx.apiBase ?? API_BASE_URL);
  } catch (error) {
    return { ok: false, message: error instanceof Error ? error.message : simCopy.characterFailed };
  }
  const characterKey = `${path.region}/${path.ruleset}/${path.slug}`;

  // An addon-sourced read carries the addon's own export string as `gear` (input.go:
  // "the source's own shape"), and that string carries everything the paste path decodes
  // -- class, race, talents and gear -- so it takes exactly the paste path, with the
  // stored name and capture time. Nothing here is guessed: the code refuses the same
  // way the paste box does.
  if (
    (input.source === 'addon' || input.source === 'blizzard') &&
    typeof input.gear === 'string' &&
    input.gear.startsWith(`${FS1_PREFIX}:`)
  ) {
    const code = input.gear;
    const decoded = decodeFS1(code);
    if (!decoded.ok) return { ok: false, message: decoded.message };
    let talents;
    let classes;
    let races;
    try {
      [talents, { classes, races }] = await Promise.all([
        loadTalents(ctx.treeVersion, decoded.build.classSlug),
        reference(ctx),
      ]);
    } catch {
      return { ok: false, message: simCopy.characterFailed };
    }
    const result = characterFromFs1(
      code,
      talents,
      classes,
      races,
      { kind: input.source, ref: characterKey, captured_at: input.captured_at },
      path.slug,
    );
    if (result.ok) recordCurrentCharacter(result.character, 'addon', code, storage);
    return result;
  }

  const classSlug = input.spec.split('-')[0];
  try {
    // No loadTalents here: nothing below reads a TalentFile, so fetching one would be
    // dead weight -- see the point_order comment below for why.
    const { classes, races } = await reference(ctx);
    const classRow = classes.find((row) => row.slug === classSlug);
    if (classRow === undefined) return { ok: false, message: simCopy.characterFailed };
    // `SimInput` carries no race, and a race is not guessable: Forever's racials are two
    // actives and two passives each, and they are real damage. The first draft substituted
    // `races[0]`, which sims an orc's Blood Fury for a troll and says nothing -- the exact
    // silent wrong value character.ts refuses on the addon path. The same refusal applies
    // here, in the same words. It makes this source unavailable until the companion
    // records a race, which is honest and is the parked item; the addon-export source
    // works today and carries the race in the FS1 string.
    const raceRow = races.find((row) => row.slug === input.race);
    if (raceRow === undefined) {
      return { ok: false, message: simCopy.unknownRace(input.race ?? 'none recorded') };
    }
    const name = path.slug;
    // `input.talents` is fight_metrics.talent_split -- points per tree as "31/0/20"
    // (input.go's own comment: "the addon export is an opaque string this repository
    // never parses, so a character who has never parsed has none"), never a per-talent
    // order. That is enough for the level shown on the strip -- the same total
    // fromLoggedFight's own `split` produces -- but not enough to say which talent was
    // picked in which order, so point_order is honestly empty: the same "the page says
    // the tree is empty" case fromLoggedFight already has for the same reason.
    const totalPoints = input.talents
      .split('/')
      .reduce((sum, part) => sum + (Number.parseInt(part, 10) || 0), 0);
    const character: SimCharacter = {
      name,
      spec: input.spec,
      class_slug: classSlug,
      race_slug: raceRow.slug,
      talent_level: talentLevel(new Array<number>(totalPoints)),
      tree_version: ctx.treeVersion,
      point_order: [],
      // `input.gear`'s shape is the source's own -- the addon's own export JSON for an
      // addon read, `{"trinkets": […]}` for a fight (input.go) -- and openapi.yaml
      // leaves it an opaque object rather than the planner's slot-to-item map. Nothing
      // in this repository decodes either shape yet, so gear starts empty rather than
      // guessed at: the same honest-empty choice as point_order, above.
      gear: {},
      // No enchant or suffix data on this source either, so gear_slots is the (empty) id
      // map converted -- the same "the two always agree on item ids" promise every other
      // source keeps.
      gear_slots: [],
      professions: [],
      bags: [],
      bank: [],
      sets: [],
      loadouts: [],
      // Buff ids straight through. The amended contract puts the spell-id mapping on the
      // API side -- "the web never maps spell ids itself" -- so `input.buffs` is already
      // the engine's vocabulary and anything it does not recognise surfaces as the
      // engine's own error rather than being dropped here.
      buffs: [...input.buffs],
      consumables: [],
      // The API's own word for where this came from -- never a guess, so the pill is
      // true today and true unchanged when Armory lands.
      source: { kind: input.source, ref: characterKey, captured_at: input.captured_at },
    };
    // The pointer's own source is always `'armory'` here, whatever `input.source` says (the
    // character's own `source.kind` above stays the API's word, unchanged): `'armory'` in a
    // pointer means "the site's stored character, by key" -- the same `?source=armory&ref=
    // <key>` URL LandingState.svelte already writes -- and `bootstrapSource` (store.svelte.ts)
    // resolves that key straight back through this same function. Stamping the API's literal
    // `input.source` ("fight", say) here would write a pointer no loader can resume: nothing
    // reads `?source=fight&ref=<region>/<ruleset>/<slug>` as a character key.
    recordCurrentCharacter(character, 'armory', characterKey, storage);
    return { ok: true, character };
  } catch {
    return { ok: false, message: simCopy.characterFailed };
  }
}

/**
 * The shared body of `fromAddonExport` and `fromManualCode`: an FS1 string, decoded through
 * the identical `characterFromFs1` grammar, differing only in the `CharacterSource.kind` it
 * stamps on the character and the pointer source it records -- `'addon'`/`'addon'` for a
 * pasted or pushed export, `'manual'`/`'code'` for a "Sim this build" link's own unsaved
 * code. Kept as one private helper (rather than two near-duplicate exported functions) so
 * a change to the decode step -- the talent lookup, the error message, the pointer write --
 * only has one place to make it.
 */
async function loadFs1Character(
  code: string,
  ctx: LoadContext,
  kind: 'addon' | 'manual',
  pointerSource: 'addon' | 'code',
  storage?: Storage,
): Promise<SourceResult> {
  // The class is in the string, so the talent file is chosen from it rather than guessed.
  const classSlug = code.trim().split(':')[2] ?? '';
  let talents;
  let classes;
  let races;
  try {
    [talents, { classes, races }] = await Promise.all([
      loadTalents(ctx.treeVersion, classSlug),
      reference(ctx),
    ]);
  } catch {
    return { ok: false, message: simCopy.characterFailed };
  }
  const result = characterFromFs1(code, talents, classes, races, {
    kind,
    ref: '',
    captured_at: new Date().toISOString(),
  });
  if (result.ok) recordCurrentCharacter(result.character, pointerSource, code, storage);
  return result;
}

/** An FS1 string, pasted or pushed by the companion, into an `'addon'`-sourced character. */
export async function fromAddonExport(
  code: string,
  ctx: LoadContext,
  storage?: Storage,
): Promise<SourceResult> {
  return loadFs1Character(code, ctx, 'addon', 'addon', storage);
}

/**
 * An unsaved planner build's own FS1 code, into a `'manual'`-sourced character. Moved here
 * from store.svelte.ts (Task 3, current-character spec) so the tools island -- previously
 * blind to `?code=` entirely -- can bootstrap from one too, through the same loader the sim
 * page already used. Identical to `fromAddonExport` in every step but the source it stamps
 * on the character and the pointer: an addon export and a "Sim this build" link decode
 * through the identical `FS1:…` grammar and the identical `characterFromFs1`, and only
 * differ in where the character came from, which the strip's pill has to say honestly
 * (`sourcePill` reads `'addon'` as "Addon export, …" and anything else, `'manual'`
 * included, as "Entered by hand").
 */
export async function fromManualCode(
  code: string,
  ctx: LoadContext,
  storage?: Storage,
): Promise<SourceResult> {
  return loadFs1Character(code, ctx, 'manual', 'code', storage);
}

export async function fromPlannerBuild(
  buildId: string,
  ctx: LoadContext,
  storage?: Storage,
): Promise<SourceResult> {
  let record: BuildRecord | null;
  try {
    ({ data: record } = await requestEnvelope<BuildRecord>(
      `/v1/builds/${buildId}`,
      ctx.apiBase ?? API_BASE_URL,
      { credentials: 'omit', failureMessage: simCopy.buildNotFound },
    ));
  } catch {
    return { ok: false, message: simCopy.buildNotFound };
  }
  if (record === null) return { ok: false, message: simCopy.buildNotFound };

  try {
    const { classes, races } = await reference(ctx);
    const classRow = classes.find((row) => row.id === record.class_id);
    if (classRow === undefined) return { ok: false, message: simCopy.buildNotFound };
    const talents = await loadTalents(record.tree_version, classRow.slug);
    const character = fromBuildDraft(record, talents, classes, races, {
      name: record.title === undefined || record.title === '' ? classRow.name : record.title,
      source: { kind: 'build', ref: record.id, captured_at: record.created_at },
    });
    if (character === null) return { ok: false, message: simCopy.buildNotFound };
    recordCurrentCharacter(character, 'build', record.id, storage);
    return { ok: true, character };
  } catch {
    return { ok: false, message: simCopy.buildNotFound };
  }
}

/**
 * A logged fight. The gear comes from the fight's COMBATANT_INFO row through the report
 * lane's own gearFromCombatant, so the sim and the planner read a fight the same way.
 * Talents are only usable when the log recorded one rank per talent in tab order; when it
 * did not, the character still loads with its gear and the page says the tree is empty.
 */
export async function fromLoggedFight(
  ref: string,
  ctx: LoadContext,
  storage?: Storage,
): Promise<SourceResult> {
  const parsed = parseFightRef(ref);
  if (parsed === null) return { ok: false, message: simCopy.fightRefInvalid };

  try {
    const meta = await fetchReportMeta(parsed.reportId, ctx.apiBase ?? API_BASE_URL);
    const summary = await fetchSummary(meta.data_base_url, parsed.fightIndex);
    const roster = selectRosterRow(summary.roster, parsed.guid);
    const combatant = summary.combatants.find((row) => row.guid === roster?.guid);
    if (roster === undefined || combatant === undefined) {
      return { ok: false, message: simCopy.fightNoCombatant };
    }

    // `find` filters on `class !== undefined` but TypeScript does not narrow through it.
    const classSlug = classSlugFromName(roster.class ?? '');
    const [talents, { classes, races }] = await Promise.all([
      loadTalents(ctx.treeVersion, classSlug),
      reference(ctx),
    ]);
    // A combat log records no race: `CombatantRow` is guid, name, spec, item level, gear,
    // talents and buffs, and nothing else. So this source does not invent one either --
    // `race_slug` comes back EMPTY and the character strip turns its race line into a
    // select the player answers once (Task 11), with the run control disabled until they
    // do. That is one click, and it is the difference between a number that is right and a
    // number computed with an orc's Blood Fury on a troll.
    //
    // The FS1 round-trip below only needs *a* race slug to satisfy the decoder's grammar;
    // the character it produces is used for its point order alone, and `race_slug` is set
    // from `PENDING_RACE` afterwards.
    const anyRace = races[0].slug;

    const treeSizes = talents.trees.map((tree) => tree.talents.length);
    const treeRanks = treeRanksFromTalents(combatant.talents, treeSizes);
    const split =
      treeRanks === null ? treeSizes.map(() => 0) : treeRanks.map((tree) => tree.reduce((a, b) => a + b, 0));
    const order =
      treeRanks === null
        ? []
        : characterFromFs1(
            `FS1:${ctx.treeVersion}:${classSlug}:${anyRace}:${treeRanks
              .map(
                (tree) =>
                  tree
                    .map((rank) => rank.toString(36))
                    .join('')
                    .replace(/0+$/, '') || '0',
              )
              .join('/')}:`,
            talents,
            classes,
            races,
            { kind: 'fight', ref, captured_at: meta.created_at },
          );
    const point_order = Array.isArray(order) ? [] : order.ok ? order.character.point_order : [];

    const character: SimCharacter = {
      name: roster.name,
      spec: specForSplit(classSlug, split),
      class_slug: classSlug,
      // Empty on purpose: the log has no race and this function will not guess one.
      race_slug: PENDING_RACE,
      talent_level: talentLevel(point_order),
      tree_version: ctx.treeVersion,
      point_order,
      gear: gearFromCombatant(combatant.gear),
      // A fight's gear carries no enchant or suffix either, so gear_slots is the id map
      // converted -- the same "the two always agree on item ids" promise every other
      // source keeps.
      gear_slots: gearSlots(gearFromCombatant(combatant.gear)),
      professions: [],
      bags: [],
      bank: [],
      sets: [],
      loadouts: [],
      // Same as above: the fight records spell ids and CharacterSpec wants buff ids, so
      // the settings bar's preset stands in until one vocabulary maps onto the other.
      buffs: [],
      consumables: [],
      source: { kind: 'fight', ref, captured_at: meta.created_at },
    };
    recordCurrentCharacter(character, 'fight', ref, storage);
    return { ok: true, character };
  } catch {
    return { ok: false, message: simCopy.fightNoCombatant };
  }
}
