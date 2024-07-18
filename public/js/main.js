document.addEventListener('DOMContentLoaded', () => {
    const taskForm = document.getElementById('task-form');
    const taskList = document.getElementById('task-list');

    document.getElementById('add-task').addEventListener('click', addTask);

    function addTask() {
        const title = document.getElementById('task-title').value;
        const description = document.getElementById('task-description').value;
        const dueDate = document.getElementById('task-due-date').value;
        const category = document.getElementById('task-category').value;

        const task = {
            title,
            description,
            due_date: dueDate,
            category,
            is_completed: false
        };

        fetch('/tasks', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(task)
        })
            .then(response => response.json())
            .then(data => {
                if (data.error) {
                    alert('Error: ' + data.error);
                } else {
                    renderTask(data);
                    clearForm();
                }
            })
            .catch(error => console.error('Error:', error));
    }

    function renderTask(task) {
        const taskDiv = document.createElement('div');
        taskDiv.classList.add('task');
        taskDiv.innerHTML = `
            <h3>${task.title}</h3>
            <p>${task.description}</p>
            <p>Due: ${task.due_date}</p>
            <p>Category: ${task.category}</p>
            <button onclick="markAsComplete(${task.id})">Complete</button>
            <button onclick="deleteTask(${task.id})">Delete</button>
        `;
        taskList.appendChild(taskDiv);
    }

    function clearForm() {
        document.getElementById('task-title').value = '';
        document.getElementById('task-description').value = '';
        document.getElementById('task-due-date').value = '';
        document.getElementById('task-category').value = 'Work';
    }

    window.markAsComplete = function(taskId) {
        fetch(`/tasks/${taskId}`, {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ is_completed: true })
        })
            .then(response => response.json())
            .then(data => {
                if (data.error) {
                    alert('Error: ' + data.error);
                } else {
                    alert('Task marked as complete!');
                }
            })
            .catch(error => console.error('Error:', error));
    }

    window.deleteTask = function(taskId) {
        fetch(`/tasks/${taskId}`, {
            method: 'DELETE'
        })
            .then(response => response.json())
            .then(data => {
                if (data.error) {
                    alert('Error: ' + data.error);
                } else {
                    alert('Task deleted!');
                    location.reload();
                }
            })
            .catch(error => console.error('Error:', error));
    }
});
