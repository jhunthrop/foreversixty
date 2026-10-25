import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import { characterHistoryFixture } from '../../fixtures/character-history';
import type { CharacterHistory } from '../../lib/character/history-api';
import type { ChartSize } from '../../lib/character/history-scales';
import ParseChart from './ParseChart.svelte';

const DESKTOP: ChartSize = { width: 588, height: 294, phone: false };

function parse(history: CharacterHistory, sizeOverride: ChartSize | null = DESKTOP): string {
  return render(ParseChart, { props: { history, sizeOverride } }).body.replace(/<!--[^>]*-->/g, '');
}

describe('ParseChart, a boss with several kills', () => {
  const html = parse(characterHistoryFixture());

  it('opens on the boss of the newest kill, with a tablist of pills and their counts', () => {
    expect(html).toContain('Parse by boss');
    expect(html).toContain('role="tablist"');
    expect(html).toMatch(/role="tab"[^>]*aria-selected="true"[^>]*aria-label="Golemagg, 5 kills"/);
    expect(html).toMatch(/aria-label="Gehennas, 1 kill"/);
    expect(html).toMatch(/aria-selected="false"[^>]*tabindex="-1"/);
  });

  it('reads out the newest kill with its parse colour', () => {
    expect(html).toMatch(/data-testid="history-readout-value">88</);
    expect(html).toContain('parse');
    expect(html).toContain('11 Dec 2026 · kill 5 of 5');
  });

  it('draws the chart without tier bands or gaps, with the ladder-bounded axis', () => {
    expect(html).toContain(
      'aria-label="Parse chart for Golemagg. Use the arrow keys to move between kills."',
    );
    expect(html).not.toContain('history-band-label');
    expect(html).not.toContain('history-gap"');
    expect(html).toContain('Up 34 since 17 Nov');
    expect(html).toMatch(/aria-label="Later kill"[^>]*aria-disabled="true"/);
  });

  it('lists the kills with the date first and the newest marked', () => {
    expect(html).toContain('Kills of Golemagg');
    expect(html.match(/data-marked="true"/g)).toHaveLength(1);
    expect(html).toContain('Parse by kill, oldest first');
  });
});

describe('ParseChart, a boss with one kill', () => {
  const base = characterHistoryFixture();
  const history = { ...base, encounters: [base.encounters[1]] };
  const html = parse(history);

  it('writes the one-kill sentence with no chart and no readout, and the one-row list', () => {
    expect(html).toContain('One kill so far. A line needs two.');
    expect(html).not.toContain('history-plot');
    expect(html).not.toContain('history-readout-value');
    expect(html).toContain('Kills of Gehennas');
  });
});

describe('ParseChart, no kills', () => {
  it('writes one sentence and no picker', () => {
    const html = parse(characterHistoryFixture({ encounters: [] }));
    expect(html).toContain('No kills on record yet. Parse charts start with your first kill.');
    expect(html).not.toContain('role="tablist"');
    expect(html).not.toContain('history-plot');
  });
});
