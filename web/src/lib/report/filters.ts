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

/** The same rule over several encounter names at once: the whole night's bosses. */
export function bossGuidsOf(units: Unit[], fightNames: readonly string[]): Set<string> {
  const names = new Set(fightNames.filter((name) => name !== ''));
  return new Set(
    units.filter((unit) => unit.kind !== 'player' && names.has(unit.name)).map((unit) => unit.guid),
  );
}

export function playerGuids(units: Unit[]): Set<string> {
  return new Set(units.filter((unit) => unit.kind === 'player').map((unit) => unit.guid));
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
  return [...totals.values()].sort((a, b) => b.total - a.total || a.name.localeCompare(b.name));
}

export function targetOptions(actors: Actor[]): FilterOption[] {
  const totals = new Map<string, FilterOption>();
  for (const actor of actors) {
    for (const target of actor.targets) {
      const found = totals.get(target.guid);
      if (found === undefined)
        totals.set(target.guid, { id: target.guid, name: target.name, total: target.total });
      else found.total += target.total;
    }
  }
  return [...totals.values()].sort((a, b) => b.total - a.total || a.name.localeCompare(b.name));
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
  return { ...actor, abilities, targets, total: gross, effective };
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

  return actors
    .filter((actor) => !filters.playersOnly || context.players.has(actor.guid))
    .map((actor) => {
      let targets = actor.targets;
      if (filters.target !== '') targets = targets.filter((target) => target.guid === filters.target);
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
      }));

      let series = actor.series;
      const diedAt = deathBucket.get(actor.guid);
      if (filters.ignoreAfterDeath && diedAt !== undefined) series = series.slice(0, diedAt);

      return { ...rebuild(actor, scaled, targets, filters.countOverkill), series };
    })
    .filter((actor) => actor.total > 0 || actor.effective > 0)
    .sort((a, b) => b.effective - a.effective || a.guid.localeCompare(b.guid));
}
