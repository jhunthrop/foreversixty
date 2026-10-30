// web/src/lib/ui/select-options.test.ts
// The shared option shapes (controls-charts spec §A10): every player row carries the crest's
// place and its class colour, even when the log never named the class.
import { describe, expect, it } from 'vitest';
import { classOption, iconOption } from './select-options';

describe('classOption', () => {
  it('a known class: the ringed crest and the class colour', () => {
    expect(classOption({ value: 'g', label: 'Baelgrim', className: 'Warrior' })).toMatchObject({
      icon: { kind: 'crest', src: '/icons/hd/crests/warrior.webp' },
      colour: '--color-class-warrior',
    });
  });

  it('a class the log never named keeps the crest slot, empty, and no colour', () => {
    const option = classOption({ value: 't', label: 'Thalgrit', className: undefined });
    expect(option.icon).toEqual({ kind: 'crest', src: '' });
    expect(option.colour).toBeUndefined();
  });
});

describe('iconOption', () => {
  it('draws an icon only when there is one', () => {
    expect(iconOption('a', 'Slam', '/i.webp', 'spell').icon).toEqual({ kind: 'spell', src: '/i.webp' });
    expect(iconOption('a', 'Slam', null).icon).toBeUndefined();
  });
});
