import DesignSystem
import SwiftUI

/// Gear signed out (flows spec 4.3, mock ios-gear-out): the nine classes with their crests; a
/// class opens its list with the picker on it, for the spec and the level range.
struct GearClassesView: View {
    /// The class slugs in the game's order.
    let classes: [String]
    let name: (String) -> String
    var open: (String) -> Void = { _ in }

    /// `.ap section { margin-top: 28px }`.
    static let sectionGap: CGFloat = 28

    var body: some View {
        ScrollView {
            VStack(alignment: .leading, spacing: 0) {
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
                .padding(.top, Self.sectionGap)
            }
            .padding(.top, FSSpace.space2)
            .padding(.horizontal, FSSpace.gutterPhone)
            .padding(.bottom, FSSpace.space6)
            .frame(maxWidth: .infinity, alignment: .leading)
            .parityContent()
        }
        .background(FSColor.bg)
    }
}
