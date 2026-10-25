import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import { characterHistoryFixture } from '../../fixtures/character-history';
import { parseCharacterPath, type CharacterPath } from '../../lib/characters';
import type { CharacterHistory } from '../../lib/character/history-api';
import type { HistoryController, HistoryStatus } from '../../lib/character/history-controller.svelte';
import CharacterHistorySection from './CharacterHistory.svelte';

function fixturePath(): CharacterPath {
  const parsed = parseCharacterPath('/character/us/hardcore/elyra-duskvale');
  if (parsed === null) throw new Error('fixture path');
  return parsed;
}
const path = fixturePath();

const stub = (status: HistoryStatus, history: CharacterHistory | null): HistoryController => ({
  status,
  history,
  load: () => undefined,
});

function section(controller: HistoryController): string {
  return render(CharacterHistorySection, {
    props: { path, controller, sizeOverride: { width: 588, height: 294, phone: false } },
  }).body.replace(/<!--[^>]*-->/g, '');
}

describe('CharacterHistory', () => {
  it('is the #history section with the heading, the caption and one panel, in both columns when ready', () => {
    const html = section(stub('ready', characterHistoryFixture()));
    expect(html).toMatch(/<section[^>]*id="history"/);
    expect(html).toContain('-order-2');
    expect(html).toContain('>History</h2>');
    expect(html).toContain('Includes your private and guild reports. Only you see this section.');
    expect(html).toContain('xl:grid-cols-2');
    expect(html).toContain('data-testid="history-score"');
    expect(html).toContain('data-testid="history-parse"');
    expect(html).not.toContain('history-reserve');
  });

  it('shows the skeleton at the reserve while loading, busy and named for a screen reader', () => {
    const html = section(stub('loading', null));
    expect(html).toContain('history-reserve');
    expect(html).toContain('aria-busy="true"');
    expect(html).toContain('Reading your history…');
    expect(html).toContain('data-testid="history-skeleton"');
    expect(html).not.toContain('history-plot');
  });

  it('shows the load error with Try again inside the panel', () => {
    const html = section(stub('error', null));
    expect(html).toContain('Your history did not load.');
    expect(html).toContain('data-testid="history-error-retry"');
    expect(html).toContain('>History</h2>');
  });

  it('replaces both columns with one sentence when no score point exists', () => {
    const html = section(stub('ready', characterHistoryFixture({ score: [] })));
    expect(html).toContain('No history yet. Each rated fight you upload adds a point.');
    expect(html).not.toContain('data-testid="history-score"');
    expect(html).not.toContain('data-testid="history-parse"');
  });

  it('starts as the loading skeleton before any read when it makes its own controller', () => {
    const html = render(CharacterHistorySection, { props: { path } }).body;
    expect(html).toContain('data-testid="history-skeleton"');
  });
});
