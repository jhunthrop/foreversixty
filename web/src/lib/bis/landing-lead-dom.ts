// web/src/lib/bis/landing-lead-dom.ts
// Writes a `LeadView` into the lead's prerendered markup (components/bis/BisLead.astro). Every
// value goes in through textContent or an attribute, never as HTML. The markup carries each
// variant of the count column; this shows the one the view names.
import { factionLogoSrc, factionName } from '../faction-mark';
import type { CountView, LeadView } from './landing-lead';

/** The one element carrying `data-lead-<name>`; a missing one is a template bug, not a state. */
function field(root: HTMLElement, name: string): HTMLElement {
  const node = root.querySelector<HTMLElement>(`[data-lead-${name}]`);
  if (node === null) throw new Error(`BisLead markup has no [data-lead-${name}]`);
  return node;
}

function setPartsHidden(root: HTMLElement, name: string, hidden: boolean): void {
  for (const node of root.querySelectorAll<HTMLElement>(`[data-lead-${name}]`)) node.hidden = hidden;
}

function showCrest(root: HTMLElement, classSlug: string): void {
  for (const slot of root.querySelectorAll<HTMLElement>('[data-lead-crest]')) {
    slot.hidden = slot.dataset.leadCrest !== classSlug;
  }
}

function writeDescriptor(root: HTMLElement, view: LeadView): void {
  field(root, 'spec-class').textContent = view.specClass;
  const hasLevel = view.level !== undefined;
  setPartsHidden(root, 'level-part', !hasLevel);
  if (hasLevel) field(root, 'level').textContent = String(view.level);
  setPartsHidden(root, 'faction-part', view.faction === undefined);
  if (view.faction !== undefined) {
    const logo = field(root, 'faction-logo');
    if (logo instanceof HTMLImageElement) logo.src = factionLogoSrc(view.faction);
    field(root, 'faction').textContent = factionName(view.faction);
  }
}

function writeSync(root: HTMLElement, view: LeadView): void {
  const sync = field(root, 'sync');
  sync.hidden = view.sync === null;
  if (view.sync === null) return;
  field(root, 'sync-text').textContent = view.sync.text;
  const stale = field(root, 'sync-stale');
  stale.hidden = view.sync.staleTail === null;
  field(root, 'sync-stale-text').textContent = view.sync.staleTail ?? '';
}

function writeCount(root: HTMLElement, count: CountView): void {
  field(root, 'count').hidden = count.kind !== 'count';
  field(root, 'nosync').hidden = count.kind !== 'nosync';
  field(root, 'under20').hidden = count.kind !== 'under20';
  if (count.kind !== 'count') return;
  field(root, 'count-label').textContent = count.label;
  field(root, 'worn').textContent = String(count.worn);
  field(root, 'of').textContent = ` of ${count.total} slots`;
  field(root, 'fill').style.width = `${count.percent.toFixed(1)}%`;
  const note = field(root, 'note');
  note.hidden = count.note === null;
  note.textContent = count.note ?? '';
}

/** Shows the lead's content for `view`, replacing the skeleton in the same box. */
export function renderLead(root: HTMLElement, view: LeadView): void {
  field(root, 'skeleton').hidden = true;
  field(root, 'body').hidden = false;
  showCrest(root, view.classSlug);
  const name = field(root, 'name');
  name.textContent = view.name;
  name.style.color = view.classColor;
  field(root, 'link').setAttribute('href', view.href);
  field(root, 'cta').setAttribute('href', view.href);
  field(root, 'cta').textContent = view.buttonLabel;
  writeDescriptor(root, view);
  writeSync(root, view);
  writeCount(root, view.count);
}

/** No character to speak to: the lead is omitted entirely, and nothing says "failed". */
export function hideLead(root: HTMLElement): void {
  root.hidden = true;
}

export function showLeadShell(root: HTMLElement): void {
  root.hidden = false;
}
