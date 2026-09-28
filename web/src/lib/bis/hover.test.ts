// web/src/lib/bis/hover.test.ts
import { afterEach, describe, expect, it, vi } from 'vitest';
import {
  bandForLevel,
  bisAssetUrl,
  bisPageHref,
  fetchBisFile,
  MAX_BAND,
  MIN_BAND,
  resetBisHoverCache,
  slotHoverDiff,
} from './hover';
import type { BisBand, BisFile, BisSlot } from './types';

function slot(item_id: number, item_name: string): BisSlot {
  return {
    slot: 'head',
    item_id,
    item_name,
    source: 'Vendor: Someone',
    source_kind: 'vendor',
    score: 1,
    verified: true,
  };
}

function band(bandLevel: number, faction: BisBand['faction'], slots: BisSlot[]): BisBand {
  return {
    spec: 'hunter-marksmanship',
    band: bandLevel,
    faction,
    race: 'Dwarf',
    talents: '',
    talent_points: 1,
    weights: [],
    slots,
    set_dps: 0,
    no_source_count: 0,
    new_at_band: [],
    weights_run_seconds: 0,
    verify_run_seconds: 0,
  };
}

function fileWith(bands: BisBand[]): BisFile {
  return {
    spec: 'hunter-marksmanship',
    build: 'b1',
    engine_version: 'e1',
    generated_at: '2026-01-01',
    bands,
  };
}

describe('bandForLevel', () => {
  it('rounds down to the nearest 5', () => {
    expect(bandForLevel(27)).toBe(25);
    expect(bandForLevel(30)).toBe(30);
    expect(bandForLevel(34)).toBe(30);
  });

  it('clamps below the lowest band to the lowest band', () => {
    expect(bandForLevel(1)).toBe(MIN_BAND);
    expect(bandForLevel(9)).toBe(MIN_BAND);
  });

  it('clamps at and above the top band to the top band', () => {
    expect(bandForLevel(60)).toBe(MAX_BAND);
    expect(bandForLevel(75)).toBe(MAX_BAND);
  });
});

describe('bisAssetUrl', () => {
  it('builds the per-spec static asset path', () => {
    expect(bisAssetUrl('1.60.1.70009', 'hunter-marksmanship')).toBe(
      '/data/1.60.1.70009/bis/hunter-marksmanship.json',
    );
  });
});

describe('slotHoverDiff', () => {
  it('is not new at the file’s lowest band, even with a pick', () => {
    const file = fileWith([band(10, 'alliance', [slot(1, 'Old Cap')])]);
    const diff = slotHoverDiff(file, 10, 'alliance', 'head');
    expect(diff.pick?.item_name).toBe('Old Cap');
    expect(diff.previous).toBeUndefined();
    expect(diff.isNewAtBand).toBe(false);
  });

  it('marks a slot new when the pick changes from the previous band', () => {
    const file = fileWith([
      band(10, 'alliance', [slot(1, 'Old Cap')]),
      band(15, 'alliance', [slot(2, 'New Cap')]),
    ]);
    const diff = slotHoverDiff(file, 15, 'alliance', 'head');
    expect(diff.pick?.item_name).toBe('New Cap');
    expect(diff.previous).toEqual({ band: 10, pick: expect.objectContaining({ item_name: 'Old Cap' }) });
    expect(diff.isNewAtBand).toBe(true);
  });

  it('is not new when the pick is unchanged from the previous band', () => {
    const file = fileWith([
      band(10, 'alliance', [slot(1, 'Same Cap')]),
      band(15, 'alliance', [slot(1, 'Same Cap')]),
    ]);
    const diff = slotHoverDiff(file, 15, 'alliance', 'head');
    expect(diff.isNewAtBand).toBe(false);
  });

  it('is new when the slot had no known source at the previous band', () => {
    const file = fileWith([band(10, 'alliance', []), band(15, 'alliance', [slot(1, 'First Cap')])]);
    const diff = slotHoverDiff(file, 15, 'alliance', 'head');
    expect(diff.isNewAtBand).toBe(true);
    expect(diff.previous?.pick).toBeUndefined();
  });

  it('is not new when the slot still has no known source', () => {
    const file = fileWith([band(10, 'alliance', []), band(15, 'alliance', [])]);
    const diff = slotHoverDiff(file, 15, 'alliance', 'head');
    expect(diff.pick).toBeUndefined();
    expect(diff.isNewAtBand).toBe(false);
  });

  it('skips a band the file has no entry for, using the previous band that does exist', () => {
    // Alliance has 10 and 20 only; band 15 (a different faction's band) is not this
    // faction's neighbour.
    const file = fileWith([
      band(10, 'alliance', [slot(1, 'Old Cap')]),
      band(15, 'horde', [slot(9, 'Horde-only Cap')]),
      band(20, 'alliance', [slot(2, 'New Cap')]),
    ]);
    const diff = slotHoverDiff(file, 20, 'alliance', 'head');
    expect(diff.previous?.band).toBe(10);
    expect(diff.previous?.pick?.item_name).toBe('Old Cap');
  });

  it('reads faction independently: the same band number, different faction, different pick', () => {
    const file = fileWith([band(10, 'alliance', [slot(1, 'A Cap')]), band(10, 'horde', [slot(2, 'H Cap')])]);
    expect(slotHoverDiff(file, 10, 'alliance', 'head').pick?.item_name).toBe('A Cap');
    expect(slotHoverDiff(file, 10, 'horde', 'head').pick?.item_name).toBe('H Cap');
  });
});

describe('bisPageHref', () => {
  it('links to the spec page with the faction query and band anchor', () => {
    expect(bisPageHref('hunter-marksmanship', 'alliance', 30)).toBe(
      '/bis/hunter/marksmanship?faction=alliance#band-alliance-30',
    );
  });

  it('falls back to the index for a spec key the generated list does not carry', () => {
    expect(bisPageHref('not-a-real-spec', 'horde', 10)).toBe('/bis');
  });
});

describe('fetchBisFile', () => {
  afterEach(() => {
    resetBisHoverCache();
    vi.unstubAllGlobals();
  });

  it('parses and returns the file on a 200', async () => {
    const file = fileWith([band(10, 'alliance', [slot(1, 'Cap')])]);
    vi.stubGlobal(
      'fetch',
      vi.fn(async () => new Response(JSON.stringify(file), { status: 200 })),
    );
    await expect(fetchBisFile('b1', 'hunter-marksmanship')).resolves.toEqual(file);
  });

  it('resolves null on a 404 rather than throwing', async () => {
    vi.stubGlobal(
      'fetch',
      vi.fn(async () => new Response('not found', { status: 404 })),
    );
    await expect(fetchBisFile('b1', 'priest-shadow')).resolves.toBeNull();
  });

  it('resolves null on a network failure rather than throwing', async () => {
    vi.stubGlobal(
      'fetch',
      vi.fn(async () => {
        throw new TypeError('offline');
      }),
    );
    await expect(fetchBisFile('b1', 'hunter-marksmanship')).resolves.toBeNull();
  });

  it('fetches a given (build, spec) at most once, sharing the cached promise', async () => {
    const file = fileWith([band(10, 'alliance', [slot(1, 'Cap')])]);
    const fetchSpy = vi.fn(async () => new Response(JSON.stringify(file), { status: 200 }));
    vi.stubGlobal('fetch', fetchSpy);

    await fetchBisFile('b1', 'hunter-marksmanship');
    await fetchBisFile('b1', 'hunter-marksmanship');
    await fetchBisFile('b1', 'hunter-marksmanship');

    expect(fetchSpy).toHaveBeenCalledTimes(1);
  });

  it('fetches each spec independently', async () => {
    const fetchSpy = vi.fn(async (input: string) => {
      const body = fileWith([band(10, 'alliance', [slot(1, String(input))])]);
      return new Response(JSON.stringify(body), { status: 200 });
    });
    vi.stubGlobal('fetch', fetchSpy);

    await fetchBisFile('b1', 'hunter-marksmanship');
    await fetchBisFile('b1', 'warrior-fury');

    expect(fetchSpy).toHaveBeenCalledTimes(2);
  });
});
