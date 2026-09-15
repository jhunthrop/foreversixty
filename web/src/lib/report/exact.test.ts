import { describe, expect, it } from 'vitest';
import { exactSplitSql, exactTableSql, rowsSql } from './exact';

const pets = { pets: new Map([['Pet-7', 'Player-1']]) };

describe('rowsSql', () => {
  it('lands a pet’s damage on its owner, by the report’s units before the advanced field', () => {
    const sql = rowsSql('damage-done', { startMs: 5000, endMs: 25_000 }, pets);
    expect(sql).toContain(
      `CASE source_guid WHEN 'Pet-7' THEN 'Player-1' ELSE coalesce(nullif(nullif(adv_owner_guid, ''), '0000000000000000'), source_guid) END AS actor`,
    );
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

  it('counts overkill only when asked', () => {
    expect(rowsSql('damage-done', { startMs: 0, endMs: 1000 }, { countOverkill: true })).toContain(
      'amount AS effective',
    );
  });

  it('reads damage taken from the victim and never folds pets', () => {
    const sql = rowsSql('damage-taken', { startMs: 0, endMs: 1000 }, pets);
    expect(sql).toContain('SELECT dest_guid AS actor, source_guid AS other_guid');
    expect(sql).not.toContain('CASE');
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

describe('exactSplitSql', () => {
  it('scopes one actor’s split, pets included, and narrows by target', () => {
    const sql = exactSplitSql(
      'damage-done',
      'Player-1',
      { startMs: 5000, endMs: 25_000 },
      'Creature-9',
      pets,
    );
    expect(sql.abilities).toContain("actor = 'Player-1' AND other_guid = 'Creature-9'");
    expect(sql.abilities).toContain("WHEN 'Pet-7' THEN 'Player-1'");
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
