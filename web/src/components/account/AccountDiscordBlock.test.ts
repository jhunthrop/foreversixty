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

  it.each([true, false])('off (no notifications key, entitled %s): renders nothing at all', (entitled) => {
    const { notifications: _omitted, ...bare } = meWith(entitled, { discord: null });
    expect(text(html(bare)).trim()).toBe('');
    expect(html(bare)).not.toContain('id="discord"');
  });

  it('carries the #discord anchor in every state', () => {
    expect(html(meWith(false, { discord: null }))).toContain('id="discord"');
    expect(html(meWith(true, { discord: null }))).toContain('id="discord"');
  });

  it('empty: the paste form with both boxes on and the only filled button', () => {
    const body = html(meWith(true, { discord: null }));
    expect(body).toContain('data-testid="account-discord"');
    expect(text(body)).toContain(notificationsCopy.intro);
    for (const step of notificationsCopy.howSteps) expect(text(body)).toContain(step);
    expect(body).toContain('discord-step-1');
    expect(body).toContain('discord-step-2');
    expect(body).toContain('discord-step-3');
    expect(body).toContain('id="discord-webhook"');
    expect(body).toContain('data-testid="discord-webhook-error"');
    expect(body).toMatch(/<input[^>]*checked[^>]*discord-event-sim/);
    expect(body).toMatch(/<input[^>]*checked[^>]*discord-event-upload/);
    expect(body).not.toMatch(/<input[^>]*checked[^>]*discord-event-private/);
    expect(text(body)).toContain(notificationsCopy.eventsNote);
    expect(filledButtons(body)).toBe(1);
    expect(body).not.toContain('discord-replace');
  });

  it('empty: the steps come before the field and the error line sits directly under the input', () => {
    const body = html(meWith(true, { discord: null }));
    expect(body.indexOf('discord-step-3')).toBeLessThan(body.indexOf('id="discord-webhook"'));
    expect(body).toMatch(/<input[^>]*id="discord-webhook"[^>]*>\s*<div[^>]*discord-webhook-error/);
    expect(body.indexOf('discord-webhook-error')).toBeLessThan(body.indexOf(notificationsCopy.trustHelper));
  });

  it('empty: the three rows in order, the private box indented and enabled while the log box is on', () => {
    const body = html(meWith(true, { discord: null }));
    const order = ['discord-event-sim', 'discord-event-upload', 'discord-event-private'].map((id) =>
      body.indexOf(id),
    );
    expect(order).toEqual([...order].sort((a, b) => a - b));
    expect(text(body)).toContain(notificationsCopy.eventUploadPrivate);
    expect(body).toMatch(/class="[^"]*ml-7[^"]*"[^>]*>\s*<input[^>]*discord-event-private/);
    expect(body).not.toMatch(/<input[^>]*disabled[^>]*discord-event-private/);
  });

  it('saved: channel id, two outline buttons, no filled button', () => {
    const body = html(meWith(true, { discord: discordBlock() }));
    expect(text(body)).toContain('Posts through “Raid pings”');
    expect(text(body)).toContain(notificationsCopy.channelLine('1098765432101234567'));
    expect(text(body)).not.toContain('1098765432101234567');
    expect(body).toContain('data-testid="discord-test"');
    expect(body).toContain('data-testid="discord-remove"');
    expect(body).toContain('data-testid="discord-replace"');
    expect(filledButtons(body)).toBe(0);
    expect(body).not.toContain('<form');
    expect(body).not.toContain('discord-notice');
  });

  it('saved: the lead and the quoted name are one sentence once the tags are gone', () => {
    const body = html(meWith(true, { discord: discordBlock({ name: 'Raid pings' }) }));
    const named = notificationsCopy.destinationNamed('Raid pings');
    expect(body.replace(/<!--.*?-->/g, '').replace(/<[^>]+>/g, '')).toContain(`${named.lead} ${named.name}`);
  });

  it('saved without a name: the unnamed line', () => {
    const body = html(meWith(true, { discord: discordBlock({ name: null }) }));
    expect(text(body)).toContain(notificationsCopy.destinationUnnamed);
    expect(text(body)).not.toContain('Posts through');
  });

  it('saved: a name with markup, a link and markdown is plain text, cut at 40 characters', () => {
    const name = `<b>bold</b> [x](https://evil.example) **${'n'.repeat(60)}`;
    const body = html(meWith(true, { discord: discordBlock({ name }) }));
    expect(body).not.toContain('<b>');
    expect(body).not.toContain('href="https://evil.example');
    expect(body).toContain('&lt;b>bold&lt;/b>');
    const named = notificationsCopy.destinationNamed(name);
    expect(text(body)).toContain(`${named.lead} ${named.name}`.replaceAll('<', '&lt;'));
    expect(text(body)).not.toContain('n'.repeat(60));
    expect(text(body)).toContain('…”');
  });

  it('saved: the private box is disabled, value kept, while the log box is off', () => {
    const events = { sim_finished: true, upload_parsed: false, upload_private: true };
    const body = html(meWith(true, { discord: discordBlock({ events }) }));
    expect(body).toMatch(/<input[^>]*disabled[^>]*discord-event-private/);
    expect(body).toMatch(/<input[^>]*checked[^>]*discord-event-private/);
    expect(body).not.toMatch(/<input[^>]*disabled[^>]*discord-event-sim/);
    expect(text(body)).toContain(notificationsCopy.eventsNote);
  });

  it('saved: the replace link is present (hidden only while the remove question is open)', () => {
    expect(html(meWith(true, { discord: discordBlock() }))).toContain('discord-replace');
  });

  it('saved with 3 failures shows the warning sentence', () => {
    const body = html(
      meWith(true, { discord: discordBlock({ failures: 3, last_error: 'Missing Permissions' }) }),
    );
    const { failures, lastError } = notificationsCopy.warningLines(3, 'Missing Permissions');
    expect(text(body)).toContain(failures);
    expect(text(body)).toContain(lastError);
  });

  it('saved with failures: the count line in text-text, the Discord error line in text-muted', () => {
    const body = html(
      meWith(true, { discord: discordBlock({ failures: 3, last_error: 'Missing Permissions' }) }),
    );
    const { failures, lastError } = notificationsCopy.warningLines(3, 'Missing Permissions');
    expect(body).toMatch(new RegExp(`<p class="text-text[^"]*"[^>]*>\\s*${failures}\\s*</p>`));
    expect(body).toMatch(new RegExp(`<p class="text-muted[^"]*"[^>]*>\\s*${lastError}\\s*</p>`));
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
    expect(text(body)).toContain(notificationsCopy.switchedOffFix);
    expect(text(body)).not.toContain('Sends a test message. If Discord accepts it');
    expect(text(body)).toContain(notificationsCopy.testAndSwitchOn);
    expect(filledButtons(body)).toBe(1);
    expect(body).toContain('discord-replace');
  });

  it('switched off: four lines in the notice, the replace link inside it, before the actions and events', () => {
    const body = html(
      meWith(true, {
        discord: discordBlock({
          disabled_reason: 'failed five times',
          failures: 5,
          last_error: 'Unknown Webhook',
        }),
      }),
    );
    const notice = body.indexOf('discord-notice');
    const order = [
      notificationsCopy.switchedOff,
      'Last error from Discord',
      notificationsCopy.switchedOffFix,
      'discord-replace',
    ].map((part) => body.indexOf(part, notice));
    expect(order).toEqual([...order].sort((a, b) => a - b));
    expect(order.every((at) => at > notice)).toBe(true);
    expect(body.indexOf('discord-test')).toBeGreaterThan(body.indexOf('discord-replace'));
    expect(body.indexOf('discord-event-sim')).toBeGreaterThan(body.indexOf('discord-test'));
  });

  it('switched off with no last error: three lines', () => {
    const body = html(
      meWith(true, { discord: discordBlock({ disabled_reason: 'failed five times', last_error: null }) }),
    );
    expect(text(body)).not.toContain('Last error from Discord');
    expect(text(body)).toContain(notificationsCopy.switchedOffFix);
  });

  it('lapsed: gold notice, disabled boxes, Remove only', () => {
    const body = html(meWith(false, { discord: discordBlock() }));
    expect(body).toContain('border-gold');
    expect(text(body)).toContain(notificationsCopy.lapsedTitle);
    expect(text(body)).toContain(notificationsCopy.lapsedBody);
    expect(body).toContain('href="/premium"');
    expect(body).toMatch(/<input[^>]*disabled[^>]*discord-event-sim/);
    expect(body).toMatch(/<input[^>]*disabled[^>]*discord-event-upload/);
    expect(body).toMatch(/<input[^>]*disabled[^>]*discord-event-private/);
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
