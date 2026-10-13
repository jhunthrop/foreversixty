// web/src/lib/gear/qualities.test.ts
// Task L8.6 fix round 1 ruling 2: the Gear section's item facts (quality, icon, two-handed)
// come from the build's items/<class>.json. One read per class while it succeeds; a failed
// read is not remembered, so the next open retries; a name the table lacks, or a failed read,
// takes the body colour, never white.
import { beforeEach, describe, expect, it, vi } from 'vitest';
import { GEAR_FIXTURE_ITEMS } from '../../fixtures/character-gear';

const loadItems = vi.fn();
vi.mock('../planner/load', () => ({ loadItems: (...args: unknown[]) => loadItems(...args) }));

const { itemFacts, nameClass, iconOf, resetItemFacts } = await import('./qualities');

beforeEach(() => {
  loadItems.mockReset();
  resetItemFacts();
});

describe('itemFacts (L8.6 fix round 1 ruling 2)', () => {
  it('a hit: reads the class’s file once and serves the second call from it', async () => {
    loadItems.mockResolvedValue({ build: 'b', class_slug: 'priest', items: GEAR_FIXTURE_ITEMS });
    const first = await itemFacts('b', 'priest');
    const second = await itemFacts('b', 'priest');
    expect(loadItems).toHaveBeenCalledTimes(1);
    expect(loadItems).toHaveBeenCalledWith('b', 'priest');
    expect(second).toBe(first);
    expect(first?.get(16921)).toEqual({ quality: 4, icon: 'inv_helmet_24', twoHand: false });
    expect(first?.get(13000)?.twoHand).toBe(true);
    expect(nameClass(first ?? null, 16921)).toBe('text-rarity-epic-text');
    expect(nameClass(first ?? null, 10504)).toBe('text-rarity-rare-text');
  });

  it('a miss: an id the table lacks takes the body colour and no icon', async () => {
    loadItems.mockResolvedValue({ build: 'b', class_slug: 'priest', items: GEAR_FIXTURE_ITEMS });
    const facts = await itemFacts('b', 'priest');
    expect(facts?.get(99001)).toBeUndefined();
    expect(nameClass(facts ?? null, 99001)).toBe('text-text');
    expect(iconOf(facts ?? null, 99001)).toBeNull();
  });

  it('a failure resolves to null, and every name takes the body colour', async () => {
    loadItems.mockRejectedValue(new Error('Talent data did not load (500)'));
    const facts = await itemFacts('b', 'priest');
    expect(facts).toBeNull();
    expect(nameClass(facts, 16921)).toBe('text-text');
    expect(iconOf(facts, 16921)).toBeNull();
  });

  it('a retry: the failure is not cached, so the next open reads the file again', async () => {
    loadItems.mockRejectedValueOnce(new Error('offline'));
    loadItems.mockResolvedValueOnce({ build: 'b', class_slug: 'priest', items: GEAR_FIXTURE_ITEMS });
    expect(await itemFacts('b', 'priest')).toBeNull();
    const retried = await itemFacts('b', 'priest');
    expect(loadItems).toHaveBeenCalledTimes(2);
    expect(retried?.get(16921)?.quality).toBe(4);
  });

  it('the questionmark placeholder is no icon (the /bis rule)', async () => {
    loadItems.mockResolvedValue({
      build: 'b',
      class_slug: 'priest',
      items: [{ ...GEAR_FIXTURE_ITEMS[0], icon: 'inv_misc_questionmark' }],
    });
    const facts = await itemFacts('b', 'priest');
    expect(iconOf(facts, GEAR_FIXTURE_ITEMS[0].id)).toBeNull();
  });
});
