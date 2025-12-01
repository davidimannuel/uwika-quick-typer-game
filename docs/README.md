# Quick Typer Game - Architecture Documentation

Dokumentasi arsitektur sistem Quick Typer Game menggunakan PlantUML.

## 📐 Diagram yang Tersedia

### High-Level Architecture (Recommended untuk Overview)

#### 1. **c4-level1-context.puml** - C4 Level 1: System Context ⭐
- Menunjukkan sistem Quick Typer dan aktornya (Player, Admin)
- External systems (Ngrok)
- Big picture dari sistem
- **Best for:** Presentasi ke stakeholder, overview sistem

#### 2. **c4-level2-container.puml** - C4 Level 2: Container Diagram ⭐
- Mobile App, Admin Web, API Server, Database
- Hubungan antar container
- Technology stack per container
- Key features dan responsibilities
- **Best for:** Understanding system components

**👍 Gunakan ini untuk:**
- Presentasi ke stakeholder
- Onboarding developer baru
- Overview sistem secara keseluruhan

### Detailed Architecture (Untuk Developer)

#### 3. **layered-architecture-detail.puml** - Layered Architecture Detail
Arsitektur berlapis (layered architecture) dari sistem dengan detail teknis:
- **Presentation Layer**: HTTP Router, Middleware, Handlers, DTO
- **Application Layer**: Services (orchestration)
- **Domain Layer**: Models, Domain Services, Repository Interfaces
- **Infrastructure Layer**: Repository Implementations, Database Connection

**👍 Gunakan ini untuk:**
- Understanding code structure
- Developer yang mau kontribusi
- Technical documentation

### Database & Data Model

#### 4. **database-schema.puml** - Database Schema
Entity Relationship Diagram dari database:
- Tables: users, personal_access_tokens, themes, stages, phrases, scores
- Relationships dan foreign keys
- Indexes dan constraints

### Behavior & Flow Diagrams

#### 5. **game-flow-sequence.puml** - Game Flow Sequence
Sequence diagram untuk flow utama game:
1. Authentication
2. Get Active Stages
3. Get Stage Details with Phrases
4. Play Game
5. Submit Score
6. View Leaderboard

#### 6. **score-calculation.puml** - Score Calculation Flow
Activity diagram menjelaskan detail perhitungan score:
- Raw metrics calculation (accuracy, WPM)
- Validation
- Score formula dengan bonuses dan penalties
- Domain service logic

### Deployment

#### 7. **deployment-local.puml** - Local Development Deployment
Deployment diagram untuk local development dengan ngrok:
- Docker container untuk PostgreSQL
- API running di localhost
- Ngrok tunnel untuk mobile access

#### 8. **deployment-production.puml** - Production Deployment (Docker)
Deployment diagram untuk production:
- Full Docker Compose setup
- API, Admin Web, PostgreSQL containers
- Production-ready configuration
- Scaling considerations

## 🎯 Quick Start - Mana Diagram yang Harus Dilihat?

### Untuk Non-Technical (PM, Stakeholder)
1. **c4-level1-context.puml** - System big picture
2. **c4-level2-container.puml** - System components
3. **game-flow-sequence.puml** - Understand user flow

### Untuk Developer Baru
1. **c4-level1-context.puml** - System overview
2. **c4-level2-container.puml** - Container overview
3. **layered-architecture-detail.puml** - Code structure
4. **database-schema.puml** - Data model
5. **game-flow-sequence.puml** - Main flow

### Untuk Developer yang Mau Kontribusi
1. **layered-architecture-detail.puml** - Understand layers
2. **component-diagram.puml** - Component dependencies (jika ada)
3. **score-calculation.puml** - Business logic detail

### Untuk DevOps / Deployment
1. **deployment-local.puml** - Local development setup (with ngrok)
2. **deployment-production.puml** - Production setup (Docker only)
3. **database-schema.puml** - Database requirements

## 🛠️ Cara Melihat Diagram

### Online (Recommended)
1. Buka [PlantUML Online Server](http://www.plantuml.com/plantuml/uml/)
2. Copy-paste isi file `.puml`
3. Lihat hasilnya

### VS Code
1. Install extension: **PlantUML** by jebbs
2. Buka file `.puml`
3. Tekan `Alt+D` atau `Cmd+D` untuk preview

### CLI (Local)
```bash
# Install PlantUML
brew install plantuml

# Generate PNG
plantuml c4-level1-context.puml

# Generate SVG
plantuml -tsvg c4-level2-container.puml
```

### Generate All Diagrams
```bash
cd docs

# Generate semua diagram ke PNG
plantuml -tpng *.puml

# Atau ke SVG (better quality)
plantuml -tsvg *.puml

# Generate specific diagram
plantuml -tpng c4-level1-context.puml
plantuml -tpng c4-level2-container.puml
```

## 📚 Arsitektur Highlights

### Clean Architecture / Layered Architecture
- **Separation of Concerns**: Setiap layer punya tanggung jawab jelas
- **Dependency Rule**: Dependencies mengalir ke dalam (ke domain)
- **Testability**: Domain layer bisa di-test tanpa infrastructure

### Domain-Driven Design (DDD)
- **Domain Models**: Pure Go structs tanpa JSON tags
- **Domain Services**: `ScoreCalculator` - pure business logic
- **Repository Pattern**: Interface di domain, implementation di infrastructure

### Key Design Decisions

1. **Score Calculator sebagai Domain Service**
   - Pure business logic
   - Testable tanpa dependencies
   - Reusable calculation rules

2. **Multiple Score Attempts**
   - Migration 000003 mengubah scores table
   - Dari composite PK → serial ID
   - Allows multiple submissions per user/stage

3. **Repository Pattern**
   - Interfaces di `domain/repositories`
   - Implementations di `infrastructure/persistence/postgres`
   - Easy to mock untuk testing

4. **Layered Architecture**
   - Clear separation: Presentation → Application → Domain → Infrastructure
   - Domain tidak depend pada infrastructure
   - Infrastructure depend pada domain interfaces

## 🔄 Data Flow

```
Mobile App → Handler → Service → Repository → Database
                ↓
         Domain Service (Score Calculator)
```

## 🎯 Scoring System

**Formula:**
```
Base Score = (accuracy / 100) × typing_speed × 10 × base_multiplier

Bonuses:
- Accuracy ≥ 95%: +50%
- Accuracy ≥ 90%: +25%
- Accuracy ≥ 85%: +10%
- Speed ≥ 80 WPM: +30%
- Speed ≥ 60 WPM: +15%

Penalties:
- Error rate: -(error_rate × 50%)
- Time > 180s: -20%
- Time > 120s: -10%

Final Score = Base + Bonuses - Penalties
```

## 📝 Notes

- Semua diagram mengikuti konvensi PlantUML
- Untuk diagram yang lebih kompleks, bisa di-split jadi beberapa view
- Update diagram seiring evolusi sistem

