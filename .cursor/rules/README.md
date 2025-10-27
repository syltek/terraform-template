# Cursor rule guidelines

## Core principles

- `Cohesion`: Group related rules together.
- `Granularity`: Keep rules concise (under 500 lines) and focused on a single, well-defined concern.
- `Consistency`: Apply the same rule structure every time.
- `Discoverability`: Names should be intuitive and indicate purpose.
- `Scalability`: The structure should accommodate growth.
- `Specificity`: Avoid vague guidance. Provide concrete examples and references.

> [!IMPORTANT]
>
> - Use Cursor's AI to create new rules. Don't always write them manually.
> - Create rules that you can reuse when you find yourself repeating prompts in chat.
> - NEVER include verbose explanations or redundant context that increases AI token overhead

## File naming convention

- Use kebab-case for filenames.
- Always use .mdc extension
- Make names a concise description of the rule's purpose.

## File organization

General rule files are placed in `PROJECT_ROOT/.cursor/rules/`

Sub-directories structure:
    - `global/`: Project-wide rules and guidelines
    - `terraform/`: Terraform-specific rules
    - ...

## Examples of bad file names

- `my_rule.mdc` (Too generic)
- `big_long_description_of_something_very_specific_that_takes_up_too_much_space.mdc` (Too long, hard to scan)
- `auth.mdc` (Not descriptive enough; could be auth-api-keys.mdc, auth-middleware.mdc, etc.)
- `RuleAboutReact.mdc` (Inconsistent casing, redundant "RuleAbout")

## File content structure

### YAML front matter

Start the file with a front matter YAML at the top. This is crucial for Cursor to understand when to apply each rule.

```yaml
---
description: This rule explains how to create new .mdc project rule files for the Cursor agent.
globs: '.cursor/rules/**/*.mdc'
alwaysApply: true
---
```

- `description`: A human-readable summary of the rule's purpose. It will be used
by Cursor to dynamically select the rule when needed
- `globs`: Crucial for defining when the rule applies. For example:
  - `"**"`: Applies to all files/contexts in the project (e.g., global guidelines).
  - `"**/*.md"`: Only applies to Markdown files (useful for documentation rules).
  - `"**/*.swift"`: Only applies to Swift files.
  - You can also specify multiple globs: `["**/*.swift", "**/*.xib"]`
- `alwaysApply`: If true, the rule will always be applied. If false, the rule will
only be applied when the `globs` condition is met or when the rule is explicitly
invoked.

More info in [cursor's official documentation](https://docs.cursor.com/context/rules#adding-a-new-setting-in-cursor).

### Rule content

- Use Markdown for clarity.
- Keep rules as short as possible.
- Keep examples as short as possible to clearly convey the positive or negative example.
- Start with a clear heading (e.g., `# Rule Name`).
- Provide context, explanations, and examples.
- Be explicit with instructions for the AI, especially in a dedicated
"Instructions for AI" section at the end. Use imperative verbs.
- Avoid conversational filler.
