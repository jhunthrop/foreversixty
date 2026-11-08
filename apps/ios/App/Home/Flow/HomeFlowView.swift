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

/// The time-aware Home (flows spec 4.1; mocks ios-home-*): one head, then the server's blocks in
/// its order, so the moment's 3-second answer follows the head directly.
struct HomeFlowView: View {
    let content: HomeFlowContent
    let build: String
    var actions = HomeFlowActions()

    /// `.ap section { margin-top: 28px }`; `.blk { margin-top: 24px }`.
    static let sectionGap: CGFloat = 28
    static let blockGap: CGFloat = 24
    static let skeletons = 5
    static let skeletonHeight: CGFloat = 56

    var body: some View {
        ScrollView {
            VStack(alignment: .leading, spacing: 0) {
                if let stamp = content.offlineStamp { OfflineStrip(text: Copy.offlineStrip(time: stamp)) }
                if let stale = content.staleStamp { StaleStrip(time: stale, retry: actions.retry) }
                HomeHeadView(head: content.head)
                switch content.phase {
                case .loading: loading
                case .failed: failed
                case let .loaded(blocks):
                    ForEach(blocks) { block in self.block(block) }
                }
            }
            .padding(.top, FSSpace.space2)
            .padding(.horizontal, FSSpace.gutterPhone)
            .padding(.bottom, FSSpace.space6)
            .parityContent()
        }
        .background(FSColor.bg)
    }

    private var loading: some View {
        VStack(spacing: FSSpace.space3) {
            ForEach(0 ..< Self.skeletons, id: \.self) { _ in FSSkeleton(height: Self.skeletonHeight) }
        }
        .parityRegion(ParityRegionName("section-loading"))
        .padding(.top, Self.blockGap)
        .accessibilityElement(children: .ignore)
        .accessibilityLabel(Copy.homeAnswerLoading)
    }

    private var failed: some View {
        FSRow(title: Copy.homeErrorTitle, sub: Copy.homeErrorLine, inset: 0, trailPlacement: .below) {
            EmptyView()
        } trail: {
            Button(Copy.commonTryAgain, action: actions.retry).buttonStyle(FSButtonStyle(fullWidth: false, small: true))
        }
        .parityRegion(ParityRegionName("section-error"))
        .padding(.top, Self.blockGap)
    }

    @ViewBuilder
    private func block(_ block: HomeBlockContent) -> some View {
        switch block {
        case let .upgrades(upgrades):
            HomeUpgradesSection(content: upgrades, build: build, open: actions.openItem, openGear: actions.openGear)
                .padding(.top, Self.sectionGap)
        case let .raidCard(raid):
            NextRaidCardView(content: raid, actions: actions.raid).padding(.top, Self.blockGap)
        case let .raidLine(raid):
            NextRaidLineView(content: raid, build: build, actions: actions.raid).padding(.top, Self.blockGap)
        case let .raidLive(live):
            RaidLiveCard(content: live, open: actions.openRaidNight).padding(.top, Self.blockGap)
        case let .lastNight(night):
            LastNightSection(content: night, build: build, actions: actions).padding(.top, Self.sectionGap)
        case let .newForYou(new):
            NewForYouSection(content: new, build: build, open: actions.openActivity, all: actions.allActivity)
                .padding(.top, Self.sectionGap)
        case let .findGuild(line):
            FindGuildBlock(line: line, find: actions.findGuild).padding(.top, Self.blockGap)
        case let .notice(text):
            SmallText(text).padding(.top, Self.blockGap)
        }
    }
}

/// The head: the class crest, the name in the class colour, the fact line (`Level 60 Fury Warrior`,
/// `Olympus XXVII`; a leveler's levels to the next band).
struct HomeHeadView: View {
    let head: HomeHeadContent

    var body: some View {
        PageHeader(title: head.name, titleColor: ClassColor.of(head.className)) {
            ClassCrest(className: head.className, size: FSSize.crest56, flat: true)
        } meta: {
            WrapLayout(spacing: FSSpace.space4, lineSpacing: 6) {
                ForEach(head.facts, id: \.self) { SmallText($0) }
            }
        }
    }
}

/// Next upgrades (F5): three rows with their gain and source, and the Best run line.
struct HomeUpgradesSection: View {
    let content: HomeUpgradesContent
    let build: String
    let open: (HomeUpgradeItem) -> Void
    let openGear: () -> Void

    var body: some View {
        FSSection(Copy.homeUpgradesNext) {
            SectionLink(title: Copy.homeUpgradesGear, action: openGear)
        } content: {
            VStack(alignment: .leading, spacing: 0) {
                ForEach(content.rows) { row in
                    Button { open(row) } label: {
                        ItemLineView(
                            name: row.look.name ?? Copy.itemsNotInData, quality: row.look.quality, icon: row.look.icon,
                            source: row.source, gain: row.gain.map { ($0.figure, $0.unit) }, build: build
                        )
                    }
                    .buttonStyle(RowPressStyle())
                }
                if let note = content.note {
                    FSRow(title: note, density: .slim, inset: 0) { EmptyView() }
                }
                if let bestRun = content.bestRun {
                    FSRow(title: bestRun, density: .slim, inset: 0) { EmptyView() }
                }
            }
        }
    }
}

/// A section's action at the right (`Gear ›`, `All activity (3) ›`).
struct SectionLink: View {
    let title: String
    let action: () -> Void

    var body: some View {
        Button(title, action: action)
            .font(.custom(FSFont.bodySemibold, size: FSType.small.size, relativeTo: .footnote))
            .foregroundStyle(FSColor.strong)
    }
}

/// Raid live (F2, mock ios-home-raid-live): pulled how long ago, the counts, your answer, Open raid night.
struct RaidLiveCard: View {
    let content: RaidLiveContent
    let open: () -> Void

    var body: some View {
        FSCountdown(eyebrow: content.eyebrow, live: true) {
            VStack(alignment: .leading, spacing: FSSpace.space2) {
                SmallText(content.line)
                Button(Copy.homeRaidOpenNight, action: open)
                    .font(.custom(FSFont.bodySemibold, size: 15, relativeTo: .body))
                    .foregroundStyle(FSColor.gold)
            }
        }
        .parityRegion(.countdowns)
    }
}

/// Last night (F2, mock ios-home-after-raid): your loot with its gain, the report, Sim again.
struct LastNightSection: View {
    let content: LastNightContent
    let build: String
    let actions: HomeFlowActions

    var body: some View {
        FSSection(Copy.homeLastNight) {
            VStack(spacing: 0) {
                ForEach(content.loot) { loot in
                    Button { actions.openLoot(loot) } label: {
                        FSRow(title: loot.look.name ?? "", titleColor: ItemQuality(value: loot.look.quality).text, sub: loot.line, inset: 0) {
                            ItemIconView(icon: loot.look.icon, quality: loot.look.quality, build: build, size: FSSize.icon40)
                                .parityItem(.image)
                        }
                    }
                    .buttonStyle(RowPressStyle())
                    .accessibilityHint(Copy.itemsA11yHint)
                }
                if let report = content.report, let reportID = content.reportID {
                    Button { actions.openReport(reportID) } label: {
                        FSRow(title: report.title, sub: report.sub.isEmpty ? nil : report.sub, inset: 0) {
                            ItemIconView(icon: content.reportIcon, quality: nil, build: build, size: FSSize.icon40).parityItem(.image)
                        }
                    }
                    .buttonStyle(RowPressStyle())
                }
                if let sim = content.simAgain {
                    Button(action: actions.simAgain) {
                        FSRow(title: sim.title, sub: sim.sub, density: .slim, inset: 0) { EmptyView() }
                    }
                    .buttonStyle(RowPressStyle())
                }
            }
        }
    }
}

/// New for you (F4, M7): the three newest Activity rows, unread dotted, All activity (N).
struct NewForYouSection: View {
    /// The descent under an inline 40 pt icon in a web row.
    static let iconDescent: CGFloat = 3

    let content: NewForYouContent
    let build: String
    let open: (ActivityRow) -> Void
    let all: () -> Void

    var body: some View {
        FSSection(Copy.homeNewForYou) {
            SectionLink(title: content.unread > 0 ? Copy.homeNewForYouAllUnread(n: content.unread) : Copy.homeNewForYouAll, action: all)
        } content: {
            VStack(spacing: 0) {
                ForEach(content.rows) { item in
                    Button { open(item.row) } label: {
                        FSRow(title: item.row.title, sub: item.line, inset: 0) {
                            ItemIconView(icon: item.row.icon, quality: nil, build: build, size: FSSize.icon40)
                                .parityItem(.image)
                                // The web's inline icon sits on the text baseline and takes its line's
                                // descent below it, so a two-line row is the icon plus that gap.
                                .padding(.bottom, Self.iconDescent)
                        }
                        .overlay(alignment: .topLeading) {
                            if !item.row.read {
                                Circle().fill(FSColor.gold)
                                    .frame(width: ActivityView.dotSize, height: ActivityView.dotSize)
                                    .offset(x: -ActivityView.dotInset, y: 24)
                                    .accessibilityHidden(true)
                            }
                        }
                        .accessibilityValue(item.row.read ? "" : Copy.activityUnread)
                    }
                    .buttonStyle(RowPressStyle())
                }
            }
        }
    }
}

/// Find a guild (no guild at 60): the one line, with the first raids date before they open.
struct FindGuildBlock: View {
    let line: String
    let find: () -> Void

    var body: some View {
        VStack(alignment: .leading, spacing: 0) {
            SmallText(line)
            FSButton(Copy.homeFindGuildButton, action: find).padding(.top, FSSpace.space5)
        }
        .parityRegion(.section(Copy.homeFindGuildTitle))
    }
}

/// Signed in with no character yet (mock ios-home-new): the one thing to do.
struct HomeNoCharacterView: View {
    let checkAgain: () -> Void

    var body: some View {
        ScrollView {
            VStack(alignment: .leading, spacing: 0) {
                PageHeader(title: Copy.homeNewTitle) { SmallText(Copy.homeNewLine) }
                FSButton(Copy.homeNewCheck, kind: .primary, action: checkAgain).padding(.top, FSSpace.space5)
            }
            .padding(.top, FSSpace.space2)
            .padding(.horizontal, FSSpace.gutterPhone)
            .parityContent()
        }
        .background(FSColor.bg)
    }
}
