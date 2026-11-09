// The /logs sample row and the report page's subtitle say the same things about the sample,
// because both read the one bundled meta through reportFacts.
import { describe, expect, it } from 'vitest';
import sampleMeta from '../../../public/sample-report/sample/meta.json';
import { sampleRowFromMeta } from '../logs/sample-row';
import { reportFacts } from './heading';
import type { ReportMeta } from './types';

const meta = sampleMeta as unknown as ReportMeta;

describe('the sample, one source', () => {
  const facts = reportFacts(meta);
  const row = sampleRowFromMeta(meta);

  it('says the same date: 8 Nov 2026', () => {
    expect(facts.date).toBe('8 Nov 2026');
    expect(row.dateLabel).toBe(facts.date);
  });

  it('says the same zone and fight count', () => {
    expect(row.summary).toContain(facts.zone);
    expect(row.summary).toContain(facts.fights);
    expect(facts.fights).toBe(`${meta.fights.length} fights`);
  });

  it('says the bosses the same way on both: the night killed six', () => {
    expect(facts.bosses).toBe('6 bosses');
    expect(row.summary).toContain(facts.bosses);
  });

  it('leaves the bosses out when none died, on both', () => {
    const trashOnly = {
      ...meta,
      fights: meta.fights.filter((fight) => fight.kind !== 'encounter'),
    } as ReportMeta;
    expect(reportFacts(trashOnly).bosses).toBe('');
    expect(sampleRowFromMeta(trashOnly).summary).not.toMatch(/boss/i);
  });

  it('is titled without "sample log"; the pill says it is a sample', () => {
    expect(meta.title).toBe('Sample dungeon night, 8 Nov 2026');
    expect(row.title).toBe(meta.title);
  });
});
