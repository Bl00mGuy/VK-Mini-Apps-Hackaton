// src/api.js
const API_BASE_URL = 'https://example.com/api'; // Замените на ваш бэкенд

export const fetchTasks = async () => {
    const response = await fetch(`${API_BASE_URL}/tasks`);
    if (!response.ok) {
        throw new Error('Ошибка при получении задач');
    }
    return response.json();
};

export const completeTask = async (taskId) => {
    const response = await fetch(`${API_BASE_URL}/tasks/${taskId}/complete`, {
        method: 'POST',
    });
    if (!response.ok) {
        throw new Error('Ошибка при выполнении задачи');
    }
};

export const fetchAchievements = async () => {
    const response = await fetch(`${API_BASE_URL}/achievements`);
    if (!response.ok) {
        throw new Error('Ошибка при получении достижений');
    }
    return response.json();
};

export const fetchLeaderboard = async () => {
    const response = await fetch(`${API_BASE_URL}/leaderboard`);
    if (!response.ok) {
        throw new Error('Ошибка при получении лидерборда');
    }
    return response.json();
};

export const fetchStatistics = async () => {
    const response = await fetch(`${API_BASE_URL}/statistics`);
    if (!response.ok) {
        throw new Error('Ошибка при получении статистики');
    }
    return response.json();
};
