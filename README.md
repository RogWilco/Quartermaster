# Quartermaster

A Go application using the Charm Bracelet ecosystem of TUI libraries.

## About

Quartermaster is a terminal-based application built with several libraries from the Charm Bracelet ecosystem:

- [Bubble Tea](https://github.com/charmbracelet/bubbletea): A powerful TUI framework for building terminal applications
- [Bubbles](https://github.com/charmbracelet/bubbles): A collection of common UI components for Bubble Tea
- [Lipgloss](https://github.com/charmbracelet/lipgloss): A styling library for terminal applications

## Project Structure

The project follows the [Standard Go Project Layout](https://github.com/golang-standards/project-layout):

```
.
├── bin/                  # Compiled application binaries
├── cmd/                  # Main applications for this project
│   └── qm/               # The Quartermaster CLI application
├── internal/             # Private application and library code
│   ├── app/              # Application-specific code
│   │   ├── model/        # Application model and business logic
│   │   └── ui/           # User interface code
│   └── pkg/              # Private library code
│       └── tui/          # Terminal UI utilities
```

## Getting Started

### Prerequisites

- Go 1.24 or later

### Installation

1. Clone the repository
2. Install dependencies:
   ```
   go get
   ```
3. Build the application:
   ```
   go build -o bin/qm cmd/qm/main.go
   ```

### Running the Application

To run the application after building:

```
./bin/qm
```

Or directly with Go:

```
go run cmd/qm/main.go
```

This will start an interactive application where you can:
- Use the up/down arrow keys to increment/decrement a counter
- Enter text and press Enter to add items to a list
- Press `+` to add the current text as an item
- Press `?` to toggle the help menu
- Press `q` to quit the application

## Components Used

The application demonstrates several components from the Bubbles library:

1. **TextInput**: A text input field for entering data
2. **Spinner**: A loading indicator that animates while operations are in progress
3. **Help**: A help menu that displays available keybindings
4. **Key**: A system for defining and handling keyboard shortcuts

Styling is handled with Lipgloss, which provides:
- Color management
- Text styling (bold, italic, etc.)
- Layout utilities (padding, margins, borders)

## Architecture

The application follows a Model-View-Update (MVU) architecture, which is common for Bubble Tea applications:

- **Model**: Represents the application state (in `internal/app/model`)
- **Update**: Handles messages and updates the model
- **View**: Renders the current state as a string

To learn more about the Charm Bracelet ecosystem, visit:
- [Bubble Tea](https://github.com/charmbracelet/bubbletea)
- [Bubbles](https://github.com/charmbracelet/bubbles)
- [Lipgloss](https://github.com/charmbracelet/lipgloss)
