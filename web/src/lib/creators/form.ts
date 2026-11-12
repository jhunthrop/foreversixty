// web/src/lib/creators/form.ts
// The owner's add, edit and approve form on /admin/creators: its values from a creator or a
// request, and the API input from its values. Pure, so the page and its tests share it.
import type { Creator, CreatorInput, CreatorRequest } from './model';

export interface CreatorFormValues {
  name: string;
  channel_url: string;
  slug: string;
  code: string;
  starts_on: string;
  /** Whole dollars, as typed. */
  cap: string;
}

/** The default monthly cap, in dollars (the API's creators.DefaultCapCents). */
export const DEFAULT_CAP_DOLLARS = '500';

/** "2026-12-09" for an instant, in UTC (the API's start dates are UTC days). */
export function utcDay(now: Date): string {
  return now.toISOString().slice(0, 10);
}

export function emptyForm(now: Date): CreatorFormValues {
  return { name: '', channel_url: '', slug: '', code: '', starts_on: utcDay(now), cap: DEFAULT_CAP_DOLLARS };
}

export function formFromCreator(c: Creator): CreatorFormValues {
  return {
    name: c.name,
    channel_url: c.channel_url,
    slug: c.slug,
    code: c.code,
    starts_on: c.starts_on,
    cap: String(Math.round(c.cap_cents / 100)),
  };
}

/** A request's Approve form: the API's suggestions, starting today. */
export function formFromRequest(r: CreatorRequest, now: Date): CreatorFormValues {
  return {
    ...emptyForm(now),
    name: r.suggested.name,
    channel_url: r.channel_url,
    slug: r.suggested.slug,
    code: r.suggested.code,
  };
}

/** The API input; a cap that is not a whole number of dollars goes as NaN cents for the API to refuse. */
export function inputFromForm(v: CreatorFormValues): CreatorInput {
  const dollars = v.cap.trim() === '' ? Number(DEFAULT_CAP_DOLLARS) : Number(v.cap.trim());
  return {
    name: v.name.trim(),
    channel_url: v.channel_url.trim(),
    slug: v.slug.trim(),
    code: v.code.trim().toUpperCase(),
    starts_on: v.starts_on.trim(),
    cap_cents: Number.isInteger(dollars) ? dollars * 100 : -1,
  };
}
