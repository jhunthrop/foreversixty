# Guild tab: interaction inventory (lane guild-tab, 2026-11-11)

Spec: design/specs/2026-11-10-guild-tab.md. Captures: `ios/<dark|light>/` (UITests/GuildTabCaptureTests, the shared
simulator, fixture app) and `android/<dark|light>/` (capture-android.sh, the shared emulator, debug APK in fixture mode),
each step at tap, +400 ms and settled. The mocked states, the assignments and the at-60 tab are the parity goldens
(apps/ios/AppTests/Guild/__Snapshots__/GuildTabParityTests, apps/android/app/src/test/screenshots/parity/android-guild-tab*),
rendered from the same views.

Both apps are dark only (`preferredColorScheme(.dark)`, the dark theme): the light captures show the system in light mode
with the app unchanged, as the site is. That is the expected result, not a gap.

## Fixture limits (not app behaviour)
- The fixture account's characters are Gorvash (58, Olympus XXVII) and Yellfury (31, no guild): the device walk covers
  under 60 in a guild, the officer (the officer raid and badge fixtures), and no guild under 60. At 60 and no guild at 60
  have no fixture character; their states are the parity renders (`guild-tab`, `guild-tab-none`).
- The fixture raids carry no assignments (`assignments: null`), so Your job and Assignments show in the parity renders
  (`guild-tab`, `guild-tab-officer`) and are left out on the device, which is the app's rule with no pick.
- Both walks read the 38-member roster home (`guild.home` variant `roster`); its viewer is a member (`officer_tools: false`),
  so the officer rows follow the officer badge. Full-page renders of every state: `parity/<ios|android>-full-*.png`.
- Site links (Plan a raid, Setup, Search guilds, Claim) open the browser on the live site.

## Steps

| Step | iOS | Android | What happens |
|---|---|---|---|
| Open the Guild tab | leveler-01-top | leveler-01-top | Under 60: the guild head (crest, name, members, faction) and the spread (at 60, 55 to 59, you) lead; the raid head answers live; `Raids open at 60 · you're 58` under it; Who is coming shows each role's count alone (`Healers 5`, no target), and Before you pull is left out (a 60's check, round 3) |
| Answer Late, then In | leveler-02, -03 | leveler-02, -03 | The four answers work at any level (coordinator ruling 2026-11-10); the tapped answer selects and `Saved.` shows (round 3: fixture mode echoes the write, SignupEcho, as the server does) |
| Roster › | leveler-04, -05 | leveler-04, -05 | Recently on and Close to 60 rows, then the roster page pushed on the Guild tab, back to the root |
| Progression | leveler-06 | leveler-06 | The tier's line (`1 of 1 down · N pulls this tier`); a tap opens the guild page's Progression on the site |
| Calendar | leveler-07 | leveler-07 | The planned raids pushed (iOS) or the site (Android, as before this lane) |
| Pull to refresh | leveler-08 | leveler-08 | The tab reloads the raids, the guild home and the badge |
| Officer: the tab | officer-01, -06 | officer-01, -06 | Needs you, then Officer: Plan a raid (the next two planned raids), Applications (`1 waiting`), Setup (`4 of 8 done`, the steps still to do) |
| Officer: Ping on Discord | officer-07 | officer-07 | The not-answered list opens on the Guild tab. Mention posts in Discord and the screen keeps a result line: `Mentioned 2 healers in Discord.`, with `1 was mentioned in the last 6 hours and was left out.` when the server skipped someone, `Nobody to mention: none of them has linked Discord.` when nothing posted, and a refusal line for a repeat too soon. The fixture raid carries no `not_signed_raiders`, so the device list is empty and Mention cannot be tapped there; the result lines are pinned by RaidFlowTests (iOS, mentionStatus) and NotAnsweredTest (Android, MentionWords) |
| Officer: Applications | officer-03 | officer-03 | The newest applicant's sheet, one tap |
| Officer: Setup | officer-04 | officer-04 | The site's setup checklist (`{guild page}/setup`) |
| Officer: Plan a raid | officer-05 | officer-05 | The guild page's Raids tab on the site (`#raids`), where Plan a raid lives |
| Officer: Assignments › Open in planner | parity `guild-tab-officer` | parity `guild-tab-officer` | The raid's page on the site (`{guild page}/raids/{raid}`); `Assigning is on the site.` |
| Roster at 60 | parity `guild-tab`, `guild-tab-officer` | parity `guild-tab`, `guild-tab-officer` | Round 3: Your attendance (`7 of 8 nights`) for everyone; officers also get Missing raids and Gear gaps (the addon's gap, largest three, `41 DPS`); your own loot award is gold |
| Your job | parity `guild-tab`, raid page | parity `guild-tab`, raid page | `Whelps, west cave with Kaelra and Thalrin`, `Onyxia · <purpose>`; gone with no slot |
| No guild | noguild-01 | noguild-01 | `Raids open at 60 · you're 31`, the guilds recruiting the spec (Apply, or `Applied Mon`), Find your guild |
| Search guilds on the site | noguild-02 | noguild-02 | The site's guild finder `/guild` (the app's old `/guilds` link was a 404) |
| Claim your guild | noguild-03 | noguild-03 | `/guild#claim` |
| A raid link | link-01, -02 | link-01, -02 | `/app/raid/2/31` (a push's tap-through) selects Guild and pushes the raid page; back returns to the Guild root |

## Tab bar
The Guild door (design/assets/icons/ui/guild.svg, the web header's Guild glyph) at 24 pt / 24 dp with the 52-unit stroke,
tinted by the tab bar exactly as Home, Gear and Activity: gold selected, the muted ink unselected, the same opacity. Every
tab-root capture above shows the four icons together, selected and unselected.
