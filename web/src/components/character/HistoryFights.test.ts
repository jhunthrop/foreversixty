import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import { characterHistoryFixture } from '../../fixtures/character-history';
import type { HistoryFight } from '../../lib/character/history-api';
import HistoryFights from './HistoryFights.svelte';

function fights(count: number): HistoryFight[] {
  const [template] = characterHistoryFixture().fights;
  return Array.from({ length: count }, (_, index) => ({
    ...template,
    report_id: `rpt${String(index).padStart(9, '0')}`,
    visibility: index === 1 ? 'private' : index === 2 ? 'guild' : 'public',
  }));
}

function list(count: number, marked: string[] = [], extra: Record<string, unknown> = {}): string {
  return render(HistoryFights, {
    props: {
      fights: fights(count),
      label: 'Fights behind the score',
      kind: 'fights',
      marked: new Set(marked),
      testid: 'history-score-list',
      ...extra,
    },
  }).body.replace(/<!--[^>]*-->/g, '');
}

const rowCount = (html: string): number =>
  (html.match(/<li class="border-line-soft grid min-h-11/g) ?? []).length;

describe('HistoryFights', () => {
  it('shows eight rows and Show all N when more exist', () => {
    const html = list(31);
    expect(rowCount(html)).toBe(8);
    expect(html).toContain('Show all 31 fights');
    expect(html).toContain('Fights behind the score');
    expect(html).not.toContain('<h2');
  });

  it('has no button at eight rows or fewer', () => {
    const html = list(8);
    expect(rowCount(html)).toBe(8);
    expect(html).not.toContain('history-score-list-toggle');
  });

  it('names kills for the parse list', () => {
    expect(list(12, [], { kind: 'kills', label: 'Kills of Golemagg' })).toContain('Show all 12 kills');
  });

  it('caps Show all at a hundred', () => {
    expect(list(250)).toContain('Show all 100 fights');
  });

  it('marks the selected point s rows', () => {
    const html = list(5, ['rpt000000000:7']);
    expect(html.match(/data-marked="true"/g)).toHaveLength(1);
  });

  it('writes Private and Guild after the boss name, and nothing for a public report', () => {
    const html = list(5);
    expect(html.match(/data-testid="character-fight-note">Private</g)).toHaveLength(1);
    expect(html.match(/data-testid="character-fight-note">Guild</g)).toHaveLength(1);
    expect(html.match(/character-fight-note/g)).toHaveLength(2);
  });

  it('links each boss to its report fight', () => {
    expect(list(1)).toContain('href="/reports/rpt000000000?fight=7"');
  });

  it('leads with the date for a kills list', () => {
    const html = list(2, [], { lead: 'date', kind: 'kills' });
    expect(html).not.toContain('Golemagg the Incinerator');
    expect(html).toContain('11 Dec 2026');
  });
});
