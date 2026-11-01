import DesignSystem
import Models
import Routing
import SwiftUI

/// What the character page hands back to its screen.
struct CharacterActions {
    var retry: () -> Void = {}
    var openHow: () -> Void = {}
    var openSocket: (PaperdollSlot) -> Void = { _ in }
    /// Nil while the reports switch is off: recent fights do not open.
    var openReport: ((FightLine) -> Void)?
    var openHistory: () -> Void = {}
    var searchAgain: () -> Void = {}
    /// Lane F15 (F11.4 follow-up): `Open in planner` under Builds seen; nil while `planner` is off.
    var openPlanner: ((DeepLink) -> Void)?
}

/// F11.4 to F11.7 as drawn from the store: stateless, so snapshots drive every state. The bar
/// (back, the Watch star, share) belongs to the screen.
struct CharacterView: View {
    let ref: CharacterRef
    let phase: CharacterPhase
    let items: [Int: Item]
    let account: RankingsAccount
    let specs: SpecCatalog
    let build: String
    /// The star's failed toggle: `char.watchFailed` under the head.
    var watchFailed = false
    let now: Date
    let clock: ReportsClock
    var actions = CharacterActions()

    var body: some View {
        VStack(alignment: .leading, spacing: 0) {
            switch phase {
            case .loading:
                CharacterSkeleton()
            case .failed:
                ReportsErrorCard(text: Copy.charError, retry: actions.retry).padding(16)
            case .notFound:
                notFound.padding(16)
            case let .loaded(content):
                if let since = content.offlineSince {
                    OfflineStrip(text: Copy.offlineStrip(time: OfflineStamp.format(since, now: now)))
                }
                if let since = content.staleSince {
                    StaleStrip(time: OfflineStamp.format(since, now: now), retry: actions.retry)
                }
                loaded(content.profile)
            }
        }
        .frame(maxWidth: .infinity, alignment: .leading)
        .padding(.bottom, SimsListView.tabBarClearance)
    }

    private func loaded(_ profile: CharacterProfile) -> some View {
        VStack(alignment: .leading, spacing: 22) {
            CharacterHead(
                profile: profile, specs: specs,
                yours: CharacterPage.isYours(profile.character, keys: Set(account.characterKeys))
            )
            if watchFailed {
                Text(Copy.charWatchFailed(name: profile.character.name)).font(ScaledType.meta).foregroundStyle(FSColor.wipe)
                    .fixedSize(horizontal: false, vertical: true)
            }
            if CharacterPage.hasScore(profile) {
                ScoreCard(profile: profile, openHow: actions.openHow)
            } else {
                ReportStateCard(title: Copy.charNoLogs, detail: Copy.charNoScore(name: profile.character.name), color: FSColor.strong)
            }
            let bests = CharacterPage.bests(profile, now: now)
            if !bests.isEmpty { section(Copy.charBests) { fights(bests, recent: false) } }
            let recent = CharacterPage.recent(profile, now: now)
            if !recent.isEmpty { section(Copy.charRecent) { fights(recent, recent: true) } }
            section(Copy.charGear) { gear(profile) }
            let builds = CharacterPage.builds(profile, clock: clock)
            if !builds.isEmpty { section(Copy.charBuilds) { buildRows(builds) } }
            if let openPlanner = actions.openPlanner, let link = CharacterPage.plannerLink(profile) {
                // In the app: a chevron, not the external arrow.
                linkRow(Copy.charOpenPlanner, glyph: "chevron.right") { openPlanner(link) }
            }
            if account.history, CharacterPage.isYours(profile.character, keys: Set(account.characterKeys)) {
                linkRow(Copy.charParseHistory, action: actions.openHistory)
            }
        }
        .dynamicTypeSize(...RankingsLayout.rowTypeCap)
        .padding(.horizontal, 16)
        .padding(.top, 8)
    }

    private var notFound: some View {
        VStack(alignment: .leading, spacing: 12) {
            VStack(alignment: .leading, spacing: 4) {
                Text(Copy.charNotFound(name: ref.name)).font(ScaledType.lead).foregroundStyle(FSColor.strong)
                    .fixedSize(horizontal: false, vertical: true)
                Text(Copy.charNotFoundHelp).font(ScaledType.meta).foregroundStyle(FSColor.muted)
                    .fixedSize(horizontal: false, vertical: true)
            }
            SecondaryButton(title: Copy.charSearchAgain, action: actions.searchAgain)
        }
        .padding(16)
        .frame(maxWidth: .infinity, alignment: .leading)
        .quietCard()
    }

    private func section(_ title: String, @ViewBuilder _ body: () -> some View) -> some View {
        VStack(alignment: .leading, spacing: 8) {
            MicroLabel(title).accessibilityAddTraits(.isHeader)
            body()
        }
    }

    private func fights(_ lines: [FightLine], recent: Bool) -> some View {
        VStack(spacing: 0) {
            ForEach(Array(lines.enumerated()), id: \.element.id) { index, line in
                if index > 0 { Divider().overlay(FSColor.lineSoft) }
                if recent, let open = actions.openReport {
                    Button { open(line) } label: { FightLineView(line: line, recent: true, chevron: true) }
                        .buttonStyle(RowPressStyle())
                } else {
                    FightLineView(line: line, recent: recent, chevron: false)
                }
            }
        }
        .quietCard()
    }

    @ViewBuilder
    private func gear(_ profile: CharacterProfile) -> some View {
        if profile.gear.isEmpty {
            Text(Copy.charNoGear(name: profile.character.name)).font(ScaledType.body).foregroundStyle(FSColor.text)
                .fixedSize(horizontal: false, vertical: true)
                .padding(16)
                .frame(maxWidth: .infinity, alignment: .leading)
                .quietCard()
        } else {
            VStack(alignment: .leading, spacing: 8) {
                VStack(alignment: .leading, spacing: 10) {
                    PaperdollView(
                        frame: PaperdollModel.frame(PaperdollModel.character(gear: profile.gear, items: items)),
                        build: build, open: actions.openSocket
                    )
                    Text(Copy.charGoldDot).font(ScaledType.meta).foregroundStyle(FSColor.muted)
                        .fixedSize(horizontal: false, vertical: true)
                }
                .padding(.horizontal, 8)
                .padding(.vertical, 14)
                .quietCard()
                if let gap = CharacterPage.gap(profile.gearGapTotal) {
                    HStack {
                        Text(Copy.charGap).font(ScaledType.bodyStrong).foregroundStyle(FSColor.strong)
                        Spacer()
                        Text(gap).font(ScaledType.mono).foregroundStyle(FSColor.strong)
                    }
                    .padding(.horizontal, 16)
                    .frame(minHeight: 52)
                    .quietCard()
                    .accessibilityElement(children: .combine)
                }
            }
        }
    }

    private func buildRows(_ builds: [BuildLine]) -> some View {
        VStack(spacing: 0) {
            ForEach(Array(builds.enumerated()), id: \.offset) { index, build in
                if index > 0 { Divider().overlay(FSColor.lineSoft) }
                HStack(spacing: 12) {
                    Image(systemName: "point.3.connected.trianglepath.dotted").font(.system(size: 16))
                        .foregroundStyle(FSColor.goldHover).frame(width: 36, height: 36)
                        .background(GameColors.controlFill, in: .rect(cornerRadius: FSRadius.control))
                        .accessibilityHidden(true)
                    VStack(alignment: .leading, spacing: 2) {
                        Text(build.trees).font(ScaledType.mono.weight(.bold)).foregroundStyle(FSColor.strong)
                        Text(build.meta).font(ScaledType.meta).foregroundStyle(FSColor.muted)
                            .fixedSize(horizontal: false, vertical: true)
                    }
                    .frame(maxWidth: .infinity, alignment: .leading)
                }
                .padding(.horizontal, 16)
                .padding(.vertical, 10)
                .frame(minHeight: 60)
                .accessibilityElement(children: .combine)
            }
        }
        .quietCard()
    }

    private func linkRow(_ title: String, glyph: String = "arrow.up.right", action: @escaping () -> Void) -> some View {
        Button(action: action) {
            HStack {
                Text(title).font(ScaledType.bodyStrong).foregroundStyle(FSColor.gold)
                Spacer()
                Image(systemName: glyph).foregroundStyle(FSColor.muted).accessibilityHidden(true)
            }
            .padding(.horizontal, 16)
            .frame(minHeight: 52)
            .contentShape(Rectangle())
        }
        .buttonStyle(RowPressStyle())
        .quietCard()
    }
}

/// The head (spec F11.4): ringed crest 64, the name in the class colour, the descriptor, where,
/// a `You` tag on your own and the supporter mark.
struct CharacterHead: View {
    let profile: CharacterProfile
    let specs: SpecCatalog
    let yours: Bool

    var body: some View {
        HStack(alignment: .center, spacing: 14) {
            ClassCrest(className: profile.character.className, size: RankingsLayout.headCrest)
            VStack(alignment: .leading, spacing: 2) {
                HStack(spacing: 8) {
                    Text(profile.character.name.uppercased()).font(ScaledType.h2)
                        .foregroundStyle(ClassColor.of(profile.character.className))
                        .fixedSize(horizontal: false, vertical: true)
                        .accessibilityAddTraits(.isHeader)
                    if yours { TagView(text: Copy.charYou, color: FSColor.gold) }
                    if profile.supporter { TagView(text: Copy.charSupporter, color: FSColor.goldHover, systemImage: "heart.fill") }
                }
                Text(CharacterPage.descriptor(profile, specs: specs)).font(ScaledType.meta).foregroundStyle(FSColor.muted)
                    .fixedSize(horizontal: false, vertical: true)
                Text(CharacterPage.whereLine(profile.character)).font(ScaledType.meta).foregroundStyle(FSColor.muted)
                    .fixedSize(horizontal: false, vertical: true)
            }
            .frame(maxWidth: .infinity, alignment: .leading)
        }
        .accessibilityElement(children: .combine)
    }
}

/// The Forever Score card (`--surface-card`): the figure with its tier, the date, the sparkline,
/// the five bars and `How is this calculated?`.
struct ScoreCard: View {
    let profile: CharacterProfile
    let openHow: () -> Void

    var body: some View {
        VStack(alignment: .leading, spacing: 12) {
            HStack(alignment: .top) {
                VStack(alignment: .leading, spacing: 2) {
                    MicroLabel(Copy.charScore, color: FSColor.gold)
                    ViewThatFits(in: .horizontal) {
                        HStack(alignment: .firstTextBaseline, spacing: 8) { figure; tier }
                        VStack(alignment: .leading, spacing: 4) { figure; tier }
                    }
                    if let asOf = CharacterPage.asOf(profile.asOf) {
                        Text(asOf).font(ScaledType.meta).foregroundStyle(FSColor.muted)
                    }
                }
                .accessibilityElement(children: .ignore)
                .accessibilityLabel(CharacterPage.scoreA11y(profile))
                Spacer(minLength: 8)
                Sparkline(points: CharacterPage.sparkline(profile.scoreHistory), color: ItemQuality(value: CharacterPage.tierQuality(profile.tier)).text)
                    .frame(width: RankingsLayout.sparkWidth, height: RankingsLayout.sparkHeight)
                    .accessibilityHidden(true)
            }
            VStack(spacing: 8) {
                ForEach(CharacterPage.bars(profile)) { ScoreBarRow(bar: $0) }
            }
            // The link's 44 pt target reaches past its line, so the card ends at its own padding.
            Button(Copy.charHow, action: openHow)
                .font(ScaledType.bodyStrong)
                .foregroundStyle(FSColor.gold)
                .contentShape(Rectangle().inset(by: -RankingsLayout.linkSlop))
        }
        .padding(.horizontal, 16)
        .padding(.vertical, 14)
        .frame(maxWidth: .infinity, alignment: .leading)
        .primaryCard()
    }
}

private extension ScoreCard {
    /// The figure never breaks (`1,83` / `0`); at the largest sizes the tier drops under it.
    var figure: some View {
        Text(CharacterPage.score(profile)).font(ScaledType.scoreFigure).foregroundStyle(FSColor.strong).fixedSize()
    }

    @ViewBuilder
    var tier: some View {
        if let quality = CharacterPage.tierQuality(profile.tier) {
            TagView(text: CharacterPage.tierName(profile.tier), color: ItemQuality(value: quality).text)
        }
    }
}

/// One 8 pt component bar with its label and value; at the accessibility sizes the label sits over it.
struct ScoreBarRow: View {
    let bar: ScoreBar
    @Environment(\.dynamicTypeSize) private var typeSize

    var body: some View {
        let layout = typeSize.isAccessibilitySize
            ? AnyLayout(VStackLayout(alignment: .leading, spacing: 4))
            : AnyLayout(HStackLayout(alignment: .center, spacing: 10))
        layout {
            Text(bar.label).font(ScaledType.meta).foregroundStyle(FSColor.muted)
                .frame(width: typeSize.isAccessibilitySize ? nil : RankingsLayout.barLabel, alignment: .leading)
            HStack(spacing: 10) {
                GeometryReader { proxy in
                    ZStack(alignment: .leading) {
                        Capsule().fill(FSColor.cardTop)
                        Capsule().fill(FSColor.gold).frame(width: proxy.size.width * bar.fraction)
                    }
                }
                .frame(height: RankingsLayout.bar)
                Text("\(bar.value)").font(ScaledType.monoMeta.weight(.semibold)).foregroundStyle(FSColor.strong)
                    .frame(minWidth: 34, alignment: .trailing)
                    .fixedSize()
            }
        }
        .accessibilityElement(children: .ignore)
        .accessibilityLabel(bar.a11y)
    }
}

/// The score history as a line, oldest first.
struct Sparkline: View {
    let points: [Double]
    let color: Color

    var body: some View {
        GeometryReader { proxy in
            Path { path in
                guard points.count > 1 else { return }
                let step = proxy.size.width / CGFloat(points.count - 1)
                for (index, point) in points.enumerated() {
                    let location = CGPoint(x: CGFloat(index) * step, y: proxy.size.height * (1 - point))
                    if index == 0 { path.move(to: location) } else { path.addLine(to: location) }
                }
            }
            .stroke(color, style: StrokeStyle(lineWidth: 2, lineCap: .round, lineJoin: .round))
        }
        .padding(2)
    }
}

/// A Bests or Recent fights row: the kill or wipe dot (recent), boss, meta, value and unit, and
/// on a best the parse number in its ladder colour.
struct FightLineView: View {
    let line: FightLine
    let recent: Bool
    let chevron: Bool
    @Environment(\.dynamicTypeSize) private var typeSize

    var body: some View {
        HStack(alignment: .center, spacing: 12) {
            if recent {
                Circle().fill(line.kill ? FSColor.kill : FSColor.wipe).frame(width: 10, height: 10).accessibilityHidden(true)
            }
            if typeSize.isAccessibilitySize {
                VStack(alignment: .leading, spacing: 4) { words; figure(.leading) }
                    .frame(maxWidth: .infinity, alignment: .leading)
            } else {
                words.frame(maxWidth: .infinity, alignment: .leading)
                figure(.trailing)
            }
            if chevron { ReportsChevron() }
        }
        .padding(.horizontal, 16)
        .padding(.vertical, 10)
        .frame(minHeight: recent ? RankingsLayout.fightRow : RankingsLayout.bestRow)
        .contentShape(Rectangle())
        .accessibilityElement(children: .combine)
    }

    private var words: some View {
        VStack(alignment: .leading, spacing: 2) {
            Text(line.encounter).font(ScaledType.bodyStrong).foregroundStyle(FSColor.strong)
                .fixedSize(horizontal: false, vertical: true)
            Text(line.meta).font(ScaledType.meta).foregroundStyle(FSColor.muted)
                .fixedSize(horizontal: false, vertical: true)
        }
    }

    private func figure(_ alignment: HorizontalAlignment) -> some View {
        VStack(alignment: alignment, spacing: 2) {
            HStack(alignment: .firstTextBaseline, spacing: 4) {
                Text(line.value).font(ScaledType.mono).foregroundStyle(FSColor.strong)
                Text(line.unit).font(ScaledType.micro).foregroundStyle(FSColor.muted)
            }
            if !recent, let percentile = line.percentile {
                HStack(alignment: .firstTextBaseline, spacing: 4) {
                    Text("\(percentile)").font(ReportsLayout.parseNumber).foregroundStyle(ParseLadder.color(ParseTier(percentile)))
                    Text(Copy.charPercentile).font(ScaledType.meta).foregroundStyle(FSColor.muted)
                }
            }
        }
        .fixedSize()
    }
}

/// The page while it first loads: the head, the score card and a section at their heights.
struct CharacterSkeleton: View {
    var body: some View {
        VStack(alignment: .leading, spacing: 16) {
            HStack(spacing: 14) {
                Circle().fill(FSColor.cardTop).frame(width: RankingsLayout.headCrest, height: RankingsLayout.headCrest)
                VStack(alignment: .leading, spacing: 8) { SkeletonBar(width: 160, height: 24); SkeletonBar(width: 220, height: 12) }
            }
            RoundedRectangle(cornerRadius: FSRadius.panel).fill(FSColor.cardTop).frame(height: 240)
            RoundedRectangle(cornerRadius: FSRadius.panel).fill(FSColor.cardTop).frame(height: RankingsLayout.bestRow * 2)
        }
        .padding(16)
        .accessibilityElement(children: .ignore)
        .accessibilityLabel(Copy.charScore)
    }
}
