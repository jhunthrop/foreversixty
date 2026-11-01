import Foundation
import Observation
import Routing

/// A screen the Guides tab pushes (spec §3.1): a class page, a guide, a changelog entry.
enum GuidesRoute: Hashable {
    case classPage(String)
    /// A guide by its key, `shaman-elemental`.
    case guide(String)
    /// A changelog entry by its id.
    case entry(String)
}

/// The tab's segments; Tier lists only behind `guide_tiers`.
enum GuidesSegment: String, CaseIterable, Identifiable {
    case guides, tiers, changes

    static let storageKey = "gg.foreversixty.guides.segment"

    var id: String { rawValue }

    var title: String {
        switch self {
        case .guides: Copy.guidesSegmentGuides
        case .tiers: Copy.guidesSegmentTiers
        case .changes: Copy.guidesSegmentChanges
        }
    }

    static func available(tiersOn: Bool) -> [GuidesSegment] {
        allCases.filter { $0 != .tiers || tiersOn }
    }

    /// The stored segment when it is still offered, else Guides.
    static func resolved(_ stored: String, tiersOn: Bool) -> GuidesSegment {
        GuidesSegment(rawValue: stored).flatMap { available(tiersOn: tiersOn).contains($0) ? $0 : nil } ?? .guides
    }
}

/// The Guides tab's stack, which guide links push onto.
@MainActor
@Observable
final class GuidesRouter {
    var path: [GuidesRoute] = []

    /// The guide file's key: `{class}-{spec}` (app-v1.md W3-C1).
    static func key(className: String, spec: String) -> String { "\(className)-\(spec)" }

    static func handles(_ link: DeepLink) -> Bool {
        switch link {
        case .guides, .guide: true
        default: false
        }
    }

    /// `/app/guides` shows the tab's root; `/app/guides/{class}/{spec}` that guide over it.
    func open(_ link: DeepLink) {
        switch link {
        case .guides: path = []
        case let .guide(className, spec): path = [.guide(Self.key(className: className, spec: spec))]
        default: return
        }
    }
}
