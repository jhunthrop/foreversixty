import { describe, expect, it } from 'vitest';
import { pillClassFor, pillLabelFor, sourceTitle } from './sources';

describe('source pills', () => {
  it('maps each kind to a class and label', () => {
    expect(pillClassFor('blizzard')).toBe('pill-blizzard');
    expect(pillClassFor('datamined')).toBe('pill-datamined');
    expect(pillClassFor('community')).toBe('pill-community');
    expect(pillClassFor('site')).toBe('pill-site');
    expect(pillLabelFor('site')).toBe('This site');
    expect(pillLabelFor('blizzard')).toBe('Blizzard');
  });
});

describe('sourceTitle', () => {
  it('drops the publisher the pill already names, and keeps a community site', () => {
    expect(sourceTitle('Blizzard, Class Deep Dives: Hunter and Druid')).toBe(
      'Class Deep Dives: Hunter and Druid',
    );
    expect(sourceTitle('Forever Sixty, race and class pairs')).toBe('Race and class pairs');
    expect(sourceTitle('Warcraft Tavern, Hunter and Druid Deep Dive')).toBe(
      'Warcraft Tavern, Hunter and Druid Deep Dive',
    );
  });
});
