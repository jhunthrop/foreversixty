// web/src/lib/ui/reading-wrap.ts
// A reading section's lane wraps under its prose once it runs past it (owner 2026-11-09: no dead
// space in the left column beside a tall lane). From 1280 a ReadingSection sets its prose beside a
// 380 px lane; when the lane ends more than WRAP_GAP_PX below the prose, the lane's trailing rows
// move, last first, into the section's `.ui-rd-more` under the prose (columns of ROW_COLUMN_PX),
// until both columns end within a row of each other. Below 1280, on a resize or when a lane swaps
// its content, every row goes back where it was first.
//
// What can move is marked in the markup, never guessed: a `[data-wrap-list]` element's children
// are rows; an element holding a list is opened up the same way, so a block's own label moves
// with its rows only once every row has gone. A moved row keeps its classes; each ancestor up to
// the lane is mirrored by a shallow copy (classes kept, ids and test ids dropped), so the row
// keeps its list's spacing. A `[data-wrap-repeat]` child (a table's header row) is copied into the
// mirror, so a table split in two keeps its heads. Hydrated islands and templates never move.

export const WRAP_FROM_PX = 1280;
/** A left-column gap at least this tall wraps (the owner's ~200 px). */
export const WRAP_GAP_PX = 200;
/** The columns stop within this much of each other (about one row). */
export const WRAP_SLACK_PX = 48;

const LIST = '[data-wrap-list]';
const FIXED = 'astro-island, template';
/** Set on a row while it sits under the prose: it never splits across the columns there. */
const WRAPPED = 'data-wrapped';
const DROPPED_ATTRIBUTES = ['id', 'data-testid', 'data-bis-current', 'data-panel', 'data-region'];

interface Moved {
  node: Element;
  parent: Node;
  next: Node | null;
}

interface WrapState {
  moved: Moved[];
  mirrors: Map<Element, Element>;
}

const states = new WeakMap<HTMLElement, WrapState>();

/** The rows that may move, in document order: the leaves under the lane, opening every element
 *  that is or holds a list. */
export function wrapUnits(lane: Element): Element[] {
  const units: Element[] = [];
  const open = (parent: Element): void => {
    for (const child of parent.children) {
      if (child.matches(FIXED) || child.hasAttribute('data-wrap-repeat')) continue;
      if (child.matches(LIST) || child.querySelector(LIST) !== null) open(child);
      else units.push(child);
    }
  };
  open(lane);
  return units;
}

/** The units that may still move, from the end: a unit that holds an island stops the run. */
export function movableTail(units: readonly Element[]): Element[] {
  const tail: Element[] = [];
  for (let i = units.length - 1; i > 0; i -= 1) {
    const unit = units[i]!;
    if (unit.querySelector(FIXED) !== null) break;
    tail.push(unit);
  }
  return tail;
}

function stripped(element: Element): Element {
  const copy = element.cloneNode(false) as Element;
  for (const name of DROPPED_ATTRIBUTES) copy.removeAttribute(name);
  return copy;
}

function mirrorOf(state: WrapState, lane: Element, more: Element, original: Element): Element {
  if (original === lane) return more;
  const known = state.mirrors.get(original);
  if (known !== undefined) return known;
  const mirror = stripped(original);
  for (const child of original.children) {
    if (child.hasAttribute('data-wrap-repeat')) mirror.append(child.cloneNode(true));
  }
  const parent = original.parentElement;
  if (parent === null) throw new Error('reading-wrap: a row left its lane');
  const outer = mirrorOf(state, lane, more, parent);
  outer.insertBefore(mirror, outer.firstChild);
  state.mirrors.set(original, mirror);
  return mirror;
}

function move(state: WrapState, lane: Element, more: Element, unit: Element): void {
  const parent = unit.parentElement;
  if (parent === null) return;
  const mirror = mirrorOf(state, lane, more, parent);
  state.moved.push({ node: unit, parent, next: unit.nextSibling });
  unit.setAttribute(WRAPPED, '');
  const repeats = [...mirror.children].filter((child) => child.hasAttribute('data-wrap-repeat')).length;
  mirror.insertBefore(unit, mirror.children[repeats] ?? null);
}

function undoLast(state: WrapState): void {
  const last = state.moved.pop();
  if (last === undefined) return;
  last.node.removeAttribute(WRAPPED);
  last.parent.insertBefore(last.node, last.next);
}

interface SectionParts {
  prose: HTMLElement;
  lane: HTMLElement;
  more: HTMLElement;
}

function partsOf(section: HTMLElement): SectionParts | undefined {
  const prose = section.querySelector<HTMLElement>(':scope > .ui-rd-prose');
  const lane = section.querySelector<HTMLElement>(':scope > .ui-rd-lane');
  const more = prose?.querySelector<HTMLElement>(':scope > .ui-rd-more') ?? null;
  return prose && lane && more ? { prose, lane, more } : undefined;
}

/** Every moved row back in its place (a row whose list has gone is dropped), the mirrors cleared. */
export function unwrapSection(section: HTMLElement): void {
  const state = states.get(section);
  const parts = partsOf(section);
  if (state === undefined || parts === undefined) return;
  for (const entry of state.moved.toReversed()) {
    entry.node.removeAttribute(WRAPPED);
    if (entry.parent.isConnected) entry.parent.insertBefore(entry.node, entry.next);
  }
  parts.more.replaceChildren();
  states.delete(section);
}

/** The left-column gap: how far the lane ends below the prose (and what has wrapped under it). */
function gapOf({ prose, lane }: SectionParts): number {
  return lane.getBoundingClientRect().bottom - prose.getBoundingClientRect().bottom;
}

/** Wraps one section's lane under its prose when the gap is WRAP_GAP_PX or more (from 1280). */
export function wrapSection(section: HTMLElement, viewportWidth: number): void {
  unwrapSection(section);
  const parts = partsOf(section);
  if (parts === undefined || viewportWidth < WRAP_FROM_PX) return;
  if (gapOf(parts) < WRAP_GAP_PX) return;
  const state: WrapState = { moved: [], mirrors: new Map() };
  states.set(section, state);
  for (const unit of movableTail(wrapUnits(parts.lane))) {
    const before = gapOf(parts);
    if (before <= WRAP_SLACK_PX) break;
    move(state, parts.lane, parts.more, unit);
    const after = gapOf(parts);
    if (after < -WRAP_SLACK_PX) {
      if (-after > before) undoLast(state);
      break;
    }
  }
  if (state.moved.length === 0) unwrapSection(section);
}

const SECTION = '.ui-rd:not([data-solo])';

/** Re-wraps the section holding `element` around `change` (a lane that swaps its rows). */
export function rewrapAround(element: Element, change: () => void): void {
  const section = element.closest<HTMLElement>(SECTION);
  if (section !== null) unwrapSection(section);
  change();
  if (section !== null) wrapSection(section, window.innerWidth);
}

/** Wraps every reading section on the page now, once the fonts settle and on every width change. */
export function mountReadingWrap(root: ParentNode = document): void {
  const all = (): HTMLElement[] => [...root.querySelectorAll<HTMLElement>(SECTION)];
  const run = (): void => {
    for (const section of all()) wrapSection(section, window.innerWidth);
  };
  run();
  void document.fonts?.ready.then(run);
  let width = window.innerWidth;
  let frame = 0;
  window.addEventListener('resize', () => {
    if (window.innerWidth === width) return;
    width = window.innerWidth;
    cancelAnimationFrame(frame);
    frame = requestAnimationFrame(run);
  });
}
