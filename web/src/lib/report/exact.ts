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
import type { Ability, Mitigated, Pair } from './types';
import type { TimeWindow } from './window';

export type ActorKind = 'damage-done' | 'damage-taken' | 'healing';

export interface ExactSplit {
  abilities: Ability[];
  targets: Pair[];
}

/** An actor's measured totals: what they did in the window, gross and net, and to whom. */
export interface ExactTotals {
  effective: number;
  /** Gross amount, overheal included, so the row's overheal share stays exact. */
  total: number;
  overheal: number;
  targets: Pair[];
  mitigated: Mitigated;
  /** Whole seconds inside the actor's dead spans that still carried a line of theirs: a pet or a dot ticking. */
  deadActiveMs: number;
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
  /** Spans to leave out per actor: the time each player was dead, for "ignore events after a death". */
  exclude?: DeadSpan[];
  /**
   * The ability filter's spell id: every table's rows, splits and misses are read for
   * that spell alone, so a target and an ability together measure exactly. Healing files
   * an absorb under the shield's spell, which the raw line carries as its extra spell.
   */
  ability?: number;
}

/** One stretch a player was dead, from the death to the first cast after it (or the fight's end). */
export interface DeadSpan {
  guid: string;
  startMs: number;
  endMs: number;
}

/** The clause that leaves an actor's dead spans out, on whatever names the actor and the instant.
 *  (The instant's alias is `fight_ms`, never `at`: `at` is a keyword in DuckDB.) */
function excludeClause(exclude: DeadSpan[] | undefined, actor: string, at: string): string {
  if (exclude === undefined || exclude.length === 0) return '';
  return exclude
    .map(
      (span) =>
        // Strictly after the death's instant: the killing blow shares it and must stay counted.
        ` AND NOT (${actor} = ${quote(span.guid)} AND ${at} > ${Math.round(span.startMs)} AND ${at} < ${Math.round(span.endMs)})`,
    )
    .join('');
}
const NO_PETS: PetOwners = new Map();
/** A measure is not a preview: every row, or the totals are wrong. */
const ALL_ROWS = { maxRows: Number.POSITIVE_INFINITY };

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

/** The other side of each event: what was hit, or who did the hitting. */
function otherSide(kind: ActorKind): { guid: string; name: string } {
  return kind === 'damage-taken'
    ? { guid: 'source_guid', name: 'source_name' }
    : { guid: 'dest_guid', name: 'dest_name' };
}

/** The unit a column names, resolved to its owner when it is a pet. */
function ownerExpr(column: string, pets: PetOwners): string {
  // From the report's units only. The advanced owner field on a damage or heal line
  // describes the unit the advanced block is about, which is the target, so reading it
  // here credited a heal on someone's pet to that someone.
  if (pets.size === 0) return column;
  const arms = [...pets.entries()].map(([pet, owner]) => `WHEN ${quote(pet)} THEN ${quote(owner)}`).join(' ');
  return `CASE ${column} ${arms} ELSE ${column} END`;
}

/** Half-open, like the one-second buckets the tables sum: [start, end). */
/** The pet's name when a line is a pet's, so its abilities stay apart from its owner's; '' otherwise. */
function viaExpr(pets: PetOwners): string {
  if (pets.size === 0) return `''`;
  return `CASE WHEN source_guid IN (${[...pets.keys()].map(quote).join(', ')}) THEN source_name ELSE '' END`;
}

/** The ability filter as SQL, on the column the line files its spell under (see MeasureOptions.ability). */
function abilityClause(options: MeasureOptions, column = 'spell_id'): string {
  if (options.ability === undefined) return '';
  return ` AND ${column} = ${Math.trunc(options.ability)}`;
}

function windowClause(window: TimeWindow): string {
  return `${FIGHT_MS} >= ${Math.round(window.startMs)} AND ${FIGHT_MS} < ${Math.round(window.endMs)}`;
}

/**
 * One row per counted event, in the same shape for every kind: who it counts for (`actor`),
 * the other side, the spell it is filed under, and its effective amount. Healing takes
 * heals and absorbs; damage taken reads the victim as the actor and never folds pets.
 */
/**
 * The damage lines as the engine counts them (0.3.5): a melee swing is logged twice,
 * SWING_DAMAGE with what was thrown and SWING_DAMAGE_LANDED a few milliseconds later with
 * what the target took -- the amount after a shield soaked it, the absorb, the overkill.
 * The landed line stands in for its swing; a swing with no landed line within the echo
 * counts as thrown; a landed line with no swing counts for nothing. Every damage read
 * here goes through this, so a measured window agrees with the summary and the recap.
 */
const SWING_ECHO_NS = 250_000_000;
export const DAMAGE_LINES = `(SELECT * FROM ${EVENTS_TABLE} d
  WHERE d.kind = 'damage' AND NOT (d.event = 'SWING_DAMAGE' AND EXISTS (
    SELECT 1 FROM ${EVENTS_TABLE} l
    WHERE l.kind = 'damage_landed' AND l.source_guid = d.source_guid AND l.dest_guid = d.dest_guid
      AND l.time_unix_nano BETWEEN d.time_unix_nano AND d.time_unix_nano + ${SWING_ECHO_NS}))
  UNION ALL
  SELECT * REPLACE ('damage' AS kind, 'SWING_DAMAGE' AS event,
    CASE WHEN coalesce(l.amount, 0) = 0 AND coalesce(l.absorbed, 0) = 0 THEN (
      SELECT max(d.amount) FROM ${EVENTS_TABLE} d
      WHERE d.event = 'SWING_DAMAGE' AND d.source_guid = l.source_guid AND d.dest_guid = l.dest_guid
        AND d.time_unix_nano BETWEEN l.time_unix_nano - ${SWING_ECHO_NS} AND l.time_unix_nano)
    ELSE l.absorbed END AS absorbed) FROM ${EVENTS_TABLE} l
  WHERE l.kind = 'damage_landed' AND EXISTS (
    SELECT 1 FROM ${EVENTS_TABLE} d
    WHERE d.event = 'SWING_DAMAGE' AND d.source_guid = l.source_guid AND d.dest_guid = l.dest_guid
      AND d.time_unix_nano BETWEEN l.time_unix_nano - ${SWING_ECHO_NS} AND l.time_unix_nano))`;

export function rowsSql(kind: ActorKind, window: TimeWindow, options: MeasureOptions = {}): string {
  const at = `${windowClause(window)}${abilityClause(options)}`;
  const pets = options.pets ?? NO_PETS;
  const damage = options.countOverkill ? 'amount' : 'amount - greatest(coalesce(overkill, 0), 0)';
  if (kind === 'damage-taken') {
    return `SELECT dest_guid AS actor, source_guid AS other_guid, source_name AS other_name,
    '' AS via,
    ${FIGHT_MS} AS fight_ms, spell_id, spell_name, spell_school, event, amount,
    ${damage} AS effective,
    0 AS overheal, coalesce(absorbed, 0) AS absorbed, coalesce(blocked, 0) AS blocked, critical
  FROM ${DAMAGE_LINES} WHERE ${at}`;
  }
  if (kind === 'healing') {
    return `SELECT ${ownerExpr('source_guid', pets)} AS actor, dest_guid AS other_guid, dest_name AS other_name,
    ${viaExpr(pets)} AS via,
    ${FIGHT_MS} AS fight_ms, spell_id, spell_name, spell_school, event, amount,
    amount - coalesce(overheal, 0) AS effective,
    coalesce(overheal, 0) AS overheal, 0 AS absorbed, 0 AS blocked, critical
  FROM ${EVENTS_TABLE} WHERE kind = 'heal' AND ${at}
  UNION ALL
  SELECT ${ownerExpr('extra_guid', pets)} AS actor, dest_guid AS other_guid, dest_name AS other_name,
    '' AS via,
    ${FIGHT_MS} AS fight_ms, extra_spell_id AS spell_id, extra_spell_name AS spell_name, extra_spell_school AS spell_school, event, amount,
    amount AS effective, 0 AS overheal, amount AS absorbed, 0 AS blocked, false AS critical
  FROM ${EVENTS_TABLE} WHERE kind = 'absorbed' AND extra_guid <> '' AND ${windowClause(window)}${abilityClause(options, 'extra_spell_id')}`;
  }
  return `SELECT ${ownerExpr('source_guid', pets)} AS actor, dest_guid AS other_guid, dest_name AS other_name,
    ${viaExpr(pets)} AS via,
    ${FIGHT_MS} AS fight_ms, spell_id, spell_name, spell_school, event, amount,
    ${damage} AS effective,
    0 AS overheal, coalesce(absorbed, 0) AS absorbed, coalesce(blocked, 0) AS blocked, critical
  FROM ${DAMAGE_LINES} WHERE ${OTHER_SIDE} AND ${at}`;
}

/**
 * Damage done is damage to the other side, as the engine counts it: a hit on a friendly
 * unit (an Earthen Wall Totem taking it in a player's place) is nobody's damage done.
 * The reaction bits are friendly 0x10 and hostile 0x40; a line whose units carry neither
 * is on nobody's side and stays counted.
 */
const OTHER_SIDE = `NOT ((source_flags & 80) <> 0 AND (dest_flags & 80) <> 0 AND (source_flags & 80) = (dest_flags & 80))`;

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
  target: TargetScope | null = null,
  options: MeasureOptions = {},
): { abilities: string; misses: string; targets: string } {
  const rows = rowsSql(kind, window, options);
  const pets = options.pets ?? NO_PETS;
  const scope = `actor = ${quote(guid)}${scopeClause(target)}${excludeClause(options.exclude, 'actor', 'fight_ms')}`;
  const other = otherSide(kind);
  const own = `${
    kind === 'damage-taken'
      ? `dest_guid = ${quote(guid)}`
      : `${ownerExpr('source_guid', pets)} = ${quote(guid)}`
  }${scopeClause(target).replaceAll('other_guid', other.guid).replaceAll('other_name', other.name)}${excludeClause(
    options.exclude,
    kind === 'damage-taken' ? 'dest_guid' : ownerExpr('source_guid', pets),
    FIGHT_MS,
  )}`;
  return {
    abilities: `WITH rows AS (${rows})
SELECT spell_id, via, any_value(spell_name) AS spell_name, min(spell_school) AS school,
  sum(amount) AS total, sum(effective) AS effective,
  sum(overheal) AS overheal, sum(absorbed) AS absorbed, sum(blocked) AS blocked,
  count(*) FILTER (WHERE event NOT LIKE '%PERIODIC%') AS hits,
  count(*) FILTER (WHERE event LIKE '%PERIODIC%') AS ticks,
  count(*) FILTER (WHERE critical) AS crits,
  min(effective) AS min_hit, max(effective) AS max_hit
FROM rows
WHERE ${scope}
GROUP BY spell_id, via
ORDER BY effective DESC`,
    misses: `SELECT spell_id, any_value(spell_name) AS spell_name, any_value(spell_school) AS school,
  miss_type, count(*) AS n, coalesce(sum(amount), 0) AS amount
FROM ${EVENTS_TABLE}
WHERE kind = 'missed' AND miss_type <> '' AND ${own}${kind === 'damage-done' ? ` AND ${OTHER_SIDE}` : ''} AND ${windowClause(window)}${abilityClause(options)}
GROUP BY spell_id, miss_type`,
    targets: `WITH rows AS (${rows})
SELECT other_guid AS guid, any_value(other_name) AS name, sum(effective) AS total, sum(overheal) AS overheal
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
SELECT actor AS guid, other_guid, any_value(other_name) AS other_name, sum(effective) AS total,
  sum(amount) AS gross, sum(overheal) AS overheal, sum(absorbed) AS absorbed, sum(blocked) AS blocked
FROM rows
WHERE actor <> ''${scopeClause(scope)}${excludeClause(options.exclude, 'actor', 'fight_ms')}
GROUP BY actor, other_guid
ORDER BY total DESC`;
}

/** Every actor's avoided hits by type inside the window, on the same side as the table. */
export function exactMissesSql(
  kind: ActorKind,
  window: TimeWindow,
  scope: TargetScope | null = null,
  options: MeasureOptions = {},
): string {
  const pets = options.pets ?? NO_PETS;
  const actor = kind === 'damage-taken' ? 'dest_guid' : ownerExpr('source_guid', pets);
  const other = otherSide(kind);
  const narrowed =
    scope === null || (scope.guids.length === 0 && scope.names.length === 0)
      ? ''
      : scopeClause(scope).replaceAll('other_guid', other.guid).replaceAll('other_name', other.name);
  return `SELECT ${actor} AS guid, miss_type, count(*) AS n, coalesce(sum(amount), 0) AS amount
FROM ${EVENTS_TABLE}
WHERE kind = 'missed' AND miss_type <> '' AND ${windowClause(window)}${abilityClause(options)}${narrowed}${excludeClause(options.exclude, actor, FIGHT_MS)}
GROUP BY 1, 2`;
}

/** One caster's counts for one spell inside a window, read from the fight's cast lines. */
export interface CastCounts {
  started: number;
  succeeded: number;
  failed: number;
  fail_reasons: Record<string, number>;
}

/** The key a cast row and its measured counts share: the caster and the spell. */
export function castKey(row: { guid: string; spell_id: number }): string {
  return `${row.guid}|${row.spell_id}`;
}

/**
 * Every cast line in the window by caster, spell, kind and failure reason: the summary
 * keeps only whole-fight counts of starts and failures, and a window's cancelled casts
 * are the starts inside it that never went off.
 */
export function castCountsSql(window: TimeWindow): string {
  return `SELECT source_guid AS guid, spell_id, kind, coalesce(failed_type, '') AS reason, count(*) AS n
FROM ${EVENTS_TABLE}
WHERE kind IN ('cast_start', 'cast_success', 'cast_failed') AND ${windowClause(window)}
GROUP BY 1, 2, 3, 4`;
}

/** Measures each caster's starts, successes and failures per spell inside the window. */
export async function measureCasts(
  layer: QueryLayer,
  eventsUrl: string,
  window: TimeWindow,
): Promise<Map<string, CastCounts>> {
  const result = await layer.run(eventsUrl, castCountsSql(window), ALL_ROWS);
  const out = new Map<string, CastCounts>();
  for (const row of rowsOf(result)) {
    const key = castKey({ guid: String(row.guid ?? ''), spell_id: num(row.spell_id) });
    const found = out.get(key) ?? { started: 0, succeeded: 0, failed: 0, fail_reasons: {} };
    const n = num(row.n);
    switch (String(row.kind)) {
      case 'cast_start':
        found.started += n;
        break;
      case 'cast_success':
        found.succeeded += n;
        break;
      default: {
        found.failed += n;
        const reason = String(row.reason ?? '');
        if (reason !== '') found.fail_reasons[reason] = (found.fail_reasons[reason] ?? 0) + n;
      }
    }
    out.set(key, found);
  }
  return out;
}

/**
 * One line of the full stream: every hit, heal and avoided hit in the window, for the
 * events view. A miss is a line of its own: the parry or block that kept a swing off the
 * tank is the moment a tank goes to the events to find.
 */
export interface StreamLine {
  atMs: number;
  kind: 'damage' | 'heal' | 'missed';
  sourceGuid: string;
  sourceName: string;
  destGuid: string;
  destName: string;
  spellName: string;
  amount: number;
  overheal: number;
  absorbed: number;
  /** Damage: what a block took off the hit; a full block is a miss of type BLOCK instead. */
  blocked: number;
  /** Missed: PARRY, DODGE, BLOCK, MISS, IMMUNE, ABSORB, DEFLECT, EVADE, RESIST, REFLECT. */
  missType: string;
}

export function eventStreamSql(window: TimeWindow): string {
  return `SELECT ${FIGHT_MS} AS fight_ms, kind, source_guid, source_name, dest_guid, dest_name, spell_name,
  coalesce(amount, 0) AS amount, coalesce(overheal, 0) AS overheal, coalesce(absorbed, 0) AS absorbed,
  coalesce(blocked, 0) AS blocked, coalesce(miss_type, '') AS miss_type
FROM (SELECT * FROM ${DAMAGE_LINES} UNION ALL SELECT * FROM ${EVENTS_TABLE} WHERE kind IN ('heal', 'missed'))
WHERE ${windowClause(window)}
ORDER BY time_unix_nano, line`;
}

/** Every hit and heal in the window, time-ordered. */
export async function loadEventStream(
  layer: QueryLayer,
  eventsUrl: string,
  window: TimeWindow,
): Promise<StreamLine[]> {
  const result = await layer.run(eventsUrl, eventStreamSql(window), ALL_ROWS);
  return rowsOf(result).map((row) => ({
    atMs: num(row.fight_ms),
    kind: row.kind === 'heal' ? 'heal' : row.kind === 'missed' ? 'missed' : 'damage',
    sourceGuid: String(row.source_guid ?? ''),
    sourceName: String(row.source_name ?? ''),
    destGuid: String(row.dest_guid ?? ''),
    destName: String(row.dest_name ?? ''),
    spellName: String(row.spell_name ?? ''),
    amount: num(row.amount),
    overheal: num(row.overheal),
    absorbed: num(row.absorbed),
    blocked: num(row.blocked),
    missType: String(row.miss_type ?? ''),
  }));
}

/**
 * Per actor, the whole seconds inside their dead spans that carried a damage, heal or cast
 * line of theirs: the summary's active time counted those, and "ignore events while dead"
 * takes them back out.
 */
export function deadActiveSql(window: TimeWindow, options: MeasureOptions): string {
  const pets = options.pets ?? NO_PETS;
  const actor = ownerExpr('source_guid', pets);
  const spans = (options.exclude ?? [])
    .map(
      (span) =>
        `(${actor} = ${quote(span.guid)} AND ${FIGHT_MS} > ${Math.round(span.startMs)} AND ${FIGHT_MS} < ${Math.round(span.endMs)})`,
    )
    .join(' OR ');
  return `SELECT ${actor} AS guid, count(DISTINCT floor(${FIGHT_MS} / 1000)) AS seconds
FROM ${EVENTS_TABLE}
WHERE kind IN ('damage', 'heal', 'cast_success') AND ${windowClause(window)} AND (${spans || 'false'})
GROUP BY 1`;
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
  const empty = { columns: [], rows: [] };
  const [result, misses, deadActive] = await Promise.all([
    layer.run(eventsUrl, exactTableSql(kind, window, scope, options), ALL_ROWS),
    kind === 'healing'
      ? Promise.resolve(empty)
      : layer.run(eventsUrl, exactMissesSql(kind, window, scope, options), ALL_ROWS),
    (options.exclude ?? []).length === 0
      ? Promise.resolve(empty)
      : layer.run(eventsUrl, deadActiveSql(window, options), ALL_ROWS),
  ]);
  const out = new Map<string, ExactTotals>();
  const fresh = (): ExactTotals => ({
    effective: 0,
    total: 0,
    overheal: 0,
    targets: [],
    mitigated: { absorbed: 0, blocked: 0, misses: {} },
    deadActiveMs: 0,
  });
  for (const row of rowsOf(result)) {
    const guid = String(row.guid ?? '');
    const found = out.get(guid) ?? fresh();
    const total = num(row.total);
    found.effective += total;
    found.total += num(row.gross);
    found.overheal += num(row.overheal);
    found.mitigated.absorbed += num(row.absorbed);
    found.mitigated.blocked += num(row.blocked);
    found.targets.push({
      guid: String(row.other_guid ?? ''),
      name: String(row.other_name ?? ''),
      total,
      ...(kind === 'healing' ? { overheal: num(row.overheal) } : {}),
    });
    out.set(guid, found);
  }
  for (const row of rowsOf(deadActive)) {
    const guid = String(row.guid ?? '');
    const found = out.get(guid) ?? fresh();
    found.deadActiveMs += num(row.seconds) * 1000;
    out.set(guid, found);
  }
  for (const row of rowsOf(misses)) {
    const guid = String(row.guid ?? '');
    const found = out.get(guid) ?? fresh();
    const type = String(row.miss_type);
    found.mitigated.misses[type] = (found.mitigated.misses[type] ?? 0) + num(row.n);
    // A hit a shield ate in full is a miss of type ABSORB carrying the amount on one
    // client and a damage line landing for 0 on another; both are absorbed.
    if (type === 'ABSORB') found.mitigated.absorbed += num(row.amount);
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
  target: TargetScope | null = null,
  options: MeasureOptions = {},
): Promise<ExactSplit> {
  const sql = exactSplitSql(kind, guid, window, target, options);
  const [abilities, misses, targets] = await Promise.all([
    layer.run(eventsUrl, sql.abilities, ALL_ROWS),
    layer.run(eventsUrl, sql.misses, ALL_ROWS),
    layer.run(eventsUrl, sql.targets, ALL_ROWS),
  ]);
  const missesBySpell = new Map<number, Record<string, number>>();
  /** The absorb misses' amounts per spell: a shield's work, counted as absorbed on the row. */
  const absorbMissBySpell = new Map<number, number>();
  // An ability that never landed (every cast absorbed, dodged or parried) has no damage
  // row to hang its misses on; it keeps a row of its own, at zero, so the miss count the
  // mitigation line adds up is on the table too.
  const missedOnly = new Map<number, Ability>();
  for (const row of rowsOf(misses)) {
    const spell = num(row.spell_id);
    const found = missesBySpell.get(spell) ?? {};
    found[String(row.miss_type)] = num(row.n);
    missesBySpell.set(spell, found);
    if (String(row.miss_type) === 'ABSORB')
      absorbMissBySpell.set(spell, (absorbMissBySpell.get(spell) ?? 0) + num(row.amount));
    missedOnly.set(spell, {
      spell_id: spell,
      name: spell === 0 ? 'Melee' : String(row.spell_name ?? ''),
      school: num(row.school) || undefined,
      total: 0,
      effective: 0,
      hits: 0,
      crits: 0,
      ticks: 0,
      misses: found,
      min: 0,
      max: 0,
    });
  }
  const landedRows = rowsOf(abilities);
  for (const row of landedRows) missedOnly.delete(num(row.spell_id));
  return {
    abilities: [
      ...landedRows.map((row) => {
        const spell = num(row.spell_id);
        return {
          spell_id: spell,
          name: spell === 0 ? 'Melee' : String(row.spell_name ?? ''),
          via: String(row.via ?? '') || undefined,
          school: num(row.school) || undefined,
          total: num(row.total),
          effective: num(row.effective),
          overheal: kind === 'healing' ? num(row.overheal) : undefined,
          absorbed: num(row.absorbed) + (absorbMissBySpell.get(spell) ?? 0) || undefined,
          blocked: num(row.blocked) || undefined,
          hits: num(row.hits),
          crits: num(row.crits),
          ticks: num(row.ticks),
          misses: missesBySpell.get(spell),
          min: num(row.min_hit),
          max: num(row.max_hit),
        };
      }),
      ...missedOnly.values(),
    ],
    targets: rowsOf(targets).map((row) => ({
      guid: String(row.guid ?? ''),
      name: String(row.name ?? ''),
      total: num(row.total),
      ...(kind === 'healing' ? { overheal: num(row.overheal) } : {}),
    })),
  };
}
