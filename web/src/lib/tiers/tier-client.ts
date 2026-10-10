// web/src/lib/tiers/tier-client.ts
// The tier list page's browser behaviour: the faction pills, the arrow keys on the role tabs
// and the signed-in answer (the callout on the visitor's own role, a pointer on the others,
// and a mark on their row). The list itself is static HTML; nothing here is needed to read it.
import type { Faction } from '../bis/types';
import { onCurrentCharacterChange } from '../current-character';
import { hrefWithFaction, nextTabIndex, parseFaction } from './tier-controls';
import type { TierSpecView } from './tier-callout';
import { tiersCopy } from './tier-copy';

interface ClientConfig {
  role: string;
  views: Record<Faction, Record<string, TierSpecView>>;
  paths: Record<string, string>;
  factionParam: string;
  factionAttribute: string;
}

function readConfig(): ClientConfig | null {
  const node = document.getElementById('tier-client-config');
  return node?.textContent ? (JSON.parse(node.textContent) as ClientConfig) : null;
}

function element<K extends keyof HTMLElementTagNameMap>(
  tag: K,
  className: string,
  text?: string,
): HTMLElementTagNameMap[K] {
  const node = document.createElement(tag);
  node.className = className;
  if (text !== undefined) node.textContent = text;
  return node;
}

function crestImage(view: TierSpecView): HTMLImageElement {
  const img = element('img', 'tier-callout-crest');
  img.src = view.crestSrc;
  img.alt = '';
  img.width = 36;
  img.height = 36;
  img.style.cssText = `width:36px;height:36px;border-radius:999px;object-fit:cover;background:var(--color-raised);box-shadow:0 0 0 2px ${view.colorVar}`;
  return img;
}

/** "Fury Warrior" in the class colour, then the rest of the sentence as plain text. */
function titleNode(view: TierSpecView, tag: 'span'): HTMLElement {
  const node = element(tag, 'tier-callout-title');
  node.style.setProperty('--c', view.colorVar);
  const name = document.createElement('b');
  name.textContent = view.fullName;
  node.append(name, view.title.slice(view.fullName.length));
  return node;
}

function calloutNode(view: TierSpecView): HTMLElement {
  const box = element('div', 'tier-callout');
  box.dataset.testid = 'tier-callout';
  const text = element('span', 'tier-callout-text');
  text.append(titleNode(view, 'span'), element('span', 'tier-callout-detail', view.detail));
  const who = element('span', 'tier-callout-who');
  who.append(crestImage(view), text);
  const button = element('a', 'tier-callout-button', tiersCopy.calloutButton);
  button.setAttribute('href', view.bisHref);
  button.dataset.testid = 'tier-callout-button';
  box.append(who, button);
  return box;
}

function pointerNode(view: TierSpecView, faction: Faction): HTMLElement {
  const link = element('a', 'tier-pointer');
  link.setAttribute('href', hrefWithFaction(view.rolePath, faction));
  link.dataset.testid = 'tier-pointer';
  link.style.setProperty('--c', view.colorVar);
  const label = element('span', '');
  const name = document.createElement('b');
  name.textContent = view.fullName;
  const lead = view.pointer.slice(view.fullName.length, view.pointer.length - tiersCopy.pointerCta.length);
  label.append(name, lead, element('span', 'tier-pointer-cta', tiersCopy.pointerCta));
  link.append(crestImage(view), label);
  return link;
}

function markYourRows(key: string | undefined): void {
  for (const row of document.querySelectorAll<HTMLElement>('.tier-row')) {
    const mine = key !== undefined && row.dataset.specKey === key;
    row.classList.toggle('is-you', mine);
    const slot = row.querySelector<HTMLElement>('[data-you-slot]');
    if (slot === null) continue;
    if (!mine) {
      slot.replaceChildren();
      continue;
    }
    if (slot.childElementCount === 0) {
      const pill = element('span', 'pill pill-site', tiersCopy.yourSpec);
      pill.dataset.testid = 'tier-you-pill';
      slot.append(pill);
    }
  }
}

function currentFaction(config: ClientConfig): Faction {
  return document.documentElement.getAttribute(config.factionAttribute) === 'horde' ? 'horde' : 'alliance';
}

function renderYou(config: ClientConfig, key: string | undefined): void {
  const slot = document.getElementById('tier-you-slot');
  if (slot === null) return;
  const faction = currentFaction(config);
  const view = key === undefined ? undefined : config.views[faction][key];
  if (view === undefined) {
    slot.replaceChildren();
    slot.hidden = true;
    return;
  }
  slot.hidden = false;
  slot.replaceChildren(view.role === config.role ? calloutNode(view) : pointerNode(view, faction));
}

function applyFaction(config: ClientConfig, faction: Faction): void {
  if (faction === 'horde') document.documentElement.setAttribute(config.factionAttribute, 'horde');
  else document.documentElement.removeAttribute(config.factionAttribute);
  for (const pill of document.querySelectorAll<HTMLElement>('[data-faction-pill]')) {
    pill.setAttribute('aria-pressed', String(pill.dataset.factionPill === faction));
  }
  for (const tab of document.querySelectorAll<HTMLAnchorElement>('[data-role-tab]')) {
    tab.setAttribute('href', hrefWithFaction(config.paths[tab.dataset.roleTab ?? ''] ?? '/tiers', faction));
  }
}

/** The address bar follows the pill, so a copied link shows the same list. */
function followInAddressBar(faction: Faction): void {
  const { pathname, hash } = location;
  history.replaceState(null, '', `${hrefWithFaction(pathname, faction)}${hash}`);
}

function wireTabKeys(): void {
  const tabs = [...document.querySelectorAll<HTMLAnchorElement>('[data-role-tab]')];
  for (const [index, tab] of tabs.entries()) {
    tab.addEventListener('keydown', (event) => {
      const next = nextTabIndex(event.key, index, tabs.length);
      if (next === null) return;
      event.preventDefault();
      tabs[next]?.focus();
    });
  }
}

/** `<html data-session="1">`: Base.astro's pre-paint hint that a session cookie exists. */
function hasSessionHint(): boolean {
  return document.documentElement.dataset.session === '1';
}

/** Runs `work` once the page has loaded and the browser is idle, so the signed-in answer never
 *  competes with the first paint. */
function afterLoad(work: () => void): void {
  const schedule = (): void => {
    if ('requestIdleCallback' in window) window.requestIdleCallback(work);
    else setTimeout(work, 0);
  };
  if (document.readyState === 'complete') schedule();
  else window.addEventListener('load', schedule, { once: true });
}

export function mountTierPage(): void {
  const config = readConfig();
  if (config === null) return;
  let characterKey: string | undefined;
  const refresh = (): void => renderYou(config, characterKey);
  applyFaction(config, parseFaction(location.search));
  for (const pill of document.querySelectorAll<HTMLElement>('[data-faction-pill]')) {
    pill.addEventListener('click', () => {
      const faction = pill.dataset.factionPill === 'horde' ? 'horde' : 'alliance';
      applyFaction(config, faction);
      followInAddressBar(faction);
      refresh();
    });
  }
  wireTabKeys();
  if (!hasSessionHint()) return;
  // The answer follows the header selector: asked once after load, and again on every pointer
  // change. Only the newest question may paint, so a slow older answer never wins.
  let latest = 0;
  const answerForSelectedCharacter = async (): Promise<void> => {
    const asked = ++latest;
    const { findCharacterKey } = await import('./tier-me');
    const key = await findCharacterKey();
    if (asked !== latest) return;
    characterKey = key;
    markYourRows(characterKey);
    refresh();
  };
  afterLoad(() => {
    void answerForSelectedCharacter();
    onCurrentCharacterChange(() => void answerForSelectedCharacter());
  });
}
