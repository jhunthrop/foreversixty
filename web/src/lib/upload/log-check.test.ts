import { describe, expect, it } from 'vitest';
import {
  HEADER_BYTES,
  checkLog,
  headerLacksAdvancedLogging,
  lacksEncounterStart,
  logWarningCopy,
} from './log-check';

const HEADER_ON =
  '10/19 21:00:00.000  COMBAT_LOG_VERSION,20,ADVANCED_LOG_ENABLED,1,BUILD_VERSION,1.15.9,PROJECT_ID,5\n';
const HEADER_OFF =
  '10/19 21:00:00.000  COMBAT_LOG_VERSION,20,ADVANCED_LOG_ENABLED,0,BUILD_VERSION,1.15.9,PROJECT_ID,5\n';
const TRASH = '10/19 21:00:01.000  SWING_DAMAGE,Player-1,"Bow",0x511\n'.repeat(20);
const BOSS = '10/19 21:05:00.000  ENCOUNTER_START,1084,"Onyxia",9,20,249\n';

const file = (text: string): File => new File([text], 'WoWCombatLog.txt', { type: 'text/plain' });

describe('headerLacksAdvancedLogging', () => {
  it('is false when the header switches advanced logging on', () => {
    expect(headerLacksAdvancedLogging(HEADER_ON)).toBe(false);
    expect(headerLacksAdvancedLogging('ADVANCED_LOG_ENABLED 1')).toBe(false);
  });
  it('is true when it is off, or the header does not say', () => {
    expect(headerLacksAdvancedLogging(HEADER_OFF)).toBe(true);
    expect(headerLacksAdvancedLogging('')).toBe(true);
    expect(headerLacksAdvancedLogging('ADVANCED_LOG_ENABLED,10')).toBe(true);
  });
});

describe('lacksEncounterStart', () => {
  it('finds a boss fight anywhere in the file, past the first 64 KB', async () => {
    const padding = 'x'.repeat(HEADER_BYTES * 2);
    expect(await lacksEncounterStart(file(HEADER_ON + padding + BOSS))).toBe(false);
  });
  it('finds a marker split across two chunks', async () => {
    const text = 'a'.repeat(95) + 'ENCOUNTER_START' + 'b'.repeat(40);
    expect(await lacksEncounterStart(file(text), 100)).toBe(false);
  });
  it('is true for a log of trash only', async () => {
    expect(await lacksEncounterStart(file(HEADER_ON + TRASH))).toBe(true);
  });
});

describe('checkLog', () => {
  it('flags a trash-only log written with advanced logging off both ways', async () => {
    expect(await checkLog(file(HEADER_OFF + TRASH))).toEqual({ advancedOff: true, noEncounters: true });
  });
  it('flags nothing for a full log', async () => {
    expect(await checkLog(file(HEADER_ON + TRASH + BOSS))).toEqual({
      advancedOff: false,
      noEncounters: false,
    });
  });
  it('words the two warnings as the form shows them', () => {
    expect(logWarningCopy.advancedOff).toBe(
      'Advanced combat logging was off for this file: no gear, resources or positions in the report.',
    );
    expect(logWarningCopy.noEncounters).toBe(
      'No boss fights found: this report will have no ratings or rankings.',
    );
  });
});
