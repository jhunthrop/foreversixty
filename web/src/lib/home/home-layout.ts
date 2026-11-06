// web/src/lib/home/home-layout.ts
// Which home shows before paint (home spec 2026-11-06 §4): the session hint picks the signed-in
// island's place, and this attribute on <html> brings the signed-out page back when /v1/me says
// nobody is signed in after all. global.css reads both.
export const HOME_OUT_ATTR = 'data-home-out';
