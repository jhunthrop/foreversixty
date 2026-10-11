// web/src/lib/bis/spec-header.ts
// The BiS spec page's mount of the shared character header (spec 2026-10-11 section 4): the same
// worn-slot count the landing shows, but only for a character of THIS page's spec (a warrior's
// gear counted against a mage list would be a meaningless zero), and the secondary action that
// sends this list to the addon.
import type { MeCharacter } from '../account/api';
import type { AnswerSpec, HeaderAction } from '../character-header/types';
import { classSlugFromName } from '../report/tree-sizes';
import { bisCopy } from './copy';
import { bisAnswerSpec } from './landing-answer';

export interface SpecPage {
  /** "warrior-fury". */
  specKey: string;
  classSlug: string;
  /** "Fury". */
  specName: string;
  build: string;
}

export function characterIsOfSpec(
  character: MeCharacter,
  page: Pick<SpecPage, 'classSlug' | 'specName'>,
): boolean {
  return (
    character.class !== undefined &&
    classSlugFromName(character.class) === page.classSlug &&
    character.spec === page.specName
  );
}

/** The count for a character of this spec; nothing at all for any other (the slot is left out). */
export function specPageAnswerSpec(page: SpecPage): AnswerSpec {
  return bisAnswerSpec({
    build: page.build,
    specKeyOf: (character) => (characterIsOfSpec(character, page) ? page.specKey : undefined),
    withoutList: { kind: 'none' },
  });
}

export function specPageAction(sendListHref: string): HeaderAction {
  return { label: bisCopy.sendListToAddon, href: sendListHref };
}
