// web/src/lib/guides/signed-in-callout.ts
// `/guides` index's own signed-in callout (rebuild spec §4.A): "Your guide: Fury Warrior
// →", shown only once `me.characters[0]` resolves with a known class and spec -- the plain
// first character on the account, not the home hero's own "best character" selection
// (that answers a different question, "which character represents me across the site";
// this one answers "where's my own guide", which is the account's own lead character).
// Pure view-model only: the page's own inline bootstrap script (no Svelte island -- the
// guide rebuild's own budget keeps every page static beyond the one existing
// `GuideBuildTree` island) calls `fetchMeOnce()` itself and passes the result here.
import { classSlugFromName } from '../report/tree-sizes';
import { classCrestSrc } from '../class-crest';
import { classColorVar } from '../report/format';
import { guidesCopy } from './copy';
import type { MeCharacter } from '../account/api';
import classes from '../../data/classes.json';

export interface SignedInCalloutView {
  href: string;
  crestSrc: string;
  crestColorVar: string;
  text: string;
}

/**
 * The callout's own view-model for `me.characters[0]` -- `undefined` when there is no
 * character, or its class/spec have not been learned yet (never a guess at either, spec
 * §4.A/§5's own "omitted on failure or signed-out" rule, extended to "not yet known").
 */
export function signedInCalloutView(character: MeCharacter | undefined): SignedInCalloutView | undefined {
  if (character?.class === undefined || character.spec === undefined) return undefined;
  const classSlug = classSlugFromName(character.class);
  const specSlug = classSlugFromName(character.spec);
  const className = classes.find((c) => c.slug === classSlug)?.name;
  if (className === undefined) return undefined;
  return {
    href: `/guides/${classSlug}/${specSlug}`,
    crestSrc: classCrestSrc(classSlug),
    crestColorVar: classColorVar(className),
    text: guidesCopy.signedInCallout(character.spec, className),
  };
}
