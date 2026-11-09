// web/src/lib/account/small-button.ts
// The account page's small buttons (Sign out, Revoke, Unlink, Link Discord, Link Premium role, Set
// as main, Copy link): the system's sm Button (lib/ui/button.ts), which the mock library's
// `.btn.sm` draws at the same geometry (design/mocks/ws_lib.py ROUND5_CSS).
import { button, type ButtonTone } from '../ui/button';

export function smallButton(tone: ButtonTone = 'secondary'): string {
  return button(tone, 'sm');
}
