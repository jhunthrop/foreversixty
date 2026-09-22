// web/src/components/GuildStatus.test.ts
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import GuildStatus from './GuildStatus.svelte';

describe('GuildStatus', () => {
  it('renders a Skeleton sized to the caller while loading', () => {
    const { body } = render(GuildStatus, {
      props: {
        status: 'loading',
        error: '',
        onRetry: () => {},
        lines: 4,
        minHeight: 'min-h-[200px]',
        testid: 'guild',
      },
    });
    expect(body).toContain('data-testid="guild-skeleton"');
    expect(body).toContain('min-h-[200px]');
  });

  it('renders LoadError with the caller-prefixed testid on failure', () => {
    const { body } = render(GuildStatus, {
      props: {
        status: 'failed',
        error: 'That did not load.',
        onRetry: () => {},
        lines: 4,
        minHeight: 'min-h-[200px]',
        testid: 'guild-claim',
      },
    });
    expect(body).toContain('data-testid="guild-claim-error"');
    expect(body).toContain('That did not load.');
  });
});
