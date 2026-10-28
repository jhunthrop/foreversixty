// web/src/lib/billing/premium-source-copy.ts
// The strings of design/specs/2026-10-27-subscribe-and-join.md section 9.2, "You already have it":
// the grant and fallback plan lines, the panel title, the billing buttons and the three
// `/premium` card lines. The own and seat lines stay in ./account-plan-copy (section 4.3 reuses
// them), so the account panel and the checkout panel cannot disagree.
import { zonedDayMonthYear } from '../dates';

export const premiumSourceCopy = {
  panelTitle: 'You already have Premium.',
  grant: {
    noEnd: 'Premium · granted, no end date',
    until: (date: string, zone?: string): string =>
      `Premium · granted until ${zonedDayMonthYear(date, zone)}`,
    givenNoEnd: 'Given by Forever Sixty. Nothing is billed to you.',
    givenUntil: 'Given by Forever Sixty. It ends on its own; nothing is billed to you.',
  },
  fallback: 'Premium',
  buttons: {
    manage: 'Manage billing',
    history: 'Billing history',
    seePlan: 'See your plan',
  },
  premiumCard: {
    own: 'You have Premium through your subscription.',
    seat: 'You have Premium through a seat from your guild.',
    grant: 'You have Premium through a grant from Forever Sixty.',
  },
} as const;
