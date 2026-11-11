// web/src/components/guild/planner/planner-parts.test.ts
// The planner's parts on their own: the comp target form and one sign-up row.
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import type { RosterRow } from '../../../lib/guild/raid-roster';
import { text } from '../../../test-support/raid-fixtures';
import CompTargetForm from './CompTargetForm.svelte';
import PlannerRow from './PlannerRow.svelte';

const yes = (): Promise<boolean> => Promise.resolve(true);
const noop = (): void => {};

function form(props: Record<string, unknown> = {}): string {
  return render(CompTargetForm, {
    props: {
      target: { tanks: 2, healers: 8, melee: 10, ranged: 20, source: 'default' as const },
      onSave: yes,
      onReset: yes,
      onRefuse: noop,
      onDone: noop,
      ...props,
    },
  }).body;
}

describe('CompTargetForm', () => {
  it('starts on the target, one numeric Field per role', () => {
    const body = form();
    for (const [role, value] of [
      ['tanks', 2],
      ['healers', 8],
      ['melee', 10],
      ['ranged', 20],
    ] as const) {
      expect(body).toMatch(
        new RegExp(`inputmode="numeric"[^>]*value="${value}"[^>]*data-testid="raid-comp-${role}"`),
      );
    }
    expect(body).not.toContain('raid-comp-reset');
  });

  it('offers Back to the default once the raid has its own target', () => {
    const body = form({ target: { tanks: 1, healers: 8, melee: 10, ranged: 20, source: 'raid' } });
    expect(text(body)).toContain('Back to the default');
  });

  it('puts a field line under its field and the total under the buttons', () => {
    const body = form({ fields: { healers: 'Use 0 or more.', total: 'Ask for at least one raider.' } });
    expect(body).toContain('id="raid-comp-healers-error"');
    expect(text(body)).toContain('Use 0 or more.');
    expect(body).toMatch(/role="alert" data-testid="raid-comp-total">Ask for at least one raider./);
  });
});

const ROW: RosterRow = {
  character_key: 'us/pvp/ironbrand',
  name: 'Ironbrand',
  class: 'paladin',
  spec: 'Holy',
  role: 'healer',
  group: 'healer',
  spec_key: 'paladin-holy',
  state: 'late',
  note: '',
  own: true,
  also: 'Also: Rootbound (Restoration Druid, Bench)',
};

describe('PlannerRow', () => {
  const row = (props: Record<string, unknown> = {}): string =>
    render(PlannerRow, { props: { row: ROW, roles: true, officer: false, onSetState: yes, ...props } }).body;

  it('reads the name in its class colour, the spec line, the Also line and the Tag', () => {
    const body = row();
    expect(body).toContain('var(--color-class-paladin)');
    expect(text(body)).toBe(
      'Ironbrand Holy Paladin · Healer Also: Rootbound (Restoration Druid, Bench) Late',
    );
    expect(body).toContain('aria-current="true"');
  });

  it('drops the Tag under a group head that already names the state', () => {
    const body = row({ stateNamed: true });
    expect(text(body)).toBe('Ironbrand Holy Paladin · Healer Also: Rootbound (Restoration Druid, Bench)');
    expect(body).not.toContain('data-testid="raid-roster-state"');
  });

  it('gives an officer the state control on its own phone line, and a refusal line', () => {
    const body = row({ officer: true, refusal: 'Did not save.' });
    expect(body).toContain('max-md:col-span-full');
    expect(body).toMatch(/aria-checked="true"[^>]*data-testid="raid-roster-state-us\/pvp\/ironbrand-late"/);
    expect(body).toMatch(/role="alert" data-testid="raid-roster-refusal">Did not save./);
  });
});
