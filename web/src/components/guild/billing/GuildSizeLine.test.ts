import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import { text } from '../../../test-support/raid-fixtures';
import GuildSizeLine from './GuildSizeLine.svelte';

const html = (payer: boolean): string =>
  render(GuildSizeLine, { props: { count: 28, payer, onMove: () => {} } }).body;

describe('GuildSizeLine', () => {
  it('states the count and nothing switching off, with exactly one Move to Raid guild for a payer', () => {
    const body = html(true);
    expect(text(body)).toContain('28 verified members. Guild covers up to 25; nothing switches off.');
    expect(body.match(/Move to Raid guild/g)?.length).toBe(1);
  });

  it('draws no button for a reader who may not bill', () => {
    expect(html(false)).not.toContain('Move to Raid guild');
  });
});
