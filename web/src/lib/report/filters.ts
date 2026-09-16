// web/src/lib/report/filters.ts
// The filters beside the chart. Like the window, they run over the summary in the browser:
// every one of them is answerable from the per-ability and per-target totals the engine
// already wrote, so none of them costs a request.
//
// Two of the spec's filters are not controls here, and the filter bar says why:
//   pets     the engine folds a pet's damage into its owner, so there are no pet rows
//   weighted Forever publishes no per-encounter weighting table, so the control is
//            "Boss damage only" and is labelled as what it does
import { BUCKET_MS } from './window';
import type { Actor, Death, Unit } from './types';

export interface ReportFilters {
  /** A target GUID, or '' for every target. */
  target: string;
  /** A spell id, or null for every ability. */
  ability: number | null;
  /** Keep only damage aimed at the fight's encounter boss. */
  bossOnly: boolean;
  /** Drop rows for units that are not players. */
  playersOnly: boolean;
  /** Add each ability's overkill into its total. */
  countOverkill: boolean;
  /** Clip each actor's series at the second they died. */
  ignoreAfterDeath: boolean;
}

export const DEFAULT_FILTERS: ReportFilters = {
  target: '',
  ability: null,
  bossOnly: false,
  playersOnly: false,
  countOverkill: false,
  ignoreAfterDeath: false,
};

export interface FilterContext {
  bosses: Set<string>;
  players: Set<string>;
  deaths: Death[];
}

/**
 * The encounter boss is the unit whose name is the fight's name. The engine takes both
 * from ENCOUNTER_START and the creature's own name, so they match for every encounter and
 * for no trash fight -- which is right: "boss damage only" means nothing on trash.
 */
export function bossGuids(units: Unit[], fightName: string): Set<string> {
  return bossGuidsOf(units, [fightName]);
}

/**
 * The same rule over several encounter names at once: the whole night's bosses. An
 * encounter's name is not always the boss unit's name to the letter -- "Halkias, the
 * Sin-Stained Goliath" is the unit "Halkias", "Amarth, The Harvester" is "Amarth", and
 * the encounter "Stichflesh" is the unit "Surgeon Stitchflesh" -- so a unit matches when
 * one of its words is one of the encounter's, allowing a letter's slip in a long word.
 */
export function bossGuidsOf(units: Unit[], fightNames: readonly string[]): Set<string> {
  const wanted = fightNames.filter((name) => name !== '').map((name) => ({ name, words: nameWords(name) }));
  return new Set(
    units
      .filter(
        (unit) =>
          unit.kind !== 'player' &&
          wanted.some(
            ({ name, words }) =>
              unit.name === name ||
              nameWords(unit.name).some((word) => words.some((other) => sameWord(word, other))),
          ),
      )
      .map((unit) => unit.guid),
  );
}

/** The words of a unit or encounter name worth matching: four letters or more, lowercased. */
function nameWords(name: string): string[] {
  return name
    .toLowerCase()
    .split(/[^a-z']+/)
    .filter((word) => word.length >= 4 && !STOP_WORDS.has(word));
}

const STOP_WORDS = new Set(['the', 'lord', 'high', 'grand', 'surgeon', 'general', 'commander', 'executor']);

/** Equal, or one letter off in a word long enough for that to be a typo rather than a different word. */
function sameWord(a: string, b: string): boolean {
  if (a === b) return true;
  if (a.length < 6 || b.length < 6 || Math.abs(a.length - b.length) > 1) return false;
  let i = 0;
  let j = 0;
  let edits = 0;
  while (i < a.length && j < b.length) {
    if (a[i] === b[j]) {
      i += 1;
      j += 1;
      continue;
    }
    edits += 1;
    if (edits > 1) return false;
    if (a.length > b.length) i += 1;
    else if (b.length > a.length) j += 1;
    else {
      i += 1;
      j += 1;
    }
  }
  return edits + (a.length - i) + (b.length - j) <= 1;
}

/** The players, and only the players: the rows a friendlies table is made of. */
export function playerGuids(units: Unit[]): Set<string> {
  return new Set(units.filter((unit) => unit.kind === 'player').map((unit) => unit.guid));
}

/**
 * The players and everything they own: a pet or a totem is on the players' side, so an
 * "enemies" scope must not list the raid's own totems among the trash -- but neither is
 * a totem a row of the friendlies table, which is why this is not playerGuids.
 */
export function friendlyGuids(units: Unit[]): Set<string> {
  const players = playerGuids(units);
  return new Set(
    units
      .filter(
        (unit) => players.has(unit.guid) || (unit.owner_guid !== undefined && players.has(unit.owner_guid)),
      )
      .map((unit) => unit.guid),
  );
}

export interface FilterOption {
  id: string;
  name: string;
  total: number;
}

/** Every ability present in the rows, largest first, so the list opens on what matters. */
export function abilityOptions(actors: Actor[]): FilterOption[] {
  const totals = new Map<number, FilterOption>();
  for (const actor of actors) {
    for (const ability of actor.abilities) {
      const found = totals.get(ability.spell_id);
      if (found === undefined) {
        totals.set(ability.spell_id, {
          id: String(ability.spell_id),
          name: ability.name,
          total: ability.total,
        });
      } else {
        found.total += ability.total;
      }
    }
  }
  const options = [...totals.values()].sort((a, b) => b.total - a.total || a.name.localeCompare(b.name));
  // Two spells with one name (a cast and its heal, say) are told apart by their id.
  const names = new Map<string, number>();
  for (const option of options) names.set(option.name, (names.get(option.name) ?? 0) + 1);
  return options.map((option) =>
    (names.get(option.name) ?? 0) > 1
      ? { ...option, name: option.id === '0' ? `${option.name} swing` : `${option.name} #${option.id}` }
      : option,
  );
}

/**
 * One option per unit name: a trash pack is thirteen "Mindbender" GUIDs, and thirteen
 * options of the same word help nobody. The option's id is the first GUID seen, and
 * applyActorFilters matches every unit of that name (see targetsNamed).
 */
export function targetOptions(actors: Actor[]): FilterOption[] {
  const totals = new Map<string, FilterOption & { guids: Set<string> }>();
  for (const actor of actors) {
    for (const target of actor.targets) {
      const found = totals.get(target.name);
      if (found === undefined)
        totals.set(target.name, {
          id: target.guid,
          name: target.name,
          total: target.total,
          guids: new Set([target.guid]),
        });
      else {
        found.total += target.total;
        found.guids.add(target.guid);
      }
    }
  }
  return [...totals.values()]
    .sort((a, b) => b.total - a.total || a.name.localeCompare(b.name))
    .map(({ guids, ...option }) =>
      guids.size > 1 ? { ...option, name: `${option.name} ×${guids.size}` } : option,
    );
}

/** The name the target filter's GUID stands for, so every unit of that name matches. */
function targetName(actors: Actor[], guid: string): string | null {
  for (const actor of actors) {
    const found = actor.targets.find((target) => target.guid === guid);
    if (found !== undefined) return found.name;
  }
  return null;
}

/** Rebuilds a row's totals from whatever abilities and targets survived the filters. */
function rebuild(
  actor: Actor,
  abilities: Actor['abilities'],
  targets: Actor['targets'],
  countOverkill: boolean,
): Actor {
  const effective = abilities.reduce(
    (total, ability) => total + ability.effective + (countOverkill ? (ability.overkill ?? 0) : 0),
    0,
  );
  const gross = abilities.reduce(
    (total, ability) => total + ability.total + (countOverkill ? (ability.overkill ?? 0) : 0),
    0,
  );
  // Overheal and absorbs are per-ability figures too, so they follow the same share.
  const overheal = abilities.some((ability) => ability.overheal !== undefined)
    ? abilities.reduce((total, ability) => total + (ability.overheal ?? 0), 0)
    : actor.overheal;
  return { ...actor, abilities, targets, total: gross, effective, overheal };
}

export function applyActorFilters(actors: Actor[], filters: ReportFilters, context: FilterContext): Actor[] {
  const untouched =
    filters.target === '' &&
    filters.ability === null &&
    !filters.bossOnly &&
    !filters.playersOnly &&
    !filters.countOverkill &&
    !filters.ignoreAfterDeath;
  if (untouched) return actors;

  const deathBucket = new Map<string, number>();
  for (const death of context.deaths) deathBucket.set(death.guid, Math.ceil(death.at_ms / BUCKET_MS));
  const wantedName = filters.target === '' ? null : targetName(actors, filters.target);

  return actors
    .filter((actor) => !filters.playersOnly || context.players.has(actor.guid))
    .map((actor) => {
      let targets = actor.targets;
      if (filters.target !== '')
        targets = targets.filter((target) => target.guid === filters.target || target.name === wantedName);
      if (filters.bossOnly) targets = targets.filter((target) => context.bosses.has(target.guid));

      let abilities = actor.abilities;
      if (filters.ability !== null)
        abilities = abilities.filter((ability) => ability.spell_id === filters.ability);

      // A target filter cannot be resolved per ability from the summary, so the row is
      // scaled to the surviving targets' share -- the same trade the window makes, and the
      // reason the Queries view exists.
      const targetShare =
        actor.targets.length === 0 || targets.length === actor.targets.length
          ? 1
          : targets.reduce((sum, target) => sum + target.total, 0) /
            Math.max(
              actor.targets.reduce((sum, target) => sum + target.total, 0),
              1,
            );

      const scaled = abilities.map((ability) => ({
        ...ability,
        total: Math.round(ability.total * targetShare),
        effective: Math.round(ability.effective * targetShare),
        overheal: ability.overheal === undefined ? undefined : Math.round(ability.overheal * targetShare),
        // Not prorated either: what a shield soaked off one target is not the damage share
        // of that target; a pull measures it, the night leaves it out under a target filter.
        absorbed: undefined,
        blocked: undefined,
        hits: Math.round(ability.hits * targetShare),
        crits: Math.round(ability.crits * targetShare),
        ticks: Math.round(ability.ticks * targetShare),
        // Not prorated: an avoided hit has no damage share to scale by, and scaling a
        // count by one gives noise (eight absorbs on a boss read as "~29 parry"). A pull
        // measures them from its events; the night leaves them out under a target filter.
        misses: undefined,
      }));

      let series = actor.series;
      const diedAt = deathBucket.get(actor.guid);
      if (filters.ignoreAfterDeath && diedAt !== undefined) series = series.slice(0, diedAt);

      return { ...rebuild(actor, scaled, targets, filters.countOverkill), series };
    })
    .filter((actor) => actor.total > 0 || actor.effective > 0)
    .sort((a, b) => b.effective - a.effective || a.guid.localeCompare(b.guid));
}
