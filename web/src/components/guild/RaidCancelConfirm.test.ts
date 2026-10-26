// web/src/components/guild/RaidCancelConfirm.test.ts
// The cancel confirm rendered with svelte/server (raids spec 4 "Cancel confirm"): the question, the
// remembered Discord box, Cancel the raid (Cancelling… in flight) and Keep it. Focus on Keep it and
// Escape-as-Keep-it run in the browser, in the raid e2e.
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import { raidCopy } from '../../lib/guild/raid-copy';
import { text } from '../../test-support/raid-fixtures';
import RaidCancelConfirm from './RaidCancelConfirm.svelte';

const DISABLED = / disabled(=| |>)/;

function html(postCancel: boolean, busy: boolean): string {
  return render(RaidCancelConfirm, {
    props: {
      postCancel,
      busy,
      buttonClass: 'btn',
      onPostChange: () => {},
      onConfirm: () => {},
      onKeep: () => {},
    },
  }).body;
}

describe('RaidCancelConfirm', () => {
  it('asks, offers the Discord post and both answers', () => {
    const body = html(false, false);
    expect(text(body)).toBe(
      [
        raidCopy.cancelConfirm.question,
        raidCopy.cancelConfirm.postToDiscord,
        raidCopy.cancelConfirm.confirm,
        raidCopy.cancelConfirm.keep,
      ].join(' '),
    );
    expect(body).not.toMatch(/data-testid="raid-cancel-discord"[^>]*checked/);
    expect(body).not.toMatch(DISABLED);
  });

  it('ticks the remembered choice', () => {
    expect(html(true, false)).toMatch(/<input[^>]*checked[^>]*data-testid="raid-cancel-discord"/);
  });

  it('reads Cancelling… and holds the confirm while the cancel is in flight, Keep it live', () => {
    const body = html(false, true);
    expect(text(body)).toContain(raidCopy.cancelConfirm.confirming);
    expect(body.match(/<button[^>]* disabled/g)?.length).toBe(1);
  });
});
