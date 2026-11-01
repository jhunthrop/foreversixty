import DesignSystem
import Models
import SwiftUI

/// F9.1 under the segments, as drawn from a phase: `Your guide`, the class grid, the footer.
struct GuidesHomeView: View {
    let phase: StaticPhase<GuidesHome>
    let character: AccountCharacter?
    var openGuide: (GuideSpecRef) -> Void = { _ in }
    var openClass: (GuideClassCell) -> Void = { _ in }
    var retry: () -> Void = {}

    var body: some View {
        switch phase {
        case .loading: GuidesHomeSkeleton()
        case .notCached: GuidesNotCached(text: Copy.guidesListNotCached)
        case .failed: ReportsErrorCard(text: Copy.guidesError, retry: retry)
        case let .loaded(home, _): loaded(home)
        }
    }

    private func loaded(_ home: GuidesHome) -> some View {
        VStack(alignment: .leading, spacing: 0) {
            if let mine = GuidesIndexLogic.yourGuide(home.index, character: character) {
                YourGuideCard(spec: mine, build: home.build, saved: home.saved.isSaved(mine.spec)) { openGuide(mine) }
                    .padding(.bottom, 4)
            }
            GuideSectionLabel(text: Copy.guidesClasses)
            ClassGrid(cells: GuidesIndexLogic.cells(home.index), open: openClass)
            Text(home.saved.footer(home.index))
                .font(ScaledType.meta).foregroundStyle(FSColor.muted)
                .fixedSize(horizontal: false, vertical: true)
                .padding(.top, 12)
        }
    }
}

/// `Your guide` (the one `--surface-card`): one 80 pt target.
struct YourGuideCard: View {
    let spec: GuideSpecRef
    let build: String
    let saved: Bool
    let open: () -> Void

    var body: some View {
        Button(action: open) {
            HStack(alignment: .center, spacing: 14) {
                SpecIcon(icon: spec.icon, build: build, size: GuidesLayout.headIcon)
                VStack(alignment: .leading, spacing: 4) {
                    MicroLabel(Copy.guidesYours, color: FSColor.gold)
                    Text(spec.title.uppercased()).font(ScaledType.sec).foregroundStyle(ClassColor.of(spec.classSlug))
                        .fixedSize(horizontal: false, vertical: true)
                    Text(GuideCopy.cardMeta(spec)).font(ScaledType.meta).foregroundStyle(FSColor.muted)
                        .fixedSize(horizontal: false, vertical: true)
                    if saved {
                        TagView(text: Copy.guidesSaved, color: FSColor.kill, systemImage: "checkmark").padding(.top, 4)
                    }
                }
                .frame(maxWidth: .infinity, alignment: .leading)
                Image(systemName: "chevron.right").font(.system(size: 14, weight: .semibold)).foregroundStyle(FSColor.muted)
                    .accessibilityHidden(true)
            }
            .padding(GuidesLayout.cardPadding)
            .frame(maxWidth: .infinity, minHeight: GuidesLayout.yourCard, alignment: .leading)
            .contentShape(Rectangle())
        }
        .buttonStyle(RowPressStyle())
        .primaryCard()
        .accessibilityElement(children: .combine)
    }
}

/// `All classes`: nine ringed crests, three across (two at the accessibility sizes).
struct ClassGrid: View {
    let cells: [GuideClassCell]
    let open: (GuideClassCell) -> Void
    @Environment(\.dynamicTypeSize) private var typeSize

    private var columns: [GridItem] {
        Array(repeating: GridItem(.flexible(), spacing: 8), count: typeSize.isAccessibilitySize ? 2 : 3)
    }

    var body: some View {
        LazyVGrid(columns: columns, spacing: 8) {
            ForEach(cells) { cell in
                Button { open(cell) } label: {
                    VStack(spacing: 6) {
                        ClassCrest(className: cell.classSlug, size: GuidesLayout.crest)
                        Text(cell.name).font(ScaledType.meta.weight(.semibold)).foregroundStyle(ClassColor.of(cell.classSlug))
                            .multilineTextAlignment(.center)
                            .fixedSize(horizontal: false, vertical: true)
                    }
                    .frame(maxWidth: .infinity, minHeight: 88)
                    .contentShape(Rectangle())
                }
                .buttonStyle(RowPressStyle())
                .accessibilityLabel(cell.name)
            }
        }
        .padding(GuidesLayout.cardPadding)
        .quietCard()
    }
}

/// F9.7 for the tab: the card and the grid at their real heights; the segments stay live.
struct GuidesHomeSkeleton: View {
    var body: some View {
        VStack(alignment: .leading, spacing: 16) {
            SkeletonBar(width: nil, height: 104)
            SkeletonBar(width: 90, height: 11).padding(.top, 8)
            SkeletonBar(width: nil, height: 300)
        }
        .accessibilityElement(children: .ignore)
        .accessibilityLabel(Copy.guidesTitle)
    }
}

/// F9.6: a class's guides, each Saved (kill, check) or Not saved (muted, cloud-down).
struct GuideClassView: View {
    let classSlug: String
    let home: GuidesHome
    let offlineSince: Date?
    let now: Date
    let open: (GuideSpecRef) -> Void

    private var specs: [GuideSpecRef] { GuidesIndexLogic.specs(of: classSlug, in: home.index) }

    var body: some View {
        VStack(alignment: .leading, spacing: 0) {
            GuidesOfflineStrip(since: offlineSince, now: now)
            Text((specs.first?.className ?? classSlug).uppercased())
                .font(ScaledType.title)
                .foregroundStyle(ClassColor.of(classSlug))
                .accessibilityAddTraits(.isHeader)
                .padding(.horizontal, GuidesLayout.gutter).padding(.top, 8).padding(.bottom, 12)
            GuideCardRows(items: specs) { spec in
                Button { open(spec) } label: { GuideSpecRow(spec: spec, build: home.build, saved: home.saved.isSaved(spec.spec)) }
                    .buttonStyle(RowPressStyle())
            }
            .padding(.horizontal, GuidesLayout.gutter)
        }
        .frame(maxWidth: .infinity, alignment: .leading)
    }
}

struct GuideSpecRow: View {
    let spec: GuideSpecRef
    let build: String
    let saved: Bool
    @Environment(\.dynamicTypeSize) private var typeSize

    var body: some View {
        HStack(alignment: .center, spacing: 12) {
            SpecIcon(icon: spec.icon, build: build, size: GuidesLayout.rowIcon)
            VStack(alignment: .leading, spacing: 3) {
                Text(spec.specName).font(ScaledType.name).foregroundStyle(FSColor.strong)
                Text(GuideCopy.rowMeta(spec)).font(ScaledType.meta).foregroundStyle(FSColor.muted)
                if typeSize.isAccessibilitySize { savedMark }
            }
            .fixedSize(horizontal: false, vertical: true)
            .frame(maxWidth: .infinity, alignment: .leading)
            if !typeSize.isAccessibilitySize { savedMark }
            Image(systemName: "chevron.right").font(.system(size: 13, weight: .semibold)).foregroundStyle(FSColor.muted)
                .accessibilityHidden(true)
        }
        .padding(.horizontal, GuidesLayout.cardPadding)
        .padding(.vertical, 12)
        .frame(minHeight: 64)
        .contentShape(Rectangle())
        .accessibilityElement(children: .combine)
    }

    @ViewBuilder
    private var savedMark: some View {
        if saved {
            TagView(text: Copy.guidesSavedShort, color: FSColor.kill, systemImage: "checkmark")
        } else {
            HStack(spacing: 4) {
                Image(systemName: "icloud.and.arrow.down").font(.system(size: 14)).accessibilityHidden(true)
                Text(Copy.guidesNotSaved).font(ScaledType.meta)
            }
            .foregroundStyle(FSColor.muted)
        }
    }
}
