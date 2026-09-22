// The one shared look every island's request-triggering button (Run, Save, Submit,
// Upload, Join, Claim...) takes on while its own request is in flight: dimmed, a progress
// cursor, no spinner glyph and no label swap (design 2026-09-22 spec section 3.2). Applied
// alongside `aria-busy` and the button's existing `disabled` flag, never instead of them.
export const BUSY_CLASS = 'opacity-60 cursor-progress';
