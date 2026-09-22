// web/src/lib/rating/moments.test.ts
import { describe, expect, it } from 'vitest';
import { momentHref } from './moments';
import type { RatingMoment } from './types';

const GUID = 'Player-4184-000000A1';

describe('momentHref', () => {
  it('a death moment links to the Deaths tab using the API’s own anchor', () => {
    const moment: RatingMoment = {
      kind: 'death',
      at_ms: 140_000,
      spell_id: 19712,
      spell_name: 'Arcane Explosion',
      avoidable: true,
      anchor: `death-${GUID}-140000`,
    };
    expect(momentHref(3, GUID, moment)).toBe(`?fight=3&tab=deaths#death-${GUID}-140000`);
  });

  it('an interrupt moment links to the Interrupts tab via the exchange-<guid>-<spellId> convention', () => {
    const moment: RatingMoment = { kind: 'interrupt', spell_id: 20066, spell_name: 'Repentance' };
    expect(momentHref(3, GUID, moment)).toBe(`?fight=3&tab=interrupts#exchange-${GUID}-20066`);
  });

  it('a dispel moment links to the Dispels tab the same way', () => {
    const moment: RatingMoment = { kind: 'dispel', spell_id: 6205, spell_name: 'Gehennas’ Curse' };
    expect(momentHref(3, GUID, moment)).toBe(`?fight=3&tab=dispels#exchange-${GUID}-6205`);
  });

  it('a carried-debuff moment links to the Debuffs tab via the aura-<guid>-<spellId> convention', () => {
    const moment: RatingMoment = { kind: 'carried-debuff', spell_id: 6205, spell_name: 'Gehennas’ Curse' };
    expect(momentHref(3, GUID, moment)).toBe(`?fight=3&tab=debuffs#aura-${GUID}-6205`);
  });

  it('a utility-uptime moment links to the Buffs tab the same way', () => {
    const moment: RatingMoment = { kind: 'utility-uptime', spell_id: 1160, spell_name: 'Demoralizing Shout' };
    expect(momentHref(3, GUID, moment)).toBe(`?fight=3&tab=buffs#aura-${GUID}-1160`);
  });

  it('an unrecognised kind, or one missing the spell id a link needs, has no href', () => {
    expect(momentHref(3, GUID, { kind: 'avoidable-hit', spell_name: 'Standing in fire' })).toBeNull();
    expect(momentHref(3, GUID, { kind: 'interrupt' })).toBeNull();
  });

  it('an unresolved player guid (Ruling 3’s no-match case) has no href even for a death', () => {
    expect(momentHref(3, '', { kind: 'death', anchor: 'death-x-1' })).toBeNull();
  });
});
