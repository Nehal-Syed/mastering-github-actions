const API_URL = 'http://localhost:8080/api';

async function fetchUsers() {
    try {
        const response = await fetch(`${API_URL}/users`);
        const users = await response.json();
        displayUsers(users);
    } catch (error) {
        console.error('Error fetching users:', error);
    }
}

function displayUsers(users) {
    const usersDiv = document.getElementById('users');
    if (users.length === 0) {
        usersDiv.innerHTML = '<p>No users found. Add your first user!</p>';
        return;
    }

    usersDiv.innerHTML = users.map(user => `
        <div class="user-card" data-id="${user.id}">
            <div class="user-info">
                <h3>${user.name}</h3>
                <p>Email: ${user.email}</p>
                <p>Age: ${user.age}</p>
                <p>Created: ${new Date(user.created_at).toLocaleDateString()}</p>
            </div>
            <div class="user-actions">
                <button class="edit-btn" onclick="editUser(${user.id})">Edit</button>
                <button class="delete-btn" onclick="deleteUser(${user.id})">Delete</button>
            </div>
        </div>
    `).join('');
}

async function createUser() {
    const name = document.getElementById('name').value;
    const email = document.getElementById('email').value;
    const age = parseInt(document.getElementById('age').value);

    if (!name || !email || !age) {
        alert('Please fill all fields');
        return;
    }

    try {
        const response = await fetch(`${API_URL}/users`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ name, email, age })
        });

        if (response.ok) {
            document.getElementById('name').value = '';
            document.getElementById('email').value = '';
            document.getElementById('age').value = '';
            fetchUsers();
        }
    } catch (error) {
        console.error('Error creating user:', error);
    }
}

async function deleteUser(id) {
    if (confirm('Are you sure you want to delete this user?')) {
        try {
            const response = await fetch(`${API_URL}/users/${id}`, {
                method: 'DELETE'
            });
            if (response.ok) {
                fetchUsers();
            }
        } catch (error) {
            console.error('Error deleting user:', error);
        }
    }
}

async function editUser(id) {
    const newName = prompt('Enter new name:');
    const newEmail = prompt('Enter new email:');
    const newAge = parseInt(prompt('Enter new age:'));

    if (newName && newEmail && newAge) {
        try {
            const response = await fetch(`${API_URL}/users/${id}`, {
                method: 'PUT',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ name: newName, email: newEmail, age: newAge })
            });
            if (response.ok) {
                fetchUsers();
            }
        } catch (error) {
            console.error('Error updating user:', error);
        }
    }
}

// Load users on page load
fetchUsers();