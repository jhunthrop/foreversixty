// web/src/lib/guides/guide-answer-mount.ts
// Paints the signed-in answer (ui/CharacterAnswer's markup, via GuideAnswer.astro) after load,
// only when a session is hinted, and again on every change of the header selector's character;
// only the newest read paints. Plain DOM: the guide pages carry no island above the fold.
import { fetchMeOnce } from '../account/api';
import { onCurrentCharacterChange, readCurrent } from '../current-character';
import { ANSWER_CREST_PX, paintCrest } from '../class-crest';
import { classColorVar } from '../report/format';
import { afterLoad, hasSessionHint } from '../page-load';
import { answerDetail, guideAnswerFrom, type GuideAnswer, type GuideAnswerPage } from './guide-answer';
import { publishGuideCharacter } from './guide-character';

type Field = 'crest' | 'name' | 'detail' | 'who' | 'next' | 'links' | 'message' | 'action';

function fieldOf(root: HTMLElement): (name: Field) => HTMLElement {
  return (name) => {
    const node = root.querySelector<HTMLElement>(`[data-field="${name}"]`);
    if (node === null) throw new Error(`guide answer: no [data-field="${name}"]`);
    return node;
  };
}

function linkNode(label: string, href: string, testid: string): HTMLAnchorElement {
  const link = document.createElement('a');
  link.href = href;
  link.textContent = label;
  link.dataset.testid = testid;
  link.className = 'guide-answer-link';
  return link;
}

function openCharacterSelector(): void {
  const trigger = [...document.querySelectorAll<HTMLElement>('[data-testid="character-selector"]')].find(
    (node) => node.offsetParent !== null,
  );
  trigger?.click();
}

/** Paints one answer into `root`; returns what the action button should do. */
export function paintGuideAnswer(root: HTMLElement, answer: GuideAnswer, retry: () => void): void {
  const field = fieldOf(root);
  root.dataset.state = answer.state;
  root.hidden = answer.state === 'hidden';
  const body = root.querySelector<HTMLElement>('[data-answer-body]');
  const skeleton = root.querySelector<HTMLElement>('[data-answer-skeleton]');
  if (skeleton !== null) skeleton.hidden = true;
  if (body !== null) body.hidden = false;
  const ready = answer.state === 'ready';
  field('who').hidden = !ready;
  field('links').hidden = !(ready || answer.state === 'none');
  field('next').hidden = !(ready && answer.next !== undefined);
  field('message').hidden = ready || answer.state === 'hidden';
  field('action').hidden = !(answer.state === 'unknown' || answer.state === 'error');
  field('links').replaceChildren();
  if (answer.state === 'ready') {
    const { character } = answer;
    paintCrest(field('crest') as HTMLImageElement, character.className, ANSWER_CREST_PX);
    field('name').textContent = character.name;
    field('name').style.color = classColorVar(character.className);
    field('detail').textContent = answerDetail(character.level, character.identity, character.className);
    field('next').textContent = answer.next ?? '';
    field('links').append(
      ...answer.links.map((link, i) => linkNode(link.label, link.href, `guide-answer-link-${i}`)),
    );
  } else if (answer.state === 'none') {
    field('message').textContent = answer.message;
    field('links').append(linkNode(`${answer.link.label} ›`, answer.link.href, 'guide-answer-link-0'));
  } else if (answer.state === 'unknown') {
    field('message').textContent = answer.message;
    field('action').textContent = `${answer.linkLabel} ›`;
    field('action').onclick = openCharacterSelector;
  } else if (answer.state === 'error') {
    field('message').textContent = answer.message;
    field('action').textContent = answer.retry;
    field('action').onclick = retry;
  }
  if (answer.state === 'error') field('message').setAttribute('role', 'alert');
  else field('message').removeAttribute('role');
}

const SPEC_ROW = 'guide-spec-row-';

/** Marks the header character's spec row on /guides, a class landing and /bis (gold ground, bar,
 *  "Your spec" or "Yours"); every other row is unmarked. `prefix` is the rows' testid before the
 *  `<class>/<spec>` key (components/guides/GuideSpecRow.astro). */
export function markSpecRows(scope: ParentNode, key: string | undefined, prefix: string = SPEC_ROW): void {
  for (const row of scope.querySelectorAll<HTMLElement>(`[data-testid^="${prefix}"]`)) {
    const mine = key !== undefined && row.dataset.testid === `${prefix}${key}`;
    if (mine) row.setAttribute('aria-current', 'true');
    else row.removeAttribute('aria-current');
    const tag = row.querySelector<HTMLElement>('[data-your-spec]');
    if (tag !== null) tag.hidden = !mine;
  }
}

/** Boots the answer in `#guide-answer`: the page's facts come from its `data-page` JSON. */
export function mountGuideAnswer(): void {
  const root = document.getElementById('guide-answer');
  if (root === null) return;
  const page = JSON.parse(root.dataset.page ?? '{}') as GuideAnswerPage;
  if (!hasSessionHint()) {
    root.dataset.state = 'hidden';
    publishGuideCharacter(null);
    return;
  }
  let latest = 0;
  const answer = async (): Promise<void> => {
    const asked = ++latest;
    let found: GuideAnswer;
    try {
      found = guideAnswerFrom(await fetchMeOnce(), readCurrent(), page);
    } catch {
      found = guideAnswerFrom(null, null, page, true);
    }
    if (asked !== latest) return;
    paintGuideAnswer(root, found, () => void answer());
    markSpecRows(
      document,
      found.state === 'ready' ? `${found.character.classSlug}/${found.character.specSlug}` : undefined,
    );
    publishGuideCharacter(found.state === 'ready' ? found.character : null);
  };
  afterLoad(() => {
    void answer();
    onCurrentCharacterChange(() => void answer());
  });
}
