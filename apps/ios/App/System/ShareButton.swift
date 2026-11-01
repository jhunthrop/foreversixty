import DesignSystem
import SwiftUI

/// The system share sheet with the page link and its message (spec §5.7 Share). Every share in
/// the app goes through it, so each carries the table's text.
struct ShareButton<Label: View>: View {
    let item: ShareItem
    @ViewBuilder let label: () -> Label

    var body: some View {
        ShareLink(item: item.url, subject: Text(item.text), message: Text(item.text), label: label)
    }
}

extension ShareButton where Label == AnyView {
    /// The bar's share glyph with its spoken name.
    init(_ item: ShareItem, accessibility: String) {
        self.init(item: item) {
            AnyView(Image(systemName: "square.and.arrow.up").accessibilityLabel(accessibility))
        }
    }
}
