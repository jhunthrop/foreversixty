// web/src/lib/guild/report-spell-icons.test.ts
// The death causes' spell icons: the build's table, nothing for an id it lacks or a swing.
import { describe, expect, it } from 'vitest';
import { NO_SPELL_ICONS, spellIconOf } from './report-spell-icons';

describe('spellIconOf', () => {
  const iconOf = spellIconOf('1.60.1', { '133': 'spell_fire_flamebolt', '9': 'inv_misc_questionmark' });

  it('draws the build icon for an id the table carries', () => {
    expect(iconOf(133)).toBe('/data/1.60.1/icons/spell_fire_flamebolt.webp');
  });

  it('draws nothing for an unknown id, the placeholder, or a swing', () => {
    expect(iconOf(18584)).toBeNull();
    expect(iconOf(9)).toBeNull();
    expect(iconOf(null)).toBeNull();
    expect(NO_SPELL_ICONS(133)).toBeNull();
  });
});
