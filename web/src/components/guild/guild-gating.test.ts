// The gating spec's small components, rendered (spec 2026-11-05 §3.1 to §3.3).
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import { gateFactsFixture, GATE_CLAIM_HREF, GATE_SETTINGS_HREF } from '../../test-support/gate-facts';
import GuildClaimRow from './GuildClaimRow.svelte';
import GuildClaimedBy from './GuildClaimedBy.svelte';
import GuildInviteLink from './GuildInviteLink.svelte';
import GuildPlanGate from './GuildPlanGate.svelte';

const text = (markup: string): string =>
  markup
    .replace(/<!--[\s\S]*?-->/g, '')
    .replace(/<[^>]+>/g, '')
    .replace(/\s+/g, ' ')
    .trim();

describe('GuildPlanGate', () => {
  const gate = (tab: 'raids' | 'loot', over = {}): string =>
    render(GuildPlanGate, { props: { tab, facts: gateFactsFixture(over) } }).body;
  const holder = { name: 'Thornhide', class: 'druid', character_key: 'us/pvp/thornhide' };

  it('reads as one sentence with the link last, and carries the id the controls name', () => {
    const body = gate('loot');
    expect(text(body)).toBe('Awarding loot is part of Raid guild. Start the 7-day trial.');
    expect(body).toContain(`href="${GATE_SETTINGS_HREF}?plan=raid_guild#billing"`);
    expect(body).toContain('id="officer-tools-gate-loot"');
  });

  it('on a granted Guild plan reads as a note asking us, never the dead Move link (guild round 3)', () => {
    const body = gate('loot', { guildTier: true, granted: true });
    expect(text(body)).toBe(
      'Awarding loot is part of Raid guild. Your plan was granted by Forever Sixty. Ask us on Discord to move it to Raid guild.',
    );
    expect(body).not.toContain('upgrade=1');
    expect(body).toMatch(
      /<a[^>]*href="https:\/\/discord\.gg\/[^"]+"[^>]*target="_blank"[^>]*>Ask us on Discord<\/a>/,
    );
  });

  it('asks the claim holder by crest and class colour, with no link', () => {
    const body = gate('loot', {
      guildMaster: false,
      claim: { state: 'claimed', yours: false, claimed_by: holder },
    });
    expect(text(body)).toBe('Awarding loot is part of Raid guild. Ask Thornhide to turn it on.');
    expect(body).toContain('/crests/druid');
    expect(body).toContain('color: var(--color-class-druid)');
    expect(body).not.toContain('<a ');
    // 4 px each side of the crest: a margin before it (no text space) and the 4 px gap after it.
    expect(body).toMatch(/Ask<span class="ml-1 inline-flex[^"]*gap-1[^>]*>(<!---->)? </);
  });

  it('asks by role when the holder has no character on the roster', () => {
    const body = gate('raids', {
      guildMaster: false,
      claim: { state: 'claimed', yours: false, claimed_by: null },
    });
    expect(text(body)).toBe(
      'Planning raids and events is part of the Guild and Raid guild plans. Ask your guild master to turn it on.',
    );
  });

  it('an unclaimed guild: the guild master claims, an officer reads who does', () => {
    const claim = { state: 'unclaimed', yours: false, claimed_by: null };
    expect(text(gate('raids', { claim }))).toContain('Claim the guild first.');
    expect(gate('raids', { claim })).toContain(`href="${GATE_CLAIM_HREF}"`);
    const officer = gate('raids', { claim, guildMaster: false });
    expect(text(officer)).toContain('Your guild master claims the guild first.');
    expect(officer).not.toContain('<a ');
  });
});

describe('GuildClaimRow', () => {
  const row = (kind: 'guild-master' | 'officer' | 'pending'): string =>
    render(GuildClaimRow, {
      props: { kind, guildName: 'Olympus', claimHref: GATE_CLAIM_HREF },
    }).body;

  it('the guild master: the sentence, then the gold outline button to the claim page, never a fill', () => {
    const body = row('guild-master');
    expect(text(body)).toContain('You lead Olympus in game.');
    expect(text(body)).toContain('Claim Olympus');
    expect(body).toContain(`href="${GATE_CLAIM_HREF}"`);
    expect(body).toContain('border-gold');
    expect(body).not.toContain('linear-gradient');
  });

  it('an officer: the copy link, its href still the claim page', () => {
    const body = row('officer');
    expect(text(body)).toContain('Nobody has claimed Olympus here yet. Your guild master claims it.');
    expect(text(body)).toContain('Copy the claim link for your guild master');
    expect(body).toContain('data-testid="guild-claim-copy-link"');
  });

  it('pending: opens the claim', () => {
    const body = row('pending');
    expect(text(body)).toContain('A claim on Olympus is waiting for a second officer.');
    expect(text(body)).toContain('Open the claim');
  });
});

describe('GuildClaimedBy', () => {
  it('names the character, in class colour, with its crest', () => {
    const body = render(GuildClaimedBy, {
      props: { holder: { name: 'Thornhide', class: 'druid', character_key: 'k' } },
    }).body;
    expect(text(body)).toBe('Claimed by Thornhide');
    expect(body).toContain('/crests/druid');
    expect(body).toContain('font-semibold');
  });

  it('with no character on the roster says the guild master', () => {
    expect(text(render(GuildClaimedBy, { props: { holder: null } }).body)).toBe(
      'Claimed by the guild master',
    );
  });
});

describe('GuildInviteLink', () => {
  const link = (rotatedAt: string | null, disabled = false): string =>
    render(GuildInviteLink, { props: { guildId: 7, guildName: 'Olympus', rotatedAt, disabled } }).body;

  it('sits under #invite and says what the link does', () => {
    const body = link(null);
    expect(body).toContain('id="invite"');
    expect(body).toContain('scroll-mt-20');
    expect(text(body)).toContain(
      'Anyone with this link joins Olympus as a verified member: no officer has to approve them, and they can sign up for raids and see the roster, loot and readiness at once.',
    );
  });

  it('no link yet: a secondary with no confirm, the plan card keeping the one fill', () => {
    const body = link(null);
    expect(text(body)).toContain('Make a link to share in guild chat or Discord.');
    expect(text(body)).toContain('Make the invite link');
    const make = /<button[^>]*class="([^"]*)"[^>]*data-testid="guild-invite-make"/.exec(body)?.[1] ?? '';
    expect(make).toContain('bg-white/[.045]');
    expect(make).not.toContain('linear-gradient');
    expect(body).not.toContain('data-testid="guild-invite-rotate"');
    expect(body).not.toContain('data-testid="guild-invite-url"');
  });

  it('a later visit: no field and no URL, the date, and Make a new link', () => {
    const body = link('2026-11-12T10:00:00Z');
    expect(text(body)).toContain(
      'Made 12 Nov. It still works, but it cannot be shown again. Make a new one to copy it.',
    );
    expect(text(body)).toContain('Make a new link');
    expect(body).not.toContain('data-testid="guild-invite-url"');
    expect(body).not.toContain('/guild/invite/');
  });

  it('a contested claim disables the press', () => {
    expect(link('2026-11-12T10:00:00Z', true)).toMatch(/data-testid="guild-invite-rotate"|disabled/);
    expect(link(null, true)).toMatch(/\sdisabled/);
  });
});
