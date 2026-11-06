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

/** The answer slot and rows as TierPage.astro renders them, reduced to what the script reads. */
const ANSWER = `<div id="tier-answer"><div data-answer-skeleton></div><div data-answer-body hidden>
  <span data-field="label"></span><img data-field="crest" hidden><span data-field="name"></span>
  <span data-field="detail"></span>
  <span data-field="figure"><span class="ui-figure"><span>0<span>of 0</span></span></span></span>
  <span data-field="links"><a data-field="bis"></a><a data-field="guide"></a></span>
  <a data-field="other" hidden></a>
  <span data-field="message" hidden><span></span><button data-field="choose" hidden></button><button data-field="retry" hidden></button></span>
</div></div>`;

function renderPage(): void {
  document.documentElement.dataset.session = '1';
  const rows = Object.keys(views)
    .map((key) => `<li data-rank-key="${key}"><span data-you-tag hidden></span></li>`)
    .join('');
  const config = JSON.stringify({
    role: 'dps',
    views: { alliance: views, horde: views },
    factionParam: 'faction',
    factionAttribute: 'data-faction',
  });
  document.body.innerHTML = `${ANSWER}${rows}<script type="application/json" id="tier-client-config">${config}</script>`;
}

async function settle(): Promise<void> {
  await vi.advanceTimersByTimeAsync(10);
}

const field = (name: string) => document.querySelector<HTMLElement>(`[data-field="${name}"]`)!;
const answerText = (): string =>
  document.querySelector<HTMLElement>('[data-answer-body]')!.hidden
    ? ''
    : `${field('name').textContent} ${field('detail').textContent}`;
const marked = () =>
  [...document.querySelectorAll('[aria-current="true"]')].map((n) => n.getAttribute('data-rank-key'));
const shownTags = () => document.querySelectorAll('[data-you-tag]:not([hidden])').length;

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

  it('answers in the header and moves the mark when the selected character changes', async () => {
    findTierMe.mockResolvedValue(spec('warrior/arms'));
    const { mountTierPage } = await import('./tier-client');
    mountTierPage();
    await settle();
    expect(answerText()).toContain('Arms Warrior');
    expect(field('figure').textContent).toBe('1stof 3');
    expect(field('bis').getAttribute('href')).toBe('/bis/warrior/arms');
    expect(marked()).toEqual(['warrior/arms']);

    findTierMe.mockResolvedValue(spec('hunter/beast-mastery'));
    window.dispatchEvent(new Event(CURRENT_CHARACTER_CHANGED));
    await settle();

    expect(answerText()).toContain('Beast Mastery Hunter');
    expect(field('figure').textContent).toBe('2ndof 3');
    expect(marked()).toEqual(['hunter/beast-mastery']);
    expect(shownTags()).toBe(1);
  });

  it('points a spec on another role to its list, with no place and no marked row', async () => {
    findTierMe.mockResolvedValue(spec('paladin/protection'));
    const { mountTierPage } = await import('./tier-client');
    mountTierPage();
    await settle();
    expect(field('detail').textContent).toBe('Protection Paladin is on the Tank list.');
    expect(field('other').hidden).toBe(false);
    expect(field('other').getAttribute('href')).toBe('/tiers/tank');
    expect(field('figure').hidden).toBe(true);
    expect(marked()).toEqual([]);
  });

  it('says when the spec is not known yet, and when the read failed', async () => {
    findTierMe.mockResolvedValue({ kind: 'unknown', name: 'Thandrel' });
    const { mountTierPage } = await import('./tier-client');
    mountTierPage();
    await settle();
    expect(field('message').textContent).toContain("Thandrel's spec is not known yet.");
    expect(field('choose').hidden).toBe(false);
    findTierMe.mockResolvedValue({ kind: 'error' });
    window.dispatchEvent(new Event(CURRENT_CHARACTER_CHANGED));
    await settle();
    expect(field('message').textContent).toContain('We could not load your character.');
    expect(field('retry').hidden).toBe(false);
    expect(field('choose').hidden).toBe(true);
  });

  it('hides the answer when the new selection is no character', async () => {
    findTierMe.mockResolvedValue(spec('warrior/arms'));
    const { mountTierPage } = await import('./tier-client');
    mountTierPage();
    await settle();
    findTierMe.mockResolvedValue({ kind: 'none' });
    window.dispatchEvent(new Event(CURRENT_CHARACTER_CHANGED));
    await settle();
    expect(document.getElementById('tier-answer')!.hidden).toBe(true);
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
