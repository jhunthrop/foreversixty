import DesignSystem
import Models
import SwiftUI

/// The class an Activity row is about, drawn as its crest (one crest language, never a square spec
/// icon; lane M4a): the spec icon the server sends for a guildmate (`guildsim.passed`), else the spec
/// a tier move names at its start (`Fury Warrior moved to S`). Nil for every other row: it keeps its
/// game icon.
enum ActivityCrest {
    /// The kinds whose title opens with the spec they are about.
    static let namedSpecKinds: Set<String> = ["tier.moved"]

    static func className(_ row: ActivityRow, specs: SpecCatalog) -> String? {
        // The server names the other player's class when the row is about one.
        if let crest = row.crest, !crest.isEmpty { return crest }
        if let spec = specs.specs.first(where: { $0.icon == row.icon }) { return spec.classSlug }
        guard namedSpecKinds.contains(row.kind) else { return nil }
        let title = row.title.lowercased()
        return specs.specs.first { title.hasPrefix("\($0.name) \($0.classSlug) ".lowercased()) }?.classSlug
    }
}

/// An Activity row's 40 pt lead: the crest for a row about a class, else the row's game icon.
struct ActivityLead: View {
    let row: ActivityRow
    let build: String
    var specs: SpecCatalog = .main

    var body: some View {
        if let className = ActivityCrest.className(row, specs: specs) {
            ClassCrest(className: className, size: FSSize.icon40, flat: true)
        } else {
            ItemIconView(icon: row.icon, quality: nil, build: build, size: FSSize.icon40).parityItem(.image)
        }
    }
}

extension SpecCatalog {
    /// The bundled catalog, read once.
    static let main = bundled()
}
