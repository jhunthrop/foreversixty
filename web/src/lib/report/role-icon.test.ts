// web/src/lib/report/role-icon.test.ts
import { describe, expect, it } from 'vitest';
import { roleIcon, roleOf } from './role-icon';

describe('role icons', () => {
  it('knows the three roles and nothing else', () => {
    expect(roleOf('tank')).toBe('tank');
    expect(roleOf('healer')).toBe('healer');
    expect(roleOf('dps')).toBe('dps');
    expect(roleOf(undefined)).toBeNull();
    expect(roleOf('other')).toBeNull();
  });

  it('draws each role with the game icon the client uses for it, named for a reader', () => {
    expect(roleIcon('tank', '1.60.1')).toEqual({
      url: '/data/1.60.1/icons/ability_warrior_defensivestance.webp',
      label: 'Tank',
    });
    expect(roleIcon('healer', '1.60.1').url).toContain('spell_holy_flashheal');
    expect(roleIcon('dps', '1.60.1').label).toBe('Damage');
  });
});
