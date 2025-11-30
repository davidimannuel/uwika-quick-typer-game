-- Remove unique constraint dari scores table agar bisa insert multiple attempts
-- (Primary key composite sudah di-drop, jadi sekarang hanya perlu ALTER TABLE)

-- Drop primary key constraint first
ALTER TABLE scores DROP CONSTRAINT IF EXISTS scores_pkey;

-- Add auto-increment ID as primary key
ALTER TABLE scores ADD COLUMN IF NOT EXISTS id SERIAL PRIMARY KEY;

-- Add indexes untuk query performance
CREATE INDEX IF NOT EXISTS idx_scores_user_stage_score ON scores(user_id, stage_id, final_score DESC);
CREATE INDEX IF NOT EXISTS idx_scores_stage_score ON scores(stage_id, final_score DESC);
CREATE INDEX IF NOT EXISTS idx_scores_user_score ON scores(user_id, final_score DESC);

