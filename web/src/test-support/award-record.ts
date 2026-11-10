// web/src/test-support/award-record.ts
// An award's decision record before the council kept records (api AwardRecord): every field null.
// The night report's award fixtures spread it.
import type { AwardRecord } from '../lib/guild/loot-council-api';

export const NO_AWARD_RECORD: AwardRecord = {
  ms_os: null,
  reason: null,
  votes_yes: null,
  votes_no: null,
  votes_officers: null,
};
