import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import { exportHref } from '../../../lib/billing/guild-billing-api';
import GuildRecords from './GuildRecords.svelte';

describe('GuildRecords', () => {
  it('links both CSV downloads to the API export routes', () => {
    const body = render(GuildRecords, { props: { guildId: 7 } }).body;
    expect(body).toContain(`href="${exportHref(7, 'attendance')}"`);
    expect(body).toContain(`href="${exportHref(7, 'loot')}"`);
    expect(body).toContain('Download attendance (CSV)');
    expect(body).toContain('Download loot (CSV)');
  });
});
