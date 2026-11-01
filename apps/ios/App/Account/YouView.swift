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
            VStack(alignment: .leading, spacing: FSSpacing.space5) {
                accountHead
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
                NavigationLink(value: YouRoute.deleteAccount) {
                    Text(Copy.accountDeleteRow)
                        .font(ScaledType.bodyStrong)
                        .foregroundStyle(FSColor.wipe)
                        .frame(maxWidth: .infinity, minHeight: 44)
                        .contentShape(.rect)
                }
                .padding(.top, FSSpacing.space2)
                Text(summary.version).font(ScaledType.micro).foregroundStyle(FSColor.muted)
                    .frame(maxWidth: .infinity)
            }
            .buttonStyle(.plain)
            .padding(FSSpacing.space4)
        }
        .background(FSColor.bg)
        .navigationTitle(Copy.youTitle)
        .confirmationDialog(
            Copy.youSignOutConfirm(device: Copy.accountDeviceWord),
            isPresented: $confirmingSignOut,
            titleVisibility: .visible
        ) {
            Button(Copy.accountSignOutAction, role: .destructive, action: onSignOut)
            Button(Copy.commonCancel, role: .cancel) {}
        }
    }

    private func characterRow(_ character: YouSummary.CharacterLine) -> some View {
        HStack(spacing: FSSpacing.space3) {
            ClassCrest(className: character.className, size: 32)
            VStack(alignment: .leading, spacing: 2) {
                Text(Copy.accountYouCharacter).font(ScaledType.bodyStrong).foregroundStyle(FSColor.strong)
                Text("\(Text(character.name).foregroundStyle(ClassColor.of(character.className)))\(Text(" · \(character.place)").foregroundStyle(FSColor.muted))")
                    .font(ScaledType.meta)
            }
            Spacer(minLength: 0)
        }
        .padding(.horizontal, FSSpacing.space4)
        .frame(minHeight: 56)
        .accessibilityElement(children: .combine)
    }

    /// Lane F8: the gold Premium tag sits beside the name, or under it when the name needs the room.
    private var accountHead: some View {
        ViewThatFits(in: .horizontal) {
            HStack(spacing: FSSpacing.space2) {
                identity
                Spacer(minLength: FSSpacing.space2)
                premiumTag
            }
            VStack(alignment: .leading, spacing: FSSpacing.space2) {
                identity
                premiumTag
            }
        }
        .accessibilityElement(children: .combine)
    }

    @ViewBuilder
    private var premiumTag: some View {
        if summary.plan.isPremium { TagView(text: Copy.planPremium, color: FSColor.gold) }
    }

    private var identity: some View {
        HStack(spacing: FSSpacing.space3) {
            Image(systemName: "person")
                .font(.title3)
                .foregroundStyle(FSColor.gold)
                .frame(width: 44, height: 44)
                .background(FSColor.gold.opacity(0.12), in: .rect(cornerRadius: FSRadius.control))
                .accessibilityHidden(true)
            VStack(alignment: .leading, spacing: 2) {
                Text(summary.battletag).font(ScaledType.name).foregroundStyle(FSColor.strong)
                Text(Copy.youSignedInWith).font(ScaledType.meta).foregroundStyle(FSColor.muted)
            }
        }
    }
}

/// You › About Forever Sixty: the version, the build and the site footer's Blizzard line.
struct AboutView: View {
    let version: String

    var body: some View {
        ScrollView {
            VStack(alignment: .leading, spacing: FSSpacing.space4) {
                Seal(size: 64)
                Text(Copy.appName).font(ScaledType.wordmark).textCase(.uppercase).tracking(2).foregroundStyle(FSColor.strong)
                Text(version).font(ScaledType.body).foregroundStyle(FSColor.text)
                Text(Copy.aboutNotAffiliated).font(ScaledType.meta).foregroundStyle(FSColor.muted)
                    .fixedSize(horizontal: false, vertical: true)
            }
            .frame(maxWidth: .infinity, alignment: .leading)
            .padding(FSSpacing.space4)
        }
        .background(FSColor.bg)
        .navigationTitle(Copy.youAbout)
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
        HStack(spacing: FSSpacing.space3) {
            // Accessibility sizes: the tile sits above the words, so they keep the row's width
            // and never break inside a word.
            (typeSize.isAccessibilitySize
                ? AnyLayout(VStackLayout(alignment: .leading, spacing: FSSpacing.space2))
                : AnyLayout(HStackLayout(spacing: FSSpacing.space3))) {
                Image(systemName: icon)
                    .font(.callout)
                    .foregroundStyle(tint == FSColor.strong ? FSColor.gold : tint)
                    .frame(width: tile, height: tile)
                    .background(FSColor.gold.opacity(0.12), in: .rect(cornerRadius: FSRadius.pill))
                    .accessibilityHidden(true)
                VStack(alignment: .leading, spacing: 2) {
                    Text(title).font(ScaledType.bodyStrong).foregroundStyle(tint)
                    if let subtitle {
                        Text(subtitle).font(ScaledType.meta).foregroundStyle(FSColor.muted)
                    }
                }
            }
            Spacer(minLength: FSSpacing.space2)
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
        .padding(.horizontal, FSSpacing.space4)
        .padding(.vertical, FSSpacing.space2)
        .frame(minHeight: 56)
        .contentShape(.rect)
        .accessibilityElement(children: .combine)
    }
}
