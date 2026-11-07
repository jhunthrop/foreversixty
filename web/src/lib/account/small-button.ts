// web/src/lib/account/small-button.ts
// The account page's small buttons (Sign out, Revoke, Unlink, Link Discord, Link Premium role, Set
// as main, Copy link) at the mock library's sm geometry (design/mocks/ws_lib.py `.btn.sm`: 12 px
// padding, 12 px label). The system's BUTTON_SM pads 18 px with a 13 px label, which wraps the head's
// meta line at 390 and the rail's rows at 340 px. One place, so it goes when BUTTON_SM matches the
// mock (raised with the coordinator, premium round 2026-11-08).
import { button, type ButtonTone } from '../ui/button';

export function smallButton(tone: ButtonTone = 'secondary'): string {
  return `${button(tone, 'sm')} px-3! text-[12px]!`;
}
