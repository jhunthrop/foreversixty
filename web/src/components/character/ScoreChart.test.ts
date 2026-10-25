import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import { characterHistoryFixture } from '../../fixtures/character-history';
import type { CharacterHistory } from '../../lib/character/history-api';
import type { ChartSize } from '../../lib/character/history-scales';
import ScoreChart from './ScoreChart.svelte';

const DESKTOP: ChartSize = { width: 588, height: 294, phone: false };
const PHONE: ChartSize = { width: 320, height: 213, phone: true };

function score(history: CharacterHistory, sizeOverride: ChartSize | null = DESKTOP): string {
  return render(ScoreChart, { props: { history, sizeOverride } }).body.replace(/<!--[^>]*-->/g, '');
}

describe('ScoreChart, the owner with data', () => {
  const html = score(characterHistoryFixture());

  it('writes the sub-heading, the note with the rings key, and the readout of the newest day', () => {
    expect(html).toContain('Score over time');
    expect(html).toContain(
      "One point a day: the day's last rated fight. Rings are private or guild reports.",
    );
    expect(html).toMatch(/data-testid="history-readout-value">1830</);
    expect(html).toContain('Epic');
    expect(html).toContain('11 Dec 2026 · 2 fights');
  });

  it('draws the plot as a keyboard group, the line, tier bands, a gap band with its reason and the guide', () => {
    expect(html).toContain('role="group"');
    expect(html).toContain('tabindex="0"');
    expect(html).toContain('aria-label="Score chart. Use the arrow keys to move between points."');
    expect(html).toMatch(/<path d="M[^"]+" fill="none" stroke="var\(--color-gold\)"/);
    expect(html).toContain('data-testid="history-band-label"');
    expect(html).toContain('data-testid="history-gap"');
    expect(html).toContain('No rated');
    expect(html).toContain('data-testid="history-guide"');
  });

  it('draws the private and guild points as rings and the newest as the selected dot', () => {
    expect(html.match(/data-testid="history-ring"/g)).toHaveLength(2);
    expect(html.match(/data-testid="history-selected-dot"/g)).toHaveLength(1);
  });

  it('labels the x axis and breaks the line at the gap into two runs', () => {
    expect(html).toContain('>12 Oct<');
    expect(html).toContain('>11 Dec<');
    expect(html.match(/<path d="M[^"]*" fill="none" stroke="var\(--color-gold\)"/g)).toHaveLength(2);
  });

  it('writes the net change and the step buttons, the later one disabled at the newest point', () => {
    expect(html).toContain('Up 1090 since 12 Oct');
    expect(html).toContain('aria-label="Earlier point"');
    expect(html).toMatch(/aria-label="Later point"[^>]*aria-disabled="true"/);
    expect(html).not.toMatch(/aria-label="Earlier point"[^>]*aria-disabled/);
  });

  it('lists the fights behind the score with the newest day s rows marked', () => {
    expect(html).toContain('Fights behind the score');
    expect(html).toContain('Show all 9 fights');
    expect(html.match(/data-marked="true"/g)).toHaveLength(2);
  });

  it('carries the screen-reader table with the gap as a row', () => {
    expect(html).toContain('Forever Score by date, oldest first');
    expect(html).toContain('Gap from 26 Oct to 17 Nov: No rated fights');
  });
});

describe('ScoreChart, the narrow gap', () => {
  it('writes the reason under the chart when the band is too narrow for it', () => {
    const history = characterHistoryFixture({
      gaps: [{ from: '2026-10-26', to: '2026-11-17', reason: 'Report flagged' }],
    });
    const html = score(history, { width: 160, height: 107, phone: true });
    expect(html).toContain('Gap 26 Oct to 17 Nov: Report flagged');
    expect(html).not.toContain('data-testid="history-gap-label"');
  });
});

describe('ScoreChart, a phone', () => {
  it('draws three x labels', () => {
    expect(score(characterHistoryFixture(), PHONE).match(/data-testid="history-x-label"/g)).toHaveLength(3);
  });
});

describe('ScoreChart, before the first measure', () => {
  it('holds an empty plot box and still writes the readout', () => {
    const html = score(characterHistoryFixture(), null);
    expect(html).toContain('data-testid="history-plot"');
    expect(html).not.toContain('<svg class="absolute');
    expect(html).toContain('history-readout-value');
  });
});

describe('ScoreChart, one point', () => {
  const history = characterHistoryFixture({
    score: characterHistoryFixture().score.slice(0, 1),
    gaps: [],
  });
  const html = score(history);

  it('has no chart box, writes the first-score sentence and keeps the readout and the list', () => {
    expect(html).not.toContain('data-testid="history-plot"');
    expect(html).toContain('First score recorded 12 Oct. A line needs a second point.');
    expect(html).toContain('history-readout-value');
    expect(html).toContain('Fights behind the score');
  });
});

describe('ScoreChart, public points only', () => {
  it('drops the rings key from the note', () => {
    const base = characterHistoryFixture();
    const history = {
      ...base,
      score: base.score.map((point) => ({ ...point, visibility: 'public' as const })),
    };
    const html = score(history);
    expect(html).not.toContain('Rings are private');
    expect(html).not.toContain('history-ring');
  });
});
