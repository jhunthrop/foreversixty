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

// The paste state's HTML as it shipped before the connect lane (Svelte's comment markers removed):
// with connect off the block must still draw exactly this.
const PASTE_STATE_BEFORE_CONNECT = `<div id="discord" class="scroll-mt-20" role="presentation"><div class="bg-raised border-line rounded-panel flex flex-col gap-4 border px-4 py-4 md:px-6 md:py-5" data-testid="account-discord"><div class="flex items-center gap-[10px]"><span class="bg-gold h-2 w-2 rounded-full shadow-[0_0_10px_rgba(229,185,85,.8)]" aria-hidden="true"></span> <span class="label text-gold">Discord</span> </div>  <form class="flex flex-col gap-4" novalidate=""><p class="text-muted text-[13px]">Get a message in a Discord channel when a server sim finishes or a log finishes parsing.</p> <ol class="text-muted m-0 flex list-decimal flex-col gap-1 pl-5 text-[13px] marker:tabular-nums"><li data-testid="discord-step-1">In Discord, open the channel's settings, then Integrations, then Webhooks.</li><li data-testid="discord-step-2">Choose New Webhook, name it something you will recognize, then Copy Webhook URL.</li><li data-testid="discord-step-3">Paste it below.</li></ol> <div class="flex flex-col"><label class="label text-muted mb-2" for="discord-webhook">Webhook URL</label> <input type="text" name="discord-webhook" id="discord-webhook" class="bg-raised rounded-control text-text h-11 w-full border px-3 text-[15px] border-line-warm" inputmode="url" autocomplete="off" autocapitalize="off" autocorrect="off" spellcheck="false" enterkeyhint="done" placeholder="https://discord.com/api/webhooks/…" aria-describedby="discord-webhook-error discord-webhook-help" aria-invalid="false" value=""/> <div class="text-wipe min-h-[21px] pt-1 text-[13px]" id="discord-webhook-error" data-testid="discord-webhook-error"></div> <p id="discord-webhook-help" class="text-muted text-[13px]">We store the URL encrypted, never show it again, and delete it with your account.</p></div> <fieldset class="m-0 flex flex-col border-0 p-0" aria-busy="false"><legend class="label text-muted mb-1 p-0">Post when</legend> <label class="flex min-h-11 items-center gap-3 text-[14px]  "><input type="checkbox" checked="" data-testid="discord-event-sim"/> A server sim finishes</label><label class="flex min-h-11 items-center gap-3 text-[14px]  "><input type="checkbox" checked="" data-testid="discord-event-upload"/> A combat log you upload finishes parsing</label><label class="flex min-h-11 items-center gap-3 text-[14px] ml-7 "><input type="checkbox" data-testid="discord-event-private"/> Include private logs</label> <p class="text-muted text-[13px]">Off, private logs never post. On, everyone in the channel sees the log's name and “Only you”.</p></fieldset> <div class="flex flex-wrap items-center gap-x-4"><button type="submit" class="rounded-control border-gold bg-gold text-bg hover:bg-gold-hover inline-flex h-11 items-center border text-[12px] font-bold tracking-[0.06em] uppercase px-5" aria-busy="false" data-testid="discord-save">Save webhook</button> </div></form></div></div>`;

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
  return (body.match(/<(?:button|a)[^>]*bg-gold[^>]*>/g) ?? []).length;
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

type Notifications = NonNullable<Me['notifications']>;

const connectOn = (discord: Notifications['discord']): Notifications => ({ connect: true, discord });

describe('connect off: the paste state is unchanged', () => {
  it.each<[string, Notifications]>([
    ['connect false', { connect: false, discord: null }],
    ['no connect key', { discord: null }],
  ])('draws the shipped HTML (%s)', (_name, notifications) => {
    expect(uncommented(html(meWith(true, notifications)))).toBe(PASTE_STATE_BEFORE_CONNECT);
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
      /<a[^>]*href="https:\/\/api\.foreversixty\.gg\/v1\/notifications\/discord\/connect"[^>]*bg-gold[^>]*data-testid="discord-connect"/,
    );
    expect(filledControls(body)).toBe(1);
    expect(body).not.toContain('target=');
  });

  it('draws no paste form, steps, field or checkboxes', () => {
    for (const part of ['<form', 'discord-step-1', 'id="discord-webhook"', 'discord-event-sim'])
      expect(body).not.toContain(part);
  });

  it('uses no Discord branding', () => {
    expect(body).not.toMatch(/blurple|5865f2|<svg|<img/i);
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

  it('says Connected to a Discord channel over the shortened channel id', () => {
    expect(text(body)).toContain('Connected to a Discord channel');
    expect(text(body)).toContain(notificationsCopy.channelLine(connected.channel_id));
    expect(text(body)).not.toContain('Posts through');
    expect(text(body)).not.toContain(notificationsCopy.destinationUnnamed);
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
      'Connected to a Discord channel',
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
