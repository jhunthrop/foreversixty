import { experimental_AstroContainer as AstroContainer } from 'astro/container';
import { describe, expect, it } from 'vitest';
import EntryTable from './EntryTable.astro';

const rows = [
  {
    href: '/dungeons/hall-of-thanes',
    title: 'Hall of Thanes',
    meta: 'Dun Morogh',
    levelMin: 24,
    levelMax: 32,
  },
  {
    href: '/dungeons/alcaz-prison',
    title: 'Alcaz Prison',
    meta: 'Dustwallow Marsh',
    levelMin: 56,
    levelMax: 60,
  },
  { href: '/dungeons/drowned-city', title: 'Drowned City', meta: 'Vashjir' },
];

describe('EntryTable', () => {
  it('renders the headings, one linked row per entry, and rarity-coloured ranges', async () => {
    const c = await AstroContainer.create();
    const html = await c.renderToString(EntryTable, {
      props: { nameHeading: 'Dungeon', metaHeading: 'Zone', rows },
    });
    expect(html).toContain('>Dungeon<');
    expect(html).toContain('>Zone<');
    expect(html).toContain('href="/dungeons/hall-of-thanes"');
    expect(html).toContain('24–32');
    expect(html).toContain('text-rarity-uncommon');
    expect(html).toContain('56–60');
    expect(html).toContain('text-rarity-epic');
    expect(html).toContain('—');
    expect(html).toContain('text-muted');
  });
});
