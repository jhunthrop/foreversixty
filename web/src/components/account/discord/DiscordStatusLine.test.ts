import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import DiscordStatusLine from './DiscordStatusLine.svelte';

const html = (props: Record<string, unknown>): string =>
  render(DiscordStatusLine, { props: { status: null, ...props } }).body;

describe('DiscordStatusLine', () => {
  it('reserves its 21 px line by default, empty without a result', () => {
    expect(html({})).toContain('min-h-[21px]');
  });

  it('with reserve off draws nothing until there is a result (no dead band)', () => {
    expect(html({ reserve: false })).not.toContain('discord-status');
    const body = html({ reserve: false, status: { ok: true, message: 'Sent.' } });
    expect(body).toContain('data-testid="discord-status"');
    expect(body).not.toContain('min-h-[21px]');
    expect(body).toContain('role="status"');
  });

  it('a failure is an alert in the failure colour', () => {
    const body = html({ status: { ok: false, message: 'Nope.' } });
    expect(body).toContain('role="alert"');
    expect(body).toContain('text-wipe');
  });
});
