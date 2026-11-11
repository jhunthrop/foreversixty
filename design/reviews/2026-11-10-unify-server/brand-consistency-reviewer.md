# Brand consistency review, lane C (unify-server)

**CONSISTENT**

- Discord: consistent. Guild time reads "8 pm CST" (site clock); When/Was stay `<t:unix:t>`; ring crest avatar and class icons unchanged. (after/discord-post-full-desktop.png)
- Server pages: consistent. /b and /b-unavailable match the web reference; the app sign-in page uses the Marcellus wordmark and title, Open Sans body, the flat seal, the web plate (1px black edge plus inset hairline), turquoise primary button and muted note. (after/app-confirm-390.png, -1440.png; after/b-unavailable-1440.jpg vs header-width/after)
- Companion: consistent. Clock, "Sun 18 Oct", "14 Oct, 9:05 am" and compact ages follow apps/shared/formats.json; type, plates, nav and buttons match the site. (after/companion-reports-en-US.png, companion-device-connected-en-GB.png)

## Findings

Must: none.

Should
1. The Discord embed title prints the long form "Thursday 10 December" (after/discord-post-full-desktop.png). The web and formats.json use weekday_day_month ("Thu 10 Dec") for raid dates; the long form is for headings. Use the short form unless the owner wants the long one on Discord. Also check the title day is computed in the guild zone: formats.json shows 02:00Z reads "Wed 9 Dec" in Chicago, while the post pairs "Thursday 10 December" with "8 pm CST".

Nit
2. app_confirm.go gives the h1 text-shadow:0 1px 2px rgba(0,0,0,.45); the web page titles carry none. Drop it. (after/app-confirm-1440.png)
3. app_confirm.go hardcodes #000 for the panel edge and the rgba shadow; route them through brand tokens (renderGo) so they cannot drift. Every other colour comes from tokens_gen.go.
4. The confirm page ground is flat tan under the sky gradient; the web ground carries the parchment texture. Acceptable for a data:-only CSP page. (after/app-confirm-1440.png)
5. The confirm page has no site header or footer; fine for a pre-auth interstitial, noted so it is not mistaken for a regression.
6. Discord "When 8:00 PM" and "Today at 4:00 PM" are Discord own <t:> rendering in the reader locale; allowed by the owner ruling, no fix.
7. The lane b-unavailable capture rasterises the nav and subtitles with a different font renderer than the reference capture; markup is unchanged (TestSharedBuildParityFixtures), so a capture environment difference.
