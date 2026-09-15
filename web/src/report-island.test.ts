// @vitest-environment jsdom
// web/src/report-island.test.ts
// The same two things src/planner-island.test.ts asserts about the planner's bundle: the
// entry pulls in the site's fonts and global stylesheet, so the Worker-served shell is not
// left in Georgia and Arial, and it reads the report id from the place the shell puts it.
import { readFile } from 'node:fs/promises';
import path from 'node:path';
import { describe, expect, it } from 'vitest';
import { reportIdFrom } from './report-island';

describe('the report island entry', () => {
  it('imports the site’s fonts and global stylesheet', async () => {
    // import.meta.url is an http url under the jsdom environment this file needs for
    // `document`, so the source is read by path, the way src/planner-island.test.ts does.
    const source = await readFile(path.join(import.meta.dirname, 'report-island.ts'), 'utf8');
    expect(source).toContain("import './styles/fonts.css'");
    expect(source).toContain("import './styles/global.css'");
  });

  it('prefers the mount element’s id over the path', () => {
    const element = document.createElement('div');
    element.dataset.reportId = 'fixture2abcd';
    expect(reportIdFrom(element, '/reports/somethingelse')).toBe('fixture2abcd');
  });

  it('falls back to the id in the path, which is what the Worker-served shell has', () => {
    expect(reportIdFrom(document.createElement('div'), '/reports/fixture2abcd')).toBe('fixture2abcd');
    expect(reportIdFrom(document.createElement('div'), '/reports/fixture2abcd/')).toBe('fixture2abcd');
  });

  it('returns an empty id for a path with no report in it', () => {
    expect(reportIdFrom(document.createElement('div'), '/reports')).toBe('');
    expect(reportIdFrom(document.createElement('div'), '/logs')).toBe('');
  });
});
