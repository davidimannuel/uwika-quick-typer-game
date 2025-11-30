-- Create users table
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    username VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    role VARCHAR(50) NOT NULL DEFAULT 'user',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Create index on username for faster lookups
CREATE INDEX IF NOT EXISTS idx_users_username ON users(username);

-- Create personal_access_tokens table
CREATE TABLE IF NOT EXISTS personal_access_tokens (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    token VARCHAR(255) UNIQUE NOT NULL,
    expires_at TIMESTAMP NOT NULL,
    revoked_at TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- Create index on token for faster lookups
CREATE INDEX IF NOT EXISTS idx_tokens_token ON personal_access_tokens(token);
CREATE INDEX IF NOT EXISTS idx_tokens_user_id ON personal_access_tokens(user_id);

-- Create themes table
CREATE TABLE IF NOT EXISTS themes (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) UNIQUE NOT NULL,
    description TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Create stages table
CREATE TABLE IF NOT EXISTS stages (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    theme_id UUID NOT NULL,
    difficulty VARCHAR(50) NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT true,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (theme_id) REFERENCES themes(id) ON DELETE RESTRICT
);

-- Create index for filtering stages by theme
CREATE INDEX IF NOT EXISTS idx_stages_theme_id ON stages(theme_id);

-- Create phrases table
CREATE TABLE IF NOT EXISTS phrases (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    stage_id UUID NOT NULL,
    text TEXT NOT NULL,
    sequence_number INTEGER NOT NULL,
    base_multiplier DECIMAL(10, 2) NOT NULL DEFAULT 1.0,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (stage_id) REFERENCES stages(id) ON DELETE CASCADE,
    UNIQUE(stage_id, sequence_number)
);

-- Create index for ordering phrases
CREATE INDEX IF NOT EXISTS idx_phrases_stage_sequence ON phrases(stage_id, sequence_number);

-- Create scores table with composite unique key
CREATE TABLE IF NOT EXISTS scores (
    user_id UUID NOT NULL,
    stage_id UUID NOT NULL,
    final_score DECIMAL(10, 2) NOT NULL,
    total_time_ms INTEGER NOT NULL,
    total_errors INTEGER NOT NULL,
    completed_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, stage_id),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (stage_id) REFERENCES stages(id) ON DELETE CASCADE
);

-- Create index for leaderboard queries
CREATE INDEX IF NOT EXISTS idx_scores_stage_score ON scores(stage_id, final_score DESC);
