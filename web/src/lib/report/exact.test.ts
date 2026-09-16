import { describe, expect, it } from 'vitest';
import {
  castCountsSql,
  castKey,
  exactMissesSql,
  exactSplitSql,
  exactTableSql,
  measureCasts,
  rowsSql,
  eventStreamSql,
} from './exact';
import type { QueryLayer } from './query';

const pets = { pets: new Map([['Pet-7', 'Player-1']]) };

describe('rowsSql', () => {
  it('lands a pet’s damage on its owner, by the report’s units', () => {
    const sql = rowsSql('damage-done', { startMs: 5000, endMs: 25_000 }, pets);
    expect(sql).toContain(`CASE source_guid WHEN 'Pet-7' THEN 'Player-1' ELSE source_guid END AS actor`);
    expect(sql).toContain('>= 5000 AND');
    expect(sql).toContain('< 25000');
    expect(sql).toContain('amount - greatest(coalesce(overkill, 0), 0) AS effective');
  });

  it('counts an absorb as healing by the shield’s caster under the shield’s spell', () => {
    const sql = rowsSql('healing', { startMs: 0, endMs: 1000 });
    expect(sql).toContain("kind = 'heal'");
    expect(sql).toContain('UNION ALL');
    expect(sql).toContain("kind = 'absorbed' AND extra_guid <> ''");
    expect(sql).toContain('extra_spell_id AS spell_id');
    expect(sql).toContain('amount - coalesce(overheal, 0) AS effective');
  });

  it('leaves friendly fire out of damage done, on both sides of the flags', () => {
    const sql = rowsSql('damage-done', { startMs: 0, endMs: 1000 });
    expect(sql).toContain('(source_flags & 80) = (dest_flags & 80)');
    expect(rowsSql('damage-taken', { startMs: 0, endMs: 1000 })).not.toContain('dest_flags');
  });

  it('reads one ability alone when the ability filter is set', () => {
    expect(rowsSql('damage-taken', { startMs: 0, endMs: 1000 }, { ability: 331415 })).toContain(
      'AND spell_id = 331415',
    );
    expect(rowsSql('damage-done', { startMs: 0, endMs: 1000 }, { ability: 0 })).toContain('AND spell_id = 0');
    expect(rowsSql('healing', { startMs: 0, endMs: 1000 }, { ability: 116 })).toContain('AND spell_id = 116');
  });

  it('counts overkill only when asked', () => {
    expect(rowsSql('damage-done', { startMs: 0, endMs: 1000 }, { countOverkill: true })).toContain(
      'amount AS effective',
    );
  });

  it('reads damage taken from the victim and never folds pets', () => {
    const sql = rowsSql('damage-taken', { startMs: 0, endMs: 1000 }, pets);
    expect(sql).toContain('SELECT dest_guid AS actor, source_guid AS other_guid');
    // No owner expression on the victim's side: a pet's damage taken is the pet's.
    expect(sql).not.toContain('CASE source_guid');
    expect(sql).not.toContain('CASE dest_guid');
  });
});

describe('exactTableSql', () => {
  it('groups every actor by what they hit, inside the window and target scope', () => {
    const sql = exactTableSql(
      'damage-done',
      { startMs: 44_000, endMs: 64_000 },
      { guids: ['Creature-1'], names: ['Ravenous Dreadbat'] },
    );
    expect(sql).toContain(`other_guid IN ('Creature-1') OR other_name IN ('Ravenous Dreadbat')`);
    expect(sql).toContain('GROUP BY actor, other_guid');
  });

  it('narrows damage taken by source without a name clause when none was given', () => {
    const sql = exactTableSql('damage-taken', { startMs: 0, endMs: 1000 }, { guids: ['Boss-1'], names: [] });
    expect(sql).toContain(`other_guid IN ('Boss-1')`);
    expect(sql).not.toContain('other_name IN');
  });
});

describe('dead spans', () => {
  it('leave an actor’s lines out while they were dead, in the split and in the table', () => {
    const exclude = [{ guid: 'Player-1', startMs: 63_800, endMs: 457_000 }];
    const split = exactSplitSql('damage-taken', 'Player-1', { startMs: 0, endMs: 724_000 }, null, {
      exclude,
    });
    expect(split.abilities).toContain(
      "AND NOT (actor = 'Player-1' AND fight_ms > 63800 AND fight_ms < 457000)",
    );
    expect(split.misses).toContain("AND NOT (dest_guid = 'Player-1' AND");
    const table = exactTableSql('damage-taken', { startMs: 0, endMs: 724_000 }, null, { exclude });
    expect(table).toContain("AND NOT (actor = 'Player-1' AND fight_ms > 63800 AND fight_ms < 457000)");
  });
});

describe('exactMissesSql', () => {
  it('counts avoided hits by type on the table’s own side of the events', () => {
    const sql = exactMissesSql(
      'damage-taken',
      { startMs: 0, endMs: 1000 },
      { guids: ['Boss-1'], names: ['Kaal'] },
    );
    expect(sql).toContain('SELECT dest_guid AS guid, miss_type, count(*) AS n');
    expect(sql).toContain("kind = 'missed'");
    expect(sql).toContain(`source_guid IN ('Boss-1') OR source_name IN ('Kaal')`);
  });
});

describe('exactSplitSql', () => {
  it('scopes one actor’s split, pets included, and narrows by target', () => {
    const sql = exactSplitSql(
      'damage-done',
      'Player-1',
      { startMs: 5000, endMs: 25_000 },
      { guids: ['Creature-9'], names: ['General Kaal'] },
      pets,
    );
    expect(sql.abilities).toContain(
      `actor = 'Player-1' AND (other_guid IN ('Creature-9') OR other_name IN ('General Kaal'))`,
    );
    expect(sql.abilities).toContain("WHEN 'Pet-7' THEN 'Player-1'");
    expect(sql.misses).toContain(`dest_guid IN ('Creature-9') OR dest_name IN ('General Kaal')`);
    expect(sql.targets).toContain('other_guid AS guid');
  });

  it('reads misses from the actor’s own lines', () => {
    const sql = exactSplitSql('damage-taken', 'Player-1', { startMs: 0, endMs: 1000 });
    expect(sql.misses).toContain("kind = 'missed' AND miss_type <> '' AND dest_guid = 'Player-1'");
  });

  it('quotes a name with an apostrophe', () => {
    const sql = exactSplitSql('healing', "Player-O'Neil", { startMs: 0, endMs: 1000 });
    expect(sql.abilities).toContain("actor = 'Player-O''Neil'");
  });
});

describe('measureCasts', () => {
  it('reads each caster’s starts, successes and failures per spell inside the window', async () => {
    const sql = castCountsSql({ startMs: 120_000, endMs: 180_000 });
    expect(sql).toContain("kind IN ('cast_start', 'cast_success', 'cast_failed')");
    expect(sql).toContain('>= 120000 AND');
    const layer = {
      run: async () => ({
        columns: ['guid', 'spell_id', 'kind', 'reason', 'n'],
        rows: [
          ['Player-1', 116n, 'cast_start', '', 4n],
          ['Player-1', 116n, 'cast_success', '', 3n],
          ['Player-1', 116n, 'cast_failed', 'Interrupted', 1n],
          ['Player-1', 116n, 'cast_failed', 'Not enough mana', 2n],
          ['Player-2', 8092n, 'cast_success', '', 5n],
        ],
      }),
    } as unknown as QueryLayer;
    const counts = await measureCasts(layer, 'events.parquet', { startMs: 120_000, endMs: 180_000 });
    expect(counts.get(castKey({ guid: 'Player-1', spell_id: 116 }))).toEqual({
      started: 4,
      succeeded: 3,
      failed: 3,
      fail_reasons: { Interrupted: 1, 'Not enough mana': 2 },
    });
    expect(counts.get(castKey({ guid: 'Player-2', spell_id: 8092 }))).toEqual({
      started: 0,
      succeeded: 5,
      failed: 0,
      fail_reasons: {},
    });
  });
});

describe('the ability filter on healing', () => {
  it('reads heals by their spell and absorbs by the shield they came from', () => {
    const sql = rowsSql('healing', { startMs: 0, endMs: 1000 }, { ability: 116670 });
    expect(sql).toContain("kind = 'heal' AND");
    expect(sql).toMatch(/kind = 'heal' AND [^\n]*AND spell_id = 116670/);
    expect(sql).toMatch(/kind = 'absorbed' AND[^\n]*AND extra_spell_id = 116670/);
    expect(sql).not.toMatch(/kind = 'absorbed' AND[^\n]*AND spell_id = 116670/);
  });
});

describe('damage lines', () => {
  it('reads a melee swing as it landed, on every damage table and the event stream', () => {
    // A SWING_DAMAGE with a landed twin is left out and the twin stands in as a damage
    // line; a swing with no twin stays. The stream reads the same lines.
    for (const kind of ['damage-done', 'damage-taken'] as const) {
      const sql = rowsSql(kind, { startMs: 0, endMs: 1000 });
      expect(sql).toContain("l.kind = 'damage_landed'");
      expect(sql).toContain("REPLACE ('damage' AS kind, 'SWING_DAMAGE' AS event,");
      // A landed line at 0 with no absorb of its own is soaked in full: what was thrown.
      expect(sql).toContain('coalesce(l.amount, 0) = 0 AND coalesce(l.absorbed, 0) = 0');
      expect(sql).not.toContain("FROM read_parquet('events.parquet') WHERE kind = 'damage'");
    }
    expect(eventStreamSql({ startMs: 0, endMs: 1000 })).toContain("l.kind = 'damage_landed'");
  });
});

describe('eventStreamSql', () => {
  it('reads aura refreshes alongside the hits, heals and misses', () => {
    const sql = eventStreamSql({ startMs: 0, endMs: 10_000 });
    expect(sql).toContain("kind IN ('heal', 'missed', 'aura_refresh')");
  });
});
