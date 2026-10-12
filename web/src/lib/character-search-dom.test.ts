// @vitest-environment jsdom
// web/src/lib/character-search-dom.test.ts
// Fix round 2 items 4 and 7: the lookup forms' shared browser code -- the "Look up another
// character" prefill, and the check that keeps "/" from stealing a key someone is typing.
import { describe, expect, it } from 'vitest';
import { isTypingTarget, LOOKUP_NAME_PARAM, prefillFromQuery } from './character-search-dom';
import { scoreCopy } from './score/copy';

function form(): HTMLFormElement {
  document.body.innerHTML = '<form><input name="name" type="search" /></form>';
  return document.querySelector('form') as HTMLFormElement;
}

describe('prefillFromQuery', () => {
  it('fills the name a "Look up another character" link carries', () => {
    const href = scoreCopy.lookupHref('Elyra Duskvale');
    expect(href).toBe('/?character=Elyra%20Duskvale#lookup');
    const f = form();
    prefillFromQuery(f, new URL(href, 'https://example.test').search);
    expect((f.querySelector('input') as HTMLInputElement).value).toBe('Elyra Duskvale');
  });

  it('leaves the box empty with no name, and never submits', () => {
    const f = form();
    let submitted = false;
    f.addEventListener('submit', () => (submitted = true));
    prefillFromQuery(f, '');
    prefillFromQuery(f, `?${LOOKUP_NAME_PARAM}=%20`);
    expect((f.querySelector('input') as HTMLInputElement).value).toBe('');
    expect(submitted).toBe(false);
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
