// web/src/lib/guides/bis-lane.ts
// The guide's Best in slot lane follows the answer's character (guides spec 2026-11-08 §4.3,
// G5; bug B2: any character the header selector names, a Battle.net one included): a character
// of this class sees its own band and faction, anyone else the guide's level 60 picks. Every
// band is in the page as a <template>; only the chosen one is in the document.
import { specLinkHref } from '../bis/landing-links';
import { readCurrent } from '../current-character';
import type { GuideCharacter } from './guide-answer';
import { bandForLevel, onGuideCharacter } from './guide-character';

/** A stored addon or code build's band (its spent points + 9 is its level), for a visitor whose
 *  pointer names no account character (a pasted export): undefined for another class or a
 *  pointer that is not an FS1 code. */
export function ownBand(ref: string, classSlug: string): number | undefined {
  const parts = ref.split(':');
  if (parts[0] !== 'FS1' || parts[2] !== classSlug) return undefined;
  let points = 0;
  for (const tree of (parts[4] ?? '').split('|')[0]!.split('/')) {
    for (const digit of tree) {
      const value = Number(digit);
      if (Number.isInteger(value)) points += value;
    }
  }
  return bandForLevel(Math.min(60, points + 9));
}

/** The panel key a character reads on a guide of `classSlug`: `<band>-<faction>`. */
export function bisPanelKey(
  character: GuideCharacter | null,
  classSlug: string,
  defaultFaction: string,
): string {
  if (character === null || character.classSlug !== classSlug) return `60-${defaultFaction}`;
  return `${bandForLevel(character.level)}-${character.faction ?? defaultFaction}`;
}

/** Swaps the lane's panel to the character's whenever the answer settles. */
export function mountBisLane(root: HTMLElement): void {
  const classSlug = root.dataset.guideClass ?? '';
  const defaultFaction = root.dataset.defaultFaction ?? 'alliance';
  const current = root.querySelector<HTMLElement>('[data-bis-current]');
  if (current === null) return;
  // The band the server drew has no template of its own: keep a copy, so the lane can return to it.
  const first = current.dataset.panel ?? '';
  if (root.querySelector(`template[data-bis-panel="${first}"]`) === null) {
    const keep = document.createElement('template');
    keep.dataset.bisPanel = first;
    keep.content.append(...[...current.childNodes].map((node) => node.cloneNode(true)));
    root.append(keep);
  }
  onGuideCharacter((character) => {
    const pointer = character === null ? readCurrent() : null;
    const pasted =
      pointer !== null && (pointer.source === 'addon' || pointer.source === 'code')
        ? ownBand(pointer.ref, classSlug)
        : undefined;
    const key =
      pasted === undefined
        ? bisPanelKey(character, classSlug, defaultFaction)
        : `${pasted}-${defaultFaction}`;
    if (current.dataset.panel === key) return;
    const template = root.querySelector<HTMLTemplateElement>(`template[data-bis-panel="${key}"]`);
    if (template === null) return;
    current.replaceChildren(template.content.cloneNode(true));
    current.dataset.panel = key;
    retargetFullList(root, classSlug, key);
  });
}

/** The full-list link follows the shown band: `/bis/<class>/<spec>?faction=<f>#band-<f>-<band>`. */
export function retargetFullList(root: HTMLElement, classSlug: string, key: string): void {
  const link = root.querySelector<HTMLAnchorElement>('[data-testid="guide-bis-full-list"]');
  const specSlug = root.dataset.specSlug;
  const [band, faction] = key.split('-');
  if (link === null || specSlug === undefined || (faction !== 'alliance' && faction !== 'horde')) return;
  link.href = specLinkHref({ classSlug, specSlug, band: Number(band), faction });
}
