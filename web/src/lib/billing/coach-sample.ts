// web/src/lib/billing/coach-sample.ts
// The sample coach review in /premium's third column (polish spec 2026-10-28 section 3.3): the
// coach's Fury model answer (web/src/fixtures/coach/fury-model.json) as src/data/premium-coach-
// sample.json, which differs from it only by the player's ruling that the sample's figures agree
// (coach-sample.test.ts pins the whole difference), plus the zone the context line names. Cut down
// to what a visitor should read: drill text, link sentences, evidence chips and the "couldn't
// check" lines are left out. The page labels it a sample (plans-page-copy.ts).
import classes from '../../data/classes.json';
import sample from '../../data/premium-coach-sample.json';
import { firstSpellId } from '../coach/chips';
import { normalizeFightCoach } from '../coach/normalize';
import type { CoachAnchor, CoachFull } from '../coach/types';
import { SPECS } from '../sim/specs';
import { plansPageCopy } from './plans-page-copy';

/** The icon url for a spell id, or null when the build has none. */
export type IconOf = (spellId: number | null) => string | null;

export interface SampleLine {
  readonly title: string;
  readonly why: string;
  readonly icon: string | null;
}

export interface SampleWell {
  readonly text: string;
  readonly icon: string | null;
}

export interface CoachSampleView {
  readonly classSlug: string;
  readonly zone: string;
  readonly title: string;
  readonly context: string;
  readonly opening: string;
  readonly fixFirst: SampleLine | null;
  readonly then: readonly SampleLine[];
  readonly doingWell: readonly SampleWell[];
}

export function loadCoachSampleAnswer(): CoachFull {
  const answer = normalizeFightCoach(sample.answer);
  if (answer === null || answer.access === 'teaser') {
    throw new Error('coach sample: fury-model.json is not a full coach answer');
  }
  return answer;
}

function classOf(answer: CoachFull): { slug: string; name: string } {
  const spec = SPECS.find((row) => row.name === answer.spec && row.role === answer.role);
  const cls = classes.find((row) => row.slug === spec?.class_slug);
  if (spec === undefined || cls === undefined) {
    throw new Error(`coach sample: no ${answer.role} spec named ${answer.spec} in the catalogue`);
  }
  return { slug: cls.slug, name: cls.name };
}

function line(
  item: { title: string; why: string; anchors: readonly string[] },
  anchors: Readonly<Record<string, CoachAnchor>>,
  iconOf: IconOf,
): SampleLine {
  return { title: item.title, why: item.why, icon: iconOf(firstSpellId(item.anchors, anchors)) };
}

export function coachSampleView(answer: CoachFull, iconOf: IconOf): CoachSampleView {
  const cls = classOf(answer);
  const { review, anchors } = answer;
  return {
    classSlug: cls.slug,
    zone: sample.zone,
    title: answer.persona.label,
    context: plansPageCopy.coachSample.context(answer.spec, cls.name, sample.zone, answer.size, answer.kill),
    opening: review.opening,
    fixFirst: review.fixFirst === null ? null : line(review.fixFirst, anchors, iconOf),
    then: review.thenFix.map((item) => line(item, anchors, iconOf)),
    doingWell: review.doingWell.map((item) => ({
      text: item.text,
      icon: iconOf(firstSpellId(item.anchors, anchors)),
    })),
  };
}
