// web/src/lib/planner/url.test.ts
import { describe, expect, it } from 'vitest';
import { plannerSearchFor } from './url';

describe('plannerSearchFor', () => {
  it('writes the class and race a bare planner has moved to', () => {
    expect(plannerSearchFor('', 'mage', 'gnome', null)).toBe('?class=mage&race=gnome');
  });

  it('is null when the address already says the same thing', () => {
    expect(plannerSearchFor('?class=mage&race=gnome', 'mage', 'gnome', null)).toBeNull();
  });

  it('replaces a stale class and race', () => {
    expect(plannerSearchFor('?class=warrior&race=dwarf', 'mage', 'gnome', null)).toBe(
      '?class=mage&race=gnome',
    );
  });

  it('drops the race when none is chosen', () => {
    expect(plannerSearchFor('?class=warrior&race=dwarf', 'mage', '', null)).toBe('?class=mage');
  });

  it('keeps a code while it still describes the build on screen', () => {
    const code = { classSlug: 'warrior', raceSlug: 'orc' };
    expect(plannerSearchFor('?code=FS1%3Ax', 'warrior', 'orc', code)).toBeNull();
  });

  it('keeps a bare address while a restored pointer still describes the build on screen', () => {
    const pointer = { classSlug: 'warrior', raceSlug: 'orc' };
    expect(plannerSearchFor('', 'warrior', 'orc', pointer)).toBeNull();
    expect(plannerSearchFor('', 'mage', 'gnome', pointer)).toBe('?class=mage&race=gnome');
  });

  it('drops the code once the visitor moves off its class or race', () => {
    const code = { classSlug: 'warrior', raceSlug: 'orc' };
    expect(plannerSearchFor('?code=FS1%3Ax', 'mage', 'gnome', code)).toBe('?class=mage&race=gnome');
    expect(plannerSearchFor('?code=FS1%3Ax', 'warrior', 'tauren', code)).toBe('?class=warrior&race=tauren');
  });

  it('drops a code that never decoded', () => {
    expect(plannerSearchFor('?code=broken', 'warrior', 'orc', null)).toBe('?class=warrior&race=orc');
  });

  it('leaves unrelated parameters alone', () => {
    expect(plannerSearchFor('?utm=x&class=warrior', 'mage', 'gnome', null)).toBe(
      '?utm=x&class=mage&race=gnome',
    );
  });
});
