-- Remove all seed data
-- Delete in correct order to respect foreign keys

-- First delete scores (references users and stages)
DELETE FROM scores;

-- Delete phrases (references stages)
DELETE FROM phrases;

-- Delete stages (references themes)  
DELETE FROM stages;

-- Delete themes
DELETE FROM themes;

-- Delete users
DELETE FROM users;
