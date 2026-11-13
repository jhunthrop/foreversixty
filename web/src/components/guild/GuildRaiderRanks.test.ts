// web/src/components/guild/GuildRaiderRanks.test.ts
// One static render per state of the raider ranks row (vitest hands components Svelte's server
// entry, so the ticks and the save are Playwright's: tests/e2e/guild-settings-raider-ranks.spec.ts).
import type { ComponentProps } from 'svelte';
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import { raiderRanksCopy } from '../../lib/guild/raider-ranks';
import type { SetupRankLine } from '../../lib/guild/setup-api';
import { busyCopy } from '../../lib/ui/copy';
import GuildRaiderRanks from './GuildRaiderRanks.svelte';

const LINES: SetupRankLine[] = [
  { rank_index: 0, members: 1, at_sixty: 1, names: ['Harrowgate'] },
  { rank_index: 1, members: 2, at_sixty: 2, names: ['Kestrelle', 'Dunmarrow'] },
  { rank_index: 2, members: 6, at_sixty: 5, names: ['Bramblek', 'Ostavine', 'Pellwyn'] },
  { rank_index: 3, members: 9, at_sixty: 4, names: ['Quillmere', 'Tarnwick', 'Velloran'] },
];

type Props = ComponentProps<typeof GuildRaiderRanks>;

function html(overrides: Partial<Props> = {}): string {
  return render(GuildRaiderRanks, {
    props: {
      lines: LINES,
      officerMax: 1,
      saved: null,
      busy: false,
      disabled: false,
      justSaved: false,
      onSave: () => {},
      onEdit: () => {},
      ...overrides,
    },
  }).body.replace(/<!--[\s\S]*?-->/g, '');
}

function checkbox(body: string, rank: number): string {
  const match = body.match(new RegExp(`<input[^>]*data-testid="guild-raider-rank-${rank}-check"[^>]*>`));
  if (match === null) throw new Error(`no checkbox for rank ${rank}`);
  return match[0];
}

describe('GuildRaiderRanks', () => {
  it('lists every rank highest first with its holders, nothing ticked until set', () => {
    const body = html();
    expect(body).toContain(raiderRanksCopy.unset);
    const order = [0, 1, 2, 3].map((rank) => body.indexOf(`data-testid="guild-raider-rank-${rank}"`));
    expect(order).toEqual([...order].sort((a, b) => a - b));
    expect(body).toContain('Guild master');
    expect(body).toContain(raiderRanksCopy.explain);
    expect(body).toContain('5 of 6 at 60 · Bramblek, Ostavine, Pellwyn\u00a0+3');
    for (const rank of [0, 1, 2, 3]) expect(checkbox(body, rank)).not.toContain('checked');
    expect(body).toMatch(
      /data-testid="guild-raider-ranks-save"[^>]*disabled|disabled[^>]*data-testid="guild-raider-ranks-save"/,
    );
  });

  it('preselects the saved ranks and sums what they cover', () => {
    const body = html({ saved: [2, 3] });
    expect(body).toContain('15 characters, 9 at 60');
    expect(checkbox(body, 2)).toContain('checked');
    expect(checkbox(body, 3)).toContain('checked');
    expect(checkbox(body, 1)).not.toContain('checked');
  });

  it('shows Saved. after its own save', () => {
    const body = html({ saved: [2], justSaved: true });
    expect(body).toContain('data-testid="guild-raider-ranks-saved"');
    expect(body).toContain(raiderRanksCopy.saved);
  });

  it('locks every box while a request is in flight or the claim is frozen', () => {
    for (const state of [{ busy: true }, { disabled: true }]) {
      const body = html({ saved: [2], ...state });
      for (const rank of [0, 1, 2, 3]) expect(checkbox(body, rank)).toContain('disabled');
    }
    expect(html({ busy: true })).not.toContain(busyCopy.saving);
  });

  it('reads loading, failed and an empty roster', () => {
    expect(html({ lines: null })).toContain('data-testid="settings-row-loading"');
    expect(html({ lines: 'failed' })).toContain(raiderRanksCopy.ranksFailed);
    expect(html({ lines: [] })).toContain(raiderRanksCopy.noRanks);
  });
});
