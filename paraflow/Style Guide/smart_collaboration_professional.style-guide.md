# Smart Collaboration Professional Style Guide

**Style Overview**:
A modern flat design with subtle depth for an AI DevOps Orchestrator platform, featuring deep teal as the primary color against very light cool-gray backgrounds. Boundaries are established through clear surface color contrast and soft shadows that provide gentle elevation, creating a professional technical atmosphere with clean organized structure and refined B2B sophistication.

## Colors
### Primary Colors
  - **primary-base**: `text-[#0D7C7C]` or `bg-[#0D7C7C]`
  - **primary-lighter**: `bg-[#1A9999]`
  - **primary-darker**: `text-[#096363]` or `bg-[#096363]`

### Background Colors

#### Structural Backgrounds

Choose based on layout type:

**For Vertical Layout** (Top Header + Optional Side Panels):
- **bg-nav-primary**: `bg-[hsla(200, 15%, 96%, 1)]` - Top header
- **bg-nav-secondary**: `bg-[hsla(200, 15%, 97.5%, 1)]` - Inner Left sidebar (if present)
- **bg-page**: `bg-[hsla(200, 8%, 99%, 1)]` - Page background (bg of Main Content area)

**For Horizontal Layout** (Side Navigation + Optional Top Bar):
- **bg-nav-primary**: `bg-[hsla(200, 8%, 99%, 1)] border-r border-[#E1E6EA]` - Left main sidebar
- **bg-nav-secondary**: `bg-[hsla(200, 8%, 99%, 1)] border-b border-[#E1E6EA]` - Inner Top header (if present)
- **bg-page**: `bg-[hsla(200, 8%, 99%, 1)]` - Page background (bg of Main Content area)

#### Container Backgrounds
For main content area. Adjust values when used on navigation backgrounds to ensure sufficient contrast.
- **bg-container-primary**: `bg-white`
- **bg-container-secondary**: `bg-[hsla(200, 20%, 98%, 1)]`
- **bg-container-inset**: `bg-[hsla(200, 15%, 95%, 1)]`
- **bg-container-inset-strong**: `bg-[hsla(200, 20%, 92%, 1)]`

### Text Colors
- **color-text-primary**: `text-[#1A1F24]`
- **color-text-secondary**: `text-[#4A5259]`
- **color-text-tertiary**: `text-[#6B7680]`
- **color-text-quaternary**: `text-[#9BA3AB]`
- **color-text-on-dark-primary**: `text-white/95` - Text on dark backgrounds and primary-base, accent-dark color surfaces
- **color-text-on-dark-secondary**: `text-white/70` - Text on dark backgrounds and primary-base, accent-dark color surfaces
- **color-text-link**: `text-[#0D7C7C]` - Links, text-only buttons without backgrounds, and clickable text in tables

### Functional Colors
Use **sparingly** to maintain a minimalist and neutral overall style. Used for the surfaces of text-only cards, simple cards, buttons, and tags.
  - **color-success-default**: #0F8B5D
  - **color-success-light**: #E8F7F0 - tag/label bg
  - **color-error-default**: #D93025 - alert banner bg
  - **color-error-light**: #FDECEA - tag/label bg
  - **color-warning-default**: #F59D0D - tag/label bg
  - **color-warning-light**: #FEF6E8 - tag/label bg, alert banner bg
  - **color-function-default**: #3367D6
  - **color-function-light**: #E8F0FE - tag/label bg

### Accent Colors
  - A secondary palette for occasional highlights and categorization. **Avoid overuse** to protect brand identity. Use **sparingly**.
  - **accent-blue-gray**: `text-[#5B7C92]` or `bg-[#5B7C92]`
  - **accent-cyan-muted**: `text-[#4D9FA6]` or `bg-[#4D9FA6]`

### Data Visualization Charts
For data visualization charts only.
  - Standard data colors: #0D7C7C, #1A9999, #5B7C92, #4D9FA6, #6B7680, #9BA3AB
  - Important data can use small amounts of: #3367D6, #0F8B5D, #F59D0D, #D93025

## Typography
- **Font Stack**:
  - **font-family-base**: `-apple-system, BlinkMacSystemFont, "Segoe UI"` — For regular UI copy

- **Font Size & Weight**:

  - **Caption**: `text-sm font-normal`
  - **Body**: `text-base font-normal`
  - **Body Emphasized**: `text-base font-semibold`
  - **Card Title / Subtitle**: `text-lg font-semibold`
  - **Page Title**: `text-2xl font-semibold`
  - **Headline**: `text-3xl font-semibold`

- **Line Height**: 1.5

## Border Radius
  - **Small**: 8px — Elements inside cards (e.g., photos)
  - **Medium**: 12px
  - **Large**: 16px — Cards
  - **Full**: full — Toggles, avatars, small tags, inputs, etc.

## Layout & Spacing
  - **Tight**: 8px - For closely related small internal elements, such as icons and text within buttons
  - **Compact**: 12px - For small gaps between small containers, such as a line of tags
  - **Standard**: 20px - For gaps between medium containers like list items
  - **Relaxed**: 32px - For gaps between large containers and sections
  - **Section**: 48px - For major section divisions


## Create Boundaries (contrast of surface color, borders, shadows)
Primarily relying on clear surface color contrast to create visual hierarchy, complemented by soft shadows for gentle elevation on key interactive elements and containers.

### Borders
  - **Default**: 1px solid #E1E6EA. Used for inputs, dividing navigation sections, and subtle container separation. `border border-[#E1E6EA]`
  - **Stronger**: 1px solid #CBD3DA. Used for active or focused states. `border border-[#CBD3DA]`

### Dividers
  - Use `border-t` or `border-b` `border-[#E1E6EA]` for section separation within containers and navigation areas.

### Shadows & Effects
  - **Case 1 (no shadow)**: For flat elements like tags, text-only buttons, and inline controls
  - **Case 2 (subtle shadow)**: `shadow-[0_1px_3px_rgba(0,0,0,0.08)]` - For smaller cards, dropdowns, and floating elements
  - **Case 3 (moderate shadow)**: `shadow-[0_2px_8px_rgba(0,0,0,0.10)]` - For primary cards, panels, and modal dialogs
  - **Case 4 (pronounced shadow)**: `shadow-[0_4px_12px_rgba(0,0,0,0.12)]` - For elevated modals, popovers, and prominent interactive elements

## Visual Emphasis for Containers
When containers (tags, cards, list items, rows) need visual emphasis to indicate priority, status, or category, use the following techniques:

| Technique | Implementation Notes | Best For | Avoid |
|-----------|---------------------|----------|-------|
| Background Tint | Slightly darker/lighter color or reduce transparency of backgrounds | Gentle, common approach for moderate emphasis needs | Heavy colors on large areas (e.g., red background for entire large cards) |
| Border Highlight | Use thin border with opacity for subtlety | Active/selected states, form validation | - |
| Glow/Shadow Effect | Keep shadow subtle with low opacity | Premium aesthetics, hover states | Flat/minimal designs |
| Status Tag/Label | Add colored tag/label inside container | Larger containers | - |
| Side Accent Bar | **Left edge only**, for **non-rounded containers** | Small non-rounded list items (e.g., side nav tabs), small non-rounded cards (e.g., task cards) | Large cards, wide list items, rounded containers |

## Assets
### Image

- For normal `<img>`: object-cover
- For `<img>` with:
  - Slight overlay: object-cover brightness-85
  - Heavy overlay: object-cover brightness-50

### Icon

- Use Lucide icons from Iconify.
- To ensure an aesthetic layout, each icon should be centered in a square container, typically without a background, matching the icon's size.
- Use Tailwind font size to control icon size
- Example:
  ```html
  <div class="flex items-center justify-center bg-transparent w-5 h-5">
  <iconify-icon icon="lucide:flag" class="text-base"></iconify-icon>
  </div>
  ```

### Third-Party Brand Logos:
   - Use Brand Icons from Iconify.
   - Logo Example:
     Monochrome Logo: `<iconify-icon icon="simple-icons:x"></iconify-icon>`
     Colored Logo: `<iconify-icon icon="logos:google-icon"></iconify-icon>`

### User's Own Logo:
- To protect copyright, do **NOT** use real product logos as a logo for a new product, individual user, or other company products.
- **Icon-based**:
  - **Graphic**: Use a simple, relevant icon (e.g., a `calendar` icon for a scheduling app, a `heart` icon for a dating app).

## Page Layout - Web

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
<body class="w-[1440px] min-h-[900px] font-[-apple-system,BlinkMacSystemFont,'Segoe UI'] leading-[1.5]">

  <!-- Header (Primary Nav): Fixed height -->
  <header class="w-full">
    <!-- Header content -->
  </header>

  <!-- Content Container: Must include 'flex' class -->
  <div class="w-full flex min-h-[900px]">
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
<body class="w-[1440px] min-h-[900px] flex font-[-apple-system,BlinkMacSystemFont,'Segoe UI'] leading-[1.5]">

<!-- Left Sidebar (Primary Nav): Use its ml to control left page margin -->
  <aside class="flex-shrink-0 min-w-fit">
  </aside>

  <!-- Content Container-->
  <div class="flex-1 overflow-x-hidden flex flex-col min-h-[900px]">

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

- **Button**:
  - Example 1 (text button):
    - button: flex items-center gap-2 px-4 py-2 rounded-full transition hover:opacity-80
      - span(button copy): whitespace-nowrap
  - Example 2 (icon button):
    - button: flex items-center justify-center w-10 h-10 rounded-full transition hover:opacity-80
      - icon

- **Tag Group (Filter Tags)**
  - container(scrollable): flex gap-2 overflow-x-auto [&::-webkit-scrollbar]:hidden
    - label (Tag item):
      - input: type="radio" name="tag1" class="sr-only peer" checked
      - div: bg-[hsla(200,15%,95%,1)] text-[#4A5259] px-4 py-2 rounded-full peer-checked:bg-[#0D7C7C] peer-checked:text-white/95 hover:opacity-80 transition whitespace-nowrap

### Data Entry
- **Progress bars/Slider**: h-2
- **Checkbox**
  - label: flex items-center gap-2
    - input: type="checkbox" class="sr-only peer"
    - div: w-5 h-5 bg-[hsla(200,15%,95%,1)] rounded-md flex items-center justify-center peer-checked:bg-[#0D7C7C] text-transparent peer-checked:text-white/95
      - svg(Checkmark): stroke="currentColor" stroke-width="3"
    - span(text)
- **Radio button**
  - label: flex items-center gap-2
    - input: type="radio" name="option1" class="sr-only peer"
    - div: w-5 h-5 bg-[hsla(200,15%,95%,1)] rounded-full flex items-center justify-center peer-checked:bg-[#0D7C7C] text-transparent peer-checked:text-white/95
      - svg(dot indicator): fill="currentColor"
    - span(text)
- **Switch/Toggle**
  - label: flex items-center gap-3
    - div: relative
      - input: type="checkbox" class="sr-only peer"
      - div(Toggle track): w-12 h-6 bg-[hsla(200,15%,95%,1)] rounded-full peer-checked:bg-[#0D7C7C] transition
      - div(Toggle thumb): absolute top-0.5 left-0.5 w-5 h-5 bg-white rounded-full shadow-sm peer-checked:translate-x-6 transition
    - span(text)

- **Select/Dropdown**
  - Select container: flex items-center gap-2 px-4 py-2 bg-white border border-[#E1E6EA] rounded-full
    - text
    - Dropdown icon(square container): flex items-center justify-center bg-transparent w-4 h-4
      - icon


### Container
- **Navigation Menu - horizontal**
    - Navigation with sections/grouping:
        - Nav Container: flex items-center justify-between w-full px-8
        - Left Section: flex items-center gap-8
          - Menu Item: flex items-center gap-2
        - Right Section: flex items-center gap-4
          - Menu Item: flex items-center gap-2
          - Notification (if applicable): relative flex items-center justify-center w-10 h-10
            - notification-icon: w-5 h-5
            - badge (if has unread): absolute -top-1 -right-1 w-5 h-5 rounded-full flex items-center justify-center
              - badge-count: text-xs
          - Avatar(if applicable): flex items-center gap-2
            - avatar-image: w-9 h-9 rounded-full
            - dropdown-icon (if applicable): w-4 h-4

- **Card**
    - Example 1 (Vertical card with image and text):
        - Card: bg-white shadow-[0_2px_8px_rgba(0,0,0,0.10)] rounded-2xl flex flex-col p-5 gap-4
        - Image: rounded-lg w-full object-cover
        - Text area: flex flex-col gap-3
          - card-title: text-lg font-semibold
          - card-subtitle: text-sm font-normal text-[#6B7680]
    - Example 2 (Horizontal card with image and text):
        - Card: bg-white shadow-[0_2px_8px_rgba(0,0,0,0.10)] rounded-2xl flex gap-5 p-5
        - Image: rounded-lg h-full object-cover
        - Text area: flex flex-col gap-3
          - card-title: text-lg font-semibold
          - card-subtitle: text-sm font-normal text-[#6B7680]
    - Example 3 (Image-focused card: no background or padding. Avoid rounded corners on container as they cause only top corners of image to be rounded):
        - Card: flex flex-col gap-4
        - Image: rounded-2xl w-full object-cover
        - Text area: flex flex-col gap-3
          - card-title: text-lg font-semibold
          - card-subtitle: text-sm font-normal text-[#6B7680]
    - Example 4 (text-only cards, simple cards, such as Upgrade Card, Activity Summary Cards):
        - Card: bg-white shadow-[0_2px_8px_rgba(0,0,0,0.10)] rounded-2xl flex flex-col p-5 gap-4

## Additional Notes

This style guide creates a professional technical atmosphere suitable for an AI DevOps platform. The deep teal primary color conveys trust and technical sophistication, while the very light cool-gray backgrounds maintain clarity and reduce eye strain during extended work sessions. The combination of clear surface color contrast and soft shadows provides functional hierarchy without visual heaviness, supporting focused professional workflows.

<colors_extraction>
#0D7C7C
#1A9999
#096363
#E8F0F7
#ECECF0
#FBFCFC
#FFFFFF
#F6F8FA
#F0F2F5
#E8EBED
#1A1F24
#4A5259
#6B7680
#9BA3AB
#FFFFFFF2
#FFFFFFB3
#0F8B5D
#E8F7F0
#D93025
#FDECEA
#F59D0D
#FEF6E8
#3367D6
#E8F0FE
#5B7C92
#4D9FA6
#E1E6EA
#CBD3DA
</colors_extraction>
