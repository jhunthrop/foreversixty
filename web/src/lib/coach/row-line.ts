// web/src/lib/coach/row-line.ts
// What the row strip says for a settled `line=1` answer (experience spec 3.2 and 8.2). Pure, so
// the strip and its tests share one decision. `null` renders nothing. The strip never collapses:
// every settled answer but "not reviewable" has a line, and the slot is reserved at its tallest.
import { coachCopy, fill } from './copy';
import type { CoachFull, CoachTeaser } from './types';

/**
 * The strip's reserved height, known from `me` before any answer (spec 3.2 and 7): a Premium
 * owner's one 52 px strip on phone, a free owner's two 44 px targets; 36 px on desktop for both.
 */
export const STRIP_RESERVE = {
  premium: 'min-h-[52px] md:min-h-9',
  free: 'min-h-[88px] md:min-h-9',
} as const;

export type StripReserve = keyof typeof STRIP_RESERVE;

/** What the strip is doing: waiting on the answer, showing a line, or offering a retry. */
export type StripPhase = 'loading' | 'line' | 'error';

export interface RowLine {
  /** The muted lead, when the line has a title in `--strong` after it. */
  lead: string;
  /** The strong title; empty when the line is a plain sentence. */
  title: string;
  /** The first spell, for the 20 px icon before the title. */
  spellId: number | null;
  /** The free line's second target, the gold Premium link; empty on Premium lines. */
  upgrade: string;
}

const count = (n: number, one: string, many: string): string => (n === 1 ? one : fill(many, { n }));

const UPGRADE_TAIL = ` · ${coachCopy.upgrade}`;

function plain(text: string): RowLine {
  return { lead: text, title: '', spellId: null, upgrade: '' };
}

function freeLine(answer: CoachTeaser): RowLine | null {
  if (answer.status !== 'ready') return null;
  const { finding_count: n, fix_first: first } = answer;
  if (n === 0) return plain(coachCopy.rowClean);
  if (first === null) {
    // The Premium sentence is its own target: the gold link, not muted lead text.
    const lead = count(n, coachCopy.rowFreeNoFirstOne, coachCopy.rowFreeNoFirst).replace(UPGRADE_TAIL, '');
    return { lead, title: '', spellId: null, upgrade: coachCopy.upgrade };
  }
  const template = n === 1 ? coachCopy.rowFreeOne : coachCopy.rowFree;
  const [lead = ''] = template.split('{title}');
  return {
    lead: fill(lead, { n }),
    title: first.title,
    spellId: first.spell_id,
    upgrade: coachCopy.upgrade,
  };
}

function fullLine(answer: CoachFull): RowLine | null {
  if (answer.status !== 'ready' && answer.status !== 'writing') return null;
  const first = answer.review.fixFirst;
  if (first === null) {
    const findings = answer.review.thenFix.length;
    if (findings === 0) return plain(coachCopy.rowClean);
    return plain(count(findings, coachCopy.rowPremiumNoFirstOne, coachCopy.rowPremiumNoFirst));
  }
  const [lead = ''] = coachCopy.rowPremium.split('{title}');
  const firstSpell = first.anchors
    .map((id) => answer.anchors[id]?.spell_id)
    .find((spell): spell is number => spell !== undefined);
  return { lead, title: first.title, spellId: firstSpell ?? null, upgrade: '' };
}

/**
 * `fallbackSpec` is the roster's class label, for an answer that carries no spec (a teaser from an
 * API that does not send one yet).
 */
export function rowLineFor(answer: CoachFull | CoachTeaser | null, fallbackSpec = ''): RowLine | null {
  if (answer === null) return null;
  if (answer.status === 'preparing') {
    const spec = answer.spec !== undefined && answer.spec !== '' ? answer.spec : fallbackSpec;
    return plain(fill(coachCopy.rowPreparing, { Spec: spec }));
  }
  return answer.access === 'teaser' ? freeLine(answer) : fullLine(answer);
}
