// web/src/lib/not-found/answer.ts
// The 404's signed-in answer (pages spec 2026-11-08 §4.4, ruling Q6): the header selector's
// character with its guide and its best in slot at its level band, then a smaller link to its
// own page. Every other state is the guides' answer (lib/guides/guide-answer.ts). Pure: the
// mount feeds it the account read and the pointer.
import type { Me } from '../account/api';
import { answerCopy } from '../account/answer-copy';
import { selectedCharacter } from '../account/hero-character';
import { characterHref } from '../characters';
import type { CurrentCharacter } from '../current-character';
import { guideAnswerFrom, type GuideAnswer } from '../guides/guide-answer';
import { bandForLevel } from '../guides/guide-character';
import { notFoundCopy } from './copy';

/** The index of the smaller third link (`Your character ›`). */
export const CHARACTER_LINK_INDEX = 2;

export function notFoundAnswerFrom(
  me: Pick<Me, 'characters' | 'main_character_key'> | null,
  current: CurrentCharacter | null,
  failed = false,
): GuideAnswer {
  const answer = guideAnswerFrom(me, current, {}, failed);
  if (answer.state !== 'ready' || me === null) return answer;
  const raw = selectedCharacter(current, me);
  const { classSlug, specSlug, level } = answer.character;
  const bis = `/bis/${classSlug}/${specSlug}?band=${bandForLevel(level)}`;
  return {
    ...answer,
    links: [
      ...answer.links,
      { label: `${answerCopy.bis} ›`, href: bis },
      ...(raw === null
        ? []
        : [
            {
              label: `${notFoundCopy.yourCharacter} ›`,
              href: characterHref(raw.region, raw.ruleset, raw.name),
            },
          ]),
    ],
  };
}
