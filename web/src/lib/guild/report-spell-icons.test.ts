// web/src/lib/guild/report-spell-icons.test.ts
// A log's spell icons: the chain a spell goes down until it has one, never an empty socket.
import { describe, expect, it } from 'vitest';
import { NO_SPELL_ICONS, spellIconOf, spellIconStem } from './report-spell-icons';

const TABLE = {
  '133': 'spell_fire_flamebolt',
  '9': 'inv_misc_questionmark',
  '1464': 'ability_warrior_decisivestrike',
};
const BY_NAME = { Slam: 'ability_warrior_decisivestrike' };

describe('spellIconStem', () => {
  it('draws the client attack icon for a melee swing', () => {
    expect(spellIconStem(TABLE, BY_NAME, 0)).toBe('ability_meleedamage');
    expect(spellIconStem(TABLE, BY_NAME, 0, 'Melee')).toBe('ability_meleedamage');
  });

  it('takes the client override, then the id table', () => {
    expect(spellIconStem(TABLE, BY_NAME, 75, 'Auto Shot')).toBe('ability_whirlwind');
    expect(spellIconStem(TABLE, BY_NAME, 133, 'Fireball')).toBe('spell_fire_flamebolt');
  });

  it('takes a same-named spell for another rank or a Forever id', () => {
    expect(spellIconStem(TABLE, BY_NAME, 11430, 'Slam')).toBe('ability_warrior_decisivestrike');
  });

  it('falls to the school, then the attack icon; the placeholder never counts', () => {
    expect(spellIconStem(TABLE, BY_NAME, 1259652, 'Shock', 8)).toBe('spell_nature_lightning');
    expect(spellIconStem(TABLE, BY_NAME, 1259652, 'Shock', 20)).toBe('spell_fire_fire');
    expect(spellIconStem(TABLE, BY_NAME, 9, 'Mystery')).toBe('ability_meleedamage');
    expect(spellIconStem(TABLE, BY_NAME, 7183, 'Elixir of Minor Defense')).toBe('ability_meleedamage');
  });
});

describe('spellIconOf', () => {
  const iconOf = spellIconOf('1.60.1', TABLE, BY_NAME, true);

  it("without the chain is the id table alone (the guild report's death causes)", () => {
    const tableOnly = spellIconOf('1.60.1', TABLE, BY_NAME);
    expect(tableOnly(133)).toBe('/data/1.60.1/icons/spell_fire_flamebolt.webp');
    expect(tableOnly(0)).toBeNull();
    expect(tableOnly(9)).toBeNull();
    expect(tableOnly(11430, 'Slam')).toBeNull();
  });

  it('is the build icon url, and the socket only before the tables load', () => {
    expect(iconOf(133)).toBe('/data/1.60.1/icons/spell_fire_flamebolt.webp');
    expect(iconOf(0)).toBe('/data/1.60.1/icons/ability_meleedamage.webp');
    expect(iconOf(null)).toBeNull();
    expect(NO_SPELL_ICONS(133)).toBeNull();
  });
});
