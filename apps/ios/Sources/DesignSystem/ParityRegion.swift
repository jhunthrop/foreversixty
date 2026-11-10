import SwiftUI

/// A native parity region (design/parity.md "Native"): the web's `data-region` names
/// (web/src/lib/ui/regions.ts `REGION`, held equal by web/src/lib/mobile/parity-regions.test.ts),
/// so a region of a mock pairs with the same region of the build.
public struct ParityRegionName: Hashable, Sendable {
    public let rawValue: String

    public init(_ rawValue: String) {
        self.rawValue = rawValue
    }

    public static let head = Self("head")
    public static let headTitle = Self("head-title")
    public static let headMeta = Self("head-meta")
    public static let headAnswer = Self("head-answer")
    public static let controls = Self("controls")
    public static let summary = Self("summary")
    public static let countdowns = Self("countdowns")
    public static let notes = Self("notes")
    public static let plate = Self("plate")
    public static let slots = Self("slots")
    public static let inspector = Self("inspector")
    public static let inspectorInline = Self("inspector-inline")
    public static let footerNote = Self("footer-note")
    public static let intro = Self("intro")
    public static let dialog = Self("dialog")

    /// A Section's region, `section-<slug of its title>`: the web's `sectionRegion` rule.
    public static func section(_ title: String) -> Self {
        let words = title.lowercased().split { !($0.isASCII && ($0.isLetter || $0.isNumber)) }
        return Self("section-" + words.joined(separator: "-"))
    }
}

/// What a region's inventory counts (design/parity.md "Native").
public enum ParityItemKind: String, Sendable, CaseIterable {
    case row, button, segment, field, image
}

/// One counted element and its frame in the content coordinate space.
public struct ParityItem: Equatable, Sendable {
    public let kind: ParityItemKind
    public let frame: CGRect
}

/// The coordinate space region boxes are measured in: the screen's content, under the
/// navigation bar.
public enum ParityContent {
    public static let space = "parity-content"
}

struct ParityRegionsKey: PreferenceKey {
    static let defaultValue: [String: CGRect] = [:]

    static func reduce(value: inout [String: CGRect], nextValue: () -> [String: CGRect]) {
        value.merge(nextValue()) { first, _ in first }
    }
}

struct ParityItemsKey: PreferenceKey {
    static let defaultValue: [ParityItem] = []

    static func reduce(value: inout [ParityItem], nextValue: () -> [ParityItem]) {
        value.append(contentsOf: nextValue())
    }
}

public extension View {
    /// Reports this view's frame as the named region. Costs nothing outside a parity capture.
    func parityRegion(_ name: ParityRegionName) -> some View {
        background {
            GeometryReader { geo in
                Color.clear.preference(
                    key: ParityRegionsKey.self,
                    value: [name.rawValue: geo.frame(in: .named(ParityContent.space))]
                )
            }
        }
    }

    /// Counts this view in the inventory of every region that contains it.
    func parityItem(_ kind: ParityItemKind) -> some View {
        background {
            GeometryReader { geo in
                Color.clear.preference(
                    key: ParityItemsKey.self,
                    value: [ParityItem(kind: kind, frame: geo.frame(in: .named(ParityContent.space)))]
                )
            }
        }
    }

    /// Leaves this view's items out of every region: a bar pinned over the content (the applicant's
    /// Accept / Decline) that the mock draws in the page's flow after the content.
    func parityPinned() -> some View {
        transformPreference(ParityItemsKey.self) { $0 = [] }
    }

    /// Marks the screen's content (under the navigation bar) as the space boxes are measured in.
    func parityContent() -> some View {
        coordinateSpace(.named(ParityContent.space))
    }

    /// Hands the regions and items under this view to a parity capture.
    func onParityCapture(
        regions: @escaping @Sendable ([String: CGRect]) -> Void,
        items: @escaping @Sendable ([ParityItem]) -> Void
    ) -> some View {
        onPreferenceChange(ParityRegionsKey.self, perform: regions)
            .onPreferenceChange(ParityItemsKey.self, perform: items)
    }
}
