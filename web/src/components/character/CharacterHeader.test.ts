// web/src/components/character/CharacterHeader.test.ts
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import type { Me, MeCharacter } from '../../lib/account/api';
import type { CharacterCardState } from '../../lib/bis/character-card-state.svelte';
import { mockEntitlements } from '../../test-support/billing-fixtures';
import CharacterHeader from './CharacterHeader.svelte';

const ANSWER = {
  strings: { failed: 'Could not load.' },
  load: () => Promise.reject(new Error('not used')),
};

/** The text a crawler would read: the markup with tags removed (the aria-hidden separator dots are decoration). */
function readableText(html: string): string {
  return html
    .replace(/<style[\s\S]*?<\/style>/g, '')
    .replace(/<[^>]*>/g, '')
    .replace(/&nbsp;|&middot;|\u00b7/g, '')
    .trim();
}

describe('CharacterHeader loading state (the server-rendered reserved box)', () => {
  const props = {
    action: { label: 'Open', href: '/bis' },
    answer: ANSWER,
    nameHeading: true,
    whenSignedOut: 'none' as const,
  };

  it('carries no readable text and never says "Placeholder"', () => {
    const { body } = render(CharacterHeader, { props });
    expect(body).toContain('data-testid="character-header-loading"');
    expect(body).not.toMatch(/placeholder/i);
    expect(readableText(body)).toBe('');
  });

  it('never renders the character name as a heading, even when the page asks for one', () => {
    const { body } = render(CharacterHeader, { props });
    expect(body).not.toContain('<h1');
  });
});

describe('CharacterHeader Supporter chip', () => {
  const CHARACTER: MeCharacter = {
    key: 'us/hardcore/elyra-duskvale',
    region: 'us',
    ruleset: 'hardcore',
    name: 'Elyra Duskvale',
    class: 'Priest',
    level: 60,
  };

  function headerWith(entitlements: Me['entitlements'] | undefined): string {
    const me = { user: { id: 7 }, characters: [CHARACTER], guilds: [], entitlements } as unknown as Me;
    const source: CharacterCardState = {
      status: 'ready',
      error: '',
      me,
      character: CHARACTER,
      refresh: () => {},
    };
    return render(CharacterHeader, { props: { action: { label: 'Open', href: '/bis' }, source } }).body;
  }

  it('draws the chip beside the name when the account shows the mark', () => {
    const body = headerWith(mockEntitlements({ supporter_mark: true }));
    expect(body).toContain('data-testid="supporter-chip"');
    expect(body.indexOf('Elyra Duskvale')).toBeLessThan(body.indexOf('supporter-chip'));
  });

  it('draws none when the mark is off', () => {
    expect(headerWith(mockEntitlements({ supporter_mark: false }))).not.toContain('supporter-chip');
  });

  it('draws none for an older /v1/me with no entitlements', () => {
    expect(headerWith(undefined)).not.toContain('supporter-chip');
  });

  it('never draws it in the loading ghost', () => {
    const { body } = render(CharacterHeader, {
      props: { action: { label: 'Open', href: '/bis' }, whenSignedOut: 'none' as const },
    });
    expect(body).not.toContain('supporter-chip');
  });
});

describe('CharacterHeader done line', () => {
  const CHARACTER: MeCharacter = {
    key: 'us/hardcore/elyra-duskvale',
    region: 'us',
    ruleset: 'hardcore',
    name: 'Elyra Duskvale',
    class: 'Priest',
    level: 60,
  };
  const RUN = {
    sim_id: 'sim-1',
    spec: 'mage-frost',
    title: '',
    dps: 1204,
    finished_at: '2026-10-24T11:00:00Z',
  };

  function headerWith(
    server: boolean,
    finished: unknown[] | undefined,
    layout: 'band' | 'card' = 'band',
  ): string {
    const me = {
      user: { id: 7 },
      characters: [CHARACTER],
      guilds: [],
      entitlements: mockEntitlements({ server_sims: server }),
      finished_sims: finished,
    } as unknown as Me;
    const source: CharacterCardState = {
      status: 'ready',
      error: '',
      me,
      character: CHARACTER,
      refresh: () => {},
    };
    return render(CharacterHeader, {
      props: { action: { label: 'Open', href: '/bis' }, source, layout },
    }).body;
  }

  it('draws the line inside the panel for a member with server sims and an unseen run', () => {
    const body = headerWith(true, [RUN]);
    expect(body).toContain('data-testid="finished-sim-line"');
    expect(body).toContain('ch-has-foot');
    expect(body.indexOf('character-header-action')).toBeLessThan(body.indexOf('finished-sim-line'));
  });

  it.each([
    ['no server sims', false, [RUN]],
    ['an empty list', true, []],
    ['an older API', true, undefined],
  ])('draws no line for %s', (_name, server, finished) => {
    const body = headerWith(server, finished);
    expect(body).not.toContain('finished-sim-line');
    expect(body).not.toContain('ch-has-foot');
  });

  it('draws none in the stacked card layout', () => {
    expect(headerWith(true, [RUN], 'card')).not.toContain('finished-sim-line');
  });

  it('the loading skeleton carries the reserve row, shown only under data-finished', () => {
    const { body } = render(CharacterHeader, {
      props: { action: { label: 'Open', href: '/bis' }, whenSignedOut: 'none' as const },
    });
    expect(body).toContain('ch-foot-reserve');
    expect(body).toContain('ch-loading');
    expect(body).not.toContain('finished-sim-line');
  });
});
