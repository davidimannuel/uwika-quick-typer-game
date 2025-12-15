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
('30000000-0000-0000-0000-000000000005'::uuid, 'DevOps', 'CI/CD, Docker, and cloud infrastructure'),
('30000000-0000-0000-0000-000000000006'::uuid, 'General Knowledge', 'Common phrases and sentences'),
('30000000-0000-0000-0000-000000000007'::uuid, 'Business', 'Business and professional writing'),
('30000000-0000-0000-0000-000000000008'::uuid, 'Science', 'Scientific terms and concepts'),
('30000000-0000-0000-0000-000000000009'::uuid, 'Literature', 'Famous quotes and literary excerpts')
ON CONFLICT (name) DO NOTHING;

-- ==========================================
-- PROGRAMMING THEME STAGES
-- ==========================================

INSERT INTO stages (id, name, theme_id, difficulty, is_active) VALUES
('10000000-0000-0000-0000-000000000001'::uuid, 'Java Basics', '30000000-0000-0000-0000-000000000001'::uuid, 'easy', true),
('10000000-0000-0000-0000-000000000002'::uuid, 'Python Functions', '30000000-0000-0000-0000-000000000001'::uuid, 'medium', true),
('10000000-0000-0000-0000-000000000003'::uuid, 'Go Concurrency', '30000000-0000-0000-0000-000000000001'::uuid, 'hard', true)
ON CONFLICT (id) DO NOTHING;

-- Java Basics phrases
INSERT INTO phrases (id, stage_id, text, sequence_number, base_multiplier) VALUES
('20000000-0000-0000-0000-000000000001'::uuid, '10000000-0000-0000-0000-000000000001'::uuid, 'public class HelloWorld', 1, 1.0),
('20000000-0000-0000-0000-000000000002'::uuid, '10000000-0000-0000-0000-000000000001'::uuid, 'System.out.println("Hello");', 2, 1.2),
('20000000-0000-0000-0000-000000000003'::uuid, '10000000-0000-0000-0000-000000000001'::uuid, 'int number = 42;', 3, 1.0),
('20000000-0000-0000-0000-000000000004'::uuid, '10000000-0000-0000-0000-000000000001'::uuid, 'String name = "Java";', 4, 1.1),
('20000000-0000-0000-0000-000000000005'::uuid, '10000000-0000-0000-0000-000000000001'::uuid, 'boolean isActive = true;', 5, 1.0)
ON CONFLICT (id) DO NOTHING;

-- Python Functions phrases
INSERT INTO phrases (id, stage_id, text, sequence_number, base_multiplier) VALUES
('20000000-0000-0000-0000-000000000006'::uuid, '10000000-0000-0000-0000-000000000002'::uuid, 'def calculate_sum(a, b):', 1, 1.5),
('20000000-0000-0000-0000-000000000007'::uuid, '10000000-0000-0000-0000-000000000002'::uuid, 'return a + b', 2, 1.3),
('20000000-0000-0000-0000-000000000008'::uuid, '10000000-0000-0000-0000-000000000002'::uuid, 'result = calculate_sum(10, 20)', 3, 1.5),
('20000000-0000-0000-0000-000000000009'::uuid, '10000000-0000-0000-0000-000000000002'::uuid, 'for item in collection:', 4, 1.4),
('20000000-0000-0000-0000-000000000010'::uuid, '10000000-0000-0000-0000-000000000002'::uuid, 'if condition is True:', 5, 1.3)
ON CONFLICT (id) DO NOTHING;

-- Go Concurrency phrases
INSERT INTO phrases (id, stage_id, text, sequence_number, base_multiplier) VALUES
('20000000-0000-0000-0000-000000000011'::uuid, '10000000-0000-0000-0000-000000000003'::uuid, 'go func() { process() }()', 1, 2.0),
('20000000-0000-0000-0000-000000000012'::uuid, '10000000-0000-0000-0000-000000000003'::uuid, 'channel := make(chan int)', 2, 1.8),
('20000000-0000-0000-0000-000000000013'::uuid, '10000000-0000-0000-0000-000000000003'::uuid, 'select { case msg := <-ch: }', 3, 2.2),
('20000000-0000-0000-0000-000000000014'::uuid, '10000000-0000-0000-0000-000000000003'::uuid, 'var wg sync.WaitGroup', 4, 1.9),
('20000000-0000-0000-0000-000000000015'::uuid, '10000000-0000-0000-0000-000000000003'::uuid, 'defer wg.Done()', 5, 1.7)
ON CONFLICT (id) DO NOTHING;

-- ==========================================
-- WEB DEVELOPMENT THEME STAGES
-- ==========================================

INSERT INTO stages (id, name, theme_id, difficulty, is_active) VALUES
('10000000-0000-0000-0000-000000000004'::uuid, 'HTML Fundamentals', '30000000-0000-0000-0000-000000000002'::uuid, 'easy', true),
('10000000-0000-0000-0000-000000000005'::uuid, 'CSS Styling', '30000000-0000-0000-0000-000000000002'::uuid, 'medium', true),
('10000000-0000-0000-0000-000000000006'::uuid, 'JavaScript DOM', '30000000-0000-0000-0000-000000000002'::uuid, 'hard', true)
ON CONFLICT (id) DO NOTHING;

-- HTML Fundamentals phrases
INSERT INTO phrases (id, stage_id, text, sequence_number, base_multiplier) VALUES
('20000000-0000-0000-0000-000000000016'::uuid, '10000000-0000-0000-0000-000000000004'::uuid, '<!DOCTYPE html>', 1, 1.0),
('20000000-0000-0000-0000-000000000017'::uuid, '10000000-0000-0000-0000-000000000004'::uuid, '<html lang="en">', 2, 1.1),
('20000000-0000-0000-0000-000000000018'::uuid, '10000000-0000-0000-0000-000000000004'::uuid, '<head><title>My Page</title></head>', 3, 1.3),
('20000000-0000-0000-0000-000000000019'::uuid, '10000000-0000-0000-0000-000000000004'::uuid, '<div class="container">', 4, 1.2),
('20000000-0000-0000-0000-000000000020'::uuid, '10000000-0000-0000-0000-000000000004'::uuid, '<a href="https://example.com">Link</a>', 5, 1.4)
ON CONFLICT (id) DO NOTHING;

-- CSS Styling phrases
INSERT INTO phrases (id, stage_id, text, sequence_number, base_multiplier) VALUES
('20000000-0000-0000-0000-000000000021'::uuid, '10000000-0000-0000-0000-000000000005'::uuid, 'display: flex; justify-content: center;', 1, 1.5),
('20000000-0000-0000-0000-000000000022'::uuid, '10000000-0000-0000-0000-000000000005'::uuid, 'background-color: #3498db;', 2, 1.4),
('20000000-0000-0000-0000-000000000023'::uuid, '10000000-0000-0000-0000-000000000005'::uuid, 'border-radius: 8px; padding: 16px;', 3, 1.5),
('20000000-0000-0000-0000-000000000024'::uuid, '10000000-0000-0000-0000-000000000005'::uuid, 'box-shadow: 0 4px 6px rgba(0,0,0,0.1);', 4, 1.7),
('20000000-0000-0000-0000-000000000025'::uuid, '10000000-0000-0000-0000-000000000005'::uuid, '@media (max-width: 768px) { }', 5, 1.6)
ON CONFLICT (id) DO NOTHING;

-- JavaScript DOM phrases
INSERT INTO phrases (id, stage_id, text, sequence_number, base_multiplier) VALUES
('20000000-0000-0000-0000-000000000026'::uuid, '10000000-0000-0000-0000-000000000006'::uuid, 'document.getElementById("app")', 1, 1.8),
('20000000-0000-0000-0000-000000000027'::uuid, '10000000-0000-0000-0000-000000000006'::uuid, 'element.addEventListener("click", handler)', 2, 2.0),
('20000000-0000-0000-0000-000000000028'::uuid, '10000000-0000-0000-0000-000000000006'::uuid, 'const items = document.querySelectorAll(".item")', 3, 2.2),
('20000000-0000-0000-0000-000000000029'::uuid, '10000000-0000-0000-0000-000000000006'::uuid, 'fetch("/api/data").then(res => res.json())', 4, 2.3),
('20000000-0000-0000-0000-000000000030'::uuid, '10000000-0000-0000-0000-000000000006'::uuid, 'localStorage.setItem("token", value)', 5, 2.0)
ON CONFLICT (id) DO NOTHING;

-- ==========================================
-- DATABASE THEME STAGES
-- ==========================================

INSERT INTO stages (id, name, theme_id, difficulty, is_active) VALUES
('10000000-0000-0000-0000-000000000007'::uuid, 'SQL Basics', '30000000-0000-0000-0000-000000000003'::uuid, 'easy', true),
('10000000-0000-0000-0000-000000000008'::uuid, 'SQL Joins', '30000000-0000-0000-0000-000000000003'::uuid, 'medium', true),
('10000000-0000-0000-0000-000000000009'::uuid, 'Advanced Queries', '30000000-0000-0000-0000-000000000003'::uuid, 'hard', true)
ON CONFLICT (id) DO NOTHING;

-- SQL Basics phrases
INSERT INTO phrases (id, stage_id, text, sequence_number, base_multiplier) VALUES
('20000000-0000-0000-0000-000000000031'::uuid, '10000000-0000-0000-0000-000000000007'::uuid, 'SELECT * FROM users;', 1, 1.0),
('20000000-0000-0000-0000-000000000032'::uuid, '10000000-0000-0000-0000-000000000007'::uuid, 'INSERT INTO products (name, price) VALUES', 2, 1.3),
('20000000-0000-0000-0000-000000000033'::uuid, '10000000-0000-0000-0000-000000000007'::uuid, 'UPDATE users SET status = "active"', 3, 1.2),
('20000000-0000-0000-0000-000000000034'::uuid, '10000000-0000-0000-0000-000000000007'::uuid, 'DELETE FROM orders WHERE id = 1', 4, 1.2),
('20000000-0000-0000-0000-000000000035'::uuid, '10000000-0000-0000-0000-000000000007'::uuid, 'CREATE TABLE customers (id SERIAL PRIMARY KEY)', 5, 1.5)
ON CONFLICT (id) DO NOTHING;

-- SQL Joins phrases
INSERT INTO phrases (id, stage_id, text, sequence_number, base_multiplier) VALUES
('20000000-0000-0000-0000-000000000036'::uuid, '10000000-0000-0000-0000-000000000008'::uuid, 'SELECT * FROM orders INNER JOIN users ON', 1, 1.6),
('20000000-0000-0000-0000-000000000037'::uuid, '10000000-0000-0000-0000-000000000008'::uuid, 'LEFT JOIN products ON orders.product_id = products.id', 2, 1.8),
('20000000-0000-0000-0000-000000000038'::uuid, '10000000-0000-0000-0000-000000000008'::uuid, 'GROUP BY category HAVING COUNT(*) > 5', 3, 1.7),
('20000000-0000-0000-0000-000000000039'::uuid, '10000000-0000-0000-0000-000000000008'::uuid, 'ORDER BY created_at DESC LIMIT 10', 4, 1.5),
('20000000-0000-0000-0000-000000000040'::uuid, '10000000-0000-0000-0000-000000000008'::uuid, 'WHERE status IN ("pending", "processing")', 5, 1.6)
ON CONFLICT (id) DO NOTHING;

-- Advanced Queries phrases
INSERT INTO phrases (id, stage_id, text, sequence_number, base_multiplier) VALUES
('20000000-0000-0000-0000-000000000041'::uuid, '10000000-0000-0000-0000-000000000009'::uuid, 'WITH ranked AS (SELECT *, ROW_NUMBER() OVER)', 1, 2.2),
('20000000-0000-0000-0000-000000000042'::uuid, '10000000-0000-0000-0000-000000000009'::uuid, 'PARTITION BY category ORDER BY sales DESC', 2, 2.0),
('20000000-0000-0000-0000-000000000043'::uuid, '10000000-0000-0000-0000-000000000009'::uuid, 'COALESCE(discount, 0) AS final_discount', 3, 1.9),
('20000000-0000-0000-0000-000000000044'::uuid, '10000000-0000-0000-0000-000000000009'::uuid, 'CASE WHEN amount > 100 THEN "high" ELSE "low" END', 4, 2.3),
('20000000-0000-0000-0000-000000000045'::uuid, '10000000-0000-0000-0000-000000000009'::uuid, 'CREATE INDEX CONCURRENTLY idx_users_email ON users(email)', 5, 2.5)
ON CONFLICT (id) DO NOTHING;

-- ==========================================
-- DEVOPS THEME STAGES
-- ==========================================

INSERT INTO stages (id, name, theme_id, difficulty, is_active) VALUES
('10000000-0000-0000-0000-000000000010'::uuid, 'Docker Basics', '30000000-0000-0000-0000-000000000005'::uuid, 'easy', true),
('10000000-0000-0000-0000-000000000011'::uuid, 'Kubernetes', '30000000-0000-0000-0000-000000000005'::uuid, 'hard', true),
('10000000-0000-0000-0000-000000000012'::uuid, 'CI/CD Pipelines', '30000000-0000-0000-0000-000000000005'::uuid, 'medium', true)
ON CONFLICT (id) DO NOTHING;

-- Docker Basics phrases
INSERT INTO phrases (id, stage_id, text, sequence_number, base_multiplier) VALUES
('20000000-0000-0000-0000-000000000046'::uuid, '10000000-0000-0000-0000-000000000010'::uuid, 'docker build -t myapp:latest .', 1, 1.3),
('20000000-0000-0000-0000-000000000047'::uuid, '10000000-0000-0000-0000-000000000010'::uuid, 'docker run -d -p 8080:80 nginx', 2, 1.4),
('20000000-0000-0000-0000-000000000048'::uuid, '10000000-0000-0000-0000-000000000010'::uuid, 'FROM golang:1.21-alpine AS builder', 3, 1.5),
('20000000-0000-0000-0000-000000000049'::uuid, '10000000-0000-0000-0000-000000000010'::uuid, 'COPY --from=builder /app/main /app/', 4, 1.4),
('20000000-0000-0000-0000-000000000050'::uuid, '10000000-0000-0000-0000-000000000010'::uuid, 'docker-compose up -d --build', 5, 1.3)
ON CONFLICT (id) DO NOTHING;

-- Kubernetes phrases
INSERT INTO phrases (id, stage_id, text, sequence_number, base_multiplier) VALUES
('20000000-0000-0000-0000-000000000051'::uuid, '10000000-0000-0000-0000-000000000011'::uuid, 'kubectl apply -f deployment.yaml', 1, 1.8),
('20000000-0000-0000-0000-000000000052'::uuid, '10000000-0000-0000-0000-000000000011'::uuid, 'apiVersion: apps/v1 kind: Deployment', 2, 2.0),
('20000000-0000-0000-0000-000000000053'::uuid, '10000000-0000-0000-0000-000000000011'::uuid, 'spec: replicas: 3 selector: matchLabels:', 3, 2.2),
('20000000-0000-0000-0000-000000000054'::uuid, '10000000-0000-0000-0000-000000000011'::uuid, 'kubectl get pods -n production -o wide', 4, 1.9),
('20000000-0000-0000-0000-000000000055'::uuid, '10000000-0000-0000-0000-000000000011'::uuid, 'helm install myrelease charts/myapp --values values.yaml', 5, 2.5)
ON CONFLICT (id) DO NOTHING;

-- CI/CD Pipelines phrases
INSERT INTO phrases (id, stage_id, text, sequence_number, base_multiplier) VALUES
('20000000-0000-0000-0000-000000000056'::uuid, '10000000-0000-0000-0000-000000000012'::uuid, 'name: CI Pipeline on: [push, pull_request]', 1, 1.7),
('20000000-0000-0000-0000-000000000057'::uuid, '10000000-0000-0000-0000-000000000012'::uuid, 'runs-on: ubuntu-latest steps:', 2, 1.5),
('20000000-0000-0000-0000-000000000058'::uuid, '10000000-0000-0000-0000-000000000012'::uuid, 'uses: actions/checkout@v4', 3, 1.4),
('20000000-0000-0000-0000-000000000059'::uuid, '10000000-0000-0000-0000-000000000012'::uuid, 'run: npm ci && npm run build && npm test', 4, 1.8),
('20000000-0000-0000-0000-000000000060'::uuid, '10000000-0000-0000-0000-000000000012'::uuid, 'env: DATABASE_URL: ${{ secrets.DB_URL }}', 5, 1.9)
ON CONFLICT (id) DO NOTHING;

-- ==========================================
-- ALGORITHMS THEME STAGES
-- ==========================================

INSERT INTO stages (id, name, theme_id, difficulty, is_active) VALUES
('10000000-0000-0000-0000-000000000013'::uuid, 'Sorting Algorithms', '30000000-0000-0000-0000-000000000004'::uuid, 'medium', true),
('10000000-0000-0000-0000-000000000014'::uuid, 'Data Structures', '30000000-0000-0000-0000-000000000004'::uuid, 'hard', true)
ON CONFLICT (id) DO NOTHING;

-- Sorting Algorithms phrases
INSERT INTO phrases (id, stage_id, text, sequence_number, base_multiplier) VALUES
('20000000-0000-0000-0000-000000000061'::uuid, '10000000-0000-0000-0000-000000000013'::uuid, 'function quickSort(arr) {', 1, 1.6),
('20000000-0000-0000-0000-000000000062'::uuid, '10000000-0000-0000-0000-000000000013'::uuid, 'if (arr.length <= 1) return arr;', 2, 1.5),
('20000000-0000-0000-0000-000000000063'::uuid, '10000000-0000-0000-0000-000000000013'::uuid, 'const pivot = arr[Math.floor(arr.length / 2)];', 3, 2.0),
('20000000-0000-0000-0000-000000000064'::uuid, '10000000-0000-0000-0000-000000000013'::uuid, 'const left = arr.filter(x => x < pivot);', 4, 1.8),
('20000000-0000-0000-0000-000000000065'::uuid, '10000000-0000-0000-0000-000000000013'::uuid, 'return [...quickSort(left), pivot, ...quickSort(right)];', 5, 2.2)
ON CONFLICT (id) DO NOTHING;

-- Data Structures phrases
INSERT INTO phrases (id, stage_id, text, sequence_number, base_multiplier) VALUES
('20000000-0000-0000-0000-000000000066'::uuid, '10000000-0000-0000-0000-000000000014'::uuid, 'class Node { constructor(value) { this.value = value; } }', 1, 2.2),
('20000000-0000-0000-0000-000000000067'::uuid, '10000000-0000-0000-0000-000000000014'::uuid, 'this.left = null; this.right = null;', 2, 1.9),
('20000000-0000-0000-0000-000000000068'::uuid, '10000000-0000-0000-0000-000000000014'::uuid, 'insert(value) { if (value < this.value) {', 3, 2.0),
('20000000-0000-0000-0000-000000000069'::uuid, '10000000-0000-0000-0000-000000000014'::uuid, 'const stack = []; stack.push(element);', 4, 1.8),
('20000000-0000-0000-0000-000000000070'::uuid, '10000000-0000-0000-0000-000000000014'::uuid, 'while (!queue.isEmpty()) { node = queue.dequeue(); }', 5, 2.3)
ON CONFLICT (id) DO NOTHING;

-- ==========================================
-- GENERAL KNOWLEDGE THEME STAGES
-- ==========================================

INSERT INTO stages (id, name, theme_id, difficulty, is_active) VALUES
('10000000-0000-0000-0000-000000000015'::uuid, 'Common Phrases', '30000000-0000-0000-0000-000000000006'::uuid, 'easy', true),
('10000000-0000-0000-0000-000000000016'::uuid, 'Pangrams Collection', '30000000-0000-0000-0000-000000000006'::uuid, 'medium', true)
ON CONFLICT (id) DO NOTHING;

-- Common Phrases
INSERT INTO phrases (id, stage_id, text, sequence_number, base_multiplier) VALUES
('20000000-0000-0000-0000-000000000071'::uuid, '10000000-0000-0000-0000-000000000015'::uuid, 'The quick brown fox jumps over the lazy dog.', 1, 1.0),
('20000000-0000-0000-0000-000000000072'::uuid, '10000000-0000-0000-0000-000000000015'::uuid, 'Practice makes perfect.', 2, 0.8),
('20000000-0000-0000-0000-000000000073'::uuid, '10000000-0000-0000-0000-000000000015'::uuid, 'A journey of a thousand miles begins with a single step.', 3, 1.2),
('20000000-0000-0000-0000-000000000074'::uuid, '10000000-0000-0000-0000-000000000015'::uuid, 'Time flies like an arrow; fruit flies like a banana.', 4, 1.3),
('20000000-0000-0000-0000-000000000075'::uuid, '10000000-0000-0000-0000-000000000015'::uuid, 'To be or not to be, that is the question.', 5, 1.1)
ON CONFLICT (id) DO NOTHING;

-- Pangrams Collection
INSERT INTO phrases (id, stage_id, text, sequence_number, base_multiplier) VALUES
('20000000-0000-0000-0000-000000000076'::uuid, '10000000-0000-0000-0000-000000000016'::uuid, 'Pack my box with five dozen liquor jugs.', 1, 1.4),
('20000000-0000-0000-0000-000000000077'::uuid, '10000000-0000-0000-0000-000000000016'::uuid, 'How vexingly quick daft zebras jump!', 2, 1.5),
('20000000-0000-0000-0000-000000000078'::uuid, '10000000-0000-0000-0000-000000000016'::uuid, 'The five boxing wizards jump quickly.', 3, 1.3),
('20000000-0000-0000-0000-000000000079'::uuid, '10000000-0000-0000-0000-000000000016'::uuid, 'Sphinx of black quartz, judge my vow.', 4, 1.6),
('20000000-0000-0000-0000-000000000080'::uuid, '10000000-0000-0000-0000-000000000016'::uuid, 'Two driven jocks help fax my big quiz.', 5, 1.4)
ON CONFLICT (id) DO NOTHING;

-- ==========================================
-- BUSINESS THEME STAGES
-- ==========================================

INSERT INTO stages (id, name, theme_id, difficulty, is_active) VALUES
('10000000-0000-0000-0000-000000000017'::uuid, 'Email Writing', '30000000-0000-0000-0000-000000000007'::uuid, 'easy', true),
('10000000-0000-0000-0000-000000000018'::uuid, 'Business Terms', '30000000-0000-0000-0000-000000000007'::uuid, 'medium', true)
ON CONFLICT (id) DO NOTHING;

-- Email Writing phrases
INSERT INTO phrases (id, stage_id, text, sequence_number, base_multiplier) VALUES
('20000000-0000-0000-0000-000000000081'::uuid, '10000000-0000-0000-0000-000000000017'::uuid, 'Dear Sir or Madam, I hope this email finds you well.', 1, 1.2),
('20000000-0000-0000-0000-000000000082'::uuid, '10000000-0000-0000-0000-000000000017'::uuid, 'Please find attached the requested documents.', 2, 1.1),
('20000000-0000-0000-0000-000000000083'::uuid, '10000000-0000-0000-0000-000000000017'::uuid, 'I am writing to follow up on our previous conversation.', 3, 1.3),
('20000000-0000-0000-0000-000000000084'::uuid, '10000000-0000-0000-0000-000000000017'::uuid, 'Thank you for your prompt response to this matter.', 4, 1.2),
('20000000-0000-0000-0000-000000000085'::uuid, '10000000-0000-0000-0000-000000000017'::uuid, 'Please do not hesitate to contact me if you have any questions.', 5, 1.4)
ON CONFLICT (id) DO NOTHING;

-- Business Terms phrases
INSERT INTO phrases (id, stage_id, text, sequence_number, base_multiplier) VALUES
('20000000-0000-0000-0000-000000000086'::uuid, '10000000-0000-0000-0000-000000000018'::uuid, 'Return on investment (ROI) is a key performance indicator.', 1, 1.5),
('20000000-0000-0000-0000-000000000087'::uuid, '10000000-0000-0000-0000-000000000018'::uuid, 'The quarterly earnings report exceeded expectations.', 2, 1.4),
('20000000-0000-0000-0000-000000000088'::uuid, '10000000-0000-0000-0000-000000000018'::uuid, 'Our market share increased by 15% year-over-year.', 3, 1.5),
('20000000-0000-0000-0000-000000000089'::uuid, '10000000-0000-0000-0000-000000000018'::uuid, 'The stakeholders approved the proposed budget allocation.', 4, 1.6),
('20000000-0000-0000-0000-000000000090'::uuid, '10000000-0000-0000-0000-000000000018'::uuid, 'Synergy between departments drives organizational success.', 5, 1.7)
ON CONFLICT (id) DO NOTHING;

-- ==========================================
-- SCIENCE THEME STAGES
-- ==========================================

INSERT INTO stages (id, name, theme_id, difficulty, is_active) VALUES
('10000000-0000-0000-0000-000000000019'::uuid, 'Scientific Method', '30000000-0000-0000-0000-000000000008'::uuid, 'easy', true),
('10000000-0000-0000-0000-000000000020'::uuid, 'Chemistry Formulas', '30000000-0000-0000-0000-000000000008'::uuid, 'hard', true)
ON CONFLICT (id) DO NOTHING;

-- Scientific Method phrases
INSERT INTO phrases (id, stage_id, text, sequence_number, base_multiplier) VALUES
('20000000-0000-0000-0000-000000000091'::uuid, '10000000-0000-0000-0000-000000000019'::uuid, 'The hypothesis must be tested through experimentation.', 1, 1.3),
('20000000-0000-0000-0000-000000000092'::uuid, '10000000-0000-0000-0000-000000000019'::uuid, 'Data collection is crucial for scientific research.', 2, 1.2),
('20000000-0000-0000-0000-000000000093'::uuid, '10000000-0000-0000-0000-000000000019'::uuid, 'Peer review ensures the validity of research findings.', 3, 1.4),
('20000000-0000-0000-0000-000000000094'::uuid, '10000000-0000-0000-0000-000000000019'::uuid, 'The control group remained unchanged during the experiment.', 4, 1.5),
('20000000-0000-0000-0000-000000000095'::uuid, '10000000-0000-0000-0000-000000000019'::uuid, 'Results were statistically significant with p < 0.05.', 5, 1.6)
ON CONFLICT (id) DO NOTHING;

-- Chemistry Formulas phrases
INSERT INTO phrases (id, stage_id, text, sequence_number, base_multiplier) VALUES
('20000000-0000-0000-0000-000000000096'::uuid, '10000000-0000-0000-0000-000000000020'::uuid, 'H2O + CO2 -> H2CO3 (carbonic acid)', 1, 2.0),
('20000000-0000-0000-0000-000000000097'::uuid, '10000000-0000-0000-0000-000000000020'::uuid, '2NaOH + H2SO4 -> Na2SO4 + 2H2O', 2, 2.2),
('20000000-0000-0000-0000-000000000098'::uuid, '10000000-0000-0000-0000-000000000020'::uuid, 'CH4 + 2O2 -> CO2 + 2H2O (combustion)', 3, 2.1),
('20000000-0000-0000-0000-000000000099'::uuid, '10000000-0000-0000-0000-000000000020'::uuid, 'C6H12O6 -> 2C2H5OH + 2CO2 (fermentation)', 4, 2.4),
('20000000-0000-0000-0000-000000000100'::uuid, '10000000-0000-0000-0000-000000000020'::uuid, 'ATP -> ADP + Pi + Energy (cellular respiration)', 5, 2.3)
ON CONFLICT (id) DO NOTHING;

-- ==========================================
-- LITERATURE THEME STAGES
-- ==========================================

INSERT INTO stages (id, name, theme_id, difficulty, is_active) VALUES
('10000000-0000-0000-0000-000000000021'::uuid, 'Famous Quotes', '30000000-0000-0000-0000-000000000009'::uuid, 'easy', true),
('10000000-0000-0000-0000-000000000022'::uuid, 'Book Opening Lines', '30000000-0000-0000-0000-000000000009'::uuid, 'medium', true)
ON CONFLICT (id) DO NOTHING;

-- Famous Quotes phrases
INSERT INTO phrases (id, stage_id, text, sequence_number, base_multiplier) VALUES
('20000000-0000-0000-0000-000000000101'::uuid, '10000000-0000-0000-0000-000000000021'::uuid, 'Be the change you wish to see in the world.', 1, 1.1),
('20000000-0000-0000-0000-000000000102'::uuid, '10000000-0000-0000-0000-000000000021'::uuid, 'The only thing we have to fear is fear itself.', 2, 1.2),
('20000000-0000-0000-0000-000000000103'::uuid, '10000000-0000-0000-0000-000000000021'::uuid, 'In the middle of difficulty lies opportunity.', 3, 1.1),
('20000000-0000-0000-0000-000000000104'::uuid, '10000000-0000-0000-0000-000000000021'::uuid, 'Stay hungry, stay foolish.', 4, 0.9),
('20000000-0000-0000-0000-000000000105'::uuid, '10000000-0000-0000-0000-000000000021'::uuid, 'The unexamined life is not worth living.', 5, 1.2)
ON CONFLICT (id) DO NOTHING;

-- Book Opening Lines phrases
INSERT INTO phrases (id, stage_id, text, sequence_number, base_multiplier) VALUES
('20000000-0000-0000-0000-000000000106'::uuid, '10000000-0000-0000-0000-000000000022'::uuid, 'It was the best of times, it was the worst of times.', 1, 1.4),
('20000000-0000-0000-0000-000000000107'::uuid, '10000000-0000-0000-0000-000000000022'::uuid, 'Call me Ishmael.', 2, 0.8),
('20000000-0000-0000-0000-000000000108'::uuid, '10000000-0000-0000-0000-000000000022'::uuid, 'All happy families are alike; each unhappy family is unhappy in its own way.', 3, 1.8),
('20000000-0000-0000-0000-000000000109'::uuid, '10000000-0000-0000-0000-000000000022'::uuid, 'It is a truth universally acknowledged that a single man must be in want of a wife.', 4, 2.0),
('20000000-0000-0000-0000-000000000110'::uuid, '10000000-0000-0000-0000-000000000022'::uuid, 'In a hole in the ground there lived a hobbit.', 5, 1.3)
ON CONFLICT (id) DO NOTHING;
