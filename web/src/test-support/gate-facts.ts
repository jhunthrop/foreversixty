// web/src/test-support/gate-facts.ts
// The gate sentence's facts for a component test: the claim holder of a guild with no plan, on a
// fresh trial, unless a test says otherwise.
import type { GateFacts } from '../lib/guild/plan-gate';

export const GATE_SETTINGS_HREF = '/guild/us/pvp/olympus-xxvii/settings';
export const GATE_CLAIM_HREF = '/guild/us/pvp/olympus-xxvii/claim';

export function gateFactsFixture(over: Partial<GateFacts> = {}): GateFacts {
  return {
    claim: {
      state: 'claimed',
      yours: true,
      claimed_by: { name: 'Thornhide', class: 'druid', character_key: 'us/pvp/thornhide' },
    },
    guildMaster: true,
    trialUsed: false,
    lapsed: false,
    guildTier: false,
    granted: false,
    settingsHref: GATE_SETTINGS_HREF,
    claimHref: GATE_CLAIM_HREF,
    ...over,
  };
}
