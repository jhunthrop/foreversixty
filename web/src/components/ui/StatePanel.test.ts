import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import { createRawSnippet } from 'svelte';
import StatePanel from './StatePanel.svelte';

const CHILDREN = createRawSnippet(() => ({
  render: () => '<p data-testid="inner">Inside</p>',
}));

describe('StatePanel', () => {
  it('renders the glowing gold dot, the label, and the children inside the panel box', () => {
    const { body } = render(StatePanel, {
      props: { label: 'Devices', testid: 'account-devices', children: CHILDREN },
    });
    expect(body).toContain('data-testid="account-devices"');
    expect(body).toContain('bg-raised');
    expect(body).toContain('rounded-panel');
    expect(body).toContain('bg-gold');
    expect(body).toContain('>Devices<');
    expect(body).toContain('data-testid="inner"');
  });

  it('shows an Updated stamp on the right when given one', () => {
    const { body } = render(StatePanel, {
      props: { label: 'Devices', updated: '3 hours ago', testid: 'account-devices', children: CHILDREN },
    });
    expect(body).toContain('data-testid="account-devices-updated"');
    expect(body).toContain('Updated');
    expect(body).toContain('3 hours ago');
  });

  it('shows no Updated stamp when none is given', () => {
    const { body } = render(StatePanel, { props: { label: 'You', children: CHILDREN } });
    expect(body).not.toContain('Updated');
  });

  it('renders an aside snippet instead of the Updated stamp when both are given', () => {
    const aside = createRawSnippet(() => ({
      render: () => '<a data-testid="refresh" href="/refresh">Refresh from Battle.net</a>',
    }));
    const { body } = render(StatePanel, {
      props: { label: 'Characters', updated: '4 hours ago', aside, children: CHILDREN },
    });
    expect(body).toContain('data-testid="refresh"');
    expect(body).not.toContain('Updated 4 hours ago');
  });
});
