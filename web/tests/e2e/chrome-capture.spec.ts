// web/tests/e2e/chrome-capture.spec.ts
// Not a test: the tool that keeps the mocks' site header and footer the shipped ones
// (design/parity.md, "The site chrome"). Skipped unless CHROME_CAPTURE=1:
//
//   CHROME_CAPTURE=1 npx playwright test tests/e2e/chrome-capture.spec.ts --project=desktop
//
// For every parity state (and four templates: signed out, signed in, signed in with no
// character, the session loading) it opens the built page as the parity spec does, waits for
// the header's island to hydrate, and writes
// design/mocks/data/chrome/<mock>.json: the <html> attributes and the header and footer
// markup, scripts removed and images inlined. design/mocks/data/chrome/chrome.css holds the
// site CSS that styles them, scoped under .fs-chrome so it never reaches a mock's own blocks.
// The mock library (design/mocks/ws_lib.py) draws its header and footer from these files.
import { mkdirSync, writeFileSync } from 'node:fs';
import path from 'node:path';
import type { Browser, Page } from '@playwright/test';
import { test } from './support/test';
import { openBis } from './support/bis-page';
import { openCharacter, type CharacterState } from './support/character-page';
import { openHome, type HomeState } from './support/home-page';
import { openSimState, type SimParityState } from './support/sim-parity';
import { openLogsState, type LogsParityState } from './support/logs-parity';
import { openTiersState, type TiersParityState } from './support/tiers-parity';
import { openGuildState } from './support/guild-parity';
import { REPO_ROOT, readParityConfig, waitForChromeHydrated } from './support/parity';

const OUT = path.join(REPO_ROOT, 'design/mocks/data/chrome');
const BASE = `http://localhost:${process.env.E2E_PORT ?? '4321'}`;

interface Capture {
  html: Record<string, string>;
  /** <body>'s data attributes: the mock puts them on the [data-fs-body] wrapper. */
  body: Record<string, string>;
  header: string;
  footer: string;
  css: string;
}

/** Runs in the page: the chrome's markup and every CSS rule that styles it, scoped. */
async function capture(page: Page): Promise<Capture> {
  return page.evaluate(async () => {
    const roots = [document.querySelector('[data-site-header]'), document.querySelector('footer')].filter(
      (el): el is HTMLElement => el !== null,
    );
    const SCOPE = '.fs-chrome';
    const pseudo = /(?<!\\)::?[a-zA-Z-]+(?:\((?:[^()]|\([^()]*\))*\))?/g;
    const matches = (selector: string): boolean => {
      const plain = selector.replace(pseudo, '').trim() || '*';
      try {
        return roots.some((root) => root.matches(plain) || root.querySelector(plain) !== null);
      } catch {
        return false;
      }
    };
    const splitList = (list: string): string[] => {
      const out: string[] = [];
      let depth = 0;
      let start = 0;
      for (let i = 0; i < list.length; i += 1) {
        const ch = list[i];
        if (ch === '(' || ch === '[') depth += 1;
        else if (ch === ')' || ch === ']') depth -= 1;
        else if (ch === ',' && depth === 0) {
          out.push(list.slice(start, i).trim());
          start = i + 1;
        }
      }
      out.push(list.slice(start).trim());
      return out;
    };
    const scoped = (selector: string): string | null => {
      const lead = /^((?:html|:root)(?:\[[^\]]*\]|:not\([^)]*\)|:is\([^)]*\)|:where\([^)]*\))*)\s+(.*)$/.exec(
        selector,
      );
      if (lead !== null) return `${lead[1]} ${SCOPE} ${lead[2]}`;
      const onBody = /^body((?:\[[^\]]*\])*)(\s.*)?$/.exec(selector);
      if (onBody !== null) return `${SCOPE} [data-fs-body]${onBody[1]}${onBody[2] ?? ''}`;
      if (/^(html|:root)\b/.test(selector)) return null;
      return `${SCOPE} ${selector}`;
    };
    const customProps = (style: CSSStyleDeclaration): string => {
      const out: string[] = [];
      for (let i = 0; i < style.length; i += 1) {
        const name = style[i];
        if (name.startsWith('--')) out.push(`${name}:${style.getPropertyValue(name)}`);
      }
      return out.join(';');
    };
    /** A kept rule, its nested rules flattened (Tailwind's variants nest a media query in the
     *  utility): `&` takes the scoped selector, a nested media query wraps it. */
    const nested = (selector: string, rule: CSSStyleRule | CSSMediaRule): string => {
      let out =
        rule instanceof CSSStyleRule && rule.style.length > 0 ? `${selector}{${rule.style.cssText}}\n` : '';
      for (const child of Array.from(rule.cssRules)) {
        if (child instanceof CSSMediaRule) {
          const inner = nested(selector, child).replace(/\n/g, '');
          if (inner !== '') out += `@media ${child.conditionText}{${inner}}\n`;
        } else if (child instanceof CSSStyleRule) {
          const sub = splitList(child.selectorText)
            .map((part) =>
              part.includes('&') ? part.replace(/&/g, `:is(${selector})`) : `:is(${selector}) ${part}`,
            )
            .join(',');
          out += nested(sub, child);
        } else if ('style' in child) {
          out += `${selector}{${(child as CSSStyleRule).style.cssText}}\n`;
        }
      }
      return out;
    };
    const walk = (rules: CSSRuleList): string => {
      let css = '';
      for (const rule of Array.from(rules)) {
        if (rule instanceof CSSStyleRule) {
          const list = splitList(rule.selectorText);
          if (list.every((s) => /^(html|:root|:host)(\[[^\]]*\])*$/.test(s))) {
            const props = customProps(rule.style);
            if (props !== '') css += `${rule.selectorText}{${props}}\n`;
            continue;
          }
          const kept = list
            .filter(matches)
            .map(scoped)
            .filter((s): s is string => s !== null);
          if (kept.length > 0) css += nested(kept.join(','), rule);
        } else if (rule instanceof CSSMediaRule) {
          const inner = walk(rule.cssRules);
          if (inner !== '') css += `@media ${rule.conditionText}{${inner.replace(/\n/g, '')}}\n`;
        } else if (rule instanceof CSSSupportsRule) {
          const inner = walk(rule.cssRules);
          if (inner !== '') css += `@supports ${rule.conditionText}{${inner.replace(/\n/g, '')}}\n`;
        } else if (rule instanceof CSSLayerBlockRule || rule instanceof CSSContainerRule) {
          css += walk(rule.cssRules);
        } else if (rule instanceof CSSKeyframesRule || rule instanceof CSSPropertyRule) {
          css += `${rule.cssText.replace(/\n/g, ' ')}\n`;
        }
      }
      return css;
    };
    // What the chrome inherits from <body> on the site (the mock's body is its own).
    const body = getComputedStyle(document.body);
    // Unitless, as the site sets it: a px value would be inherited as is by smaller text.
    const lineHeight =
      body.lineHeight === 'normal'
        ? 'normal'
        : String(parseFloat(body.lineHeight) / parseFloat(body.fontSize));
    let css = `${SCOPE}{font-family:${body.fontFamily};font-size:${body.fontSize};line-height:${lineHeight};color:${body.color};letter-spacing:${body.letterSpacing}}\n`;
    for (const sheet of Array.from(document.styleSheets)) {
      try {
        css += walk(sheet.cssRules);
      } catch {
        // A cross-origin sheet (none today) has no readable rules.
      }
    }
    const ISLAND_ATTRS = [
      'component-url',
      'component-export',
      'renderer-url',
      'props',
      'ssr',
      'client',
      'opts',
      'before-hydration-url',
      'await-children',
      'uid',
    ];
    const clean = async (root: HTMLElement): Promise<string> => {
      const copy = root.cloneNode(true) as HTMLElement;
      copy.querySelectorAll('script, template').forEach((el) => el.remove());
      copy
        .querySelectorAll('astro-island')
        .forEach((el) => ISLAND_ATTRS.forEach((a) => el.removeAttribute(a)));
      for (const img of Array.from(copy.querySelectorAll('img'))) {
        const src = img.getAttribute('src');
        // The site's own files, from design/mocks/web/ (a mock opens as file://).
        if (src !== null && src.startsWith('/')) img.setAttribute('src', `../../../web/public${src}`);
        img.removeAttribute('srcset');
        img.setAttribute('loading', 'eager');
      }
      copy.querySelectorAll('a[href]').forEach((a) => a.setAttribute('href', '#'));
      return copy.outerHTML;
    };
    const dataAttrs = (el: HTMLElement): Record<string, string> =>
      Object.fromEntries(
        Array.from(el.attributes)
          .filter((a) => a.name.startsWith('data-'))
          .map((a) => [a.name, a.value]),
      );
    return {
      html: dataAttrs(document.documentElement),
      body: dataAttrs(document.body),
      header: await clean(roots[0]),
      footer: await clean(roots[1]),
      css,
    };
  });
}

/** The CSS of every capture, in first-seen order: pages load different bundles. */
const CSS_RULES = new Set<string>();

function write(name: string, value: Capture): void {
  mkdirSync(OUT, { recursive: true });
  for (const rule of value.css.split('\n')) if (rule.trim() !== '') CSS_RULES.add(rule);
  writeFileSync(path.join(OUT, 'chrome.css'), `${[...CSS_RULES].join('\n')}\n`);
  const { html, body, header, footer } = value;
  writeFileSync(
    path.join(OUT, `${name.replace(/\.html$/, '')}.json`),
    `${JSON.stringify({ html, body, header, footer }, null, 1)}\n`,
  );
}

const TEMPLATES: Record<string, HomeState> = {
  'template-out': {},
  'template-in': { me: 'me-34.json', feed: 'feed-34.json', home: 'home-34.json' },
  'template-new': { me: 'me-new.json' },
  'template-pending': { me: 'pending' },
};

test.describe('site chrome capture', () => {
  test.skip(process.env.CHROME_CAPTURE !== '1', 'a tool, run by hand: CHROME_CAPTURE=1');
  test.skip(({ isMobile }) => isMobile, 'captured at 1440');
  test.use({ viewport: { width: 1440, height: 1000 } });

  /** Each state in a fresh context: a remembered character must never leak into the next. */
  async function each<T>(
    browser: Browser,
    states: [string, T][],
    open: (page: Page, state: T) => Promise<void>,
  ) {
    for (const [name, state] of states) {
      const context = await browser.newContext({ viewport: { width: 1440, height: 1000 }, baseURL: BASE });
      const page = await context.newPage();
      await open(page, state);
      // The header as a reader sees it after load: the selector island hydrated, as the parity
      // specs wait for it on the build side (support/parity.ts waitForChromeHydrated).
      await waitForChromeHydrated(page);
      // And wired (Header.astro's after-paint script, data-header-wired), so every capture holds the
      // same header attributes whatever the state's timing.
      await page.locator('[data-site-header][data-header-wired]').waitFor({ state: 'attached' });
      write(name, await capture(page));
      await context.close();
    }
  }

  test('templates and home', async ({ browser }) => {
    const home = readParityConfig('home-parity.json').states as (HomeState & { mock: string })[];
    await each(
      browser,
      [...Object.entries(TEMPLATES), ...home.map((s): [string, HomeState] => [s.mock, s])],
      openHome,
    );
  });

  test('best in slot', async ({ browser }) => {
    const states = readParityConfig('bis-parity.json').states.map((s): [string, string] => [s.mock, s.path]);
    await each(browser, states, openBis);
  });

  test('simulator', async ({ browser }) => {
    const states = readParityConfig('sim-parity.json').states.map((s): [string, SimParityState] => [
      s.mock,
      s.id as SimParityState,
    ]);
    await each(browser, states, openSimState);
  });

  test('logs, reports, rankings and ratings', async ({ browser }) => {
    const states = readParityConfig('logs-parity.json').states.map((s): [string, LogsParityState] => [
      s.mock,
      s.id as LogsParityState,
    ]);
    await each(browser, states, openLogsState);
  });

  test('tier list', async ({ browser }) => {
    const states = readParityConfig('tiers-parity.json').states.map((s): [string, TiersParityState] => [
      s.mock,
      s.id as TiersParityState,
    ]);
    await each(browser, states, openTiersState);
  });

  test('guild', async ({ browser }) => {
    const states = readParityConfig('guild-parity.json').states.map((s): [string, string] => [s.mock, s.id]);
    await each(browser, states, openGuildState);
  });

  test('character', async ({ browser }) => {
    const states = readParityConfig('character-parity.json').states as (CharacterState & { mock: string })[];
    await each(
      browser,
      states.map((s): [string, CharacterState] => [s.mock, s]),
      openCharacter,
    );
  });
});
