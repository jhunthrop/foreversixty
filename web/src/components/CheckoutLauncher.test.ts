// web/src/components/CheckoutLauncher.test.ts
// The static render (what the page ships before hydration): the head with its plan box held and
// the Steps in the form column, row 2 busy, so nothing moves when the launcher reads the URL
// (premium spec 2026-11-08 section 13).
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import { checkoutCopy } from '../lib/billing/copy';
import CheckoutLauncher from './CheckoutLauncher.svelte';

function html(): string {
  return render(CheckoutLauncher, { props: {} }).body.replace(/<!--.*?-->/g, '');
}

describe('CheckoutLauncher static render', () => {
  it('draws the head and the form column before the URL is read', () => {
    const body = html();
    expect(body).toContain('data-testid="checkout-band"');
    expect(body).toMatch(/data-layout="form"[^>]*data-testid="checkout-body"/);
  });

  it('draws the three steps, Pay with Stripe busy with the redirect line', () => {
    const body = html();
    expect(body).toContain('data-testid="checkout-steps"');
    for (const title of [
      checkoutCopy.steps.signIn,
      checkoutCopy.steps.pay,
      checkoutCopy.steps.on('Premium'),
    ]) {
      expect(body).toContain(title);
    }
    expect(body).toContain('aria-label="Step 2: In progress"');
    expect(body).toContain('data-testid="checkout-message"');
    expect(body).toContain(checkoutCopy.stripeLine);
  });
});
