// @vitest-environment jsdom
// web/src/lib/character-search-dom.test.ts
// Fix round 2 items 4 and 7, round 3 item 3: the lookup forms' shared browser code -- the
// "Look up another character" prefill (name, region, ruleset, then focus), and the check that
// keeps "/" from stealing a key someone is typing.
import { describe, expect, it } from 'vitest';
import { focusIfIdle, isTypingTarget, LOOKUP_NAME_PARAM, prefillFromQuery } from './character-search-dom';
import { scoreCopy } from './score/copy';

function form(): HTMLFormElement {
  document.body.innerHTML = [
    '<form>',
    '<input name="name" type="search" />',
    '<select name="region"><option value="us" selected>US</option><option value="eu">EU</option></select>',
    '<select name="ruleset"><option value="normal" selected>Normal</option><option value="hardcore">Hardcore</option></select>',
    '</form>',
  ].join('');
  return document.querySelector('form') as HTMLFormElement;
}

const field = (f: HTMLFormElement, selector: string): HTMLInputElement | HTMLSelectElement =>
  f.querySelector(selector) as HTMLInputElement | HTMLSelectElement;

describe('prefillFromQuery', () => {
  it('fills the name, region and ruleset a "Look up another character" link carries, and focuses the box', () => {
    const href = scoreCopy.lookupHref('Elyra Duskvale', 'eu', 'hardcore');
    expect(href).toBe('/?character=Elyra%20Duskvale&region=eu&ruleset=hardcore#lookup');
    const f = form();
    prefillFromQuery(f, new URL(href, 'https://example.test').search);
    expect(field(f, 'input').value).toBe('Elyra Duskvale');
    expect(field(f, 'select[name="region"]').value).toBe('eu');
    expect(field(f, 'select[name="ruleset"]').value).toBe('hardcore');
    expect(document.activeElement).toBe(field(f, 'input'));
  });

  it('ignores a region or ruleset the site does not know, keeping the defaults', () => {
    const f = form();
    prefillFromQuery(f, '?character=Elyra&region=moon&ruleset=chaos');
    expect(field(f, 'select[name="region"]').value).toBe('us');
    expect(field(f, 'select[name="ruleset"]').value).toBe('normal');
  });

  it('leaves the box empty with no name, and never submits', () => {
    const f = form();
    let submitted = false;
    f.addEventListener('submit', () => (submitted = true));
    prefillFromQuery(f, '');
    prefillFromQuery(f, `?${LOOKUP_NAME_PARAM}=%20`);
    expect(field(f, 'input').value).toBe('');
    expect(submitted).toBe(false);
  });
});

describe('focusIfIdle (fix round 4 item 1)', () => {
  it('takes the cursor when nothing has it (focus on the page itself)', () => {
    const f = form();
    (document.activeElement as HTMLElement | null)?.blur();
    expect(document.activeElement).toBe(document.body);
    focusIfIdle(field(f, 'input'));
    expect(document.activeElement).toBe(field(f, 'input'));
  });

  it('never steals focus from a field the player is already in', () => {
    const f = form();
    const other = document.createElement('input');
    document.body.append(other);
    other.focus();
    focusIfIdle(field(f, 'input'));
    expect(document.activeElement).toBe(other);
  });

  it('never steals focus from a button the player tabbed to', () => {
    const f = form();
    const button = document.createElement('button');
    document.body.append(button);
    button.focus();
    prefillFromQuery(f, '?character=Elyra');
    expect(field(f, 'input').value).toBe('Elyra');
    expect(document.activeElement).toBe(button);
  });
});

describe('isTypingTarget', () => {
  it('is true in a field and false elsewhere', () => {
    document.body.innerHTML =
      '<input id="a" /><textarea id="b"></textarea><select id="c"></select><p id="d"></p>';
    for (const id of ['a', 'b', 'c']) expect(isTypingTarget(document.getElementById(id))).toBe(true);
    expect(isTypingTarget(document.getElementById('d'))).toBe(false);
    expect(isTypingTarget(null)).toBe(false);
  });
});
