import DesignSystem
import SwiftUI
import UIKit

/// The Next upgrade snippet (board F15.2): the item icon in its quality border, the name in its
/// quality colour, `{Slot} · item level {n} · {source}` and the gain line.
struct UpgradeSnippetView: View {
    let answer: UpgradeAnswer
    let image: UIImage?

    private static let iconSize: CGFloat = 44

    var body: some View {
        HStack(alignment: .top, spacing: FSSpacing.space3) {
            WidgetItemIcon(icon: answer.icon ?? "", quality: answer.quality ?? 1, size: Self.iconSize, icons: WidgetIcons { _ in image })
            VStack(alignment: .leading, spacing: 4) {
                Text(answer.itemName)
                    .font(ScaledType.bodyStrong)
                    .foregroundStyle(ItemQuality(value: answer.quality).text)
                Text(answer.meta).font(ScaledType.meta).foregroundStyle(FSColor.muted)
                Text(answer.gain).font(ScaledType.body).foregroundStyle(FSColor.text)
            }
            .fixedSize(horizontal: false, vertical: true)
            Spacer(minLength: 0)
        }
        .padding(FSSpacing.space4)
        .frame(maxWidth: .infinity, alignment: .leading)
        .background(FSColor.raised)
        .accessibilityElement(children: .combine)
    }
}

/// A spoken answer as Siri draws it under the app's name: the review captures' stand-in for the
/// system's own sheet (F15.1, F15.2), and the confirmation with its two buttons.
struct IntentAnswerCard: View {
    let text: String
    var buttons: (accept: String, deny: String)?
    var snippet: UpgradeAnswer?

    var body: some View {
        VStack(alignment: .leading, spacing: FSSpacing.space3) {
            HStack(spacing: FSSpacing.space2) {
                Image(uiImage: BrandImages.seal ?? UIImage()).resizable().frame(width: 20, height: 20).clipShape(.circle).accessibilityHidden(true)
                Text(Copy.appName).font(ScaledType.meta).foregroundStyle(FSColor.muted)
            }
            Text(text).font(ScaledType.body).foregroundStyle(FSColor.strong).fixedSize(horizontal: false, vertical: true)
            if let snippet { UpgradeSnippetView(answer: snippet, image: nil).clipShape(.rect(cornerRadius: 12)) }
            if let buttons {
                HStack(spacing: FSSpacing.space2) {
                    Text(buttons.deny).frame(maxWidth: .infinity, minHeight: 44).background(FSColor.line, in: .capsule)
                    Text(buttons.accept).frame(maxWidth: .infinity, minHeight: 44).background(FSColor.gold, in: .capsule).foregroundStyle(FSColor.bg)
                }
                .font(ScaledType.bodyStrong)
            }
        }
        .padding(FSSpacing.space4)
        .background(.ultraThinMaterial, in: .rect(cornerRadius: 24))
        .environment(\.colorScheme, .dark)
    }
}
