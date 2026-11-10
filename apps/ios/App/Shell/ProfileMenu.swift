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

    static let youPortrait: CGFloat = 56
    static let rowPortrait: CGFloat = 40
    /// The band runs under the status bar; You sits 18 under it (the mock's `--status`, 54).
    static let statusBar: CGFloat = 54
    static let side: CGFloat = 18
    static let youTop: CGFloat = 18
    static let youBottom: CGFloat = 12
    static let rowMin: CGFloat = 48
    static let rowPad: CGFloat = 7

    init(summary: ProfileMenuSummary, signingIn: Bool = false, showsLowLevel: Bool = false, actions: ProfileMenuActions = ProfileMenuActions()) {
        self.summary = summary
        self.signingIn = signingIn
        self.actions = actions
        _showsLowLevel = State(initialValue: showsLowLevel)
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
                        rowsSection(Copy.menuAccount, summary.account, lines: true)
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
    }

    private func youBlock(_ you: AccountCharacter) -> some View {
        VStack(alignment: .leading, spacing: 0) {
            HStack(spacing: FSSpace.space4) {
                CharacterPortrait(you, size: Self.youPortrait)
                VStack(alignment: .leading, spacing: 0) {
                    Text(you.name)
                        .type(TypeRole(face: FSFont.title, size: 26, lineHeight: 32, scalesWith: .title2))
                        .foregroundStyle(ClassColor.of(you.className))
                        .accessibilityAddTraits(.isHeader)
                    VStack(alignment: .leading, spacing: FSSpace.space1) {
                        Text(AccountCopy.characterLine(you)).type(FSType.small).foregroundStyle(FSColor.text)
                        if summary.isMain { MainTag() }
                    }
                    .parityRegion(.headMeta)
                }
                .frame(maxWidth: .infinity, alignment: .leading)
            }
            .parityRegion(.headTitle)
            .parityRegion(.head)
            if let stamp = summary.stamp {
                HStack(spacing: 0) {
                    Text(stamp + " · ").type(TypeRole(face: FSFont.body, size: 12, lineHeight: 16, scalesWith: .caption))
                    Button(Copy.menuRefresh, action: actions.refresh)
                        .font(.custom(FSFont.bodySemibold, size: 12, relativeTo: .caption))
                        .foregroundStyle(FSColor.gold)
                }
                .foregroundStyle(FSColor.text)
                .frame(minHeight: FSTarget.min + FSSpace.space1)
                .padding(.top, FSSpace.space1)
            }
        }
    }

    private var signedOut: some View {
        VStack(alignment: .leading, spacing: FSSpace.space1) {
            Text(Copy.appName)
                .type(TypeRole(face: FSFont.title, size: 26, lineHeight: 32, scalesWith: .title2))
                .foregroundStyle(FSColor.cream)
                .accessibilityAddTraits(.isHeader)
            Text(Copy.signinBody).type(FSType.small).foregroundStyle(FSColor.text).parityRegion(.headMeta)
            FSButton(Copy.signinButton, kind: .primary, action: actions.signIn)
                .disabled(signingIn)
                .padding(.top, FSSpace.space3)
                .padding(.bottom, FSSpace.space1)
        }
        .parityRegion(.headTitle)
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
                            title: character.name, titleColor: ClassColor.of(character.className), sub: AccountCopy.characterLine(character),
                            last: index == shown.count - 1 && summary.lowLevel.isEmpty
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
                    .frame(minHeight: Self.rowMin)
                }
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

/// `Main`: the gold outline tag under the line.
private struct MainTag: View {
    var body: some View {
        Text(Copy.menuCharacterMain)
            .font(.custom(FSFont.bodySemibold, size: 12, relativeTo: .caption))
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
                .type(TypeRole(face: FSFont.title, size: 16, lineHeight: 22, scalesWith: .headline))
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
                Text(title).type(FSType.item).foregroundStyle(titleColor)
                if let sub { Text(sub).type(RowType.sub).foregroundStyle(FSColor.muted) }
            }
            .frame(maxWidth: .infinity, alignment: .leading)
        }
        .padding(.vertical, ProfileMenuView.rowPad)
        .frame(minHeight: ProfileMenuView.rowMin)
        .overlay(alignment: .bottom) { if !last { Hairline(FSColor.line) } }
        .contentShape(.rect)
        .accessibilityElement(children: .combine)
        .parityItem(.row)
    }
}
