// The server render of the two /app islands (the vitest config hands mount() Svelte's server
// entry, so the presses are covered by lib/app-beta/form.test.ts and tests/e2e/app-beta.spec.ts).
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import { appBetaCopy as copy } from '../../lib/app-beta/copy';
import StoreButtons from './StoreButtons.svelte';
import TesterForm from './TesterForm.svelte';

const decode = (s: string): string =>
  s.replaceAll('&#39;', "'").replaceAll('&quot;', '"').replaceAll('&amp;', '&');

describe('TesterForm', () => {
  const body = decode(render(TesterForm).body);

  it('asks for the email and the phone, iPhone until the browser says otherwise', () => {
    expect(body).toMatch(/<input[^>]*type="email"[^>]*autocomplete="email"/);
    expect(body).toContain(copy.emailLabel);
    expect(body).toMatch(/role="radiogroup"[^>]*aria-label="Your phone"/);
    expect(body).toMatch(/aria-checked="true"[^>]*data-testid="app-beta-platform-ios"/);
  });

  it('carries the one primary and the consent line', () => {
    expect(body).toContain(copy.submit);
    expect(body).toContain(copy.consent);
    expect(copy.consent).toBe("We'll only email you about the app beta.");
  });
});

describe('StoreButtons', () => {
  it('draws one button per live store, the first primary', () => {
    const stores = [
      { platform: 'ios' as const, url: 'https://apps.apple.com/app/id1' },
      {
        platform: 'android' as const,
        url: 'https://play.google.com/store/apps/details?id=gg.foreversixty.app',
      },
    ];
    const body = decode(render(StoreButtons, { props: { stores } }).body);
    expect(body).toContain('href="https://apps.apple.com/app/id1"');
    expect(body).toContain(copy.stores.ios);
    expect(body).toContain(copy.stores.android);
    expect(body.match(/data-testid="app-beta-store-/g)).toHaveLength(2);
  });
});

describe('the copy', () => {
  it('has no em dash anywhere', () => {
    const every = JSON.stringify(copy) + copy.created('ios') + copy.exists('android');
    expect(every).not.toContain('—');
  });
});
