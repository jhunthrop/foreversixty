import DesignSystem
import Models
import SwiftUI

/// An item of the gear grid, as the build's items file names it (nil while loading or unknown).
struct SimGearLook: Equatable, Sendable {
    let name: String
    let icon: String
    let quality: Int
}

/// F5.3 as drawn from a phase: stateless, so snapshots drive every state.
struct SimResultView: View {
    let phase: SimResultPhase
    let looks: [Int: SimGearLook]
    let build: String?
    let now: Date
    var openItem: (SimGearCell) -> Void = { _ in }
    var openBreakdown: (URL) -> Void = { _ in }
    var retry: () -> Void = {}
    /// Lane F15: the stale strip's time.
    var staleSince: Date?

    private var title: String {
        if case let .loaded(model, _) = phase { return model.title }
        return Copy.simsTitle
    }

    static let gridColumns = 6
    static let gearIcon: CGFloat = 40
    static let gridGap: CGFloat = 8

    var body: some View {
        VStack(alignment: .leading, spacing: 0) {
            SimsTitle(text: title)
            switch phase {
            case .loading: SimResultSkeleton()
            case .failed:
                VStack(alignment: .leading, spacing: 12) {
                    Text(Copy.simsResultFailed).font(ScaledType.lead).foregroundStyle(FSColor.wipe)
                    SecondaryButton(title: Copy.commonTryAgain, systemImage: "arrow.clockwise", action: retry)
                }
                .padding(16)
            case let .loaded(model, since):
                if let since { OfflineStrip(text: Copy.offlineStrip(time: OfflineStamp.format(since, now: now))) }
                if let staleSince { StaleStrip(time: OfflineStamp.format(staleSince, now: now), retry: retry) }
                content(model)
            }
        }
        .frame(maxWidth: .infinity, alignment: .leading)
        .padding(.bottom, SimsListView.tabBarClearance)
    }

    private func content(_ model: SimResultModel) -> some View {
        VStack(alignment: .leading, spacing: 0) {
            MicroLabel(model.kind)
            meta(model).padding(.top, 4)
            answer(model.answer).padding(.top, 12)
            Text(model.errorLine).font(ScaledType.monoMeta).foregroundStyle(FSColor.muted)
                .fixedSize(horizontal: false, vertical: true).padding(.top, 4)
            if !model.gear.isEmpty { gear(model).padding(.top, 22) }
            Button { openBreakdown(model.breakdownURL) } label: {
                HStack(spacing: 6) {
                    Text(Copy.simsBreakdown)
                    Image(systemName: "arrow.up.right.square").accessibilityHidden(true)
                }
                .font(ScaledType.bodyStrong)
                .foregroundStyle(FSColor.gold)
                .frame(minHeight: 44)
            }
            .buttonStyle(.plain)
            .padding(.top, 8)
        }
        .padding(.horizontal, 16)
        .padding(.top, 4)
    }

    private func meta(_ model: SimResultModel) -> some View {
        let level = model.level.map { " · " + Copy.simsLevel(n: $0) } ?? ""
        let className = Text(model.className).foregroundStyle(ClassColor.of(model.classSlug))
        return Text("\(model.specName) \(className)\(level)")
            .font(ScaledType.meta)
            .foregroundStyle(FSColor.muted)
    }

    @ViewBuilder
    private func answer(_ answer: SimResultModel.Answer) -> some View {
        switch answer {
        case let .figure(value):
            ViewThatFits(in: .horizontal) {
                HStack(alignment: .firstTextBaseline, spacing: 6) { figure(value); dpsWord }
                VStack(alignment: .leading, spacing: 0) { figure(value); dpsWord }
            }
            .accessibilityElement(children: .combine)
        case let .sentence(text):
            Text(text).font(ScaledType.stat).foregroundStyle(FSColor.strong)
                .fixedSize(horizontal: false, vertical: true)
        }
    }

    private func figure(_ value: String) -> some View {
        Text(value).font(SimResultView.figureFont).foregroundStyle(FSColor.strong)
    }

    private var dpsWord: some View {
        Text(Copy.bisUnitDps).font(ScaledType.body).foregroundStyle(FSColor.muted)
    }

    /// The 52 pt figure, growing with the text size (to 64 at AX5, spec F5.3).
    static let figureFont = Font.custom("BarlowCondensed-ExtraBold", size: 52, relativeTo: .largeTitle)

    private func gear(_ model: SimResultModel) -> some View {
        VStack(alignment: .leading, spacing: 8) {
            MicroLabel(Copy.simsGear)
            LazyVGrid(
                columns: Array(repeating: GridItem(.fixed(Self.gearIcon), spacing: Self.gridGap), count: Self.gridColumns),
                alignment: .leading, spacing: Self.gridGap
            ) {
                ForEach(model.gear) { cell in gearIcon(cell) }
            }
            if model.showsChangedLegend {
                Text(Copy.simsChanged).font(ScaledType.meta).foregroundStyle(FSColor.muted)
            }
        }
    }

    @ViewBuilder
    private func gearIcon(_ cell: SimGearCell) -> some View {
        let look = looks[cell.itemID]
        let icon = ItemIconView(icon: look?.icon ?? Self.unknownIcon, quality: look?.quality, build: build ?? "", size: Self.gearIcon)
            .overlay(alignment: .topTrailing) { if cell.changed { ChangedCorner() } }
        if let look {
            Button { openItem(cell) } label: { icon }
                .buttonStyle(.plain)
                .frame(minWidth: 44, minHeight: 44)
                .accessibilityLabel(SimResultView.gearLabel(look, cell: cell))
        } else {
            icon.accessibilityLabel(ItemVocabulary.slotLabel(cell.slot))
        }
    }

    static func gearLabel(_ look: SimGearLook, cell: SimGearCell) -> String {
        let base = "\(look.name), \(ItemVocabulary.slotLabel(cell.slot))"
        return cell.changed ? "\(base), \(Copy.simsChangedLabel)" : base
    }

    /// The game's question-mark icon, for an id the build's items file does not list.
    static let unknownIcon = "inv_misc_questionmark"
}

/// The 12 pt gold corner on a slot the result changes.
struct ChangedCorner: View {
    var body: some View {
        Path { path in
            path.move(to: .zero)
            path.addLine(to: CGPoint(x: 12, y: 0))
            path.addLine(to: CGPoint(x: 12, y: 12))
            path.closeSubpath()
        }
        .fill(FSColor.gold)
        .frame(width: 12, height: 12)
        .accessibilityHidden(true)
    }
}

/// F5.3 loading: the title, kind and figure as blocks, then six squares.
struct SimResultSkeleton: View {
    var body: some View {
        VStack(alignment: .leading, spacing: 10) {
            bar(80, 12)
            bar(160, 14)
            bar(140, 52)
            bar(240, 12)
            HStack(spacing: SimResultView.gridGap) {
                ForEach(0..<SimResultView.gridColumns, id: \.self) { _ in bar(SimResultView.gearIcon, SimResultView.gearIcon) }
            }
            .padding(.top, 16)
        }
        .padding(16)
        .accessibilityElement(children: .ignore)
        .accessibilityLabel(Copy.simsTitle)
    }

    private func bar(_ width: CGFloat, _ height: CGFloat) -> some View {
        RoundedRectangle(cornerRadius: 4).fill(FSColor.cardTop).frame(width: width, height: height)
    }
}
