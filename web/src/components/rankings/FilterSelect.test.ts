// web/src/components/rankings/FilterSelect.test.ts
// The optgroup output and the invalid state (vitest renders on Svelte's server entry, so this is
// the markup; the select's behaviour is Playwright's).
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import FilterSelect from './FilterSelect.svelte';

const BASE = {
  id: 'pick',
  label: 'Raid channel',
  value: '',
  options: [{ value: '', label: 'Choose a channel' }],
  onchange: () => {},
};

function html(props: Record<string, unknown>): string {
  return render(FilterSelect, { props: { ...BASE, ...props } }).body;
}

describe('FilterSelect groups', () => {
  const groups = [
    { label: null, options: [{ value: 'a', label: '#general' }] },
    { label: 'Raids', options: [{ value: 'b', label: '#raid-signups' }] },
  ];

  it('draws a null-label group bare and a named one under an optgroup', () => {
    const body = html({ groups });
    expect(body.match(/<optgroup label="Raids">/g)).toHaveLength(1);
    expect(body.indexOf('#general')).toBeLessThan(body.indexOf('<optgroup'));
    expect(body.indexOf('<optgroup')).toBeLessThan(body.indexOf('#raid-signups'));
  });

  it('draws no optgroup without groups', () => {
    expect(html({})).not.toContain('<optgroup');
  });
});

describe('FilterSelect invalid', () => {
  it('sets aria-invalid and reads the describing sentence', () => {
    const body = html({ invalid: true, describedBy: 'why' });
    expect(body).toContain('aria-invalid="true"');
    expect(body).toContain('aria-describedby="why"');
  });

  it('draws neither when valid and enabled', () => {
    const body = html({ describedBy: 'why' });
    expect(body).not.toContain('aria-invalid');
    expect(body).not.toContain('aria-describedby');
  });

  it('still reads the sentence while disabled', () => {
    expect(html({ disabled: true, describedBy: 'why' })).toContain('aria-describedby="why"');
  });
});
