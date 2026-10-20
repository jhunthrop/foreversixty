import { describe, expect, it, vi } from 'vitest';

const fetchSimInput = vi.fn();
vi.mock('./sim/api', () => ({ fetchSimInput: (...args: unknown[]) => fetchSimInput(...args) }));

const addonExport = await import('./addon-export');
const { lookupAddonExport } = addonExport;

const PATH = { region: 'us' as const, ruleset: 'normal' as const, slug: 'thrallgar' };

describe('lookupAddonExport', () => {
  it('carries the FS1 code when the newest source is an addon export', async () => {
    fetchSimInput.mockResolvedValueOnce({
      spec: 'warrior-fury',
      gear: 'FS1:1.60.1.69893:warrior:human:0/0/0:',
      talents: '',
      buffs: [],
      captured_at: '2026-09-20T00:00:00Z',
      source: 'addon',
    });

    const result = await lookupAddonExport(PATH);
    expect(result.path).toBe(PATH);
    expect(result.code).toBe('FS1:1.60.1.69893:warrior:human:0/0/0:');
  });

  it('carries the FS1 code when the newest source is a Battle.net import', async () => {
    fetchSimInput.mockResolvedValueOnce({
      spec: 'warrior-fury',
      gear: 'FS1:1.60.1.69893:warrior:human:0/0/0:',
      talents: '',
      buffs: [],
      captured_at: '2026-09-20T00:00:00Z',
      source: 'blizzard',
    });

    const result = await lookupAddonExport(PATH);
    expect(result.code).toBe('FS1:1.60.1.69893:warrior:human:0/0/0:');
  });

  it('reads no export when the newest source is a logged fight', async () => {
    fetchSimInput.mockResolvedValueOnce({
      spec: 'mage-fire',
      gear: { trinkets: [] },
      talents: '31/0/20',
      buffs: [],
      captured_at: '2026-09-20T00:00:00Z',
      source: 'fight',
    });

    const result = await lookupAddonExport(PATH);
    expect(result.code).toBeNull();
  });

  it('reads no export, not an error, when the API call fails', async () => {
    fetchSimInput.mockRejectedValueOnce(new Error('offline'));

    const result = await lookupAddonExport(PATH);
    expect(result.code).toBeNull();
  });
});

describe('exportHasGear', () => {
  const { exportHasGear, slotsHaveGear, buildHasGear, simCharacterHasGear } = addonExport;

  it('is true for an export with a worn item', () => {
    expect(exportHasGear('FS1:1.60.1.69893:warrior:human:0/0/0:head=12640')).toBe(true);
  });

  it('is true for an export with talent points and no gear', () => {
    expect(exportHasGear('FS1:1.60.1.69893:warrior:human:5/0/0:')).toBe(true);
  });

  it('is false for an export with neither gear nor talents', () => {
    expect(exportHasGear('FS1:1.60.1.69893:warrior:human:0/0/0:')).toBe(false);
  });

  it('is false for nothing, an empty string or a string that does not decode', () => {
    expect(exportHasGear(null)).toBe(false);
    expect(exportHasGear(undefined)).toBe(false);
    expect(exportHasGear('')).toBe(false);
    expect(exportHasGear('not a code')).toBe(false);
  });

  it('reads the API gear panel, a /v1/me build and a sim character by the same rule', () => {
    expect(slotsHaveGear([{ worn: null }, { worn: null }])).toBe(false);
    expect(slotsHaveGear([{ worn: null }, { worn: { id: 1 } }])).toBe(true);
    expect(buildHasGear(undefined)).toBe(false);
    expect(buildHasGear({})).toBe(false);
    expect(buildHasGear({ gear: {}, talents: { points: [0, 0, 0] } })).toBe(false);
    expect(buildHasGear({ gear: { head: 5 } })).toBe(true);
    expect(buildHasGear({ talents: { points: [0, 3, 0] } })).toBe(true);
    expect(simCharacterHasGear({ gear: {}, point_order: [] })).toBe(false);
    expect(simCharacterHasGear({ gear: { head: 5 }, point_order: [] })).toBe(true);
    expect(simCharacterHasGear({ gear: {}, point_order: [0] })).toBe(true);
  });
});
