// web/src/lib/billing/example-image.ts
// The one source for the premium example picture: /premium and the /sim landing card both read
// it, so the files, the reserved box and the alt text can never disagree.
import { serverRunCaption } from './premium-example-copy';
import { serverRunExample, type ServerRunExample } from './premium-examples';

export const EXAMPLE_IMAGE_WIDTH = 440;
export const EXAMPLE_IMAGE_HEIGHT = 330;

const EXAMPLE_IMAGE_BASE = '/premium/example-server-run';

/**
 * Spec B4 (gate fix, row 21): a failed load hides the img, so the reserved box stays plain
 * `--card-top` with no broken-image glyph and no alt text repeating the caption under it. The
 * static plans page carries it as an inline handler; the /sim card sets the same style.
 */
export const EXAMPLE_IMAGE_FAILED_STYLE = 'visibility: hidden';
export const EXAMPLE_IMAGE_ON_ERROR = "this.style.visibility='hidden'";

export interface ExampleImage {
  avif: string;
  webp: string;
  width: typeof EXAMPLE_IMAGE_WIDTH;
  height: typeof EXAMPLE_IMAGE_HEIGHT;
  alt: string;
  loading: 'lazy';
  decoding: 'async';
}

export function exampleImage(ex: ServerRunExample = serverRunExample): ExampleImage {
  return {
    avif: `${EXAMPLE_IMAGE_BASE}.avif`,
    webp: `${EXAMPLE_IMAGE_BASE}.webp`,
    width: EXAMPLE_IMAGE_WIDTH,
    height: EXAMPLE_IMAGE_HEIGHT,
    alt: serverRunCaption(ex),
    loading: 'lazy',
    decoding: 'async',
  };
}
