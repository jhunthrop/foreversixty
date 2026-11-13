// web/src/lib/tiers/tier-card.test.ts
// The tier list's share card: the published list in the page's order and letters, on the
// default faction, a version that follows the data, and a 1200 by 630 PNG.
import { mkdirSync, writeFileSync } from 'node:fs';
import { describe, expect, it } from 'vitest';
import { renderTierOg } from '../og-tier';
import { tierCardModel, tierCardVersion } from './tier-card';
import { ROLE_TABS, specFullName, tiersCopy } from './tier-copy';
import { TIER_ROLES, groupByTier } from './tier-list';
import { loadTierData } from './tier-load';

const data = loadTierData();

/** PNG width and height from the IHDR chunk. */
function pngSize(png: Uint8Array): { width: number; height: number } {
  const view = new DataView(png.buffer, png.byteOffset, png.byteLength);
  return { width: view.getUint32(16), height: view.getUint32(20) };
}

describe('tierCardModel', () => {
  it('names every damage spec in the page order under its tier letter', () => {
    const model = tierCardModel(data, 'dps');
    const expected = groupByTier(data.lists.alliance.dps).map((group) => ({
      letter: group.letter,
      names: group.rows.map(specFullName),
    }));
    expect(model.groups.map((g) => ({ letter: g.letter, names: g.specs.map((s) => s.name) }))).toEqual(
      expected,
    );
    expect(model.groups.flatMap((g) => g.specs).length).toBe(data.lists.alliance.dps.length);
    expect(model.title).toBe('Tier list at 60');
    expect(model.role).toBe(ROLE_TABS.dps);
  });

  it('carries the page meta line and the Updated stamp', () => {
    const { meta } = tierCardModel(data, 'dps');
    expect(meta.startsWith('Best in slot, raid buffs · vs a level 63 boss')).toBe(true);
    expect(meta).toMatch(/· Updated \d{1,2} [A-Z][a-z]{2}$/);
  });

  it('ranks the tank and healer lists without letters, in the page order', () => {
    for (const role of ['tank', 'healer'] as const) {
      const [group] = tierCardModel(data, role).groups;
      expect(group?.letter).toBeNull();
      expect(group?.specs.map((s) => s.name)).toEqual(data.lists.alliance[role].map(specFullName));
      expect(group?.specs.map((s) => s.rank)).toEqual(data.lists.alliance[role].map(tiersCopy.rankLabel));
    }
  });

  it('versions the card by what it draws', () => {
    const model = tierCardModel(data, 'dps');
    expect(tierCardVersion(model)).toBe(tierCardVersion(tierCardModel(data, 'dps')));
    const moved = { ...model, groups: [...model.groups].reverse() };
    expect(tierCardVersion(moved)).not.toBe(tierCardVersion(model));
    expect(tierCardVersion(model)).toMatch(/^[0-9a-f]{10}$/);
  });
});

describe('renderTierOg', () => {
  it.each(TIER_ROLES)(
    'paints the %s card at 1200 by 630',
    async (role) => {
      const png = await renderTierOg(tierCardModel(data, role));
      expect(pngSize(png)).toEqual({ width: 1200, height: 630 });
      expect(png.byteLength).toBeGreaterThan(20_000);
      const out = process.env.TIER_CARD_OUT;
      if (out !== undefined) {
        mkdirSync(out, { recursive: true });
        writeFileSync(`${out}/tiers-${role}.png`, png);
      }
    },
    30_000,
  );
});
