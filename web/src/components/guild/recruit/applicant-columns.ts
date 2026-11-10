// web/src/components/guild/recruit/applicant-columns.ts
// The open applications' figure columns (item level, best kill, days waiting), one
// track list shared by the list's head line and every applicant row.
export const APPLICANT_COLUMNS =
  'grid-cols-[72px_76px_56px] items-center gap-x-3 text-right tabular-nums whitespace-nowrap';

export const APPLICANT_HEADS: readonly string[] = ['Item level', 'Best %', 'Waiting'];

/** The votes and the two decisions keep one width from 1024, so the figures line up under the head. */
export const APPLICANT_ACTIONS = 'lg:w-[272px] lg:flex-nowrap';
