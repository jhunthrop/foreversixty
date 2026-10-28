// Server render only (the vitest config mounts Svelte's server entry): the markup the block and
// the openers own. The decisions behind clicks are lib/billing/claim-flow.test.ts.
import { render } from 'svelte/server';
import { describe, expect, it, vi } from 'vitest';
import type { MeGuild } from '../../lib/account/api';
import type { ClaimControls } from '../../lib/billing/claim-flow';
import { OLYMPUS } from '../../test-support/subscribe-fixtures';
import ClaimConfirmBlock from './ClaimConfirmBlock.svelte';
import GuildStepBlock from './GuildStepBlock.svelte';

const GUILD = { id: OLYMPUS.id, name: OLYMPUS.name, claim: null };
const IO = { claim: vi.fn(), confirm: vi.fn() };
const plain = (html: string): string => html.replace(/<!--.*?-->/g, '');

function block(kind: 'claim' | 'start_claim' | 'confirm_claim', error: string | null = null): string {
  return plain(
    render(ClaimConfirmBlock, {
      props: {
        kind,
        guild: GUILD,
        io: IO,
        error,
        onDone: () => {},
        onRefused: () => {},
        onCancel: () => {},
      },
    }).body,
  );
}

describe('ClaimConfirmBlock', () => {
  it('draws the sentence, the commit as the primary fill, and a 44 px Cancel', () => {
    const body = block('claim');
    expect(body).toContain('Claim OLYMPUS XXVII? You may attempt one claim every 30 days');
    expect(body).toMatch(/bg-gold[^>]*>\s*Yes, claim it/);
    expect(body).toContain('border-line-soft');
    expect(body).toMatch(/min-h-11[^>]*underline|underline[^>]*min-h-11/);
    expect(body).toContain('Cancel');
    expect(body).toContain('aria-busy="false"');
  });

  it.each([
    ['start_claim', 'Yes, start the claim'],
    ['confirm_claim', 'Yes, confirm it'],
  ] as const)('%s commit label', (kind, label) => {
    expect(block(kind)).toContain(label);
  });

  it('shows the refusal as an alert and nothing when there is none', () => {
    expect(block('claim')).not.toContain('role="alert"');
    const body = block('claim', 'You made a claim in the last 30 days.');
    expect(body).toMatch(/role="alert"[^>]*>\s*You made a claim in the last 30 days\./);
  });
});

function controls(open: boolean, locked = false): ClaimControls {
  return {
    open,
    locked,
    error: null,
    rereading: false,
    io: IO,
    onOpen: () => {},
    onCancel: () => {},
    onDone: () => {},
    onRefused: () => {},
  };
}

function card(step: string, claim: ClaimControls | undefined): string {
  const guild = {
    id: 2,
    name: 'OLYMPUS XXVII',
    region: 'us',
    ruleset: 'forever',
    slug: 'olympus-xxvii',
    verified: true,
    claim: null,
    subscribe: { step },
  } as unknown as MeGuild & { slug: string };
  return plain(
    render(GuildStepBlock, {
      props: {
        guild,
        plan: 'guild',
        interval: 'yearly',
        selfUrl: '/premium/checkout?plan=guild',
        viewerIsOfficer: true,
        guildPageHref: '/guild/us/forever/olympus-xxvii',
        onCheckAgain: () => Promise.resolve(true),
        claim,
      },
    }).body,
  );
}

describe('claim openers on the card', () => {
  it('draws the gold-outline opener, collapsed, and the block when open', () => {
    const closed = card('claim', controls(false));
    expect(closed).toContain('Claim OLYMPUS XXVII');
    expect(closed).toContain('border-gold text-gold w-full justify-center');
    expect(closed).toContain('aria-expanded="false"');
    expect(closed).not.toContain('claim-confirm');
    expect(card('claim', controls(true))).toContain('guild-2-claim-confirm');
  });

  it('disables the opener while another card holds the confirm', () => {
    expect(card('start_claim', controls(false, true))).toMatch(/<button[^>]*disabled[^>]*>\s*Start a claim/);
  });

  it('claim_link_bnet is a link to the Battle.net start with this card as next', () => {
    const body = card('claim_link_bnet', undefined);
    expect(body).toContain('/v1/auth/battlenet/start?next=');
    expect(body).toContain(encodeURIComponent('/premium/checkout?plan=guild#guild-2'));
    expect(body).toContain('Link Battle.net');
  });
});
