// web/src/lib/bis/landing-build.test.ts
import { describe, expect, it } from 'vitest';
import classes from '../../data/classes.json';
import activeBuild from '../../data/active-build.json';
import { SPECS } from '../sim/specs';
import { LANDING_BANDS } from './landing-bands';
import { loadLandingData } from './landing-build';

const data = loadLandingData(activeBuild.build);

describe('loadLandingData', () => {
  it('lists the classes in game order, each with its specs', () => {
    expect(data.cards.map((card) => card.slug)).toEqual(classes.map((cls) => cls.slug));
    expect(data.cards.flatMap((card) => card.specs)).toHaveLength(SPECS.length);
  });

  it('gives every spec row its tab icon from the build’s icon tree and its role', () => {
    const fury = data.cards
      .find((card) => card.slug === 'warrior')
      ?.specs.find((spec) => spec.specSlug === 'fury');
    expect(fury?.iconSrc).toMatch(new RegExp(`^/data/${activeBuild.build}/icons/[a-z0-9_]+\\.webp$`));
    expect(fury?.role).toBe('dps');
    const bear = data.cards
      .find((card) => card.slug === 'druid')
      ?.specs.find((spec) => spec.specSlug === 'feral-bear');
    expect(bear?.role).toBe('tank');
  });

  it('has one caption per band, the first always naming its sources', () => {
    expect(data.captions.map((caption) => caption.band)).toEqual([...LANDING_BANDS]);
    expect(data.captions[0]?.same).toBe(false);
  });

  it('keys the browser config by class slug and spec key, with each class’s first spec as its default', () => {
    expect(Object.keys(data.config.classes)).toEqual(classes.map((cls) => cls.slug));
    expect(data.config.classes.warrior?.defaultSpecKey).toBe(data.cards[0]?.specs[0]?.key);
    expect(data.config.specs['warrior-fury']).toEqual({
      classSlug: 'warrior',
      specSlug: 'fury',
      name: 'Fury',
    });
    expect(data.config.build).toBe(activeBuild.build);
  });

  it('reads the raid label as a sentence continues it', () => {
    expect(data.raidLabel).toBe(data.raidLabel.charAt(0).toLowerCase() + data.raidLabel.slice(1));
  });
});
