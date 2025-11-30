// API Configuration
const API_URL = window.location.hostname === 'localhost' 
    ? 'http://localhost:8080' 
    : window.location.origin;

let authToken = localStorage.getItem('authToken');
let stages = [];
let phrases = [];

// Utility Functions
function showMessage(message, isError = false) {
    const messageDiv = document.getElementById('message');
    messageDiv.textContent = message;
    messageDiv.className = isError ? 'error' : 'success';
    messageDiv.classList.remove('hidden');
    setTimeout(() => {
        messageDiv.classList.add('hidden');
    }, 5000);
}

function showTab(tabName) {
    // Hide all tabs
    document.querySelectorAll('.tab-content').forEach(tab => {
        tab.classList.remove('active');
    });
    document.querySelectorAll('.tab').forEach(tab => {
        tab.classList.remove('active');
    });

    // Show selected tab
    document.getElementById(tabName + 'Tab').classList.add('active');
    event.target.classList.add('active');

    // Load data for the selected tab
    if (tabName === 'stages') {
        loadStages();
    } else if (tabName === 'phrases') {
        loadStagesForDropdown();
        loadPhrases();
    }
}

// Authentication
document.getElementById('loginForm').addEventListener('submit', async (e) => {
    e.preventDefault();
    
    const username = document.getElementById('username').value;
    const password = document.getElementById('password').value;

    try {
        const response = await fetch(`${API_URL}/api/auth/login`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ username, password }),
        });

        const data = await response.json();

        if (!response.ok) {
            throw new Error(data.error || 'Login failed');
        }

        // Check if user is admin
        const profileResponse = await fetch(`${API_URL}/api/auth/profile`, {
            headers: {
                'Authorization': `Bearer ${data.access_token}`,
            },
        });

        const profile = await profileResponse.json();

        if (profile.role !== 'admin') {
            throw new Error('Admin access required');
        }

        authToken = data.access_token;
        localStorage.setItem('authToken', authToken);

        document.getElementById('loginSection').classList.add('hidden');
        document.getElementById('mainContent').classList.remove('hidden');
        
        loadStages();
    } catch (error) {
        const errorDiv = document.getElementById('loginError');
        errorDiv.textContent = error.message;
        errorDiv.classList.remove('hidden');
    }
});

function logout() {
    authToken = null;
    localStorage.removeItem('authToken');
    document.getElementById('loginSection').classList.remove('hidden');
    document.getElementById('mainContent').classList.add('hidden');
}

// API Helper
async function apiRequest(url, options = {}) {
    const headers = {
        'Content-Type': 'application/json',
        ...options.headers,
    };

    if (authToken) {
        headers['Authorization'] = `Bearer ${authToken}`;
    }

    const response = await fetch(API_URL + url, {
        ...options,
        headers,
    });

    if (response.status === 401) {
        logout();
        throw new Error('Session expired. Please login again.');
    }

    const data = await response.json();

    if (!response.ok) {
        throw new Error(data.error || 'Request failed');
    }

    return data;
}

// Stages Management
async function loadStages() {
    try {
        stages = await apiRequest('/admin/stages');
        const themes = await apiRequest('/admin/themes');
        
        // Map theme names to stages
        stages = stages.map(stage => {
            const theme = themes.find(t => t.id === stage.theme_id);
            return {
                ...stage,
                theme_name: theme ? theme.name : 'Unknown'
            };
        });
        
        renderStages();
    } catch (error) {
        showMessage('Error loading stages: ' + error.message, true);
    }
}

function renderStages() {
    const tbody = document.getElementById('stagesTableBody');
    
    if (stages.length === 0) {
        tbody.innerHTML = '<tr><td colspan="5">No stages found</td></tr>';
        return;
    }

    tbody.innerHTML = stages.map(stage => `
        <tr>
            <td>${stage.name}</td>
            <td>${stage.theme_name || stage.theme_id}</td>
            <td><span class="badge badge-${stage.difficulty}">${stage.difficulty}</span></td>
            <td><span class="badge ${stage.is_active ? 'badge-success' : 'badge-danger'}">${stage.is_active ? 'Active' : 'Inactive'}</span></td>
            <td class="action-buttons">
                <button class="btn btn-small" onclick="editStage('${stage.id}')">Edit</button>
                <button class="btn btn-small btn-danger" onclick="deleteStage('${stage.id}')">Delete</button>
            </td>
        </tr>
    `).join('');
}

let editingStageId = null;

document.getElementById('stageForm').addEventListener('submit', async (e) => {
    e.preventDefault();

    const stageData = {
        name: document.getElementById('stageName').value,
        theme_id: document.getElementById('stageTheme').value,
        difficulty: document.getElementById('stageDifficulty').value,
        is_active: document.getElementById('stageIsActive').checked,
    };

    try {
        if (editingStageId) {
            // Update existing stage
            await apiRequest(`/admin/stage/${editingStageId}`, {
                method: 'PUT',
                body: JSON.stringify(stageData),
            });
            showMessage('Stage updated successfully!');
            editingStageId = null;
            document.getElementById('stageFormTitle').textContent = 'Create New Stage';
            document.querySelector('#stageForm button[type="submit"]').textContent = 'Create Stage';
        } else {
            // Create new stage
            await apiRequest('/admin/stage', {
                method: 'POST',
                body: JSON.stringify(stageData),
            });
            showMessage('Stage created successfully!');
        }

        document.getElementById('stageForm').reset();
        loadStages();
        loadStagesForDropdown();
    } catch (error) {
        showMessage('Error saving stage: ' + error.message, true);
    }
});

function cancelEdit() {
    editingStageId = null;
    document.getElementById('stageForm').reset();
    document.getElementById('stageFormTitle').textContent = 'Create New Stage';
    document.querySelector('#stageForm button[type="submit"]').textContent = 'Create Stage';
    document.getElementById('cancelEditBtn').style.display = 'none';
}

function editStage(stageId) {
    const stage = stages.find(s => s.id === stageId);
    if (!stage) {
        showMessage('Stage not found', true);
        return;
    }

    // Populate form with stage data
    document.getElementById('stageName').value = stage.name;
    document.getElementById('stageTheme').value = stage.theme_id;
    document.getElementById('stageDifficulty').value = stage.difficulty;
    document.getElementById('stageIsActive').checked = stage.is_active;

    // Update form UI
    editingStageId = stageId;
    document.getElementById('stageFormTitle').textContent = 'Edit Stage';
    document.querySelector('#stageForm button[type="submit"]').textContent = 'Update Stage';
    document.getElementById('cancelEditBtn').style.display = 'inline-block';

    // Scroll to form
    document.getElementById('stageForm').scrollIntoView({ behavior: 'smooth' });
}

async function deleteStage(stageId) {
    if (!confirm('Are you sure you want to delete this stage?')) {
        return;
    }

    try {
        await apiRequest(`/admin/stage/${stageId}`, {
            method: 'DELETE',
        });

        showMessage('Stage deleted successfully!');
        loadStages();
    } catch (error) {
        showMessage('Error deleting stage: ' + error.message, true);
    }
}

// Phrases Management
async function loadStagesForDropdown() {
    try {
        const result = await apiRequest('/admin/stages');
        stages = result || [];
        
        const phraseStageSelect = document.getElementById('phraseStageId');
        const filterStageSelect = document.getElementById('filterStageId');
        
        if (stages.length === 0) {
            phraseStageSelect.innerHTML = '<option value="">No stages available</option>';
            filterStageSelect.innerHTML = '<option value="">No stages available</option>';
            return;
        }
        
        const options = stages.map(stage => 
            `<option value="${stage.id}">${stage.name} (${stage.difficulty})</option>`
        ).join('');

        phraseStageSelect.innerHTML = '<option value="">Select a stage</option>' + options;
        filterStageSelect.innerHTML = '<option value="">All Stages</option>' + options;
    } catch (error) {
        console.error('Error loading stages for dropdown:', error);
        showMessage('Error loading stages: ' + error.message, true);
    }
}

async function loadPhrases() {
    const stageId = document.getElementById('filterStageId')?.value || '';
    
    try {
        if (stageId) {
            const result = await apiRequest(`/admin/phrases?stage_id=${stageId}`);
            phrases = result || [];
        } else {
            // Load all phrases from all stages
            const allStages = await apiRequest('/admin/stages');
            phrases = [];
            
            if (allStages && Array.isArray(allStages)) {
                for (const stage of allStages) {
                    const stagePhrases = await apiRequest(`/admin/phrases?stage_id=${stage.id}`);
                    if (stagePhrases && Array.isArray(stagePhrases)) {
                        phrases.push(...stagePhrases.map(p => ({ ...p, stageName: stage.name })));
                    }
                }
            }
        }
        renderPhrases();
    } catch (error) {
        console.error('Error loading phrases:', error);
        showMessage('Error loading phrases: ' + error.message, true);
        phrases = [];
        renderPhrases();
    }
}

function renderPhrases() {
    const tbody = document.getElementById('phrasesTableBody');
    
    if (phrases.length === 0) {
        tbody.innerHTML = '<tr><td colspan="6">No phrases found</td></tr>';
        return;
    }

    tbody.innerHTML = phrases.map(phrase => {
        const stage = stages.find(s => s.id === phrase.stage_id);
        const stageName = stage ? stage.name : (phrase.stageName || 'Unknown');
        
        return `
            <tr>
                <td>${stageName}</td>
                <td>${phrase.text}</td>
                <td>${phrase.sequence_number}</td>
                <td>${phrase.multiplier || phrase.base_multiplier}</td>
                <td class="action-buttons">
                    <button class="btn btn-small" onclick="editPhrase('${phrase.id}')">Edit</button>
                    <button class="btn btn-small btn-danger" onclick="deletePhrase('${phrase.id}')">Delete</button>
                </td>
            </tr>
        `;
    }).join('');
}

let editingPhraseId = null;

document.getElementById('phraseForm').addEventListener('submit', async (e) => {
    e.preventDefault();

    const phraseData = {
        stage_id: document.getElementById('phraseStageId').value,
        text: document.getElementById('phraseText').value,
        sequence_number: parseInt(document.getElementById('phraseSequence').value),
        base_multiplier: parseFloat(document.getElementById('phraseMultiplier').value),
    };

    try {
        if (editingPhraseId) {
            // Update existing phrase
            await apiRequest(`/admin/phrase/${editingPhraseId}`, {
                method: 'PUT',
                body: JSON.stringify(phraseData),
            });
            showMessage('Phrase updated successfully!');
            editingPhraseId = null;
            document.getElementById('phraseFormTitle').textContent = 'Create New Phrase';
            document.querySelector('#phraseForm button[type="submit"]').textContent = 'Create Phrase';
            document.getElementById('cancelPhraseEditBtn').style.display = 'none';
        } else {
            // Create new phrase
            await apiRequest('/admin/phrase', {
                method: 'POST',
                body: JSON.stringify(phraseData),
            });
            showMessage('Phrase created successfully!');
        }

        document.getElementById('phraseForm').reset();
        loadPhrases();
    } catch (error) {
        showMessage('Error saving phrase: ' + error.message, true);
    }
});

function editPhrase(phraseId) {
    const phrase = phrases.find(p => p.id === phraseId);
    if (!phrase) {
        showMessage('Phrase not found', true);
        return;
    }

    // Populate form with phrase data
    document.getElementById('phraseStageId').value = phrase.stage_id;
    document.getElementById('phraseText').value = phrase.text;
    document.getElementById('phraseSequence').value = phrase.sequence_number;
    document.getElementById('phraseMultiplier').value = phrase.multiplier || phrase.base_multiplier;

    // Update form UI
    editingPhraseId = phraseId;
    document.getElementById('phraseFormTitle').textContent = 'Edit Phrase';
    document.querySelector('#phraseForm button[type="submit"]').textContent = 'Update Phrase';
    document.getElementById('cancelPhraseEditBtn').style.display = 'inline-block';

    // Switch to phrases tab if not already there
    showTab('phrases');

    // Scroll to form
    document.getElementById('phraseForm').scrollIntoView({ behavior: 'smooth' });
}

function cancelPhraseEdit() {
    editingPhraseId = null;
    document.getElementById('phraseForm').reset();
    document.getElementById('phraseFormTitle').textContent = 'Create New Phrase';
    document.querySelector('#phraseForm button[type="submit"]').textContent = 'Create Phrase';
    document.getElementById('cancelPhraseEditBtn').style.display = 'none';
}

async function deletePhrase(phraseId) {
    if (!confirm('Are you sure you want to delete this phrase?')) {
        return;
    }

    try {
        await apiRequest(`/admin/phrase/${phraseId}`, {
            method: 'DELETE',
        });

        showMessage('Phrase deleted successfully!');
        loadPhrases();
    } catch (error) {
        showMessage('Error deleting phrase: ' + error.message, true);
    }
}

// Initialize
if (authToken) {
    // Verify token is still valid
    apiRequest('/api/auth/profile')
        .then(profile => {
            if (profile.role === 'admin') {
                document.getElementById('loginSection').classList.add('hidden');
                document.getElementById('mainContent').classList.remove('hidden');
                loadStages();
                loadThemes(); // Load themes on init
            } else {
                logout();
            }
        })
        .catch(() => {
            logout();
        });
}

// Load themes
async function loadThemes() {
    try {
        const themes = await apiRequest('/admin/themes');
        const themeSelect = document.getElementById('stageTheme');
        
        themeSelect.innerHTML = '<option value="">Select a theme</option>' + 
            themes.map(theme => `<option value="${theme.id}">${theme.name}</option>`).join('');
    } catch (error) {
        console.error('Error loading themes:', error);
    }
}

