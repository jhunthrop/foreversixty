// web/src/components/account/AccountDiscordBlock.test.ts
// One static render per state (this project's vitest config hands components Svelte's server
// entry, so there is no client click to simulate; HelpNote.test.ts documents the same limit).
// The flows are covered in lib/account/discord-actions.test.ts and the clicks in Playwright.
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import type { Me } from '../../lib/account/api';
import { removeQuestionFor } from '../../lib/account/notifications';
import { notificationsCopy } from '../../lib/account/notifications-copy';
import { mockEntitlements } from '../../test-support/billing-fixtures';
import { GUILD_ID, discordBlock } from '../../test-support/notifications-api';
import AccountDiscordBlock from './AccountDiscordBlock.svelte';
import DiscordActions from './discord/DiscordActions.svelte';
import DiscordConnect from './discord/DiscordConnect.svelte';
import DiscordPasteForm from './discord/DiscordPasteForm.svelte';

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

function uncommented(body: string): string {
  return body.replace(/<!--.*?-->/g, '');
}

function filledControls(body: string): number {
  return (body.match(/<(?:button|a)[^>]*bg-\[linear-gradient[^>]*>/g) ?? []).length;
}

function filledButtons(body: string): number {
  return (body.match(/<button[^>]*bg-\[linear-gradient[^>]*>/g) ?? []).length;
}

/** The opening tag of the element carrying a data-testid. */
function tagOf(body: string, testid: string): string {
  return new RegExp(`<[a-z]+[^>]*data-testid="${testid}"[^>]*>`).exec(body)?.[0] ?? '';
}

const isOn = (body: string, testid: string): boolean => tagOf(body, testid).includes('aria-checked="true"');
const isDisabled = (body: string, testid: string): boolean => tagOf(body, testid).includes('disabled=""');

describe('AccountDiscordBlock', () => {
  it('free: the card head, the sentence and See plans as an action link, no tag, no form', () => {
    const body = html(meWith(false, { discord: null }));
    expect(body).toContain('data-testid="discord-free"');
    expect(text(body)).toContain(notificationsCopy.freeSentence);
    expect(body).toContain('href="/premium#notifications"');
    expect(body).toContain('account-discord"');
    expect(body).not.toContain('discord-state-tag');
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

  it('empty: the paste form with both switches on and the only filled button', () => {
    const body = html(meWith(true, { discord: null }));
    expect(body).toContain('data-testid="account-discord"');
    expect(text(body)).toContain(notificationsCopy.intro);
    for (const step of notificationsCopy.howSteps) expect(text(body)).toContain(step);
    expect(body).toContain('discord-step-1');
    expect(body).toContain('discord-step-2');
    expect(body).toContain('discord-step-3');
    expect(body).toContain('id="discord-webhook"');
    expect(body).toContain('data-testid="discord-webhook-error"');
    expect(isOn(body, 'discord-event-sim')).toBe(true);
    expect(isOn(body, 'discord-event-upload')).toBe(true);
    expect(isOn(body, 'discord-event-private')).toBe(false);
    expect(tagOf(body, 'discord-event-sim')).toContain('role="switch"');
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

  it('empty: the three rows in order, the private switch nested and enabled while the log switch is on', () => {
    const body = html(meWith(true, { discord: null }));
    const order = ['discord-event-sim', 'discord-event-upload', 'discord-event-private'].map((id) =>
      body.indexOf(id),
    );
    expect(order).toEqual([...order].sort((a, b) => a - b));
    expect(text(body)).toContain(notificationsCopy.eventUploadPrivate);
    expect(tagOf(body, 'discord-event-private-row')).toContain('padding-left: 46px');
    expect(isDisabled(body, 'discord-event-private')).toBe(false);
  });

  it('saved: no channel id, a secondary test button and a danger Remove, no filled button', () => {
    const body = html(meWith(true, { discord: discordBlock() }));
    expect(text(body)).toContain('Posts through “Raid pings”');
    expect(text(body)).not.toContain('Channel 1098');
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

  it('saved without a name: the channel line, and never an id', () => {
    const body = html(meWith(true, { discord: discordBlock({ name: null }) }));
    expect(text(body)).toContain(notificationsCopy.destinationChannel);
    expect(text(body)).not.toContain('Channel 1198');
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

  it('saved: the private switch is disabled, value kept, while the log switch is off', () => {
    const events = { sim_finished: true, upload_parsed: false, upload_private: true };
    const body = html(meWith(true, { discord: discordBlock({ events }) }));
    expect(isDisabled(body, 'discord-event-private')).toBe(true);
    expect(isOn(body, 'discord-event-private')).toBe(true);
    expect(isDisabled(body, 'discord-event-sim')).toBe(false);
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

  it('switched off: the notice, then the replace row, before the actions and events', () => {
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
    ].map((part) => body.indexOf(part, notice));
    expect(order).toEqual([...order].sort((a, b) => a - b));
    expect(order.every((at) => at > notice)).toBe(true);
    expect(body.indexOf('discord-replace')).toBeGreaterThan(order[2] ?? 0);
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

  it('lapsed: gold notice, disabled switches, Remove only', () => {
    const body = html(meWith(false, { discord: discordBlock() }));
    expect(body).toContain('border-gold');
    expect(text(body)).toContain(notificationsCopy.lapsedTitle);
    expect(text(body)).toContain(notificationsCopy.lapsedBody);
    expect(body).toContain('href="/premium#notifications"');
    for (const id of ['discord-event-sim', 'discord-event-upload', 'discord-event-private'])
      expect(isDisabled(body, id)).toBe(true);
    expect(tagOf(body, 'discord-state-tag')).toContain('data-tone="muted"');
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

type Notifications = NonNullable<Me['notifications']>;

const connectOn = (discord: Notifications['discord']): Notifications => ({ connect: true, discord });

describe('connect off: the paste state', () => {
  it.each<[string, Notifications]>([
    ['connect false', { connect: false, discord: null }],
    ['no connect key', { discord: null }],
  ])('draws the paste form with no connect control (%s)', (_name, notifications) => {
    const body = uncommented(html(meWith(true, notifications)));
    expect(body).toContain('<form');
    expect(body).toContain('data-testid="discord-save"');
    expect(body).not.toContain('discord-connect');
  });

  it('quiet (a seat offer holds the page fill): Save webhook is a secondary button, not a fill', () => {
    const me = meWith(true, { connect: false, discord: null });
    const body = render(AccountDiscordBlock, { props: { me, quiet: true } }).body;
    expect(filledButtons(body)).toBe(0);
    expect(tagOf(body, 'discord-save')).toContain('border-line-card');
  });

  it('has no connect control, disclosure or back link', () => {
    const body = html(meWith(true, { connect: false, discord: null }));
    for (const id of ['discord-connect', 'discord-paste-instead', 'discord-connect-instead'])
      expect(body).not.toContain(id);
  });
});

describe('connect state', () => {
  const body = html(meWith(true, connectOn(null)));

  it('shows the intro, the gold link, its explanation, the status line and the paste disclosure, in order', () => {
    const order = [
      notificationsCopy.intro,
      notificationsCopy.connect,
      notificationsCopy.connectExplanation,
      'discord-status',
      notificationsCopy.pasteInstead,
    ].map((part) => body.indexOf(part));
    expect(order).toEqual([...order].sort((a, b) => a - b));
    expect(order.every((at) => at >= 0)).toBe(true);
  });

  it('is a plain same-tab link to the connect route, styled as the one filled control', () => {
    expect(body).toMatch(
      /<a[^>]*href="https:\/\/api\.foreversixty\.gg\/v1\/notifications\/discord\/connect"[^>]*bg-\[linear-gradient[^>]*data-testid="discord-connect"/,
    );
    expect(filledControls(body)).toBe(1);
    expect(body).not.toContain('target=');
  });

  it('draws no paste form, steps, field or switches', () => {
    for (const part of ['<form', 'discord-step-1', 'id="discord-webhook"', 'discord-event-sim'])
      expect(body).not.toContain(part);
  });

  it('uses no Discord branding', () => {
    expect(body).not.toMatch(/blurple|5865f2|<img/i);
  });

  it('is never shown to a free member', () => {
    expect(html(meWith(false, connectOn(null)))).not.toContain('discord-connect"');
  });
});

describe('connect state, status line', () => {
  const connectHtml = (status: { ok: boolean; message: string } | null): string =>
    render(DiscordConnect, { props: { status, onStuck: () => {}, onPaste: () => {}, onCancel: () => {} } })
      .body;

  it('is the reserved line, empty without an outcome', () => {
    const body = connectHtml(null);
    expect(body).toContain('min-h-[21px]');
    expect(body).not.toContain('role=');
  });

  it.each([
    [notificationsCopy.returnConnected, true, 'status', 'text-text'],
    [notificationsCopy.returnDenied, true, 'status', 'text-text'],
    [notificationsCopy.returnFailed, false, 'alert', 'text-wipe'],
  ])('draws %s as role=%s', (message, ok, role, tone) => {
    const body = connectHtml({ ok, message });
    expect(body).toMatch(new RegExp(`<p class="${tone}" role="${role}"`));
    expect(text(body)).toContain(message);
  });
});

describe('connect replace view', () => {
  const body = render(DiscordConnect, {
    props: { status: null, replacing: true, onStuck: () => {}, onPaste: () => {}, onCancel: () => {} },
  }).body;

  it('omits the intro and ends with Cancel', () => {
    expect(text(body)).not.toContain(notificationsCopy.intro);
    expect(body.indexOf('discord-cancel')).toBeGreaterThan(body.indexOf('discord-paste-instead'));
  });
});

describe('paste view when connect is on', () => {
  const props = { webhook: '', events: discordBlock().events, onSubmit: () => {}, onCancel: () => {} };

  it('adds the back link under Save webhook', () => {
    const body = render(DiscordPasteForm, { props: { ...props, onConnect: () => {} } }).body;
    expect(body.indexOf('discord-connect-instead')).toBeGreaterThan(body.indexOf('discord-save'));
    expect(text(body)).toContain(notificationsCopy.connectInstead);
  });

  it('has none without onConnect', () => {
    expect(render(DiscordPasteForm, { props }).body).not.toContain('discord-connect-instead');
  });
});

describe('saved, connected', () => {
  const connected = discordBlock({ source: 'connected', name: null, guild_id: GUILD_ID });
  const body = html(meWith(true, connectOn(connected)));

  it('says it posts to a channel in your server, with no channel id', () => {
    expect(text(body)).toContain(notificationsCopy.destinationChannel);
    expect(text(body)).not.toContain(notificationsCopy.channelLine(connected.channel_id));
    expect(text(body)).not.toContain('Posts through');
  });

  it('never shows the Discord-side name', () => {
    const named = html(meWith(true, connectOn({ ...connected, name: 'Forever Sixty' })));
    expect(named).not.toContain('Forever Sixty');
  });

  it('links Open in Discord to the channel, in a new tab', () => {
    expect(body).toMatch(
      new RegExp(
        `<a[^>]*href="https://discord.com/channels/${GUILD_ID}/${connected.channel_id}"[^>]*target="_blank"[^>]*rel="noopener"[^>]*data-testid="discord-open"`,
      ),
    );
    expect(text(body)).toContain(notificationsCopy.openInDiscord);
  });

  it('replaces with a link to the connect route and the note, under the status line', () => {
    expect(body).toMatch(
      /<a[^>]*href="https:\/\/api\.foreversixty\.gg\/v1\/notifications\/discord\/connect"[^>]*data-testid="discord-replace"/,
    );
    expect(text(body)).toContain(notificationsCopy.reconnect);
    expect(text(body)).toContain(notificationsCopy.reconnectNote);
    expect(text(body)).not.toContain(notificationsCopy.replace);
    expect(body.indexOf('discord-replace')).toBeGreaterThan(body.indexOf('discord-status'));
    expect(filledControls(body)).toBe(0);
  });

  it('draws the same panel when connect is off', () => {
    expect(text(html(meWith(true, { connect: false, discord: connected })))).toContain(
      notificationsCopy.destinationChannel,
    );
  });
});

describe('saved, pasted, with connect on', () => {
  it('shows Open in Discord only once the guild is known', () => {
    expect(html(meWith(true, connectOn(discordBlock())))).not.toContain('discord-open');
    const known = html(meWith(true, connectOn(discordBlock({ guild_id: GUILD_ID }))));
    expect(known).toContain('discord-open');
    expect(text(known)).toContain('Posts through “Raid pings”');
  });

  it('keeps Use a different webhook as a button, not a link', () => {
    const body = html(meWith(true, connectOn(discordBlock())));
    expect(body).toMatch(/<button[^>]*data-testid="discord-replace"/);
    expect(text(body)).toContain(notificationsCopy.replace);
    expect(text(body)).not.toContain(notificationsCopy.reconnectNote);
  });
});

describe('switched off, connected', () => {
  it('offers Connect a different channel inside the notice', () => {
    const discord = discordBlock({ source: 'connected', disabled_reason: 'failed five times', failures: 5 });
    const body = html(meWith(true, connectOn(discord)));
    expect(body.indexOf('discord-replace')).toBeGreaterThan(body.indexOf('discord-notice'));
    expect(body.indexOf('discord-replace')).toBeLessThan(body.indexOf('discord-test'));
    expect(text(body)).toContain(notificationsCopy.reconnect);
  });
});

describe('the remove question', () => {
  const question = (discord: ReturnType<typeof discordBlock>): string =>
    text(
      render(DiscordActions, {
        props: {
          state: 'saved',
          confirming: true,
          removeQuestion: removeQuestionFor(discord),
          busy: null,
          onTest: () => {},
          onAskRemove: () => {},
          onRemove: () => {},
          onCancel: () => {},
        },
      }).body,
    );

  it('asks the connected question for a connected webhook', () => {
    expect(question(discordBlock({ source: 'connected' }))).toContain(
      notificationsCopy.removeQuestionConnected,
    );
  });

  it('asks the pasted question for a pasted webhook', () => {
    expect(question(discordBlock())).toContain(notificationsCopy.removeQuestionPasted);
  });
});
