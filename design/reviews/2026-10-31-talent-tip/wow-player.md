# Talent tooltip fix - wow-player verdict

Verdict: SHIP

Seen: after captures at 1024, 1440, 2000 and 360. I did not compare the before captures in detail.

- Every tooltip in the captures reads in full: Improved Thunder Clap and Defiance (right column, Protection) at 1024, 1440 and 2000, and Improved Overpower (right column, Arms) on the phone. The "Improved T..." cut-off is gone. At 1024 the tooltip hangs over the middle of the Protection tree and the rail stays out of the way. At 1440 and 2000 it sits over the Level 60 build panel and covers its "See what differs" button for as long as you hover.
- On the phone the tooltip stays inside the 360 px screen (about 8 px clear on the right).
- It reads like a talent tooltip: name, "Rank 0 of 3", "NEXT RANK" text, then the build line. I'd add the "Requires N points in Arms Talents" line the client shows for locked talents, but that is not a blocker.

Not blockers:
- On desktop the hovered talent has no highlight ring. The phone capture shows one on Improved Overpower, so the two are inconsistent.
- No capture shows the flip-above case near the bottom of the viewport. The 360 bottom capture is identical to the 360 top one, so I could not verify the flip.
- The tooltip covers talents under it while you hover. The client does the same.
