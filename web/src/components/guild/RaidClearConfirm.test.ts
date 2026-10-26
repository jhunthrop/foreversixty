// web/src/components/guild/RaidClearConfirm.test.ts
// The clear confirm rendered with svelte/server (attendance spec 3.1): the question, Clear the match
// in the failure outline (Clearing… in flight) and Keep it. Focus on Keep it and Escape-as-Keep-it
// run in the browser, as the cancel confirm's do.
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import { attendanceCopy } from '../../lib/guild/attendance-copy';
import { RAID_BUTTON, RAID_WIPE_BUTTON } from '../../lib/guild/raid-buttons';
import { text } from '../../test-support/raid-fixtures';
import RaidClearConfirm from './RaidClearConfirm.svelte';

/** A `disabled` attribute, not the `disabled:` class variants. */
const DISABLED = / disabled(=| |>)/;

function html(busy: boolean): string {
  return render(RaidClearConfirm, { props: { busy, onConfirm: () => {}, onKeep: () => {} } }).body;
}

describe('RaidClearConfirm', () => {
  it('asks, then offers Clear the match and Keep it', () => {
    const copy = attendanceCopy.clearConfirm;
    expect(text(html(false))).toBe(`${copy.prompt} ${copy.confirm} ${copy.keep}`);
  });

  it('wears the failure outline on the clear and the plain one on Keep it', () => {
    const body = html(false);
    expect(body).toContain(`class="${RAID_WIPE_BUTTON}"`);
    expect(body).toContain(`class="${RAID_BUTTON}"`);
  });

  it('says Clearing… in flight, disables the clear and keeps Keep it live', () => {
    const body = html(true);
    expect(text(body)).toContain(attendanceCopy.match.clearing);
    expect(/<button[^>]*data-testid="raid-clear-confirm-button"/.exec(body)?.[0]).toMatch(DISABLED);
    expect(/<button[^>]*data-testid="raid-clear-keep"/.exec(body)?.[0]).not.toMatch(DISABLED);
    expect(body).toContain('aria-busy="true"');
  });
});
