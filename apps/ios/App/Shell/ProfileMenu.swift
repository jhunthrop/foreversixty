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
    /// `Updated 2 h ago from Battle.net`, nil before the first import.
    let stamp: String?
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
        var notificationGroups: [String]?
        var plan: String?
        var devices: Int?
        var offlineBytes: Int64?
    }

    init(me: Me?, current: AccountCharacter?, facts: AccountFacts, version: String, now: Date) {
        self.version = version
        operatorLine = Copy.menuLegalOperator(year: String(Calendar(identifier: .gregorian).component(.year, from: now)))
        guard let me, let current else {
            you = nil
            isMain = false
            stamp = nil
            others = []
            lowLevel = []
            tools = MenuDestination.signedOutTools.map(Self.tool)
            account = []
            return
        }
        you = current
        isMain = current.key == me.mainCharacterKey
        stamp = me.bnetImportedAt.map { Copy.sharedCharactersUpdatedFrom(ago: RelativeAge(from: $0, to: now).phrase) }
        let rest = me.characters.filter { $0.key != current.key }
        others = rest.filter { ($0.level ?? 0) >= Self.switchLevel }
        lowLevel = rest.filter { ($0.level ?? 0) < Self.switchLevel }
        tools = MenuDestination.tools.map(Self.tool)
        account = MenuDestination.account.map { Self.account($0, facts: facts) }
    }

    private static func tool(_ destination: MenuDestination) -> Row {
        switch destination {
        case .guides: Row(destination: destination, title: Copy.menuToolGuides, line: Copy.menuToolGuidesLine)
        case .tiers: Row(destination: destination, title: Copy.menuToolTiers, line: Copy.menuToolTiersLine)
        case .patchNotes: Row(destination: destination, title: Copy.menuToolPatchNotes, line: Copy.menuToolPatchNotesLine)
        case .sims: Row(destination: destination, title: Copy.menuToolSims, line: Copy.menuToolSimsLine)
        case .reports: Row(destination: destination, title: Copy.menuToolReports, line: Copy.menuToolReportsLine)
        case .planner: Row(destination: destination, title: Copy.menuToolPlanner, line: Copy.menuToolPlannerLine)
        case .rankings: Row(destination: destination, title: Copy.menuToolRankings, line: Copy.menuToolRankingsLine)
        default: Row(destination: destination, title: "", line: "")
        }
    }

    private static func account(_ destination: MenuDestination, facts: AccountFacts) -> Row {
        switch destination {
        case .notifications:
            Row(destination: destination, title: Copy.menuNotifications, line: notificationsLine(facts.notificationGroups))
        case .plan:
            Row(destination: destination, title: Copy.menuPlan, line: facts.plan ?? "")
        case .devices:
            Row(destination: destination, title: Copy.menuDevices, line: facts.devices.map { Copy.menuDevicesLine(n: $0) } ?? "")
        case .offlineData:
            Row(destination: destination, title: Copy.menuOffline, line: facts.offlineBytes.map(offlineLine) ?? "")
        default:
            Row(destination: destination, title: Copy.menuSettings, line: Copy.menuSettingsLine)
        }
    }

    /// `Raids, gear and sims, logs on`: the groups that are on, in their order, sentence case.
    static func notificationsLine(_ groups: [String]?) -> String {
        guard let groups else { return "" }
        guard let first = groups.first else { return Copy.menuNotificationsOff }
        let joined = ([first] + groups.dropFirst().map { $0.lowercased() }).joined(separator: ", ")
        return Copy.menuNotificationsOn(groups: joined)
    }

    static func offlineLine(_ bytes: Int64) -> String {
        Copy.menuOfflineLine(size: ByteCountFormatter.string(fromByteCount: bytes, countStyle: .file))
    }
}

/// The menu's rows and links, handed in by the shell.
struct ProfileMenuActions {
    var pick: (AccountCharacter) -> Void = { _ in }
    var refresh: () -> Void = {}
    var open: (MenuDestination) -> Void = { _ in }
    var signIn: () -> Void = {}
    var help: () -> Void = {}
    var privacy: () -> Void = {}
}

/// The profile menu's content (mocks ios-menu, -one, -low, -out): You, Switch character, Tools,
/// Account and the footer; signed out, Sign in and the tools that work without an account.
struct ProfileMenuView: View {
    let summary: ProfileMenuSummary
    var signingIn = false
    var actions = ProfileMenuActions()
    @State private var showsLowLevel = false

    static let youPortrait: CGFloat = 56
    static let rowPortrait: CGFloat = 40
    /// `.drawer { padding-top: calc(var(--status) + 12px) }`: the mock's drawer starts at the
    /// content's top and pads by the status bar again, so the head sits 66 under the safe area.
    static let topInset: CGFloat = 66
    /// `.drawer section { margin-top: 24px }`.
    static let sectionGap: CGFloat = 24

    init(summary: ProfileMenuSummary, signingIn: Bool = false, showsLowLevel: Bool = false, actions: ProfileMenuActions = ProfileMenuActions()) {
        self.summary = summary
        self.signingIn = signingIn
        self.actions = actions
        _showsLowLevel = State(initialValue: showsLowLevel)
    }

    var body: some View {
        ScrollView {
            VStack(alignment: .leading, spacing: 0) {
                if let you = summary.you {
                    youBlock(you)
                    switchSection
                    rowsSection(Copy.menuTools, summary.tools)
                    rowsSection(Copy.menuAccount, summary.account)
                } else {
                    signedOut
                    rowsSection(Copy.menuTools, summary.tools)
                }
                footer
            }
            .padding(.top, Self.topInset)
            .padding(.horizontal, FSSpace.gutterPhone)
            .padding(.bottom, FSSpace.space6)
            .parityContent()
        }
        .background(FSColor.raised)
    }

    private func youBlock(_ you: AccountCharacter) -> some View {
        VStack(alignment: .leading, spacing: 0) {
            PageHeader(title: you.name, titleColor: ClassColor.of(you.className)) {
                CharacterPortrait(you, size: Self.youPortrait)
            } meta: {
                WrapLayout(spacing: 6) {
                    SmallText(AccountCopy.characterLine(you))
                    if summary.isMain { FSTag(Copy.menuCharacterMain, tone: .gold) }
                }
            }
            .padding(.horizontal, FSSpace.gutterPhone)
            if let stamp = summary.stamp {
                HStack(spacing: 0) {
                    SmallText(stamp + " · ")
                    Button(Copy.menuRefresh, action: actions.refresh)
                        .font(.custom(FSFont.bodySemibold, size: FSType.small.size, relativeTo: .footnote))
                        .foregroundStyle(FSColor.gold)
                }
                .padding(.top, 6)
            }
        }
    }

    @ViewBuilder
    private var switchSection: some View {
        if !summary.others.isEmpty || !summary.lowLevel.isEmpty {
            FSSection(Copy.menuSwitch) {
                VStack(alignment: .leading, spacing: 0) {
                    ForEach(summary.others + (showsLowLevel ? summary.lowLevel : []), id: \.key) { character in
                        Button { actions.pick(character) } label: {
                            FSRow(
                                title: character.name, titleColor: ClassColor.of(character.className),
                                sub: AccountCopy.characterLine(character), inset: 0
                            ) { CharacterPortrait(character, size: Self.rowPortrait) }
                        }
                        .buttonStyle(RowPressStyle())
                    }
                    if !summary.lowLevel.isEmpty {
                        Button(showsLowLevel ? Copy.menuLowLevelHide : Copy.menuLowLevelShow(n: summary.lowLevel.count)) {
                            showsLowLevel.toggle()
                        }
                        .font(.custom(FSFont.bodySemibold, size: FSType.item.size, relativeTo: .body))
                        .foregroundStyle(FSColor.gold)
                        .frame(minHeight: FSTarget.min)
                    }
                }
                .padding(.top, FSSpace.space1)
            }
            .padding(.top, Self.sectionGap)
        }
    }

    private func rowsSection(_ title: String, _ rows: [ProfileMenuSummary.Row]) -> some View {
        FSSection(title) {
            VStack(spacing: 0) {
                ForEach(rows) { row in
                    Button { actions.open(row.destination) } label: {
                        FSRow(title: row.title, sub: row.line.isEmpty ? nil : row.line, inset: 0) { EmptyView() }
                    }
                    .buttonStyle(RowPressStyle())
                }
            }
            .padding(.top, FSSpace.space1)
        }
        .padding(.top, Self.sectionGap)
    }

    private var signedOut: some View {
        VStack(alignment: .leading, spacing: 0) {
            PageHeader(title: Copy.appName) { SmallText(Copy.signinBody) }
                .padding(.horizontal, FSSpace.gutterPhone)
            FSButton(Copy.signinButton, kind: .primary, action: actions.signIn)
                .disabled(signingIn)
                .padding(.top, FSSpace.space5)
        }
    }

    private var footer: some View {
        VStack(alignment: .leading, spacing: 6) {
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
        .padding(.top, FSSpace.space7)
    }
}
