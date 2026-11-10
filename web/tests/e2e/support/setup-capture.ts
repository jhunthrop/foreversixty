// The /setup companion capture's expected aspect: the narrow crop under 640 px (Capture.svelte), the wide one above.
import { addonCopy } from '../../../src/lib/addon/copy';

const NARROW_BELOW = 640;

export function companionCaptureRatio(viewportWidth: number): number {
  const capture = addonCopy.setupCompanionCapture;
  const shown = viewportWidth < NARROW_BELOW ? capture.narrow : capture;
  return shown.width / shown.height;
}
