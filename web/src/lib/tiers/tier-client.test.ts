// web/src/lib/tiers/tier-client.test.ts
// @vitest-environment jsdom
// The tier list answers for whoever the header selector names, and answers again when the
// selection changes -- no reload.
import { afterEach, beforeEach, describe, expect, it, vi } from 'vitest';
import { CURRENT_CHARACTER_CHANGED } from '../current-character';
import { dpsInput, tankInput } from './tier-test-support';
import { rankRole } from './tier-list';
import { tierSpecViews } from './tier-callout';

const findCharacterKey = vi.fn<() => Promise<string | undefined>>();
vi.mock('./tier-me', () => ({ findCharacterKey: () => findCharacterKey() }));

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

function renderPage(): void {
  document.documentElement.dataset.session = '1';
  const rows = Object.keys(views)
    .map((key) => `<div class="tier-row" data-spec-key="${key}"><span data-you-slot></span></div>`)
    .join('');
  const config = JSON.stringify({
    role: 'dps',
    views: { alliance: views, horde: views },
    paths: {},
    factionParam: 'faction',
    factionAttribute: 'data-faction',
  });
  document.body.innerHTML = `<div id="tier-you-slot" hidden></div>${rows}<script type="application/json" id="tier-client-config">${config}</script>`;
}

async function settle(): Promise<void> {
  await vi.advanceTimersByTimeAsync(10);
}

function calloutText(): string | undefined {
  return document.querySelector('[data-testid="tier-callout"]')?.textContent ?? undefined;
}

describe('mountTierPage', () => {
  beforeEach(() => {
    vi.useFakeTimers();
    renderPage();
  });
  afterEach(() => {
    vi.useRealTimers();
    findCharacterKey.mockReset();
    vi.resetModules();
  });

  it('moves the callout and the YOUR SPEC badge when the selected character changes', async () => {
    findCharacterKey.mockResolvedValue('warrior/arms');
    const { mountTierPage } = await import('./tier-client');
    mountTierPage();
    await settle();
    expect(calloutText()).toContain('Arms Warrior');
    expect(document.querySelector('.tier-row.is-you')?.getAttribute('data-spec-key')).toBe('warrior/arms');

    findCharacterKey.mockResolvedValue('hunter/beast-mastery');
    window.dispatchEvent(new Event(CURRENT_CHARACTER_CHANGED));
    await settle();

    expect(calloutText()).toContain('Beast Mastery Hunter');
    expect(calloutText()).not.toContain('Arms Warrior is');
    const mine = document.querySelectorAll('.tier-row.is-you');
    expect(mine).toHaveLength(1);
    expect(mine[0]?.getAttribute('data-spec-key')).toBe('hunter/beast-mastery');
    expect(document.querySelectorAll('[data-testid="tier-you-pill"]')).toHaveLength(1);
  });

  it('clears the answer when the new selection has no spec on this list', async () => {
    findCharacterKey.mockResolvedValue('warrior/arms');
    const { mountTierPage } = await import('./tier-client');
    mountTierPage();
    await settle();
    findCharacterKey.mockResolvedValue(undefined);
    window.dispatchEvent(new Event(CURRENT_CHARACTER_CHANGED));
    await settle();
    expect(calloutText()).toBeUndefined();
    expect(document.querySelectorAll('.tier-row.is-you')).toHaveLength(0);
  });

  it('lets the newest selection win when an older answer lands last', async () => {
    let releaseFirst: (key: string | undefined) => void = () => undefined;
    findCharacterKey.mockImplementationOnce(() => new Promise((resolve) => (releaseFirst = resolve)));
    const { mountTierPage } = await import('./tier-client');
    mountTierPage();
    await settle();
    findCharacterKey.mockResolvedValue('hunter/beast-mastery');
    window.dispatchEvent(new Event(CURRENT_CHARACTER_CHANGED));
    await settle();
    releaseFirst('warrior/arms');
    await settle();
    expect(document.querySelector('.tier-row.is-you')?.getAttribute('data-spec-key')).toBe(
      'hunter/beast-mastery',
    );
  });
});
