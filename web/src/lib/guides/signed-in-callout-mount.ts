// web/src/lib/guides/signed-in-callout-mount.ts
// The /guides index's signed-in callout, drawn into its slot and kept in step with the header
// character selector: the shared `/v1/me` read is awaited once per draw, and every pointer
// change draws again. A static page with no island, so this is plain DOM.
import { fetchMeOnce } from '../account/api';
import { selectedCharacter } from '../account/hero-character';
import { onCurrentCharacterChange, readCurrent } from '../current-character';
import { signedInCalloutView, type SignedInCalloutView } from './signed-in-callout';

const CREST_SIZE_PX = 36;

function calloutLink(view: SignedInCalloutView): HTMLAnchorElement {
  const link = document.createElement('a');
  link.className = 'guides-callout';
  link.href = view.href;
  link.dataset.testid = 'guides-signed-in-callout';
  const img = document.createElement('img');
  img.src = view.crestSrc;
  img.alt = '';
  img.width = CREST_SIZE_PX;
  img.height = CREST_SIZE_PX;
  img.className = 'guides-callout-crest';
  img.style.setProperty('--c', view.crestColorVar);
  const text = document.createElement('span');
  text.className = 'guides-callout-text';
  text.textContent = view.text;
  link.append(img, text);
  return link;
}

/** The view for the character the header selector shows: the pointer, else the account main
 *  (owner-reported defect 2026-10-07: `characters[0]` named an alt while a warrior was
 *  selected). */
async function currentView(): Promise<SignedInCalloutView | undefined> {
  const me = await fetchMeOnce().catch(() => null);
  return signedInCalloutView(me === null ? undefined : (selectedCharacter(readCurrent(), me) ?? undefined));
}

/** Draws the callout into `root` now and after every pointer change; a character with no
 *  known class or spec, or a signed-out visitor, leaves the slot empty and hidden. Only the
 *  newest read paints. Returns the unsubscribe. */
export function mountSignedInCallout(root: HTMLElement): () => void {
  let latest = 0;
  const draw = async (): Promise<void> => {
    const asked = ++latest;
    const view = await currentView();
    if (asked !== latest) return;
    root.hidden = view === undefined;
    root.replaceChildren(...(view === undefined ? [] : [calloutLink(view)]));
  };
  void draw();
  return onCurrentCharacterChange(() => void draw());
}
