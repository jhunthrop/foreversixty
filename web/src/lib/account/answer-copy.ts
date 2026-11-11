// web/src/lib/account/answer-copy.ts
// The signed-in answer's shared strings (guides spec 2026-11-08 gap G2): tiers, the guides and
// later pages answer "your character" with the same words.
import { uiCopy } from '../ui/copy';
export const answerCopy = {
  bis: 'Your best in slot',
  sim: 'Sim your character',
  unknown: (name: string): string => `${name}'s spec is not known yet.`,
  unknownLink: 'Choose a character',
  error: 'We could not load your character.',
  retry: uiCopy.retry,
} as const;
