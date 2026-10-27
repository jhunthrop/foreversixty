import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import type { TransferView } from '../../../lib/billing/contract';
import { TRANSFER_DECLINED, TRANSFER_OFFERED, planView } from '../../../test-support/billing-fixtures';
import { text } from '../../../test-support/raid-fixtures';
import GuildWhoPays from './GuildWhoPays.svelte';

const CANDIDATES = [{ user_id: 40, name: 'Mirelle' }];
const html = (transfer: TransferView | null, candidates = CANDIDATES): string =>
  render(GuildWhoPays, {
    props: { guildId: 7, plan: planView(), transfer, candidates, onChanged: () => Promise.resolve() },
  }).body;

describe('GuildWhoPays', () => {
  it('offers the transfer button with the help line', () => {
    const body = html(null);
    expect(text(body)).toContain('Who pays');
    expect(body).toContain('guild-transfer-open');
    expect(body).not.toMatch(/\sdisabled(=|\s|>)/);
  });

  it('with no officer to send to, disables the button and says why', () => {
    const body = html(null, []);
    expect(body).toContain('disabled');
    expect(text(body)).toContain('No other officer has a Forever Sixty account yet.');
  });

  it('while the offer waits, shows who and Withdraw instead of the button', () => {
    const body = html(TRANSFER_OFFERED);
    expect(text(body)).toContain('Offered to Mirelle on');
    expect(text(body)).toContain('Waiting for them.');
    expect(body).toContain('guild-transfer-withdraw');
    expect(body).not.toContain('guild-transfer-open');
  });

  it('after a decline, names who declined and brings the button back', () => {
    const body = html(TRANSFER_DECLINED);
    expect(text(body)).toContain('Mirelle declined.');
    expect(body).toContain('guild-transfer-open');
  });
});
