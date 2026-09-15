import { describe, expect, it } from 'vitest';
import { exactSplitSql } from './exact';

describe('exactSplitSql', () => {
  it('scopes damage done to the actor and their pets, inside the window', () => {
    const sql = exactSplitSql('damage-done', 'Player-1', { startMs: 5000, endMs: 25_000 });
    expect(sql.abilities).toContain(
      "kind = 'damage' AND (source_guid = 'Player-1' OR adv_owner_guid = 'Player-1')",
    );
    expect(sql.abilities).toContain('BETWEEN 5000 AND 25000');
    expect(sql.targets).toContain('dest_guid AS guid');
  });

  it('reads damage taken from the other side, and a target filter narrows the source', () => {
    const sql = exactSplitSql('damage-taken', 'Player-1', { startMs: 0, endMs: 1000 }, 'Creature-9');
    expect(sql.abilities).toContain("dest_guid = 'Player-1'");
    expect(sql.abilities).toContain("source_guid = 'Creature-9'");
    expect(sql.targets).toContain('source_guid AS guid');
  });

  it('measures healing net of overhealing and quotes a name with an apostrophe', () => {
    const sql = exactSplitSql('healing', "Player-O'Neil", { startMs: 0, endMs: 1000 });
    expect(sql.abilities).toContain('sum(amount - coalesce(overheal, 0)) AS effective');
    expect(sql.abilities).toContain("'Player-O''Neil'");
    expect(sql.abilities).toContain("kind = 'heal'");
  });
});
