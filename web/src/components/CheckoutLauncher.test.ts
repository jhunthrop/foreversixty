// web/src/components/CheckoutLauncher.test.ts
// The static render (what the page ships before hydration): the band's ground and reserve,
// and the body slot reserved at the PremiumOwnedPanel's height, so nothing shifts later.
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import { CHECKOUT_BAND_RESERVE, CHECKOUT_BODY_RESERVE } from '../lib/billing/checkout-layout';
import CheckoutLauncher from './CheckoutLauncher.svelte';

function html(): string {
  return render(CheckoutLauncher, { props: {} }).body.replace(/<!--.*?-->/g, '');
}

describe('CheckoutLauncher static render', () => {
  it('draws the full-bleed band with its reserve before the URL is read', () => {
    const body = html();
    expect(body).toContain('data-testid="checkout-band"');
    expect(body).toContain(CHECKOUT_BAND_RESERVE);
    // No container of its own: the narrow body takes the site's .measure-form (640 px on this page).
    expect(body).not.toContain('mx-auto');
    expect(body).toMatch(/class="measure-form [^"]*"[^>]*data-testid="checkout-body"/);
  });

  it('reserves the body slot at the panel height around the one-line message', () => {
    const body = html();
    expect(body).toMatch(
      new RegExp(
        `class="[^"]*${CHECKOUT_BODY_RESERVE.replace(/[[\]]/g, '\\$&')}[^"]*"[^>]*data-testid="checkout-body"`,
      ),
    );
    expect(body).toContain('data-testid="checkout-message"');
  });
});
