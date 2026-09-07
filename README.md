# 💣 SaperGoLang

**SaperGoLang** is a classic Minesweeper game written in **Go** using [raylib-go](https://github.com/gen2brain/raylib-go).

The project focuses on clean game logic, separation of concerns, unit testing, and CI automation.

<p align="center">
  <img src="assets/gameplay.gif" alt="SaperGoLang gameplay" width="600">
</p>

## ✨ Features

* 💣 Classic Minesweeper gameplay
* 🎚️ Three difficulty levels
* 🚩 Flag / unflag cells
* 🔢 Mines-around counter
* 🖱️ Mouse controls
* ⚡ Automatic opening of empty areas
* 🔗 Chording — open neighboring cells when the number of flags matches the number of surrounding mines
* ⏱️ Game timer
* 🏆 Win detection
* 💥 Game-over detection
* 🔄 Restart game
* 🧪 Unit tests for core game logic
* 🔁 GitHub Actions CI

## 🎮 Difficulty Levels

| Level           |   Board | Mines |
| --------------- | ------: | ----: |
| 🟢 Beginner     |   9 × 9 |    10 |
| 🟡 Intermediate | 16 × 16 |    40 |
| 🔴 Expert       | 30 × 30 |    99 |

## 🕹️ Controls

| Action                 | Control                    |
| ---------------------- | -------------------------- |
| Open cell              | Left mouse button          |
| Mark / unmark cell     | Right mouse button         |
| Open neighboring cells | Left + right mouse buttons |
| Restart                | Restart button             |

## 🛠️ Tech Stack

* **Go 1.26**
* **raylib-go** — graphics, window management and input
* **raygui** — UI components
* **Go testing** — unit tests
* **GitHub Actions** — CI

## 🚀 Getting Started

### Requirements

* Go 1.26+
* C compiler
* Native dependencies required by raylib

### Clone

```bash
git clone https://github.com/Avigdor-Kolonimus/sapergolang.git
cd sapergolang
```

### Install dependencies

```bash
go mod download
```

### Run

```bash
go run ./cmd/sapergame
```

## 🧪 Testing

Run all tests:

```bash
go test ./...
```

Run tests with verbose output:

```bash
go test -v ./...
```

Run static analysis:

```bash
go vet ./...
```

## 🏗️ Project Structure

```text
.
├── cmd/
│   └── sapergame/
│       └── saper.go          # Application entry point
│
├── game/
│   ├── field.go              # Core Minesweeper logic
│   ├── field_test.go         # Unit tests
│   └── ...
│
├── .github/
│   └── workflows/
│       └── unit_test.yml     # GitHub Actions workflow
│
├── go.mod
├── go.sum
└── LICENSE
```

## 🧩 Architecture

The application is split into two main parts:

```text
cmd/sapergame
       │
       ▼
   GameState
       │
       ├── Game lifecycle
       ├── Field management
       ├── Cell interactions
       ├── Win / lose detection
       └── Timer
       │
       ▼
    raylib / raygui
       │
       ├── Rendering
       ├── Mouse input
       └── UI
```

### Game Logic

The core state is represented by `GameState`.

Each cell contains its current state:

```go
type point struct {
    hasMine     bool
    opened      bool
    marked      bool
    minesAround int
}
```

The game logic is responsible for:

* revealing cells;
* detecting mines;
* calculating neighboring cells;
* opening empty areas recursively;
* marking cells;
* implementing chording;
* detecting win / loss conditions.

UI and rendering are kept separate from the core game operations where practical.

## 🧠 Interesting Implementation Details

### Recursive Empty-Cell Reveal

When a cell has no neighboring mines, its neighbors are automatically revealed.

The implementation uses a shared neighbor iterator:

```go
func (g *GameState) doForNeighbours(
    x, y int,
    do func(x, y int),
)
```

This keeps boundary checks in one place and avoids duplicating the same eight-neighbor logic across the game.

### Chording

The game supports the classic Minesweeper interaction where an already opened cell can reveal its neighbors when:

```text
number of marked neighbors == minesAround
```

This allows several safe cells to be opened with a single action.

### Game State

The game explicitly tracks terminal states:

```text
Playing
   │
   ├── Mine revealed ──► Game Over
   │
   └── All safe cells ──► Game Won
```

This prevents further interaction after the game has finished.

## 🧪 Testing Strategy

The core game logic is covered with unit tests without requiring a graphical environment.

Tests cover cases such as:

* revealing a mine;
* revealing an already opened cell;
* revealing a marked cell;
* interaction after game over;
* interaction after winning;
* recursive neighbor revealing;
* chording with an incorrect number of flags;
* chording with the correct number of flags;
* neighbor detection for corners, edges and center cells.

Example:

```go
func TestGameState_RevealTile_Mine(t *testing.T) {
    g := &GameState{
        rows: 1,
        cols: 1,
        field: [][]point{
            {{hasMine: true}},
        },
    }

    g.revealTile(0, 0)

    if !g.gameOver {
        t.Error("gameOver should be true")
    }
}
```

## 🔄 Continuous Integration

Every push and pull request runs automated checks using GitHub Actions.

The CI pipeline verifies:

```text
┌───────────────┐
│ Checkout code │
└───────┬───────┘
        ▼
┌────────────────────┐
│ Install raylib deps│
└────────┬───────────┘
         ▼
┌────────────────┐
│ Download modules│
└───────┬────────┘
        ▼
┌──────────────┐
│ go build ./...│
└───────┬──────┘
        ▼
┌──────────────┐
│ go test ./... │
└───────┬──────┘
        ▼
┌──────────────┐
│ go vet ./...  │
└──────────────┘
```

This ensures that changes are compiled, tested, and statically checked automatically.

## 📌 What I Practiced

This project was built as a practical exercise in:

* Go project structure
* Struct-based state management
* Recursive algorithms
* 2D grid traversal
* Event handling
* Separation of game logic and rendering
* Unit testing
* GitHub Actions
* Working with a native graphics library from Go

## 📄 License

This project is licensed under the MIT License.

See [LICENSE](LICENSE) for details.

---

<p align="center">
  Built with ❤️ and Go
</p>
