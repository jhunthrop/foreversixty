import { describe, expect, it } from 'vitest';
import { humanise, humaniseKey } from './humanise';

describe('humanise', () => {
  it('is empty for an empty string', () => {
    expect(humanise('')).toBe('');
  });

  it('capitalises a single word with no underscore', () => {
    expect(humanise('attack')).toBe('Attack');
  });

  it('splits underscores into spaces and capitalises only the first word', () => {
    expect(humanise('rage_gain')).toBe('Rage gain');
    expect(humanise('attack_power')).toBe('Attack power');
  });

  it('leaves an already-capitalised first letter alone', () => {
    expect(humanise('Attack_power')).toBe('Attack power');
  });
});

describe('humaniseKey', () => {
  // D48: a raw `spell:<id>`/`item:<id>` key reaching the summary sentence and the report
  // tables, because the build's own table never carried the id (or had not loaded yet).
  // humaniseKey never shows the colon-joined key itself -- it humanises the kind and keeps
  // the numeric id as a plain number, the same "Item <id>" shape combos.ts's own
  // substitutionLabel already uses for an unnamed item substitution.
  it('humanises a spell or item id with no raw colon in the result', () => {
    expect(humaniseKey('spell:20662')).toBe('Spell 20662');
    expect(humaniseKey('item:14554')).toBe('Item 14554');
  });

  it('humanises a multi-segment, hyphenated key, keeping the trailing numeric id plain', () => {
    // The newcomer review's own repro: a Droptimizer boss id whose source name never rode
    // along, falling back to the loot-table id itself -- `dungeon:blackrock-spire:175245`.
    expect(humaniseKey('dungeon:blackrock-spire:175245')).toBe('Dungeon Blackrock spire 175245');
  });

  it('is plain humanise() for a key with no colon at all', () => {
    expect(humaniseKey('rage_gain')).toBe('Rage gain');
  });
});
