// web/src/components/ui/state.test.ts
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import EmptyState from './EmptyState.svelte';
import LoadError from './LoadError.svelte';
import Skeleton from './Skeleton.svelte';
import { uiCopy } from '../../lib/ui/copy';

describe('Skeleton', () => {
  it('reserves height, marks the region busy, and tells screen readers once', () => {
    const { body } = render(Skeleton, { props: { lines: 4, minHeight: 'min-h-[320px]' } });
    expect(body).toContain('aria-busy="true"');
    expect(body).toContain('min-h-[320px]');
    expect((body.match(/skeleton-block/g) ?? []).length).toBe(4);
    expect((body.match(new RegExp(uiCopy.loading, 'g')) ?? []).length).toBe(1);
  });
});

describe('LoadError', () => {
  it('shows the message and a Try again button when a retry is offered', () => {
    const { body } = render(LoadError, { props: { message: 'Rankings did not load.', onRetry: () => {} } });
    expect(body).toContain('role="alert"');
    expect(body).toContain('Rankings did not load.');
    expect(body).toContain(uiCopy.retry);
  });

  it('shows no button without a retry', () => {
    const { body } = render(LoadError, { props: { message: 'Gone.' } });
    expect(body).not.toContain('<button');
  });
});

describe('EmptyState', () => {
  it('shows the sentence and one action', () => {
    const { body } = render(EmptyState, {
      props: { message: 'No reports yet.', action: { label: 'Upload a log', href: '/logs' } },
    });
    expect(body).toContain('No reports yet.');
    expect(body).toContain('href="/logs"');
  });
});
