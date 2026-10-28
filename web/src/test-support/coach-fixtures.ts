// web/src/test-support/coach-fixtures.ts
// Coach API answers as the API sends them (raw JSON shapes), for the client, the normaliser and
// the panel tests. Each builder returns a fresh object, so a test may edit its copy.
import type { CoachFull, CoachTeaser } from '../lib/coach/types';

export const FIXTURE_GUID = 'Player-4184-000000A1';

export function coachTeaser(overrides: Partial<CoachTeaser> = {}): CoachTeaser {
  return {
    status: 'ready',
    access: 'teaser',
    finding_count: 3,
    fix_first: { title: 'Bloodthirst: 9 casts lost', code: 'core_casts', spell_id: 23881 },
    ...overrides,
  };
}

/** A full model review: a Fix first, one Then, one Doing well, one Couldn't check. */
export function coachModelReview(overrides: Partial<CoachFull> = {}): CoachFull {
  return {
    report_id: 'rep1',
    fight_index: 3,
    player_key: 'us-ruleset-Thrall',
    spec: 'Fury',
    role: 'dps',
    size: 40,
    kill: true,
    status: 'ready',
    source: 'model',
    busy: false,
    access: 'full',
    persona: { label: 'Your Fury coach', status: 'full' },
    reviewed_at: '2026-10-27T20:00:00Z',
    review: {
      opening: 'A clean pull with one habit to fix.',
      fixFirst: {
        finding: 'f1',
        title: 'Bloodthirst: 9 casts lost',
        why: 'Each gap cost a Bloodthirst.',
        drill: 'Watch the cooldown.',
        drillText: 'Press it on ready.',
        anchors: ['a1'],
        links: [],
      },
      thenFix: [
        { finding: 'f2', title: 'No flask at the pull', why: 'Pulled bare.', anchors: [], links: [] },
      ],
      doingWell: [{ strength: 's1', text: 'Whirlwind on cooldown.', anchors: ['a2'] }],
      couldntCheck: ["Spell ranks: this log doesn't tell us your level."],
    },
    anchors: {
      a1: { kind: 'cast', spell_id: 23881 },
      a2: { kind: 'death', death: `death-${FIXTURE_GUID}-200000` },
    },
    basis_labels: ['your log', 'top Fury logs'],
    ...overrides,
  };
}

export function coachRulesReview(overrides: Partial<CoachFull> = {}): CoachFull {
  return coachModelReview({
    source: 'rules',
    persona: { label: 'Your Fury review', status: 'full' },
    ...overrides,
  });
}

export const coachBusyReview = (): CoachFull => coachRulesReview({ busy: true });
export const coachWritingReview = (): CoachFull => coachRulesReview({ status: 'writing' });
export const coachOfficerReview = (): CoachFull => coachModelReview({ access: 'officer' });
export const coachNotReviewable = (): CoachTeaser =>
  coachTeaser({ status: 'not_reviewable', fix_first: null, finding_count: 0 });
export const coachPreparing = (): CoachTeaser =>
  coachTeaser({ status: 'preparing', fix_first: null, finding_count: 0 });
