// web/src/lib/sim/server-runs-copy.test.ts
import { describe, expect, it } from 'vitest';
import { serverRunsCopy } from './server-runs-copy';

describe('serverRunsCopy', () => {
  it('carries the fixed strings of spec section 6', () => {
    expect(serverRunsCopy.heading).toBe('Server runs');
    expect(serverRunsCopy.empty).toBe(
      'No server runs yet. Choose Run on our servers when you start a sim, and it keeps going after you close this tab.',
    );
    expect(serverRunsCopy.loadingSr).toBe('Reading your server runs…');
    expect(serverRunsCopy.loadFailed).toBe('Your server runs did not load.');
    expect(serverRunsCopy.overflow).toBe('Older finished runs are in Your sims below.');
    expect(serverRunsCopy.queued).toBe('Starting');
    expect(serverRunsCopy.running).toBe('Running');
    expect(serverRunsCopy.unreadSr).toBe('Not opened yet');
    expect(serverRunsCopy.runAgain).toBe('Run again');
    expect(serverRunsCopy.runAgainPending).toBe('Starting…');
    expect(serverRunsCopy.runAgainDone).toBe('Started');
    expect(serverRunsCopy.runAgainFailed).toBe('Could not start it again. Try again.');
    expect(serverRunsCopy.planLapsed).toBe('Server sims are part of Premium.');
    expect(serverRunsCopy.seePlans).toBe('See plans');
    expect(serverRunsCopy.bandOpen).toBe('Open result →');
    expect(serverRunsCopy.bandOpenPhone).toBe('→');
    expect(serverRunsCopy.dismissFailed).toBe('Could not dismiss that. Try again.');
  });

  it('runningLine is the muted suffix: the estimate, the overdue words, or nothing', () => {
    expect(serverRunsCopy.runningLine(180, 120)).toBe(' · takes about 3 minutes');
    expect(serverRunsCopy.runningLine(60, 10)).toBe(' · takes about a minute');
    expect(serverRunsCopy.runningLine(7200, 10)).toBe(' · takes about 2 hours');
    expect(serverRunsCopy.runningLine(180, 271)).toBe(' · taking longer than usual');
    expect(serverRunsCopy.runningLine(undefined, 540)).toBe('');
  });

  it('time lines', () => {
    expect(serverRunsCopy.startedLine('just now')).toBe('Started just now');
    expect(serverRunsCopy.startedLine('2 minutes ago')).toBe('Started 2 minutes ago');
    expect(serverRunsCopy.finishedLine('14 minutes ago')).toBe('Finished 14 minutes ago');
    expect(serverRunsCopy.endedLine('2 hours ago')).toBe('Ended 2 hours ago');
  });

  it('failedLine appends the API reason verbatim, or stands alone', () => {
    expect(serverRunsCopy.failedLine('stopped before completion')).toBe('Failed: stopped before completion');
    expect(serverRunsCopy.failedLine(undefined)).toBe('Failed');
    expect(serverRunsCopy.failedLine('')).toBe('Failed');
  });

  it('bandLineParts and dismissLabel', () => {
    expect(serverRunsCopy.bandLineParts('Frost Mage', 1204)).toEqual({
      prefix: 'Your Frost Mage sim finished: ',
      figure: '1,204 DPS',
    });
    expect(serverRunsCopy.dismissLabel('Frost Mage')).toBe('Dismiss: your Frost Mage sim finished');
  });

  it('moreLabel', () => {
    expect(serverRunsCopy.moreLabel(2)).toBe('and 2 more');
  });
});
