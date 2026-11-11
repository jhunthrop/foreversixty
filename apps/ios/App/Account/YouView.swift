import Auth
import DesignSystem
import Models
import SwiftUI

/// What the You tab shows, already worded (spec F8.1: the account head; Plan, Character and
/// Devices; Privacy, Terms and About; Sign out; Delete account; the version line). F7 and F8 add
/// the Notifications, Characters and Discord rows.
struct YouSummary: Equatable {
    let battletag: String
    /// Lane F8: the plan in the site's words.
    let plan: PlanSummary
    /// The current character and where it plays (`Obnoxious Yell`, `US · PvP`).
    let character: CharacterLine?
    /// `This iPhone and 2 more`, nil until the device list answers.
    let devices: String?
    /// `Version 0.1.0 (1) · data 1.60.1.70658`.
    let version: String

    struct CharacterLine: Equatable {
        let name: String
        let className: String?
        let place: String
    }

    init(battletag: String, plan: PlanSummary, character: CharacterLine?, devices: String?, version: String) {
        self.battletag = battletag
        self.plan = plan
        self.character = character
        self.devices = devices
        self.version = version
    }

    init(me: Me, current: AccountCharacter?, deviceCount: Int?, version: String, now: Date = .now) {
        self.init(
            battletag: me.user.battletag ?? "",
            plan: PlanSummary(me: me, now: now),
            character: current.map { CharacterLine(name: $0.name, className: $0.className, place: AccountCopy.place($0)) },
            devices: deviceCount.map { Copy.youDevicesSub(device: Copy.accountDeviceWord, n: max(0, $0 - 1)) },
            version: version
        )
    }
}

enum YouRoute: Hashable {
    case plan
    case devices
    case deleteAccount
    case about
    /// Lane F14: You › Live surfaces.
    case liveSurfaces
    /// Lane F16: You › Calendar.
    case calendar
    /// Lane F15: You › Offline data.
    case offlineData
}

/// The You tab. Pure: the snapshot tests render it with a summary.
struct YouView: View {
    let summary: YouSummary
    /// Lane F8: whether the store allows the link to web checkout (`PlansLink`).
    var mayLink = false
    /// Lane F14: the Live surfaces row, behind the widgets switch.
    var showsLiveSurfaces = false
    /// Lane F16: the Calendar row, while `guild_extras` is on.
    var showsCalendar = false
    /// Lane F15: the Offline data row, while `system` is on.
    var showsOfflineData = false
    let onSignOut: () -> Void
    @Environment(\.openURL) private var openURL
    @State private var confirmingSignOut = false

    var body: some View {
        ScrollView {
            VStack(alignment: .leading, spacing: 0) {
                // The head names the screen and the account on the band (N3); the bar has no title at rest.
                BandHead(title: Copy.youTitle) {
                    Text(summary.battletag).font(ScaledType.name).foregroundStyle(FSColor.strong)
                    BandFact(Copy.youSignedInWith)
                }
                .accessibilityElement(children: .combine)
                rows
            }
        }
        .screenGround()
        .pushedBandBar(Copy.youTitle)
        .confirmationDialog(
            Copy.youSignOutConfirm(device: Copy.accountDeviceWord),
            isPresented: $confirmingSignOut,
            titleVisibility: .visible
        ) {
            Button(Copy.accountSignOutAction, role: .destructive, action: onSignOut)
            Button(Copy.commonCancel, role: .cancel) {}
        }
    }

    private var rows: some View {
        VStack(alignment: .leading, spacing: FSSpace.space5) {
            VStack(spacing: 0) {
                PlanRow(plan: summary.plan, mayLink: mayLink)
                if let character = summary.character {
                    Divider().overlay(FSColor.lineSoft)
                    characterRow(character)
                }
                Divider().overlay(FSColor.lineSoft)
                NavigationLink(value: YouRoute.devices) {
                    AccountRow(icon: "iphone", title: Copy.youDevices, subtitle: summary.devices, chevron: true)
                }
                if showsLiveSurfaces {
                    Divider().overlay(FSColor.lineSoft)
                    NavigationLink(value: YouRoute.liveSurfaces) {
                        AccountRow(icon: "rectangle.stack", title: Copy.liveSettings, chevron: true)
                    }
                }
                if showsCalendar {
                    Divider().overlay(FSColor.lineSoft)
                    NavigationLink(value: YouRoute.calendar) {
                        AccountRow(icon: "calendar", title: Copy.calendarTitle, subtitle: Copy.calendarYouRowSub, chevron: true)
                    }
                }
                if showsOfflineData {
                    Divider().overlay(FSColor.lineSoft)
                    NavigationLink(value: YouRoute.offlineData) {
                        AccountRow(icon: "arrow.down.circle", title: Copy.offlineDataTitle, subtitle: Copy.offlineDataYouRowSub, chevron: true)
                    }
                }
            }
            .quietCard()
            VStack(spacing: 0) {
                Button { openURL(AccountLinks.privacy) } label: { AccountRow(icon: "hand.raised", title: Copy.youPrivacy, external: true) }
                Divider().overlay(FSColor.lineSoft)
                Button { openURL(AccountLinks.terms) } label: { AccountRow(icon: "doc.text", title: Copy.youTerms, external: true) }
                Divider().overlay(FSColor.lineSoft)
                NavigationLink(value: YouRoute.about) {
                    AccountRow(icon: "info.circle", title: Copy.youAbout, chevron: true)
                }
            }
            .quietCard()
            Button(Copy.youSignOut) { confirmingSignOut = true }
                .buttonStyle(BrandButtonStyle(.quiet, height: 48))
            // The destructive action on a panel, as Android's AccountCard: wipe reads on the dark, not loose on the ground.
            NavigationLink(value: YouRoute.deleteAccount) {
                Text(Copy.sharedAccountDeleteTitle)
            }
            .buttonStyle(BrandButtonStyle(.destructive, height: 48))
            .padding(FSSpace.space4)
            .quietCard()
            Text(summary.version).font(ScaledType.label).foregroundStyle(FSColor.mutedGround)
                .frame(maxWidth: .infinity)
        }
        .buttonStyle(.plain)
        .padding(FSSpace.space4)
    }

    private func characterRow(_ character: YouSummary.CharacterLine) -> some View {
        HStack(spacing: FSSpace.space3) {
            ClassCrest(className: character.className, size: 32)
            VStack(alignment: .leading, spacing: 2) {
                Text(Copy.accountYouCharacter).font(ScaledType.bodyStrong).foregroundStyle(FSColor.strong)
                Text("\(Text(character.name).foregroundStyle(ClassColor.of(character.className)))\(Text(" · \(character.place)").foregroundStyle(FSColor.muted))")
                    .font(ScaledType.meta)
            }
            Spacer(minLength: 0)
        }
        .padding(.horizontal, FSSpace.space4)
        .frame(minHeight: 56)
        .accessibilityElement(children: .combine)
    }

}

/// You › About Forever Sixty: the version, the build and the site footer's Blizzard line.
struct AboutView: View {
    let version: String

    var body: some View {
        ScrollView {
            VStack(alignment: .leading, spacing: 0) {
                BandHead(title: Copy.youAbout)
                about
            }
        }
        .screenGround()
        .pushedBandBar(Copy.youAbout)
    }

    private var about: some View {
        VStack(alignment: .leading, spacing: FSSpace.space4) {
            Seal(size: 64)
            Text(Copy.appName).font(ScaledType.wordmark).tracking(2).foregroundStyle(FSColor.strong)
            Text(version).font(ScaledType.body).foregroundStyle(FSColor.text)
            Text(Copy.aboutNotAffiliated).font(ScaledType.meta).foregroundStyle(FSColor.mutedGround)
                .fixedSize(horizontal: false, vertical: true)
        }
        .frame(maxWidth: .infinity, alignment: .leading)
        .padding(FSSpace.space4)
    }
}

/// One list row: a gold tile, the title and its sub line, a trailing value, tag or chevron.
struct AccountRow: View {
    let icon: String
    let title: String
    var subtitle: String?
    var trailing: String?
    var tag: String?
    var chevron = false
    /// Opens the browser: the external glyph instead of the chevron.
    var external = false
    var tint: Color = FSColor.strong
    /// The icon tile grows with the text, so a large glyph never spills over the title.
    @ScaledMetric(relativeTo: .callout) private var tile: CGFloat = 32
    @Environment(\.dynamicTypeSize) private var typeSize

    var body: some View {
        HStack(spacing: FSSpace.space3) {
            // Accessibility sizes: the tile sits above the words, so they keep the row's width
            // and never break inside a word.
            (typeSize.isAccessibilitySize
                ? AnyLayout(VStackLayout(alignment: .leading, spacing: FSSpace.space2))
                : AnyLayout(HStackLayout(spacing: FSSpace.space3))) {
                Image(systemName: icon)
                    .font(.callout)
                    .foregroundStyle(tint == FSColor.strong ? FSColor.gold : tint)
                    .frame(width: tile, height: tile)
                    .background(FSColor.gold.opacity(0.12), in: .rect(cornerRadius: FSRadius.legacyPill))
                    .accessibilityHidden(true)
                VStack(alignment: .leading, spacing: 2) {
                    Text(title).font(ScaledType.bodyStrong).foregroundStyle(tint)
                    if let subtitle {
                        Text(subtitle).font(ScaledType.meta).foregroundStyle(FSColor.muted)
                    }
                }
            }
            Spacer(minLength: FSSpace.space2)
            if let tag { TagView(text: tag, color: FSColor.gold) }
            if let trailing {
                Text(trailing).font(ScaledType.meta).foregroundStyle(FSColor.muted).multilineTextAlignment(.trailing)
            }
            if external {
                Image(systemName: "arrow.up.right.square").font(.footnote.weight(.semibold)).foregroundStyle(FSColor.muted)
                    .accessibilityHidden(true)
            }
            if chevron {
                Image(systemName: "chevron.right").font(.footnote.weight(.semibold)).foregroundStyle(FSColor.muted)
                    .accessibilityHidden(true)
            }
        }
        .padding(.horizontal, FSSpace.space4)
        .padding(.vertical, FSSpace.space2)
        .frame(minHeight: 56)
        .contentShape(.rect)
        .accessibilityElement(children: .combine)
    }
}
