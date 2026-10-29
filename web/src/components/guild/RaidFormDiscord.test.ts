// web/src/components/guild/RaidFormDiscord.test.ts
// The raid form's Discord block rendered with svelte/server (guided-cards spec P.3): the switch
// with its helper, the reminder under it, and the three post rows. Toggling is clicked in the e2e.
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import { raidCopy } from '../../lib/guild/raid-copy';
import type { DiscordHelper, RaidFormValues } from '../../lib/guild/raid-form';
import { text } from '../../test-support/raid-fixtures';
import RaidFormDiscord from './RaidFormDiscord.svelte';

const VALUES: RaidFormValues = {
  zone: 'onyxias-lair',
  title: '',
  date: '2026-12-03',
  time: '20:00',
  duration: 180,
  timezone: 'America/Chicago',
  notes: '',
  repeat: false,
  postToDiscord: true,
  remindDayBefore: true,
};
const ON: DiscordHelper = {
  post: raidCopy.form.postNow('#raid-signups'),
  remind: raidCopy.form.remind,
  postRow: 'switch',
  remindRow: 'live',
};

function html(helper: DiscordHelper = ON, values: Partial<RaidFormValues> = {}, busy = false): string {
  return render(RaidFormDiscord, {
    props: { values: { ...VALUES, ...values }, helper, busy, set: () => {} },
  }).body;
}

const tag = (markup: string, testid: string): string =>
  markup.match(new RegExp(`<input[^>]*data-testid="${testid}"[^>]*>`))?.[0] ?? '';

describe('RaidFormDiscord', () => {
  it('draws the switch on with its helper and the reminder under it', () => {
    const markup = html();
    const sw = tag(markup, 'raid-form-discord');
    expect(sw).toContain('role="switch"');
    expect(sw).toContain('aria-checked="true"');
    expect(sw).toContain('checked');
    const describedBy = sw.match(/aria-describedby="([^"]+)"/)?.[1] ?? '';
    expect(describedBy).not.toBe('');
    expect(markup).toMatch(
      new RegExp(`id="${describedBy}"[^>]*>\\s*${raidCopy.form.postNow('#raid-signups')}`),
    );
    expect(text(markup)).toContain(raidCopy.form.postToDiscord);
    const remind = tag(markup, 'raid-form-remind');
    expect(remind).toContain('role="switch"');
    expect(remind).not.toMatch(/ disabled(=|\s|\/|>)/);
    expect(text(markup)).toContain(raidCopy.form.remindDayBefore);
    expect(text(markup)).toContain(raidCopy.form.remind);
  });

  it('switch off: the reminder is disabled, still checked, and says why', () => {
    const off: DiscordHelper = { ...ON, post: raidCopy.form.postOff, remind: raidCopy.form.remindNeedsPost };
    const markup = html(off, { postToDiscord: false });
    expect(tag(markup, 'raid-form-discord')).toContain('aria-checked="false"');
    const remind = tag(markup, 'raid-form-remind');
    expect(remind).toMatch(/ disabled(=|\s|\/|>)/);
    expect(remind).toContain('checked');
    expect(remind).toContain('aria-checked="true"');
    expect(remind).not.toContain('checked:bg-gold');
    expect(tag(markup, 'raid-form-discord')).toContain('checked:bg-gold');
    expect(text(markup)).toContain(raidCopy.form.remindNeedsPost);
  });

  it('reserves its height so the defaults loading shift nothing', () => {
    expect(html()).toContain('min-h-[164px]');
    expect(html()).toContain('sm:min-h-[144px]');
  });

  it('a bot post replaces the switch with the posted line', () => {
    const line = raidCopy.form.posted('#raid-signups');
    const markup = html({ ...ON, postRow: 'posted', post: line });
    expect(markup).not.toContain('data-testid="raid-form-discord"');
    expect(markup).toContain('data-testid="raid-form-discord-posted"');
    expect(text(markup)).toContain(line);
    expect(tag(markup, 'raid-form-remind')).not.toMatch(/ disabled(=|\s|\/|>)/);
  });

  it('a webhook post keeps the if-the-time-changes box', () => {
    const markup = html({ ...ON, postRow: 'webhookMoved', post: null }, { postToDiscord: false });
    expect(text(markup)).toContain(raidCopy.form.postToDiscordIfMoved);
    expect(tag(markup, 'raid-form-discord')).not.toContain('role="switch"');
    expect(tag(markup, 'raid-form-remind')).not.toMatch(/ disabled(=|\s|\/|>)/);
  });

  it('a sent reminder reads Reminder sent in place of the row', () => {
    const markup = html({ ...ON, remindRow: 'reminded', remind: raidCopy.form.reminded });
    expect(markup).not.toContain('data-testid="raid-form-remind"');
    expect(text(markup)).toContain(raidCopy.form.reminded);
  });

  it('while saving both rows are disabled', () => {
    const markup = html(ON, {}, true);
    expect(tag(markup, 'raid-form-discord')).toMatch(/ disabled(=|\s|\/|>)/);
    expect(tag(markup, 'raid-form-remind')).toMatch(/ disabled(=|\s|\/|>)/);
  });
});
