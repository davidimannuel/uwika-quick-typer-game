-- Insert default admin user (password: admin123)
-- Password hash for "admin123"
INSERT INTO users (id, username, password_hash, role) 
VALUES (
    '00000000-0000-0000-0000-000000000001'::uuid,
    'admin',
    '$2a$10$Q8rej5GLz75dvXe0btNxkeEzBLZEf1ulcP72pXz4gz3CxckrfdkmS', -- password: admin123
    'admin'
) ON CONFLICT (username) DO NOTHING;

-- Insert themes
INSERT INTO themes (id, name, description) VALUES
('30000000-0000-0000-0000-000000000001'::uuid, 'Programming', 'Programming languages and concepts'),
('30000000-0000-0000-0000-000000000002'::uuid, 'Web Development', 'HTML, CSS, JavaScript and web technologies'),
('30000000-0000-0000-0000-000000000003'::uuid, 'Database', 'SQL and database management'),
('30000000-0000-0000-0000-000000000004'::uuid, 'Algorithms', 'Data structures and algorithms'),
('30000000-0000-0000-0000-000000000005'::uuid, 'DevOps', 'CI/CD, Docker, and cloud infrastructure')
ON CONFLICT (name) DO NOTHING;

-- Insert sample stages
INSERT INTO stages (id, name, theme_id, difficulty, is_active) VALUES
('10000000-0000-0000-0000-000000000001'::uuid, 'Java Basics', '30000000-0000-0000-0000-000000000001'::uuid, 'easy', true),
('10000000-0000-0000-0000-000000000002'::uuid, 'Python Functions', '30000000-0000-0000-0000-000000000001'::uuid, 'medium', true),
('10000000-0000-0000-0000-000000000003'::uuid, 'Advanced Algorithms', '30000000-0000-0000-0000-000000000004'::uuid, 'hard', true)
ON CONFLICT (id) DO NOTHING;

-- Insert sample phrases for Java Basics stage
INSERT INTO phrases (id, stage_id, text, sequence_number, base_multiplier) VALUES
('20000000-0000-0000-0000-000000000001'::uuid, '10000000-0000-0000-0000-000000000001'::uuid, 'public class HelloWorld', 1, 1.0),
('20000000-0000-0000-0000-000000000002'::uuid, '10000000-0000-0000-0000-000000000001'::uuid, 'System.out.println("Hello");', 2, 1.2),
('20000000-0000-0000-0000-000000000003'::uuid, '10000000-0000-0000-0000-000000000001'::uuid, 'int number = 42;', 3, 1.0)
ON CONFLICT (id) DO NOTHING;

-- Insert sample phrases for Python Functions stage
INSERT INTO phrases (id, stage_id, text, sequence_number, base_multiplier) VALUES
('20000000-0000-0000-0000-000000000004'::uuid, '10000000-0000-0000-0000-000000000002'::uuid, 'def calculate_sum(a, b):', 1, 1.5),
('20000000-0000-0000-0000-000000000005'::uuid, '10000000-0000-0000-0000-000000000002'::uuid, 'return a + b', 2, 1.3),
('20000000-0000-0000-0000-000000000006'::uuid, '10000000-0000-0000-0000-000000000002'::uuid, 'result = calculate_sum(10, 20)', 3, 1.5)
ON CONFLICT (id) DO NOTHING;

-- Insert sample phrases for Advanced Algorithms stage
INSERT INTO phrases (id, stage_id, text, sequence_number, base_multiplier) VALUES
('20000000-0000-0000-0000-000000000007'::uuid, '10000000-0000-0000-0000-000000000003'::uuid, 'function quickSort(arr) {', 1, 2.0),
('20000000-0000-0000-0000-000000000008'::uuid, '10000000-0000-0000-0000-000000000003'::uuid, 'if (arr.length <= 1) return arr;', 2, 2.0),
('20000000-0000-0000-0000-000000000009'::uuid, '10000000-0000-0000-0000-000000000003'::uuid, 'const pivot = arr[Math.floor(arr.length / 2)];', 3, 2.5)
ON CONFLICT (id) DO NOTHING;
