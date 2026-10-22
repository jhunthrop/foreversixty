# Round 6 onboarding walk — setup's done lines against the real local API (21 Oct 2026, 00:55)

Local stack: API 8080 (pre-G.7 binary), web dev 4321 on main at `ef21f761` (W.9 merged). Browser pane, desktop width. Dev sign-in per `r4/local-stack.md`.

| # | Who | Surface | Read | Verdict |
|---|---|---|---|---|
| 1 | hunthrop#1894 (GM, two characters) | `/setup` | Step 1 "Signed in as hunthrop#1894" with a tick; step 2 "Pipsqueak arrived from the addon 2 hours ago" (the most recently saved character, the alt, not the main); step 3 "No companion signed in yet". One gold button per step. | OK. Note: the arrived line names whichever character saved last; a player with alts may read the alt's name and wonder about the main. Carry as polish: "Pipsqueak and 1 more arrived…" or name the main. |
| 2 | Newcomer#1111 (a bare `users` row inserted for the walk, no characters, no companion) | `/setup` | "Signed in as Newcomer#1111"; step 2 "No character yet — log out once with the addon on"; step 3 "No companion signed in yet"; chips "Signed in ✓ · 2 Addon · 3 Companion". | OK with the real API. The step-2 line does not mention the companion or the export route (P6-1) — W.10 is rewording it. |
| 3 | — | The fresh-account proof (install → in-game welcome → logout → character on the site) | Not walkable here: dev sign-in only mints sessions for existing accounts, and the game is the owner's. | Owner item (P6-2): one fresh Battle.net account on production, screenshots of the three chat lines, the logout and the character page. |

The pane was signed back in as hunthrop#1894 afterwards. The `Newcomer#1111` row stays in `forever_dev` (id 15) for later walks; `tools/dev-stack/seed.sh` does not know it.
