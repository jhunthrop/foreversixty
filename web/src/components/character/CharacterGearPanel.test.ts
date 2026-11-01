// web/src/components/character/CharacterGearPanel.test.ts
// Task L8.6 and its fix round 1: the profile's Gear section, rendered with svelte/server (no
// effects run). `preloaded` stands in for the build's items file having loaded (or failed);
// without it the section is still waiting for that file. The e2e spec character-gear.spec.ts
// checks the same in a browser.
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import { GEAR_FIXTURE, GEAR_FIXTURE_ITEMS } from '../../fixtures/character-gear';
import { gearCopy, notReadCopy } from '../../lib/gear/copy';
import type { GearItemTable } from '../../lib/gear/qualities';
import CharacterGearPanel from './CharacterGearPanel.svelte';

const FACTS: GearItemTable = new Map(
  GEAR_FIXTURE_ITEMS.map((item) => [
    item.id,
    { quality: item.quality, icon: item.icon, twoHand: item.two_hand === true },
  ]),
);

const PROPS = {
  identity: {
    name: 'Sunwick',
    classSlug: 'priest',
    race: 'Human',
    spec: 'Discipline',
    level: 60,
    faction: 'alliance' as const,
    source: 'Addon export',
  },
  gear: GEAR_FIXTURE,
  total: 70,
  classSlug: 'priest',
  bisHref: '/bis/priest/discipline?faction=alliance#band-alliance-60',
  spec: 'Discipline',
  band: 60,
  preloaded: FACTS,
};

function text(html: string): string {
  return html
    .replace(/<!--[\s\S]*?-->/g, '')
    .replace(/<[^>]+>/g, '')
    .replace(/&#39;/g, "'")
    .replace(/&lt;/g, '<')
    .replace(/\s+/g, ' ')
    .trim();
}

/** The markup without Svelte's hydration comments. */
function bare(html: string): string {
  return html.replace(/<!--[\s\S]*?-->/g, '');
}

function classesOf(html: string, testid: string): string[] {
  return new RegExp(`<[a-z]+ class="([^"]*)"[^>]*data-testid="${testid}"`).exec(html)?.[1].split(' ') ?? [];
}

describe('CharacterGearPanel as the paperdoll (paperdoll spec call site 2)', () => {
  it('is the card the Gear link opens, landing 96px under the top', () => {
    const { body } = render(CharacterGearPanel, { props: PROPS });
    const anchor = /<div class="([^"]*)" id="gear"/.exec(body);
    expect(anchor?.[1].split(' ')).toEqual(expect.arrayContaining(['scroll-mt-24', '-order-2']));
    expect(body).toContain('data-testid="character-gear"');
  });

  it('heads the card "Gear" with the gain total in gold, the compact head (§3.2)', () => {
    const { body } = render(CharacterGearPanel, { props: PROPS });
    expect(text(/<h2[^>]*data-testid="character-gear-title"[\s\S]*?<\/h2>/.exec(body)?.[0] ?? '')).toBe(
      'Gear',
    );
    expect(classesOf(body, 'character-gear-summary')).toEqual(
      expect.arrayContaining(['text-gold', 'font-semibold']),
    );
    expect(text(/data-testid="character-gear-summary">[^<]*/.exec(body)?.[0] ?? '')).toContain(
      'data-testid="character-gear-summary">+70 points across 7 slots',
    );
  });

  it('draws every slot in the frame, the worn names, the picks with their gains and "Best in slot"', () => {
    const body = bare(render(CharacterGearPanel, { props: PROPS }).body);
    expect(body).toContain('data-testid="character-gear-paperdoll"');
    for (const slot of ['head', 'neck', 'wrist', 'trinket2', 'main_hand', 'off_hand', 'ranged'])
      expect(body).toContain(`data-testid="paperdoll-slot-${slot}"`);
    expect(body).toContain('aria-label="Head: Green Lens"');
    expect(body).toContain('aria-label="Head upgrade: Halo of Transcendence, 18.4 points to gain"');
    expect(text(body)).toContain('+18.4');
    expect(body).toContain('data-testid="paperdoll-best"');
  });

  it('never offers the worn item as a gain: the rings worn the other way round read "Best in slot"', () => {
    const body = bare(render(CharacterGearPanel, { props: PROPS }).body);
    const ring =
      /data-testid="paperdoll-slot-finger1"[\s\S]*?(?=data-testid="paperdoll-(?:slot|up)-)/.exec(body)?.[0] ??
      '';
    expect(ring).toContain('paperdoll-best');
    expect(ring).not.toContain('paperdoll-gain');
  });

  it('draws an item our data lacks with its slot word and "Not scored", and hides the item level', () => {
    const body = bare(render(CharacterGearPanel, { props: PROPS }).body);
    expect(text(body)).toContain('Legs · Item');
    expect(body).toContain('data-testid="paperdoll-unscored"');
    expect(body).not.toContain('data-testid="paperdoll-ilvl"');
    expect(body).toContain('data-testid="paperdoll-note"');
  });

  it('shows the item level once every worn item is known', () => {
    const known = GEAR_FIXTURE.map((slot) =>
      slot.worn !== null && slot.worn.name === null ? { ...slot, worn: null } : slot,
    );
    const body = bare(render(CharacterGearPanel, { props: { ...PROPS, gear: known } }).body);
    expect(body).toContain('data-testid="paperdoll-ilvl"');
    expect(body).not.toContain('data-testid="paperdoll-note"');
  });

  it('marks the empty off hand under the worn two-hander "Two-hand in main hand"', () => {
    const body = bare(render(CharacterGearPanel, { props: PROPS }).body);
    expect(body).toMatch(/data-testid="paperdoll-slot-off_hand"[^>]*data-state="two-hand"/);
    expect(text(body)).toContain(gearCopy.paperdoll.twoHand);
  });

  it('draws the two-hander greyed in the off-hand slot, as the character pane does, not an empty socket', () => {
    const body = bare(render(CharacterGearPanel, { props: PROPS }).body);
    const mainHandId = GEAR_FIXTURE.find((slot) => slot.slot === 'main_hand')!.worn!.id;
    const icon = FACTS.get(mainHandId)!.icon;
    const off = body.slice(body.indexOf('data-testid="paperdoll-slot-off_hand"'));
    const cell = off.slice(0, off.indexOf('data-testid="paperdoll-slot-', 10));
    // paperdoll.css greys the socket of a data-state="two-hand" cell.
    expect(cell).toContain('data-state="two-hand"');
    expect(cell).toContain(`/icons/${icon}.webp`);
    expect(cell).not.toContain('/icons/hd/slots/');
  });

  it('mid-load: the frame shimmers and the BiS link holds its place', () => {
    const body = bare(render(CharacterGearPanel, { props: { ...PROPS, preloaded: undefined } }).body);
    expect(body).toMatch(/data-testid="paperdoll-slot-head"[^>]*data-state="loading"/);
    expect(classesOf(body, 'character-gear-bis')).toContain('invisible');
  });

  it('a failed items file: names in the body colour and the note says so once', () => {
    const body = bare(render(CharacterGearPanel, { props: { ...PROPS, preloaded: null } }).body);
    expect(text(body)).toContain(gearCopy.paperdoll.noteFailed);
    expect(body).toContain('pd-plain');
  });

  it('keeps the BiS link in the card foot', () => {
    const body = bare(render(CharacterGearPanel, { props: PROPS }).body);
    expect(body).toMatch(/data-testid="character-gear-footer"[\s\S]*Full Discipline BiS list at 60/);
  });

  it('`gear: []`: the "Gear" heading over one line that keeps the anchor', () => {
    const { body } = render(CharacterGearPanel, { props: { ...PROPS, gear: [], total: null } });
    expect(body).not.toContain('character-gear-row-');
    const heading = /<h2[^>]*tabindex="-1"[^>]*data-testid="character-gear-title"[^>]*>([\s\S]*?)<\/h2>/.exec(
      body,
    );
    expect(text(heading?.[1] ?? '')).toBe('Gear');
    const line = /<p class="([^"]*)" id="gear" data-testid="character-gear-none">([\s\S]*?)<\/p>/.exec(body);
    // The default viewer owns the character: sync, and nothing else (W.10).
    expect(text(line?.[2] ?? '')).toBe('No gear on file yet. Sync the addon.');
    expect(body).not.toContain('character-gear-planner');
  });

  it('`gear: []`, a visitor: the one quiet line says sync or sign in; with a score the planner is the one button (W.10)', () => {
    const visitor = render(CharacterGearPanel, {
      props: { ...PROPS, gear: [], total: null, owned: false },
    }).body;
    const line = /<p class="[^"]*" id="gear" data-testid="character-gear-none">([^]*?)[<][/]p>/.exec(visitor);
    expect(text(line?.[1] ?? '')).toBe(gearCopy.noGear);
    expect(line?.[1]).toContain('href="/setup"');
    expect(line?.[1]).toContain('href="/login"');
    expect(visitor).not.toContain('character-gear-planner');
    const scored = render(CharacterGearPanel, {
      props: { ...PROPS, gear: [], total: null, owned: false, planner: true },
    }).body;
    expect(scored).toMatch(/href="[/]planner"[^>]*data-testid="character-gear-planner">Open the planner</);
  });
});

describe('CharacterGearPanel on an export that carried no gear and no talents (W.3)', () => {
  const UNREAD = GEAR_FIXTURE.map((slot) => ({ ...slot, worn: null, gap: slot.pick === null ? null : 5 }));
  const props = { ...PROPS, gear: UNREAD, unread: true, capturedAt: '2026-10-19T21:40:00', band: 20 };

  it('says the gear has not been read, in place of the slot list, the Empty rows and the totals', () => {
    const body = bare(render(CharacterGearPanel, { props }).body);
    expect(body).toContain('data-testid="export-not-read"');
    expect(text(body)).toContain(notReadCopy.lead);
    expect(text(body)).toContain('carried no gear or talents.');
    expect(text(body)).toContain(notReadCopy.steps);
    expect(body).not.toContain('character-gear-empty');
    expect(body).not.toContain('character-gear-totals');
    expect(body).not.toContain('character-gear-gap');
    expect(text(body)).not.toContain(gearCopy.empty);
    expect(text(body)).not.toContain(gearCopy.noPick);
    expect(text(body)).not.toContain('pts');
  });

  it('keeps the band’s best in slot below it, titled for the band, picks only', () => {
    const body = bare(render(CharacterGearPanel, { props }).body);
    expect(text(/data-testid="character-gear-title"[\s\S]*?<\/h2>/.exec(body)?.[0] ?? '')).toContain(
      'Best in slot for level 20 to 29',
    );
    expect(body.indexOf('export-not-read')).toBeLessThan(body.indexOf('character-gear-title'));
    expect(body).toContain('character-gear-row-head');
    expect(body).toContain('Halo of Transcendence');
    expect(body).not.toContain('character-gear-worn');
  });

  it('speaks to a visitor without the owner’s steps', () => {
    const body = text(bare(render(CharacterGearPanel, { props: { ...props, owned: false } }).body));
    expect(body).toContain(notReadCopy.visitorLead);
    expect(body).not.toContain(notReadCopy.steps);
  });

  it('titles the list by band: 60 alone, a range below it, plain with none', () => {
    expect(notReadCopy.bestInSlot(60)).toBe('Best in slot for level 60');
    expect(notReadCopy.bestInSlot(20)).toBe('Best in slot for level 20 to 29');
    expect(notReadCopy.bestInSlot(null)).toBe('Best in slot');
    expect(notReadCopy.synced('19 Oct, 21:40')).toBe(
      'The last sync from the game (19 Oct, 21:40) carried no gear or talents.',
    );
  });
});

describe('the spec the picks are for (A.2)', () => {
  it('names the spec and its source after the heading when the talents named it', () => {
    const body = render(CharacterGearPanel, {
      props: { ...PROPS, specLine: 'Fury · from your talents' },
    }).body;
    expect(bare(body)).toContain('data-testid="character-gear-spec"');
    expect(text(body)).toContain('Fury · from your talents');
  });

  it('draws nothing extra when a fight named the spec', () => {
    const body = render(CharacterGearPanel, { props: PROPS }).body;
    expect(bare(body)).not.toContain('character-gear-spec');
  });
});
