# Addon growth: what DungeonJournal does that we don't (2026-11-14)

## For the owner: the project name

CurseForge's search matches the project name and the author only, not the summary
(https://docs.curseforge.com/rest-api/). "Forever Sixty" matches nothing a player types.

- Proposed: **Forever Sixty: Gear Upgrades, BiS and Talents**
- Alternative 1: Forever Sixty: BiS, Talents and Rotation Helper
- Alternative 2: Forever Sixty: BiS Gear and Talent Planner

The text is ready in design/listing/2026-11-14-growth-listing.md. Nothing is published.

## The real gap

CurseForge counts every file download, so each app auto-update counts as one
(https://www.reddit.com/r/CurseForge/comments/1jf960f/how_does_curseforge_count_downloads/). The
fair way to compare is downloads per file, which is roughly how many installs update. Numbers are
from the public cfwidget mirror of the CurseForge API (https://api.cfwidget.com/1707971 and
https://api.cfwidget.com/1706689), read 2026-11-14:

| | Total | Per file (recent) | Rough active installs |
|---|---|---|---|
| DungeonJournal (1707971) | 11,998,628 | 450K to 1.64M | about 700K |
| Forever Sixty (1706689) | 868 | 0 to 61 | about 50 |

So the gap is not inflated by their release cadence. It is real, and it is about four orders of
magnitude.

## Top findings

1. **Nobody can find us by searching.** Search reads the name and the author only
   (https://docs.curseforge.com/rest-api/), and by user reports it ranks by popularity before
   relevance (https://curseforge-ideas.overwolf.com/ideas/CF-I-6526). For "forever talent" we
   rank 46th of 107, and we are not on page 1 for "gear upgrade"
   (https://www.curseforge.com/wow/search?search=forever%20talent,
   https://www.curseforge.com/wow/search?search=gear%20upgrade). The winners carry the job in
   their name: DungeonJournal, GearQuest Forever, Talents Forever, What's Training?
   (https://www.curseforge.com/wow/search?sortBy=popularity&version=1.60.1).
2. **One streamer and the curated lists made DungeonJournal.** Towelliee recommended it and it
   took 78K downloads in four days
   (https://dalaranpostasi.com/haber/towelliee-forever-icin-dungeonjournal-addonunu-onerdi/, and
   it is on https://towelliee.com/addon/). It was the favourite in the top r/classicwow beta addon
   thread (https://www.reddit.com/r/classicwow/comments/1wr3egt/). Players pick from lists:
   mein-mmo's top 10
   (https://mein-mmo.de/en/wow-forever-the-10-most-important-addons-and-how-to-install-them,1593699/),
   kami-labs (https://kami-labs.fr/en/addons-wow-forever/), timesaver.gg, wowcampfire and
   Muchamiel. We are on none of them, and no player has mentioned Forever Sixty anywhere we could
   find.
3. **Our direct rivals are big and narrow.** GearQuest Forever (2.73M, "right-click any gear slot
   to see the best upgrades for your class, spec, and level 1-60",
   https://www.curseforge.com/wow/addons/gearquest-forever) and Talents Forever (738K,
   "talentsforever.com inside the game", https://www.curseforge.com/wow/addons/talents-forever)
   each own one of our jobs with a name that says so. Talents Forever is our own model, a website
   brought into the game, and it is 850 times our size.
4. **Their basics are our basics.** Read from the packages (list below), DungeonJournal, Leatrix
   Plus and GearQuest all have a minimap button, an AddOns-menu (compartment) entry, an
   IconTexture, slash commands and a first-run message or guide. We already ship every one: the
   minimap button and compartment entry (Minimap.lua), IconTexture in the TOC, /fs and
   /foreversixty, the first-login welcome (SiteStatus.welcome) and the first open of the window on
   Me's Set up (surfaces/Onboarding.lua). The addon's discovery plumbing is not the problem.
5. **The 2.0.40 on our page was a moderation lag, not a stuck release.** Every 2.0.41 to 2.0.52
   upload returned "Success!" as a Release for 1.60.1 (runs 43363234307 to 43635956657), and the
   CurseForge API now lists all of them as Release, with v2.0.52 as the main file. Five of them
   (2.0.43, 2.0.44, 2.0.45, 2.0.48, 2.0.50) have 0 downloads because each was replaced within an
   hour, before review let it out. Every upload sits "Under Review" first
   (https://support.curseforge.com/en/support/solutions/articles/9000197905-project-statuses-101),
   and DungeonJournal's 2.0.6 also had 1 download hours after upload.

## What we learned from their code

Read, not copied. Each package was downloaded into its own scratch folder and nothing was run.

| Addon | Licence | Pattern worth learning |
|---|---|---|
| DungeonJournal 2.0.6 (CurseForge 1707971) | All rights reserved (no licence file) | Version check over the hidden addon channel: each copy says its version to guild and group, and a player with an older copy is told once per session to update. Ten client languages. A button on the client's Looking For Group window and the world map, frames every player already opens. |
| Questie 12.9.0 (GitHub Questie/Questie) | GPLv3 | The data addon (QuestieDB) ships in the same zip as Questie, so one install gets both. A "Welcome to Questie" choice frame at first run, queued until out of combat. `## Category` in the TOC. |
| GearQuest Forever 1.1.6 (CurseForge 1698950) | All rights reserved (no licence file) | One zip, eleven folders: the class data and the translations are LoadOnDemand addons beside the main one. A plain "Leveling gear upgrades for WoW Forever" Notes line, and a French Notes line. A portrait guide for onboarding. |
| Leatrix Plus 1.60.30 (CurseForge 94855) | All rights reserved | TOC AddonCompartmentFunc, and one first-run chat line naming the slash command and the minimap button. |

We could not download Talents Forever (cfwidget had no entry for its slug).

## Ranked changes

| # | Change | Expected effect | Effort | Status |
|---|---|---|---|---|
| 1 | Put the job in the project name (owner's call, top of this report) | Shows up for gear, upgrade, BiS, talent searches; the single largest lever since search reads only the name | S | Prepared |
| 2 | Summary that leads with the green arrow and the gain, not the Forever Score | More clicks from search and the app list | S | Prepared |
| 3 | Publish the approved categories with Bags & Inventory for Minimap (adds Quests & Leveling) | Browse exposure where levelling players look | S | Prepared |
| 4 | Publish the approved description and 11 plates | The page shows the value in one picture, as GearQuest's does | S | Waiting on the Chrome extension |
| 5 | An "Install the addon" link on every BiS, talent and guide page of foreversixty.gg | Turns site visitors into installs, the Talents Forever route | S | Proposed (web lane) |
| 6 | Ask the curated lists (kami-labs takes suggestions on Discord, mein-mmo, timesaver, wowcampfire, Muchamiel) to consider us | Lists are where players choose their short set | S | Owner's call; we contacted no one |
| 7 | One author post on r/classicwow with a 20 second clip, and a creator ask through the creator program | DungeonJournal's lift came from one streamer | M | Owner's call; nothing posted |
| 8 | Fewer, bigger releases: about one a day | App users get each build once review lets it out; five of our last twelve never reached anyone | S | Proposed (release habit) |
| 9 | Ship Forever Sixty Data inside the main zip (Questie and GearQuest pattern), keeping the nightly data project | Hand installers and Wago users get one download | M | Proposed, below |
| 10 | Update notice over the addon channel: the hello already reaches the guild; add the version to it and tell an older copy once per session (DungeonJournal pattern) | Players on hand installs stay current | M | Proposed (Comms.lua) |
| 11 | TOC Notes that say the job, and `## Category: Class` | The AddOns list says what we do | S | Shipped |
| 12 | Translate (German, French, Spanish, Russian, Korean first) | EU sites like mein-mmo, and the EU half of the realms | L | Proposed |

### The Data dependency (change 9)

The main addon already loads without it (`## OptionalDeps: ForeverSixtyInbox, ForeverSixtyData`)
and the CurseForge app installs a required relation by itself, so app users are fine. The cost is
for hand installers and the page itself, which shows "Required dependency". Proposal: copy the
current ForeverSixtyData folder into the main zip at release time (addon-release.yml already does
this for ForeverSixtyInbox through .pkgmeta move-folders), and keep the nightly Data project for
fresher ratings. The question to settle first is the CurseForge app's behaviour when two projects
ship the same folder; until it is tested, the relation stays and the description's first line
says "Install with the app and Forever Sixty Data comes with it".

### For after the brand lane lands

Nothing in this lane needed Theme.lua, Tokens, Widgets, HubFrame, surfaces/Hub.lua or the tabs.

## Sources not reachable

Icy Veins and the CurseForge project page refused automated fetches (403); Wago's pages render
only with JavaScript. The CurseForge numbers come from the cfwidget mirror of its API.
