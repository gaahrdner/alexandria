# 11: Refining the Oracle (`11_refining_the_oracle.md`)

## Story

The Alexandria is feature-complete, but it's rough around the edges. The command-line interface is functional but clunky.  Error messages are cryptic.  The code itself has become somewhat disorganized as new features were added.

Anya Sharma gathers the team for a review. "We've built something amazing," she says, "but it's not yet ready for prime time. We need to polish the user experience and refactor the code."

Dr. Vance agrees. "The library must be intuitive and user-friendly.  Patrons shouldn't have to struggle to find what they're looking for."

Kai Tanaka, ever focused on efficiency, adds, "And we need to clean up the code. There's some technical debt we need to address.  Duplicated code, inconsistent error handling, inefficient algorithms..."

Isabelle Dubois points out, "The error messages are confusing.  If something goes wrong, the user needs to know *what* happened and *how* to fix it."

Marcus Cole emphasizes, "And we need to do a final security audit.  Make sure there are no vulnerabilities we've overlooked."

Anya nods. "It's time for a refinement phase.  We'll focus on three key areas: user interface, error handling, and code refactoring."

## Your Job

Refine the user interface, improve error handling, and refactor the code.

### Tasks

* **Objective:** Polish the UI, enhance error handling, and refactor.

* **Theory:**

  * **User Interface (UI) Design Principles:**
    * **Clarity:**  The interface should be easy to understand and use.
    * **Consistency:**  Use consistent terminology, formatting, and interaction patterns.
    * **Efficiency:**  Users should be able to accomplish tasks quickly and easily.
    * **Feedback:**  Provide clear feedback to the user about the results of their actions.
    * **Error Prevention:**  Design the interface to minimize the possibility of errors.
    * **Error Recovery:**  Make it easy for users to recover from errors.

  * **Error Handling Best Practices:**
    * **Use descriptive error messages:**  Tell the user *what* went wrong and *how* to fix it.
    * **Use custom error types:**  Define your own error types (using `errors.New` or by creating custom error structs) to provide more context about the error.
    * **Wrap errors:** Use `fmt.Errorf` with the `%w` verb to wrap errors, preserving the original error context.
    * **Handle errors appropriately:** Don't just ignore them!  Log errors, display user-friendly messages, and take corrective action where possible.
    * **Don't Panic:** Unless it's truly unrecoverable.

  * **Code Refactoring Techniques:**
    * **Extract Functions:** Break down large functions into smaller, more manageable functions.
    * **Eliminate Duplication:**  Use functions and loops to avoid repeating code.
    * **Improve Naming:** Use clear and descriptive names for variables, functions, and structs.
    * **Add Comments:**  Explain *why* the code is doing what it's doing, not just *what* it's doing.
    * **Use Constants and Enumerations:**  Define constants for magic numbers and strings. Use enumerations (using `iota`) for related constants.

* **Activities:**

    1. **`cmd/alexandria/main.go` (UI Improvements):**
        * Improve the menu layout and presentation.  Use consistent spacing and formatting.
        * Add more descriptive prompts for user input.
        * Provide clearer feedback to the user after each action (e.g., "Book added successfully!", "Patron not found.").
        * Consider using a library like `promptui` or `survey` for a more interactive CLI experience (optional).

    2. **`cmd/alexandria/main.go` (and other files) (Error Handling):**
        * Review all error handling in the code.
        * Replace generic error messages with more descriptive ones.
        * Define custom error types for common error conditions (e.g., `ErrPatronNotFound`, `ErrBookNotAvailable`).
        * Wrap errors where appropriate to provide more context.

    3. **Refactor (All Files):**
        * Identify and extract common code into reusable functions.
        * Improve variable and function names.
        * Add comments to explain complex logic.
        * Look for opportunities to simplify the code.
        * Address any TODOs or FIXME comments.

    4. **String Methods:**
        * Review all String() Methods and make sure that their output is sufficient.

    5. **Run and Test:** Compile/run. Thoroughly test all functionality, paying close attention to the user experience and error handling.
