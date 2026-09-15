// web/src/lib/report/exact.ts
// The exact per-ability and per-target split of one actor inside a window, and the exact
// totals of a whole table, measured from the fight's own events rather than prorated from
// the whole-fight summary. The summary keeps one per-second series per actor and none per
// ability, so window.ts can only scale an ability's whole-fight figure by the window's
// share of the actor's total -- which credits a healer with spells she never cast in that
// window. This asks DuckDB the real question, through the same engine the Queries view
// uses, and hands back rows in the summary's own Ability and Pair shapes.
//
// It counts what the engine counts (logs/engine/summary/damage.go): a pet's damage and
// healing land on its owner, read from the report's units since the advanced owner field
// is empty on most lines; and an absorb is healing done by the shield's caster, under the
// shield's spell, since that is the only place an absorbed amount is counted.
import { EVENTS_TABLE, FIGHT_MS, createDuckDbEngine, createQueryLayer, type QueryLayer } from './query';
import type { Ability, Pair } from './types';
import type { TimeWindow } from './window';

export type ActorKind = 'damage-done' | 'damage-taken' | 'healing';

export interface ExactSplit {
  abilities: Ability[];
  targets: Pair[];
}

/** An actor's measured totals: what they did in the window, and to whom. */
export interface ExactTotals {
  effective: number;
  targets: Pair[];
}

/** The other side a table is narrowed to: by GUID, or by name where a filter matched a name. */
export interface TargetScope {
  guids: string[];
  names: string[];
}

/** Pet GUID to owner GUID, from the report's units. */
export type PetOwners = ReadonlyMap<string, string>;

/** What the measure counts the way the tables count it. */
export interface MeasureOptions {
  pets?: PetOwners;
  /** Count a killing blow's overkill as damage, as the "Count overkill" filter does. */
  countOverkill?: boolean;
}
const NO_PETS: PetOwners = new Map();

let shared: QueryLayer | null = null;

/** One engine for the page: the Queries view and the tables' exact measure share it. */
export function sharedQueryLayer(): QueryLayer {
  shared ??= createQueryLayer({
    load: createDuckDbEngine,
    fetchBytes: async (url) => {
      const response = await fetch(url);
      if (!response.ok) throw new Error('This fight’s event file did not load.');
      return new Uint8Array(await response.arrayBuffer());
    },
  });
  return shared;
}

function quote(value: string): string {
  return `'${value.replaceAll("'", "''")}'`;
}

/** The log's "no unit" GUID, which the advanced owner field carries on a player's own lines. */
const NO_GUID = '0000000000000000';

/** The unit a column names, resolved to its owner when it is a pet. */
function ownerExpr(column: string, pets: PetOwners): string {
  const fallback = `coalesce(nullif(nullif(adv_owner_guid, ''), '${NO_GUID}'), ${column})`;
  if (pets.size === 0) return fallback;
  const arms = [...pets.entries()].map(([pet, owner]) => `WHEN ${quote(pet)} THEN ${quote(owner)}`).join(' ');
  return `CASE ${column} ${arms} ELSE ${fallback} END`;
}

/** Half-open, like the one-second buckets the tables sum: [start, end). */
function windowClause(window: TimeWindow): string {
  return `${FIGHT_MS} >= ${Math.round(window.startMs)} AND ${FIGHT_MS} < ${Math.round(window.endMs)}`;
}

/**
 * One row per counted event, in the same shape for every kind: who it counts for (`actor`),
 * the other side, the spell it is filed under, and its effective amount. Healing takes
 * heals and absorbs; damage taken reads the victim as the actor and never folds pets.
 */
export function rowsSql(kind: ActorKind, window: TimeWindow, options: MeasureOptions = {}): string {
  const at = windowClause(window);
  const pets = options.pets ?? NO_PETS;
  const damage = options.countOverkill ? 'amount' : 'amount - greatest(coalesce(overkill, 0), 0)';
  if (kind === 'damage-taken') {
    return `SELECT dest_guid AS actor, source_guid AS other_guid, source_name AS other_name,
    spell_id, spell_name, spell_school, event, amount,
    ${damage} AS effective,
    0 AS overheal, coalesce(absorbed, 0) AS absorbed, coalesce(blocked, 0) AS blocked, critical
  FROM ${EVENTS_TABLE} WHERE kind = 'damage' AND ${at}`;
  }
  if (kind === 'healing') {
    return `SELECT ${ownerExpr('source_guid', pets)} AS actor, dest_guid AS other_guid, dest_name AS other_name,
    spell_id, spell_name, spell_school, event, amount,
    amount - coalesce(overheal, 0) AS effective,
    coalesce(overheal, 0) AS overheal, 0 AS absorbed, 0 AS blocked, critical
  FROM ${EVENTS_TABLE} WHERE kind = 'heal' AND ${at}
  UNION ALL
  SELECT ${ownerExpr('extra_guid', pets)} AS actor, dest_guid AS other_guid, dest_name AS other_name,
    extra_spell_id AS spell_id, extra_spell_name AS spell_name, extra_spell_school AS spell_school, event, amount,
    amount AS effective, 0 AS overheal, amount AS absorbed, 0 AS blocked, false AS critical
  FROM ${EVENTS_TABLE} WHERE kind = 'absorbed' AND extra_guid <> '' AND ${at}`;
  }
  return `SELECT ${ownerExpr('source_guid', pets)} AS actor, dest_guid AS other_guid, dest_name AS other_name,
    spell_id, spell_name, spell_school, event, amount,
    ${damage} AS effective,
    0 AS overheal, coalesce(absorbed, 0) AS absorbed, coalesce(blocked, 0) AS blocked, critical
  FROM ${EVENTS_TABLE} WHERE kind = 'damage' AND ${at}`;
}

function scopeClause(scope: TargetScope | null): string {
  if (scope === null || (scope.guids.length === 0 && scope.names.length === 0)) return '';
  const parts = [
    ...(scope.guids.length > 0 ? [`other_guid IN (${scope.guids.map(quote).join(', ')})`] : []),
    ...(scope.names.length > 0 ? [`other_name IN (${scope.names.map(quote).join(', ')})`] : []),
  ];
  return ` AND (${parts.join(' OR ')})`;
}

/** The three statements one actor's split needs: abilities, misses by type, targets. */
export function exactSplitSql(
  kind: ActorKind,
  guid: string,
  window: TimeWindow,
  target: string | null = null,
  options: MeasureOptions = {},
): { abilities: string; misses: string; targets: string } {
  const rows = rowsSql(kind, window, options);
  const pets = options.pets ?? NO_PETS;
  const scope = `actor = ${quote(guid)}${target === null ? '' : ` AND other_guid = ${quote(target)}`}`;
  const own =
    kind === 'damage-taken'
      ? `dest_guid = ${quote(guid)}`
      : `${ownerExpr('source_guid', pets)} = ${quote(guid)}`;
  return {
    abilities: `WITH rows AS (${rows})
SELECT spell_id, any_value(spell_name) AS spell_name, min(spell_school) AS school,
  sum(amount) AS total, sum(effective) AS effective,
  sum(overheal) AS overheal, sum(absorbed) AS absorbed, sum(blocked) AS blocked,
  count(*) FILTER (WHERE event NOT LIKE '%PERIODIC%') AS hits,
  count(*) FILTER (WHERE event LIKE '%PERIODIC%') AS ticks,
  count(*) FILTER (WHERE critical) AS crits,
  min(amount) AS min_hit, max(amount) AS max_hit
FROM rows
WHERE ${scope}
GROUP BY spell_id
ORDER BY effective DESC`,
    misses: `SELECT spell_id, miss_type, count(*) AS n
FROM ${EVENTS_TABLE}
WHERE kind = 'missed' AND miss_type <> '' AND ${own} AND ${windowClause(window)}
GROUP BY spell_id, miss_type`,
    targets: `WITH rows AS (${rows})
SELECT other_guid AS guid, any_value(other_name) AS name, sum(effective) AS total
FROM rows
WHERE ${scope}
GROUP BY other_guid
ORDER BY total DESC`,
  };
}

/**
 * Every actor's effective total per other unit inside the window, narrowed to the targets
 * (or sources) a filter chose: one statement for the whole table, so a windowed "Boss
 * damage only" or target filter reads from events instead of prorating.
 */
export function exactTableSql(
  kind: ActorKind,
  window: TimeWindow,
  scope: TargetScope | null = null,
  options: MeasureOptions = {},
): string {
  return `WITH rows AS (${rowsSql(kind, window, options)})
SELECT actor AS guid, other_guid, any_value(other_name) AS other_name, sum(effective) AS total
FROM rows
WHERE actor <> ''${scopeClause(scope)}
GROUP BY actor, other_guid
ORDER BY total DESC`;
}

type Row = Record<string, unknown>;

function rowsOf(result: { columns: string[]; rows: unknown[][] }): Row[] {
  return result.rows.map((row) => Object.fromEntries(result.columns.map((column, i) => [column, row[i]])));
}

const num = (value: unknown): number => (typeof value === 'bigint' ? Number(value) : Number(value ?? 0));

/** Measures every actor's totals and pairs inside the window, keyed by actor GUID. */
export async function measureTable(
  layer: QueryLayer,
  eventsUrl: string,
  kind: ActorKind,
  window: TimeWindow,
  scope: TargetScope | null = null,
  options: MeasureOptions = {},
): Promise<Map<string, ExactTotals>> {
  const result = await layer.run(eventsUrl, exactTableSql(kind, window, scope, options));
  const out = new Map<string, ExactTotals>();
  for (const row of rowsOf(result)) {
    const guid = String(row.guid ?? '');
    const found = out.get(guid) ?? { effective: 0, targets: [] };
    const total = num(row.total);
    found.effective += total;
    found.targets.push({ guid: String(row.other_guid ?? ''), name: String(row.other_name ?? ''), total });
    out.set(guid, found);
  }
  return out;
}

/** Measures one actor's split inside the window. Rejects with the engine's own message. */
export async function measureExact(
  layer: QueryLayer,
  eventsUrl: string,
  kind: ActorKind,
  guid: string,
  window: TimeWindow,
  target: string | null = null,
  options: MeasureOptions = {},
): Promise<ExactSplit> {
  const sql = exactSplitSql(kind, guid, window, target, options);
  const [abilities, misses, targets] = await Promise.all([
    layer.run(eventsUrl, sql.abilities),
    layer.run(eventsUrl, sql.misses),
    layer.run(eventsUrl, sql.targets),
  ]);
  const missesBySpell = new Map<number, Record<string, number>>();
  for (const row of rowsOf(misses)) {
    const spell = num(row.spell_id);
    const found = missesBySpell.get(spell) ?? {};
    found[String(row.miss_type)] = num(row.n);
    missesBySpell.set(spell, found);
  }
  return {
    abilities: rowsOf(abilities).map((row) => {
      const spell = num(row.spell_id);
      return {
        spell_id: spell,
        name: spell === 0 ? 'Melee' : String(row.spell_name ?? ''),
        school: num(row.school) || undefined,
        total: num(row.total),
        effective: num(row.effective),
        overheal: kind === 'healing' ? num(row.overheal) : undefined,
        absorbed: num(row.absorbed) || undefined,
        blocked: num(row.blocked) || undefined,
        hits: num(row.hits),
        crits: num(row.crits),
        ticks: num(row.ticks),
        misses: missesBySpell.get(spell),
        min: num(row.min_hit),
        max: num(row.max_hit),
      };
    }),
    targets: rowsOf(targets).map((row) => ({
      guid: String(row.guid ?? ''),
      name: String(row.name ?? ''),
      total: num(row.total),
    })),
  };
}
