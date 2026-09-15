// web/src/lib/report/exact.ts
// The exact per-ability and per-target split of one actor inside a window, measured from
// the fight's own events rather than prorated from the whole-fight summary. The summary
// keeps one per-second series per actor and none per ability, so window.ts can only
// scale an ability's whole-fight figure by the window's share of the actor's total --
// which credits a healer with spells she never cast in that window. This asks DuckDB the
// real question, through the same engine the Queries view uses, and hands back rows in
// the summary's own Ability and Pair shapes so the table renders them unchanged.
import { EVENTS_TABLE, FIGHT_MS, createDuckDbEngine, createQueryLayer, type QueryLayer } from './query';
import type { Ability, Pair } from './types';
import type { TimeWindow } from './window';

export type ActorKind = 'damage-done' | 'damage-taken' | 'healing';

export interface ExactSplit {
  abilities: Ability[];
  targets: Pair[];
}

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

/** The rows that are this actor's, for the table in question. Pets count for their owner. */
function actorClause(kind: ActorKind, guid: string): string {
  const g = quote(guid);
  if (kind === 'damage-taken') return `kind = 'damage' AND dest_guid = ${g}`;
  if (kind === 'healing') return `kind = 'heal' AND (source_guid = ${g} OR adv_owner_guid = ${g})`;
  return `kind = 'damage' AND (source_guid = ${g} OR adv_owner_guid = ${g})`;
}

/** The other side of each event: what was hit, or who did the hitting. */
function otherSide(kind: ActorKind): { guid: string; name: string } {
  return kind === 'damage-taken'
    ? { guid: 'source_guid', name: 'source_name' }
    : { guid: 'dest_guid', name: 'dest_name' };
}

function windowClause(window: TimeWindow): string {
  return `${FIGHT_MS} BETWEEN ${Math.round(window.startMs)} AND ${Math.round(window.endMs)}`;
}

/** The three statements the split needs: abilities, misses by type, targets. */
export function exactSplitSql(
  kind: ActorKind,
  guid: string,
  window: TimeWindow,
  target: string | null = null,
): { abilities: string; misses: string; targets: string } {
  const other = otherSide(kind);
  const scope = `${actorClause(kind, guid)} AND ${windowClause(window)}${
    target === null ? '' : ` AND ${other.guid} = ${quote(target)}`
  }`;
  const effective =
    kind === 'healing'
      ? 'sum(amount - coalesce(overheal, 0))'
      : 'sum(amount - greatest(coalesce(overkill, 0), 0))';
  return {
    abilities: `SELECT spell_id, any_value(spell_name) AS spell_name, min(spell_school) AS school,
  sum(amount) AS total, ${effective} AS effective,
  sum(coalesce(overheal, 0)) AS overheal, sum(coalesce(absorbed, 0)) AS absorbed, sum(coalesce(blocked, 0)) AS blocked,
  count(*) FILTER (WHERE event NOT LIKE '%PERIODIC%') AS hits,
  count(*) FILTER (WHERE event LIKE '%PERIODIC%') AS ticks,
  count(*) FILTER (WHERE critical) AS crits,
  min(amount) AS min_hit, max(amount) AS max_hit
FROM ${EVENTS_TABLE}
WHERE ${scope}
GROUP BY spell_id
ORDER BY effective DESC`,
    misses: `SELECT spell_id, miss_type, count(*) AS n
FROM ${EVENTS_TABLE}
WHERE kind = 'missed' AND miss_type <> '' AND ${
      kind === 'damage-taken'
        ? `dest_guid = ${quote(guid)}`
        : `(source_guid = ${quote(guid)} OR adv_owner_guid = ${quote(guid)})`
    } AND ${windowClause(window)}
GROUP BY spell_id, miss_type`,
    targets: `SELECT ${other.guid} AS guid, any_value(${other.name}) AS name, ${effective} AS total
FROM ${EVENTS_TABLE}
WHERE ${scope}
GROUP BY ${other.guid}
ORDER BY total DESC`,
  };
}

type Row = Record<string, unknown>;

function rowsOf(result: { columns: string[]; rows: unknown[][] }): Row[] {
  return result.rows.map((row) => Object.fromEntries(result.columns.map((column, i) => [column, row[i]])));
}

const num = (value: unknown): number => (typeof value === 'bigint' ? Number(value) : Number(value ?? 0));

/** Measures one actor's split inside the window. Rejects with the engine's own message. */
export async function measureExact(
  layer: QueryLayer,
  eventsUrl: string,
  kind: ActorKind,
  guid: string,
  window: TimeWindow,
  target: string | null = null,
): Promise<ExactSplit> {
  const sql = exactSplitSql(kind, guid, window, target);
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
