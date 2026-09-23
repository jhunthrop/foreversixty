// web/src/components/HomeAccountPanel.test.ts
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import HomeAccountPanel from './HomeAccountPanel.svelte';

describe('HomeAccountPanel', () => {
  it('renders only the empty, occlusion-ready root before the session resolves (SSR)', () => {
    const { body } = render(HomeAccountPanel, { props: {} });
    expect(body).toContain('[grid-area:1/1]');
    expect(body).not.toContain('data-testid="home-account-panel"');
  });
});
