import { describe, expect, it } from 'vitest';
import { pickerClasses, specSummary } from './class-picker';
import { CLASS_SLUGS } from './ui/primitives';

describe('pickerClasses', () => {
  const classes = pickerClasses();

  it('lists the nine classes in the game order, linked to their guides', () => {
    expect(classes.map((c) => c.slug)).toEqual(CLASS_SLUGS);
    expect(classes[2]).toMatchObject({ name: 'Hunter', href: '/guides/hunter' });
  });

  it('lists each class in talent-tree order with every spec linked', () => {
    const hunter = classes.find((c) => c.slug === 'hunter');
    expect(hunter?.specs.map((s) => s.name)).toEqual(['Beast Mastery', 'Marksmanship', 'Survival']);
    expect(hunter?.specs[0]).toMatchObject({
      key: 'hunter/beast-mastery',
      href: '/guides/hunter/beast-mastery',
    });
  });

  it('takes other links and icons', () => {
    const bis = pickerClasses({ specHref: (c, s) => `/bis/${c}/${s}`, specIcon: (s) => `/i/${s.icon}` });
    expect(bis[0].href).toBeUndefined();
    expect(bis[0].specs[0].href).toBe('/bis/warrior/arms');
    expect(bis[0].specs[0].icon).toMatch(/^\/i\//);
  });
});

describe('specSummary', () => {
  it('names a role variant that shares a tree once', () => {
    const druid = pickerClasses().find((c) => c.slug === 'druid')!;
    expect(specSummary(druid, ', ')).toBe('Balance, Feral, Restoration');
  });
});
