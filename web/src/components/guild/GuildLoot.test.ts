// web/src/components/guild/GuildLoot.test.ts
// The Loot tab rendered with svelte/server (initial state, no effects): who sees Award and Undo,
// the awarded line said once, a shared item name listed once, and the honest empty candidate
// list. The ask-first confirm and the undo are clicked through in guild-centre.spec.ts.
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import { buildMockLoot } from '../../fixtures/guild/mock-guild';
import GuildLoot from './GuildLoot.svelte';

function html(officer: boolean): string {
  return render(GuildLoot, {
    props: {
      loot: buildMockLoot(),
      officer,
      memberReadOnly: !officer,
      busyItemId: null,
      onAward: () => {},
      onUndo: () => {},
    },
  }).body;
}

function text(markup: string): string {
  return markup
    .replace(/<!--[\s\S]*?-->/g, '')
    .replace(/<[^>]+>/g, ' ')
    .replace(/&nbsp;/g, ' ')
    .replace(/\s+/g, ' ')
    .trim();
}

describe('GuildLoot', () => {
  const awarded = buildMockLoot().items.find((item) => item.awarded_to !== null);

  it('offers the officer Undo beside the awarded line, and a member nothing', () => {
    expect(awarded).toBeDefined();
    const id = awarded?.item_id;
    expect(html(true)).toContain(`data-testid="guild-loot-undo-${id}"`);
    expect(html(false)).not.toContain('guild-loot-undo-');
    expect(html(false)).not.toContain('guild-loot-award-');
  });

  it('says "Already awarded" on the other candidates, without a live Award button', () => {
    const markup = html(true);
    const start = markup.indexOf(`data-testid="guild-loot-item-${awarded?.item_id}"`);
    const panel = markup.slice(start, markup.indexOf('data-testid="guild-loot-item-', start + 10));
    expect(panel).not.toContain('data-testid="guild-loot-award-');
    expect(text(panel)).toContain('Awarded to');
    expect(text(panel)).toContain('Undo');
    const others = (awarded?.candidates.length ?? 0) - 1;
    expect(text(panel).split('Already awarded').length - 1).toBe(others);
  });

  it('lists "Head of Onyxia" once, for both copies', () => {
    const markup = html(true);
    expect(markup.split('Head of Onyxia').length - 1).toBe(1);
    expect(text(markup)).toContain('Head of Onyxia x2');
  });

  it('shows no Award confirm until Award is pressed', () => {
    expect(html(true)).not.toContain('guild-loot-award-confirm-');
  });

  it('never offers a paladin helm to a warrior', () => {
    const markup = html(true);
    const crown = markup.slice(markup.indexOf('Judgement Crown'));
    const panel = crown.slice(0, crown.indexOf('data-testid="guild-loot-item-', 10));
    expect(panel).toContain('No one on the roster can use this.');
    expect(panel).not.toContain('data-testid="guild-loot-award-');
  });
});
