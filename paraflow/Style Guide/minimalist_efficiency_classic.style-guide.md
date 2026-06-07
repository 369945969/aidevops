# Minimalist Efficiency Classic Style Guide

**Style Overview**:
A clean, outlined minimalist **light theme** centered on calm navy blue, using pure white backgrounds with fine-line borders to create precise structural boundaries, complemented by warm amber accents for strategic focus points—delivering distraction-free clarity ideal for technical workflows.

## Colors
### Primary Colors
  - **primary-base**: `text-[#2C5F8D]` or `bg-[#2C5F8D]`
  - **primary-lighter**: `bg-[#4A7BA7]`
  - **primary-darker**: `text-[#1E4463]` or `bg-[#1E4463]`

### Background Colors

#### Structural Backgrounds

Choose based on layout type:

**For Vertical Layout** (Top Header + Optional Side Panels):
- **bg-nav-primary**: `bg-white border-b border-[#E8EBED]` - Top header
- **bg-nav-secondary**: `bg-white border-r border-[#E8EBED]` - Inner Left sidebar (if present)
- **bg-page**: `bg-white` - Page background (bg of Main Content area)

**For Horizontal Layout** (Side Navigation + Optional Top Bar):
- **bg-nav-primary**: `bg-white border-r border-[#E8EBED]` - Left main sidebar
- **bg-nav-secondary**: `bg-white border-b border-[#E8EBED]` - Inner Top header (if present)
- **bg-page**: `bg-white` - Page background (bg of Main Content area)

#### Container Backgrounds
For main content area. Adjust values when used on navigation backgrounds to ensure sufficient contrast.
- **bg-container-primary**: `bg-white`
- **bg-container-secondary**: `bg-[#FAFBFC]`
- **bg-container-inset**: `bg-[#F5F7F9]`
- **bg-container-inset-strong**: `bg-[#E8EBED]`

### Text Colors
- **color-text-primary**: `text-[#1A1F24]`
- **color-text-secondary**: `text-[#4A5568]`
- **color-text-tertiary**: `text-[#6B7280]`
- **color-text-quaternary**: `text-[#9CA3AF]`
- **color-text-on-dark-primary**: `text-white/95` - Text on dark backgrounds and primary-base color surfaces
- **color-text-on-dark-secondary**: `text-white/75` - Text on dark backgrounds and primary-base color surfaces
- **color-text-link**: `text-[#2C5F8D]` - Links, text-only buttons without backgrounds, and clickable text in tables

### Functional Colors
Use **sparingly** to maintain a minimalist and neutral overall style. Used for the surfaces of text-only cards, simple cards, buttons, and tags.
  - **color-success-default**: #10B981 - button bg, icon color
  - **color-success-light**: #D1FAE5 - tag/label bg
  - **color-error-default**: #EF4444 - button bg, icon color
  - **color-error-light**: #FEE2E2 - tag/label bg, alert banner bg
  - **color-warning-default**: #F59E0B - button bg, icon color
  - **color-warning-light**: #FEF3C7 - tag/label bg, alert banner bg
  - **color-function-default**: #2C5F8D - button bg
  - **color-function-light**: #DBEAFE - tag/label bg

### Accent Colors
  - A secondary palette for occasional highlights and categorization. **Avoid overuse** to protect brand identity. Use **sparingly**.
  - **accent-amber**: `text-[#D97706]` or `bg-[#D97706]`
  - **accent-amber-light**: `bg-[#FEF3C7]`

### Data Visualization Charts
For data visualization charts only.
  - Standard data colors: #2C5F8D, #4A7BA7, #6B93BE, #8AAED5, #D97706, #F59E0B
  - Status indicators: #10B981 (success), #F59E0B (warning), #EF4444 (error)

## Typography
- **Font Stack**:
  - **font-family-base**: `-apple-system, BlinkMacSystemFont, "Segoe UI", "Helvetica Neue", Arial, sans-serif` — For regular UI copy

- **Font Size & Weight**:
  - **Caption**: `text-sm font-normal`
  - **Body**: `text-base font-normal`
  - **Body Emphasized**: `text-base font-semibold`
  - **Card Title / Subtitle**: `text-lg font-semibold`
  - **Page Title**: `text-2xl font-semibold`
  - **Headline**: `text-3xl font-semibold`

- **Line Height**: 1.5

## Border Radius
  - **Small**: 4px — Elements inside cards (e.g., inputs, small buttons)
  - **Medium**: 8px — Cards, containers
  - **Large**: 12px — Large containers, modals
  - **Full**: full — Toggles, avatars, small tags

## Layout & Spacing
  - **Tight**: 8px - For closely related small internal elements, such as icons and text within buttons
  - **Compact**: 12px - For small gaps between small containers, such as a line of tags
  - **Standard**: 16px - For gaps between medium containers like list items
  - **Relaxed**: 24px - For gaps between large containers and sections
  - **Section**: 32px - For major section divisions

## Create Boundaries (contrast of surface color, borders, shadows)
Pure outlined design - surfaces match background with fine-line borders for clear structural definition, no shadows for distraction-free focus.

### Borders
  - **Default**: 1px solid #E8EBED. Used for cards, containers, inputs. `border border-[#E8EBED]`
  - **Stronger**: 1px solid #D1D5DB. Used for active or focused states. `border border-[#D1D5DB]`
  - **Emphasis**: 1px solid #2C5F8D. Used for selected states or primary emphasis. `border border-[#2C5F8D]`

### Dividers
  - **Horizontal**: `border-t border-[#E8EBED]`
  - **Vertical**: `border-l border-[#E8EBED]`

### Shadows & Effects
  - **Case 1**: No shadow - maintain pure outlined aesthetic for maximum clarity

## Visual Emphasis for Containers
When containers (tags, cards, list items, rows) need visual emphasis to indicate priority, status, or category, use the following techniques:

| Technique | Implementation Notes | Best For | Avoid |
|-----------|---------------------|----------|-------|
| Background Tint | Slightly darker/lighter color or reduce transparency of backgrounds | Gentle, common approach for moderate emphasis needs | Heavy colors on large areas (e.g., red background for entire large cards) |
| Border Highlight | Use thin border with opacity for subtlety | Active/selected states, form validation | - |
| Status Tag/Label | Add colored tag/label inside container | Larger containers | - |
| Side Accent Bar | **Left edge only**, for **non-rounded containers** | Small non-rounded list items (e.g., side nav tabs), small non-rounded cards (e.g., task cards) | Large cards, wide list items, rounded containers |

## Assets
### Image

- For normal `<img>`: object-cover
- For `<img>` with:
  - Slight overlay: object-cover brightness-90
  - Heavy overlay: object-cover brightness-75

### Icon

- Use Lucide icons from Iconify.
- To ensure an aesthetic layout, each icon should be centered in a square container, typically without a background, matching the icon's size.
- Use Tailwind font size to control icon size
- Example:
  ```html
  <div class="flex items-center justify-center bg-transparent w-5 h-5">
  <iconify-icon icon="lucide:activity" class="text-base"></iconify-icon>
  </div>
  ```

### Third-Party Brand Logos:
   - Use Brand Icons from Iconify.
   - Logo Example:
     Monochrome Logo: `<iconify-icon icon="simple-icons:docker"></iconify-icon>`
     Colored Logo: `<iconify-icon icon="logos:docker-icon"></iconify-icon>`

### User's Own Logo:
- To protect copyright, do **NOT** use real product logos as a logo for a new product, individual user, or other company products.
- **Icon-based**:
  - **Graphic**: Use a simple, relevant icon (e.g., a `workflow` icon for an orchestration platform, a `server` icon for infrastructure tools).

## Page Layout - Web (*EXTREMELY* important)
### Determine Layout Type
- Choose between Vertical or Horizontal layout based on whether the primary navigation is a full-width top header or a full-height sidebar (left/right).
- User requirements typically indicate the layout preference. If unclear, consider:
  - Marketing/content sites typically use Vertical Layout.
  - Functional/dashboard sites can use either, depending on visual style. Sidebars accommodate more complex navigation than top bars. For complex navigation needs with a preference for minimal chrome (Vertical Layout adds an extra fixed header), choose Horizontal Layout (omits the fixed top header).
- Vertical Layout Diagram:
┌──────────────────────────────────────────────────────┐
│  Header (Primary Nav)                                │
├──────────┬──────────────────────────────┬────────────┤
│Left      │ Sub-header (Tertiary Nav)    │ Right      │
│Sidebar   │ (optional)                   │ Sidebar    │
│(Secondary├──────────────────────────────┤ (Utility   │
│Nav)      │ Main Content                 │ Panel)     │
│(optional)│                              │ (optional) │
│          │                              │            │
└──────────┴──────────────────────────────┴────────────┘
- Horizontal Layout Diagram:
┌──────────┬──────────────────────────────┬───────────┐
│          │ Header (Secondary Nav)       │           │
│ Left     │ (optional)                   │ Right     │
│ Sidebar  ├──────────────────────────────┤ Sidebar   │
│ (Primary │ Main Content                 │ (Utility  │
│ Nav)     │                              │ Panel)    │
│          │                              │ (optional)│
│          │                              │           │
└──────────┴──────────────────────────────┴───────────┘
### Detailed Layout Code
**Vertical Layout**
```html
<!-- Body: Adjust width (w-[1440px]) based on target screen size -->
<body class="w-[1440px] min-h-[700px] font-[-apple-system,BlinkMacSystemFont,'Segoe_UI','Helvetica_Neue',Arial,sans-serif] leading-[1.5]">

  <!-- Header (Primary Nav): Fixed height -->
  <header class="w-full">
    <!-- Header content -->
  </header>

  <!-- Content Container: Must include 'flex' class -->
  <div class="w-full flex min-h-[700px]">
    <!-- Left Sidebar (Secondary Nav) (Optional): Remove if not needed. If Left Sidebar exists, use its ml to control left page margin -->
    <aside class="flex-shrink-0 min-w-fit">

    </aside>

    <!-- Main Content Area:
     Use Main Content Area's horizontal padding (px) to control distance from main content to sidebars or page edges.
     For pages without sidebars (like Marketing Pages, simple content pages such as help centers, privacy policies) use larger values (px-30 to px-80), for pages with sidebars (Functional/Dashboard Pages, complex content pages with multi-level navigation like knowledge base articles) use moderate values (px-8 to px-16) -->
    <main class="flex-1 overflow-x-hidden flex flex-col">
    <!--  Main Content -->

    </main>

    <!-- Right Sidebar (Utility Panel) (Optional): Remove if not needed. If Right Sidebar exists, use its mr to control right page margin -->
    <aside class="flex-shrink-0 min-w-fit">
    </aside>

  </div>
</body>
```

**Horizontal Layout**

```html
<!-- Body: Adjust width (w-[1440px]) based on target screen size. Must include 'flex' class -->
<body class="w-[1440px] min-h-[700px] flex font-[-apple-system,BlinkMacSystemFont,'Segoe_UI','Helvetica_Neue',Arial,sans-serif] leading-[1.5]">

<!-- Left Sidebar (Primary Nav): Use its ml to control left page margin -->
  <aside class="flex-shrink-0 min-w-fit">
  </aside>

  <!-- Content Container-->
  <div class="flex-1 overflow-x-hidden flex flex-col min-h-[700px]">

    <!-- Header (Secondary Nav) (Optional): Remove if not needed. If Header exists, use its mx to control distance to left/right sidebars or page margins -->
    <header class="w-full">
    </header>

    <!-- Main Content Area: Use Main Content Area's pl to control distance from main content to left sidebar. Use pr to control distance to right sidebar/right page edge -->
    <main class="w-full">
    </main>


  </div>

  <!-- Right Sidebar (Utility Panel) (Optional): Remove if not needed. If Right Sidebar exists, use its mr to control right page margin -->
  <aside class="flex-shrink-0 min-w-fit">
  </aside>

</body>
```

## Tailwind Component Examples (Key attributes)
**Important Note**: Use utility classes directly. Do NOT create custom CSS classes or add styles in <style> tags for the following components

### Basic

- **Button**: (Note: Use flex and items-center for the container)
  - Example 1 (Primary solid button):
    - button: flex items-center bg-[#2C5F8D] text-white/95 hover:bg-[#1E4463] transition rounded-lg px-4 py-2 gap-2
      - icon (optional)
      - span(button copy): whitespace-nowrap
  - Example 2 (Secondary outlined button):
    - button: flex items-center bg-white border border-[#E8EBED] text-[#2C5F8D] hover:border-[#2C5F8D] transition rounded-lg px-4 py-2 gap-2
      - icon (optional)
      - span(button copy): whitespace-nowrap
  - Example 3 (text button):
    - button: flex items-center text-[#2C5F8D] hover:opacity-70 transition gap-2
      - icon (optional)
      - span(button copy): whitespace-nowrap
  - Example 4 (icon button):
    - button: flex items-center justify-center w-9 h-9 bg-white border border-[#E8EBED] hover:border-[#2C5F8D] transition rounded-lg
      - icon: w-5 h-5

- **Tag Group (Filter Tags)** (Note: `overflow-x-auto` and `whitespace-nowrap` are required)
  - container(scrollable): flex overflow-x-auto gap-2 [&::-webkit-scrollbar]:hidden
    - label (Tag item 1):
      - input: type="radio" name="tag1" class="sr-only peer" checked
      - div: bg-white border border-[#E8EBED] text-[#4A5568] peer-checked:border-[#2C5F8D] peer-checked:bg-[#2C5F8D] peer-checked:text-white/95 hover:border-[#D1D5DB] transition whitespace-nowrap rounded-full px-3 py-1.5

### Data Entry
- **Progress bars/Slider**: h-2 bg-[#E8EBED] rounded-full
  - progress fill: bg-[#2C5F8D] h-2 rounded-full
- **Checkbox**
  - label: flex items-center gap-2
    - input: type="checkbox" class="sr-only peer"
    - div: w-5 h-5 bg-white border-2 border-[#E8EBED] rounded flex items-center justify-center peer-checked:border-[#2C5F8D] peer-checked:bg-[#2C5F8D] transition
      - svg(Checkmark): stroke="currentColor" stroke-width="3" class="w-3 h-3 text-transparent peer-checked:text-white"
    - span(text): text-base
- **Radio button**
  - label: flex items-center gap-2
    - input: type="radio" name="option1" class="sr-only peer"
    - div: w-5 h-5 bg-white border-2 border-[#E8EBED] rounded-full flex items-center justify-center peer-checked:border-[#2C5F8D] transition
      - svg(dot indicator): fill="currentColor" class="w-2.5 h-2.5 text-transparent peer-checked:text-[#2C5F8D]"
    - span(text): text-base
- **Switch/Toggle**
  - label: flex items-center gap-3
    - div: relative
      - input: type="checkbox" class="sr-only peer"
      - div(Toggle track): w-11 h-6 bg-[#E8EBED] peer-checked:bg-[#2C5F8D] transition rounded-full
      - div(Toggle thumb): absolute top-0.5 left-0.5 w-5 h-5 bg-white rounded-full peer-checked:translate-x-5 transition
    - span(text): text-base

- **Input Field**
  - container: flex flex-col gap-1.5
    - label: text-sm font-medium text-[#1A1F24]
    - input: w-full px-3 py-2 bg-white border border-[#E8EBED] rounded-lg focus:outline-none focus:border-[#2C5F8D] transition text-base

- **Select/Dropdown**
  - Select container: flex items-center justify-between bg-white border border-[#E8EBED] rounded-lg px-3 py-2 hover:border-[#D1D5DB] transition cursor-pointer
    - text: text-base
    - Dropdown icon(square container): flex items-center justify-center bg-transparent w-5 h-5
      - icon: text-[#6B7280]

### Container
- **Navigation Menu - horizontal**
    - Navigation with sections/grouping:
        - Nav Container: flex items-center justify-between w-full px-6 py-4
        - Left Section: flex items-center gap-8
          - Menu Item: flex items-center gap-2 text-[#4A5568] hover:text-[#2C5F8D] transition
        - Right Section: flex items-center gap-4
          - Menu Item: flex items-center gap-2 text-[#4A5568] hover:text-[#2C5F8D] transition
          - Notification (if applicable): relative flex items-center justify-center w-10 h-10
            - notification-icon: w-5 h-5 text-[#4A5568]
            - badge (if has unread): absolute -top-1 -right-1 w-5 h-5 bg-[#EF4444] rounded-full flex items-center justify-center text-white text-xs font-semibold
              - badge-count:
          - Avatar(if applicable): flex items-center gap-2
            - avatar-image: w-9 h-9 rounded-full border border-[#E8EBED]
            - dropdown-icon (if applicable): w-4 h-4 text-[#6B7280]

- **Navigation Menu - vertical (sidebar)**
    - Menu Item: flex items-center gap-3 px-4 py-3 text-[#4A5568] hover:bg-[#F5F7F9] hover:text-[#2C5F8D] transition rounded-lg
      - icon: w-5 h-5
      - text: text-base
    - Menu Item (active): flex items-center gap-3 px-4 py-3 bg-[#DBEAFE] text-[#2C5F8D] border-l-2 border-[#2C5F8D] rounded-lg
      - icon: w-5 h-5
      - text: text-base font-semibold

- **Card**
    - Example 1 (Vertical card with border):
        - Card: bg-white border border-[#E8EBED] rounded-lg flex flex-col p-5 gap-4 hover:border-[#D1D5DB] transition
        - Image (optional): rounded-md w-full object-cover
        - Text area: flex flex-col gap-2
          - card-title: text-lg font-semibold text-[#1A1F24]
          - card-subtitle: text-sm text-[#6B7280]
    - Example 2 (Horizontal card with border):
        - Card: bg-white border border-[#E8EBED] rounded-lg flex gap-4 p-5 hover:border-[#D1D5DB] transition
        - Image (optional): rounded-md h-24 w-24 object-cover
        - Text area: flex flex-col gap-2 flex-1
          - card-title: text-lg font-semibold text-[#1A1F24]
          - card-subtitle: text-sm text-[#6B7280]
    - Example 3 (List item card):
        - Card: bg-white border border-[#E8EBED] rounded-lg flex items-center justify-between px-5 py-4 hover:border-[#D1D5DB] transition
        - Left content: flex items-center gap-4
          - icon/avatar: w-10 h-10
          - text area: flex flex-col gap-1
            - title: text-base font-semibold text-[#1A1F24]
            - subtitle: text-sm text-[#6B7280]
        - Right content: flex items-center gap-3
          - action button/status indicator

- **Table**
    - Container: w-full bg-white border border-[#E8EBED] rounded-lg overflow-hidden
    - Table: w-full
      - Header row: bg-[#FAFBFC] border-b border-[#E8EBED]
        - th: px-5 py-3 text-left text-sm font-semibold text-[#1A1F24]
      - Body row: border-b border-[#E8EBED] hover:bg-[#FAFBFC] transition
        - td: px-5 py-4 text-base text-[#4A5568]

## Additional Notes

This style guide emphasizes structural clarity through outlined design, ensuring maximum information density without visual clutter. The navy blue and warm amber combination provides professional authority while maintaining approachability for technical teams. All interactive elements use subtle border transitions rather than shadows, keeping the interface clean and focused on workflow efficiency.

<colors_extraction>
#2C5F8D
#4A7BA7
#1E4463
#FFFFFF
#E8EBED
#FAFBFC
#F5F7F9
#1A1F24
#4A5568
#6B7280
#9CA3AF
#FFFFFFF2
#FFFFFFBF
#10B981
#D1FAE5
#EF4444
#FEE2E2
#F59E0B
#FEF3C7
#DBEAFE
#D97706
#6B93BE
#8AAED5
#D1D5DB
</colors_extraction>
