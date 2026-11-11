import DesignSystem
import Models
import SwiftUI

/// What the reader hands back to its screen.
struct GuideActions {
    var openGear: (GuideGear) -> Void = { _ in }
    var retry: () -> Void = {}
    var update: () -> Void = {}
    /// Lane F13: the talents section's `Open this build in the planner` row (an FS1 code); nil
    /// while `planner` is off, and the row hides.
    var openPlanner: ((_ classSlug: String, _ code: String) -> Void)?
}

/// F9.2 as drawn from a phase: stateless apart from the scroll, so snapshots drive every state.
struct GuideReaderView: View {
    let phase: StaticPhase<GuideContent>
    var newData: NewDataState = .none
    let now: Date
    var actions = GuideActions()
    /// The guide's name while it is not drawn (loading, not saved, failed): the band head says it, as
    /// the character page's states do (N3).
    var title = Copy.guidesTitle

    var body: some View {
        switch phase {
        case .loading:
            ScrollView {
                BandHead(title: title)
                GuideSkeleton().padding(.horizontal, GuidesLayout.gutter).padding(.top, BandReach.gap)
            }
        case .notCached:
            ScrollView {
                BandHead(title: title)
                GuidesNotCached(text: Copy.guidesNotCached).padding(GuidesLayout.gutter)
            }
        case .failed:
            ScrollView {
                BandHead(title: title)
                LoadError(Copy.guidesError, retryLabel: Copy.commonTryAgain, retry: actions.retry).padding(GuidesLayout.gutter)
            }
        case let .loaded(content, since):
            GuideScroll(content: content, actions: actions) {
                NewDataStrip(newData, action: actions.update)
                GuidesOfflineStrip(since: since, now: now)
            }
        }
    }
}

/// The loaded guide: the head band scrolls away, the section strip sticks under the bar, and
/// the strip follows the scroll. `strips` (new data, offline) sit under the band.
struct GuideScroll<Strips: View>: View {
    let content: GuideContent
    var actions = GuideActions()
    @ViewBuilder var strips: () -> Strips

    @State private var tops: [GuideSection: CGFloat] = [:]

    private static var space: String { "guide" }

    private var sections: [GuideSection] { GuideSection.of(content.guide) }

    /// The last section whose top has passed under the strip; the first before any has.
    private var current: GuideSection? {
        sections.last { (tops[$0] ?? .infinity) <= GuidesLayout.currentThreshold } ?? sections.first
    }

    var body: some View {
        ScrollViewReader { proxy in
            ScrollView {
                LazyVStack(alignment: .leading, spacing: 0, pinnedViews: [.sectionHeaders]) {
                    GuideHead(guide: content.guide, build: content.build)
                    strips()
                    Section {
                        GuideBody(content: content, space: Self.space, actions: actions) { tops[$0] = $1 }
                        .padding(.bottom, GuidesLayout.tabBarClearance)
                    } header: {
                        GuideSectionStrip(sections: sections, current: current) { section in
                            withAnimation { proxy.scrollTo(section, anchor: .top) }
                        }
                    }
                }
            }
            .coordinateSpace(.named(Self.space))
            .screenGround()
        }
    }
}

/// The head band (`--surface-band`): spec icon 56, `{Spec} {Class}` in the class colour, the meta.
struct GuideHead: View {
    let guide: GuideFile
    let build: String

    var body: some View {
        HStack(alignment: .top, spacing: 14) {
            SpecIcon(icon: guide.icon, build: build, size: GuidesLayout.headIcon)
            VStack(alignment: .leading, spacing: 4) {
                Text(GuideCopy.title(guide))
                    .font(ScaledType.h2)
                    .foregroundStyle(ClassColor.of(guide.classSlug))
                    .fixedSize(horizontal: false, vertical: true)
                    .accessibilityAddTraits(.isHeader)
                Text(GuideCopy.headMeta(guide)).font(ScaledType.meta).foregroundStyle(FSColor.mutedGround)
                    .fixedSize(horizontal: false, vertical: true)
            }
            .frame(maxWidth: .infinity, alignment: .leading)
        }
        .padding(GuidesLayout.gutter)
        .frame(maxWidth: .infinity, alignment: .leading)
        // The band runs up under the clear bar and the status bar (N3).
        .bandUnderBars()
    }
}

/// The sticky strip: one 36 pt pill per section the guide has; the current one gold. A tap
/// scrolls; the scroll moves the gold; the current pill is kept in view.
struct GuideSectionStrip: View {
    let sections: [GuideSection]
    let current: GuideSection?
    let select: (GuideSection) -> Void

    var body: some View {
        ScrollViewReader { proxy in
            ScrollView(.horizontal, showsIndicators: false) {
                HStack(spacing: 8) {
                    ForEach(sections) { section in
                        chip(section).id(section)
                    }
                }
                .padding(.horizontal, GuidesLayout.gutter)
                .padding(.vertical, 8)
            }
            .onAppear { if let current { proxy.scrollTo(current, anchor: .center) } }
            .onChange(of: current) { _, section in
                guard let section else { return }
                withAnimation { proxy.scrollTo(section, anchor: .center) }
            }
        }
        .screenGround()
        .overlay(alignment: .bottom) { FSColor.lineSoft.frame(height: 1) }
        .haptic(.selection, trigger: current)
    }

    private func chip(_ section: GuideSection) -> some View {
        let isCurrent = section == current
        let shape = RoundedRectangle(cornerRadius: GuidesLayout.chip / 2)
        return Button { select(section) } label: {
            Text(section.title)
                .font(ScaledType.bodyStrong)
                .foregroundStyle(isCurrent ? FSColor.gold : FSColor.nav)
                .lineLimit(1)
                .fixedSize()
                .padding(.horizontal, 14)
                .padding(.vertical, 4)
                .frame(minHeight: GuidesLayout.chip)
                .background(isCurrent ? FSColor.gold.opacity(0.12) : GameColors.controlFill, in: shape)
                .overlay(shape.strokeBorder(isCurrent ? FSColor.gold : FSColor.lineCard, lineWidth: 1))
                .frame(minHeight: 44)
                .contentShape(Rectangle())
        }
        .buttonStyle(.plain)
        .accessibilityLabel(GuideSection.a11y(section, in: sections))
        .accessibilityAddTraits(isCurrent ? .isSelected : [])
    }
}

/// F9.7 for a guide: the head, the strip and the at-a-glance card at their real heights.
struct GuideSkeleton: View {
    var body: some View {
        VStack(alignment: .leading, spacing: 14) {
            HStack(spacing: 14) {
                SkeletonBar(width: GuidesLayout.headIcon, height: GuidesLayout.headIcon)
                VStack(alignment: .leading, spacing: 8) { SkeletonBar(width: 200, height: 24); SkeletonBar(width: 240, height: 12) }
            }
            HStack(spacing: 8) { ForEach(0..<4, id: \.self) { _ in SkeletonBar(width: 76, height: GuidesLayout.chip) } }
            SkeletonBar(width: nil, height: 180)
            SkeletonBar(width: nil, height: 240)
        }
        .accessibilityElement(children: .ignore)
        .accessibilityLabel(Copy.guidesTitle)
    }
}
