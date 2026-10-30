// web/src/components/ui/Select.test.ts
// The shared dropdown's server markup (design/specs/2026-10-30-controls-charts.md §A3 to §A5):
// the native select carrying every option's data, the face over it, the sizes and the states.
// The enhancer's behaviour on this markup is lib/ui/listbox.test.ts.
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import { selectCopy } from '../../lib/ui/copy';
import Select from './Select.svelte';

const OPTIONS = [
  { value: 'us', label: 'US' },
  { value: 'eu', label: 'EU' },
];

function select(props: Record<string, unknown> = {}): string {
  return render(Select, { props: { id: 'region', options: OPTIONS, label: 'Region', ...props } }).body;
}

describe('Select', () => {
  it('is the native select under a styled face, marked for the listbox', () => {
    const body = select({ value: 'eu' });
    expect(body).toMatch(/<div[^>]*class="sel"[^>]*data-listbox/);
    expect(body).toMatch(/<select class="sel-native"[^>]*id="region"/);
    expect(body).toContain('<option value="us"');
    expect(body).toMatch(/<span class="sel-face" aria-hidden="true">/);
    expect(body).toMatch(/class="sel-label"[^>]*>EU</);
    expect(body).toContain('data-testid="select-chevron"');
  });

  it('takes the three sizes', () => {
    expect(select()).toContain('data-size="default"');
    expect(select({ size: 'compact' })).toContain('data-size="compact"');
    expect(select({ size: 'dense' })).toContain('data-size="dense"');
  });

  it('names itself with aria-label', () => {
    expect(select()).toContain('aria-label="Region"');
  });

  it('carries each option’s data for the enhancer, and groups as optgroups', () => {
    const body = select({
      value: 'm',
      options: [
        {
          value: 'm',
          label: 'Morrigan',
          group: 'Damage',
          colour: '--color-class-mage',
          icon: { kind: 'crest', src: '/crest/mage.png' },
          meta: '612',
          you: true,
          description: 'Fire',
        },
        { value: 's', label: 'Sword', icon: { kind: 'item', src: '/i.png', quality: 'epic' } },
      ],
    });
    expect(body).toMatch(/<optgroup label="Damage">(<!--[^>]*-->)*<option value="m"/);
    expect(body).toContain('data-icon="/crest/mage.png"');
    expect(body).toContain('data-icon-kind="crest"');
    expect(body).toContain('data-colour="--color-class-mage"');
    expect(body).toContain('data-meta="612"');
    expect(body).toContain('data-you="true"');
    expect(body).toContain('data-description="Fire"');
    expect(body).toContain('data-quality="epic"');
    // The closed field shows the picked option's icon and class colour.
    expect(body).toMatch(/class="sel-icon" data-kind="crest"/);
    expect(body).toMatch(/class="sel-label"[^>]*style="color: var\(--color-class-mage\);"/);
  });

  it('placeholder: a first empty option, muted while nothing is chosen', () => {
    const body = select({ placeholder: 'Choose a region', value: '' });
    expect(body).toMatch(/<option value=""[^>]*>Choose a region/);
    expect(body).toMatch(/class="sel-label is-placeholder"/);
  });

  it('disabled: a real disabled select', () => {
    const body = select({ disabled: true });
    expect(body).toContain('disabled=""');
    expect(body).toContain('data-disabled="true"');
  });

  it('error: invalid, described by the field error line', () => {
    const body = select({ error: 'Pick a region.' });
    expect(body).toContain('data-error="true"');
    expect(body).toContain('aria-invalid="true"');
    expect(body).toContain('aria-describedby="region-error"');
    expect(body).toMatch(/id="region-error"[^>]*>Pick a region\./);
  });

  it('loading: the busy label, a spinner for the chevron, and it does not open', () => {
    const body = select({ loading: true });
    expect(body).toContain(`>${selectCopy.loading}<`);
    expect(body).toContain('data-testid="spinner"');
    expect(body).not.toContain('data-testid="select-chevron"');
    expect(body).toContain('aria-busy="true"');
    expect(body).toContain('disabled=""');
  });

  it('empty: the call site’s empty label, no chevron, disabled', () => {
    const body = select({ options: [], emptyLabel: selectCopy.noEnemies });
    expect(body).toContain(`>${selectCopy.noEnemies}<`);
    expect(body).not.toContain('data-testid="select-chevron"');
    expect(body).toContain('data-empty="true"');
  });

  it('searchable and right-slot fields tell the enhancer', () => {
    const body = select({ searchable: true, searchPlaceholder: 'Filter players', align: 'end' });
    expect(body).toContain('data-searchable="true"');
    expect(body).toContain('data-search-placeholder="Filter players"');
    expect(body).toContain('data-align="end"');
  });

  it('a searchable field without its placeholder fails fast', () => {
    expect(() => select({ searchable: true })).toThrow(/searchPlaceholder/);
  });

  it('multiple: a hidden multi select, the closed summary, and a sheet button for a coarse pointer', () => {
    const body = select({
      multiple: true,
      values: ['us'],
      fieldLabel: 'Top 8 of 40',
      footLabel: 'Top 8',
      footCount: '9 drawn',
      testid: 'players',
    });
    expect(body).toMatch(/<select class="sel-native"[^>]*multiple[^>]*hidden/);
    expect(body).toMatch(/<option value="us" selected/);
    expect(body).toContain('data-multiple="true"');
    expect(body).toContain('data-foot-label="Top 8"');
    expect(body).toContain('data-testid="players-sheet"');
    expect(body).toMatch(/class="sel-label"[^>]*>Top 8 of 40</);
  });

  it('draws no error line at rest', () => {
    const body = select();
    expect(body).not.toContain('aria-invalid');
    expect(body).not.toContain('region-error');
  });
});
