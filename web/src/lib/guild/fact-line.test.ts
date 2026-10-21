// web/src/lib/guild/fact-line.test.ts
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import FactLine from '../../components/guild/FactLine.svelte';
import { FACT_SEPARATOR, FACT_TAIL, asFacts, bandFacts, factsText } from './fact-line';

/** The rendered words with tags dropped and no-break spaces read as spaces. */
function words(html: string): string {
  return html
    .replace(/<!--.*?-->/g, '')
    .replace(/<[^>]+>/g, '')
    .replace(/\u00a0/g, ' ');
}

describe('factsText', () => {
  it('joins facts with a space on each side of the dot, never glued', () => {
    expect(FACT_SEPARATOR).toBe(' · ');
    expect(
      factsText(['5 raiders with gear on the site', '0 waiting for approval', '0 below the rating floor']),
    ).toBe('5 raiders with gear on the site · 0 waiting for approval · 0 below the rating floor');
    expect(factsText([{ text: '1 kill' }, { text: '1 wipe' }])).toBe('1 kill · 1 wipe');
  });
});

describe('bandFacts', () => {
  it('is the place and the tier’s two counts as separate facts', () => {
    expect(bandFacts('PvP', 'us', { killed: 1, pulls: 2 })).toEqual(['PvP US', '1 boss down', '2 pulls']);
    expect(bandFacts('PvP', 'us', { killed: 3, pulls: 1 })).toEqual(['PvP US', '3 bosses down', '1 pull']);
    expect(bandFacts('Hardcore', 'eu', null)).toEqual(['Hardcore EU']);
    expect(bandFacts('PvP', 'us', { killed: 1, pulls: 2 }, 'Updated today')).toEqual([
      'PvP US',
      '1 boss down',
      '2 pulls',
      'Updated today',
    ]);
  });
});

describe('FactLine', () => {
  const FACTS = ['1 boss (Onyxia)', '1 kill', '2 trash packs', '1 wipe'];
  const { body } = render(FactLine, { props: { facts: FACTS, testid: 'line' } });

  it('reads as one sentence with " · " between facts', () => {
    expect(words(body)).toBe('1 boss (Onyxia) · 1 kill · 2 trash packs · 1 wipe');
  });

  it('makes each fact one nowrap piece that carries its own dot, the last none', () => {
    const pieces = body.match(/<span class="whitespace-nowrap">/g) ?? [];
    expect(pieces).toHaveLength(4);
    expect((body.match(new RegExp(FACT_TAIL, 'g')) ?? []).length).toBe(3);
    expect(body).not.toContain('font-mono');
  });

  it('keeps a fact’s hover, class and testid', () => {
    const { body: html } = render(FactLine, {
      props: {
        facts: [{ text: 'Item level 66', title: 'Equipped average', class: 'text-gold', testid: 'chip' }],
      },
    });
    expect(html).toContain('title="Equipped average"');
    expect(html).toContain('class="text-gold"');
    expect(html).toContain('data-testid="chip"');
    expect(asFacts(['a'])).toEqual([{ text: 'a' }]);
  });
});
