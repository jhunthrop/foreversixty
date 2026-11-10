// web/src/lib/tiers/tier-client.test.ts
// @vitest-environment jsdom
// The tier list answers for whoever the header selector names, and answers again when the
// selection changes -- no reload.
import { afterEach, beforeEach, describe, expect, it, vi } from 'vitest';
import { CURRENT_CHARACTER_CHANGED } from '../current-character';
import { dpsInput, tankInput } from './tier-test-support';
import { rankRole } from './tier-list';
import { tierSpecViews } from './tier-callout';
import type { TierMe } from './tier-me';

const findTierMe = vi.fn<() => Promise<TierMe>>();
vi.mock('./tier-me', () => ({ findTierMe: () => findTierMe() }));
const spec = (key: string): TierMe => ({ kind: 'spec', key });

const dps = rankRole(
  [
    dpsInput('warrior', 'arms', 'Arms', 95),
    dpsInput('hunter', 'beast-mastery', 'Beast Mastery', 90),
    dpsInput('mage', 'fire', 'Fire', 80),
  ],
  'dps',
);
const tank = rankRole(
  [tankInput('paladin', 'protection', 'Protection', { dtps: 3, effective_health: 1, tps: 1 })],
  'tank',
);
const views = tierSpecViews({ dps, tank, healer: [] });

/** The spec panel and rows as TierSpecCard.astro and TierPage.astro render them, reduced to what the
 *  script reads. */
const CARD = `<div id="tier-answer"><div data-answer-skeleton></div><div data-answer-body hidden>
  <h2 data-field="name"></h2><p><span data-field="standing"></span><a data-field="other" hidden></a></p>
  <div><a data-field="bis"></a><a data-field="guide"></a>
  <button data-field="choose" hidden></button><button data-field="retry" hidden></button></div>
</div></div>`;

function renderPage(): void {
  document.documentElement.dataset.session = '1';
  const rows = Object.keys(views)
    .map((key) => `<li data-rank-key="${key}"></li>`)
    .join('');
  const config = JSON.stringify({
    role: 'dps',
    views: { alliance: views, horde: views },
    factionParam: 'faction',
    factionAttribute: 'data-faction',
  });
  document.body.innerHTML = `${CARD}${rows}<script type="application/json" id="tier-client-config">${config}</script>`;
}

async function settle(): Promise<void> {
  await vi.advanceTimersByTimeAsync(10);
}

const field = (name: string) => document.querySelector<HTMLElement>(`[data-field="${name}"]`)!;
const card = (): string =>
  document.querySelector<HTMLElement>('[data-answer-body]')!.hidden
    ? ''
    : `${field('name').textContent} | ${field('standing').textContent}`;
const shown = () => ['bis', 'guide', 'other', 'choose', 'retry'].filter((name) => !field(name).hidden);
const marked = () =>
  [...document.querySelectorAll('[aria-current="true"]')].map((n) => n.getAttribute('data-rank-key'));

describe('mountTierPage', () => {
  beforeEach(() => {
    vi.useFakeTimers();
    renderPage();
  });
  afterEach(() => {
    vi.useRealTimers();
    findTierMe.mockReset();
    vi.resetModules();
  });

  it('names the spec in its class colour, marks its row (no line repeats it), and moves the mark', async () => {
    findTierMe.mockResolvedValue(spec('warrior/arms'));
    const { mountTierPage } = await import('./tier-client');
    mountTierPage();
    await settle();
    expect(card()).toBe('Arms Warrior | ');
    expect(field('standing').parentElement!.hidden).toBe(true);
    expect(field('name').style.color).toBe('var(--color-class-warrior)');
    expect(shown()).toEqual(['bis', 'guide']);
    expect(field('bis').getAttribute('href')).toBe('/bis/warrior/arms/60');
    expect(field('guide').getAttribute('href')).toBe('/guides/warrior/arms');
    expect(marked()).toEqual(['warrior/arms']);

    findTierMe.mockResolvedValue(spec('hunter/beast-mastery'));
    window.dispatchEvent(new Event(CURRENT_CHARACTER_CHANGED));
    await settle();

    expect(card()).toBe('Beast Mastery Hunter | ');
    expect(marked()).toEqual(['hunter/beast-mastery']);
  });

  it('points a spec on another role to its list, with no marked row', async () => {
    findTierMe.mockResolvedValue(spec('paladin/protection'));
    const { mountTierPage } = await import('./tier-client');
    mountTierPage();
    await settle();
    expect(card()).toBe('Protection Paladin | ');
    expect(shown()).toEqual(['bis', 'guide', 'other']);
    expect(field('other').textContent).toBe('On the Tanks list ›');
    expect(field('guide').getAttribute('href')).toBe('/guides/paladin/protection');
    expect(field('other').getAttribute('href')).toBe('/tiers/tank');
    expect(marked()).toEqual([]);
  });

  it('says when the spec is not known yet, and when the read failed', async () => {
    findTierMe.mockResolvedValue({ kind: 'unknown', name: 'Thandrel' });
    const { mountTierPage } = await import('./tier-client');
    mountTierPage();
    await settle();
    expect(card()).toBe("Your spec | Thandrel's spec is not known yet.");
    expect(shown()).toEqual(['choose']);
    findTierMe.mockResolvedValue({ kind: 'error' });
    window.dispatchEvent(new Event(CURRENT_CHARACTER_CHANGED));
    await settle();
    expect(card()).toBe('Your spec | We could not load your character.');
    expect(shown()).toEqual(['retry']);
  });

  it('offers the character choice when the new selection is no character', async () => {
    findTierMe.mockResolvedValue(spec('warrior/arms'));
    const { mountTierPage } = await import('./tier-client');
    mountTierPage();
    await settle();
    findTierMe.mockResolvedValue({ kind: 'none' });
    window.dispatchEvent(new Event(CURRENT_CHARACTER_CHANGED));
    await settle();
    expect(card()).toBe('Your spec | No character chosen.');
    expect(shown()).toEqual(['choose']);
    expect(marked()).toEqual([]);
  });

  it('lets the newest selection win when an older answer lands last', async () => {
    let releaseFirst: (me: TierMe) => void = () => undefined;
    findTierMe.mockImplementationOnce(() => new Promise((resolve) => (releaseFirst = resolve)));
    const { mountTierPage } = await import('./tier-client');
    mountTierPage();
    await settle();
    findTierMe.mockResolvedValue(spec('hunter/beast-mastery'));
    window.dispatchEvent(new Event(CURRENT_CHARACTER_CHANGED));
    await settle();
    releaseFirst(spec('warrior/arms'));
    await settle();
    expect(marked()).toEqual(['hunter/beast-mastery']);
  });
});
