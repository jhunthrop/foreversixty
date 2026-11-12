// web/src/lib/creators/model.test.ts
import { describe, expect, it } from 'vitest';
import { PLANS } from '../billing/plans';
import creators from '../../../tests/e2e/fixtures/creators/creators.json';
import me from '../../../tests/e2e/fixtures/creators/me.json';
import page from '../../../tests/e2e/fixtures/creators/statements.json';
import requests from '../../../tests/e2e/fixtures/creators/requests.json';
import {
  PROGRAM_TERMS,
  adminCreatorsCopy,
  applyCopy,
  creatorStatsCopy,
  payoutsCopy,
  programCopy,
} from './copy';
import { emptyForm, formFromCreator, formFromRequest, inputFromForm } from './form';
import {
  canApprove,
  capDollars,
  channelTitle,
  creatorFigures,
  approvedByDay,
  closedMonthLine,
  dayLabel,
  earningsRows,
  firstYearEarnings,
  lineText,
  money,
  monthLabel,
  previousMonth,
  recentMonths,
  requestSub,
  statementFigures,
  statementStatusLine,
  statsFigures,
  toPay,
  weekRow,
  type Creator,
  type CreatorMe,
  type CreatorRequest,
  type Statement,
} from './model';

const list = creators as Creator[];
const statements = page.statements as Statement[];
const mine = me as CreatorMe;
const waiting = requests as CreatorRequest[];

describe('money and dates', () => {
  it('formats cents, caps and days the way the pages show them', () => {
    expect(money(16860)).toBe('$168.60');
    expect(money(200)).toBe('$2.00');
    expect(capDollars(50000)).toBe('$500');
    expect(capDollars(5050)).toBe('$50.50');
    expect(monthLabel('2026-11')).toBe('Nov 2026');
    expect(previousMonth('2026-01')).toBe('2025-12');
    expect(approvedByDay('2026-12')).toBe('7 Dec');
    expect(dayLabel('2026-11-08')).toBe('8 Nov 2026');
    expect(dayLabel('2026-12-02T23:30:00Z')).toBe('2 Dec 2026');
  });

  it('offers the open month and the twelve closed ones, newest first', () => {
    const months = recentMonths(new Date('2026-12-09T12:00:00Z'));
    expect(months).toHaveLength(13);
    expect(months.slice(0, 3)).toEqual(['2026-12', '2026-11', '2026-10']);
    expect(months[12]).toBe('2025-12');
  });
});

describe('the figures say what they cover', () => {
  it('leaves the four creator counts bare (the list says all time once) and notes the owed figure with its month and cap', () => {
    const figs = creatorFigures(list[0]);
    expect(figs.slice(0, 4).every((f) => f.note === '')).toBe(true);
    expect(adminCreatorsCopy.tableMeta(3)).toBe('3 creators · counts all time');
    expect(figs[4]).toMatchObject({ value: '$184.20', note: 'Dec 2026 · of $500 cap', lead: true });
    expect(creatorFigures(list[1])[4].note).toBe('Dec 2026 · $50 cap reached');
  });

  it('leads a statement with its total, and counts the bounties', () => {
    const figs = statementFigures(statements[0]);
    expect(figs[0]).toMatchObject({ label: payoutsCopy.columns.total, value: '$168.60', lead: true });
    expect(figs[1].note).toBe('64 × $2.00');
    expect(statementFigures(statements[1]).find((f) => f.key === 'cap')?.value).toBe('$24.40');
  });

  it('gives the creator their month with its cap, then the counts with each one’s rate', () => {
    const figs = statsFigures(mine);
    expect(figs.map((f) => f.note)).toEqual([
      'Dec 2026 so far · of $500 cap',
      '',
      '62% of players',
      '28% of returning',
    ]);
  });

  it('agrees across the fixtures: the stats page and the owner see the same December', () => {
    expect(mine.month_cents).toBe(list[0].figures.owed_cents);
    expect(mine.returning).toBe(list[0].figures.returning);
    expect(mine.weeks.reduce((n, w) => n + w.signed_up, 0)).toBe(mine.players);
    expect(mine.weeks.reduce((n, w) => n + (w.returning ?? 0), 0)).toBe(mine.returning);
  });
});

describe('a player in a creator first year', () => {
  it('is the bounty plus half of a yearly payment', () => {
    expect(firstYearEarnings('premium', 'yearly', true)).toBe(
      PROGRAM_TERMS.bountyUsd + PLANS.premium.founderYearly / 2,
    );
  });

  it('is the bounty plus half of twelve monthly payments', () => {
    expect(firstYearEarnings('premium', 'monthly', false)).toBe(
      PROGRAM_TERMS.bountyUsd + (PLANS.premium.monthly * 12) / 2,
    );
  });

  it('labels each row with its plan and price and gives the dollars', () => {
    const rows = earningsRows(true);
    expect(rows.map((r) => r.label)).toEqual([
      `Premium, $${PLANS.premium.founderYearly} a year`,
      `Premium, $${PLANS.premium.founderMonthly} a month`,
      `Guild, $${PLANS.guild.founderYearly} a year`,
      `Raid guild, $${PLANS.raid_guild.founderYearly} a year`,
    ]);
    expect(rows[0]?.sub).toBe(
      `$${PROGRAM_TERMS.bountyUsd + PLANS.premium.yearly / 2} at list price after 6 Jan 2027`,
    );
    expect(earningsRows(false)[0]?.sub).toBeUndefined();
    expect(rows[0]?.value).toBe(`$${PROGRAM_TERMS.bountyUsd + PLANS.premium.founderYearly / 2}`);
  });
});

describe('statements', () => {
  it('approves only a ready month that owes money; a $0 month has no button', () => {
    expect(statements.map(canApprove)).toEqual([true, false, false]);
    expect(canApprove({ status: 'ready', total_cents: 0 })).toBe(false);
    expect(canApprove({ status: 'open', total_cents: 500 })).toBe(false);
  });

  it('adds up what is ready to pay', () => {
    expect(toPay(statements)).toBe(16860);
  });

  it('says where each statement stands', () => {
    expect(statements.map(statementStatusLine)).toEqual([
      payoutsCopy.status.ready,
      payoutsCopy.status.blocked,
      payoutsCopy.status.nothing,
    ]);
    const paid = {
      ...statements[0],
      status: 'paid' as const,
      paid_at: '2026-12-02T12:00:00Z',
      transfer_id: 'tr_1',
    };
    expect(statementStatusLine(paid)).toBe('Paid 2 Dec 2026 · transfer tr_1');
    expect(payoutsCopy.confirmTitle('Grimbrew', '$168.60', monthLabel('2026-11'))).toBe(
      'Pay Grimbrew $168.60 for Nov 2026?',
    );
  });

  it('signs a refund line', () => {
    expect(lineText(statements[0].lines[3])).toEqual({
      text: 'Player #10498 · refund 21 Nov',
      amount: '−$2.50',
      refund: true,
    });
  });
});

describe('weeks and requests', () => {
  it('shows a running week’s sign-ups alone, marks a week under 14 days old as still filling', () => {
    const now = new Date('2026-12-09T12:00:00Z');
    const running = weekRow(mine.weeks[0], now);
    expect(running).toMatchObject({ title: 'Week of 7 Dec', filling: true });
    expect(running.figures.map((f) => [f.label, f.value])).toEqual([
      ['Signed up', '18'],
      ['Returning', '–'],
      ['Paying', '–'],
    ]);
    expect(closedMonthLine('2026-12', new Date('2026-12-05T12:00:00Z'))).toBe(
      'Nov 2026 closed on the 1st. We approve it by 7 Dec.',
    );
    expect(closedMonthLine('2026-12', now)).toBe('Nov 2026 closed on the 1st. We are approving it now.');
    const last = weekRow(mine.weeks[1], now);
    expect(last).toMatchObject({ title: 'Week of 30 Nov', filling: true });
    const closed = weekRow(mine.weeks[2], now);
    expect(closed.filling).toBe(false);
    expect(closed.figures.map((f) => f.value)).toEqual(['54', '36', '10']);
  });

  it('reads a request', () => {
    expect(channelTitle(waiting[1].channel_url)).toBe('youtube.com/@saltmarshguides');
    expect(requestSub(waiting[0])).toBe('Twitch · about 1,200 · 8 Dec 2026');
    expect(requestSub(waiting[1])).toBe('YouTube · size not given · 9 Dec 2026');
  });
});

describe('the owner’s form', () => {
  const now = new Date('2026-12-09T12:00:00Z');

  it('starts empty from today at a $500 cap', () => {
    expect(emptyForm(now)).toEqual({
      name: '',
      channel_url: '',
      slug: '',
      code: '',
      starts_on: '2026-12-09',
      cap: '500',
    });
  });

  it('round-trips a creator and prefills a request', () => {
    expect(inputFromForm(formFromCreator(list[1]))).toEqual({
      name: 'Lunawick',
      channel_url: 'https://youtube.com/@lunawick',
      slug: 'lunawick',
      code: 'LUNAWICK',
      starts_on: '2026-11-12',
      cap_cents: 5000,
    });
    expect(formFromRequest(waiting[0], now)).toMatchObject({
      name: 'emberfall',
      slug: 'emberfall',
      code: 'EMBERFALL',
      starts_on: '2026-12-09',
    });
  });

  it('uppercases the code and sends a bad cap for the API to refuse', () => {
    const input = inputFromForm({ ...emptyForm(now), code: ' grimbrew ', cap: '12.5' });
    expect(input.code).toBe('GRIMBREW');
    expect(input.cap_cents).toBe(-1);
    expect(inputFromForm({ ...emptyForm(now), cap: '' }).cap_cents).toBe(50000);
  });
});

describe('the words', () => {
  const all = JSON.stringify(
    [adminCreatorsCopy, payoutsCopy, creatorStatsCopy, programCopy, applyCopy],
    (_k, v: unknown) => (typeof v === 'function' ? String(v) : v),
  );

  it('uses no em dash and no exclamation mark', () => {
    expect(all).not.toContain('—');
    expect(all).not.toMatch(/![^=]/);
  });

  it('states the returning-player rule and the free month plainly on the public page', () => {
    const text = programCopy.sections.flatMap((s) => s.paragraphs).join(' ');
    expect(text).toContain('$2 for each player you send who comes back');
    expect(text).toContain('within 14 days of signing up');
    expect(text).toContain('for 12 months from their first payment');
    expect(text).toContain('a free month of Premium, or a free month for their guild');
    expect(text).toContain('The FTC and Twitch both require it');
  });

  it('every section figure lane names its section', () => {
    const ids = programCopy.sections.map((s) => s.id);
    for (const lane of programCopy.lanes) expect(ids).toContain(lane.section);
  });
});
