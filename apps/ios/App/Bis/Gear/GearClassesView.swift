import DesignSystem
import SwiftUI

/// Gear signed out (flows spec 4.3, mock ios-gear-out): the nine classes with their crests; a
/// class opens its list with the picker on it, for the spec and the level range.
struct GearClassesView: View {
    /// The class slugs in the catalog's order.
    let classes: [String]
    let name: (String) -> String
    var open: (String) -> Void = { _ in }

    var body: some View {
        GearPage {
            PageHeader(title: Copy.gearOutTitle) { SmallText(Copy.gearOutMeta) }
            FSSection(Copy.gearOutClasses) {
                VStack(alignment: .leading, spacing: 0) {
                    ForEach(classes, id: \.self) { slug in
                        Button { open(slug) } label: {
                            FSRow(title: name(slug), titleColor: ClassColor.of(slug), inset: 0) {
                                ClassCrest(className: slug, size: FSSize.crest40, flat: true)
                            }
                        }
                        .buttonStyle(RowPressStyle())
                    }
                }
            }
            .padding(.top, GearListView.sectionGap)
        }
    }

    /// The catalog's classes, each once, in its order.
    static func classes(_ catalog: SpecCatalog) -> [String] {
        catalog.specs.reduce(into: [String]()) { slugs, spec in
            if !slugs.contains(spec.classSlug) { slugs.append(spec.classSlug) }
        }
    }
}
