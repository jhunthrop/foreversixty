// web/src/lib/sim/sources.ts
// The four ways a character reaches the simulator, and the pill that says which. Every one
// of them ends in a SimCharacter, so nothing downstream -- the strip, the request builder,
// the planner link -- knows or cares where it came from.
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
//
// Each returns a result rather than throwing: every one of these is something a player
// typed or pasted, and the page shows the reason inline beside the field.
import { API_BASE_URL } from '../planner/config';
import { loadReference, loadTalents } from '../planner/load';
import type { BuildRecord } from '../planner/types';
import { fetchReportMeta, fetchSummary } from '../report/load';
import { gearFromCombatant, treeRanksFromTalents } from '../report/planner-link';
import { classSlugFromName } from '../report/tree-sizes';
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
import type { CharacterSource } from './types';

export interface LoadContext {
  /** The active data build; every /data/<build>/ fetch and every character carries it. */
  treeVersion: string;
  apiBase?: string;
}

export type SourceResult = { ok: true; character: SimCharacter } | { ok: false; message: string };

const FIGHT_REF = /^([a-z2-7]{12}):(\d{1,6})$/;

export function parseFightRef(ref: string): { reportId: string; fightIndex: number } | null {
  const match = FIGHT_REF.exec(ref.trim());
  return match === null ? null : { reportId: match[1], fightIndex: Number.parseInt(match[2], 10) };
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

export async function fromStoredCharacter(path: CharacterPath, ctx: LoadContext): Promise<SourceResult> {
  let input;
  try {
    input = await fetchSimInput(path, ctx.apiBase ?? API_BASE_URL);
  } catch (error) {
    return { ok: false, message: error instanceof Error ? error.message : simCopy.characterFailed };
  }
  const characterKey = `${path.region}/${path.ruleset}/${path.slug}`;

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
    return {
      ok: true,
      character: {
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
      },
    };
  } catch {
    return { ok: false, message: simCopy.characterFailed };
  }
}

export async function fromAddonExport(code: string, ctx: LoadContext): Promise<SourceResult> {
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
  return characterFromFs1(code, talents, classes, races, {
    kind: 'addon',
    ref: '',
    captured_at: new Date().toISOString(),
  });
}

export async function fromPlannerBuild(buildId: string, ctx: LoadContext): Promise<SourceResult> {
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
    return character === null ? { ok: false, message: simCopy.buildNotFound } : { ok: true, character };
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
export async function fromLoggedFight(ref: string, ctx: LoadContext): Promise<SourceResult> {
  const parsed = parseFightRef(ref);
  if (parsed === null) return { ok: false, message: simCopy.fightRefInvalid };

  try {
    const meta = await fetchReportMeta(parsed.reportId, ctx.apiBase ?? API_BASE_URL);
    const summary = await fetchSummary(meta.data_base_url, parsed.fightIndex);
    const roster = summary.roster.find((row) => row.class !== undefined && row.role === 'dps');
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

    return {
      ok: true,
      character: {
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
      },
    };
  } catch {
    return { ok: false, message: simCopy.fightNoCombatant };
  }
}
