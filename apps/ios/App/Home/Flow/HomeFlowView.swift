import DesignSystem
import Models
import SwiftUI

/// What Home's blocks hand back.
struct HomeFlowActions {
    var raid = NextRaidActions()
    var openRaidNight: () -> Void = {}
    var openItem: (HomeUpgradeItem) -> Void = { _ in }
    var openGear: () -> Void = {}
    var openLoot: (LastNightLoot) -> Void = { _ in }
    var openReport: (String) -> Void = { _ in }
    var simAgain: () -> Void = {}
    var openActivity: (ActivityRow) -> Void = { _ in }
    var allActivity: () -> Void = {}
    var findGuild: () -> Void = {}
    var retry: () -> Void = {}
}

/// The time-aware Home (flows spec 4.1) in the round-5 look (mocks ios-home-*, m-ios-home): the
/// camp-road hero behind the bar and the head, then the server's blocks in its order as panels on the
/// ground, so the moment's 3-second answer follows the head directly.
struct HomeFlowView: View {
    let content: HomeFlowContent
    let build: String
    var actions = HomeFlowActions()

    /// The raid card's height first, then four rows (the mock's loading state): nothing shifts when they fill.
    static let skeletons: [CGFloat] = [365, 56, 56, 56, 56]

    var body: some View {
        ScrollView {
            BandPage(art: .hero) {
                HomeHeroHead(title: content.head.name, facts: content.head.facts, className: content.head.className)
            } notice: {
                if let stamp = content.offlineStamp { OfflineStrip(text: Copy.offlineStrip(time: stamp)) }
                if let stale = content.staleStamp { StaleStrip(time: stale, retry: actions.retry) }
            } content: {
                switch content.phase {
                case .loading: loading
                case .failed: failed
                case let .loaded(blocks):
                    ForEach(blocks) { block in self.block(block) }
                }
            }
        }
        .environment(\.panelRowTrailCell, true)
        .environment(\.answersOffline, content.offlineStamp != nil)
        .screenGround()
    }

    private var loading: some View {
        VStack(spacing: FSSpace.space3) {
            ForEach(Array(Self.skeletons.enumerated()), id: \.offset) { _, height in FSSkeleton(height: height) }
        }
        .parityRegion(ParityRegionName("section-loading"))
        .accessibilityElement(children: .ignore)
        .accessibilityLabel(Copy.homeAnswerLoading)
    }

    private var failed: some View {
        PanelBlock {
            PanelRow(Copy.homeErrorTitle, sub: Text(Copy.homeErrorLine), last: true) {
                EmptyView()
            } trail: {
                Button(Copy.commonTryAgain, action: actions.retry).buttonStyle(FSButtonStyle(fullWidth: false))
            }
        }
        .parityRegion(ParityRegionName("section-error"))
    }

    @ViewBuilder
    private func block(_ block: HomeBlockContent) -> some View {
        switch block {
        case let .upgrades(upgrades):
            HomeUpgradesSection(content: upgrades, build: build, open: actions.openItem, openGear: actions.openGear)
        case let .raidCard(raid):
            NextRaidCardView(content: raid, actions: actions.raid)
        case let .raidLine(raid):
            NextRaidLineView(content: raid, build: build, actions: actions.raid)
        case let .raidLive(live):
            RaidLiveCard(content: live, open: actions.openRaidNight)
        case let .lastNight(night):
            LastNightSection(content: night, build: build, actions: actions)
        case let .newForYou(new):
            NewForYouSection(content: new, build: build, open: actions.openActivity, all: actions.allActivity)
        case let .findGuild(line):
            FindGuildBlock(line: line, find: actions.findGuild)
        case let .notice(text):
            PanelBlock { PanelRow(text, last: true) }.parityRegion(ParityRegionName("section-notice"))
        }
    }
}

/// One head on every Home (flows spec 4.1) on the hero: the class crest, the name in the class
/// colour in the display face, the fact line (`Level 60 Fury Warrior · Olympus XXVII`; a leveler's
/// levels to the next band). Without a class (no character yet) the title stands alone in cream.
struct HomeHeroHead: View {
    let title: String
    let facts: [String]
    let className: String?

    static let top: CGFloat = 6
    static let bottom: CGFloat = 18
    static let factSeparator = " · "

    var body: some View {
        HStack(spacing: FSSpace.space3) {
            if let className {
                ClassCrest(className: className, size: FSSize.crest56, flat: true)
            }
            VStack(alignment: .leading, spacing: 0) {
                Text(title)
                    .type(TypeRole(face: FSFont.title, size: 30, lineHeight: 34, scalesWith: .title))
                    .foregroundStyle(className.map { ClassColor.of($0) } ?? FSColor.cream)
                    .shadow(color: .black.opacity(PanelShadow.opacity), radius: PanelShadow.radius, y: PanelShadow.drop)
                    .accessibilityAddTraits(.isHeader)
                if !facts.isEmpty {
                    Text(facts.joined(separator: HomeHeroHead.factSeparator))
                        .type(TypeRole(face: FSFont.body, size: 14, lineHeight: 20, scalesWith: .subheadline))
                        .foregroundStyle(FSColor.text)
                        .frame(maxWidth: .infinity, alignment: .leading)
                        .parityRegion(.headMeta)
                }
            }
        }
        .frame(maxWidth: .infinity, alignment: .leading)
        .parityRegion(.headTitle)
        .parityRegion(.head)
        .padding(.horizontal, FSSpace.space4)
        .padding(.top, Self.top)
        .padding(.bottom, Self.bottom)
    }
}

/// Next upgrades (F5) on the plain-framed panel: three items with their gain and source, the Best run line.
struct HomeUpgradesSection: View {
    let content: HomeUpgradesContent
    let build: String
    let open: (HomeUpgradeItem) -> Void
    let openGear: () -> Void

    var body: some View {
        PanelSection(Copy.homeUpgradesNext, action: (Copy.homeUpgradesGear, openGear), frame: .plain) {
            ForEach(Array(content.rows.enumerated()), id: \.element.id) { index, row in
                if index > 0 { Hairline(FSColor.line) }
                Button { open(row) } label: {
                    ItemLineView(
                        name: row.look.name ?? Copy.itemsNotInData, quality: row.look.quality, icon: row.look.icon,
                        source: row.source, gain: row.gain.map { ($0.figure, $0.unit) }, build: build, inset: 0
                    )
                }
                .buttonStyle(RowPressStyle())
            }
            ForEach([content.note, content.bestRun].compactMap { $0 }, id: \.self) { line in
                Hairline(FSColor.line)
                Button(action: openGear) { PanelRow(line, slim: true, last: true) }.buttonStyle(RowPressStyle())
            }
        }
    }
}

/// Raid live (F2, mock ios-home-raid-live) on the lead panel: the green dot, the guild and zone, the
/// pull and your answer, Open raid night.
struct RaidLiveCard: View {
    let content: RaidLiveContent
    /// Open raid night; nil where the card is raid night already (the Raid tab).
    let open: (() -> Void)?

    static let liveDot: CGFloat = 8
    /// The link's line (`.link`, 15/24).
    static let linkLine: CGFloat = 24

    var body: some View {
        LeadPanel {
            HStack(spacing: FSSpace.space2) {
                Circle().fill(FSColor.kill).frame(width: Self.liveDot, height: Self.liveDot)
                    .accessibilityHidden(true)
                Text(content.eyebrow).type(FSType.bodyStrong).foregroundStyle(FSColor.strong)
            }
            Text(content.line).type(FSType.small).foregroundStyle(FSColor.text).padding(.top, FSSpace.space2)
            if let open {
                Button(Copy.homeRaidOpenNight, action: open)
                    .font(.custom(FSFont.bodySemibold, size: 15, relativeTo: .body))
                    .foregroundStyle(FSColor.gold)
                    .frame(minHeight: FSTarget.min)
                    .contentShape(.rect)
                    .padding(.vertical, -(FSTarget.min - Self.linkLine) / 2)
                    .padding(.top, FSSpace.space2)
            }
        }
        .parityRegion(.countdowns)
    }
}

/// Last night (F2, mock ios-home-after-raid) on the plain-framed panel: your loot with its gain, the
/// report, Sim again.
struct LastNightSection: View {
    let content: LastNightContent
    let build: String
    let actions: HomeFlowActions

    var body: some View {
        PanelSection(Copy.homeLastNight, frame: .plain) {
            ForEach(content.loot) { loot in
                Button { actions.openLoot(loot) } label: {
                    PanelRow(loot.look.name ?? "", titleColor: ItemQuality(value: loot.look.quality).text, sub: Text(loot.line)) {
                        ItemIconView(icon: loot.look.icon, quality: loot.look.quality, build: build, size: FSSize.icon40).parityItem(.image)
                    }
                }
                .buttonStyle(RowPressStyle())
                .accessibilityHint(Copy.itemsA11yHint)
            }
            if let report = content.report, let reportID = content.reportID {
                Button { actions.openReport(reportID) } label: {
                    PanelRow(report.title, sub: report.sub.isEmpty ? nil : Text(report.sub), last: content.simAgain == nil) {
                        ItemIconView(icon: content.reportIcon, quality: nil, build: build, size: FSSize.icon40).parityItem(.image)
                    }
                }
                .buttonStyle(RowPressStyle())
            }
            if let sim = content.simAgain {
                Button(action: actions.simAgain) { PanelRow(sim.title, sub: Text(sim.sub), slim: true, last: true) }
                    .buttonStyle(RowPressStyle())
            }
        }
    }
}

/// New for you (F4, M7) on a plain panel: the three newest Activity rows, unread dotted, All activity (N).
struct NewForYouSection: View {
    let content: NewForYouContent
    let build: String
    let open: (ActivityRow) -> Void
    let all: () -> Void

    var body: some View {
        let link = content.unread > 0 ? Copy.homeNewForYouAllUnread(n: content.unread) : Copy.homeNewForYouAll
        PanelSection(Copy.homeNewForYou, action: (link, all)) {
            ForEach(Array(content.rows.enumerated()), id: \.element.id) { index, item in
                Button { open(item.row) } label: {
                    PanelRow(
                        item.row.title, sub: Text(item.line), unread: item.row.read ? nil : Copy.activityUnread,
                        last: index == content.rows.count - 1
                    ) {
                        ActivityLead(row: item.row, build: build)
                    }
                }
                .buttonStyle(RowPressStyle())
            }
        }
    }
}

/// Find a guild (no guild at 60) on a plain panel: the one line, with the first raids date before
/// they open, and the button.
struct FindGuildBlock: View {
    let line: String
    let find: () -> Void

    var body: some View {
        PanelBlock {
            Text(line).type(FSType.small).foregroundStyle(FSColor.text).padding(.top, FSSpace.space2)
            FSButton(Copy.homeFindGuildButton, action: find).padding(.vertical, FSSpace.space3)
        }
        .parityRegion(.section(Copy.homeFindGuildTitle))
    }
}

/// Signed in with no character yet (mock ios-home-new): the title on the hero, then the one thing to do.
struct HomeNoCharacterView: View {
    let checkAgain: () -> Void

    var body: some View {
        ScrollView {
            BandPage(art: .hero) {
                HomeHeroHead(title: Copy.homeNewTitle, facts: [Copy.homeNewLine], className: nil)
            } content: {
                FSButton(Copy.homeNewCheck, kind: .primary, action: checkAgain)
            }
        }
        .screenGround()
    }
}

/// A pre-round-5 section's action at the right (`Roster ›`): the raid pages not yet restyled.
struct SectionLink: View {
    let title: String
    let action: () -> Void

    var body: some View {
        Button(title, action: action)
            .font(.custom(FSFont.bodySemibold, size: FSType.small.size, relativeTo: .footnote))
            .foregroundStyle(FSColor.strong)
    }
}
