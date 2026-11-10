// web/src/lib/report/copy-link.test.ts
import { describe, expect, it } from 'vitest';
import { copyReportLink } from './copy-link';
import { reportCopy } from './copy';

describe('copyReportLink', () => {
  it('writes the url and says Copied', async () => {
    const written: string[] = [];
    const clipboard = { writeText: async (text: string) => void written.push(text) };
    await expect(copyReportLink('https://x/reports/a?fight=3', clipboard)).resolves.toBe(reportCopy.copied);
    expect(written).toEqual(['https://x/reports/a?fight=3']);
  });

  it('says Copy failed when the clipboard refuses or is missing', async () => {
    const clipboard = {
      writeText: async () => {
        throw new Error('denied');
      },
    };
    await expect(copyReportLink('u', clipboard)).resolves.toBe(reportCopy.copyFailed);
    await expect(copyReportLink('u', null)).resolves.toBe(reportCopy.copyFailed);
  });
});
