# Creator outreach

A DM template and a list of creators. Drafts only: the owner sends each message by hand through
the creator's own public channel for business contact, such as a YouTube "About" business email
button, a Twitch panel or a Discord they run. This file holds no personal contact details.

## The offer

| What | How it is delivered |
|---|---|
| Free Premium for as long as they cover Forever | `api grant --user <battletag> --plan premium --note "creator: <name>"` with no `--until`, or one year. Run it from the deployed sha only, since the grant CLI migrates at startup. |
| Early sim data | The nightly tier list and leveling BiS for their class, sent the day before it goes live when a client build changes the rankings. Raw numbers come from the nightly run; never promise a figure in advance. |
| Their own tracked link | `foreversixty.gg/go/yt-<name>`, pointing at the guide for their main class (targets are below). All 20 were created in production on 3 November 2026 and each one 302s to its target. The owner's /admin/links page shows the clicks and sign-ups each link brings. |

**Owner decisions:**

1. Premium for how long: open-ended, a year, or to the end of Phase 1.
2. Whether "early sim data" means a day-ahead file by DM, or only a heads-up when a ranking moves.
   Day-ahead files are a promise someone has to keep every night.
3. Whether to offer anything for a sponsored mention. The template asks for nothing in return, on
   purpose.

## DM template

Keep it under 120 words. Change the bracketed parts for every creator. Never send two creators
the same first line.

> Hi [name], I loved your [specific video title, and one specific thing from it].
>
> I run Forever Sixty, a free helper for Forever players: class guides, leveling
> BiS at every level range ranked by our own sim, a tier list, and an addon that marks upgrades in
> your bags.
>
> If it is useful for your [class] content, I would like to give you:
>
> - Premium free, for as long as you cover Forever
> - Our sim results for [class] before they go live, when a patch moves them
> - Your own link, foreversixty.gg/go/yt-[name], so you can see what your viewers use
>
> No ask in return, and no reply needed if it is not a fit. If you spot something wrong in the
> [class] guide, tell me and I will fix it.
>
> [owner's name], Forever Sixty

## Creators

Found by web search on 3 November 2026. Subscriber counts were read off each YouTube channel page
that day. "Forever" means we found Forever videos from them dated September to November 2026.
Public channel links only.

| # | Creator | Channel | Subs | Forever | Why they fit | Link slug, target |
|---|---|---|---|---|---|---|
| 1 | Kargoz | https://www.youtube.com/@kargozgaming | 123K | Yes | Forever tier lists, duo leveling, a BM hunter 1 to 60 guide. A min-maxer audience. | `yt-kargoz` → /guides/hunter/beast-mastery |
| 2 | NickTew | https://www.youtube.com/@NickTew | 240K | Yes | Shaman 1 to 60 leveling guide with talents, rotation and gear. | `yt-nicktew` → /guides/shaman/enhancement |
| 3 | WillE | https://www.youtube.com/@WillEMMO | 197K | Yes | A class-by-class Forever series and an all-class overview. | `yt-wille` → /guides |
| 4 | Crix Forever | https://www.youtube.com/@crixforeverguides | 44.1K | Yes | Warlock specialist who tests things properly. A theorycraft audience that will check our numbers. | `yt-crix` → /guides/warlock |
| 5 | Tiqqle | https://www.youtube.com/@Tiqqle | 36.9K | Yes | Forever shaman leveling guide; streams on Twitch. | `yt-tiqqle` → /guides/shaman |
| 6 | Bajheera | https://www.youtube.com/@BajheeraWoW | 464K | Yes | Warrior specialist; a Forever Protection warrior leveling build. At the top of the size range. | `yt-bajheera` → /guides/warrior/protection |
| 7 | Venruki | https://www.youtube.com/@Venruki | 109K | Yes | Mage specialist: Forever mage changes, races, datamining. | `yt-venruki` → /guides/mage |
| 8 | hammerdance | https://www.youtube.com/@hammerdancegaming | 30.4K | Yes | Forever news and class guides; guest on the official Forever podcast. Note: `@hammerdance` is a different channel. | `yt-hammerdance` → /guides |
| 9 | SignsOfKelani | https://www.youtube.com/@SignsOfKelani | 281K | Yes | WoW guides and how-tos, with Forever class rework coverage. | `yt-signsofkelani` → /guides |
| 10 | Supatease | https://www.youtube.com/@Supatease | 73.3K | Yes | Educational and news; Forever class change breakdowns. | `yt-supatease` → /guides |
| 11 | Frostadamus | https://www.youtube.com/@Frostadamus | 79.7K | Yes | Long-running Classic channel with Forever launch videos. | `yt-frostadamus` → /bis |
| 12 | Icosiol | https://www.youtube.com/@Icosiol | 20K | Yes | Forever-only channel on leveling, dungeons and professions. A strong fit for leveling BiS. | `yt-icosiol` → /bis |
| 13 | Savix | https://www.youtube.com/@Savixirl | 227K | Yes | PvP; guest on the official Forever podcast. Fits the planner and the sim more than the guides. | `yt-savix` → /planner |
| 14 | Tips Out | https://www.youtube.com/@TipsOutBaby | 57.9K | Yes | Classic duel scene; ran the Forever Open duel tournament. | `yt-tipsout` → /planner |
| 15 | Shobek | https://www.youtube.com/@Shobek | 154K | Yes | Classic streamer, guest on the official podcast. Less guide-focused. | `yt-shobek` → / |
| 16 | NOHITJEROME | https://www.youtube.com/@NOHITJEROME | 49.1K | Yes | Full-time Classic streamer, Forever winners and losers videos. **Caution:** parts of r/classicwow dislike him after the April lockdown. | `yt-nohitjerome` → /tiers |
| 17 | Aeghis | https://www.youtube.com/@Aeghis | 24.9K | Not checked | Mage specialist with daily educational videos. | `yt-aeghis` → /guides/mage |
| 18 | HydraMist | https://www.youtube.com/@HydramistTV (old channel @HydraMist, 67.4K) | 13K | Partly | Priest specialist. Listed among Forever broadcasters on one fan site; no Forever videos found. | `yt-hydramist` → /guides/priest |
| 19 | Jokerd | https://www.youtube.com/@Jokerdtv | 43.1K | Not checked | Hardcore Classic streamer; mostly on Twitch. | `yt-jokerd` → /guides |
| 20 | Taliesin & Evitel | https://www.youtube.com/@TaliesinEvitel | 295K | Yes | News and the Weekly Reset podcast. Commentary rather than guides; better for a launch mention than for a guide link. | `yt-taliesin` → / |

Start with 1 to 12: they make class guides, so the offer fits what they already do. Check 17 to 19
for Forever content before writing to them.

**Near misses:** DoubleZug (15.2K; Forever Holy paladin and Resto druid videos) and Ahlaundoh (18.2K;
writes a Forever warrior leveling guide) fit well but fall under the size range. Both are worth a
message if the owner drops the floor. CaptainCoach made one Forever druid video, but the channel is
mostly Marvel Rivals.

Sources: each channel page above, plus https://mobalytics.gg/wow-forever/creator-class-guides,
https://mobalytics.gg/wow-forever/leveling-guides, https://forg.gg/en/forever/videos and the Forever
podcast episode pages on bluetracker.gg and wowhead.com.
