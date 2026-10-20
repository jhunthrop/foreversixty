import { describe, expect, it } from 'vitest';
import { INFERRED_CLASS_TITLE, specCell, unitClassMap } from './roster-class';
import type { Unit } from './types';

const unit = (guid: string, klass?: string, source?: string): Unit => ({
  guid,
  name: guid,
  kind: 'player',
  flags: 0,
  class: klass,
  class_source: source,
  first_seen: '',
  last_seen: '',
});

describe('specCell', () => {
  const units = unitClassMap([
    unit('p1', 'Warrior', 'inferred'),
    unit('p2'),
    unit('p3', 'Mage', 'combatant_info'),
  ]);

  it('shows the spec when the log has one', () => {
    expect(specCell({ guid: 'p1', class: 'Warrior', spec: 'Fury' }, units)).toMatchObject({
      text: 'Fury',
      muted: false,
      classColoured: false,
    });
  });
  it('shows the inferred class, class-coloured, with the hover that says how it was read', () => {
    expect(specCell({ guid: 'p1' }, units)).toEqual({
      class: 'Warrior',
      text: 'Warrior',
      title: INFERRED_CLASS_TITLE,
      muted: false,
      classColoured: true,
    });
  });
  it('does not claim a class from the spells when the log named it', () => {
    expect(specCell({ guid: 'p3' }, units).title).toBeUndefined();
  });
  it('is a muted dash, never the word Unknown, when not even the class is known', () => {
    const cell = specCell({ guid: 'p2' }, units);
    expect(cell).toMatchObject({ text: '—', muted: true, class: undefined });
    expect(cell.text).not.toMatch(/unknown/i);
  });
});
