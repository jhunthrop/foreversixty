import { describe, expect, it } from 'vitest';
import { specString } from '../../test-support/spec-table';
import { notificationsCopy } from '../account/notifications-copy';
import { guildBotCopy, shortServerName } from './discord-bot-copy';
import { guildDiscordCopy } from './discord-copy';

const SPEC = new URL('../../../../design/specs/2026-10-27-discord-bot.md', import.meta.url);
const HEADING = '### 6.1 On the site';
const row = (label: string, index: number, fill: Record<string, string> = {}): string =>
  specString(SPEC, HEADING, label, index, fill);
const SERVER = 'Olympus Discord';
/** Segments the spec lists that are data, not copy. */
const DATA = '{Server}';

function expectRow(label: string, texts: readonly string[]): void {
  texts.forEach((text, index) => {
    if (text !== DATA) expect(text).toBe(row(label, index));
  });
}

describe('guildBotCopy, verbatim from spec 6.1', () => {
  it('pins the add path', () => {
    expectRow('Add button', [guildBotCopy.add]);
    expectRow('Add explanation', [guildBotCopy.addExplanation]);
    expectRow('Webhook path', [guildBotCopy.webhookPath, guildBotCopy.webhookPathNote]);
    expectRow('Bot offer', [guildBotCopy.offer]);
  });

  it('pins the picker', () => {
    const p = guildBotCopy.picker;
    expectRow('Picker', [
      p.added,
      DATA,
      p.serverLead,
      DATA,
      p.label,
      p.placeholder,
      p.helper,
      p.save,
      p.saving,
      p.removeBot,
      p.cancel,
      p.emptyPress,
    ]);
    expectRow('Picker list states', [p.loading, DATA, p.reload, p.loadFailed, p.retry]);
    expect(p.empty(SERVER)).toBe(row('Picker list states', 1, { Server: SERVER }));
    expectRow('Picker race', [p.race]);
  });

  it('pins the failure warning, which names the channel and never the switch-off', () => {
    const w = guildBotCopy.failureWarning;
    const channel = { channel: 'raid-signups' };
    expect(w.failures(1, 'raid-signups')).toBe(row('Failure warning', 0, channel));
    expect(w.failures(3, 'raid-signups')).toBe(
      `${row('Failure warning', 1, { n: '3', ...channel })} Send a test message to check.`,
    );
    expect(w.permissions('raid-signups')).toBe(row('Failure warning', 2, channel));
    expect(w.noAnswer).toBe(row('Failure warning', 3));
  });

  it('pins the installed state', () => {
    const i = guildBotCopy.installed;
    expectRow('Installed', [
      i.lead,
      DATA,
      i.in,
      DATA,
      i.withButtons,
      i.open,
      i.test,
      i.testing,
      i.change,
      i.remove,
      i.removing,
      i.saved,
    ]);
    expect(i.channel('raid-signups')).toBe(row('Installed', 1, { channel: 'raid-signups' }));
    expectRow('Fallback note', [i.fallbackNote]);
    expect(i.readOnly).toBe(row('Read-only', 2));
  });

  it('pins the linked members in its three forms', () => {
    const i = guildBotCopy.installed;
    expect(i.linkedMembers(9, 31)).toBe(row('Linked members', 0, { linked: '9', members: '31' }));
    expect(i.linkedMembers(1, 31)).toBe(row('Linked members', 1));
    expect(i.linkedMembers(0, 31)).toBe(row('Linked members', 2));
  });

  it('pins removed in Discord and the remove question', () => {
    const r = guildBotCopy.removedInDiscord;
    expect(r.title(SERVER)).toBe(row('Removed in Discord', 0, { Server: SERVER }));
    expectRow('Removed in Discord', [DATA, r.detail, r.withWebhook, r.withoutWebhook]);
    const q = guildBotCopy.remove;
    const lead = row('Remove question', 0, { Server: SERVER });
    expect(q.question(SERVER, true)).toBe(lead + row('Remove question', 1));
    expect(q.question(SERVER, false)).toBe(lead + row('Remove question', 2));
    expectRow('Remove question', [DATA, DATA, DATA, q.confirm, q.cancel, DATA, q.failed]);
  });

  it('pins the read-only refusal and the install returns', () => {
    expect(guildBotCopy.planEnded).toBe(row('Read-only', 3));
    const r = guildBotCopy.returns;
    expectRow('Install returns', [r.installed, r.denied, r.failed, DATA, r.takenExpired]);
    expect(r.taken('Mograine Raiders')).toBe(row('Install returns', 3, { Guild: 'Mograine Raiders' }));
  });
});

describe('strings the section shares with the webhook, pinned to the bot rows', () => {
  it('reads the same pending, removed and stuck lines', () => {
    expect(notificationsCopy.connecting).toBe(row('Add button', 1));
    expect(notificationsCopy.removed).toBe(row('Remove question', 5));
    expect(guildDiscordCopy.connectStuck).toBe(row('Install returns', 5));
  });
});

describe('shortServerName', () => {
  it('cuts at 60 characters with an ellipsis', () => {
    expect([...shortServerName('x'.repeat(80))]).toHaveLength(60);
    expect(shortServerName('x'.repeat(80)).endsWith('…')).toBe(true);
    expect(shortServerName(' Olympus ')).toBe('Olympus');
  });
});
