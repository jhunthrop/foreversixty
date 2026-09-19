import { describe, expect, it } from 'vitest';
// A plain .mjs so scripts/sync-sim-ids.mjs can import the same parser this test drives:
// the script runs under node before vitest ever starts, so the parser cannot be a .ts.
import { parseIdsMarkdown } from '../../../scripts/ids-md.mjs';

const SAMPLE = `# The sim request's id vocabulary

## Buffs

Prose the parser must skip, including a | pipe | in a sentence.

| id | lands in |
| --- | --- |
| \`battle_shout\` | RaidBuffs |
| \`battle_shout:improved\` | RaidBuffs |
| \`blessing_of_kings\` | IndividualBuffs |
| \`sunder_armor\` | Debuffs |
| \`atiesh_mage\` | PartyBuffs |
| \`songflower_serenade\` | IndividualBuffs |

## World buffs

| id |
| --- |
| \`songflower_serenade\` |

## Consumables

| id | sets |
| --- | --- |
| \`flask_of_supreme_power\` | Consumes.flask |
| \`main_hand_imbue:shadow_oil\` | Consumes.main_hand_imbue |
| \`food_grilled_squid\` | Consumes.food |

## Professions

| slug | engine enum |
| --- | --- |
| \`alchemy\` | Profession.Alchemy |

## Stats

| id | proto.Stat |
| --- | --- |
| \`attack_power\` | AttackPower |
| \`crit\` | Crit |
| \`melee_haste\` | MeleeHaste |
| \`mp5\` | MP5 |
| \`spell_haste\` | SpellHaste |
`;

describe('parseIdsMarkdown', () => {
  const parsed = parseIdsMarkdown(SAMPLE);

  it('reads every buff row with the message it lands in', () => {
    expect(parsed.buffs).toContainEqual({ id: 'battle_shout', message: 'RaidBuffs' });
    expect(parsed.buffs).toContainEqual({ id: 'sunder_armor', message: 'Debuffs' });
    expect(parsed.buffs).toContainEqual({ id: 'atiesh_mage', message: 'PartyBuffs' });
  });

  it('keeps a graded id as its own row, so the catalogue can pair it with the plain one', () => {
    expect(parsed.buffs).toContainEqual({ id: 'battle_shout:improved', message: 'RaidBuffs' });
  });

  it('reads the world-buff section when there is one', () => {
    expect(parsed.worldBuffs).toEqual(['songflower_serenade']);
  });

  it('reads consumables with the Consumes field they set', () => {
    expect(parsed.consumables).toContainEqual({
      id: 'main_hand_imbue:shadow_oil',
      sets: 'Consumes.main_hand_imbue',
    });
  });

  it('reads professions as plain slugs', () => {
    expect(parsed.professions).toEqual(['alchemy']);
  });

  it('reads the Stats section contract A7 adds, in the enum’s own snake case', () => {
    expect(parsed.stats).toEqual(['attack_power', 'crit', 'melee_haste', 'mp5', 'spell_haste']);
    // A7 and 10.8: haste is split, MP5 is `mp5`, and hit and crit are single stats.
    expect(parsed.stats).not.toContain('haste');
    expect(parsed.stats).not.toContain('melee_crit');
  });

  it('skips prose, headings and the separator row rather than reading them as ids', () => {
    expect(parsed.buffs.map((row) => row.id)).not.toContain('---');
    expect(parsed.buffs.map((row) => row.id)).not.toContain('id');
    expect(parsed.buffs).toHaveLength(6);
  });

  it('answers empty lists for a document with no tables at all', () => {
    expect(parseIdsMarkdown('# nothing here')).toEqual({
      buffs: [],
      consumables: [],
      professions: [],
      worldBuffs: [],
      stats: [],
    });
  });
});
