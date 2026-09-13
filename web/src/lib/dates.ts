const MONTHS = ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'June', 'July', 'Aug', 'Sept', 'Oct', 'Nov', 'Dec'];

export function formatDate(d: Date): string {
  return `${MONTHS[d.getUTCMonth()]} ${d.getUTCDate()}`;
}

export function isPast(d: Date, now: Date = new Date()): boolean {
  return d.getTime() < now.getTime();
}

/** The most recent of a set of dates -- an index page is "updated" when its newest entry is. */
export function latest(dates: Date[]): Date {
  if (dates.length === 0) throw new Error('latest() needs at least one date');
  return dates.reduce((newest, d) => (d.getTime() > newest.getTime() ? d : newest));
}
