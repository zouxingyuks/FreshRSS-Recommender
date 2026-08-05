# Agent Working Guidelines

|IMPORTANT: Prefer retrieval-led reasoning over pre-training-led reasoning
|Language Policy:Chinese for Q&A|English for code/comments/docs/technical artifacts
|Scope:repo-root
|Conventions:before coding, inspect current repo structure and analogous implementations before designing or editing
|Conventions:assumption discipline=do not silently choose among materially different interpretations; after retrieval, ask one precise question if ambiguity affects API/behavior/data model
|Conventions:simplicity first=minimum code that satisfies requested behavior; no speculative features/configuration/abstractions
|Conventions:surgical changes=every changed line must trace to the request; do not refactor/reformat adjacent code
|Conventions:goal-driven execution=convert work into verifiable success criteria; bugfix requires reproduction/validation path; feature requires behavior check
|Conventions:reference-first=before new feature/refactor, find similar implementations or explicitly state none exist
|Delivery:MVP development=first make the narrowest end-to-end path work, then add core logic, then add boundaries/error handling/optimization
|Testing:add regression/behavior tests only when a real requirement or bugfix provides a stable acceptance scenario; do not invent tests for formality
|Design Principles:preserve the spirit of OCP/SRP/ISP/DIP/LoD/LSP/ETC/DRY/YAGNI as pressure checks, without importing project-specific anchors
|Crash Early:use guard clauses and fail fast; do not silently swallow invalid state or errors
|Complexity Control:prefer guard clauses and early returns; avoid deep nested conditional logic; nesting deeper than 3 levels is a refactoring signal
|Self-Doc:self-explaining code > nearby comment > docs; comments should explain intent or non-obvious constraints
|Terminology:grep existing terms before adding new ones; same concept should use same name across code/docs/issues
|Business Model:code should follow the real domain model, not temporary UI wording or one-off requirement branches
|YAGNI:solve the current workflow need before adding future-facing knobs, abstractions, background workers, or integration surfaces
