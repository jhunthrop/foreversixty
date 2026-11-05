// web/src/components/guild/GuildLoot.test.ts
// The Loot tab rendered with svelte/server (initial state, no effects): who sees Award and Undo,
// the awarded line said once, the game's slot words, the two tiers, attendance hidden with no
// counted night, and the drops nobody ranks collapsed to one line. The ask-first confirm and the undo are clicked through in guild-centre.spec.ts.
import { gateFactsFixture } from '../../test-support/gate-facts';
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
  ownKeys?: readonly string[];
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
      gateFacts: gateFactsFixture({
        guildTier: (extra.officerTools ?? true) && !(extra.raidGuildTools ?? true),
        settingsHref: SETTINGS_HREF,
      }),
      onAward: () => {},
      onUndo: () => {},
      onBoss: () => {},
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

  it("prints the game's slot words, never the enum", () => {
    const markup = text(html(true));
    expect(markup).toContain('Head · Leather');
    expect(markup).toContain('One-Hand · Sword');
    expect(markup).toContain('Held In Off-hand');
    expect(markup).not.toMatch(/\b(off_hand|main_hand|OFF_HAND|MAIN_HAND|finger|trinket)\b/);
  });

  it('prints the gain of a tier-1 row with an empty slot and says Empty slot', () => {
    const loot = buildMockLoot();
    const item = loot.items.find((i) => i.candidates.length > 0 && i.awarded_to === null);
    expect(item).toBeDefined();
    if (item === undefined) return;
    item.candidates[0] = { ...item.candidates[0], replaces: '', replaces_item: null, ilvl_delta: 76 };
    const markup = html(true, { loot });
    const panel = markup.slice(markup.indexOf(`data-testid="guild-loot-item-${item.item_id}"`));
    expect(text(panel.slice(0, panel.indexOf('guild-loot-candidate-gain') + 80))).toContain('Empty slot');
    const figure = /guild-loot-candidate-gain[^>]*>\s*([^<]*)</.exec(panel);
    expect(figure?.[1].trim().replace(/\u00a0/g, ' ')).toBe('+76 item levels');
  });

  it('labels the first tier-1 row once, and never above a tier-0 row', () => {
    const loot = buildMockLoot();
    const item = loot.items.find((i) => i.candidates.length > 2 && i.awarded_to === null);
    expect(item).toBeDefined();
    if (item === undefined) return;
    item.candidates[0] = {
      ...item.candidates[0],
      tier: 0,
      gain_dps: 12.4,
      gain_unit: 'dps',
      ilvl_delta: null,
    };
    const markup = html(true, { loot });
    const panel = markup.slice(markup.indexOf(`data-testid="guild-loot-item-${item.item_id}"`));
    const card = panel.slice(0, panel.indexOf('data-testid="guild-loot-item-', 10));
    expect(card.split(`guild-loot-tier-one-${item.item_id}`).length - 1).toBe(1);
    const label = card.indexOf(`guild-loot-tier-one-${item.item_id}`);
    expect(card.indexOf('+12.4')).toBeLessThan(label);
    expect(card.slice(label)).toContain('item');
    expect(text(card)).toContain('Ranked by item level only. Not on their best-in-slot list.');
  });

  it('shows five candidates and then Show {n} more', () => {
    const loot = buildMockLoot();
    const item = loot.items.find((i) => i.candidates.length > 0 && i.awarded_to === null);
    expect(item).toBeDefined();
    if (item === undefined) return;
    const first = item.candidates[0];
    item.candidates = Array.from({ length: 7 }, (_, n) => ({
      ...first,
      character_key: `us/hardcore/row${n}`,
      name: `Row${n}`,
    }));
    const markup = html(true, { loot });
    expect(markup).toContain('guild-loot-award-' + item.item_id + '-us/hardcore/row4');
    expect(markup).not.toContain('guild-loot-award-' + item.item_id + '-us/hardcore/row5');
    expect(text(markup)).toContain('Show 2 more');
  });

  it("washes the viewer's row and says (you)", () => {
    const loot = buildMockLoot();
    const item = loot.items.find((i) => i.candidates.length > 0 && i.awarded_to === null);
    expect(item).toBeDefined();
    if (item === undefined) return;
    const key = item.candidates[0].character_key;
    const markup = html(false, { loot, ownKeys: [key] });
    expect(text(markup)).toContain(`${item.candidates[0].name} (you)`);
    // The space is rendered text, not markup whitespace Svelte trims: `Ashenveil (you)`, as on Readiness.
    expect(markup).toMatch(/>\s*\(you\)</);
    expect(markup).toMatch(/<span[^>]*> \(you\)<\/span>/);
    expect(markup).toContain('row-band-own');
    expect(html(false, { loot })).not.toContain('row-band-own');
  });

  it('hides attendance with no counted night, and shows it otherwise', () => {
    expect(html(true)).toContain('guild-loot-candidate-attendance');
    const loot = buildMockLoot();
    for (const item of loot.items) {
      for (const candidate of item.candidates) candidate.attendance = { present: 0, nights: 0 };
    }
    const markup = html(true, { loot });
    expect(markup).not.toContain('guild-loot-candidate-attendance');
    expect(text(markup)).not.toContain('raid night');
  });

  it('says what to expect while nobody is 60, and collapses every drop', () => {
    const loot = buildMockLoot();
    loot.at_level_cap = 0;
    for (const item of loot.items) {
      item.candidates = [];
      item.eligible = [];
      item.awarded_to = null;
    }
    const markup = html(true, { loot });
    expect(text(markup)).toContain(
      "Onyxia's drops need level 60. Candidates show here as your raiders reach it.",
    );
    expect(markup).not.toContain('guild-loot-items');
    expect(text(markup)).toContain(`${loot.items.length} more drops nobody here ranks for`);
    expect(markup).not.toContain('guild-loot-award-');
  });

  it('says so when someone is 60 and nobody can use a drop', () => {
    const loot = buildMockLoot();
    for (const item of loot.items) {
      item.candidates = [];
      item.awarded_to = null;
    }
    expect(text(html(true, { loot }))).toContain('Nobody at 60 can use these drops yet.');
  });

  describe('the collapsed line', () => {
    it('never draws a card for a quest drop, the bag or the scale, or for a drop nobody can use', () => {
      const markup = html(true);
      for (const id of [15410, 17966, 18422, 18705, 21108, 16955]) {
        expect(markup).not.toContain(`data-testid="guild-loot-item-${id}"`);
        expect(markup).toContain(`data-testid="guild-loot-unranked-${id}"`);
      }
    });

    it('labels each drop', () => {
      const markup = text(html(true));
      expect(markup).toContain('Quest: Victory for the Alliance / Victory for the Horde');
      expect(markup).toContain('Quest: A Proper String (Hunter)');
      expect(markup).toContain('Quest: The Only Prescription');
      expect(markup).toContain('Bag. Anyone can use it.');
      expect(markup).toContain('Skinned from Onyxia. Not a drop.');
      expect(markup).toContain('Nobody here can use it yet.');
      expect(markup).not.toContain('No one on the roster can use this.');
      expect(markup).not.toContain('The simulator has not checked these');
    });

    it('offers an officer Award on a quest drop and the bag, never on the scale', () => {
      const markup = html(true);
      expect(buttons(markup, 'guild-loot-award-18705')).toHaveLength(1);
      expect(buttons(markup, 'guild-loot-award-17966')).toHaveLength(1);
      expect(buttons(markup, 'guild-loot-award-15410')).toHaveLength(0);
      expect(buttons(html(false), 'guild-loot-award-18705')).toHaveLength(0);
    });

    it('disables that Award for an officer with the tools off, described by the sentence', () => {
      const [tag] = buttons(html(true, { officerTools: false }), 'guild-loot-award-18705');
      expect(tag).toMatch(/\sdisabled(=""|\s|>)/);
      expect(tag).toContain('aria-describedby="officer-tools-gate-loot"');
    });

    it('prints the award on an awarded quest drop, with Undo for the officer', () => {
      const loot = buildMockLoot();
      const sinew = loot.items.find((i) => i.item_id === 18705);
      if (sinew === undefined) throw new Error('fixture lost the Sinew');
      sinew.awarded_to = {
        award_id: 9,
        character_key: 'us/hardcore/x',
        name: 'Bowman',
        at: '2026-12-15T20:10:00Z',
        by_name: 'Kraggor',
      };
      const markup = html(true, { loot });
      expect(text(markup)).toContain('Awarded to Bowman');
      expect(markup).toContain('guild-loot-row-undo-18705');
      expect(buttons(markup, 'guild-loot-award-18705')).toHaveLength(0);
    });

    it('is absent when every drop has candidates', () => {
      const loot = buildMockLoot();
      loot.items = loot.items.filter((i) => i.kind === 'gear' && i.candidates.length > 0);
      expect(html(true, { loot })).not.toContain('guild-loot-unranked');
    });
  });

  it('puts the one footnote at the bottom of the tab', () => {
    const markup = html(true);
    expect(markup.indexOf('guild-loot-footnote')).toBeGreaterThan(markup.indexOf('guild-loot-unranked'));
    expect(text(markup).split('Candidates are characters who can wear the item').length - 1).toBe(1);
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
    expect(text(html(true)).split('Head of Onyxia').length - 1).toBe(1); // the icon's aria-label is not text
  });

  it('shows no Award confirm until Award is pressed', () => {
    expect(html(true)).not.toContain('guild-loot-award-confirm-');
  });

  it('never offers a paladin helm to a warrior', () => {
    const markup = html(true);
    const row = markup.slice(markup.indexOf('data-testid="guild-loot-unranked-16955"'));
    expect(row).toContain('Nobody here can use it yet.');
    expect(markup).not.toContain('data-testid="guild-loot-item-16955"');
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
      expect(text(gate)).toContain(`${guildGateCopy.lead.loot} ${guildGateCopy.startTrial}`);
      expect(gate).toContain(`href="${SETTINGS_HREF}?plan=raid_guild#billing"`);
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

    it('shows no Read-only label: the gate sentence covers an officer (loot spec §3.1)', () => {
      expect(text(readOnly())).not.toContain(guildGateCopy.readOnly);
    });

    it('shows the Read-only label to a member', () => {
      expect(text(html(false))).toContain(`Loot ${guildGateCopy.readOnly}`);
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
      expect(text(gate)).toContain(`${guildGateCopy.lead.loot} ${guildGateCopy.moveToRaidGuild}`);
      expect(gate).toContain(`href="${SETTINGS_HREF}?upgrade=1#billing"`);
      expect(text(markup)).not.toContain(guildGateCopy.readOnly);
      for (const tag of buttons(markup, 'guild-loot-award-')) {
        expect(tag).toMatch(/\sdisabled(=""|\s|>)/);
        expect(tag).toContain('aria-describedby="officer-tools-gate-loot"');
      }
    });

    it('offers no Move to a guild with no plan at all', () => {
      const markup = html(true, { officerTools: false, raidGuildTools: false });
      expect(text(markup)).toContain(guildGateCopy.startTrial);
      expect(text(markup)).not.toContain(guildGateCopy.moveToRaidGuild);
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
