// web/src/lib/bis/landing-client.ts
// The /bis landing's browser behaviour (design/specs/2026-11-07-bis-landing.md section 4): the
// ControlBar's level range (the URL is the truth; the last pick is also kept per viewer in storage)
// rewrites every spec link's range and the summary line; the head's answer names the header
// selector's character, whose range becomes the default (owner 2026-11-07), whose spec row is
// "Yours" and whose faction rides the links. The page is complete HTML without any of this.
import type { MeCharacter } from '../account/api';
import { characterSpecName } from '../character-spec';
import { markSpecRows } from '../guides/guide-answer-mount';
import { afterLoad, hasSessionHint } from '../page-load';
import { CONTROL_BAR_EVENT, checkRadio, type ControlBarChange } from '../ui/control-bar';
import { bandLabel, parseBandParam } from './landing-bands';
import { specLinkHref } from './landing-links';
import {
  landingPageView,
  type BandChoice,
  type LandingCharacter,
  type LandingPageView,
} from './landing-page-view';
import { specKeyFor, type LandingConfig } from './landing-config';

export const BAND_PARAM = 'band';
const BAND_STORAGE_KEY = 'fs.bisBand';
/** The spec rows' testid prefix: `bis-spec-row-<class>/<spec>` (components/guides/GuideSpecRow.astro). */
export const SPEC_ROW_PREFIX = 'bis-spec-row-';

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

function all<T extends HTMLElement>(selector: string): T[] {
  return [...document.querySelectorAll<T>(selector)];
}

/** Checks the range's radio and the phone bar's line without touching the URL. */
function renderBand(band: number): void {
  const radio = document.querySelector<HTMLElement>(`[data-param="${BAND_PARAM}"] [data-value="${band}"]`);
  if (radio !== null && radio.getAttribute('aria-checked') !== 'true') checkRadio(radio);
  const line = radio
    ?.closest('[data-control-bar]')
    ?.querySelector<HTMLElement>('[data-control-bar-toggle] > span');
  if (line) line.textContent = bandLabel(band);
}

function renderSummary(view: LandingPageView): void {
  const own = document.querySelector<HTMLAnchorElement>('[data-your-range]');
  if (own === null) return;
  const show = view.youBand !== undefined && view.youBand !== view.selectedBand;
  own.hidden = !show;
  if (show && view.youBand !== undefined) {
    own.dataset.band = String(view.youBand);
    own.href = `?${BAND_PARAM}=${view.youBand}`;
    own.textContent = `${own.dataset.label?.replace('{band}', bandLabel(view.youBand)) ?? ''} ›`;
  }
}

function renderSpecLinks(view: LandingPageView): void {
  for (const link of all<HTMLAnchorElement>(`a[data-testid^="${SPEC_ROW_PREFIX}"]`)) {
    const [classSlug = '', specSlug = ''] = (link.dataset.testid ?? '')
      .slice(SPEC_ROW_PREFIX.length)
      .split('/');
    link.href = specLinkHref({ classSlug, specSlug, band: view.selectedBand, faction: view.faction });
  }
}

function landingCharacterOf(config: LandingConfig, character: MeCharacter): LandingCharacter {
  return {
    level: character.level,
    faction: character.faction,
    specKey: specKeyFor(config, character.class, characterSpecName(character)),
  };
}

/** `warrior-fury` -> `warrior/fury`, the spec row's key. */
function rowKey(config: LandingConfig, specKey: string | undefined): string | undefined {
  const spec = specKey === undefined ? undefined : config.specs[specKey];
  return spec === undefined ? undefined : `${spec.classSlug}/${spec.specSlug}`;
}

export function mountBisLanding(config: LandingConfig): void {
  const choice: BandChoice = {
    url: parseBandParam(new URL(window.location.href).searchParams.get(BAND_PARAM)),
    stored: readStoredBand(),
  };
  // A range picked before the bar booted arrives as ?band= (the radio is a link): keep it too.
  if (choice.url !== null) storeBand(choice.url);
  let character: LandingCharacter | null = null;
  const render = (): void => {
    const view = landingPageView(choice, character);
    renderBand(view.selectedBand);
    renderSummary(view);
    renderSpecLinks(view);
    markSpecRows(document, rowKey(config, view.yoursSpecKey), SPEC_ROW_PREFIX);
  };
  document.addEventListener(CONTROL_BAR_EVENT, (event) => {
    const change = (event as CustomEvent<ControlBarChange>).detail;
    if (change.param !== BAND_PARAM) return;
    const band = parseBandParam(change.value);
    if (band === null) return;
    choice.url = band;
    storeBand(band);
    render();
  });
  document.querySelector('[data-your-range]')?.addEventListener('click', (event) => {
    const band = (event.currentTarget as HTMLElement).dataset.band;
    const radio = document.querySelector<HTMLElement>(`[data-param="${BAND_PARAM}"] [data-value="${band}"]`);
    if (radio === null) return;
    event.preventDefault();
    radio.click();
  });
  render();
  if (!hasSessionHint()) return;
  afterLoad(() => {
    void import('./landing-answer-mount').then(({ mountLandingAnswer }) =>
      mountLandingAnswer(config, (next) => {
        character = next === null ? null : landingCharacterOf(config, next);
        render();
      }),
    );
  });
}
