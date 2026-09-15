// web/src/lib/report/tree-sizes.test.ts
import { describe, expect, it, vi } from 'vitest';
import { DataLoadError } from '../planner/load';
import type { TalentFile } from '../planner/types';
import { classSlugFromName, resolveTreeSizes } from './tree-sizes';

function talentFile(sizes: number[]): TalentFile {
  return {
    build: 'b',
    class_id: 1,
    class_slug: 'warrior',
    trees: sizes.map((size, index) => ({
      id: index,
      name: `Tree ${index}`,
      position: index,
      talents: Array.from({ length: size }, (_, i) => ({
        id: i,
        name: `Talent ${i}`,
        icon: '',
        max_rank: 1,
        tier: 0,
        column: 0,
        prereq_talent_id: null,
        prereq_rank: null,
        ranks: [],
      })),
    })),
  };
}

describe('classSlugFromName', () => {
  it('lowercases and hyphenates', () => {
    expect(classSlugFromName('Death Knight')).toBe('death-knight');
  });
});

describe('resolveTreeSizes', () => {
  it('fetches only the classes not already known, and reports their tree sizes', async () => {
    const loadTalents = vi.fn().mockResolvedValue(talentFile([3, 2, 1]));
    const result = await resolveTreeSizes(['Warrior', 'Paladin'], new Set(['Paladin']), loadTalents, 'b');
    expect(loadTalents).toHaveBeenCalledTimes(1);
    expect(loadTalents).toHaveBeenCalledWith('b', 'warrior');
    expect(result).toEqual([['Warrior', [3, 2, 1]]]);
  });

  it('caches a 404 as an empty list: the class genuinely has no talent data', async () => {
    const loadTalents = vi.fn().mockRejectedValue(new DataLoadError('not found', { status: 404 }));
    const result = await resolveTreeSizes(['Death Knight'], new Set(), loadTalents, 'b');
    expect(result).toEqual([['Death Knight', []]]);
  });

  it('drops a non-404 failure instead of caching it as no data', async () => {
    const loadTalents = vi.fn().mockRejectedValue(new DataLoadError('server error', { status: 500 }));
    const result = await resolveTreeSizes(['Warrior'], new Set(), loadTalents, 'b');
    expect(result).toEqual([]);
  });

  it('drops a network failure (no status at all) the same way', async () => {
    const loadTalents = vi.fn().mockRejectedValue(new DataLoadError('offline'));
    const result = await resolveTreeSizes(['Warrior'], new Set(), loadTalents, 'b');
    expect(result).toEqual([]);
  });

  it('leaves a class dropped by a transient failure out of `known`, so the next call retries it', async () => {
    const loadTalents = vi.fn().mockRejectedValueOnce(new DataLoadError('server error', { status: 500 }));
    const first = await resolveTreeSizes(['Warrior'], new Set(), loadTalents, 'b');
    expect(first).toEqual([]);

    // A caller that only ever merges resolveTreeSizes' own output into its cache never adds
    // 'Warrior' after a failure like this -- `known` on the next call is exactly what it was
    // before, not the wrong answer a bare catch-and-cache would have pinned.
    loadTalents.mockResolvedValueOnce(talentFile([3, 2, 1]));
    const second = await resolveTreeSizes(['Warrior'], new Set(), loadTalents, 'b');
    expect(second).toEqual([['Warrior', [3, 2, 1]]]);
    expect(loadTalents).toHaveBeenCalledTimes(2);
  });

  it('does not refetch a class already known, even one known as having no data', async () => {
    const loadTalents = vi.fn();
    const result = await resolveTreeSizes(['Death Knight'], new Set(['Death Knight']), loadTalents, 'b');
    expect(loadTalents).not.toHaveBeenCalled();
    expect(result).toEqual([]);
  });
});
