"""Boards for design/specs/2026-10-30-mobile-apps.md. Run from the repo root:
python3 design/mocks/2026-10-30-mobile-apps/gen_mobile_apps.py && node design/mocks/2026-10-30-mobile-apps/render.mjs
One board per feature and platform: iPhone 17 (402 x 874 pt, SwiftUI idiom) and Pixel 9 (412 x 915 dp, Material 3
idiom), plus the large-text boards (iOS AX5, Android font scale 2.0) and the raid page side by side."""
import sys
from pathlib import Path

HERE = Path(__file__).resolve().parent
sys.path.insert(0, str(HERE))
from kit import DEVICES, board, cell  # noqa: E402
import screens_home as H  # noqa: E402
import screens_more as M  # noqa: E402
import screens_raids as R  # noqa: E402
import screens_talents as T  # noqa: E402

RENDERS = HERE / "renders"
NAMES = {"ios": "iPhone 17", "and": "Pixel 9"}


def boards(p: str) -> dict[str, tuple[str, list[str]]]:
    n = NAMES[p]
    return {
        f"f1-signin-{p}": ("F1 sign-in and the switcher", [
            cell("<b>F1.1 Signed out.</b> Home tab, first launch. No onboarding.", H.signed_out(p)),
            cell("<b>F1.2 System sign-in.</b> " + ("ASWebAuthenticationSession consent." if p == "ios" else "Custom Tab on Battle.net."), H.consent(p)),
            cell("<b>F1.3 Import.</b> After the code exchange, until /v1/me answers.", H.importing(p)),
            cell("<b>F1.4 Sign-in failed.</b> Back on F1.1 with the error.", H.signin_error(p)),
            cell("<b>F1.5 Switcher.</b> The header selector opened: one way to switch.", H.switcher(p))]),
        f"f2-home-{p}": ("F2 home, the answer and its states", [
            cell("<b>F2.1 Default.</b> The answer and the top 5 upgrades above the tab bar.", H.home(p)),
            cell("<b>F2.2 Loading.</b> Skeletons hold every height.", H.home(p, "loading")),
            cell("<b>F2.3 Offline.</b> The cached answer with its stamp.", H.home(p, "offline")),
            cell("<b>F2.4 No gear export.</b> Grimtusk, never synced by the addon.", H.home(p, "nogear")),
            cell("<b>F2.5 Error.</b> The answer failed; the rest still reads.", H.home(p, "error"))]),
        f"f2-items-{p}": ("F2 item tooltip and worn vs best in slot", [
            cell("<b>F2.6 Item tooltip.</b> Binds, armor type and Unique from the client tables.", H.home_tooltip(p)),
            cell("<b>F2.7 Your gear.</b> All 17 slots: worn, then → the pick by name and icon.", H.gear_full(p))]),
        f"f3-bis-{p}": ("F3 best in slot", [
            cell("<b>F3.1 Default.</b> Your spec, range and faction, marked Yours.", H.bis(p)),
            cell("<b>F3.2 Picker.</b> Class, spec, range and faction in one sheet.", H.bis_picker(p)),
            cell("<b>F3.3 Offline.</b> Static lists are cached per build.", H.bis(p, "offline"))]),
        f"f4-raids-{p}": ("F4 raids, the headline", [
            cell("<b>F4.1 Raids tab.</b> The next raid and your sign-up first.", R.raids_tab(p)),
            cell("<b>F4.2 Raid page.</b> Signed In; the answer above the fold.", R.raid_detail(p)),
            cell("<b>F4.3 Spec ask.</b> In pressed with no spec saved: pending until the pick.", R.raid_spec_ask(p)),
            cell("<b>F4.4 Reminder ask.</b> After the first In: notification permission, in context.", R.raid_detail(p, "ask")),
            cell("<b>F4.5 More menu.</b> Calendar, share, the web.", R.raid_menu(p))]),
        f"f4-raid-full-{p}": ("F4 raid page, the whole scroll", [
            cell(f"<b>F4.6 Raid page, full length.</b> {n}, scrolled top to bottom.", R.raid_detail(p, full=True, height=2860))]),
        f"f4-states-{p}": ("F4 raid page states", [
            cell("<b>F4.7 Unsigned.</b> No segment checked.", R.raid_detail(p, "unsigned")),
            cell("<b>F4.8 Locked.</b>", R.raid_detail(p, "locked")),
            cell("<b>F4.9 Cancelled.</b>", R.raid_detail(p, "cancelled")),
            cell("<b>F4.10 Offline.</b> Read-only, stamped.", R.raid_detail(p, "offline")),
            cell("<b>F4.11 Save failed.</b> Late was pressed and refused; In, the old state, is back.", R.raid_detail(p, "failed")),
            cell("<b>F4.12 Moved.</b> In on Lightningbolltz benched Vorgash (R3); the line is on his block.", R.raid_detail(p, "moved"))]),
        f"f4-tab-states-{p}": ("F4 Raids tab states", [
            cell("<b>F4.13 No guild.</b>", R.raids_tab(p, "noguild")),
            cell("<b>F4.14 Waiting for approval.</b>", R.raids_tab(p, "waiting")),
            cell("<b>F4.15 No raid planned.</b>", R.raids_tab(p, "empty")),
            cell("<b>F4.16 Error.</b>", R.raids_tab(p, "error")),
            cell("<b>F4.17 Offline.</b> Segments disabled.", R.raids_tab(p, "offline"))]),
        f"f5-sims-{p}": ("F5 sims, today's API", [
            cell("<b>F5.1 Sims, Premium.</b> Running bulk with bar, running plain without, finished unseen, failed with Run again; footer pinned.", M.sims(p)),
            cell("<b>F5.2 Sims, free.</b> No server section, no prompt.", M.sims(p, "free")),
            cell("<b>F5.2b Lapsed, after a 402.</b> Old server rows stay; the line and a gold underlined See plans (link region only).", M.sims(p, "lapsed")),
            cell("<b>F5.3 Result, plain run.</b> No rerun, no Premium line.", M.sim_result(p)),
            cell("<b>F5.3b Result, Top Gear.</b> The row headline is the answer; a gold corner marks the changed slot.", M.sim_result(p, "gear"))]),
        f"f5-states-{p}": ("F5 sims states", [
            cell("<b>F5.6 Loading.</b> Skeletons at real row heights.", M.sims(p, "loading")),
            cell("<b>F5.7 Empty.</b> Premium, nothing yet; no footer.", M.sims(p, "empty")),
            cell("<b>F5.8 Error.</b> Server runs failed alone; Your sims keeps its rows.", M.sims(p, "error")),
            cell("<b>F5.8b Offline.</b> Cached rows, no polling, Run again disabled.", M.sims(p, "offline")),
            cell("<b>F5.9 Next page loading.</b> Spinner row under the last row.", M.sims(p, "paging")),
            cell("<b>F5.9b Next page failed.</b> sims.failedLoad with Try again as the last row.", M.sims(p, "pagefail"))]),
        f"f5-rerun-{p}": ("F5 Run again states (failed Server runs row)", [
            cell("<b>R1 Starting.</b> Disabled, busy.", M.rerun(p, "starting")),
            cell("<b>R2 Started.</b> Disabled until the list re-reads.", M.rerun(p, "started")),
            cell("<b>R3 Could not start.</b> Network or 5xx.", M.rerun(p, "failed")),
            cell("<b>R4 Lapsed, link region.</b> Gold underlined See plans.", M.rerun(p, "lapsed")),
            cell("<b>R5 Lapsed, other regions.</b> The sentence alone.", M.rerun(p, "lapsed-plain")),
            cell("<b>R6 Offline.</b> Disabled with its sentence.", M.rerun(p, "offline"))]),
        f"f5-result-states-{p}": ("F5.3 result states", [
            cell("<b>F5.3.1 Loading.</b>", M.sim_result(p, state="loading")),
            cell("<b>F5.3.2 Did not load.</b> sims.resultFailed with Try again.", M.sim_result(p, state="failed")),
            cell("<b>F5.3.3 Offline.</b> Strip over a cached result.", M.sim_result(p, state="offline")),
            cell("<b>F5.3.4 No gear.</b> The section is left out.", M.sim_result(p, state="nogear")),
            cell("<b>F5.3.5 Unknown item.</b> Question mark, not tappable.", M.sim_result(p, "gear", state="unknown"))]),
        f"f6-talents-{p}": ("F6 talents, read-only", [
            cell("<b>F6.1 Matches.</b> Tab totals are summed from the drawn build: 30 + 3 + 18 = 51. No dashes, no unspent line.", T.screen_content(p, "matches")),
            cell("<b>F6.2 Differs.</b> 2 of 48 points differ; 4 cells are dashed (+2 over the guide, −2 under); the Restoration tab carries a dot; 3 unspent.", T.screen_content(p, "differs")),
            cell("<b>F6.3 Guide build.</b> The same screen on the Guide segment: totals 30, 3, 18; no dashes; the tap hint.", T.screen_content(p, "differs", on=0))]),
        f"f6-rank-{p}": ("F6 the rank sheet", [
            cell("<b>F6.4 Bottom row.</b> Elemental Mastery (tier 7). The tree scrolls so the tapped cell stays above the sheet, ringed.",
                 T.with_sheet(p, "differs", 0, 1, T.named("Elemental Mastery"), 0, 0)),
            cell("<b>F6.5 Right column, differs.</b> Call of Flame: your rank 1, the guide's 3, so the guide line shows.",
                 T.with_sheet(p, "differs", 0, 1, T.named("Call of Flame"), 1, 3)),
            cell("<b>F6.6 Guide build, maxed.</b> Elemental Fury 5/5: current text only, no Next rank.",
                 T.with_sheet(p, "differs", 0, 0, T.named("Elemental Fury"), 5, None))]),
        f"f6-states-{p}": ("F6 talents states", [
            cell("<b>F6.7 Loading.</b> Skeletons at the real sizes; the tree region is reserved at 7 tiers.", T.screen_content(p, "loading")),
            cell("<b>F6.8 No export.</b> Guide build only; no segments; the export line.", T.screen_content(p, "noexport")),
            cell("<b>F6.9 No build.</b> Your build alone (level under 20, or no list for the spec).", T.screen_content(p, "nobuild")),
            cell("<b>F6.10 No spec.</b> Your build alone; the planner opens in the browser.", T.screen_content(p, "nospec")),
            cell("<b>F6.11 Offline.</b> The strip, then the cached screen.", T.screen_content(p, "differs", offline=True)),
            cell("<b>F6.12 Error.</b> The talent file did not load and nothing is cached.", T.screen_content(p, "error"))]),
        f"f6-xl-{p}": ("F6 largest text: " + ("iOS AX5" if p == "ios" else "Android font scale 2.0"), [
            cell("<b>XL.1 Top.</b> Status wraps, tabs stack, segments stack; nothing pinned.", T.screen_content(p, "differs", xl=True)),
            cell("<b>XL.2 The tree.</b> Scrolled: icons stay 48, badges cap at 14.", T.screen_content(p, "differs", xl=True, scroll=T.XL_TREE_SCROLL[p])),
            cell("<b>XL.3 Rank sheet, large detent.</b> Opens full height; the body scrolls.", T.with_sheet(p, "differs", 0, 1, T.named("Call of Flame"), 1, 3, xl=True)),
            cell("<b>XL.4 The same sheet, scrolled to the end.</b> The last line and the inset are reachable.", T.with_sheet(p, "differs", 0, 1, T.named("Call of Flame"), 1, 3, xl=True, end=True))]),
        f"f7-push-{p}": ("F7 notifications as delivered", (
            [cell("<b>F7.1 Lock screen.</b> raid.starting at invites minus the lead, raid.posted, sim.finished.", x) if i == 0 else
             cell("<b>F7.2 Long-press raid.posted.</b> In, Late, Out run in the background.", x) if i == 1 else
             cell("<b>F7.3 After In.</b> The answer replaces the question; raid.summary after the log.", x)
             for i, x in enumerate(M.lock_board_ios())] if p == "ios" else
            [cell("<b>F7.1 Shade.</b> raid.posted with In, Late, Out; raid.starting; sim.finished.", x) if i == 0 else
             cell("<b>F7.3 After In; the spec case; raid.summary.</b> Unknown spec: Choose a spec, and Out.", x)
             for i, x in enumerate(M.lock_board_and())])),
        f"f4-officer-{p}": ("F4 officer tools on the raid page (officers only)", [
            cell("<b>F4.18 Officer, open raid.</b> Lock sign-ups and Copy the list.", R.raid_detail(p, "officer")),
            cell("<b>F4.19 Officer roster.</b> Each row ends in its place.", R.officer_roster(p)),
            cell("<b>F4.20 Move.</b> In to Bench, or off the raid.", R.officer_roster(p, menu=True)),
            cell("<b>F4.21 Locked.</b> The roster went out; Unlock sign-ups.", R.raid_detail(p, "officer-locked", overlay=R.toast_locked(p)))]),
        f"f4-night-{p}": ("F4 Your raid night, the after-raid summary", [
            cell("<b>F4.22 Your raid night, Premium.</b> Opened from raid.summary.", R.night(p)),
            cell("<b>F4.23 Your raid night, free.</b> No coach section, no prompt.", R.night(p, premium=False))]),
        f"f7-moved-{p}": ("F7 raid.roster_moved, time sensitive", [
            cell(f"<b>F7.7 Moved {'In' if i == 0 else 'to Bench' if i == 1 else 'off, and added'}.</b>" if p == "ios" else f"<b>F7.7 {'Moved In' if i == 0 else 'Moved to Bench; removed'}.</b>", x)
            for i, x in enumerate(M.moved_boards(p))]),
        f"f7-settings-{p}": ("F7 notification settings", [
            cell("<b>F7.4 Settings.</b> Per event and channel; officer and Premium sections only when they apply.", M.notif_settings(p, sections=3)),
            cell("<b>F7.5 Channel picker.</b> " + ("A Menu on the value." if p == "ios" else "A dropdown menu on the value."), M.notif_settings(p, picker=True, sections=3)),
            cell("<b>F7.6 Permission denied.</b> The one place we send people to Settings.", M.notif_settings(p, denied=True, sections=2))]),
        f"f8-you-{p}": ("F8 You, Premium status and See plans", [
            cell("<b>F8.1 Free, link region.</b> Plan · Free · See plans.", M.you(p, "free-link")),
            cell("<b>F8.2 Free, other regions.</b> Plan · Free, nothing else.", M.you(p, "free-other")),
            cell("<b>F8.3 Premium.</b>", M.you(p, "premium")),
            cell("<b>F8.4 Plan.</b> What your plan includes.", M.plan_detail(p)),
            cell("<b>F1.6 Devices.</b> Revoke confirmation.", M.devices(p, confirm=True))]),
        f"xl-{p}": ("Largest text: " + ("iOS AX5" if p == "ios" else "Android font scale 2.0"), [
            cell("<b>XL.1 Home.</b> Gains move under the name; three rows.", H.home(p, xl=True)),
            cell("<b>XL.2 Raid page.</b> Segments become a 2 x 2 grid.", R.raid_detail(p, xl=True)),
            cell("<b>XL.3 Switcher.</b> Full names wrap, never cut.", H.switcher(p, xl=True)),
            cell("<b>XL.4 Notification settings.</b>", M.notif_settings(p, xl=True, sections=2)),
            cell("<b>XL.5 Sims.</b> Titles wrap whole; state and age stack; Run again grows.", M.sims(p, xl=True)),
            cell("<b>XL.6 Sim result.</b> Figure grows; the 6-up grid holds.", M.sim_result(p, xl=True))]),
    }


def main() -> None:
    RENDERS.mkdir(exist_ok=True)
    out = {}
    for p in ("ios", "and"):
        for name, (title, cells) in boards(p).items():
            out[name] = board(title, DEVICES[p]["label"], cells)
    out["f4-pair"] = board("F4 raid page, both platforms side by side", "iPhone 17 and Pixel 9", [
        cell("<b>iOS</b> Raids tab", R.raids_tab("ios")), cell("<b>Android</b> Raids tab", R.raids_tab("and")),
        cell("<b>iOS</b> Raid page", R.raid_detail("ios")), cell("<b>Android</b> Raid page", R.raid_detail("and"))])
    for name, html in out.items():
        (RENDERS / f"{name}.html").write_text(html)
        print(name)


if __name__ == "__main__":
    main()
