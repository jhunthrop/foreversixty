// web/src/lib/tiers/tier-controls.test.ts
import { describe, expect, it } from 'vitest';
import { hrefWithFaction, initialFaction, nextTabIndex, parseFaction } from './tier-controls';

describe('parseFaction', () => {
  it('reads horde and alliance from the query', () => {
    expect(parseFaction('?faction=horde')).toBe('horde');
    expect(parseFaction('?faction=alliance')).toBe('alliance');
  });

  it('falls back to alliance for nothing or nonsense', () => {
    expect(parseFaction('')).toBe('alliance');
    expect(parseFaction('?faction=orc')).toBe('alliance');
    expect(parseFaction('?other=horde')).toBe('alliance');
  });
});

describe('hrefWithFaction', () => {
  it('carries horde and drops the default', () => {
    expect(hrefWithFaction('/tiers/tank', 'horde')).toBe('/tiers/tank?faction=horde');
    expect(hrefWithFaction('/tiers/tank', 'alliance')).toBe('/tiers/tank');
  });
});

describe('initialFaction (owner decision T3)', () => {
  it('lets the URL win over the current character', () => {
    expect(initialFaction('?faction=alliance', { faction: 'horde' })).toBe('alliance');
    expect(initialFaction('?faction=horde', { faction: 'alliance' })).toBe('horde');
  });

  it("opens on the current character's faction with no URL choice", () => {
    expect(initialFaction('', { faction: 'horde' })).toBe('horde');
    expect(initialFaction('?faction=orc', { faction: 'horde' })).toBe('horde');
  });

  it('falls back to Alliance with no pointer, an older pointer, or nonsense', () => {
    expect(initialFaction('', null)).toBe('alliance');
    expect(initialFaction('', {})).toBe('alliance');
    expect(initialFaction('', { faction: 'scourge' })).toBe('alliance');
  });
});

describe('nextTabIndex', () => {
  it('wraps with the arrow keys and jumps with Home and End', () => {
    expect(nextTabIndex('ArrowRight', 2, 3)).toBe(0);
    expect(nextTabIndex('ArrowLeft', 0, 3)).toBe(2);
    expect(nextTabIndex('Home', 2, 3)).toBe(0);
    expect(nextTabIndex('End', 0, 3)).toBe(2);
    expect(nextTabIndex('Enter', 0, 3)).toBeNull();
  });
});
