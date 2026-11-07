// web/src/lib/changelog/page-scripts.test.ts
// @vitest-environment jsdom
import { afterEach, describe, expect, it } from 'vitest';
import {
  barSyncScript,
  openTarget,
  SEEN_KEY,
  seenScript,
  SHOW_ATTR,
  SHOW_HEAD_SCRIPT,
  summaries,
  syncBar,
} from './page-scripts';

const COUNTS = { all: 39, game: 25, site: 14 };
const BAR = `<div data-testid="changelog-bar"><button data-control-bar-toggle><span>All changes · 39</span></button>
  <div role="radiogroup" data-param="show">
    <a role="radio" data-value="all" aria-checked="true"></a><a role="radio" data-value="game" aria-checked="false"></a>
  </div></div>`;

function run(script: string): void {
  new Function(script)();
}

afterEach(() => {
  document.body.innerHTML = '';
  document.documentElement.removeAttribute(SHOW_ATTR);
  localStorage.clear();
  window.history.replaceState(null, '', '/');
});

describe('the changelog scripts', () => {
  it('reads ?show= into the html attribute before paint, All for anything else', () => {
    window.history.replaceState(null, '', '/changelog?show=site');
    run(SHOW_HEAD_SCRIPT);
    expect(document.documentElement.getAttribute(SHOW_ATTR)).toBe('site');
    window.history.replaceState(null, '', '/changelog?show=x');
    run(SHOW_HEAD_SCRIPT);
    expect(document.documentElement.getAttribute(SHOW_ATTR)).toBe('all');
  });

  it('checks the chosen segment and writes its phone line, inline and after load alike', () => {
    expect(summaries(COUNTS)).toEqual({
      all: 'All changes · 39',
      game: 'Game changes · 25',
      site: 'Site changes · 14',
    });
    document.body.innerHTML = BAR;
    document.documentElement.setAttribute(SHOW_ATTR, 'game');
    run(barSyncScript(COUNTS));
    const radios = [...document.querySelectorAll('[role="radio"]')].map((r) =>
      r.getAttribute('aria-checked'),
    );
    expect(radios).toEqual(['false', 'true']);
    expect(document.querySelector('button span')?.textContent).toBe('Game changes · 25');
    syncBar(document.body, 'show', 'all', summaries(COUNTS));
    expect(document.querySelector('button span')?.textContent).toBe('All changes · 39');
  });

  it('tags entries newer than the last visit New and stores the newest on leaving', () => {
    document.body.innerHTML = ['2026-11-05', '2026-10-28', '2026-10-20']
      .map(
        (iso) =>
          `<article data-show="site"><time datetime="${iso}"></time><span data-cl-new hidden></span></article>`,
      )
      .join('');
    localStorage.setItem(SEEN_KEY, '2026-10-27');
    run(seenScript('2026-11-05'));
    const shown = [...document.querySelectorAll<HTMLElement>('[data-cl-new]')].map((t) => !t.hidden);
    expect(shown).toEqual([true, true, false]);
    window.dispatchEvent(new Event('pagehide'));
    expect(localStorage.getItem(SEEN_KEY)).toBe('2026-11-05');
  });

  it('tags nothing on a first visit', () => {
    document.body.innerHTML =
      '<article data-show="game"><time datetime="2026-11-05"></time><span data-cl-new hidden></span></article>';
    run(seenScript('2026-11-05'));
    expect(document.querySelector<HTMLElement>('[data-cl-new]')?.hidden).toBe(true);
  });

  it('opens the folded entry a permalink points at', () => {
    document.body.innerHTML = '<details id="2026-10-28-log-coach"><summary>s</summary></details>';
    const scrolled: string[] = [];
    HTMLElement.prototype.scrollIntoView = function (this: HTMLElement) {
      scrolled.push(this.id);
    };
    openTarget('#2026-10-28-log-coach');
    expect(scrolled).toEqual(['2026-10-28-log-coach']);
    expect(document.querySelector('details')?.open).toBe(true);
  });
});
