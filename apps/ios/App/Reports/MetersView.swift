import DesignSystem
import Models
import SwiftUI

/// F10.3, the large-detent sheet over the report, drawn from the store's state.
struct MetersView: View {
    let target: MetersTarget
    let metric: MeterMetric
    let phase: MetersPhase
    let showAll: Bool
    let specs: SpecCatalog
    let build: String?
    var select: (MeterMetric) -> Void = { _ in }
    var expand: () -> Void = {}
    var retry: () -> Void = {}
    @Environment(\.dynamicTypeSize) private var typeSize
    @ScaledMetric(relativeTo: .subheadline) private var firstLine = ReportsLayout.firstLine

    var body: some View {
        ScrollView {
            VStack(alignment: .leading, spacing: 12) {
                Text(title.uppercased()).font(ScaledType.sec).foregroundStyle(FSColor.strong)
                    .fixedSize(horizontal: false, vertical: true)
                    .accessibilityAddTraits(.isHeader)
                    .padding(.horizontal, 16)
                head.padding(.horizontal, 16)
                // A flagged report has nothing to switch between: only its sentence (spec F10.3).
                if phase != .flagged { MeterSegments(metric: metric, select: select).padding(.horizontal, 16) }
                content
            }
            .frame(maxWidth: .infinity, alignment: .leading)
            .padding(.top, 24)
            .padding(.bottom, 32)
        }
        .background(GameColors.sheet)
    }

    private var title: String {
        switch target.kind {
        case .boss: ReportsCopy.metersTitle(metric)
        case .trash: Copy.reportTrashTitle(zone: target.name)
        case .loneTrash: Copy.reportTrash
        }
    }

    private var meters: FightMeters? {
        if case let .loaded(meters) = phase { return meters }
        return nil
    }

    private var head: some View {
        HStack(alignment: .top, spacing: 8) {
            Circle().fill(dot).frame(width: ReportsLayout.outcomeDot, height: ReportsLayout.outcomeDot)
                .frame(height: firstLine)
            let meta = Text(ReportsCopy.metersHead(target, meters: meters)).font(ScaledType.meta).foregroundStyle(FSColor.muted)
            if target.isTrash {
                meta.fixedSize(horizontal: false, vertical: true)
            } else {
                Text("\(Text(target.name).font(ScaledType.bodyStrong).foregroundStyle(FSColor.strong))  \(meta)")
                    .fixedSize(horizontal: false, vertical: true)
            }
        }
        .accessibilityElement(children: .combine)
    }

    private var dot: Color {
        if target.isTrash { return FSColor.muted }
        return (meters?.kill ?? target.kill) ? FSColor.kill : FSColor.wipe
    }

    @ViewBuilder
    private var content: some View {
        switch phase {
        case .loading:
            MeterSkeleton()
        case .failed:
            ReportsErrorCard(text: Copy.metersFailed, retry: retry).padding(.horizontal, 16)
        case .flagged:
            Text(Copy.metersFlagged).font(ScaledType.body).foregroundStyle(FSColor.text).padding(.horizontal, 16)
        case let .loaded(meters) where meters.players.isEmpty:
            Text(Copy.metersEmpty).font(ScaledType.body).foregroundStyle(FSColor.muted).padding(.horizontal, 16)
        case let .loaded(meters):
            rows(meters.players)
        }
    }

    private var limit: Int { typeSize.isAccessibilitySize ? MeterRows.largeVisible : MeterRows.defaultVisible }

    private func rows(_ players: [MeterPlayer]) -> some View {
        let top = players.first.map(metric.value) ?? 0
        return VStack(alignment: .leading, spacing: 2) {
            ForEach(MeterRows.visible(players, limit: limit, showAll: showAll), id: \.rank) { row in
                MeterRowView(
                    rank: row.rank, player: row.player, metric: metric,
                    fraction: MeterRows.fraction(metric.value(row.player), top: top),
                    icon: row.player.spec.flatMap { specs.spec($0)?.icon }, build: build
                )
            }
            if MeterRows.hasMore(players, limit: limit, showAll: showAll) {
                Button(Copy.metersShowAll(n: players.count), action: expand)
                    .font(ScaledType.bodyStrong)
                    .foregroundStyle(FSColor.gold)
                    .frame(minHeight: ReportsLayout.meterRow)
                    .padding(.horizontal, 16)
            }
        }
    }
}

/// Damage, Healing, Damage taken; a change plays the selection haptic.
struct MeterSegments: View {
    let metric: MeterMetric
    let select: (MeterMetric) -> Void
    @Environment(\.dynamicTypeSize) private var typeSize

    var body: some View {
        let layout = typeSize.isAccessibilitySize ? AnyLayout(VStackLayout(spacing: 0)) : AnyLayout(HStackLayout(spacing: 0))
        layout {
            ForEach(MeterMetric.allCases) { item in
                let checked = item == metric
                Button { select(item) } label: {
                    Text(ReportsCopy.segment(item))
                        .font(ScaledType.segment)
                        .foregroundStyle(checked ? FSColor.strong : FSColor.muted)
                        .frame(maxWidth: .infinity, minHeight: 44)
                        .background(checked ? FSColor.gold.opacity(0.12) : .clear)
                        .overlay(alignment: .bottom) { if checked { FSColor.gold.frame(height: 2) } }
                        .contentShape(Rectangle())
                }
                .buttonStyle(.plain)
                .accessibilityAddTraits(checked ? .isSelected : [])
            }
        }
        .background(FSColor.raised)
        .clipShape(.rect(cornerRadius: FSRadius.control))
        .overlay(RoundedRectangle(cornerRadius: FSRadius.control).stroke(FSColor.line, lineWidth: 1))
        .haptic(.selection, trigger: metric)
    }
}

/// One Details!-style row (spec §4 Meter row): the bar at 22% of the class colour behind rank,
/// icon, name and the value; your row on the `--row` ground with the gold left edge. At the
/// accessibility sizes the value sits under the name.
struct MeterRowView: View {
    let rank: Int
    let player: MeterPlayer
    let metric: MeterMetric
    let fraction: Double
    let icon: String?
    let build: String?
    @Environment(\.dynamicTypeSize) private var typeSize
    @ScaledMetric(relativeTo: .subheadline) private var iconSize = ReportsLayout.meterIcon

    var body: some View {
        HStack(alignment: .center, spacing: 10) {
            Text("\(rank)").font(ScaledType.monoMeta).foregroundStyle(FSColor.muted).frame(minWidth: 20, alignment: .leading)
            mark
            if typeSize.isAccessibilitySize {
                VStack(alignment: .leading, spacing: 2) { name; value }
                    .frame(maxWidth: .infinity, alignment: .leading)
            } else {
                name.frame(maxWidth: .infinity, alignment: .leading)
                value
            }
        }
        .padding(.horizontal, 16)
        .padding(.vertical, 6)
        .frame(minHeight: ReportsLayout.meterRow)
        .background(alignment: .leading) { bar }
        .background(player.mine ? GameColors.rowWash : .clear)
        .overlay(alignment: .leading) {
            if player.mine { FSColor.gold.frame(width: ReportsLayout.mineEdge) }
        }
        .accessibilityElement(children: .ignore)
        .accessibilityLabel(ReportsCopy.meterA11y(rank: rank, player: player, metric: metric))
    }

    private var color: Color { ClassColor.of(player.className) }

    private var bar: some View {
        GeometryReader { proxy in
            color.opacity(ReportsLayout.barOpacity).frame(width: proxy.size.width * fraction)
        }
    }

    @ViewBuilder
    private var mark: some View {
        let size = min(iconSize, ReportsLayout.meterIconCap)
        if let icon, let build {
            SpecIcon(icon: icon, build: build, size: size)
        } else {
            ClassCrest(className: player.className, size: size)
        }
    }

    private var name: some View {
        Text(player.name).font(ScaledType.bodyStrong).foregroundStyle(color)
            .fixedSize(horizontal: false, vertical: true)
    }

    private var value: some View {
        HStack(alignment: .firstTextBaseline, spacing: 4) {
            Text(ReportsCopy.figure(metric.value(player))).font(ScaledType.mono).foregroundStyle(FSColor.strong)
            Text(ReportsCopy.unit(metric.unit)).font(ScaledType.micro).foregroundStyle(FSColor.muted)
        }
        .fixedSize()
    }
}

/// F10.3 loading: eight meter rows at height.
struct MeterSkeleton: View {
    static let rows = 8
    /// Each skeleton bar is this much shorter than the one above, as a meter tapers.
    static let taper = 0.08

    var body: some View {
        VStack(spacing: 2) {
            ForEach(0..<Self.rows, id: \.self) { index in
                HStack(spacing: 10) {
                    SkeletonBar(width: ReportsLayout.meterIcon, height: ReportsLayout.meterIcon)
                    SkeletonBar(width: 140, height: 14)
                    Spacer()
                    SkeletonBar(width: 48, height: 14)
                }
                .padding(.horizontal, 16)
                .frame(height: ReportsLayout.meterRow)
                .background(alignment: .leading) {
                    GeometryReader { proxy in
                        FSColor.raised.frame(width: proxy.size.width * (1 - Double(index) * Self.taper))
                    }
                }
            }
        }
        .accessibilityElement(children: .ignore)
        .accessibilityLabel(Copy.metersTitleDamage)
    }
}
