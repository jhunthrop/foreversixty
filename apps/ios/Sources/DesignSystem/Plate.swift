import SwiftUI

/// `ui/Plate`: the one raised surface on a screen (`raised` ground, 1 pt `line`, 6 radius, no
/// shadow). Reports `plate`; the border is inside the box, as the web's border-box.
public struct Plate<Content: View>: View {
    private let content: Content

    public init(@ViewBuilder content: () -> Content) {
        self.content = content()
    }

    public var body: some View {
        VStack(alignment: .leading, spacing: 0) { content }
            .padding(FSStroke.hairline)
            .frame(maxWidth: .infinity, alignment: .leading)
            .background(FSColor.raised, in: .rect(cornerRadius: FSRadius.plate))
            .clipShape(.rect(cornerRadius: FSRadius.plate))
            .overlay(RoundedRectangle(cornerRadius: FSRadius.plate).strokeBorder(FSColor.line, lineWidth: FSStroke.hairline))
            .parityRegion(.plate)
    }
}

/// The padding of a message inside a Plate (the web's `.plate .empty`): 18 over, 16 at the sides
/// and under.
enum PlateMessageInset {
    static let top: CGFloat = 18
    static let side: CGFloat = FSSpace.space4
    static let bottom: CGFloat = FSSpace.space4
}

/// `ui/EmptyState`: one sentence and its actions (at most one button and a link) in a Plate.
public struct EmptyState<Actions: View>: View {
    private let text: String
    private let actions: Actions

    public init(_ text: String, @ViewBuilder actions: () -> Actions) {
        self.text = text
        self.actions = actions()
    }

    public var body: some View {
        VStack(alignment: .leading, spacing: FSSpace.space3) {
            Text(text).type(FSType.body).foregroundStyle(FSColor.muted).fixedSize(horizontal: false, vertical: true)
            if Actions.self != EmptyView.self {
                HStack(spacing: FSSpace.space4) { actions }
            }
        }
        .padding(.top, PlateMessageInset.top)
        .padding(.horizontal, PlateMessageInset.side)
        .padding(.bottom, PlateMessageInset.bottom)
        .frame(maxWidth: .infinity, alignment: .leading)
    }
}

public extension EmptyState where Actions == EmptyView {
    init(_ text: String) {
        self.init(text, actions: { EmptyView() })
    }
}

/// `ui/LoadError`: `... did not load.` and a small Try again, beside the sentence when `inline`
/// (a report's error) or under it (a list's).
public struct LoadError: View {
    private let text: String
    private let retryTitle: String
    private let inline: Bool
    private let retry: () -> Void

    public init(_ text: String, retryTitle: String, inline: Bool = false, retry: @escaping () -> Void) {
        self.text = text
        self.retryTitle = retryTitle
        self.inline = inline
        self.retry = retry
    }

    public var body: some View {
        let layout = inline
            ? AnyLayout(HStackLayout(alignment: .center, spacing: FSSpace.space1))
            : AnyLayout(VStackLayout(alignment: .leading, spacing: FSSpace.space3))
        layout {
            Text(text).type(FSType.body).foregroundStyle(FSColor.muted).fixedSize(horizontal: false, vertical: true)
            Button(retryTitle, action: retry).buttonStyle(FSButtonStyle(.secondary, fullWidth: false, small: true))
        }
        .padding(.top, PlateMessageInset.top)
        .padding(.horizontal, PlateMessageInset.side)
        .padding(.bottom, PlateMessageInset.bottom)
        .frame(maxWidth: .infinity, alignment: .leading)
        .accessibilityElement(children: .contain)
    }
}

/// `ui/Notice`: a standing sentence in the flow of the page (the offline Notice, neutral tone),
/// 14 over and 4 under, no box.
public struct Notice: View {
    private let text: String

    public init(_ text: String) {
        self.text = text
    }

    public var body: some View {
        Text(text)
            .type(FSType.body)
            .foregroundStyle(FSColor.text)
            .fixedSize(horizontal: false, vertical: true)
            .padding(.top, 14)
            .padding(.bottom, FSSpace.space1)
            .frame(maxWidth: .infinity, alignment: .leading)
    }
}

/// `ui/Switch` with its label beside it: the system toggle tinted gold, hugging its label.
public struct LabelledSwitch: View {
    private let title: String
    @Binding private var isOn: Bool

    public init(_ title: String, isOn: Binding<Bool>) {
        self.title = title
        _isOn = isOn
    }

    public var body: some View {
        Toggle(isOn: $isOn) {
            Text(title).type(FSType.small).foregroundStyle(FSColor.text)
        }
        .tint(FSColor.gold)
        .fixedSize()
        .frame(minHeight: FSTarget.min)
    }
}

/// `ui/ActionLink`: gold 600 with `›`, padded to the 44 pt target. Not a button in the inventory.
public struct ActionLink: View {
    private let title: String
    private let action: () -> Void

    public init(_ title: String, action: @escaping () -> Void) {
        self.title = title
        self.action = action
    }

    public var body: some View {
        Button(action: action) {
            Text("\(title) \u{203A}")
                .font(.custom(FSFont.bodySemibold, size: 14, relativeTo: .subheadline))
                .foregroundStyle(FSColor.gold)
                .frame(minHeight: FSTarget.min)
                .contentShape(.rect)
        }
        .buttonStyle(.plain)
    }
}
