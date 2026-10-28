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

const NUMBER_RUN = /(\d{1,3}(?:,\d{3})+|\d+)/;

export interface CaptionPart {
  text: string;
  mono: boolean;
}

/** Splits a caption into plain and number runs; both example cards set the numbers in mono. */
export function splitCaptionNumbers(caption: string): CaptionPart[] {
  return caption
    .split(NUMBER_RUN)
    .map((text, index) => ({ text, mono: index % 2 === 1 }))
    .filter((part) => part.text !== '');
}
