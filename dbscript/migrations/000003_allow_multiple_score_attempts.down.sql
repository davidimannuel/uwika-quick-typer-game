-- Remove indexes
DROP INDEX IF EXISTS idx_scores_user_score;
DROP INDEX IF EXISTS idx_scores_stage_score;
DROP INDEX IF EXISTS idx_scores_user_stage_score;

-- Remove id column
ALTER TABLE scores DROP COLUMN IF EXISTS id;

-- Re-add composite primary key
ALTER TABLE scores ADD CONSTRAINT scores_pkey PRIMARY KEY (user_id, stage_id);

