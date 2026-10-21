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

  it('says the same date: 17 Sept 2026', () => {
    expect(facts.date).toBe('17 Sept 2026');
    expect(row.dateLabel).toBe(facts.date);
  });

  it('says the same zone and fight count', () => {
    expect(row.summary).toContain(facts.zone);
    expect(row.summary).toContain(facts.fights);
    expect(facts.fights).toBe(`${meta.fights.length} fights`);
  });

  it('leaves the bosses out when none died, on both', () => {
    expect(facts.bosses).toBe('');
    expect(row.summary).not.toMatch(/boss/i);
    expect(row.summary).toBe(`${facts.zone} · ${facts.fights}`);
  });

  it('says "N bosses" the same way on both when some did', () => {
    const withKills = {
      ...meta,
      fights: [{ ...meta.fights[0], kind: 'encounter', kill: true }, ...meta.fights.slice(1)],
    } as ReportMeta;
    expect(reportFacts(withKills).bosses).toBe('1 boss');
    expect(sampleRowFromMeta(withKills).summary).toContain('1 boss');
  });

  it('is titled without "sample log"; the pill says it is a sample', () => {
    expect(meta.title).toBe('Beta night, 17 Sept 2026');
    expect(row.title).toBe(meta.title);
  });
});
