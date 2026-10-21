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

  it('offers the officer one Undo, on the winner row, and a member nothing', () => {
    expect(awarded).toBeDefined();
    const id = awarded?.item_id;
    // The awarded chip carries no Undo of its own: two links for one action read as two actions.
    expect(html(true)).not.toContain('guild-loot-undo-');
    expect(html(true).split(`data-testid="guild-loot-row-undo-${id}"`).length - 1).toBe(1);
    expect(html(false)).not.toContain('guild-loot-row-undo-');
    expect(html(false)).not.toContain('guild-loot-award-');
  });

  it('prints no gain rider for an empty slot, and keeps the rider as one non-breaking chunk otherwise', () => {
    const loot = buildMockLoot();
    const item = loot.items.find((i) => i.candidates.length > 0 && i.awarded_to === null);
    expect(item).toBeDefined();
    if (item === undefined) return;
    item.candidates[0] = { ...item.candidates[0], replaces: '', ilvl_delta: 76, gain_dps: null };
    item.candidates[1] = {
      ...item.candidates[1],
      replaces: 'Beastmaster Cap',
      ilvl_delta: 16,
      gain_dps: null,
    };
    const markup = render(GuildLoot, {
      props: {
        loot,
        officer: true,
        memberReadOnly: false,
        busyItemId: null,
        onAward: () => {},
        onUndo: () => {},
      },
    }).body;
    const panel = markup.slice(markup.indexOf(`data-testid="guild-loot-item-${item.item_id}"`));
    const firstLine = panel.slice(panel.indexOf('guild-loot-candidate-line'));
    const firstRow = firstLine.slice(0, firstLine.indexOf('Replaces'));
    expect(text(firstRow)).toContain('Empty slot');
    expect(text(firstRow)).not.toContain('item levels');
    // Exactly one rider on the item: the second candidate's, whole.
    const riders = [...panel.matchAll(/guild-loot-candidate-gain[^>]*>([^<]*)</g)].map((m) => m[1]);
    expect(riders[0]).toBe('+16 item levels');
  });

  it('says "Already awarded" on the other candidates, without a live Award button', () => {
    const markup = html(true);
    const start = markup.indexOf(`data-testid="guild-loot-item-${awarded?.item_id}"`);
    const panel = markup.slice(start, markup.indexOf('data-testid="guild-loot-item-', start + 10));
    expect(panel).not.toContain('data-testid="guild-loot-award-');
    expect(text(panel)).toContain('Awarded to');
    expect(text(panel)).toContain('Awarded · Undo');
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

  it('reads each candidate as what they replace, their spec and their raid nights', () => {
    const first = text(html(true).slice(html(true).indexOf('data-testid="guild-loot-candidate-line"')));
    expect(first).toMatch(/^[^]*Replaces [A-Z][^·]+ · [A-Za-z ]+ · \d+ of 8 raid nights/);
    expect(text(html(true))).not.toContain('over their current');
  });

  it('puts the awarded line on its own full-width row, no dot before Undo, and Awarded · Undo on the winner', () => {
    const markup = html(true);
    expect(markup).not.toMatch(/aria-hidden="true">·<\/span>\s*<button/);
    const awardedAt = markup.indexOf(`data-testid="guild-loot-awarded-${awarded?.item_id}"`);
    const chip = markup.slice(markup.lastIndexOf('<span', awardedAt), awardedAt);
    expect(chip).toContain('basis-full');
    expect(chip).not.toContain('uppercase');
    expect(markup).toContain(`data-testid="guild-loot-row-undo-${awarded?.item_id}"`);
    expect(text(html(false))).toContain('Awarded');
    expect(html(false)).not.toContain('guild-loot-row-undo-');
  });
});
