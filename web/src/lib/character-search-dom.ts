// web/src/lib/character-search-dom.ts
// The few lines of browser code behind every character lookup form (Task L8.2 fix round 2
// item 7): the home page's full box (CharacterSearch.astro) and the compact one in the site
// header (HeaderSearch.astro). A form marked `data-character-search` routes its name -- and
// its region and ruleset, where it has those fields -- through lib/character-search.ts. The
// "/" key belongs to the header (it is on every page): it focuses the home page's box where
// there is one, and otherwise opens the header's own.
import { characterSearchHref, DEFAULT_REGION, DEFAULT_RULESET } from './character-search';
import { isRegion, isRuleset } from './characters';
import { wireTypeahead } from './character-typeahead-dom';

/** The query parameter a "Look up another character" link fills the home box with. */
export const LOOKUP_NAME_PARAM = 'character';

/** Everything a lookup form does: the submit, and the suggestion list under its name box. */
export function wireLookupForm(form: HTMLFormElement): void {
  wireSearchForm(form);
  wireTypeahead(form);
}

/** Routes a lookup form's submit to the profile, or says what is wrong in its status line. */
export function wireSearchForm(form: HTMLFormElement): void {
  if (form.dataset.searchWired === '1') return;
  form.dataset.searchWired = '1';
  form.addEventListener('submit', (event) => {
    event.preventDefault();
    const fields = new FormData(form);
    const name = String(fields.get('name') ?? '');
    const region = String(fields.get('region') ?? DEFAULT_REGION);
    const ruleset = String(fields.get('ruleset') ?? DEFAULT_RULESET);
    const href = isRegion(region) && isRuleset(ruleset) ? characterSearchHref(name, region, ruleset) : null;
    if (href !== null) {
      window.location.assign(href);
      return;
    }
    const status = document.getElementById(form.dataset.statusId ?? '');
    if (status !== null)
      status.textContent = (name.trim() === '' ? form.dataset.empty : form.dataset.invalid) ?? '';
  });
}

/**
 * Fills a lookup form from `?character=&region=&ruleset=` -- the name, and the region and
 * ruleset selects where the form has them and the values are ones the site knows (fix round 3
 * item 3) -- then puts the cursor in the name box. Never submits: the player checks and presses.
 */
export function prefillFromQuery(form: HTMLFormElement, search: string = window.location.search): void {
  const params = new URLSearchParams(search);
  const name = params.get(LOOKUP_NAME_PARAM);
  const input = form.querySelector<HTMLInputElement>('input[name="name"]');
  if (name === null || input === null || name.trim() === '') return;
  input.value = name.slice(0, 64);
  const region = params.get('region');
  const regionSelect = form.querySelector<HTMLSelectElement>('select[name="region"]');
  if (region !== null && isRegion(region) && regionSelect !== null) regionSelect.value = region;
  const ruleset = params.get('ruleset');
  const rulesetSelect = form.querySelector<HTMLSelectElement>('select[name="ruleset"]');
  if (ruleset !== null && isRuleset(ruleset) && rulesetSelect !== null) rulesetSelect.value = ruleset;
  // The `#lookup` fragment scroll runs after this module and moves focus back to the page, so a
  // focus taken before `load` does not stick: wait for it, then a frame.
  if (document.readyState === 'complete') {
    focusIfIdle(input);
  } else {
    window.addEventListener('load', () => requestAnimationFrame(() => focusIfIdle(input)), { once: true });
  }
}

/**
 * Puts the cursor in `input` only while nothing else has it (fix round 4 item 1): focus on the
 * page itself (the body) or nowhere. Anything else -- a field the player has started typing in
 * while the page loaded, a button they tabbed to -- keeps its focus.
 */
export function focusIfIdle(input: HTMLElement): void {
  const active = document.activeElement;
  if (active !== null && active !== document.body) return;
  if (isTypingTarget(active)) return;
  input.focus();
}

/** Whether a key press came from somewhere typing is expected. */
export function isTypingTarget(target: EventTarget | null): boolean {
  return (
    target instanceof HTMLElement &&
    (target.isContentEditable || ['INPUT', 'TEXTAREA', 'SELECT'].includes(target.tagName))
  );
}
