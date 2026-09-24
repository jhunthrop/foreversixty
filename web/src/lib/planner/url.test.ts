// web/src/lib/planner/url.test.ts
import { describe, expect, it } from 'vitest';
import { plannerSearchFor } from './url';

const OPENED = { classSlug: 'warrior', raceSlug: 'human' };
const at = (
  search: string,
  cls: string,
  race: string,
  code: { classSlug: string; raceSlug: string } | null = null,
): string | null => plannerSearchFor(search, cls, race, code, OPENED);

describe('plannerSearchFor', () => {
  it('writes the class and race a bare planner has moved to', () => {
    expect(at('', 'mage', 'gnome')).toBe('?class=mage&race=gnome');
  });

  it('leaves a bare address bare while nothing has changed', () => {
    expect(at('', 'warrior', 'human')).toBeNull();
    expect(at('', 'warrior', 'dwarf')).toBe('?class=warrior&race=dwarf');
  });

  it('is null when the address already says the same thing', () => {
    expect(at('?class=mage&race=gnome', 'mage', 'gnome')).toBeNull();
  });

  it('replaces a stale class and race', () => {
    expect(at('?class=warrior&race=dwarf', 'mage', 'gnome')).toBe('?class=mage&race=gnome');
  });

  it('drops the race when none is chosen', () => {
    expect(at('?class=warrior&race=dwarf', 'mage', '')).toBe('?class=mage');
  });

  it('keeps a code while it still describes the build on screen', () => {
    const code = { classSlug: 'warrior', raceSlug: 'orc' };
    expect(at('?code=FS1%3Ax', 'warrior', 'orc', code)).toBeNull();
  });

  it('keeps a bare address while a restored pointer still describes the build on screen', () => {
    const pointer = { classSlug: 'warrior', raceSlug: 'orc' };
    expect(at('', 'warrior', 'orc', pointer)).toBeNull();
    expect(at('', 'mage', 'gnome', pointer)).toBe('?class=mage&race=gnome');
  });

  it('drops the code once the visitor moves off its class or race', () => {
    const code = { classSlug: 'warrior', raceSlug: 'orc' };
    expect(at('?code=FS1%3Ax', 'mage', 'gnome', code)).toBe('?class=mage&race=gnome');
    expect(at('?code=FS1%3Ax', 'warrior', 'tauren', code)).toBe('?class=warrior&race=tauren');
  });

  it('drops a code that never decoded', () => {
    expect(at('?code=broken', 'warrior', 'orc')).toBe('?class=warrior&race=orc');
  });

  it('leaves unrelated parameters alone', () => {
    expect(at('?utm=x&class=warrior', 'mage', 'gnome')).toBe('?utm=x&class=mage&race=gnome');
  });
});
