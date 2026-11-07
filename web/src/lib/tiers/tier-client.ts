// web/src/lib/tiers/tier-client.ts
// The tier list page's browser behaviour: the faction choice (the ControlBar writes the URL and
// announces it; this shows that faction's list and keeps it on the role links) and the signed-in
// answer in the PageHeader (the visitor's spec, its place and its row marked, or a pointer to
// the role list that has it). The lists are static HTML; nothing here is needed to read them.
import type { Faction } from '../bis/types';
import { onCurrentCharacterChange } from '../current-character-events';
import { afterLoad, hasSessionHint } from '../page-load';
import { writeFigure } from '../ui/figure-dom';
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

/** The answer slot's parts, found once in the server-rendered template (TierPage.astro). */
type AnswerField =
  | 'label'
  | 'crest'
  | 'name'
  | 'detail'
  | 'place'
  | 'figure'
  | 'links'
  | 'bis'
  | 'guide'
  | 'other'
  | 'message'
  | 'choose'
  | 'retry';

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
    const mine = key !== undefined && row.dataset.rankKey === key;
    if (mine) row.setAttribute('aria-current', 'true');
    else row.removeAttribute('aria-current');
    const tag = row.querySelector<HTMLElement>('[data-you-tag]');
    if (tag !== null) tag.hidden = !mine;
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

function fields(slot: HTMLElement): (name: AnswerField) => HTMLElement {
  return (name) => {
    const found = slot.querySelector<HTMLElement>(`[data-field="${name}"]`);
    if (found === null) throw new Error(`tier answer: the template has no ${name}`);
    return found;
  };
}

function showSpec(field: (name: AnswerField) => HTMLElement, view: TierSpecView, config: ClientConfig): void {
  const ownRole = view.role === config.role;
  const faction = currentFaction(config);
  field('label').textContent = tiersCopy.answer.label(view.role);
  const crest = field('crest') as HTMLImageElement;
  crest.src = view.crestSrc;
  crest.style.boxShadow = `0 0 0 2px ${view.colorVar}`;
  crest.hidden = false;
  field('name').textContent = view.fullName;
  field('name').style.color = view.colorVar;
  field('detail').textContent = ownRole ? view.detail : tiersCopy.answer.otherRole(view.fullName, view.role);
  field('figure').hidden = !ownRole || view.place === '';
  writeFigure(field('figure'), view.place, view.placeUnit);
  field('links').hidden = !ownRole;
  field('bis').setAttribute('href', view.bisHref);
  field('guide').setAttribute('href', view.guideHref);
  field('other').hidden = ownRole;
  field('other').setAttribute('href', hrefWithFaction(view.rolePath, faction));
  field('message').hidden = true;
}

function showMessage(
  field: (name: AnswerField) => HTMLElement,
  text: string,
  action: 'choose' | 'retry',
): void {
  for (const name of ['crest', 'figure', 'links', 'other', 'choose', 'retry'] as const)
    field(name).hidden = true;
  field('label').textContent = '';
  field('name').textContent = '';
  field('detail').textContent = '';
  field('message').hidden = false;
  field('message').firstElementChild!.textContent = text;
  field(action).hidden = false;
}

/** Paints the answer slot for one read of the account, and marks the visitor's row. */
export function renderAnswer(config: ClientConfig, me: TierMe): void {
  const slot = document.getElementById('tier-answer');
  if (slot === null) return;
  const body = slot.querySelector<HTMLElement>('[data-answer-body]');
  const skeleton = slot.querySelector<HTMLElement>('[data-answer-skeleton]');
  const view = me.kind === 'spec' ? config.views[currentFaction(config)][me.key] : undefined;
  markRows(view !== undefined && view.role === config.role && me.kind === 'spec' ? me.key : undefined);
  if (body === null) return;
  if (me.kind === 'none' || (me.kind === 'spec' && view === undefined)) {
    slot.hidden = true;
    return;
  }
  slot.hidden = false;
  const field = fields(body);
  if (me.kind === 'spec' && view !== undefined) showSpec(field, view, config);
  else if (me.kind === 'unknown') showMessage(field, tiersCopy.answer.unknown(me.name), 'choose');
  else showMessage(field, tiersCopy.answer.error, 'retry');
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
  // The bar may have booted first and taken a choice already: its checked radio is the truth.
  applyFaction(config, checkedFaction(config) ?? currentFaction(config));
  document.addEventListener(CONTROL_BAR_EVENT, (event) => {
    const change = (event as CustomEvent<ControlBarChange>).detail;
    if (change.param !== config.factionParam) return;
    applyFaction(config, change.value === 'horde' ? 'horde' : 'alliance');
    renderAnswer(config, me);
  });
  if (!hasSessionHint()) return;
  // The answer follows the header selector: asked once after load, and again on every pointer
  // change. Only the newest question may paint, so a slow older answer never wins.
  let latest = 0;
  const answer = async (): Promise<void> => {
    const asked = ++latest;
    const { findTierMe } = await import('./tier-me');
    const found = await findTierMe();
    if (asked !== latest) return;
    me = found;
    renderAnswer(config, me);
  };
  document.querySelector('[data-field="choose"]')?.addEventListener('click', openCharacterSelector);
  document.querySelector('[data-field="retry"]')?.addEventListener('click', () => void answer());
  afterLoad(() => {
    void answer();
    onCurrentCharacterChange(() => void answer());
  });
}
