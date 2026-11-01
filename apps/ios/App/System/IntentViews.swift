import DesignSystem
import SwiftUI
import UIKit

/// The Next upgrade snippet (board F15.2): the item icon in its quality border, the name in its
/// quality colour, `{Slot} · item level {n} · {source}` and the gain line.
struct UpgradeSnippetView: View {
    let answer: UpgradeAnswer
    let image: UIImage?
    @Environment(\.dynamicTypeSize) private var typeSize

    private static let iconSize: CGFloat = 44

    var body: some View {
        // Accessibility sizes: the icon sits above the words, so names wrap at word boundaries.
        let layout = typeSize.isAccessibilitySize
            ? AnyLayout(VStackLayout(alignment: .leading, spacing: FSSpacing.space3))
            : AnyLayout(HStackLayout(alignment: .top, spacing: FSSpacing.space3))
        layout {
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
    private static let buttonHeight: CGFloat = 48

    let text: String
    var buttons: (accept: String, deny: String)?
    var snippet: UpgradeAnswer?
    var snippetImage: UIImage?
    /// A question's choices (`Which raid?`, the spec question), as Siri lists them.
    var options: [String] = []

    var body: some View {
        VStack(alignment: .leading, spacing: FSSpacing.space3) {
            HStack(spacing: FSSpacing.space2) {
                Image(uiImage: BrandImages.seal ?? UIImage()).resizable().frame(width: 20, height: 20).clipShape(.circle).accessibilityHidden(true)
                Text(Copy.appName).font(ScaledType.meta).foregroundStyle(FSColor.muted)
            }
            // With a snippet, the snippet is the answer (board F15.2); Siri speaks the sentence.
            if let snippet {
                UpgradeSnippetView(answer: snippet, image: snippetImage).clipShape(.rect(cornerRadius: 12))
            } else {
                Text(text).font(ScaledType.body).foregroundStyle(FSColor.strong).fixedSize(horizontal: false, vertical: true)
            }
            ForEach(options, id: \.self) { option in
                Text(option).font(ScaledType.bodyStrong).foregroundStyle(FSColor.strong)
                    .fixedSize(horizontal: false, vertical: true)
                    .padding(.horizontal, FSSpacing.space4).padding(.vertical, FSSpacing.space3)
                    .frame(maxWidth: .infinity, minHeight: Self.buttonHeight, alignment: .leading)
                    .background(FSColor.line.opacity(0.6), in: .rect(cornerRadius: 14))
            }
            if let buttons {
                // Side by side while they fit; stacked at the largest sizes, never clipped.
                ViewThatFits(in: .horizontal) {
                    HStack(spacing: FSSpacing.space2) { choice(buttons.deny, accept: false); choice(buttons.accept, accept: true) }
                    VStack(spacing: FSSpacing.space2) { choice(buttons.accept, accept: true); choice(buttons.deny, accept: false) }
                }
                .font(ScaledType.bodyStrong)
            }
        }
        .padding(FSSpacing.space4)
        .background(.ultraThinMaterial, in: .rect(cornerRadius: 24))
        .environment(\.colorScheme, .dark)
    }

    /// A button that grows with its text: 48 pt at least, never clipped.
    private func choice(_ title: String, accept: Bool) -> some View {
        Text(title)
            .multilineTextAlignment(.center)
            .fixedSize(horizontal: false, vertical: true)
            .padding(.horizontal, FSSpacing.space4)
            .padding(.vertical, FSSpacing.space2)
            .frame(maxWidth: .infinity, minHeight: Self.buttonHeight)
            .background(accept ? FSColor.gold : FSColor.line, in: .capsule)
            .foregroundStyle(accept ? FSColor.bg : FSColor.strong)
    }
}
