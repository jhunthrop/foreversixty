// web/src/lib/billing/premium-example-copy.ts
// The server-run example's caption templates (B5 `examples.serverRun.caption`). Kept apart from
// plans-page-copy so the /sim landing card can import it without bundling the plan table.
import {
  exampleCaption,
  serverRunExample,
  type CaptionTemplates,
  type ServerRunExample,
} from './premium-examples';

export const serverRunCaptionTemplates: CaptionTemplates = {
  preset: (f) =>
    `A level-${f.level} ${f.specLabel} in the site's preset gear: ${f.combinations} combinations on our servers, ${f.ranAt}.`,
  ownGear: (f) =>
    `${f.characterName}, ${f.specLabel}: ${f.combinations} combinations on our servers, ${f.ranAt}.`,
};

/** The caption, also the image alt, read from the data file through the templates. */
export function serverRunCaption(ex: ServerRunExample = serverRunExample): string {
  return exampleCaption(ex, serverRunCaptionTemplates);
}

/** A run set in the number font: a whole `d Mon yyyy` date (one face, never `27` mono and `Oct`
 *  not), a grouped figure, or a plain number. */
const NUMBER_RUN = /(\d{1,2} [A-Z][a-z]{2} \d{4}|\d{1,3}(?:,\d{3})+|\d+)/;

export interface CaptionPart {
  text: string;
  mono: boolean;
}

/** Splits a caption into plain and number runs (a date is one run); both example cards set them in mono. */
export function splitCaptionNumbers(caption: string): CaptionPart[] {
  return caption
    .split(NUMBER_RUN)
    .map((text, index) => ({ text, mono: index % 2 === 1 }))
    .filter((part) => part.text !== '');
}
