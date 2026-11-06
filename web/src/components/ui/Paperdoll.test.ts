import { createRawSnippet } from 'svelte';
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import type { DollCell } from '../../lib/ui/paperdoll';
import Paperdoll from './Paperdoll.svelte';

const BUILD = '1.60.1.70790';
const CELLS: DollCell[] = [
  {
    slot: 'head',
    label: 'Head',
    flag: { text: 'Level 20', tone: 'muted' },
    item: { id: 1, name: "Brawler's Leather Hood", quality: 3, icon: 'a' },
    source: 'Crafted · Leatherworking (skill 100)',
    stats: '+10 Sta · +8 Agi',
    gain: {
      value: '+1.5',
      unit: 'DPS',
      basis: { prefix: 'vs', item: { id: 2, name: "Trapper's Leather Hood", quality: 2 } },
    },
  },
  {
    slot: 'trinket1',
    label: 'Trinket',
    item: null,
    empty: 'No trinket you can get at 20 to 29 raises your damage.',
  },
  { slot: 'hands', label: 'Hands', item: { id: 3, name: 'Gloves', quality: 3, icon: 'b' }, badge: 'upgrade' },
];

function doll(props: Record<string, unknown> = {}): string {
  return render(Paperdoll, { props: { cells: CELLS, mode: 'pick', build: BUILD, ...props } }).body.replace(
    /<!--[^>]*-->/g,
    '',
  );
}

describe('Paperdoll', () => {
  it('draws each slot as two lines: name and stats, then the source alone; the next best under the gain', () => {
    const body = doll();
    expect(body).toContain('role="listbox"');
    expect(body).toContain('data-testid="paperdoll-slot-head"');
    expect(body).toContain("Brawler's Leather Hood");
    expect(body).toContain('Crafted · Leatherworking (skill 100)');
    expect(body).toContain('+10 Sta · +8 Agi');
    expect(body).toContain('title="vs Trapper\'s Leather Hood"');
    expect(body.replace(/<[^>]+>/g, '')).toContain("vs Trapper's Leather Hood");
    expect(body).toContain('data-testid="paperdoll-basis-head"');
    expect(body).toMatch(
      /ui-doll-line2[^>]*>\s*<span class="ui-doll-src"[^>]*>Crafted · Leatherworking \(skill 100\)<\/span>\s*<\/span>/,
    );
    expect(body).toContain('Level 20');
  });

  it('explains an empty slot and keeps it out of selection and the tab order', () => {
    const body = doll();
    expect(body).toContain('No trinket you can get at 20 to 29 raises your damage.');
    expect(body).toMatch(/data-doll-slot="trinket1"[^>]*|aria-disabled="true"/);
  });

  it('server-renders the side inspector on the first upgrade, reserving its place', () => {
    const detail = createRawSnippet((cell: () => DollCell) => ({
      render: () => `<p>detail ${cell().slot}</p>`,
    }));
    const body = doll({ detail });
    expect(body).toContain('data-testid="paperdoll-inspector"');
    expect(body).toContain('detail hands');
    expect(body).toMatch(
      /data-doll-slot="hands"[^>]*aria-selected="true"|aria-selected="true"[^>]*data-doll-slot="hands"/,
    );
  });

  it('draws 17 skeleton rows at the final geometry while loading, and the error with a retry', () => {
    expect(doll({ loading: true }).match(/ui-doll-cell/g)?.length).toBe(17);
    const failed = doll({ failed: true, onretry: () => {} });
    expect(failed).toContain('We could not load the gear list.');
    expect(failed).toContain('<button');
  });

  it('groups slots under their source in the Where to go view', () => {
    const body = doll({
      groups: [{ key: 'wc', title: 'Wailing Caverns', count: '1 piece', slots: ['head'] }],
    });
    expect(body).toContain('data-view="where"');
    expect(body).toContain('Wailing Caverns');
    expect(body).not.toContain('paperdoll-slot-hands');
  });
});
