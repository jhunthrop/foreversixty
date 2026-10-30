// web/src/components/guild/GuildLoot.test.ts
// The Loot tab rendered with svelte/server (initial state, no effects): who sees Award and Undo,
// the awarded line said once, a shared item name listed once, and the honest empty candidate
// list. The ask-first confirm and the undo are clicked through in guild-centre.spec.ts.
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import { buildMockLoot } from '../../fixtures/guild/mock-guild';
import type { GuildLootPage } from '../../lib/guild/api';
import type { GuildPlanRefusal } from '../../lib/guild/officer-tools';
import { guildGateCopy } from '../../lib/guild/copy';
import GuildLoot from './GuildLoot.svelte';

const SETTINGS_HREF = '/guild/us/hardcore/the-last-watch/settings';

type LootProps = {
  loot?: GuildLootPage | null;
  officerTools?: boolean;
  raidGuildTools?: boolean;
  refusal?: GuildPlanRefusal | null;
};

function html(officer: boolean, extra: LootProps = {}): string {
  return render(GuildLoot, {
    props: {
      loot: buildMockLoot(),
      officer,
      officerTools: true,
      raidGuildTools: true,
      memberReadOnly: !officer,
      busyItemId: null,
      refusal: null,
      settingsHref: SETTINGS_HREF,
      onAward: () => {},
      onUndo: () => {},
      ...extra,
    },
  }).body;
}

/** Every opening `<button ...>` tag whose testid starts with the prefix. */
function buttons(markup: string, prefix: string): string[] {
  return [...markup.matchAll(/<button[^>]*>/g)]
    .map((m) => m[0])
    .filter((tag) => tag.includes(`data-testid="${prefix}`));
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
    const markup = html(true, { loot });
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
    expect(text(markup).split('Head of Onyxia').length - 1).toBe(1); // the icon's aria-label is not text
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

  it('puts the awarded line in the card head, no dot before Undo, and Awarded · Undo on the winner', () => {
    const markup = html(true);
    expect(markup).not.toMatch(/aria-hidden="true">·<\/span>\s*<button/);
    const awardedAt = markup.indexOf(`data-testid="guild-loot-awarded-${awarded?.item_id}"`);
    const chip = markup.slice(markup.lastIndexOf('<span', awardedAt), awardedAt);
    expect(chip).not.toContain('uppercase');
    expect(markup).toContain(`data-testid="guild-loot-row-undo-${awarded?.item_id}"`);
    expect(text(html(false))).toContain('Awarded');
    expect(html(false)).not.toContain('guild-loot-row-undo-');
  });

  describe('officer-read-only (officer tools off)', () => {
    const readOnly = (extra: LootProps = {}): string => html(true, { officerTools: false, ...extra });
    const selected = buildMockLoot().selected ?? 0;

    it('says the sentence once, with the settings link, above the boss picker', () => {
      const markup = readOnly();
      expect(markup.split('data-testid="guild-loot-plan-gate"').length - 1).toBe(1);
      expect(markup).toContain('id="officer-tools-gate-loot"');
      const gateAt = markup.indexOf('data-testid="guild-loot-plan-gate"');
      expect(gateAt).toBeLessThan(markup.indexOf('data-testid="guild-loot-boss-picker"'));
      const gate = markup.slice(gateAt, markup.indexOf('</p>', gateAt));
      expect(text(gate)).toContain(`${guildGateCopy.sentence.trim()} ${guildGateCopy.link}`);
      expect(gate).toContain(`href="${SETTINGS_HREF}"`);
    });

    it('keeps the link whole and gives it a phone-only 44px hit area outside its focus ring', () => {
      const markup = readOnly();
      const gateAt = markup.indexOf('data-testid="guild-loot-plan-gate"');
      const linkAt = markup.indexOf('<a', gateAt);
      const classes =
        (/class="([^"]*)"/.exec(markup.slice(linkAt, markup.indexOf('>', linkAt))) ?? [])[1] ?? '';
      const list = classes.split(/\s+/);
      expect(list).toEqual(
        expect.arrayContaining([
          'whitespace-nowrap',
          'relative',
          'after:absolute',
          'after:inset-x-0',
          'after:top-1/2',
          'after:h-11',
          'after:-translate-y-1/2',
          "after:content-['']",
          'md:after:hidden',
          'focus-visible:outline-2',
        ]),
      );
      expect(list).not.toContain('py-3');
    });

    it('says the sentence above the empty state too', () => {
      const markup = readOnly({ loot: null });
      const gateAt = markup.indexOf('data-testid="guild-loot-plan-gate"');
      expect(gateAt).toBeGreaterThan(-1);
      expect(gateAt).toBeLessThan(markup.indexOf('data-testid="guild-loot-empty"'));
    });

    it('shows the Read-only label', () => {
      expect(text(readOnly())).toContain(`Loot ${guildGateCopy.readOnly}`);
    });

    it('disables every Award and Undo, described by the sentence', () => {
      const markup = readOnly();
      const controls = [...buttons(markup, 'guild-loot-award-'), ...buttons(markup, 'guild-loot-row-undo-')];
      expect(buttons(markup, 'guild-loot-award-').length).toBeGreaterThan(0);
      expect(buttons(markup, 'guild-loot-row-undo-').length).toBeGreaterThan(0);
      for (const tag of controls) {
        expect(tag).toMatch(/\sdisabled(=""|\s|>)/);
        expect(tag).toContain('aria-describedby="officer-tools-gate-loot"');
        expect(tag).toMatch(/disabled:cursor-(default|not-allowed)/);
      }
      for (const tag of buttons(markup, 'guild-loot-row-undo-'))
        expect(tag).toContain('disabled:no-underline');
      expect(markup).not.toContain('guild-loot-award-confirm-');
    });

    it('says the Raid guild sentence, linking to #billing, on a Guild-tier guild', () => {
      const markup = html(true, { raidGuildTools: false });
      const gateAt = markup.indexOf('data-testid="guild-loot-plan-gate"');
      expect(gateAt).toBeGreaterThan(-1);
      const gate = markup.slice(gateAt, markup.indexOf('</p>', gateAt));
      expect(text(gate)).toContain(
        `${guildGateCopy.raidGuildSentence.trim()} ${guildGateCopy.raidGuildLink}`,
      );
      expect(gate).toContain(`href="${SETTINGS_HREF}#billing"`);
      expect(text(markup)).toContain(guildGateCopy.readOnly);
      for (const tag of buttons(markup, 'guild-loot-award-')) {
        expect(tag).toMatch(/\sdisabled(=""|\s|>)/);
        expect(tag).toContain('aria-describedby="officer-tools-gate-loot"');
      }
    });

    it("keeps today's sentence for a lapsed guild, even if Raid guild is also off", () => {
      const markup = html(true, { officerTools: false, raidGuildTools: false });
      expect(text(markup)).toContain(guildGateCopy.link);
      expect(text(markup)).not.toContain(guildGateCopy.raidGuildLink);
    });

    it('shows a member nothing of the Raid guild gate', () => {
      expect(html(false, { raidGuildTools: false })).not.toContain('guild-loot-plan-gate');
    });

    it('keeps the officer view free of the gate while tools are on', () => {
      const markup = html(true);
      expect(markup).not.toContain('guild-loot-plan-gate');
      expect(markup).not.toContain('aria-describedby="officer-tools-gate-loot"');
      expect(text(markup)).not.toContain(guildGateCopy.readOnly);
    });

    it('leaves a moderator as before: the page passes officerTools true, so no sentence and live Award', () => {
      const markup = html(true, { officerTools: true });
      expect(markup).not.toContain('guild-loot-plan-gate');
      expect(text(markup)).not.toContain(guildGateCopy.readOnly);
      for (const tag of buttons(markup, 'guild-loot-award-')) expect(tag).not.toMatch(/ disabled(=| |>)/);
    });

    it('leaves a member exactly as before: no sentence, no controls', () => {
      const markup = html(false, { officerTools: false });
      expect(markup).not.toContain('guild-loot-plan-gate');
      expect(markup).not.toContain('guild-loot-award-');
      expect(markup).not.toContain('guild-loot-row-undo-');
    });

    it('puts one refused line inside the refused item card, for an award', () => {
      const item = buildMockLoot().items.find((i) => i.awarded_to === null && i.candidates.length > 0);
      expect(item).toBeDefined();
      const id = item?.item_id ?? 0;
      const markup = readOnly({ refusal: { encounterId: selected, itemId: id, kind: 'award' } });
      const testid = `data-testid="guild-loot-plan-refused-${id}"`;
      expect(markup.split(testid).length - 1).toBe(1);
      expect(markup.split('guild-loot-plan-refused-').length - 1).toBe(1);
      const cardAt = markup.indexOf(`data-testid="guild-loot-item-${id}"`);
      const nextCard = markup.indexOf('data-testid="guild-loot-item-', cardAt + 10);
      const lineAt = markup.indexOf(testid);
      expect(lineAt).toBeGreaterThan(cardAt);
      if (nextCard > -1) expect(lineAt).toBeLessThan(nextCard);
      const tag = markup.slice(markup.lastIndexOf('<p', lineAt), markup.indexOf('>', lineAt) + 1);
      expect(tag).toContain('role="alert"');
      expect(tag).toContain('tabindex="-1"');
      expect(tag).toContain('text-[13px]');
      expect(tag).toContain('leading-5');
      expect(tag).toContain('text-wipe');
      expect(text(markup.slice(lineAt, markup.indexOf('</p>', lineAt)))).toContain(
        guildGateCopy.awardRefused,
      );
    });

    it('renders no refused line when another boss is selected', () => {
      const item = buildMockLoot().items.find((i) => i.awarded_to === null && i.candidates.length > 0);
      const markup = readOnly({
        refusal: { encounterId: selected + 1000, itemId: item?.item_id ?? 0, kind: 'award' },
      });
      expect(markup).not.toContain('guild-loot-plan-refused-');
      expect(text(markup)).not.toContain(guildGateCopy.awardRefused);
    });

    it('says the undo refusal for an undo', () => {
      const id = awarded?.item_id ?? 0;
      const markup = readOnly({ refusal: { encounterId: selected, itemId: id, kind: 'undo' } });
      expect(text(markup)).toContain(guildGateCopy.undoRefused);
      expect(text(markup)).not.toContain(guildGateCopy.awardRefused);
    });
  });
});
