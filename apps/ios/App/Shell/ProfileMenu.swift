import Auth
import DesignSystem
import Models
import SwiftUI

/// What the profile menu shows (flows spec 3), worked out from the account so the view only draws.
struct ProfileMenuSummary: Equatable {
    struct Row: Equatable, Identifiable {
        let destination: MenuDestination
        let title: String
        let line: String

        var id: MenuDestination { destination }
    }

    /// The selected character; nil signed out.
    let you: AccountCharacter?
    let isMain: Bool
    /// `Fury Warrior · 58 · Officer`: the kind, the level and the guild rank (the line without a guild).
    let youLine: String
    /// `Updated 2 h ago from Battle.net`, nil before the first import.
    let stamp: String?
    /// The import is past `SyncAge.staleAfter`: the stamp takes the warning tone (row 12).
    let stale: Bool
    /// The last Battle.net refresh failed: this replaces the stamp, in wipe red with Try again.
    let syncError: String?
    /// The other characters at level 10 and up, then those below it.
    let others: [AccountCharacter]
    let lowLevel: [AccountCharacter]
    let tools: [Row]
    let account: [Row]
    let version: String
    /// The site footer's operator line, `© {year} Commish LLC d/b/a Forever Sixty`, in the year of `now`.
    let operatorLine: String

    /// The level the Switch character list starts at; the rest fold.
    static let switchLevel = 10

    struct AccountFacts: Equatable {
        /// The groups the player gets, each on or off (`GET /v1/me/notifications`); nil until loaded.
        var notifications: [NotificationGroupsSummary.Row]?
        var plan: String?
        var devices: Int?
        var offlineBytes: Int64?
        /// The newest signed-in device's name, for `3 signed in · newest iPhone 17`.
        var newestDevice: String?
    }

    init(me: Me?, current: AccountCharacter?, facts: AccountFacts, version: String, now: Date) {
        self.version = version
        operatorLine = Copy.menuLegalOperator(year: String(Calendar(identifier: .gregorian).component(.year, from: now)))
        guard let me, let current else {
            you = nil
            youLine = ""
            isMain = false
            stamp = nil
            stale = false
            syncError = nil
            others = []
            lowLevel = []
            tools = MenuDestination.signedOutTools.map(Self.tool)
            account = []
            return
        }
        you = current
        youLine = Self.youLine(current)
        isMain = current.key == me.mainCharacterKey
        stamp = me.bnetImportedAt.map { Copy.menuUpdated(age: RelativeAge(from: $0, to: now).phrase) }
        stale = me.bnetImportedAt.map { SyncAge.isStale($0, now: now) } ?? false
        syncError = current.build?.syncError.map { _ in Copy.menuSyncError }
        let rest = me.characters.filter { $0.key != current.key }
        others = rest.filter { ($0.level ?? 0) >= Self.switchLevel }
        lowLevel = rest.filter { ($0.level ?? 0) < Self.switchLevel }
        tools = MenuDestination.tools.map(Self.tool)
        account = MenuDestination.account.map { Self.account($0, facts: facts) }
    }

    private static func tool(_ destination: MenuDestination) -> Row {
        // Tools are labels alone (owner, 2026-11-09).
        let title = switch destination {
        case .guides: Copy.menuGuides
        case .tiers: Copy.menuTiers
        case .patchNotes: Copy.menuPatchNotes
        case .sims: Copy.menuSims
        case .reports: Copy.menuLogs
        case .planner: Copy.menuPlanner
        case .rankings: Copy.menuRankings
        default: ""
        }
        return Row(destination: destination, title: title, line: "")
    }

    private static func account(_ destination: MenuDestination, facts: AccountFacts) -> Row {
        switch destination {
        case .notifications:
            Row(destination: destination, title: Copy.menuNotifications, line: notificationsLine(facts.notifications))
        case .plan:
            Row(destination: destination, title: Copy.menuPlan, line: facts.plan ?? "")
        case .devices:
            Row(destination: destination, title: Copy.menuDevices, line: devicesLine(facts.devices, newest: facts.newestDevice))
        case .offlineData:
            Row(destination: destination, title: Copy.menuOfflineData, line: facts.offlineBytes.map(offlineLine) ?? "")
        default:
            Row(destination: destination, title: Copy.menuSettings, line: Copy.menuSettingsLine)
        }
    }

    /// `Fury Warrior · 58 · Officer` with a guild rank, else the character's line.
    static func youLine(_ character: AccountCharacter) -> String {
        guard let rank = character.guild?.rank, !rank.isEmpty else { return AccountCopy.characterLine(character) }
        let kind = [character.spec, character.className].compactMap { $0.map(AccountCopy.words) }.joined(separator: " ")
        return Copy.menuCharacterLineRank(kind: kind, n: character.level ?? 0, rank: AccountCopy.words(rank))
    }

    /// A Switch row's line: the character's line, then `item level 63` when the import has it.
    static func switchLine(_ character: AccountCharacter) -> String {
        let line = AccountCopy.characterLine(character)
        guard let itemLevel = character.itemLevel, itemLevel > 0 else { return line }
        return line + " · " + Copy.menuCharacterItemLevel(n: itemLevel)
    }

    /// `3 signed in · newest iPhone 17`; the count alone without a name.
    static func devicesLine(_ count: Int?, newest: String?) -> String {
        guard let count else { return "" }
        guard let newest, !newest.isEmpty else { return Copy.menuDevicesLine(n: count) }
        return Copy.menuDevicesLineNewest(n: count, name: newest)
    }

    /// The facts the Devices row reads from `GET /v1/devices`: how many, and the newest by sign-in.
    static func deviceFacts(_ devices: [Device], into facts: AccountFacts) -> AccountFacts {
        var updated = facts
        updated.devices = devices.count
        updated.newestDevice = devices.max { ($0.createdAt ?? .distantPast) < ($1.createdAt ?? .distantPast) }?.name
        return updated
    }

    /// What a player misses (round 6, row 1): `All on` when every group they get is on, `All off`
    /// when none is, else the groups that are off in their order, sentence case (`Logs, tiers off`).
    static func notificationsLine(_ rows: [NotificationGroupsSummary.Row]?) -> String {
        guard let rows, !rows.isEmpty else { return "" }
        let off = rows.filter { !$0.on }.map(\.group.title)
        guard let first = off.first else { return Copy.menuNotificationsAllOn }
        guard off.count < rows.count else { return Copy.menuNotificationsOff }
        let joined = ([first] + off.dropFirst().map { $0.lowercased() }).joined(separator: Self.listSeparator)
        return Copy.menuNotificationsSomeOff(groups: joined)
    }

    private static let listSeparator = ", "

    static func offlineLine(_ bytes: Int64) -> String {
        Copy.menuOfflineDataLine(size: ByteCountFormatter.string(fromByteCount: bytes, countStyle: .file))
    }
}

/// The menu's rows and links, handed in by the shell.
struct ProfileMenuActions {
    var pick: (AccountCharacter) -> Void = { _ in }
    /// The Battle.net refresh; false when it failed.
    var refresh: () async -> Bool = { true }
    var open: (MenuDestination) -> Void = { _ in }
    var signIn: () -> Void = {}
    var help: () -> Void = {}
    var privacy: () -> Void = {}
    /// `GET /v1/devices`, for the Devices row's newest device; nil keeps the summary's line.
    var loadDevices: () async -> [Device]? = { nil }
    /// `GET /v1/me/notifications`, for the Notifications row's value; nil keeps the summary's line.
    var loadNotifications: () async -> NotificationSettings? = { nil }
}

/// The profile menu's content in the round-5 look (mocks ios-menu, -one, -low, -out and m-ios-menu):
/// You on the painted band (the portrait, the name in the display face in the class colour, the
/// line, Main, the stamp with Refresh), then on the drawer's panel Switch character, Tools as labels
/// alone (owner ruling 2026-11-09), Account and the footer; signed out, Sign in and the tools that work
/// without an account.
struct ProfileMenuView: View {
    let summary: ProfileMenuSummary
    var signingIn = false
    var actions = ProfileMenuActions()
    @State private var showsLowLevel = false
    @State private var devices: [Device]?
    @State private var notifications: NotificationSettings?
    @State private var refresh = MenuRefresh.idle

    static let youPortrait: CGFloat = 56
    static let rowPortrait: CGFloat = 40
    /// The band runs under the status bar; You sits 18 under it (the mock's `--status`, 54).
    static let statusBar: CGFloat = 54
    static let side: CGFloat = 18
    static let youTop: CGFloat = 18
    static let youBottom: CGFloat = 12
    static let rowMin: CGFloat = 48
    static let rowPad: CGFloat = 7

    init(
        summary: ProfileMenuSummary, signingIn: Bool = false, showsLowLevel: Bool = false, refresh: MenuRefresh = .idle,
        actions: ProfileMenuActions = ProfileMenuActions()
    ) {
        self.summary = summary
        self.signingIn = signingIn
        self.actions = actions
        _showsLowLevel = State(initialValue: showsLowLevel)
        _refresh = State(initialValue: refresh)
    }

    /// Row 12: `Refreshing…` with the link disabled while the call is out, then the new stamp or the error.
    private func runRefresh() {
        guard refresh != .running else { return }
        refresh = .running
        Task {
            refresh = await actions.refresh() ? .idle : .failed
        }
    }

    /// The stamp line as it stands: the failure, the sync error, or the stamp (warning tone past a week).
    @ViewBuilder
    private var stampLine: some View {
        let running = refresh == .running
        let action = running ? Copy.menuRefreshing : nil
        if refresh == .failed {
            MenuStampLine(text: Copy.menuRefreshFailed, color: FSColor.wipe, action: Copy.commonTryAgain, run: runRefresh)
        } else if let error = summary.syncError {
            MenuStampLine(text: error, color: FSColor.wipe, action: action ?? Copy.commonTryAgain, busy: running, run: runRefresh)
        } else if let stamp = summary.stamp {
            MenuStampLine(
                text: stamp, color: summary.stale ? FSColor.warn : FSColor.text, action: action ?? Copy.menuRefresh,
                busy: running, run: runRefresh
            )
        }
    }

    var body: some View {
        ScrollView {
            VStack(alignment: .leading, spacing: 0) {
                Group {
                    if let you = summary.you { youBlock(you) } else { signedOut }
                }
                .padding(.horizontal, Self.side)
                .padding(.top, Self.statusBar + Self.youTop)
                .padding(.bottom, Self.youBottom)
                .frame(maxWidth: .infinity, alignment: .leading)
                .background { BandArt() }
                .overlay(alignment: .bottom) { Color.black.frame(height: FSStroke.hairline) }
                VStack(alignment: .leading, spacing: 0) {
                    if summary.you != nil {
                        switchSection
                        rowsSection(Copy.menuTools, summary.tools, lines: false)
                        rowsSection(Copy.menuAccount, accountRows, lines: true)
                    } else {
                        rowsSection(Copy.menuTools, summary.tools, lines: false)
                    }
                    footer
                }
                .padding(.horizontal, Self.side)
            }
            .parityContent()
        }
        .ignoresSafeArea(edges: .top)
        .background(FSColor.raised)
        .task { devices = await actions.loadDevices() }
        .task { notifications = await actions.loadNotifications() }
    }

    private func youBlock(_ you: AccountCharacter) -> some View {
        VStack(alignment: .leading, spacing: 0) {
            HStack(spacing: FSSpace.space4) {
                CharacterPortrait(you, size: Self.youPortrait)
                VStack(alignment: .leading, spacing: 0) {
                    Text(you.name)
                        .type(FSType.drawerName)
                        .foregroundStyle(ClassColor.of(you.className))
                        .accessibilityAddTraits(.isHeader)
                    VStack(alignment: .leading, spacing: FSSpace.space1) {
                        Text(summary.youLine).type(MenuType.line).foregroundStyle(FSColor.text)
                        if summary.isMain { MainTag() }
                    }
                    .frame(maxWidth: .infinity, alignment: .leading)
                    .parityRegion(.headMeta)
                }
                .frame(maxWidth: .infinity, alignment: .leading)
            }
            .parityRegion(.headTitle)
            // The last refresh failed: its line in wipe red with Try again, in place of the stamp.
            stampLine
        }
        .parityRegion(.head)
    }

    private var signedOut: some View {
        VStack(alignment: .leading, spacing: 0) {
            VStack(alignment: .leading, spacing: FSSpace.space1) {
                Text(Copy.appName)
                    .type(FSType.drawerName)
                    .foregroundStyle(FSColor.cream)
                    .accessibilityAddTraits(.isHeader)
                Text(Copy.signinBody).type(MenuType.line).foregroundStyle(FSColor.text)
                    .frame(maxWidth: .infinity, alignment: .leading)
                    .parityRegion(.headMeta)
            }
            .parityRegion(.headTitle)
            FSButton(Copy.signinButton, kind: .primary, action: actions.signIn)
                .disabled(signingIn)
                .padding(.top, FSSpace.space4)
                .padding(.bottom, FSSpace.space1)
        }
        .parityRegion(.head)
    }

    @ViewBuilder
    private var switchSection: some View {
        if !summary.others.isEmpty || !summary.lowLevel.isEmpty {
            let shown = summary.others + (showsLowLevel ? summary.lowLevel : [])
            MenuSection(title: Copy.menuSwitch) {
                ForEach(Array(shown.enumerated()), id: \.element.key) { index, character in
                    Button { actions.pick(character) } label: {
                        MenuRow(
                            title: character.name, titleColor: ClassColor.of(character.className), sub: ProfileMenuSummary.switchLine(character),
                            last: index == shown.count - 1 && summary.lowLevel.isEmpty
                        ) { CharacterPortrait(character, size: Self.rowPortrait, crestFallback: true) }
                    }
                    .buttonStyle(RowPressStyle())
                }
                if !summary.lowLevel.isEmpty {
                    Button(showsLowLevel ? Copy.menuHideLow : Copy.menuShowLow(n: summary.lowLevel.count)) {
                        showsLowLevel.toggle()
                    }
                    .type(FSType.bodyStrong)
                    .foregroundStyle(FSColor.gold)
                    .frame(minHeight: Self.rowMin)
                }
            }
        }
    }

    /// The Account rows, the Devices and Notifications lines from what was loaded once it arrives.
    private var accountRows: [ProfileMenuSummary.Row] {
        summary.account.map { row in
            switch row.destination {
            case .devices:
                guard let devices else { return row }
                let facts = ProfileMenuSummary.deviceFacts(devices, into: .init())
                let line = ProfileMenuSummary.devicesLine(facts.devices, newest: facts.newestDevice)
                return ProfileMenuSummary.Row(destination: row.destination, title: row.title, line: line)
            case .notifications:
                guard let notifications else { return row }
                let line = ProfileMenuSummary.notificationsLine(NotificationGroupsSummary(notifications).rows)
                return ProfileMenuSummary.Row(destination: row.destination, title: row.title, line: line)
            default:
                return row
            }
        }
    }

    /// Tools are labels alone (owner, 2026-11-09); Account rows say what they hold.
    private func rowsSection(_ title: String, _ rows: [ProfileMenuSummary.Row], lines: Bool) -> some View {
        MenuSection(title: title) {
            ForEach(Array(rows.enumerated()), id: \.element.id) { index, row in
                Button { actions.open(row.destination) } label: {
                    MenuRow(title: row.title, sub: lines && !row.line.isEmpty ? row.line : nil, last: index == rows.count - 1) { EmptyView() }
                }
                .buttonStyle(RowPressStyle())
            }
        }
    }

    private var footer: some View {
        VStack(alignment: .leading, spacing: FSSpace.space1) {
            // The site footer's fan line and operator line (web/src/lib/footer-copy.ts), word for word.
            SmallText(Copy.menuLegalFan)
            SmallText(summary.operatorLine)
            SmallText(summary.version).padding(.top, FSSpace.space2)
            HStack(spacing: 0) {
                Button(Copy.menuHelp, action: actions.help)
                SmallText(" · ")
                Button(Copy.menuPrivacy, action: actions.privacy)
            }
            .font(.custom(FSFont.bodySemibold, size: FSType.small.size, relativeTo: .footnote))
            .foregroundStyle(FSColor.gold)
        }
        .padding(.vertical, FSSpace.space6)
    }
}

/// The stamp (or the sync error) with its action: on one line, wrapping by words; at the
/// accessibility sizes the action takes its own line, so nothing breaks mid-word or strands.
private struct MenuStampLine: View {
    let text: String
    let color: Color
    let action: String
    /// The refresh is out: the action reads `Refreshing…` and does nothing.
    var busy = false
    let run: () -> Void
    @Environment(\.dynamicTypeSize) private var typeSize

    static var type: TypeRole { RowType.detail }

    var body: some View {
        Group {
            if typeSize.isAccessibilitySize {
                VStack(alignment: .leading, spacing: FSSpace.space1) {
                    Text(text).type(Self.type).foregroundStyle(color).fixedSize(horizontal: false, vertical: true)
                    button
                }
            } else {
                HStack(alignment: .firstTextBaseline, spacing: 0) {
                    Text(text + " · ").type(Self.type).foregroundStyle(color).fixedSize(horizontal: false, vertical: true)
                    button.fixedSize()
                }
                .frame(minHeight: ProfileMenuView.rowMin)
            }
        }
        .frame(maxWidth: .infinity, alignment: .leading)
        .padding(.top, FSSpace.space1)
    }

    /// Round 4, row 10: the press and its 44 pt target are the link's alone, never the whole line.
    private var button: some View {
        Button(action: run) {
            Text(action)
                .type(FSType.tag)
                .foregroundStyle(busy ? FSColor.muted : FSColor.gold)
                .frame(minHeight: FSTarget.min)
                .contentShape(.rect)
        }
        .buttonStyle(.plain)
        .disabled(busy)
    }
}

/// `Main`: the gold outline tag under the line.
private struct MainTag: View {
    var body: some View {
        Text(Copy.menuMain)
            .type(FSType.tag)
            .foregroundStyle(FSColor.gold)
            .padding(.horizontal, FSSpace.space2)
            .padding(.vertical, 1)
            .overlay(RoundedRectangle(cornerRadius: FSRadius.control).strokeBorder(FSColor.gold, lineWidth: FSStroke.hairline))
            .fixedSize()
    }
}

/// A menu section: the cream display title at 16 (14 above, 6 under), then its rows.
private struct MenuSection<Content: View>: View {
    let title: String
    @ViewBuilder let content: Content

    var body: some View {
        VStack(alignment: .leading, spacing: 0) {
            Text(title)
                .type(FSType.drawerSection)
                .foregroundStyle(FSColor.cream)
                .accessibilityAddTraits(.isHeader)
                .padding(.top, 14)
                .padding(.bottom, 6)
            content
        }
        .parityRegion(.section(title))
    }
}

/// A menu row: 48 tall at least, 7 above and below, the title 15 bold, its line 13 muted, a rule under
/// it unless it is the last.
private struct MenuRow<Lead: View>: View {
    let title: String
    var titleColor: Color = FSColor.strong
    let sub: String?
    let last: Bool
    @ViewBuilder let lead: Lead

    var body: some View {
        HStack(spacing: FSSpace.space3) {
            lead
            VStack(alignment: .leading, spacing: 0) {
                Text(title).type(MenuType.title).foregroundStyle(titleColor)
                if let sub { Text(sub).type(MenuType.sub).foregroundStyle(FSColor.muted) }
            }
            .frame(maxWidth: .infinity, alignment: .leading)
        }
        .padding(.vertical, ProfileMenuView.rowPad)
        .frame(minHeight: sub == nil ? FSTarget.min : ProfileMenuView.rowMin)
        .overlay(alignment: .bottom) { if !last { Hairline(FSColor.line) } }
        .contentShape(.rect)
        .accessibilityElement(children: .combine)
        .parityItem(.row)
    }
}

/// The drawer's line boxes (MENU_R5_CSS): the item and sub faces at the mock's line heights, a row's
/// title 15/20.57 and line 13/18.29, You's line 13/20.57 (the drawer's parity regions are built on them).
private enum MenuType {
    static let titleLine: CGFloat = 20.57
    static let subLine: CGFloat = 18.29
    static var title: TypeRole { TypeRole(face: FSType.item.face, size: FSType.item.size, lineHeight: titleLine, scalesWith: FSType.item.scalesWith) }
    static var sub: TypeRole { TypeRole(face: FSType.small.face, size: FSType.small.size, lineHeight: subLine, scalesWith: FSType.small.scalesWith) }
    static var line: TypeRole { TypeRole(face: FSType.small.face, size: FSType.small.size, lineHeight: titleLine, scalesWith: FSType.small.scalesWith) }
}

/// The drawer's Refresh (row 12): at rest, out, or failed.
enum MenuRefresh: Equatable {
    case idle, running, failed
}
