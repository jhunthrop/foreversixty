// web/src/lib/tiers/tier-client.ts
// The tier list page's browser behaviour: the faction choice (the ControlBar writes the URL and
// announces it; this shows that faction's list and keeps it on the role links) and the signed-in
// spec panel in the rail (the visitor's spec, where it stands and its row marked, or the list
// that has it). The lists are static HTML; nothing here is needed to read them.
import type { Faction } from '../bis/types';
import { onCurrentCharacterChange } from '../current-character-events';
import { afterLoad, hasSessionHint } from '../page-load';
import { CONTROL_BAR_EVENT, type ControlBarChange } from '../ui/control-bar';
import { hrefWithFaction } from './tier-controls';
import type { TierSpecView } from './tier-callout';
import type { TierMe } from './tier-me';
import { tiersCopy } from './tier-copy';

interface ClientConfig {
  role: string;
  views: Record<Faction, Record<string, TierSpecView>>;
  factionParam: string;
  factionAttribute: string;
}

/** The spec panel's parts, found once in the server-rendered template (TierSpecCard.astro). */
type CardField = 'name' | 'standing' | 'bis' | 'guide' | 'other' | 'choose' | 'retry';

/** The panel's buttons; each state shows some of them. */
const BUTTONS = ['bis', 'guide', 'other', 'choose', 'retry'] as const satisfies readonly CardField[];
type CardButton = (typeof BUTTONS)[number];

function readConfig(): ClientConfig | null {
  const node = document.getElementById('tier-client-config');
  return node?.textContent ? (JSON.parse(node.textContent) as ClientConfig) : null;
}

function currentFaction(config: ClientConfig): Faction {
  return document.documentElement.getAttribute(config.factionAttribute) === 'horde' ? 'horde' : 'alliance';
}

function checkedFaction(config: ClientConfig): Faction | null {
  const value = document.querySelector<HTMLElement>(
    `[data-param="${config.factionParam}"] [role="radio"][aria-checked="true"]`,
  )?.dataset.value;
  return value === 'horde' || value === 'alliance' ? value : null;
}

function markRows(key: string | undefined): void {
  for (const row of document.querySelectorAll<HTMLElement>('[data-rank-key]')) {
    if (key !== undefined && row.dataset.rankKey === key) row.setAttribute('aria-current', 'true');
    else row.removeAttribute('aria-current');
  }
}

function applyFaction(config: ClientConfig, faction: Faction): void {
  if (faction === 'horde') document.documentElement.setAttribute(config.factionAttribute, 'horde');
  else document.documentElement.removeAttribute(config.factionAttribute);
  for (const radio of document.querySelectorAll<HTMLAnchorElement>('[data-param="role"] [role="radio"]')) {
    const path = radio.getAttribute('href')?.split('?')[0];
    if (path !== undefined) radio.setAttribute('href', hrefWithFaction(path, faction));
  }
}

function fields(slot: HTMLElement): (name: CardField) => HTMLElement {
  return (name) => {
    const found = slot.querySelector<HTMLElement>(`[data-field="${name}"]`);
    if (found === null) throw new Error(`tier spec panel: the template has no ${name}`);
    return found;
  };
}

/** One state of the panel: its heading, its line and the buttons it shows. */
function paint(
  field: (name: CardField) => HTMLElement,
  name: string,
  line: string,
  shown: readonly CardButton[],
) {
  field('name').textContent = name;
  field('standing').textContent = line;
  for (const button of BUTTONS) field(button).hidden = !shown.includes(button);
  // An empty line (the spec on its own list, its row marked) takes no room.
  const row = field('standing').parentElement;
  if (row !== null) row.hidden = line === '' && field('other').hidden;
}

function showSpec(field: (name: CardField) => HTMLElement, view: TierSpecView, config: ClientConfig): void {
  field('bis').setAttribute('href', view.bisHref);
  field('guide').setAttribute('href', view.guideHref);
  field('name').style.color = `var(--color-class-${view.classSlug})`;
  if (view.role === config.role) {
    // The row is marked on the list beside it: its tier, place and figure are not repeated here.
    paint(field, view.fullName, '', ['bis', 'guide']);
    return;
  }
  // On another role's list: where it stands is that list's answer, one tab away.
  field('other').textContent = `${tiersCopy.card.otherRole(view.role)} ›`;
  field('other').setAttribute('href', hrefWithFaction(view.rolePath, currentFaction(config)));
  paint(field, view.fullName, '', ['bis', 'guide', 'other']);
}

function showMessage(field: (name: CardField) => HTMLElement, me: TierMe): void {
  const { card } = tiersCopy;
  field('name').style.color = '';
  if (me.kind === 'error') paint(field, card.heading, card.error, ['retry']);
  else if (me.kind === 'unknown') paint(field, card.heading, card.unknown(me.name), ['choose']);
  else paint(field, card.heading, card.none, ['choose']);
}

/** Paints the spec panel for one read of the account, and marks the visitor's row. */
export function renderAnswer(config: ClientConfig, me: TierMe): void {
  const slot = document.getElementById('tier-answer');
  if (slot === null) return;
  const view = me.kind === 'spec' ? config.views[currentFaction(config)][me.key] : undefined;
  markRows(view !== undefined && view.role === config.role && me.kind === 'spec' ? me.key : undefined);
  const body = slot.querySelector<HTMLElement>('[data-answer-body]');
  if (body === null) return;
  const field = fields(body);
  if (view !== undefined) showSpec(field, view, config);
  else showMessage(field, me);
  const skeleton = slot.querySelector<HTMLElement>('[data-answer-skeleton]');
  if (skeleton !== null) skeleton.hidden = true;
  body.hidden = false;
}

function openCharacterSelector(): void {
  const trigger = [...document.querySelectorAll<HTMLElement>('[data-testid="character-selector"]')].find(
    (node) => node.offsetParent !== null,
  );
  trigger?.click();
}

export function mountTierPage(): void {
  const config = readConfig();
  if (config === null) return;
  let me: TierMe = { kind: 'none' };
  let answered = false;
  // The bar may have booted first and taken a choice already: its checked radio is the truth.
  applyFaction(config, checkedFaction(config) ?? currentFaction(config));
  document.addEventListener(CONTROL_BAR_EVENT, (event) => {
    const change = (event as CustomEvent<ControlBarChange>).detail;
    if (change.param !== config.factionParam) return;
    applyFaction(config, change.value === 'horde' ? 'horde' : 'alliance');
    if (answered) renderAnswer(config, me);
  });
  if (!hasSessionHint()) return;
  // The panel follows the header selector: asked once after load, and again on every pointer
  // change. Only the newest question may paint, so a slow older answer never wins.
  let latest = 0;
  const answer = async (): Promise<void> => {
    const asked = ++latest;
    const { findTierMe } = await import('./tier-me');
    const found = await findTierMe();
    if (asked !== latest) return;
    me = found;
    answered = true;
    renderAnswer(config, me);
  };
  document.querySelector('[data-field="choose"]')?.addEventListener('click', openCharacterSelector);
  document.querySelector('[data-field="retry"]')?.addEventListener('click', () => void answer());
  afterLoad(() => {
    void answer();
    onCurrentCharacterChange(() => void answer());
  });
}
