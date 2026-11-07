// web/src/lib/guides/guide-answer-mount.test.ts
// @vitest-environment jsdom
// Paints the answer into GuideAnswer.astro's own markup (read from the file, so a renamed field
// fails this test), marks the header character's spec row, and boots from the page's facts:
// the header selector's character, only the newest read painted.
import { readFileSync } from 'node:fs';
import path from 'node:path';
import { afterEach, beforeEach, describe, expect, it, vi } from 'vitest';
import type { Me, MeCharacter } from '../account/api';
import { pointerForCharacter } from '../account/main-character';
import { CURRENT_CHARACTER_CHANGED, writeCurrent } from '../current-character';
import type { GuideAnswer as Answer, GuideCharacter } from './guide-answer';
import { onGuideCharacter, resetGuideCharacter } from './guide-character';
import { markSpecRows, mountGuideAnswer, paintGuideAnswer } from './guide-answer-mount';

let session: Me | null = null;
vi.mock('../account/api', () => ({ fetchMeOnce: () => Promise.resolve(session) }));

/** GuideAnswer.astro's own markup: its body with the frontmatter, the script and the style cut
 *  and its two expressions filled (an Astro component cannot render in the jsdom environment; vitest runs from web/). */
const ANSWER_SOURCE = readFileSync(path.resolve('src/components/guides/GuideAnswer.astro'), 'utf8');
const PAGE = { classSlug: 'hunter', specSlug: 'survival', ladder: [{ level: 36, name: 'Expose Prey' }] };
const markup = ANSWER_SOURCE.split('---')[2]!
  .split('<script>')[0]!
  .replace('{guidesCopy.answerLabel}', 'Your character')
  .replace('data-page={JSON.stringify(page)}', `data-page='${JSON.stringify(PAGE)}'`)
  .replace('class={LINK}', '');

const BOW: GuideCharacter = {
  name: 'Bow',
  classSlug: 'hunter',
  specSlug: 'survival',
  className: 'Hunter',
  specName: 'Survival',
  level: 34,
  faction: 'alliance',
  levelLine: 'Level 34',
  identity: 'Dwarf Survival Hunter',
};

function mountMarkup(): HTMLElement {
  document.body.innerHTML = markup;
  return document.getElementById('guide-answer')!;
}

const field = (root: HTMLElement, name: string): HTMLElement =>
  root.querySelector<HTMLElement>(`[data-field="${name}"]`)!;

afterEach(() => {
  document.body.innerHTML = '';
});

describe('paintGuideAnswer', () => {
  it('paints a ready answer: crest, name in the class colour, the detail line, next and links', () => {
    const root = mountMarkup();
    const answer: Answer = {
      state: 'ready',
      character: BOW,
      next: 'Next point at 36: Expose Prey',
      links: [
        { label: 'Talents at level 34 ›', href: '#leveling' },
        { label: 'Your best in slot ›', href: '/bis/hunter/survival?band=30' },
      ],
    };
    paintGuideAnswer(root, answer, () => undefined);
    expect(root.dataset.state).toBe('ready');
    expect(root.hidden).toBe(false);
    expect(root.querySelector<HTMLElement>('[data-answer-skeleton]')!.hidden).toBe(true);
    expect(root.querySelector<HTMLElement>('[data-answer-body]')!.hidden).toBe(false);
    expect((field(root, 'crest') as HTMLImageElement).getAttribute('src')).toBe(
      '/icons/hd/crests/hunter.webp',
    );
    expect(field(root, 'name').textContent).toBe('Bow');
    expect(field(root, 'name').style.color).toContain('--color-class-hunter');
    expect(field(root, 'detail').textContent).toBe('Level 34 · Dwarf Survival Hunter');
    expect(field(root, 'next').hidden).toBe(false);
    expect(field(root, 'next').textContent).toBe('Next point at 36: Expose Prey');
    const links = [...field(root, 'links').querySelectorAll('a')];
    expect(links.map((link) => [link.textContent, link.getAttribute('href'), link.dataset.testid])).toEqual([
      ['Talents at level 34 ›', '#leveling', 'guide-answer-link-0'],
      ['Your best in slot ›', '/bis/hunter/survival?band=30', 'guide-answer-link-1'],
    ]);
    expect(field(root, 'message').hidden).toBe(true);
    expect(field(root, 'action').hidden).toBe(true);
  });

  it('replaces the links on a repaint and hides next when the answer has none', () => {
    const root = mountMarkup();
    const links = [{ label: 'Your guide: Survival Hunter ›', href: '/guides/hunter/survival' }];
    paintGuideAnswer(
      root,
      { state: 'ready', character: BOW, next: 'x', links: [...links, ...links] },
      () => undefined,
    );
    paintGuideAnswer(root, { state: 'ready', character: BOW, links }, () => undefined);
    expect(field(root, 'links').querySelectorAll('a')).toHaveLength(1);
    expect(field(root, 'next').hidden).toBe(true);
  });

  it('paints no character: the message and the setup link, no who', () => {
    const root = mountMarkup();
    paintGuideAnswer(
      root,
      {
        state: 'none',
        message: 'No character on this account yet.',
        link: { label: 'Get set up', href: '/setup' },
      },
      () => undefined,
    );
    expect(field(root, 'who').hidden).toBe(true);
    expect(field(root, 'message').textContent).toBe('No character on this account yet.');
    expect(field(root, 'links').querySelector('a')?.getAttribute('href')).toBe('/setup');
    expect(field(root, 'links').querySelector('a')?.textContent).toBe('Get set up ›');
    expect(field(root, 'action').hidden).toBe(true);
  });

  it('paints an error as an alert whose action retries', () => {
    const root = mountMarkup();
    const retry = vi.fn();
    paintGuideAnswer(
      root,
      { state: 'error', message: 'We could not load your character.', retry: 'Try again' },
      retry,
    );
    expect(field(root, 'message').getAttribute('role')).toBe('alert');
    expect(field(root, 'action').hidden).toBe(false);
    expect(field(root, 'action').textContent).toBe('Try again');
    field(root, 'action').click();
    expect(retry).toHaveBeenCalledOnce();
  });

  it('paints an unknown spec with the selector action, and drops the alert role', () => {
    const root = mountMarkup();
    paintGuideAnswer(root, { state: 'error', message: 'x', retry: 'Try again' }, () => undefined);
    paintGuideAnswer(
      root,
      { state: 'unknown', message: "Bow's spec is not known yet.", linkLabel: 'Choose a character' },
      () => undefined,
    );
    expect(field(root, 'message').hasAttribute('role')).toBe(false);
    expect(field(root, 'action').textContent).toBe('Choose a character ›');
    expect(field(root, 'links').hidden).toBe(true);
  });

  it('hides the whole answer when hidden', () => {
    const root = mountMarkup();
    paintGuideAnswer(root, { state: 'hidden' }, () => undefined);
    expect(root.hidden).toBe(true);
    expect(root.dataset.state).toBe('hidden');
  });

  it('throws on markup missing a field, naming it', () => {
    const root = document.createElement('div');
    expect(() => paintGuideAnswer(root, { state: 'hidden' }, () => undefined)).toThrow('[data-field="who"]');
  });
});

describe('markSpecRows', () => {
  function rows(): HTMLElement {
    const scope = document.createElement('div');
    scope.innerHTML = ['hunter/survival', 'hunter/marksmanship']
      .map((key) => `<a data-testid="guide-spec-row-${key}"><span data-your-spec hidden>Your spec</span></a>`)
      .join('');
    return scope;
  }
  const row = (scope: HTMLElement, key: string): HTMLElement =>
    scope.querySelector<HTMLElement>(`[data-testid="guide-spec-row-${key}"]`)!;

  it('marks the character’s spec row current with its tag shown, and no other', () => {
    const scope = rows();
    markSpecRows(scope, 'hunter/survival');
    expect(row(scope, 'hunter/survival').getAttribute('aria-current')).toBe('true');
    expect(row(scope, 'hunter/survival').querySelector<HTMLElement>('[data-your-spec]')!.hidden).toBe(false);
    expect(row(scope, 'hunter/marksmanship').hasAttribute('aria-current')).toBe(false);
    expect(row(scope, 'hunter/marksmanship').querySelector<HTMLElement>('[data-your-spec]')!.hidden).toBe(
      true,
    );
  });

  it('unmarks every row when the answer names nobody', () => {
    const scope = rows();
    markSpecRows(scope, 'hunter/survival');
    markSpecRows(scope, undefined);
    expect(scope.querySelectorAll('[aria-current]')).toHaveLength(0);
    expect([...scope.querySelectorAll<HTMLElement>('[data-your-spec]')].every((tag) => tag.hidden)).toBe(
      true,
    );
  });
});

describe('mountGuideAnswer', () => {
  function character(name: string, className: string, spec: string, level: number): MeCharacter {
    return {
      key: `us/normal/${name.toLowerCase()}`,
      region: 'us',
      ruleset: 'normal',
      name,
      class: className,
      spec,
      level,
    };
  }
  const armsy = character('Armsy', 'Warrior', 'Arms', 60);
  const bow = character('Bow', 'Hunter', 'Survival', 34);
  const seen: (GuideCharacter | null)[] = [];

  beforeEach(() => {
    vi.useFakeTimers();
    localStorage.clear();
    seen.length = 0;
    session = {
      user: { id: 1, battletag: null, email: null, role: 'user', anonymize: false },
      characters: [armsy, bow],
      guilds: [],
    } as Me;
    onGuideCharacter((found) => seen.push(found));
  });
  afterEach(() => {
    delete document.documentElement.dataset.session;
    resetGuideCharacter();
    vi.useRealTimers();
  });

  it('stays hidden and publishes nobody with no session hint', () => {
    const root = mountMarkup();
    mountGuideAnswer();
    expect(root.dataset.state).toBe('hidden');
    expect(seen).toEqual([null]);
  });

  it('answers for the header selector’s character and follows its changes (bug B1)', async () => {
    document.documentElement.dataset.session = '1';
    writeCurrent(pointerForCharacter(bow));
    const root = mountMarkup();
    mountGuideAnswer();
    await vi.advanceTimersByTimeAsync(5);
    expect(field(root, 'name').textContent).toBe('Bow');
    expect(field(root, 'next').textContent).toBe('Next point at 36: Expose Prey');
    expect(seen.at(-1)?.name).toBe('Bow');

    writeCurrent(pointerForCharacter(armsy));
    window.dispatchEvent(new Event(CURRENT_CHARACTER_CHANGED));
    await vi.advanceTimersByTimeAsync(5);
    expect(field(root, 'name').textContent).toBe('Armsy');
    expect(field(root, 'links').textContent).toContain('Your guide: Arms Warrior');
    expect(seen.at(-1)?.name).toBe('Armsy');
  });
});
