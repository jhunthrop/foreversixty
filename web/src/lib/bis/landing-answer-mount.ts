// web/src/lib/bis/landing-answer-mount.ts
// Paints the /bis landing's head answer (ui/CharacterAnswer with its figure, via
// components/bis/BisLandingAnswer.astro) after load, only when a session is hinted, and again on
// every change of the header selector's character; only the newest read paints. It also mounts the
// finished server sim line under the head once, and tells the page which character it follows.
// Plain DOM: the landing carries no island above the fold.
import { mount, unmount } from 'svelte';
import { fetchMeOnce, type Me, type MeCharacter } from '../account/api';
import { onCurrentCharacterChange, readCurrent } from '../current-character';
import { ANSWER_CREST_PX, paintCrest } from '../class-crest';
import { planAnswer } from '../character-header/answer-plan';
import { finishedRunsFor } from '../character-header/finished-sim-line';
import { classColorVar } from '../report/format';
import { afterLoad } from '../page-load';
import { button } from '../ui/button';
import { LINK } from '../ui/primitives';
import { writeFigure, writeMeter } from '../ui/figure-dom';
import { landingAction, landingAnswerSpec } from './landing-answer';
import { landingCopy } from './landing-copy';
import {
  emptyAnswer,
  landingAnswerCharacter,
  readyAnswer,
  type AnswerLink,
  type LandingAnswer,
} from './landing-answer-view';
import type { LandingConfig } from './landing-config';

export const ANSWER_ID = 'bis-answer';
const FINISHED_SLOT = '[data-bis-finished]';

type Field = 'crest' | 'name' | 'detail' | 'who' | 'figure' | 'note' | 'message' | 'links' | 'action';

function fieldOf(root: HTMLElement): (name: Field) => HTMLElement {
  return (name) => {
    const node = root.querySelector<HTMLElement>(`[data-field="${name}"]`);
    if (node === null) throw new Error(`bis answer: no [data-field="${name}"]`);
    return node;
  };
}

function anchor(link: AnswerLink, className: string, testid: string): HTMLAnchorElement {
  const node = document.createElement('a');
  node.href = link.href;
  node.textContent = link.label;
  node.className = className;
  node.dataset.testid = testid;
  return node;
}

function paintCharacter(
  field: (name: Field) => HTMLElement,
  answer: Extract<LandingAnswer, { state: 'ready' }>,
): void {
  const { character } = answer;
  paintCrest(field('crest') as HTMLImageElement, character.className, ANSWER_CREST_PX);
  field('name').textContent = character.name;
  field('name').style.color = classColorVar(character.className);
  field('detail').textContent = character.line;
}

/** Paints one answer into `root`. */
export function paintLandingAnswer(root: HTMLElement, answer: LandingAnswer, retry: () => void): void {
  const field = fieldOf(root);
  root.dataset.state = answer.state;
  root.querySelector<HTMLElement>('[data-answer-skeleton]')?.setAttribute('hidden', '');
  root.querySelector<HTMLElement>('[data-answer-body]')?.removeAttribute('hidden');
  const ready = answer.state === 'ready' ? answer : null;
  const body = ready?.body;
  field('who').hidden = ready === null;
  field('figure').hidden = body?.kind !== 'count';
  const message =
    answer.state === 'none' || answer.state === 'error'
      ? answer.message
      : body !== undefined && body.kind !== 'count'
        ? body.text
        : null;
  field('message').hidden = message === null;
  field('message').textContent = message ?? '';
  const failed = answer.state === 'error' || body?.kind === 'failed';
  if (failed) field('message').setAttribute('role', 'alert');
  else field('message').removeAttribute('role');
  field('action').hidden = !failed;
  field('action').textContent =
    answer.state === 'error' ? answer.retry : body?.kind === 'failed' ? body.retry : '';
  field('action').onclick = failed ? retry : null;
  const links: HTMLAnchorElement[] = [];
  if (answer.state === 'none')
    links.push(anchor({ ...answer.link, label: `${answer.link.label} ›` }, LINK, 'bis-answer-link'));
  if (body?.kind === 'sentence' && body.link !== undefined) {
    links.push(anchor({ ...body.link, label: `${body.link.label} ›` }, LINK, 'bis-answer-link'));
  }
  if (ready !== null) {
    paintCharacter(field, ready);
    links.push(anchor(ready.action, button('primary', 'md'), 'bis-answer-button'));
  }
  field('links').replaceChildren(...links);
  field('links').hidden = links.length === 0;
  if (body?.kind === 'count') {
    writeFigure(field('figure'), String(body.worn), body.unit);
    writeMeter(field('figure'), body.bar);
    // The head's compact contract (ui/PageHeader): one short line under the figure. After an old
    // sync it says the count is from that sync (it changes what the number means), else it names
    // the slots the range has no pick for; the caption and the note are the figure's tooltip. A
    // short line takes no full stop, as the head's other short lines.
    const stale = body.caption !== landingCopy.countLabel;
    const line = (stale ? body.caption : body.note)?.replace(/\.$/, '');
    field('note').textContent = line ?? '';
    field('note').hidden = line === undefined;
    field('figure').title = [body.caption, body.note].filter((part) => part !== undefined).join('. ');
  }
}

/** The answer for the selected character: its plan, the count when the plan fetches one, the action. */
async function answerFor(character: MeCharacter, config: LandingConfig): Promise<LandingAnswer> {
  const spec = landingAnswerSpec(config);
  const plan = planAnswer(character, spec);
  const fetched = plan === 'load' ? await spec.load(character).catch(() => 'failed' as const) : undefined;
  return readyAnswer({ character, plan, fetched, action: landingAction(character, config), now: new Date() });
}

/** Mounts the finished server sim line once (components/character/FinishedSimLine.svelte). */
function mountFinished(me: Me | null): void {
  const slot = document.querySelector<HTMLElement>(FINISHED_SLOT);
  const runs = finishedRunsFor(me);
  if (slot === null || runs === null || slot.dataset.mounted === 'true') return;
  slot.dataset.mounted = 'true';
  void import('../../components/character/FinishedSimLine.svelte').then(({ default: FinishedSimLine }) => {
    const line = mount(FinishedSimLine, {
      target: slot,
      props: {
        runs,
        onDismissed: () => {
          void unmount(line);
          delete document.documentElement.dataset.finished;
          document.querySelector<HTMLElement>('[data-testid="bis-answer-button"]')?.focus();
        },
      },
    });
  });
}

/**
 * Boots the answer in `#bis-answer` (call it only for a hinted session). `onCharacter` hears the
 * character the answer names, or null when it names none.
 */
export function mountLandingAnswer(
  config: LandingConfig,
  onCharacter: (character: MeCharacter | null) => void,
): void {
  const root = document.getElementById(ANSWER_ID);
  if (root === null) return;
  let latest = 0;
  const answer = async (): Promise<void> => {
    const asked = ++latest;
    let me: Me | null;
    try {
      me = await fetchMeOnce();
    } catch {
      if (asked === latest) paintLandingAnswer(root, emptyAnswer(null, true), () => void answer());
      onCharacter(null);
      return;
    }
    mountFinished(me);
    const character = landingAnswerCharacter(me, readCurrent());
    const found = character === null ? emptyAnswer(me, false) : await answerFor(character, config);
    if (asked !== latest) return;
    paintLandingAnswer(root, found, () => void answer());
    onCharacter(character);
  };
  afterLoad(() => {
    void answer();
    onCurrentCharacterChange(() => void answer());
  });
}
