import React, { useState, useEffect } from 'react';
import { Panel, PanelHeader, Group, Cell, Button, ScreenSpinner } from '@vkontakte/vkui';
import PropTypes from 'prop-types';
import { fetchTasks, completeTask } from '../api';

const TasksPanel = ({ id, goTo }) => {
    const [tasks, setTasks] = useState([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);

    useEffect(() => {
        const getTasks = async () => {
            try {
                const fetchedTasks = await fetchTasks();
                setTasks(fetchedTasks);
            } catch (err) {
                setError('Ошибка при загрузке задач');
            } finally {
                setLoading(false);
            }
        };
        getTasks();
    }, []);

    const handleCompleteTask = async (taskId) => {
        await completeTask(taskId);
        setTasks(tasks.map(task =>
            task.id === taskId ? { ...task, completed: true } : task
        ));
    };

    if (loading) return <ScreenSpinner size="large" />;
    if (error) return <div>{error}</div>;

    return (
        <Panel id={id}>
            <PanelHeader>Задания</PanelHeader>
            <Group>
                {tasks.map((task) => (
                    <Cell key={task.id} description={task.description} asideContent={
                        <Button size="m" onClick={() => handleCompleteTask(task.id)} disabled={task.completed}>
                            {task.completed ? 'Выполнено' : 'Выполнить'}
                        </Button>
                    }>
                        {task.title}
                    </Cell>
                ))}
            </Group>
        </Panel>
    );
};

TasksPanel.propTypes = {
    id: PropTypes.string.isRequired,
    goTo: PropTypes.func.isRequired,
};

export { TasksPanel };
