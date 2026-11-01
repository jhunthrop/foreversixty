import DesignSystem
import Models
import SwiftUI

/// What the reader hands back to its screen.
struct GuideActions {
    var openGear: (GuideGear) -> Void = { _ in }
    var retry: () -> Void = {}
    var update: () -> Void = {}
}

/// F9.2 as drawn from a phase: stateless apart from the scroll, so snapshots drive every state.
struct GuideReaderView: View {
    let phase: StaticPhase<GuideContent>
    var newData: NewDataState = .none
    let now: Date
    var actions = GuideActions()

    var body: some View {
        switch phase {
        case .loading:
            ScrollView { GuideSkeleton().padding(.horizontal, GuidesLayout.gutter).padding(.top, 12) }
        case .notCached:
            ScrollView { GuidesNotCached(text: Copy.guidesNotCached).padding(GuidesLayout.gutter) }
        case .failed:
            ScrollView { ReportsErrorCard(text: Copy.guidesError, retry: actions.retry).padding(GuidesLayout.gutter) }
        case let .loaded(content, since):
            GuideScroll(content: content, actions: actions)
                .safeAreaInset(edge: .top, spacing: 0) {
                    VStack(spacing: 0) {
                        NewDataStrip(newData, action: actions.update)
                        GuidesOfflineStrip(since: since, now: now)
                    }
                }
        }
    }
}

/// The loaded guide: the head band scrolls away, the section strip sticks under the bar, and
/// the strip follows the scroll.
struct GuideScroll: View {
    let content: GuideContent
    var actions = GuideActions()

    @State private var tops: [GuideSection: CGFloat] = [:]

    private static let space = "guide"

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
            .background(FSColor.bg)
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
                Text(GuideCopy.title(guide).uppercased())
                    .font(ScaledType.h2)
                    .foregroundStyle(ClassColor.of(guide.classSlug))
                    .fixedSize(horizontal: false, vertical: true)
                    .accessibilityAddTraits(.isHeader)
                Text(GuideCopy.headMeta(guide)).font(ScaledType.meta).foregroundStyle(FSColor.muted)
                    .fixedSize(horizontal: false, vertical: true)
            }
            .frame(maxWidth: .infinity, alignment: .leading)
        }
        .padding(GuidesLayout.gutter)
        .frame(maxWidth: .infinity, alignment: .leading)
        .surfaceBand()
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
        .background(FSColor.bg)
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
