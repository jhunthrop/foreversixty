// web/src/components/account/AccountDiscordBlock.test.ts
// One static render per state (this project's vitest config hands components Svelte's server
// entry, so there is no client click to simulate; HelpNote.test.ts documents the same limit).
// The flows are covered in lib/account/discord-actions.test.ts and the clicks in Playwright.
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import type { Me } from '../../lib/account/api';
import { notificationsCopy } from '../../lib/account/notifications-copy';
import { mockEntitlements } from '../../test-support/billing-fixtures';
import { discordBlock } from '../../test-support/notifications-api';
import AccountDiscordBlock from './AccountDiscordBlock.svelte';

const TOKEN = 'abcDEF_secret-token-9';

function meWith(entitled: boolean, discord: Me['notifications']): Me {
  return {
    user: { id: 1, battletag: null, email: null, role: 'member', anonymize: false },
    characters: [],
    guilds: [],
    entitlements: mockEntitlements({ notifications: entitled }),
    notifications: discord,
  };
}

function html(me: Me): string {
  return render(AccountDiscordBlock, { props: { me } }).body;
}

function text(body: string): string {
  return body
    .replace(/<!--.*?-->/g, '')
    .replace(/<[^>]+>/g, ' ')
    .replace(/\s+/g, ' ');
}

function filledButtons(body: string): number {
  return (body.match(/<button[^>]*bg-gold[^>]*>/g) ?? []).length;
}

describe('AccountDiscordBlock', () => {
  it('free: the sentence and See plans, no panel', () => {
    const body = html(meWith(false, { discord: null }));
    expect(body).toContain('data-testid="discord-free"');
    expect(text(body)).toContain(notificationsCopy.freeSentence);
    expect(body).toContain('href="/premium"');
    expect(body).not.toContain('account-discord"');
    expect(body).not.toContain('<form');
  });

  it('carries the #discord anchor in every state', () => {
    expect(html(meWith(false, { discord: null }))).toContain('id="discord"');
    expect(html(meWith(true, { discord: null }))).toContain('id="discord"');
  });

  it('empty: the paste form with both boxes on and the only filled button', () => {
    const body = html(meWith(true, { discord: null }));
    expect(body).toContain('data-testid="account-discord"');
    expect(text(body)).toContain(notificationsCopy.intro);
    expect(body).toContain('id="discord-webhook"');
    expect(body).toContain('data-testid="discord-webhook-error"');
    expect(body).toMatch(/<input[^>]*checked[^>]*discord-event-sim/);
    expect(body).toMatch(/<input[^>]*checked[^>]*discord-event-upload/);
    expect(text(body)).toContain(notificationsCopy.eventsNote);
    expect(filledButtons(body)).toBe(1);
    expect(body).not.toContain('discord-replace');
  });

  it('saved: channel id, two outline buttons, no filled button', () => {
    const body = html(meWith(true, { discord: discordBlock() }));
    expect(text(body)).toContain('Posts to channel 1098765432101234567');
    expect(body).toContain('data-testid="discord-test"');
    expect(body).toContain('data-testid="discord-remove"');
    expect(body).toContain('data-testid="discord-replace"');
    expect(filledButtons(body)).toBe(0);
    expect(body).not.toContain('<form');
    expect(body).not.toContain('discord-notice');
  });

  it('saved with 3 failures shows the warning sentence', () => {
    const body = html(
      meWith(true, { discord: discordBlock({ failures: 3, last_error: 'Missing Permissions' }) }),
    );
    expect(text(body)).toContain(notificationsCopy.warning(3, 'Missing Permissions'));
  });

  it('switched off: the notice and Test and switch on as the only filled button', () => {
    const body = html(
      meWith(true, {
        discord: discordBlock({
          disabled_reason: 'failed five times',
          failures: 5,
          last_error: 'Unknown Webhook',
        }),
      }),
    );
    expect(body).toContain('border-wipe');
    expect(text(body)).toContain(notificationsCopy.switchedOff);
    expect(text(body)).toContain('Last error from Discord: Unknown Webhook');
    expect(text(body)).toContain(notificationsCopy.switchedOffHelper);
    expect(text(body)).toContain(notificationsCopy.testAndSwitchOn);
    expect(filledButtons(body)).toBe(1);
    expect(body).toContain('discord-replace');
  });

  it('lapsed: gold notice, disabled boxes, Remove only', () => {
    const body = html(meWith(false, { discord: discordBlock() }));
    expect(body).toContain('border-gold');
    expect(text(body)).toContain(notificationsCopy.lapsedTitle);
    expect(text(body)).toContain(notificationsCopy.lapsedBody);
    expect(body).toContain('href="/premium"');
    expect(body).toMatch(/<input[^>]*disabled[^>]*discord-event-sim/);
    expect(body).toMatch(/<input[^>]*disabled[^>]*discord-event-upload/);
    expect(text(body)).not.toContain(notificationsCopy.eventsNote);
    expect(body).toContain('discord-remove');
    expect(body).not.toContain('discord-test"');
    expect(body).not.toContain('discord-replace');
  });

  it('never renders a webhook id or token, nor puts a url in a title or aria-label', () => {
    for (const me of [meWith(true, { discord: discordBlock() }), meWith(true, { discord: null })]) {
      const body = html(me);
      expect(body).not.toContain(TOKEN);
      expect(body).not.toMatch(/\b(title|aria-label)=/);
    }
  });
});
