# ux-designer verdict: addon dialogs fix (8d7a719a1)

Verdict: SHIP

Reviewed what renders: pixel-grid-1.875.png, pixel-grid-1.40625.png, both footer-2x crops, onboarding-4-guild.png, onboarding-1-loop.png. Compared against the owner's in-game capture (tour page 4 of 5, 2.0.44).

## Reasons

1. Blank button fixed. Left (2.0.44) shows the bottom-left button as an empty blue box, matching the owner's capture. Right (fixed) reads "Skip tour" at both 1.875 and 1.40625. All three footer buttons (Skip tour, Back, Next) are labelled on the fixed side.
2. Crisp. Fixed-side glyphs in the title, body and button labels have clean stems with no soft halo at either scale; the 2.0.44 side is visibly softer, which is the owner's "blurry" complaint.
3. Consistent edges. In the footer 2x crops the fixed buttons have one uniform border weight on all four sides and across all three buttons. The 2.0.44 side has slightly uneven edge weight between buttons. The panel outline and the header rule are the same weight on the fixed side.
4. Crest is crisp at both scales (128 px source). Same ring and position as 2.0.44; nothing moved.
5. Layout unchanged. Mock pages 1 and 4 keep the footer order (Skip tour left, Back/Next right, Tooltips only on page 1), header, and two-column body. No regression visible.

## Must-fix

None blocking.

## Caveats and polish (not blocking)

- The pixel-grid boards are an emulation of client rasterization, not an in-game capture. They also show only the header, one body line and the footer. The close button, the "4 of 5" counter, the page-4 columns and the other addon windows are not in any render. Owner confirmation in game at 1440p and 1080p is the real acceptance. The attach list should include an in-game page 4 capture after the next release.
- Pre-existing, not from this fix: in the page-4 mock, the right column rows start about 2 px lower than the left column's first row (Answer In... vs Tonight...), and the grey "coming" tag sits inline after "phone app" with no gap token. The owner's in-game capture shows the same. Raise as polish in a later pass.
- Mock shows "Tonight 8:00 pm", in-game shows "08:00 pm". Different copy from the mock; check against the unified locale clock format, separate from this fix.
