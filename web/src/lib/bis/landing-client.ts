// web/src/lib/bis/landing-client.ts
// The /bis landing's browser behaviour outside the lead: the band strip (a pill selects the band
// every spec link opens at; the choice is kept in the URL and, per viewer, in storage) and the
// page's answer to the selected character (the "you" pill, the "Yours" row, the faction in the
// links, the grid heading). The page is complete HTML without any of this; the shared character
// header (components/bis/BisLandingHeader.svelte) mounts after the page settles and tells the page
// which character it follows.
import { mount } from 'svelte';
import type { MeCharacter } from '../account/api';
import { nextTabIndex } from '../tiers/tier-controls';
import { afterLoad, hasSessionHint } from '../page-load';
import { bandLabel, parseBandParam } from './landing-bands';
import { landingCopy } from './landing-copy';
import { specLinkHref } from './landing-links';
import { landingPageView, type LandingCharacter, type LandingPageView } from './landing-page-view';
import { specKeyFor, type LandingConfig } from './landing-config';

const BAND_PARAM = 'band';
const BAND_STORAGE_KEY = 'fs.bisBand';

/** What the header drives; the strip drives the rest itself. */
export interface LandingController {
  /** The selected character, or null when there is none (signed out, or nothing selectable). */
  setCharacter(character: LandingCharacter | null): void;
}

function readStoredBand(): number | null {
  try {
    return parseBandParam(window.localStorage.getItem(BAND_STORAGE_KEY));
  } catch {
    return null;
  }
}

function storeBand(band: number): void {
  try {
    window.localStorage.setItem(BAND_STORAGE_KEY, String(band));
  } catch {
    // Private browsing or blocked storage: the choice still holds for this page view.
  }
}

function followInAddressBar(band: number): void {
  const url = new URL(window.location.href);
  url.searchParams.set(BAND_PARAM, String(band));
  window.history.replaceState(null, '', url);
}

function all<T extends HTMLElement>(selector: string): T[] {
  return [...document.querySelectorAll<T>(selector)];
}

function renderPills(view: LandingPageView): void {
  for (const pill of all<HTMLAnchorElement>('[data-band-pill]')) {
    const band = Number(pill.dataset.bandPill);
    const selected = band === view.selectedBand;
    const yours = band === view.youBand;
    pill.setAttribute('aria-checked', String(selected));
    if (selected) pill.setAttribute('aria-current', 'true');
    else pill.removeAttribute('aria-current');
    pill.toggleAttribute('data-you', yours);
    if (yours) pill.setAttribute('aria-label', `${bandLabel(band)}, ${landingCopy.you}`);
    else pill.removeAttribute('aria-label');
    const column = pill.closest<HTMLElement>('[data-band-col]');
    if (column === null) continue;
    if (selected) column.dataset.selected = 'true';
    else delete column.dataset.selected;
  }
}

function renderSpecLinks(view: LandingPageView): void {
  for (const link of all<HTMLAnchorElement>('[data-spec-link]')) {
    link.href = specLinkHref({
      classSlug: link.dataset.classSlug ?? '',
      specSlug: link.dataset.specSlug ?? '',
      band: view.selectedBand,
      faction: view.faction,
    });
    const yours = link.dataset.specKey === view.yoursSpecKey;
    link.classList.toggle('is-you', yours);
    const yoursPill = link.querySelector<HTMLElement>('[data-yours]');
    if (yoursPill !== null) yoursPill.hidden = !yours;
    const roleTag = link.querySelector<HTMLElement>('[data-role-tag]');
    if (roleTag !== null) roleTag.hidden = yours;
  }
}

function renderHeading(view: LandingPageView): void {
  const heading = document.querySelector('[data-grid-heading]');
  if (heading !== null) heading.textContent = view.gridHeading;
}

function wirePillKeys(pills: HTMLAnchorElement[]): void {
  for (const [index, pill] of pills.entries()) {
    pill.addEventListener('keydown', (event) => {
      const next = nextTabIndex(event.key, index, pills.length);
      if (next === null) return;
      event.preventDefault();
      pills[next]?.focus();
    });
  }
}

/** A plain click selects in place; a modified click or a middle click still opens the link. */
function isPlainClick(event: MouseEvent): boolean {
  return event.button === 0 && !event.metaKey && !event.ctrlKey && !event.shiftKey && !event.altKey;
}

const HEADER_SLOT = '[data-bis-header]';

/** Replaces the server-rendered loading header with the live one, which follows the selector. */
function mountHeader(config: LandingConfig, controller: LandingController): void {
  const slot = document.querySelector<HTMLElement>(HEADER_SLOT);
  if (slot === null) return;
  void import('../../components/bis/BisLandingHeader.svelte').then(({ default: BisLandingHeader }) => {
    slot.replaceChildren();
    mount(BisLandingHeader, {
      target: slot,
      props: {
        config,
        onCharacter: (character: MeCharacter | null) =>
          controller.setCharacter(
            character === null
              ? null
              : {
                  level: character.level,
                  faction: character.faction,
                  specKey: specKeyFor(config, character.class, character.spec),
                },
          ),
      },
    });
  });
}

export function mountBisLanding(config: LandingConfig): void {
  let explicitBand =
    parseBandParam(new URL(window.location.href).searchParams.get(BAND_PARAM)) ?? readStoredBand();
  let character: LandingCharacter | null = null;
  // The heading is prerendered for the session hint (two spans the stylesheet picks between), so
  // it is only rewritten once the character lookup has answered either way.
  let headingSettled = false;
  const render = (): void => {
    const view = landingPageView(explicitBand, character);
    renderPills(view);
    renderSpecLinks(view);
    if (headingSettled) renderHeading(view);
  };
  const controller: LandingController = {
    setCharacter(next) {
      character = next;
      headingSettled = true;
      render();
    },
  };
  const pills = all<HTMLAnchorElement>('[data-band-pill]');
  for (const pill of pills) {
    pill.addEventListener('click', (event) => {
      if (!isPlainClick(event)) return;
      event.preventDefault();
      explicitBand = Number(pill.dataset.bandPill);
      storeBand(explicitBand);
      followInAddressBar(explicitBand);
      render();
    });
  }
  wirePillKeys(pills);
  render();
  if (!hasSessionHint()) return;
  afterLoad(() => mountHeader(config, controller));
}
