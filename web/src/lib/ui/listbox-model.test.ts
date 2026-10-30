// web/src/lib/ui/listbox-model.test.ts
// The shared dropdown's pure rules (design/specs/2026-10-30-controls-charts.md §A4, §A6).
import { describe, expect, it } from 'vitest';
import {
  FLIP_THRESHOLD_PX,
  chooseSide,
  createTypeahead,
  edgeRow,
  filterRows,
  foldText,
  initialActive,
  labelColour,
  listMaxHeight,
  matchSpan,
  moveActive,
  typeaheadMatch,
  type ListRow,
} from './listbox-model';

const row = (label: string, extra: Partial<ListRow> = {}): ListRow => ({
  value: label.toLowerCase(),
  label,
  selected: false,
  ...extra,
});

const ROWS: ListRow[] = [
  row('Garrosk'),
  row('Velyn', { disabled: true }),
  row('Morrigan', { selected: true }),
  row('Ysolde'),
  row('Gorehowl'),
  row('Kael', { disabled: true }),
];

describe('moveActive', () => {
  it('moves by one and skips disabled rows', () => {
    expect(moveActive(ROWS, 0, 1)).toBe(2);
    expect(moveActive(ROWS, 2, -1)).toBe(0);
  });

  it('never wraps: the last enabled row stays put going down', () => {
    expect(moveActive(ROWS, 4, 1)).toBe(4);
    expect(moveActive(ROWS, 0, -1)).toBe(0);
  });

  it('pages by ten, clamped to the last enabled row', () => {
    expect(moveActive(ROWS, 0, 10)).toBe(4);
    expect(moveActive(ROWS, 4, -10)).toBe(0);
  });
});

describe('edgeRow and initialActive', () => {
  it('finds the first and last enabled rows', () => {
    expect(edgeRow(ROWS, 'first')).toBe(0);
    expect(edgeRow(ROWS, 'last')).toBe(4);
  });

  it('opens on the picked row, else the first enabled one', () => {
    expect(initialActive(ROWS)).toBe(2);
    expect(initialActive(ROWS.map((r) => ({ ...r, selected: false })))).toBe(0);
  });
});

describe('typeahead', () => {
  it('jumps to the first enabled row starting with the run', () => {
    expect(typeaheadMatch(ROWS, 'y', 0)).toBe(3);
    expect(typeaheadMatch(ROWS, 'mor', 0)).toBe(2);
  });

  it('cycles through rows sharing a repeated first letter', () => {
    expect(typeaheadMatch(ROWS, 'g', 0)).toBe(4);
    expect(typeaheadMatch(ROWS, 'gg', 4)).toBe(0);
  });

  it('skips disabled rows and finds nothing for an unknown run', () => {
    expect(typeaheadMatch(ROWS, 'k', 0)).toBe(-1);
    expect(typeaheadMatch(ROWS, 'zz', 0)).toBe(-1);
  });

  it('keeps a run inside 500 ms and starts a new one after', () => {
    let now = 0;
    const buffer = createTypeahead(() => now);
    expect(buffer.push('m')).toBe('m');
    now = 300;
    expect(buffer.push('o')).toBe('mo');
    now = 900;
    expect(buffer.push('y')).toBe('y');
  });
});

describe('filter', () => {
  it('matches a case- and accent-insensitive substring of the label', () => {
    const rows = [row('Blackwing Lair'), row('Lúcio'), row('Onyxia')];
    expect(filterRows(rows, 'BLACK').map((r) => r.label)).toEqual(['Blackwing Lair']);
    expect(filterRows(rows, 'luc').map((r) => r.label)).toEqual(['Lúcio']);
    expect(filterRows(rows, '  ')).toHaveLength(3);
    expect(foldText('Ýsolde')).toBe('ysolde');
  });

  it('marks the matched run', () => {
    expect(matchSpan('Lower Blackrock Spire', 'black')).toEqual({ start: 6, end: 11 });
    expect(matchSpan('Onyxia', 'naxx')).toBeUndefined();
  });
});

describe('placement', () => {
  it('opens below unless less than 240 px fits there and more fits above', () => {
    expect(chooseSide(400, 100, 300)).toBe('bottom');
    expect(chooseSide(FLIP_THRESHOLD_PX - 1, 600, 300)).toBe('top');
    expect(chooseSide(100, 80, 300)).toBe('bottom');
  });

  it('a short list that fits below never flips', () => {
    expect(chooseSide(120, 600, 100)).toBe('bottom');
  });

  it('caps the height at 360 px and at the room less 16 px', () => {
    expect(listMaxHeight(1000)).toBe(360);
    expect(listMaxHeight(200)).toBe(184);
    expect(listMaxHeight(4)).toBe(0);
  });
});

describe('labelColour', () => {
  it('takes a class token and refuses anything that is not a token', () => {
    expect(labelColour({ colour: '--color-class-mage' })).toBe('var(--color-class-mage)');
    expect(labelColour({ colour: 'red; background: url(x)' })).toBeUndefined();
  });

  it('an item takes its quality text colour', () => {
    expect(labelColour({ icon: { kind: 'item', src: 'x', quality: 'epic' } })).toBe(
      'var(--color-rarity-epic-text)',
    );
    expect(labelColour({ icon: { kind: 'item', src: 'x', quality: 'uncommon' } })).toBe(
      'var(--color-rarity-uncommon)',
    );
  });
});
