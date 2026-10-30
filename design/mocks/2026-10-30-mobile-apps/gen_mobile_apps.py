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
            cell("<b>F2.6 Item tooltip.</b> Tap an item: the sheet, medium detent.", H.home_tooltip(p)),
            cell("<b>F2.7 Your gear.</b> All 17 slots, worn and the pick.", H.gear_full(p))]),
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
            cell("<b>F4.11 Save failed.</b> Late was pressed; the old state returns.", R.raid_detail(p, "failed")),
            cell("<b>F4.12 Moved.</b> In on a second character benches the first (R3).", R.raid_detail(p, "moved"))]),
        f"f4-tab-states-{p}": ("F4 Raids tab states", [
            cell("<b>F4.13 No guild.</b>", R.raids_tab(p, "noguild")),
            cell("<b>F4.14 Waiting for approval.</b>", R.raids_tab(p, "waiting")),
            cell("<b>F4.15 No raid planned.</b>", R.raids_tab(p, "empty")),
            cell("<b>F4.16 Error.</b>", R.raids_tab(p, "error")),
            cell("<b>F4.17 Offline.</b> Segments disabled.", R.raids_tab(p, "offline"))]),
        f"f5-sims-{p}": ("F5 sims", [
            cell("<b>F5.1 Sims, Premium.</b> Server runs with progress, then your sims.", M.sims(p)),
            cell("<b>F5.2 Sims, free.</b> No server section, no prompt.", M.sims(p, premium=False)),
            cell("<b>F5.3 Result, Premium.</b>", M.sim_result(p)),
            cell("<b>F5.4 Result, free, link region.</b> One muted See plans.", M.sim_result(p, premium=False)),
            cell("<b>F5.5 Result, free, other regions.</b> Nothing about Premium.", M.sim_result(p, premium=False, link_region=False))]),
        f"f6-talents-{p}": ("F6 talents, read-only", [
            cell("<b>F6.1 Matches the guide.</b>", M.talents(p)),
            cell("<b>F6.2 Differs.</b> Dashed ranks differ from the guide.", M.talents(p, "differs"))]),
        f"f7-push-{p}": ("F7 notifications as delivered", (
            [cell("<b>F7.1 Lock screen.</b> raid.starting (time sensitive), raid.posted, sim.finished.", x) if i == 0 else
             cell("<b>F7.2 Long-press raid.posted.</b> In, Late, Out run in the background.", x) if i == 1 else
             cell("<b>F7.3 After In.</b> The thread's notification is replaced with the answer; loot.awarded.", x)
             for i, x in enumerate(M.lock_board_ios())] if p == "ios" else
            [cell("<b>F7.1 Shade.</b> raid.posted expanded with In, Late, Out; raid.starting; sim.finished.", x) if i == 0 else
             cell("<b>F7.3 After In, and the spec case.</b> Unknown spec: one action that opens the ask.", x)
             for i, x in enumerate(M.lock_board_and())])),
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
            cell("<b>XL.4 Notification settings.</b>", M.notif_settings(p, xl=True, sections=2))]),
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
