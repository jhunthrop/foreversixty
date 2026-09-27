// web/src/components/sim/CharacterStrip.test.ts
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import type { CharacterSource } from '../../lib/sim/types';
import { handoffCopy } from '../../lib/sim/handoff-copy';
import type { SimCharacter } from '../../lib/sim/character';
import CharacterStrip from './CharacterStrip.svelte';

const source: CharacterSource = { kind: 'addon', ref: '', captured_at: '2026-09-21T10:00:00Z' };

// Minimal filler; each fixture below overrides only the fields the assertion cares about --
// the same pattern character.test.ts's own `plannerGearFor` describe block uses.
const base: SimCharacter = {
  name: 'Thrallgar',
  spec: 'warrior-fury',
  class_slug: 'warrior',
  race_slug: 'orc',
  talent_level: 60,
  level: 60,
  tree_version: '1.15.9.69722',
  point_order: [],
  gear: {},
  gear_slots: [],
  buffs: [],
  consumables: [],
  source,
  professions: [],
  bags: [],
  bank: [],
  sets: [],
  loadouts: [],
};

const requiredProps = { items: new Map(), onchange: () => {}, onrace: () => {} };

// 30 points is well under the 51 a full build spends -- talent_level 39 (BASE_LEVEL 9 + 30)
// is what characterFromFs1/fromBuildDraft would actually stamp for this many points, but
// this test only needs it under `level`, which is the component's own gate. `base.level`
// is MAX_LEVEL (60), the same as a hand-built planner build's own level (design 6), so
// this reproduces exactly the case the note exists for: a planner build with fewer points
// than a full 60 would have, still simulated at 60.
const underLeveled: SimCharacter = { ...base, talent_level: 39, point_order: new Array(30).fill(0) };
const fullBuild: SimCharacter = { ...base, talent_level: 60, point_order: new Array(51).fill(0) };

describe('CharacterStrip’s honest "simmed at a level" note', () => {
  it('says the character is simmed at its own level when it has fewer than 51 points', () => {
    const { body } = render(CharacterStrip, { props: { character: underLeveled, ...requiredProps } });
    expect(body).toContain(handoffCopy.simmedAtLevel(60));
    expect(body).toContain('30 talent points');
  });

  it('says nothing extra for a full 51-point build', () => {
    const { body } = render(CharacterStrip, { props: { character: fullBuild, ...requiredProps } });
    expect(body).not.toContain(handoffCopy.simmedAtLevel(60));
  });

  // Level-aware sim design (2026-09-27): an addon export with no level section sets
  // `level` to the same talent-implied number `talent_level` already is (character.ts's
  // own `characterFromFs1`), so there is no discrepancy left between what the strip shows
  // and what actually gets simulated -- the note would be false if shown here.
  it('says nothing when the character’s own level already matches its talent spend', () => {
    const noExportLevel: SimCharacter = {
      ...base,
      talent_level: 39,
      level: 39,
      point_order: new Array(30).fill(0),
    };
    const { body } = render(CharacterStrip, { props: { character: noExportLevel, ...requiredProps } });
    expect(body).not.toContain('Simmed as a level');
  });

  // An export that DOES carry a level, above what its talent spend implies (a respec, or
  // simply fewer points than that level could have), still gets the honest note -- now
  // naming that level rather than always 60.
  it('names the export’s own level, not always 60, when there is a real discrepancy', () => {
    const respecced: SimCharacter = {
      ...base,
      talent_level: 22,
      level: 45,
      point_order: new Array(13).fill(0),
    };
    const { body } = render(CharacterStrip, { props: { character: respecced, ...requiredProps } });
    expect(body).toContain(handoffCopy.simmedAtLevel(45));
  });
});

/** Pulls the `href` of the one `data-testid="sim-open-planner"` anchor out of a render's HTML. */
function plannerLinkHref(body: string): string | null {
  const match =
    /data-testid="sim-open-planner"[^>]*href="([^"]*)"/.exec(body) ??
    /href="([^"]*)"[^>]*data-testid="sim-open-planner"/.exec(body);
  return match?.[1] ?? null;
}

// Defect fix: SavedSim.svelte's own character carries no `point_order` at all (genuinely
// unreconstructible for a saved sim), so this component's default derivation -- built for
// the live /sim page, where `point_order` is the truth -- can only ever encode zeroed
// talents from it. `plannerHref` is the caller's own escape hatch: when given, it wins
// outright, `character.point_order` and the talent-file fetch this component would
// otherwise need are never consulted for the link at all.
describe('CharacterStrip’s plannerHref override', () => {
  it('uses the precomputed plannerHref verbatim, ignoring point_order entirely', () => {
    const { body } = render(CharacterStrip, {
      props: {
        character: base, // point_order: [] -- would zero every talent through the default path
        ...requiredProps,
        plannerHref: '/planner?code=fs1v2:1.60.1.69893:warrior:orc:0/5530515/0:head=12640',
      },
    });
    expect(plannerLinkHref(body)).toBe('/planner?code=fs1v2:1.60.1.69893:warrior:orc:0/5530515/0:head=12640');
  });

  it('falls back to its own class+race-only link when no plannerHref is given (unchanged default)', () => {
    const { body } = render(CharacterStrip, { props: { character: base, ...requiredProps } });
    // No TalentIndex resolves during an SSR render (onMount never fires), so the default
    // path's own null-index fallback is what a fresh render always shows first -- exactly
    // the live page's own first paint, `CharacterStrip`'s file header already documents.
    // SSR-escaped: `&` renders as `&amp;` in the raw HTML this test reads.
    expect(plannerLinkHref(body)).toBe('/planner?class=warrior&amp;race=orc');
  });
});

// 2026-09-21 result-page review round 3, newcomer's own finding: the live page shows "51
// points" beside Open in planner; SavedSim.svelte's own character (point_order: [], the
// same reason plannerHref needs an override above) made this component's own
// `point_order.length` read 0 and the count was hidden entirely. `talentPoints` is
// SavedSim's own escape hatch, the same shape as `plannerHref`.
describe('CharacterStrip’s talentPoints override', () => {
  it('shows the precomputed talentPoints when point_order is empty', () => {
    const { body } = render(CharacterStrip, {
      props: { character: base, ...requiredProps, talentPoints: 51 },
    });
    expect(body).toContain('data-testid="sim-talent-count"');
    expect(body).toContain('51 points');
  });

  it('prefers a real point_order over the override when both are present', () => {
    const { body } = render(CharacterStrip, {
      props: { character: fullBuild, ...requiredProps, talentPoints: 3 },
    });
    expect(body).toContain('51 points');
    expect(body).not.toContain('3 points');
  });

  it('shows nothing, not "0 points", when neither point_order nor talentPoints is known', () => {
    const { body } = render(CharacterStrip, { props: { character: base, ...requiredProps } });
    expect(body).not.toContain('data-testid="sim-talent-count"');
    expect(body).not.toContain('0 points');
  });
});

describe('CharacterStrip’s class icon', () => {
  it('shows the class icon, framed in the class colour, instead of a bare coloured square', () => {
    const { body } = render(CharacterStrip, { props: { ...requiredProps, character: base } });
    expect(body).toContain('data-testid="sim-character-class-icon"');
    expect(body).toContain('classicon_');
  });
});
